<p align="center">
  <img src="banner.png" width="900" alt="style.md" />
</p>

<p align="center">
  A themeable component library for Go, built on gui.md.
</p>

---

## What is style.md?

**style.md** is a UI component system split into two layers:

1. **[core.md](./core.md)** — Headless components with minimal base styles. Renders semantic HTML with `data-*` attributes for state. No opinions on visual design.
2. **Themes** — CSS-only layers that target the same `data-*` selectors to apply a complete design language. Swap a `<link>` tag to switch themes at runtime.

The first theme is **[industrial.md](./industrial.md)** — monospace typography, bold orange accents, hard shadows, high contrast.

```
┌─────────────────────────────────────────────────┐
│  Your App (Go)                                  │
│    imports core.md components                   │
│    uses data-* attributes for state             │
├─────────────────────────────────────────────────┤
│  core.md/styles.css      │  industrial.md/      │
│  (minimal defaults)      │  theme.css           │
│                          │  (design layer)      │
└─────────────────────────────────────────────────┘
```

## Quick Start

### Use headless components with base styles

```go
import (
    gui "github.com/readmedotmd/gui.md"
    coremd "github.com/readmedotmd/core.md"
)

func App() gui.Node {
    return coremd.Stack("lg",
        coremd.Heading(1, "", gui.Text("Dashboard")),
        coremd.Card(coremd.CardProps{},
            coremd.HStack("md",
                coremd.Badge("", coremd.BadgeSuccess, "Running"),
                coremd.Paragraph("", gui.Text("All systems operational.")),
            ),
        ),
        coremd.Button(coremd.ButtonProps{Variant: "primary"}, gui.Text("Deploy")),
    )
}
```

```html
<link rel="stylesheet" href="core.md/styles.css">
```

### Apply a theme

Add a single CSS file to switch from base styles to a full design system:

```html
<link rel="stylesheet" href="core.md/styles.css">
<link rel="stylesheet" href="industrial.md/theme.css">
```

No HTML or Go code changes required.

### Switch themes at runtime

```html
<link rel="stylesheet" href="core.md/styles.css" id="base">
<link rel="stylesheet" href="" id="theme">

<select onchange="document.getElementById('theme').href = this.value">
  <option value="">Base</option>
  <option value="industrial.md/theme.css">Industrial</option>
</select>
```

## Packages

### core.md — Headless Components

60+ components with `data-*` attributes and minimal base styles. Includes layout primitives so your UI never needs external CSS.

| Category        | Components |
|-----------------|------------|
| **Primitives**  | Stack, HStack, Grid, Center, Spacer, Card, Badge, Divider, Heading, Paragraph, CodeBlock, InlineCode, Link, Image, UnorderedList, OrderedList, Quote, Muted, Mono, Truncate, SrOnly |
| **Buttons**     | Button (primary, danger, toolbar; medium, small) |
| **Forms**       | FormGroup, TextInput, TextArea, SelectInput, Checkbox, FeatureRow, VariableRow, ErrorMessage, SuccessMessage |
| **Input**       | ChatInput, AutocompletePopup, MessageQueue, SearchInputField |
| **Display**     | MessageBubble, ThinkingIndicator, ThinkingCollapsible, ToolBadge, QuestionPrompt, StatusBadge, StatusDot, LabelBadge, UsageBadge, DiffViewer, DataTable, EmptyState, ClusterStatsBar |
| **Lists**       | ConversationItem, InstanceCard, ServiceRow, RunnerRow, FileTree |
| **Navigation**  | NavLink, TabBar, BottomTabBar |
| **Overlay**     | SearchOverlay, ContextMenu |
| **Panels**      | ServicesPanel, RunnerPanel, GitPanel, SkillsPanel, TerminalPanel, FileBrowser |
| **Layout**      | AppShell, Navbar, Sidebar, Panel, Modal, ModalBackdrop, DragHandle |
| **Pages**       | LoginPage, SetupWizard, DashboardPage, SettingsCard |
| **Utility**     | Spinner, Icon |

**CSS primitives** in `styles.css` cover typography (h1-h6, p, code, pre, blockquote, kbd, mark), links, lists, images (`data-rounded`, `data-avatar`), layout (`data-stack`, `data-hstack`, `data-grid`, `data-align`, `data-justify`, `data-wrap`, `data-center`, `data-spacer`), cards (`data-card`), badges (`data-badge`), dividers, and utilities (`data-truncate`, `data-muted`, `data-mono`, `data-sr-only`).

### industrial.md — Industrial Monospace Theme

A CSS-only theme layer plus Go wrappers that re-export every core.md component with pre-configured BEM class names.

- Space Mono typography
- `#FF5500` orange accents
- Hard shadows, 2px borders
- Uppercase headings, labels, and badges
- Square bullet points, thick dividers
- Full dark mode support

**Two ways to use it:**

