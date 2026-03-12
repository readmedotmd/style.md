<p align="center">
  <img src="../../design/devbox-banner.png" width="900" alt="devbox.md" />
</p>

<p align="center">
  Developer tools theme for core.md — Inter typography, green accents, soft shadows, dark-first.
</p>

---

## What is devbox.md?

**devbox.md** is a CSS-only theme for [core.md](../../core.md) components. It applies a modern developer-tools design language: Inter sans-serif typography, `#22C55E` green accents, soft rounded corners, subtle layered shadows, and a dark-first aesthetic.

### Usage

Use core.md Go components and activate the Devbox theme with CSS:

```html
<link rel="stylesheet" href="core.md/styles.css">
<link rel="stylesheet" href="themes/devbox.md/theme.css">
<html data-theme="devbox">
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
<html data-theme="devbox" data-mode="dark">
```

## Design Language

| Element | Treatment |
|---------|-----------|
| **Typography** | Inter sans-serif, -0.01em tracking, 700 headings, clean hierarchy |
| **Accent** | `#22C55E` green, used for brand, primary actions, and active states |
| **Borders** | 1px solid, subtle (`#E5E7EB` light / `#374151` dark) |
| **Shadows** | Layered soft shadows (`0 4px 6px -1px`), no hard offsets |
| **Buttons** | Medium weight, soft rounded (6px), subtle shadow on hover |
| **Cards** | 1px borders, 8px radius, shadow lift on hover |
| **Badges** | Pill-shaped (9999px radius), no borders, tinted backgrounds |
| **Lists** | Disc bullet points |
| **Links** | No underline by default, underline on hover, green color |
| **Blockquotes** | 3px green left border, muted text |
| **Tables** | Surface header, 600 weight column names |
| **Dividers** | 1px solid, minimal |
| **Active state** | Left green border + green-tinted background |
| **Status dots** | Glowing green for running, pulse animation for starting |

## Theme Tokens

devbox.md overrides all `--core-*` CSS properties under `[data-theme="devbox"]` and adds its own:

```css
[data-theme="devbox"] {
  --core-font: 'Inter', system-ui, sans-serif;
  --core-font-mono: 'JetBrains Mono', ui-monospace, monospace;
  --core-accent: #22C55E;
  --core-border: #E5E7EB;
  --core-radius: 8px;

  --dbx-sidebar-bg: #F3F4F6;
  --dbx-active-bg: rgba(34, 197, 94, 0.08);
  --dbx-active-border: #22C55E;
  --dbx-shadow-sm: 0 1px 2px rgba(0, 0, 0, 0.05);
  --dbx-shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.07), ...;
  --dbx-shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.08), ...;
}
```

## Files

```
devbox.md/
├── theme.css          CSS-only theme, scoped under [data-theme="devbox"]
└── examples/
    └── showcase.html  Interactive component showcase
```

## Showcase

<p align="center">
  <img src="../../design/screenshots/devbox-showcase-light.png" width="900" alt="devbox.md showcase" />
</p>

---

<p align="center">
  <strong>devbox.md</strong> is part of the <a href="https://github.com/readmedotmd">readme.md</a> project.
</p>
