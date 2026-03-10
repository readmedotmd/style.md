package generate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"unicode/utf8"
)

func TestBannerCard(t *testing.T) {
	svg := Banner(BannerRequest{
		Name:    "store.md",
		Tagline: "Key-value storage interface for Go",
		Desc:    "One interface. Six backends.",
		Version: "v0.1.0",
		Tags:    []string{"go", "MIT"},
	})
	assertSVG(t, svg)
	assertContains(t, svg, "STORE.md")
	assertContains(t, svg, "KEY-VALUE STORAGE INTERFACE FOR GO")
	assertContains(t, svg, "VERSION")
	assertContains(t, svg, "V0.1.0")
}

func TestBannerTerminal(t *testing.T) {
	svg := Banner(BannerRequest{
		Name:    "gui.md",
		Tagline: "Virtual DOM for Go",
		Layout:  "terminal",
	})
	assertSVG(t, svg)
	assertContains(t, svg, "GUI.MD — TERMINAL")
	assertContains(t, svg, "▸ gui.md")
}

func TestBannerMinimal(t *testing.T) {
	svg := Banner(BannerRequest{
		Name:    "style.md",
		Tagline: "Industrial Monospace Design",
		Layout:  "minimal",
	})
	assertSVG(t, svg)
	assertContains(t, svg, "style.md")
}

func TestBannerDefaults(t *testing.T) {
	svg := Banner(BannerRequest{})
	assertSVG(t, svg)
	assertContains(t, svg, "PROJECT.md")
}

func TestBannerSizes(t *testing.T) {
	for _, size := range []string{"900x280", "1800x560", "1280x640", "800x200"} {
		svg := Banner(BannerRequest{Name: "test", Size: size})
		w, _ := parseDimensions(size)
		assertContains(t, svg, fmt.Sprintf("%d", w))
	}
}

func TestIconText(t *testing.T) {
	svg := Icon(IconRequest{
		Mode:  "text",
		Label: "St",
		Size:  512,
	})
	assertSVG(t, svg)
	assertContains(t, svg, ">St</text>")
}

func TestIconRemix(t *testing.T) {
	svg := Icon(IconRequest{
		Mode: "icon",
		Icon: "terminal-box",
		Size: 256,
	})
	assertSVG(t, svg)
	assertContains(t, svg, "<path d=")
	assertContains(t, svg, "scale(")
}

func TestIconDefaults(t *testing.T) {
	svg := Icon(IconRequest{})
	assertSVG(t, svg)
	assertContains(t, svg, "1024")
}

func TestIconNoShadow(t *testing.T) {
	f := false
	svg := Icon(IconRequest{Shadow: &f, Size: 128})
	assertSVG(t, svg)
	// Without shadow, viewBox should equal size
	assertContains(t, svg, `width="128"`)
}

func TestIconRadii(t *testing.T) {
	for _, r := range []string{"0", "small", "medium", "large", "xl"} {
		svg := Icon(IconRequest{Radius: r, Size: 256})
		assertSVG(t, svg)
	}
}

