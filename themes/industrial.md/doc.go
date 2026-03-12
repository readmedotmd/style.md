// Package industrialmd provides the Industrial Monospace theme for core.md components.
//
// It wraps every core.md component (including layout primitives like Stack, Grid, Card,
// Badge, Heading, Link, Image) with pre-configured CSS class names that implement the
// industrial design system: Space Mono typography, bold orange accents (#FF5500), hard
// shadows, 2px borders, uppercase labels, and full dark mode support.
//
// Two usage modes:
//
// CSS-only — use core.md Go components with theme.css for data-attribute styling:
//
//	<link rel="stylesheet" href="core.md/styles.css">
//	<link rel="stylesheet" href="themes/industrial.md/theme.css">
//
// Go wrappers — import this package for components pre-styled with BEM classes:
//
//	import industrialmd "github.com/readmedotmd/style.md/themes/industrial.md"
//
//	btn := industrialmd.Button(industrialmd.ButtonProps{
//	    Variant: industrialmd.ButtonPrimary,
//	}, gui.Text("Click me"))
//
// Include styles.css (BEM) or theme.css (data-attributes) in your HTML.
package industrialmd
