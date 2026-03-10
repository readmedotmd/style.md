// Command generate-server runs an HTTP server for SVG banner and icon
// generation.
//
// Endpoints:
//
//	GET/POST /banner  — generate a banner SVG
//	GET/POST /icon    — generate an icon SVG
//	GET      /icons   — list available remix icon names
//	GET      /health  — health check
//
// Environment variables:
//
//	PORT             — listen port (default 8080)
//	ALLOWED_ORIGINS  — comma-separated CORS origins (default: none, blocks cross-origin)
//	API_KEY          — optional API key; if set, requests must include
//	                   X-API-Key header (preferred) or ?api_key= query param
//	RATE_LIMIT       — requests per IP per minute (default 30)
//	TRUST_PROXY      — set to "true" or "1" to trust X-Forwarded-For header
package main

import (
	"context"
	"crypto/subtle"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/readmedotmd/style.md/generate"
)

func main() {
	port := envOr("PORT", "8080")
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	apiKey := os.Getenv("API_KEY")
	rateLimit := envInt("RATE_LIMIT", 30)

	trustProxy := os.Getenv("TRUST_PROXY")
	trustFwd := trustProxy == "true" || trustProxy == "1"

	limiter := newIPRateLimiter(rateLimit, time.Minute)
	defer limiter.Stop()

	wrap := func(next http.HandlerFunc) http.HandlerFunc {
		h := next
		if apiKey != "" {
			h = requireAPIKey(h, apiKey)
		}
		h = rateMiddleware(h, limiter, trustFwd)
		h = cors(h, allowedOrigins)
		h = logMiddleware(h)
		return h
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/banner", wrap(generate.HandleBanner))
	mux.HandleFunc("/icon", wrap(generate.HandleIcon))
	mux.HandleFunc("/icons", wrap(generate.HandleIcons))
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 16, // 64KB
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		slog.Info("server starting", "port", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
	slog.Info("shutting down")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("shutdown error", "err", err)
	}
}

// ─── Middleware ───

func logMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		slog.Info("request", "method", r.Method, "path", r.URL.Path, "remote", r.RemoteAddr, "duration", time.Since(start).Round(time.Millisecond))
	}
}

func cors(next http.HandlerFunc, allowedOrigins string) http.HandlerFunc {
	var origins []string
	if allowedOrigins != "" {
		origins = strings.Split(allowedOrigins, ",")
		for i := range origins {
			origins[i] = strings.TrimSpace(origins[i])
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Origin")
		origin := r.Header.Get("Origin")
		allowed := false
		for _, o := range origins {
			if o == "*" || o == origin {
				allowed = true
				w.Header().Set("Access-Control-Allow-Origin", o)
				break
			}
		}
		if !allowed && origin != "" {
			http.Error(w, `{"error":"origin not allowed"}`, http.StatusForbidden)
			return
		}
		if origin == "" {
			for _, o := range origins {
				if o == "*" {
					w.Header().Set("Access-Control-Allow-Origin", "*")
					allowed = true
					break
				}
			}
		}
		if allowed {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key")
		}
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func requireAPIKey(next http.HandlerFunc, key string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		k := r.Header.Get("X-API-Key")
		if k == "" {
			k = r.URL.Query().Get("api_key")
		}
		if subtle.ConstantTimeCompare([]byte(k), []byte(key)) != 1 {
			http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func rateMiddleware(next http.HandlerFunc, limiter *ipRateLimiter, trustFwd bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		if host, _, err := net.SplitHostPort(ip); err == nil {
			ip = host
		}
		if trustFwd {
			if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
				ip = strings.SplitN(fwd, ",", 2)[0]
				ip = strings.TrimSpace(ip)
			}
		}
		if !limiter.allow(ip) {
			w.Header().Set("Retry-After", "60")
			http.Error(w, `{"error":"rate limit exceeded"}`, http.StatusTooManyRequests)
			return
		}
		next(w, r)
	}
}

// ─── Rate Limiter ───

type ipRateLimiter struct {
	mu     sync.Mutex
	counts map[string]*bucket
	limit  int
	window time.Duration
	stop   chan struct{}
}

type bucket struct {
	count  int
	expiry time.Time
}

func newIPRateLimiter(limit int, window time.Duration) *ipRateLimiter {
	rl := &ipRateLimiter{
		counts: make(map[string]*bucket),
		limit:  limit,
		window: window,
		stop:   make(chan struct{}),
	}
	go rl.cleanup()
	return rl
}

func (rl *ipRateLimiter) allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	now := time.Now()
	b, ok := rl.counts[ip]
	if !ok || now.After(b.expiry) {
		rl.counts[ip] = &bucket{count: 1, expiry: now.Add(rl.window)}
		return true
	}
	b.count++
	return b.count <= rl.limit
}

func (rl *ipRateLimiter) cleanup() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			rl.mu.Lock()
			now := time.Now()
			for ip, b := range rl.counts {
				if now.After(b.expiry) {
					delete(rl.counts, ip)
				}
			}
			rl.mu.Unlock()
		case <-rl.stop:
			return
		}
	}
}

// Stop terminates the background cleanup goroutine.
func (rl *ipRateLimiter) Stop() {
	close(rl.stop)
}

// ─── Helpers ───

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func envInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def
	}
	return n
}
