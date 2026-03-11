package generate

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type bannerCtx struct {
	w, h      int
	scale     float64
	name      string
	tagline   string
	desc      string
	version   string
	accent    string
	tags      []string
	showDots  bool
	showBadge bool
	showTags  bool
}

// Banner generates an SVG banner string from the given request.
func Banner(req BannerRequest) string {
	if req.Name == "" {
		req.Name = "project.md"
	}
	if req.Accent == "" {
		req.Accent = "#FF5500"
	}
	if req.Layout == "" {
		req.Layout = "card"
	}
	if req.Size == "" {
		req.Size = "1800x560"
	}

	showDots := boolDefault(req.ShowDots, true)
	showBadge := boolDefault(req.ShowBadge, true)
	showTags := boolDefault(req.ShowTags, true)

	w, h := parseDimensions(req.Size)
	scale := float64(h) / 280.0

	ctx := &bannerCtx{
		w: w, h: h, scale: scale,
		name: req.Name, tagline: req.Tagline, desc: req.Desc,
		version: req.Version, accent: req.Accent, tags: req.Tags,
		showDots: showDots, showBadge: showBadge, showTags: showTags,
	}

	var b strings.Builder
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`+"\n", w, h, w, h)
	b.WriteString(`  <defs><style>@import url('https://fonts.googleapis.com/css2?family=Space+Mono:wght@400;700&amp;display=swap');</style></defs>` + "\n")

	// Page background
	fmt.Fprintf(&b, `  <rect width="%d" height="%d" fill="%s"/>`+"\n", w, h, Colors.PageBg)

	switch req.Layout {
	case "terminal":
		b.WriteString(renderTerminalBanner(ctx))
	case "minimal":
		b.WriteString(renderMinimalBanner(ctx))
	default:
		b.WriteString(renderCardBanner(ctx))
	}

	b.WriteString("</svg>")
	return b.String()
}

func renderCardBanner(ctx *bannerCtx) string {
	var s strings.Builder

	pad := round(24 * ctx.scale)
	cardX := pad
	cardY := pad
	cardW := ctx.w - pad*2
	cardH := ctx.h - pad*2
	borderW := maxInt(4, round(4*ctx.scale))
	shadowOff := round(5 * ctx.scale)
	r := round(6 * ctx.scale)

	// Management-style: dark header ~38%
	headerH := round(float64(cardH) * 0.38)

	// Card shadow
	fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" fill="%s"/>`+"\n",
		cardX+shadowOff, cardY+shadowOff, cardW, cardH, r, Colors.Border)

	// Clip path for card shape
	fmt.Fprintf(&s, `  <clipPath id="card-clip"><rect x="%d" y="%d" width="%d" height="%d" rx="%d"/></clipPath>`+"\n",
		cardX, cardY, cardW, cardH, r)

	// Card body (white)
	fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" fill="%s"/>`+"\n",
		cardX, cardY, cardW, cardH, r, Colors.Surface)

	// Dark header clipped to card
	fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" fill="%s" clip-path="url(#card-clip)"/>`+"\n",
		cardX, cardY, cardW, headerH, Colors.Dark)

	// Header content
	headerPadX := cardX + round(28*ctx.scale)
	labelSize := round(10 * ctx.scale)
	titleSize := round(28 * ctx.scale)
	labelY := cardY + round(float64(headerH)*0.35)
	titleY := cardY + round(float64(headerH)*0.72)

	// Muted label
	labelText := "PROJECT"
	if ctx.tagline != "" {
		labelText = strings.ToUpper(ctx.tagline)
	}
	fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="#888888" font-weight="700" letter-spacing="2">%s</text>`+"\n",
		headerPadX, labelY, Colors.Font, labelSize, escXML(labelText))

	// Title — NAME.ext format
	displayName := formatDisplayName(ctx.name)
	fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="%s" font-weight="700" letter-spacing="1">%s</text>`+"\n",
		headerPadX, titleY, Colors.Font, titleSize, ctx.accent, escXML(displayName))

	// Version badge in header (right side)
	if ctx.showBadge && ctx.version != "" {
		badgeFontSize := round(10 * ctx.scale)
		countSize := round(24 * ctx.scale)
		badgeRightX := cardX + cardW - round(28*ctx.scale)
		fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="#888888" font-weight="700" letter-spacing="2" text-anchor="end">VERSION</text>`+"\n",
			badgeRightX, labelY, Colors.Font, badgeFontSize)
		fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="%s" font-weight="700" letter-spacing="1" text-anchor="end">%s</text>`+"\n",
			badgeRightX, titleY, Colors.Font, countSize, Colors.Surface, escXML(strings.ToUpper(ctx.version)))
	}

	// Status dot
	if ctx.showDots {
		dotR := round(4 * ctx.scale)
		dotY := titleY - round(float64(titleSize)*0.35)
		dotStart := headerPadX + utf8.RuneCountInString(displayName)*round(19*ctx.scale) + round(20*ctx.scale)
		fmt.Fprintf(&s, `  <circle cx="%d" cy="%d" r="%d" fill="%s"/>`+"\n",
			dotStart, dotY, dotR, Colors.DotGreen)
	}

	// White body content
	contentX := cardX + round(28*ctx.scale)
	contentTop := cardY + headerH + round(16*ctx.scale)

	// Description
	if ctx.desc != "" {
		descSize := round(12 * ctx.scale)
		descY := contentTop + descSize
		fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="%s" letter-spacing="0.5">%s</text>`+"\n",
			contentX, descY, Colors.Font, descSize, Colors.Muted, escXML(ctx.desc))
	}

	// Tags row
	if ctx.showTags && len(ctx.tags) > 0 {
		tagFontSize := round(10 * ctx.scale)
		tagH := round(22 * ctx.scale)
		tagY := cardY + cardH - round(22*ctx.scale)
		tagR := round(3 * ctx.scale)
		tagGap := round(8 * ctx.scale)
		tagBgs := []string{ctx.accent, "#007AFF", "#34C759", "#FFCC00", "#AF52DE", "#FF3B30"}
		tagFgs := []string{"#FFFFFF", "#FFFFFF", "#FFFFFF", "#1A1A1A", "#FFFFFF", "#FFFFFF"}
		tagX := contentX

		for i, tag := range ctx.tags {
			tagText := formatUpperWithExt(tag)
			tagBg := tagBgs[i%len(tagBgs)]
			tagFg := tagFgs[i%len(tagFgs)]
			tagW := maxInt(utf8.RuneCountInString(tagText)*round(8*ctx.scale)+round(20*ctx.scale), round(70*ctx.scale))
			fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" fill="%s" stroke="%s" stroke-width="%d"/>`+"\n",
				tagX, tagY-tagH, tagW, tagH, tagR, tagBg, Colors.Border, borderW)
			fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="%s" text-anchor="middle" font-weight="700" letter-spacing="1">%s</text>`+"\n",
				tagX+tagW/2, tagY-tagH/2+round(float64(tagFontSize)*0.35), Colors.Font, tagFontSize, tagFg, escXML(tagText))
			tagX += tagW + tagGap
		}
	}

	// Orange accent bar at bottom
	barH := round(6 * ctx.scale)
	fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" fill="%s" clip-path="url(#card-clip)"/>`+"\n",
		cardX, cardY+cardH-barH, cardW, barH, ctx.accent)

	// Card border (on top)
	fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" fill="none" stroke="%s" stroke-width="%d"/>`+"\n",
		cardX, cardY, cardW, cardH, r, Colors.Border, borderW)
	// Header bottom border
	fmt.Fprintf(&s, `  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="%s" stroke-width="%d"/>`+"\n",
		cardX, cardY+headerH, cardX+cardW, cardY+headerH, Colors.Border, borderW)

	return s.String()
}

