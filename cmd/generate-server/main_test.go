package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// dummy handler used to wrap with middleware
func okHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

// ─── API Key Middleware ───

func TestRequireAPIKey_ValidHeader(t *testing.T) {
	h := requireAPIKey(http.HandlerFunc(okHandler), "secret123")
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("X-API-Key", "secret123")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("valid API key: status = %d, want 200", w.Code)
	}
}

func TestRequireAPIKey_InvalidHeader(t *testing.T) {
	h := requireAPIKey(http.HandlerFunc(okHandler), "secret123")
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("X-API-Key", "wrongkey")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("invalid API key: status = %d, want 401", w.Code)
	}
}

func TestRequireAPIKey_MissingKey(t *testing.T) {
	h := requireAPIKey(http.HandlerFunc(okHandler), "secret123")
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusUnauthorized {
		t.Errorf("missing API key: status = %d, want 401", w.Code)
	}
}

func TestRequireAPIKey_QueryParam(t *testing.T) {
	h := requireAPIKey(http.HandlerFunc(okHandler), "secret123")
	req := httptest.NewRequest("GET", "/test?api_key=secret123", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("query param API key: status = %d, want 200", w.Code)
	}
}

// ─── Rate Limiter ───

func TestRateLimiter_AllowUpToLimit(t *testing.T) {
	limiter := newIPRateLimiter(5, time.Minute)
	defer limiter.Stop()

	for i := 0; i < 5; i++ {
		if !limiter.allow("192.168.1.1") {
			t.Fatalf("request %d should be allowed", i+1)
		}
	}
}

func TestRateLimiter_RejectAfterLimit(t *testing.T) {
	limiter := newIPRateLimiter(3, time.Minute)
	defer limiter.Stop()

	for i := 0; i < 3; i++ {
		limiter.allow("10.0.0.1")
	}
	if limiter.allow("10.0.0.1") {
		t.Error("request after limit should be rejected")
	}
}

func TestRateLimiter_SeparateBuckets(t *testing.T) {
	limiter := newIPRateLimiter(2, time.Minute)
	defer limiter.Stop()

	// Exhaust IP A
	limiter.allow("10.0.0.1")
	limiter.allow("10.0.0.1")
	if limiter.allow("10.0.0.1") {
		t.Error("IP A should be rate limited")
	}

	// IP B should still be allowed
	if !limiter.allow("10.0.0.2") {
		t.Error("IP B should not be rate limited")
	}
}

func TestRateMiddleware_Returns429(t *testing.T) {
	limiter := newIPRateLimiter(1, time.Minute)
	defer limiter.Stop()

	h := rateMiddleware(http.HandlerFunc(okHandler), limiter, false)

	// First request - allowed
	req1 := httptest.NewRequest("GET", "/test", nil)
	req1.RemoteAddr = "1.2.3.4:1234"
	w1 := httptest.NewRecorder()
	h.ServeHTTP(w1, req1)
	if w1.Code != http.StatusOK {
		t.Errorf("first request: status = %d, want 200", w1.Code)
	}

	// Second request - should be rate limited
	req2 := httptest.NewRequest("GET", "/test", nil)
	req2.RemoteAddr = "1.2.3.4:1234"
	w2 := httptest.NewRecorder()
	h.ServeHTTP(w2, req2)
	if w2.Code != http.StatusTooManyRequests {
		t.Errorf("second request: status = %d, want 429", w2.Code)
	}
	if ra := w2.Header().Get("Retry-After"); ra != "60" {
		t.Errorf("Retry-After = %q, want %q", ra, "60")
	}
}

func TestRateMiddleware_TrustProxy(t *testing.T) {
	limiter := newIPRateLimiter(1, time.Minute)
	defer limiter.Stop()

	h := rateMiddleware(http.HandlerFunc(okHandler), limiter, true)

	// Request with X-Forwarded-For
	req1 := httptest.NewRequest("GET", "/test", nil)
	req1.RemoteAddr = "proxy:8080"
	req1.Header.Set("X-Forwarded-For", "5.6.7.8, proxy")
	w1 := httptest.NewRecorder()
	h.ServeHTTP(w1, req1)
	if w1.Code != http.StatusOK {
		t.Errorf("first forwarded request: status = %d, want 200", w1.Code)
	}

	// Second request from same forwarded IP
	req2 := httptest.NewRequest("GET", "/test", nil)
	req2.RemoteAddr = "proxy:8080"
	req2.Header.Set("X-Forwarded-For", "5.6.7.8, proxy")
	w2 := httptest.NewRecorder()
	h.ServeHTTP(w2, req2)
	if w2.Code != http.StatusTooManyRequests {
		t.Errorf("second forwarded request: status = %d, want 429", w2.Code)
	}
}

// ─── CORS Middleware ───

func TestCORS_WildcardAllowsAnyOrigin(t *testing.T) {
	h := cors(http.HandlerFunc(okHandler), "*")
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "https://example.com")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "*" {
		t.Errorf("ACAO = %q, want %q", got, "*")
	}
}

func TestCORS_SpecificOriginAllowed(t *testing.T) {
	h := cors(http.HandlerFunc(okHandler), "https://example.com,https://other.com")
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "https://example.com")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
	if got := w.Header().Get("Access-Control-Allow-Origin"); got != "https://example.com" {
		t.Errorf("ACAO = %q, want %q", got, "https://example.com")
	}
	if got := w.Header().Get("Vary"); got != "Origin" {
		t.Errorf("Vary = %q, want %q", got, "Origin")
	}
}

func TestCORS_OriginRejected(t *testing.T) {
	h := cors(http.HandlerFunc(okHandler), "https://allowed.com")
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "https://evil.com")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusForbidden {
		t.Errorf("status = %d, want 403", w.Code)
	}
}

func TestCORS_NoOriginPassesThrough(t *testing.T) {
	h := cors(http.HandlerFunc(okHandler), "https://allowed.com")
	req := httptest.NewRequest("GET", "/test", nil)
	// No Origin header (non-browser request)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("status = %d, want 200", w.Code)
	}
}

func TestCORS_EmptyConfigBlocksCrossOrigin(t *testing.T) {
	h := cors(http.HandlerFunc(okHandler), "")
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Set("Origin", "https://any.com")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusForbidden {
		t.Errorf("empty config should block cross-origin: status = %d, want 403", w.Code)
	}
}

func TestCORS_PreflightBypassesAuth(t *testing.T) {
	inner := requireAPIKey(http.HandlerFunc(okHandler), "secret")
	h := cors(inner, "*")
	req := httptest.NewRequest("OPTIONS", "/test", nil)
	req.Header.Set("Origin", "https://example.com")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("OPTIONS with auth should return 204, got %d", w.Code)
	}
}

func TestCORS_OptionsReturnsNoContent(t *testing.T) {
	h := cors(http.HandlerFunc(okHandler), "*")
	req := httptest.NewRequest("OPTIONS", "/test", nil)
	req.Header.Set("Origin", "https://example.com")
	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Errorf("OPTIONS status = %d, want 204", w.Code)
	}
}
