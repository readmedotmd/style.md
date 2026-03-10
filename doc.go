// Package stylemd is the module root for the style.md design system.
//
// style.md is a themeable UI component system split into two layers:
//
//   - core.md — headless components with minimal base styles and data-* attributes
//   - themes (e.g. industrial.md) — CSS-only design layers that target those attributes
//
// Import core.md for components and a theme package for styling:
//
//	import "github.com/readmedotmd/core.md"
//	import "github.com/readmedotmd/industrial.md"
package stylemd