func renderTerminalBanner(ctx *bannerCtx) string {
	var s strings.Builder

	pad := round(20 * ctx.scale)
	cardX := pad
	cardY := pad
	cardW := ctx.w - pad*2
	cardH := ctx.h - pad*2
	headerH := round(32 * ctx.scale)
	borderW := maxInt(2, round(2*ctx.scale))
	shadowOff := round(5 * ctx.scale)
	r := round(6 * ctx.scale)

	// Shadow
	fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" fill="%s"/>`+"\n",
		cardX+shadowOff, cardY+shadowOff, cardW, cardH, r, Colors.Dark)

	// Terminal body
	fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" rx="%d" fill="%s" stroke="%s" stroke-width="%d"/>`+"\n",
		cardX, cardY, cardW, cardH, r, Colors.Dark, Colors.Border, borderW)

	// Header bar
	fmt.Fprintf(&s, `  <line x1="%d" y1="%d" x2="%d" y2="%d" stroke="#333333" stroke-width="%d"/>`+"\n",
		cardX, cardY+headerH, cardX+cardW, cardY+headerH, borderW)

	// Status dots
	if ctx.showDots {
		dotR := round(5 * ctx.scale)
		dotY := cardY + round(float64(headerH)/2)
		dotStart := cardX + round(16*ctx.scale)
		dotGap := round(16 * ctx.scale)
		fmt.Fprintf(&s, `  <circle cx="%d" cy="%d" r="%d" fill="%s"/>`+"\n", dotStart, dotY, dotR, Colors.DotRed)
		fmt.Fprintf(&s, `  <circle cx="%d" cy="%d" r="%d" fill="%s"/>`+"\n", dotStart+dotGap, dotY, dotR, Colors.DotYellow)
		fmt.Fprintf(&s, `  <circle cx="%d" cy="%d" r="%d" fill="%s"/>`+"\n", dotStart+dotGap*2, dotY, dotR, Colors.DotGreen)
	}

	// Header title
	headerTextSize := round(10 * ctx.scale)
	fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="#666666" font-weight="700" text-anchor="middle" letter-spacing="2">%s — TERMINAL</text>`+"\n",
		ctx.w/2, cardY+round(float64(headerH)/2)+round(float64(headerTextSize)*0.35), Colors.Font, headerTextSize, escXML(strings.ToUpper(ctx.name)))

	// Terminal content
	lineH := round(20 * ctx.scale)
	cX := cardX + round(24*ctx.scale)
	cY := cardY + headerH + round(24*ctx.scale)
	fontSize := round(13 * ctx.scale)

	// Prompt line
	fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="%s" font-weight="700">▸ %s</text>`+"\n",
		cX, cY, Colors.Font, fontSize, ctx.accent, escXML(ctx.name))
	cY += lineH

	// Tagline
	if ctx.tagline != "" {
		fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="#AAAAAA"># %s</text>`+"\n",
			cX, cY, Colors.Font, fontSize, escXML(ctx.tagline))
		cY += lineH
	}

	// Blank line
	cY += round(8 * ctx.scale)

	// Description
	if ctx.desc != "" {
		fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="#666666">%s</text>`+"\n",
			cX, cY, Colors.Font, round(11*ctx.scale), escXML(ctx.desc))
		cY += lineH
	}

	// Version and tags at bottom
	badgeY := cardY + cardH - round(20*ctx.scale)
	badgeH := round(20 * ctx.scale)
	badgeFontSize := round(10 * ctx.scale)

	if ctx.showBadge && ctx.version != "" {
		vw := utf8.RuneCountInString(ctx.version)*round(8*ctx.scale) + round(20*ctx.scale)
		fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" rx="3" fill="none" stroke="%s" stroke-width="%d"/>`+"\n",
			cX, badgeY-badgeH, vw, badgeH, ctx.accent, borderW)
		fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="%s" text-anchor="middle" font-weight="700" letter-spacing="1">%s</text>`+"\n",
			cX+vw/2, badgeY-badgeH/2+round(float64(badgeFontSize)*0.35), Colors.Font, badgeFontSize, ctx.accent, escXML(ctx.version))
	}

	if ctx.showTags && len(ctx.tags) > 0 {
		txPos := cX
		if ctx.showBadge && ctx.version != "" {
			txPos = cX + utf8.RuneCountInString(ctx.version)*round(8*ctx.scale) + round(20*ctx.scale) + round(8*ctx.scale)
		}
		for _, tag := range ctx.tags {
			tagText := formatUpperWithExt(tag)
			tw := utf8.RuneCountInString(tagText)*round(8*ctx.scale) + round(20*ctx.scale)
			fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" rx="3" fill="none" stroke="%s" stroke-width="%d"/>`+"\n",
				txPos, badgeY-badgeH, tw, badgeH, ctx.accent, borderW)
			fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="%s" text-anchor="middle" font-weight="700" letter-spacing="1">%s</text>`+"\n",
				txPos+tw/2, badgeY-badgeH/2+round(float64(badgeFontSize)*0.35), Colors.Font, badgeFontSize, ctx.accent, escXML(tagText))
			txPos += tw + round(8*ctx.scale)
		}
	}

	// Cursor blink
	fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" fill="%s" opacity="0.7"/>`+"\n",
		cX, cY-fontSize, round(8*ctx.scale), round(16*ctx.scale), ctx.accent)

	return s.String()
}

