package main

import (
	_ "embed"
	"encoding/base64"
	"encoding/json"
	"sort"
	"strings"

	"github.com/readmedotmd/style.md/generate"
)

//go:embed root.json
var rootJSON string

// Main is the DO Functions entry point.
// Dispatch on the "type" parameter: "root", "banner" (default), "icon", or "icons".
func Main(args map[string]interface{}) map[string]interface{} {
	typ := strArg(args, "type", "banner")

	switch typ {
	case "root":
		if strArg(args, "api_key", "") == "" {
			return jsonResponse(`{"error":"unauthorized","message":"API key required for root endpoint"}`)
		}
		return rootResponse()

	case "icon":
		req := generate.IconRequest{
			Mode:   strArg(args, "mode", ""),
			Label:  strArg(args, "label", ""),
			Icon:   strArg(args, "icon", ""),
			Size:   intArg(args, "size", 0),
			Accent: strArg(args, "accent", ""),
			Radius: strArg(args, "radius", ""),
			Shadow: boolPtrArg(args, "shadow"),
		}
		req.Sanitize()
		svg := generate.Icon(req)
		return svgResponse(svg)

	case "icons":
		names := make([]string, 0, len(generate.RemixIcons))
		for name := range generate.RemixIcons {
			names = append(names, name)
		}
		sort.Strings(names)
		body, _ := json.Marshal(names)
		return jsonResponse(string(body))

	default:
		req := generate.BannerRequest{
			Name:      strArg(args, "name", ""),
			Tagline:   strArg(args, "tagline", ""),
			Desc:      strArg(args, "desc", ""),
			Version:   strArg(args, "version", ""),
			Accent:    strArg(args, "accent", ""),
			Layout:    strArg(args, "layout", ""),
			Size:      strArg(args, "size", ""),
			Tags:      tagsArg(args, "tags"),
			ShowDots:  boolPtrArg(args, "show_dots"),
			ShowBadge: boolPtrArg(args, "show_badge"),
			ShowTags:  boolPtrArg(args, "show_tags"),
		}
		req.Sanitize()
		svg := generate.Banner(req)
		return svgResponse(svg)
	}
}

// svgResponse wraps an SVG string in the DO Functions response format.
func svgResponse(svg string) map[string]interface{} {
	return map[string]interface{}{
		"statusCode": 200,
		"headers": map[string]interface{}{
			"Content-Type":  "image/svg+xml",
			"Cache-Control": "private, max-age=86400",
		},
		"body": base64.StdEncoding.EncodeToString([]byte(svg)),
	}
}

// jsonResponse wraps a JSON string in the DO Functions response format.
func jsonResponse(body string) map[string]interface{} {
	return map[string]interface{}{
		"statusCode": 200,
		"headers": map[string]interface{}{
			"Content-Type": "application/json",
		},
		"body": body,
	}
}

// rootResponse returns AI-readable instructions for connecting to and using
// the style.md SVG generation API, including a ready-to-use GitHub Action.
func rootResponse() map[string]interface{} {
	return map[string]interface{}{
		"statusCode": 200,
		"headers": map[string]interface{}{
			"Content-Type":  "application/json",
			"Cache-Control": "public, max-age=3600",
		},
		"body": rootJSON,
	}
}

// --- DO Functions arg-parsing helpers ---

func strArg(args map[string]interface{}, key, def string) string {
	v, ok := args[key]
	if !ok {
		return def
	}
	s, ok := v.(string)
	if !ok || s == "" {
		return def
	}
	return s
}

func intArg(args map[string]interface{}, key string, def int) int {
	v, ok := args[key]
	if !ok {
		return def
	}
	switch val := v.(type) {
	case float64:
		return int(val)
	case int:
		return val
	default:
		return def
	}
}

func boolPtrArg(args map[string]interface{}, key string) *bool {
	v, ok := args[key]
	if !ok {
		return nil
	}
	var b bool
	switch val := v.(type) {
	case bool:
		b = val
	case string:
		b = val == "true" || val == "1" || val == "yes"
	default:
		return nil
	}
	return &b
}

func tagsArg(args map[string]interface{}, key string) []string {
	v, ok := args[key]
	if !ok {
		return nil
	}
	var raw []string
	switch val := v.(type) {
	case string:
		parts := strings.Split(val, ",")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p != "" {
				raw = append(raw, p)
			}
		}
	case []interface{}:
		for _, item := range val {
			if s, ok := item.(string); ok && s != "" {
				raw = append(raw, s)
			}
		}
	default:
		return nil
	}
	return raw
}
