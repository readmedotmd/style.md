// Package coremd provides headless UI components with minimal base styles, built on gui.md.
//
// Components render semantic HTML with data-* attributes for state (data-variant,
// data-active, data-status, data-size, etc.) so themes can target them purely
// through CSS selectors. Include styles.css for functional base styles.
//
// The package includes layout primitives (Stack, HStack, Grid, Card, Badge, Divider),
// typography helpers (Heading, Paragraph, CodeBlock, Link, Image), and 120+ application
// components (Button, forms, modals, tables, navigation, panels, chat, settings, git,
// overlays) — everything a UI needs without external CSS.
//
// Use core.md directly for fully custom designs, or layer a theme on top:
//
//	<link rel="stylesheet" href="core.md/styles.css">
//	<link rel="stylesheet" href="industrial.md/theme.css">
package coremd