1. **CSS-only** — Load `core.md/styles.css` + `industrial.md/theme.css`. Use core.md Go components.
2. **Go wrappers** — Import `industrial.md` directly. Components come pre-styled with BEM classes.

## Data Attributes

Components communicate state through `data-*` attributes, which themes target via CSS:

| Attribute | Values | Used by |
|-----------|--------|---------|
| `data-variant` | `primary`, `danger`, `toolbar` | Button |
| `data-size` | `small`, `large` | Button, Spinner |
| `data-active` | `true` | NavLink, TabBar, ConversationItem |
| `data-status` | `running`, `stopped`, `starting`, `pending`, `error` | StatusBadge, StatusDot |
| `data-error` | `true` | TextInput |
| `data-streaming` | `true` | MessageBubble, ChatInput |
| `data-open` | `true` | Sidebar |
| `data-expanded` | `true` | Panel |
| `data-stack` | `xs`, `sm`, `md`, `lg`, `xl`, `none` | Stack layout |
| `data-hstack` | `xs`, `sm`, `md`, `lg`, `xl`, `none` | HStack layout |
| `data-grid` | `1`-`6` | Grid layout |
| `data-card` | `true`, `surface`, `flush` | Card |
| `data-badge` | `true`, `accent`, `success`, `danger`, `warning` | Badge |
| `data-align` | `start`, `center`, `end`, `stretch`, `baseline` | Alignment modifier |
| `data-justify` | `start`, `center`, `end`, `between`, `around`, `evenly` | Justification modifier |

## CSS Custom Properties

Themes override these tokens defined in `core.md/styles.css`:

```css
:root {
  --core-font:       system-ui, sans-serif;
  --core-font-mono:  ui-monospace, monospace;
  --core-text:       #1a1a1a;
  --core-text-muted: #6b7280;
  --core-bg:         #ffffff;
  --core-surface:    #f9fafb;
  --core-border:     #d1d5db;
  --core-accent:     #3b82f6;
  --core-danger:     #ef4444;
  --core-success:    #22c55e;
  --core-warning:    #f59e0b;
  --core-radius:     6px;
  --core-space:      8px;
  --core-transition: 150ms ease;
}
```

Dark mode activates via `prefers-color-scheme: dark` or `data-theme="dark"` on `<html>`.

## Creating a Theme

A theme is a CSS file that:

1. Overrides `--core-*` custom properties (fonts, colors, spacing)
2. Targets `data-*` attribute selectors to add decorative styles
3. Optionally adds theme-specific tokens

```css
/* mytheme.css */
:root {
  --core-font: 'Inter', sans-serif;
  --core-accent: #8b5cf6;
  --core-radius: 12px;
}

button[data-variant="primary"] {
  box-shadow: 0 4px 12px color-mix(in srgb, var(--core-accent) 40%, transparent);
}

[data-card] {
  box-shadow: 0 2px 8px rgba(0,0,0,0.08);
}
```

## Theme Switcher Demo

The [`examples/theme-switcher.html`](./examples/theme-switcher.html) file provides an interactive showcase with dropdowns to switch between Base and Industrial themes, plus light/dark mode toggle.

<p align="center">
  <img src="design/screenshots/theme-switcher-base-primitives.png" width="440" alt="Base theme" />
  <img src="design/screenshots/theme-switcher-industrial-primitives.png" width="440" alt="Industrial theme" />
</p>

## Project Structure

```
style.md/
├── core.md/                   Headless component library
│   ├── styles.css             Base styles + layout primitives
│   ├── primitives.go          Stack, Grid, Card, Badge, Heading, Link, Image, ...
│   ├── button.go              Button component
│   ├── form.go                Form components
│   ├── layout.go              App shell, navbar, sidebar, modal
│   ├── display.go             Status badges, diff viewer, data table
│   ├── ...                    (14 Go files total)
│   └── examples/showcase.html Core.md showcase
├── industrial.md/             Industrial monospace theme
│   ├── theme.css              CSS-only theme (targets data-* selectors)
│   ├── styles.css             BEM class-based stylesheet
│   ├── tokens.go              260+ CSS class constants
│   ├── primitives.go          Re-exported primitives with theme wrappers
│   ├── ...                    (14 Go files total)
│   └── examples/showcase.html Industrial showcase
├── examples/
│   └── theme-switcher.html    Interactive theme switching showcase
├── generate/                  SVG banner & icon generation server
├── cmd/                       CLI tools
└── design/                    Screenshots and assets
```

## Testing

Components are testable with `gui.md/testing`:

```go
func TestPrimaryButton(t *testing.T) {
    s := guitesting.Render(coremd.Button(coremd.ButtonProps{Variant: "primary"}, gui.Text("Save")))
    btn := s.GetByRole("button")
    guitesting.AssertNode(t, btn).HasText("Save")
}
```

---

<p align="center">
  <strong>style.md</strong> is part of the <a href="https://github.com/readmedotmd">readme.md</a> project.
</p>
