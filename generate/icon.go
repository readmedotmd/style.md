package generate

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// RemixIcons maps icon names to their 24x24 viewBox SVG path data.
var RemixIcons = map[string]string{
	"database-2":     "M21 9.5V12.5C21 14.985 16.971 17 12 17C7.029 17 3 14.985 3 12.5V9.5C3 11.985 7.029 14 12 14C16.971 14 21 11.985 21 9.5ZM3 14.5C3 16.985 7.029 19 12 19C16.971 19 21 16.985 21 14.5V17.5C21 19.985 16.971 22 12 22C7.029 22 3 19.985 3 17.5V14.5ZM12 12C16.971 12 21 9.985 21 7.5C21 5.015 16.971 3 12 3C7.029 3 3 5.015 3 7.5C3 9.985 7.029 12 12 12Z",
	"code-s-slash":   "M24 12L18.343 17.657L16.929 16.243L21.172 12L16.929 7.757L18.343 6.343L24 12ZM2.828 12L7.071 16.243L5.657 17.657L0 12L5.657 6.343L7.071 7.757L2.828 12ZM9.788 21H7.66L14.212 3H16.34L9.788 21Z",
	"terminal-box":   "M2 4C2 3.448 2.448 3 3 3H21C21.552 3 22 3.448 22 4V20C22 20.552 21.552 21 21 21H3C2.448 21 2 20.552 2 20V4ZM4 5V19H20V5H4ZM11.293 15.121L7.757 11.586L9.172 10.172L12.707 13.707L11.293 15.121ZM14 15H18V17H14V15Z",
	"git-branch":     "M7.105 15.21A3.001 3.001 0 1 0 5 15.17V8.83a3.001 3.001 0 1 0 2 0V12C8.775 10.772 10.7 10 13 10a1 1 0 0 0 1-1V8.83a3.001 3.001 0 1 0 2 0V9a3 3 0 0 1-3 3c-1.856 0-3.583.618-4.895 1.21ZM6 17a1 1 0 1 0 0 2 1 1 0 0 0 0-2ZM6 5a1 1 0 1 0 0 2 1 1 0 0 0 0-2Zm9 0a1 1 0 1 0 0 2 1 1 0 0 0 0-2Z",
	"settings-3":     "M3.34 17a10.018 10.018 0 0 1-.978-2.326 3 3 0 0 0 .002-5.347A9.99 9.99 0 0 1 4.865 4.99a3 3 0 0 0 4.631-2.674 9.99 9.99 0 0 1 5.007.002 3 3 0 0 0 4.632 2.672A9.99 9.99 0 0 1 20.66 7c.456.747.838 1.542 1.138 2.376a2.99 2.99 0 0 0-.002 5.247c-.3.835-.683 1.631-1.14 2.378a3 3 0 0 0-4.63 2.673 9.99 9.99 0 0 1-5.007-.002 3 3 0 0 0-4.631-2.672A10.018 10.018 0 0 1 3.34 17ZM12 15a3 3 0 1 0 0-6 3 3 0 0 0 0 6Z",
	"rocket-2":       "M5.33 15.929A13.064 13.064 0 0 1 5 13c0-5.088 2.903-9.436 7-11.182C16.097 3.564 19 7.912 19 13c0 1.01-.114 1.991-.33 2.929l2.02 1.796a.5.5 0 0 1 .097.63l-2.458 4.096a.5.5 0 0 1-.577.213l-2.765-.921a.5.5 0 0 1-.332-.38L14.2 19H9.8l-.455 2.361a.5.5 0 0 1-.332.381l-2.765.921a.5.5 0 0 1-.577-.214L3.213 18.354a.5.5 0 0 1 .098-.631l2.02-1.794ZM12 13a2 2 0 1 0 0-4 2 2 0 0 0 0 4Z",
	"send-plane":     "M1.946 9.315c-.522-.174-.527-.455.01-.634L21.044 2.32c.529-.176.832.12.684.638L16.764 21.44c-.15.529-.455.547-.68.045L12 14l6-8-8 6-8.054-2.685Z",
	"shield-check":   "M12 1l8.217 1.826a1 1 0 0 1 .783.976v9.987a6 6 0 0 1-2.672 4.992L12 23l-6.328-4.219A6 6 0 0 1 3 13.79V3.802a1 1 0 0 1 .783-.976L12 1Zm-1.452 12.134-1.966-1.966-1.414 1.414 3.38 3.38 5.656-5.656-1.414-1.414-4.242 4.242Z",
	"cpu":            "M6 18H18V6H6V18ZM14 2H16V6H14V2ZM14 18H16V22H14V18ZM2 14V16H6V14H2ZM18 14V16H22V14H18ZM8 2H10V6H8V2ZM8 18H10V22H8V18ZM2 8V10H6V8H2ZM18 8V10H22V8H18Z",
	"cloud":          "M17 21H7A6 6 0 0 1 5.008 9.339a7 7 0 1 1 13.984 0A6 6 0 0 1 17 21Z",
	"flask":          "M16 2V6.529L20.216 14.78A5 5 0 0 1 15.723 22H8.277A5 5 0 0 1 3.784 14.78L8 6.528V2H16ZM14 2V7L10 15H8L14 2Z",
	"puzzle-2":       "M4 8V2H20V8H18.005C18.002 6.895 17.107 6 16.003 6C14.898 6 14.003 6.895 14 7.999V8H10V7.999C9.997 6.895 9.102 6 7.997 6C6.893 6 5.998 6.895 5.995 7.999V8H4ZM14 10V16H14.001C14.004 14.895 14.899 14 16.003 14C17.108 14 18.003 14.895 18.005 16H20V10H14ZM12 10H4V16H5.995C5.998 14.895 6.893 14 7.997 14C9.102 14 9.997 14.895 10 16H12V10ZM4 18V22H20V18H18.005C18.002 19.105 17.107 20 16.003 20C14.898 20 14.003 19.105 14 18H10C9.997 19.105 9.102 20 7.997 20C6.893 20 5.998 19.105 5.995 18H4Z",
	"store-2":        "M22 20V22H2V20H3V13.242a5.965 5.965 0 0 1-1-.998V6l2.019-5.048A1 1 0 0 1 4.948 0h14.104a1 1 0 0 1 .929.632L22 6v6.243c-.346.38-.737.717-1.165 1V20H22ZM18 20H6V13.964c.24.023.484.036.73.036.803 0 1.563-.16 2.27-.453V20H15V13.547a5.93 5.93 0 0 0 2.27.453c.246 0 .49-.013.73-.036V20Z",
	"heart":          "M12.001 4.529a5.998 5.998 0 0 1 8.242.228 6 6 0 0 1 .236 8.236l-8.48 8.492-8.478-8.492a6 6 0 0 1 8.48-8.464Z",
	"lightning":      "M13 10h7l-9 13v-9H4l9-13v9Z",
	"eye":            "M12 2C17.523 2 22 6.477 22 12s-4.477 10-10 10S2 17.523 2 12 6.477 2 12 2Zm0 2a8 8 0 1 0 0 16 8 8 0 0 0 0-16Zm0 3a5 5 0 1 1 0 10 5 5 0 0 1 0-10Zm0 2a3 3 0 1 0 0 6 3 3 0 0 0 0-6Z",
	"download":       "M3 19H21V21H3V19ZM13 13.172L19.071 7.1L20.485 8.514L12 17L3.515 8.515L4.929 7.1L11 13.17V2H13V13.172Z",
	"folder-3":       "M22 11V20C22 20.552 21.552 21 21 21H3C2.448 21 2 20.552 2 20V4C2 3.448 2.448 3 3 3H10L12 5H21C21.552 5 22 5.448 22 6V7H4V11H22Z",
	"notification-3": "M20 17H22V19H2V17H4V10C4 5.582 7.582 2 12 2C16.418 2 20 5.582 20 10V17ZM9 21H15C15 22.657 13.657 24 12 24C10.343 24 9 22.657 9 21Z",
	"chat-3":         "M7.291 20.824L2 22L3.176 16.709A9.965 9.965 0 0 1 2 12C2 6.477 6.477 2 12 2C17.523 2 22 6.477 22 12C22 17.523 17.523 22 12 22C10.39 22 8.874 21.579 7.291 20.824Z",
	"user":           "M4 22a8 8 0 1 1 16 0H4Zm8-9c-3.315 0-6-2.685-6-6s2.685-6 6-6 6 2.685 6 6-2.685 6-6 6Z",
	"lock":           "M18 8h2a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h2V7a6 6 0 1 1 12 0v1Zm-2 0V7a4 4 0 0 0-8 0v1h8Zm-5 6v2h4v-2h-4Z",
	"key-2":          "M10.313 11.566l7.94-7.94 2.121 2.121-1.414 1.414 2.121 2.121-3.535 3.536-2.121-2.121-2.99 2.99a5.002 5.002 0 0 1-7.97 1.932 5 5 0 0 1 5.848-7.053ZM6.929 15.071a2 2 0 1 0 2.828-2.828 2 2 0 0 0-2.828 2.828Z",
}

