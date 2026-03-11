package main

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
)

// ---------------------------------------------------------------------------
// Response helpers
// ---------------------------------------------------------------------------

func respHeaders(resp map[string]interface{}) map[string]interface{} {
	h, _ := resp["headers"].(map[string]interface{})
	return h
}

func respBody(resp map[string]interface{}) string {
	b, _ := resp["body"].(string)
	return b
}

// respSVGBody decodes the base64-encoded SVG body.
func respSVGBody(resp map[string]interface{}) string {
	b, _ := resp["body"].(string)
	decoded, err := base64.StdEncoding.DecodeString(b)
	if err != nil {
		return b
	}
	return string(decoded)
}

func respContentType(resp map[string]interface{}) string {
	h := respHeaders(resp)
	ct, _ := h["Content-Type"].(string)
	return ct
}

// ---------------------------------------------------------------------------
// 1. TestMain_BannerDefault
// ---------------------------------------------------------------------------

func TestMain_BannerDefault(t *testing.T) {
	resp := Main(map[string]interface{}{})

	if ct := respContentType(resp); ct != "image/svg+xml" {
		t.Fatalf("expected Content-Type image/svg+xml, got %q", ct)
	}

	body := respSVGBody(resp)
	if !strings.HasPrefix(body, "<svg") {
		t.Fatalf("body should start with <svg, got prefix %q", body[:40])
	}
	if !strings.HasSuffix(body, "</svg>") {
		t.Fatalf("body should end with </svg>")
	}
}

// ---------------------------------------------------------------------------
// 2. TestMain_BannerWithParams
// ---------------------------------------------------------------------------

func TestMain_BannerWithParams(t *testing.T) {
	resp := Main(map[string]interface{}{
		"name":    "TestProject",
		"tagline": "A test tagline",
		"layout":  "terminal",
	})

	if ct := respContentType(resp); ct != "image/svg+xml" {
		t.Fatalf("expected Content-Type image/svg+xml, got %q", ct)
	}

	body := respSVGBody(resp)
	if !strings.Contains(body, "TestProject") {
		t.Error("body should contain the project name")
	}
	if !strings.Contains(body, "A test tagline") {
		t.Error("body should contain the tagline")
	}
	if !strings.HasPrefix(body, "<svg") {
		t.Error("body should start with <svg")
	}
	if !strings.HasSuffix(body, "</svg>") {
		t.Error("body should end with </svg>")
	}
}

// ---------------------------------------------------------------------------
// 3. TestMain_Icon
// ---------------------------------------------------------------------------

func TestMain_Icon(t *testing.T) {
	resp := Main(map[string]interface{}{
		"type":  "icon",
		"mode":  "text",
		"label": "Ab",
	})

	if ct := respContentType(resp); ct != "image/svg+xml" {
		t.Fatalf("expected Content-Type image/svg+xml, got %q", ct)
	}

	body := respSVGBody(resp)
	if !strings.HasPrefix(body, "<svg") {
		t.Error("body should start with <svg")
	}
	if !strings.HasSuffix(body, "</svg>") {
		t.Error("body should end with </svg>")
	}
	if !strings.Contains(body, "Ab") {
		t.Error("body should contain the label text")
	}
}

// ---------------------------------------------------------------------------
// 4. TestMain_Icons
// ---------------------------------------------------------------------------

func TestMain_Icons(t *testing.T) {
	resp := Main(map[string]interface{}{
		"type": "icons",
	})

	if ct := respContentType(resp); ct != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", ct)
	}

	body := respBody(resp)
	var names []string
	if err := json.Unmarshal([]byte(body), &names); err != nil {
		t.Fatalf("failed to parse JSON body: %v", err)
	}
	if len(names) == 0 {
		t.Fatal("expected at least one icon name")
	}
}

// ---------------------------------------------------------------------------
// 5. TestMain_Root
// ---------------------------------------------------------------------------

func TestMain_Root(t *testing.T) {
	resp := Main(map[string]interface{}{
		"type": "root",
	})

	if ct := respContentType(resp); ct != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", ct)
	}

	body := respBody(resp)
	var info map[string]interface{}
	if err := json.Unmarshal([]byte(body), &info); err != nil {
		t.Fatalf("failed to parse JSON body: %v", err)
	}

	if _, ok := info["service"]; !ok {
		t.Error("response should contain a 'service' key")
	}
	if _, ok := info["endpoints"]; !ok {
		t.Error("response should contain an 'endpoints' key")
	}
}

// ---------------------------------------------------------------------------
// 5b. TestMain_APIKeyAuth
// ---------------------------------------------------------------------------

func TestMain_AuthRejectsWithoutKey(t *testing.T) {
	resp := Main(map[string]interface{}{
		"STYLEMD_API_KEY": "secret123",
		"type":            "banner",
	})

	if sc := resp["statusCode"]; sc != 401 {
		t.Fatalf("expected 401, got %v", sc)
	}
	body := respBody(resp)
	if !strings.Contains(body, "unauthorized") {
		t.Errorf("expected unauthorized error, got %q", body)
	}
}

func TestMain_AuthRejectsWrongKey(t *testing.T) {
	resp := Main(map[string]interface{}{
		"STYLEMD_API_KEY": "secret123",
		"api_key":         "wrong",
		"type":            "banner",
	})

	if sc := resp["statusCode"]; sc != 401 {
		t.Fatalf("expected 401, got %v", sc)
	}
}

