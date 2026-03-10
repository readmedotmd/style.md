// Package devboxmd provides the Devbox theme for core.md components.
//
// It wraps every core.md component (including layout primitives like Stack, Grid, Card,
// Badge, Heading, Link, Image) with pre-configured CSS class names that implement the
// Devbox design system: Inter/system sans-serif typography, green accents (#22C55E),
// soft rounded corners, subtle shadows, and a dark-first developer-tools aesthetic.
//
// Two usage modes:
//
// CSS-only — use core.md Go components with theme.css for data-attribute styling:
//
//	<link rel="stylesheet" href="core.md/styles.css">
//	<link rel="stylesheet" href="devbox.md/theme.css">
//
// Go wrappers — import this package for components pre-styled with BEM classes:
//
//	import devboxmd "github.com/readmedotmd/devbox.md"
//
//	btn := devboxmd.Button(devboxmd.ButtonProps{
//	    Variant: devboxmd.ButtonPrimary,
//	}, gui.Text("Click me"))
//
// Include styles.css (BEM) or theme.css (data-attributes) in your HTML.
package devboxmd