func renderMinimalBanner(ctx *bannerCtx) string {
	var s strings.Builder

	borderW := maxInt(2, round(2*ctx.scale))

	// Accent left bar
	barW := round(6 * ctx.scale)
	fmt.Fprintf(&s, `  <rect x="0" y="0" width="%d" height="%d" fill="%s"/>`+"\n", barW, ctx.h, ctx.accent)

	// Bottom border
	fmt.Fprintf(&s, `  <line x1="0" y1="%d" x2="%d" y2="%d" stroke="%s" stroke-width="%d"/>`+"\n",
		ctx.h-borderW/2, ctx.w, ctx.h-borderW/2, Colors.Dark, borderW)

	// Content
	contentX := barW + round(40*ctx.scale)
	titleSize := round(48 * ctx.scale)
	titleY := round(float64(ctx.h)*0.42) + round(float64(titleSize)*0.3)

	fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" font-weight="700" fill="%s" letter-spacing="-1">%s</text>`+"\n",
		contentX, titleY, Colors.Font, titleSize, Colors.Dark, escXML(ctx.name))

	if ctx.tagline != "" {
		tagSize := round(13 * ctx.scale)
		fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="%s" font-weight="700" letter-spacing="2">%s</text>`+"\n",
			contentX, titleY+round(24*ctx.scale), Colors.Font, tagSize, ctx.accent, escXML(strings.ToUpper(ctx.tagline)))
	}

	if ctx.desc != "" {
		descSize := round(11 * ctx.scale)
		descY := titleY + round(46*ctx.scale)
		fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="%s">%s</text>`+"\n",
			contentX, descY, Colors.Font, descSize, Colors.Muted, escXML(ctx.desc))
	}

	// Version badge right side
	if ctx.showBadge && ctx.version != "" {
		badgeFontSize := round(10 * ctx.scale)
		vw := utf8.RuneCountInString(ctx.version)*round(8*ctx.scale) + round(20*ctx.scale)
		vx := ctx.w - round(40*ctx.scale) - vw
		vy := round(float64(ctx.h)/2) - round(11*ctx.scale)
		badgeH := round(22 * ctx.scale)
		fmt.Fprintf(&s, `  <rect x="%d" y="%d" width="%d" height="%d" rx="3" fill="none" stroke="%s" stroke-width="%d"/>`+"\n",
			vx, vy, vw, badgeH, Colors.Dark, borderW)
		fmt.Fprintf(&s, `  <text x="%d" y="%d" font-family="%s" font-size="%d" fill="%s" text-anchor="middle" font-weight="700" letter-spacing="1">%s</text>`+"\n",
			vx+vw/2, vy+badgeH/2+round(float64(badgeFontSize)*0.35), Colors.Font, badgeFontSize, Colors.Dark, escXML(strings.ToUpper(ctx.version)))
	}

	return s.String()
}

// formatDisplayName formats "store.md" as "STORE.md".
func formatDisplayName(name string) string {
	parts := strings.Split(name, ".")
	if len(parts) > 1 {
		return strings.ToUpper(strings.Join(parts[:len(parts)-1], ".")) + "." + strings.ToLower(parts[len(parts)-1])
	}
	return strings.ToUpper(name)
}
