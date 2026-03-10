// Package generate produces SVG banners and icons matching the style.md
// industrial monospace design system.
package generate

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"unicode/utf8"
)

// Input limits to prevent abuse.
const (
	MaxStringLen  = 200  // max length for any single text field
	MaxTags       = 10   // max number of tags
	MaxTagLen     = 40   // max length per tag
	MaxDimension  = 4096 // max width or height in pixels
	MinDimension  = 50   // min width or height in pixels
	MaxIconSize   = 2048 // max icon size in pixels
	MinIconSize   = 16   // min icon size in pixels
)

// validHexColor matches 3 or 6 digit hex color codes with # prefix.
var validHexColor = regexp.MustCompile(`^#([0-9A-Fa-f]{3}|[0-9A-Fa-f]{6})$`)

// truncate returns s truncated to maxLen.
func truncate(s string, maxLen int) string {
	if utf8.RuneCountInString(s) > maxLen {
		r := []rune(s)
		return string(r[:maxLen])
	}
	return s
}

// clampInt clamps v to [min, max].
func clampInt(v, min, max int) int {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}

// sanitizeColor returns the color if it's a valid hex color, otherwise returns the fallback.
func sanitizeColor(color, fallback string) string {
	if validHexColor.MatchString(color) {
		return color
	}
	return fallback
}

// Color constants matching the style.md design tokens.
var Colors = struct {
	PageBg      string
	Surface     string
	Dark        string
	Border      string
	Muted       string
	DotRed      string
	DotYellow   string
	DotGreen    string
	ShadowAlpha string
	Font        string
}{
	PageBg:      "#FFFFFF",
	Surface:     "#FFFFFF",
	Dark:        "#2D333B",
	Border:      "#000000",
	Muted:       "#888888",
	DotRed:      "#FF3B30",
	DotYellow:   "#FFCC00",
	DotGreen:    "#34C759",
	ShadowAlpha: "rgba(0,0,0,0.25)",
	Font:        "'Space Mono', 'Courier New', monospace",
}

// BannerRequest defines parameters for generating a banner SVG.
type BannerRequest struct {
	Name      string   `json:"name"`
	Tagline   string   `json:"tagline"`
	Desc      string   `json:"desc"`
	Version   string   `json:"version"`
	Accent    string   `json:"accent"`
	Layout    string   `json:"layout"`
	Tags      []string `json:"tags"`
	Size      string   `json:"size"`
	ShowDots  *bool    `json:"show_dots"`
	ShowBadge *bool    `json:"show_badge"`
	ShowTags  *bool    `json:"show_tags"`
}

// IconRequest defines parameters for generating an icon SVG.
type IconRequest struct {
	Mode      string `json:"mode"`
	Label     string `json:"label"`
	Icon      string `json:"icon"`
	Size      int    `json:"size"`
	Accent    string `json:"accent"`
	Radius    string `json:"radius"`
	Shadow    *bool  `json:"shadow"`
}

// Sanitize clamps and truncates all fields to safe limits.
func (r *BannerRequest) Sanitize() {
	r.Name = truncate(r.Name, MaxStringLen)
	r.Tagline = truncate(r.Tagline, MaxStringLen)
	r.Desc = truncate(r.Desc, MaxStringLen)
	r.Version = truncate(r.Version, MaxStringLen)
	r.Accent = sanitizeColor(r.Accent, "")
	r.Layout = truncate(r.Layout, 20)
	if len(r.Tags) > MaxTags {
		r.Tags = r.Tags[:MaxTags]
	}
	for i, t := range r.Tags {
		r.Tags[i] = truncate(t, MaxTagLen)
	}
}

// Sanitize clamps and truncates all fields to safe limits.
func (r *IconRequest) Sanitize() {
	r.Label = truncate(r.Label, MaxStringLen)
	r.Icon = truncate(r.Icon, MaxStringLen)
	r.Mode = truncate(r.Mode, 20)
	r.Accent = sanitizeColor(r.Accent, "")
	r.Radius = truncate(r.Radius, 10)
	if r.Size != 0 {
		r.Size = clampInt(r.Size, MinIconSize, MaxIconSize)
	}
}

func boolDefault(b *bool, def bool) bool {
	if b == nil {
		return def
	}
	return *b
}

func round(f float64) int {
	return int(math.Round(f))
}

func escXML(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	s = strings.ReplaceAll(s, "'", "&#39;")
	return s
}

var lowerWords = map[string]bool{
	"by": true, "on": true, "for": true, "the": true,
	"a": true, "an": true, "of": true, "in": true, "to": true,
}

// formatUpperWithExt formats a string with smart case: NAME.ext for dotted
// words, lowercase for articles/prepositions, UPPER for everything else.
func formatUpperWithExt(s string) string {
	words := strings.Fields(strings.TrimSpace(s))
	for i, word := range words {
		parts := strings.Split(word, ".")
		if len(parts) > 1 {
			name := strings.ToUpper(strings.Join(parts[:len(parts)-1], "."))
			ext := strings.ToLower(parts[len(parts)-1])
			words[i] = name + "." + ext
		} else if lowerWords[strings.ToLower(word)] {
			words[i] = strings.ToLower(word)
		} else {
			words[i] = strings.ToUpper(word)
		}
	}
	return strings.Join(words, " ")
}

// parseDimensions parses a "WxH" size string into width and height.
// Dimensions are clamped to [MinDimension, MaxDimension].
func parseDimensions(size string) (int, int) {
	parts := strings.SplitN(size, "x", 2)
	if len(parts) != 2 {
		return 1800, 560
	}
	w, h := 0, 0
	fmt.Sscanf(parts[0], "%d", &w)
	fmt.Sscanf(parts[1], "%d", &h)
	if w == 0 || h == 0 {
		return 1800, 560
	}
	return clampInt(w, MinDimension, MaxDimension), clampInt(h, MinDimension, MaxDimension)
}