func TestFormatDisplayName(t *testing.T) {
	cases := []struct{ in, want string }{
		{"store.md", "STORE.md"},
		{"gui.md", "GUI.md"},
		{"myproject", "MYPROJECT"},
		{"my.cool.lib", "MY.COOL.lib"},
	}
	for _, c := range cases {
		got := formatDisplayName(c.in)
		if got != c.want {
			t.Errorf("formatDisplayName(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestFormatUpperWithExt(t *testing.T) {
	cases := []struct{ in, want string }{
		{"by README.md", "by README.md"},
		{"go", "GO"},
		{"MIT", "MIT"},
	}
	for _, c := range cases {
		got := formatUpperWithExt(c.in)
		if got != c.want {
			t.Errorf("formatUpperWithExt(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

func TestEscXML(t *testing.T) {
	got := escXML(`<a & "b">`)
	want := "&lt;a &amp; &quot;b&quot;&gt;"
	if got != want {
		t.Errorf("escXML = %q, want %q", got, want)
	}
}

func TestHandlerBannerGET(t *testing.T) {
	req := httptest.NewRequest("GET", "/banner?name=test.md&layout=card&size=900x280", nil)
	w := httptest.NewRecorder()
	HandleBanner(w, req)

	if w.Code != 200 {
		t.Fatalf("status = %d", w.Code)
	}
	if ct := w.Header().Get("Content-Type"); ct != "image/svg+xml" {
		t.Fatalf("content-type = %s", ct)
	}
	assertSVG(t, w.Body.String())
}

func TestHandlerBannerPOST(t *testing.T) {
	body := `{"name":"api.md","tagline":"REST API","layout":"terminal","tags":["go","http"]}`
	req := httptest.NewRequest("POST", "/banner", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	HandleBanner(w, req)

	if w.Code != 200 {
		t.Fatalf("status = %d", w.Code)
	}
	assertSVG(t, w.Body.String())
	assertContains(t, w.Body.String(), "API.MD — TERMINAL")
}

func TestHandlerIconGET(t *testing.T) {
	req := httptest.NewRequest("GET", "/icon?mode=text&label=Db&size=256", nil)
	w := httptest.NewRecorder()
	HandleIcon(w, req)

	if w.Code != 200 {
		t.Fatalf("status = %d", w.Code)
	}
	assertSVG(t, w.Body.String())
	assertContains(t, w.Body.String(), ">Db</text>")
}

func TestHandlerIcons(t *testing.T) {
	req := httptest.NewRequest("GET", "/icons", nil)
	w := httptest.NewRecorder()
	HandleIcons(w, req)

	var names []string
	if err := json.NewDecoder(w.Body).Decode(&names); err != nil {
		t.Fatal(err)
	}
	if len(names) != len(RemixIcons) {
		t.Errorf("got %d icons, want %d", len(names), len(RemixIcons))
	}
}


func TestSanitizeColor(t *testing.T) {
	cases := []struct{ color, fallback, want string }{
		{"#FFF", "x", "#FFF"},
		{"#FF5500", "x", "#FF5500"},
		{"#abc", "x", "#abc"},
		// Invalid colors - should return fallback
		{"#ABCD", "#000", "#000"},
		{"#ABCDE", "#000", "#000"},
		{"red", "#000", "#000"},
		{"rgb(0,0,0)", "#000", "#000"},
		{"", "#000", "#000"},
		{"#GGG", "#000", "#000"},
	}
	for _, c := range cases {
		got := sanitizeColor(c.color, c.fallback)
		if got != c.want {
			t.Errorf("sanitizeColor(%q, %q) = %q, want %q", c.color, c.fallback, got, c.want)
		}
	}
}

func TestTruncate(t *testing.T) {
	if got := truncate("hello", 3); got != "hel" {
		t.Errorf("truncate = %q, want %q", got, "hel")
	}
	if got := truncate("hi", 10); got != "hi" {
		t.Errorf("truncate = %q, want %q", got, "hi")
	}
	// Multi-byte: truncate by runes, not bytes
	if got := truncate("日本語テスト", 3); got != "日本語" {
		t.Errorf("truncate multibyte = %q, want %q", got, "日本語")
	}
}

func TestClampInt(t *testing.T) {
	cases := []struct{ v, min, max, want int }{
		{5, 0, 10, 5},
		{-1, 0, 10, 0},
		{100, 0, 10, 10},
	}
	for _, c := range cases {
		if got := clampInt(c.v, c.min, c.max); got != c.want {
			t.Errorf("clampInt(%d, %d, %d) = %d, want %d", c.v, c.min, c.max, got, c.want)
		}
	}
}

func TestParseDimensions(t *testing.T) {
	cases := []struct {
		size          string
		wantW, wantH int
	}{
		{"1800x560", 1800, 560},
		{"100x100", 100, 100},
		{"invalid", 1800, 560},
		{"0x0", 1800, 560},
		{"axb", 1800, 560},
		{"100x", 1800, 560},
		{"x100", 1800, 560},
		{"10x10", 50, 50},
		{"9999x9999", 4096, 4096},
	}
	for _, c := range cases {
		w, h := parseDimensions(c.size)
		if w != c.wantW || h != c.wantH {
			t.Errorf("parseDimensions(%q) = (%d, %d), want (%d, %d)", c.size, w, h, c.wantW, c.wantH)
		}
	}
}

func TestBannerSanitize(t *testing.T) {
	long := strings.Repeat("a", 300)
	req := BannerRequest{
		Name:    long,
		Tagline: long,
		Desc:    long,
		Version: long,
		Accent:  "not-a-color",
		Tags:    []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
	}
	req.Sanitize()
	if len(req.Name) != MaxStringLen {
		t.Errorf("Name not truncated: len=%d", len(req.Name))
	}
	if len(req.Tags) != MaxTags {
		t.Errorf("Tags not limited: len=%d", len(req.Tags))
	}
	if req.Accent != "" {
		t.Errorf("Invalid accent not cleared: %q", req.Accent)
	}
}

func TestIconSanitize(t *testing.T) {
	req := IconRequest{
		Size:   99999,
		Accent: "invalid",
		Label:  strings.Repeat("x", 300),
	}
	req.Sanitize()
	if req.Size != MaxIconSize {
		t.Errorf("Size not clamped: %d", req.Size)
	}
	if req.Accent != "" {
		t.Errorf("Invalid accent not cleared: %q", req.Accent)
	}
	if len(req.Label) != MaxStringLen {
		t.Errorf("Label not truncated: len=%d", len(req.Label))
	}
}

func TestEscXMLInjection(t *testing.T) {
	cases := []struct{ in, notContain string }{
		{`<script>alert(1)</script>`, "<script>"},
		{`" onload="alert(1)`, `" onload="`},
		{`' onmouseover='alert(1)`, `' onmouseover='`},
	}
	for _, c := range cases {
		got := escXML(c.in)
		if strings.Contains(got, c.notContain) {
			t.Errorf("escXML(%q) still contains %q: %s", c.in, c.notContain, got)
		}
	}
}

func TestHandlerInvalidMethod(t *testing.T) {
	req := httptest.NewRequest("DELETE", "/banner", nil)
	w := httptest.NewRecorder()
	HandleBanner(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("DELETE /banner status = %d, want 405", w.Code)
	}
}

func TestHandlerSecurityHeaders(t *testing.T) {
	req := httptest.NewRequest("GET", "/banner?name=test", nil)
	w := httptest.NewRecorder()
	HandleBanner(w, req)

	checks := map[string]string{
		"X-Content-Type-Options":  "nosniff",
		"X-Frame-Options":        "DENY",
		"Content-Security-Policy": "default-src 'none'",
	}
	for header, want := range checks {
		if got := w.Header().Get(header); got != want {
			t.Errorf("header %s = %q, want %q", header, got, want)
		}
	}
}

func TestHandlerInvalidJSON(t *testing.T) {
	req := httptest.NewRequest("POST", "/banner", strings.NewReader("{invalid"))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	HandleBanner(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("invalid JSON status = %d, want 400", w.Code)
	}
}

func TestIconUnknownName(t *testing.T) {
	svg := Icon(IconRequest{Mode: "icon", Icon: "nonexistent-icon", Size: 256})
	assertSVG(t, svg)
}

func TestHandlerBodyLimit(t *testing.T) {
	body := strings.Repeat("x", MaxRequestBodyBytes+1)
	req := httptest.NewRequest("POST", "/banner", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	HandleBanner(w, req)
	if w.Code != http.StatusBadRequest {
		t.Errorf("oversized body status = %d, want 400", w.Code)
	}
}

func assertSVG(t *testing.T, svg string) {
	t.Helper()
	if !strings.HasPrefix(svg, "<svg") {
		t.Errorf("expected SVG, got: %.80s...", svg)
	}
	if !strings.HasSuffix(svg, "</svg>") {
		t.Errorf("expected SVG to end with </svg>")
	}
}

func assertContains(t *testing.T, s, substr string) {
	t.Helper()
	if !strings.Contains(s, substr) {
		t.Errorf("expected SVG to contain %q", substr)
	}
}

func FuzzBannerSanitize(f *testing.F) {
	f.Add("myproject", "A tagline", "Some description", "v1.0.0", "#FF5500", "go,web,api")
	f.Add("", "", "", "", "", "")
	f.Add(strings.Repeat("x", 500), strings.Repeat("y", 500), strings.Repeat("z", 500), strings.Repeat("w", 500), "invalid", "a,b,c,d,e,f,g,h,i,j,k,l,m")

	f.Fuzz(func(t *testing.T, name, tagline, desc, version, accent, tagsStr string) {
		tags := strings.Split(tagsStr, ",")
		req := BannerRequest{
			Name:    name,
			Tagline: tagline,
			Desc:    desc,
			Version: version,
			Accent:  accent,
			Tags:    tags,
		}
		req.Sanitize()

		if utf8.RuneCountInString(req.Name) > MaxStringLen {
			t.Errorf("Name rune count %d > MaxStringLen %d", utf8.RuneCountInString(req.Name), MaxStringLen)
		}
		if len(req.Tags) > MaxTags {
			t.Errorf("Tags length %d > MaxTags %d", len(req.Tags), MaxTags)
		}
		for i, tag := range req.Tags {
			if utf8.RuneCountInString(tag) > MaxTagLen {
				t.Errorf("Tag[%d] rune count %d > MaxTagLen %d", i, utf8.RuneCountInString(tag), MaxTagLen)
			}
		}
		if req.Accent != "" && !validHexColor.MatchString(req.Accent) {
			t.Errorf("Accent %q is non-empty but not a valid hex color", req.Accent)
		}
	})
}

func FuzzIconSanitize(f *testing.F) {
	f.Add("St", 512, "#abc")
	f.Add("", 0, "")
	f.Add(strings.Repeat("z", 500), 99999, "bad")

	f.Fuzz(func(t *testing.T, label string, size int, accent string) {
		req := IconRequest{
			Label:  label,
			Size:   size,
			Accent: accent,
		}
		req.Sanitize()

		if utf8.RuneCountInString(req.Label) > MaxStringLen {
			t.Errorf("Label rune count %d > MaxStringLen %d", utf8.RuneCountInString(req.Label), MaxStringLen)
		}
		if req.Size != 0 && (req.Size < MinIconSize || req.Size > MaxIconSize) {
			t.Errorf("Size %d not 0 and outside [%d, %d]", req.Size, MinIconSize, MaxIconSize)
		}
		if req.Accent != "" && !validHexColor.MatchString(req.Accent) {
			t.Errorf("Accent %q is non-empty but not a valid hex color", req.Accent)
		}
	})
}

func FuzzBannerGeneration(f *testing.F) {
	f.Add("store.md", "Key-value storage", "One interface")
	f.Add("", "", "")
	f.Add("<script>alert(1)</script>", "tag&line", `"desc"`)

	f.Fuzz(func(t *testing.T, name, tagline, desc string) {
		svg := Banner(BannerRequest{
			Name:    name,
			Tagline: tagline,
			Desc:    desc,
		})
		if !strings.HasPrefix(svg, "<svg") {
			t.Errorf("output does not start with <svg: %.80s", svg)
		}
		if !strings.HasSuffix(svg, "</svg>") {
			t.Errorf("output does not end with </svg>")
		}
	})
}

func FuzzEscXML(f *testing.F) {
	f.Add("hello world")
	f.Add(`<script>alert("xss")</script>`)
	f.Add(`a & b "c" 'd'`)
	f.Add("")

	f.Fuzz(func(t *testing.T, input string) {
		got := escXML(input)
		if strings.ContainsAny(got, `<>"'`) {
			t.Errorf("escXML(%q) = %q contains unescaped special character", input, got)
		}
		// Check for unescaped & (ampersands that are not part of XML entities)
		for i := 0; i < len(got); i++ {
			if got[i] == '&' {
				rest := got[i:]
				if !strings.HasPrefix(rest, "&amp;") &&
					!strings.HasPrefix(rest, "&lt;") &&
					!strings.HasPrefix(rest, "&gt;") &&
					!strings.HasPrefix(rest, "&quot;") &&
					!strings.HasPrefix(rest, "&apos;") &&
					!strings.HasPrefix(rest, "&#") {
					t.Errorf("escXML(%q) = %q contains unescaped & at position %d", input, got, i)
				}
			}
		}
	})
}
