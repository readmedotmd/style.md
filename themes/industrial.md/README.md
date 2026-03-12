<p align="center">
  <img src="../../design/industrial-banner.png" width="900" alt="industrial.md" />
</p>

<p align="center">
  Industrial monospace theme for core.md — Space Mono, hard shadows, bold accents.
</p>

---

## What is industrial.md?

**industrial.md** is a CSS-only theme for [core.md](../../core.md) components. It applies an industrial monospace design language: Space Mono typography, `#FF5500` orange accents, hard shadows, 2px borders, uppercase labels, and full dark mode.

### Usage

Use core.md Go components and activate the Industrial theme with CSS:

```html
<link rel="stylesheet" href="core.md/styles.css">
<link rel="stylesheet" href="themes/industrial.md/theme.css">
<html data-theme="industrial">
```

```go
import coremd "github.com/readmedotmd/style.md/core.md"

// Same Go code — the theme is applied purely through CSS
btn := coremd.Button(coremd.ButtonProps{Variant: "primary"}, gui.Text("Deploy"))
```

No Go wrapper imports needed. Switch themes at runtime by changing `data-theme` on `<html>`.

### Dark mode

Dark mode activates automatically via `prefers-color-scheme: dark`, or force it with:

```html
<html data-theme="industrial" data-mode="dark">
```

## Design Language

| Element | Treatment |
|---------|-----------|
| **Typography** | Space Mono monospace, uppercase headings, 0.04em letter-spacing |
| **Accent** | `#FF5500` orange, used for primary actions and highlights |
| **Borders** | 2px solid, high contrast (`#1A1A1A` light / `#444444` dark) |
| **Shadows** | Hard offset shadows (`3px 3px 0`), no blur |
| **Buttons** | Uppercase, bold, hard shadows with press animation |
| **Cards** | 2px borders with hard shadow on hover |
| **Badges** | Rectangular (3px radius), uppercase, bordered |
| **Lists** | Square bullet points |
| **Links** | Underlined, bold, thicker on hover |
| **Blockquotes** | 4px orange left border, uppercase |
| **Tables** | Dark header bar with orange column names |
| **Dividers** | 2px solid rules |

## Theme Tokens

industrial.md overrides all `--core-*` CSS properties under `[data-theme="industrial"]` and adds its own:

```css
[data-theme="industrial"] {
  --core-font: 'Space Mono', monospace;
  --core-accent: #FF5500;
  --core-border: #1A1A1A;
  --core-radius: 4px;

  --ind-dark: #2D333B;
  --ind-shadow-sm: 3px 3px 0 var(--core-border);
  --ind-shadow-md: 4px 4px 0 var(--core-border);
  --ind-shadow-lg: 6px 6px 0 var(--core-border);
}
```

## Files

```
industrial.md/
├── theme.css          CSS-only theme, scoped under [data-theme="industrial"]
└── examples/
    └── showcase.html  Interactive component showcase
```

## Showcase

<p align="center">
  <img src="../../design/screenshots/industrial-showcase-light.png" width="900" alt="industrial.md showcase" />
</p>

---

<p align="center">
  <strong>industrial.md</strong> is part of the <a href="https://github.com/readmedotmd">readme.md</a> project.
</p>
