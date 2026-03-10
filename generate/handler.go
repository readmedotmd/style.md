package generate

import (
	"encoding/json"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// MaxRequestBodyBytes is the maximum allowed size for POST request bodies (64KB).
const MaxRequestBodyBytes = 64 * 1024

// securityHeaders sets standard security response headers.
func securityHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("Content-Security-Policy", "default-src 'none'")
}

// HandleBanner is an HTTP handler that generates banner SVGs.
//
//	POST /banner — JSON body with BannerRequest fields
//	GET  /banner — query params: name, tagline, desc, version, accent, layout, tags (comma-sep), size
func HandleBanner(w http.ResponseWriter, r *http.Request) {
	securityHeaders(w)

	if r.Method != http.MethodGet && r.Method != http.MethodPost && r.Method != http.MethodOptions {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req BannerRequest

	if r.Method == http.MethodPost {
		body := http.MaxBytesReader(w, r.Body, MaxRequestBodyBytes)
		defer body.Close()
		if err := json.NewDecoder(body).Decode(&req); err != nil {
			http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
			return
		}
	} else {
		q := r.URL.Query()
		req.Name = q.Get("name")
		req.Tagline = q.Get("tagline")
		req.Desc = q.Get("desc")
		req.Version = q.Get("version")
		req.Accent = q.Get("accent")
		req.Layout = q.Get("layout")
		req.Size = q.Get("size")
		if tags := q.Get("tags"); tags != "" {
			for _, t := range splitTags(tags) {
				if t != "" {
					req.Tags = append(req.Tags, t)
				}
			}
		}
		req.ShowDots = queryBool(q, "show_dots")
		req.ShowBadge = queryBool(q, "show_badge")
		req.ShowTags = queryBool(q, "show_tags")
	}

	req.Sanitize()

	svg := Banner(req)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "private, max-age=86400")
	w.Write([]byte(svg))
}

// HandleIcon is an HTTP handler that generates icon SVGs.
//
//	POST /icon — JSON body with IconRequest fields
//	GET  /icon — query params: mode, label, icon, size, accent, radius, shadow
func HandleIcon(w http.ResponseWriter, r *http.Request) {
	securityHeaders(w)

	if r.Method != http.MethodGet && r.Method != http.MethodPost && r.Method != http.MethodOptions {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req IconRequest

	if r.Method == http.MethodPost {
		body := http.MaxBytesReader(w, r.Body, MaxRequestBodyBytes)
		defer body.Close()
		if err := json.NewDecoder(body).Decode(&req); err != nil {
			http.Error(w, `{"error":"invalid json"}`, http.StatusBadRequest)
			return
		}
	} else {
		q := r.URL.Query()
		req.Mode = q.Get("mode")
		req.Label = q.Get("label")
		req.Icon = q.Get("icon")
		req.Accent = q.Get("accent")
		req.Radius = q.Get("radius")
		if s := q.Get("size"); s != "" {
			var size int
			if err := json.Unmarshal([]byte(s), &size); err == nil {
				req.Size = size
			}
		}
		req.Shadow = queryBool(q, "shadow")
	}

	req.Sanitize()

	svg := Icon(req)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "private, max-age=86400")
	w.Write([]byte(svg))
}

// HandleIcons returns the list of available remix icon names.
func HandleIcons(w http.ResponseWriter, r *http.Request) {
	securityHeaders(w)

	if r.Method != http.MethodGet && r.Method != http.MethodOptions {
		http.Error(w, `{"error":"method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	names := make([]string, 0, len(RemixIcons))
	for name := range RemixIcons {
		names = append(names, name)
	}
	sort.Strings(names)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "private, max-age=86400")
	json.NewEncoder(w).Encode(names)
}

func splitTags(s string) []string {
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}

func queryBool(q url.Values, key string) *bool {
	v := q.Get(key)
	if v == "" {
		return nil
	}
	b := v == "true" || v == "1" || v == "yes"
	return &b
}