var radiusMap = map[string]float64{
	"0": 0, "small": 0.03, "medium": 0.06, "large": 0.08, "xl": 0.12,
}

// Icon generates an SVG icon string from the given request.
func Icon(req IconRequest) string {
	if req.Mode == "" {
		req.Mode = "text"
	}
	if req.Label == "" && req.Mode == "text" {
		req.Label = "St"
	}
	if req.Size == 0 {
		req.Size = 1024
	}
	if req.Accent == "" {
		req.Accent = "#FF5500"
	}
	if req.Radius == "" {
		req.Radius = "medium"
	}
	showShadow := boolDefault(req.Shadow, true)

	s := req.Size
	borderW := max(4, round(float64(s)*0.035))
	shadowOff := round(float64(s) * 0.04)

	rFrac, ok := radiusMap[req.Radius]
	if !ok {
		rFrac = 0.06
	}
	r := round(float64(s) * rFrac)

	totalSize := s
	if showShadow {
		totalSize = s + shadowOff
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`+"\n", totalSize, totalSize, totalSize, totalSize)
	b.WriteString(`  <defs><style>@import url('https://fonts.googleapis.com/css2?family=Space+Mono:wght@400;700&amp;display=swap');</style></defs>` + "\n")

	half := borderW / 2

	// Hard shadow
	if showShadow {
		fmt.Fprintf(&b, `  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" fill="%s"/>`+"\n",
			shadowOff, shadowOff, s, s, r, Colors.Dark)
	}

	// Main square
	fmt.Fprintf(&b, `  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" fill="%s" stroke="%s" stroke-width="%d"/>`+"\n",
		half, half, s-borderW, s-borderW, r, req.Accent, Colors.Border, borderW)

	center := s / 2

	if req.Mode == "text" {
		label := req.Label
		fontFrac := 0.5
		if utf8.RuneCountInString(label) > 2 {
			fontFrac = 0.35
		}
		fontSize := round(float64(s) * fontFrac)
		fmt.Fprintf(&b, `  <text x="%d" y="%d" font-family="%s" font-size="%d" font-weight="700" fill="%s" text-anchor="middle" dominant-baseline="central">%s</text>`+"\n",
			center, center, Colors.Font, fontSize, Colors.Dark, escXML(label))
	} else {
		path, ok := RemixIcons[req.Icon]
		if ok {
			iconSize := round(float64(s) * 0.55)
			iconScale := float64(iconSize) / 24.0
			iconOffset := round(float64(s-iconSize) / 2.0)
			fmt.Fprintf(&b, `  <g transform="translate(%d,%d) scale(%.4f)">`+"\n", iconOffset, iconOffset, iconScale)
			fmt.Fprintf(&b, `    <path d="%s" fill="%s"/>`+"\n", path, Colors.Dark)
			b.WriteString("  </g>\n")
		}
	}

	b.WriteString("</svg>")
	return b.String()
}