func TestMain_AuthAcceptsCorrectKey(t *testing.T) {
	resp := Main(map[string]interface{}{
		"STYLEMD_API_KEY": "secret123",
		"api_key":         "secret123",
		"type":            "banner",
	})

	if sc := resp["statusCode"]; sc != 200 {
		t.Fatalf("expected 200, got %v", sc)
	}
	if ct := respContentType(resp); ct != "image/svg+xml" {
		t.Fatalf("expected image/svg+xml, got %q", ct)
	}
}

func TestMain_NoAuthWhenKeyNotConfigured(t *testing.T) {
	resp := Main(map[string]interface{}{
		"type": "banner",
	})

	if sc := resp["statusCode"]; sc != 200 {
		t.Fatalf("expected 200 without key configured, got %v", sc)
	}
}

// ---------------------------------------------------------------------------
// 6. TestStrArg
// ---------------------------------------------------------------------------

func TestStrArg(t *testing.T) {
	args := map[string]interface{}{
		"present": "hello",
		"empty":   "",
		"long":    strings.Repeat("x", 300),
	}

	// present key
	if got := strArg(args, "present", "def"); got != "hello" {
		t.Errorf("present: expected 'hello', got %q", got)
	}

	// missing key returns default
	if got := strArg(args, "missing", "def"); got != "def" {
		t.Errorf("missing: expected 'def', got %q", got)
	}

	// empty string returns default
	if got := strArg(args, "empty", "def"); got != "def" {
		t.Errorf("empty: expected 'def', got %q", got)
	}

	// long string is passed through (Sanitize handles truncation)
	if got := strArg(args, "long", "def"); len(got) != 300 {
		t.Errorf("long: expected length 300, got %d", len(got))
	}
}

// ---------------------------------------------------------------------------
// 7. TestIntArg
// ---------------------------------------------------------------------------

func TestIntArg(t *testing.T) {
	args := map[string]interface{}{
		"float":  float64(42),
		"int":    int(99),
		"string": "nope",
	}

	// float64 (default JSON numeric type)
	if got := intArg(args, "float", 0); got != 42 {
		t.Errorf("float: expected 42, got %d", got)
	}

	// int
	if got := intArg(args, "int", 0); got != 99 {
		t.Errorf("int: expected 99, got %d", got)
	}

	// missing key returns default
	if got := intArg(args, "missing", 7); got != 7 {
		t.Errorf("missing: expected 7, got %d", got)
	}

	// non-numeric returns default
	if got := intArg(args, "string", 7); got != 7 {
		t.Errorf("string: expected 7, got %d", got)
	}
}

// ---------------------------------------------------------------------------
// 8. TestBoolPtrArg
// ---------------------------------------------------------------------------

func TestBoolPtrArg(t *testing.T) {
	args := map[string]interface{}{
		"bool_true":  true,
		"bool_false": false,
		"str_true":   "true",
		"str_false":  "false",
		"str_one":    "1",
	}

	// bool true
	if got := boolPtrArg(args, "bool_true"); got == nil || *got != true {
		t.Error("bool_true: expected ptr to true")
	}

	// bool false
	if got := boolPtrArg(args, "bool_false"); got == nil || *got != false {
		t.Error("bool_false: expected ptr to false")
	}

	// string "true"
	if got := boolPtrArg(args, "str_true"); got == nil || *got != true {
		t.Error("str_true: expected ptr to true")
	}

	// string "false"
	if got := boolPtrArg(args, "str_false"); got == nil || *got != false {
		t.Error("str_false: expected ptr to false")
	}

	// string "1"
	if got := boolPtrArg(args, "str_one"); got == nil || *got != true {
		t.Error("str_one: expected ptr to true")
	}

	// missing key returns nil
	if got := boolPtrArg(args, "missing"); got != nil {
		t.Error("missing: expected nil")
	}
}

// ---------------------------------------------------------------------------
// 9. TestTagsArg
// ---------------------------------------------------------------------------

func TestTagsArg(t *testing.T) {
	// comma-separated string
	args1 := map[string]interface{}{
		"tags": "go, svg, cli",
	}
	got := tagsArg(args1, "tags")
	if len(got) != 3 || got[0] != "go" || got[1] != "svg" || got[2] != "cli" {
		t.Errorf("comma-separated: expected [go svg cli], got %v", got)
	}

	// JSON array ([]interface{})
	args2 := map[string]interface{}{
		"tags": []interface{}{"alpha", "beta"},
	}
	got = tagsArg(args2, "tags")
	if len(got) != 2 || got[0] != "alpha" || got[1] != "beta" {
		t.Errorf("json array: expected [alpha beta], got %v", got)
	}

	// missing key returns nil
	args3 := map[string]interface{}{}
	got = tagsArg(args3, "tags")
	if got != nil {
		t.Errorf("missing: expected nil, got %v", got)
	}

	// tagsArg passes all tags through; truncation is handled by Sanitize()
	manyTags := make([]interface{}, 15)
	for i := range manyTags {
		manyTags[i] = "tag"
	}
	args4 := map[string]interface{}{
		"tags": manyTags,
	}
	got = tagsArg(args4, "tags")
	if len(got) != 15 {
		t.Errorf("all tags: expected 15, got %d", len(got))
	}

	// long tags are passed through (Sanitize handles truncation)
	args5 := map[string]interface{}{
		"tags": []interface{}{strings.Repeat("a", 50)},
	}
	got = tagsArg(args5, "tags")
	if len(got) != 1 || len(got[0]) != 50 {
		t.Errorf("long tag: expected length 50, got %d", len(got[0]))
	}
}
