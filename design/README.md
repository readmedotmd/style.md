# style.md Design Guidelines

A comprehensive guide for LLMs and developers building applications with the style.md industrial monospace design system.

## Table of Contents

- [Design Philosophy](#design-philosophy)
- [Color System](#color-system)
- [Dark Mode](#dark-mode)
- [Typography](#typography)
- [Spacing System](#spacing-system)
- [Borders & Shadows](#borders--shadows)
- [Component Reference](#component-reference)
- [Layout Patterns](#layout-patterns)
- [Do's and Don'ts](#dos-and-donts)
- [Code Examples](#code-examples)

---

## Design Philosophy

This design system delivers an industrial, equipment-inspired interface toolkit for building Go-based web applications with gui.md. The aesthetic draws from hardware control panels and professional audio/electronic equipment — functional, precise, and unapologetically mechanical.

### Core Principles

1. **Industrial Minimalism** - Clean, functional interfaces with purposeful elements. No decorative flourishes.

2. **Bold Orange Accents** - Orange (#FF5500) is the signature color, used sparingly for emphasis, highlights, and interactive elements.

3. **Monospace Typography** - Space Mono font throughout. UPPERCASE transforms for labels and headings. Wide letter-spacing creates that distinctive equipment label look.

4. **Hard Shadows** - No blur or soft edges. Shadows are solid offset shapes (2-6px), like physical objects casting shadows.

5. **High Contrast** - Primarily black (#1A1A1A) and white, with gray (#E8E8E8) backgrounds. Clear visual hierarchy.

6. **Functional Design** - Every element should serve a purpose. Interface components mirror physical hardware controls.

---

## Color System

### Core Palette

| Token            | Value                 | Usage                                           |
| ---------------- | --------------------- | ----------------------------------------------- |
| `--accent`       | `#FF5500`             | Primary accent, CTAs, highlights, active states |
| `--accent-hover` | `#E64D00`             | Hover state for orange elements                 |
| `--accent-light` | `#FF7733`             | Light orange for subtle emphasis                |
| `--accent-muted` | `rgba(255,85,0,0.15)` | Backgrounds for orange-themed alerts            |

### Neutrals

| Token        | Value     | Usage                                     |
| ------------ | --------- | ----------------------------------------- |
| `--black`    | `#1A1A1A` | Text, borders, dark backgrounds, shadows  |
| `--bg`       | `#E8E8E8` | Page background (swaps in dark mode)      |
| `--surface`  | `#F5F5F5` | Card header backgrounds (swaps)           |
| `--raised`   | `#FFFFFF` | Cards, inputs, raised surfaces (swaps)    |
| `--border`   | `#1A1A1A` | Component borders (swaps)                 |
| `--text`     | `#1A1A1A` | Primary text color (swaps)                |
| `--muted`    | `#CCCCCC` | Muted text, secondary labels (swaps)      |
| `--dark`     | `#2D333B` | Always-dark elements (navbar, panels)     |
| `--dark-text`| `#FFFFFF` | Text on always-dark elements              |
| `--dark-muted`| `#AAAAAA`| Muted text on always-dark elements        |
| `--divider`  | `#CCCCCC` | Divider lines, subtle borders (swaps)     |
| `--white`    | `#FFFFFF` | Static white (use `--raised` for surfaces)|

### Semantic Colors

| Token            | Value                   | Usage                               |
| ---------------- | ----------------------- | ----------------------------------- |
| `--danger`       | `#FF3333`               | Errors, danger actions, destructive |
| `--danger-dark`  | `#CC0000`               | Darker red for hover states         |
| `--danger-muted` | `rgba(255,51,51,0.15)`  | Error alert backgrounds             |
| `--success`      | `#00CC66`               | Success, active status              |
| `--success-dark` | `#00994D`               | Darker green for hover              |
| `--success-muted`| `rgba(0,204,102,0.15)`  | Success alert backgrounds           |
| `--blue`         | `#3399FF`               | Info, links                         |
| `--blue-dark`    | `#0066CC`               | Darker blue for hover               |
| `--blue-muted`   | `rgba(51,153,255,0.15)` | Info alert backgrounds              |
| `--warning`      | `#FFCC00`               | Warnings                            |
| `--warning-dark` | `#CC9900`               | Darker yellow                       |
| `--warning-muted`| `rgba(255,204,0,0.15)`  | Warning alert backgrounds           |

### Color Usage Guidelines

- **Primary actions**: Use `--accent` for the main CTA button
- **Destructive actions**: Use `--danger` for delete/revoke buttons
- **Status indicators**: Green=active, Orange=warning, Red=error, Gray=inactive
- **Text on dark backgrounds**: Use `--white` or `--accent` for emphasis
- **Code/technical values**: Display in `--accent` on dark backgrounds

---

## Dark Mode

The design system supports full dark mode via CSS custom property overrides. Dark mode activates automatically via `prefers-color-scheme: dark` or manually via `data-theme="dark"` on the root element.

### Activation

```html
<!-- Automatic: follows OS preference -->
<!-- No code needed, handled by @media query -->

<!-- Manual toggle via JavaScript -->
<button onclick="document.documentElement.setAttribute('data-theme', 'dark')">Dark</button>
<button onclick="document.documentElement.removeAttribute('data-theme')">Light</button>
```

### Token Overrides in Dark Mode

| Light Token      | Light Value | Dark Value | Notes                                |
| ---------------- | ----------- | ---------- | ------------------------------------ |
| `--bg`           | `#E8E8E8`  | `#1A1A1A`  | Page background                      |
| `--surface`      | `#F5F5F5`  | `#252525`  | Card headers, subtle backgrounds     |
| `--raised`       | `#FFFFFF`  | `#2A2A2A`  | Cards, inputs, raised surfaces       |
| `--border`       | `#1A1A1A`  | `#555555`  | Component borders                    |
| `--text`         | `#1A1A1A`  | `#E8E8E8`  | Primary text                         |
| `--muted`        | `#CCCCCC`  | `#777777`  | Secondary/muted text                 |
| `--dark`         | `#1A1A1A`  | `#111111`  | Always-dark elements (navbar, panels)|
| `--dark-text`    | `#FFFFFF`  | `#E8E8E8`  | Text on always-dark elements         |
| `--dark-muted`   | `#AAAAAA`  | `#888888`  | Muted text on always-dark elements   |
| `--divider`      | `#CCCCCC`  | `#444444`  | Divider lines, subtle borders        |

### Dark Mode Design Rules

- **Always-dark elements** (navbars, panel headers, terminal) use `--dark` / `--dark-text` tokens — these stay dark in both themes
- **Content surfaces** (cards, inputs, modals) use `--raised` / `--text` — these swap between light and dark
- **Accent color** (`--accent`: #FF5500) remains the same in both themes
- **Semantic colors** (danger, success, blue, warning) remain the same in both themes
- **Shadows** use `--shadow-color` which adjusts opacity for dark mode
- Never hardcode hex colors in component styles — always use CSS variables

---

## Typography

### Font Stack

```css
--font-mono: "Space Mono", ui-monospace, "SF Mono", "Cascadia Mono", "Consolas", monospace;
```

Always use the monospace font. Import Space Mono in your app entry:

```css
@import url("https://fonts.googleapis.com/css2?family=Space+Mono:wght@400;700&display=swap");
```

### Font Sizes

| Token         | Size      | Pixel | Usage                            |
| ------------- | --------- | ----- | -------------------------------- |
| `--text-xs`   | 0.6875rem | 11px  | Badges, small labels, timestamps |
| `--text-sm`   | 0.75rem   | 12px  | Labels, secondary text, code     |
| `--text-base` | 0.875rem  | 14px  | Body text, inputs                |
| `--text-md`   | 1rem      | 16px  | Large inputs, h4                 |
| `--text-lg`   | 1.25rem   | 20px  | h2, panel values                 |
| `--text-xl`   | 1.5rem    | 24px  | h1, page titles                  |
| `--text-2xl`  | 2rem      | 32px  | Large display numbers            |

### Letter Spacing

| Token              | Value  | Usage                          |
| ------------------ | ------ | ------------------------------ |
| `--tracking`       | 0.02em | Default body text              |
| `--tracking-wide`  | 0.05em | Buttons, headings              |
| `--tracking-wider` | 0.1em  | Labels, badges, uppercase text |

### Typography Patterns

**Headings**: Bold (700), uppercase, wide letter-spacing

```go
gui.H1(gui.Class("text-xl"),
    gui.Style("text-transform: uppercase; letter-spacing: var(--tracking-wide); font-weight: 700"),
)(
    gui.Text("PAGE TITLE"),
)
```

**Labels**: Regular weight, uppercase, wider letter-spacing, muted color

```go
gui.Span(gui.Class("text-xs"),
    gui.Style("text-transform: uppercase; letter-spacing: var(--tracking-wider); color: var(--muted)"),
)(
    gui.Text("LABEL TEXT"),
)
```

**Display Numbers**: Use tabular-nums for aligned digits

```go
gui.Span(gui.Style("font-variant-numeric: tabular-nums"))(
    gui.Text(fmt.Sprintf("%02d", count)),
)
```

---

## Spacing System

Based on a **4px unit** scale. Use CSS variables for consistency:

| Token        | Value   | Pixels |
| ------------ | ------- | ------ |
| `--space-1`  | 0.25rem | 4px    |
| `--space-2`  | 0.5rem  | 8px    |
| `--space-3`  | 0.75rem | 12px   |
| `--space-4`  | 1rem    | 16px   |
| `--space-5`  | 1.25rem | 20px   |
| `--space-6`  | 1.5rem  | 24px   |
| `--space-8`  | 2rem    | 32px   |
| `--space-10` | 2.5rem  | 40px   |
| `--space-12` | 3rem    | 48px   |

### Spacing Guidelines

- **Component padding**: Usually `--space-4` to `--space-6`
- **Gap between items**: `--space-2` to `--space-4`
- **Section margins**: `--space-6` to `--space-8`
- **Page padding**: `--space-8`

---

## Borders & Shadows

### Border Widths

- **Standard border**: `2px solid` - Used on cards, inputs, buttons
- **Thin border**: `1px solid` - Used for subtle dividers, badges

### Border Radius

| Token           | Value  | Usage                    |
| --------------- | ------ | ------------------------ |
| `--radius-sm`   | 3px    | Badges, small elements   |
| `--radius-md`   | 4px    | Buttons, inputs          |
| `--radius-lg`   | 6px    | Panels                   |
| `--radius-xl`   | 8px    | Cards                    |
| `--radius-full` | 9999px | Pills, circular elements |

### Hard Shadows

Shadows have **no blur** - they're solid offset shapes:

| Token         | Value                    | Usage                   |
| ------------- | ------------------------ | ----------------------- |
| `--shadow-sm` | `2px 2px 0 #1A1A1A`     | Small elements          |
| `--shadow-md` | `4px 4px 0 #1A1A1A`     | Cards (default)         |
| `--shadow-lg` | `6px 6px 0 #1A1A1A`     | Modals, prominent cards |

---

## Component Reference

All components are Go functions built on [gui.md](https://github.com/readmedotmd/gui.md). Import core.md for headless components and a theme for styling:

```go
import (
    gui "github.com/readmedotmd/gui.md"
    coremd "github.com/readmedotmd/core.md"
)
```

### Core Components

#### Button

Interactive buttons with variants and sizes.

```go
stylemd.Button(stylemd.ButtonProps{Variant: stylemd.ButtonPrimary},
    gui.Text("SUBMIT"),
)
stylemd.Button(stylemd.ButtonProps{},
    gui.Text("CANCEL"),
)
stylemd.Button(stylemd.ButtonProps{Variant: stylemd.ButtonDanger},
    gui.Text("DELETE"),
)
stylemd.Button(stylemd.ButtonProps{Variant: stylemd.ButtonGhost},
    gui.Text("LINK STYLE"),
)
stylemd.Button(stylemd.ButtonProps{Size: stylemd.ButtonSmall},
    gui.Text("SMALL"),
)
stylemd.Button(stylemd.ButtonProps{Disabled: true},
    gui.Text("DISABLED"),
)
```

| Field    | Type          | Default       | Description                    |
| -------- | ------------- | ------------- | ------------------------------ |
| Variant  | ButtonVariant | ButtonDefault | Visual style                   |
| Size     | ButtonSize    | ButtonMedium  | Button size                    |
| Disabled | bool          | false         | Shows spinner, disables button |
| OnClick  | func()        | nil           | Click handler                  |

Variants: `ButtonDefault`, `ButtonPrimary`, `ButtonDanger`, `ButtonGhost`, `ButtonToolbar`

#### TextInput

Text input field with error state.

```go
stylemd.TextInput(stylemd.TextInputProps{Placeholder: "Enter value"})
stylemd.TextInput(stylemd.TextInputProps{Error: true, Placeholder: "Invalid input"})
```

| Field       | Type   | Default | Description      |
| ----------- | ------ | ------- | ---------------- |
| Placeholder | string | ""      | Placeholder text |
| Value       | string | ""      | Current value    |
| Type        | string | "text"  | Input type       |
| Error       | bool   | false   | Shows red border |
| ID          | string | ""      | Element ID       |

#### Card (CSS pattern)

Container with optional header section. Compose with gui.md elements:

```go
gui.Div(gui.Class("card"))(
    gui.Div(gui.Class("card__header card__header--dark"))(gui.Text("TITLE")),
    gui.Div(gui.Class("card__body"))(gui.Text("Content here")),
)

gui.Div(gui.Class("card card--no-shadow"))(gui.Text("No shadow card"))
```

Header variants: `card__header--dark`, `card__header--orange`, `card__header--danger`, `card__header--success`

#### Panel (stylemd component)

Content panel with title, header actions, and body. Used for collapsible sections.

```go
stylemd.Panel(stylemd.PanelProps{Title: "DETAILS"}, headerActions,
    gui.Text("Panel content here"),
)
```

#### Panel (dark display — CSS pattern)

Dark display panel for key information (like equipment readout screens).

```go
gui.Div(gui.Class("display-panel"))(
    gui.Span(gui.Class("display-panel__label"))(gui.Text("STATUS")),
    gui.Span(gui.Class("display-panel__value display-panel__value--highlight"))(gui.Text("ACTIVE")),
)
```

| Class                             | Description       |
| --------------------------------- | ----------------- |
| `.display-panel`                  | Dark background   |
| `.display-panel__label`           | Muted label       |
| `.display-panel__value`           | White text value   |
| `.display-panel__value--highlight`| Orange text value  |

### Form Components

#### FormGroup

Container for form field with proper spacing.

```go
stylemd.FormGroup("USERNAME",
    stylemd.TextInput(stylemd.TextInputProps{ID: "username"}),
)
```

#### TextArea

Multi-line text input.

```go
stylemd.TextArea(stylemd.TextAreaProps{
    Placeholder: "Enter description...",
    Rows:        4,
})
```

#### SelectInput

Dropdown select with options.

```go
stylemd.SelectInput([]stylemd.SelectOption{
    {Label: "Production", Value: "prod", Selected: true},
    {Label: "Staging", Value: "staging"},
    {Label: "Development", Value: "dev"},
})
```

#### Checkbox

Toggle checkbox with label.

```go
stylemd.Checkbox(stylemd.CheckboxProps{
    Label:   "Enable notifications",
    Checked: true,
})
```

### Feedback Components

#### Badge (CSS pattern)

Small tag for labels and status.

```go
gui.Span(gui.Class("badge"))(gui.Text("DEFAULT"))
gui.Span(gui.Class("badge badge--orange"))(gui.Text("NEW"))
gui.Span(gui.Class("badge badge--success"))(gui.Text("ACTIVE"))
gui.Span(gui.Class("badge badge--danger"))(gui.Text("REVOKED"))
gui.Span(gui.Class("badge badge--outline"))(gui.Text("TAG"))
```

Variants: `badge--orange`, `badge--success`, `badge--danger`, `badge--info`, `badge--outline`

#### StatusBadge (stylemd component)

Status indicator with colored background.

```go
stylemd.StatusBadge(stylemd.StatusRunning)   // Green
stylemd.StatusBadge(stylemd.StatusStopped)   // Gray
stylemd.StatusBadge(stylemd.StatusStarting)  // Yellow
stylemd.StatusBadge(stylemd.StatusPending)   // Yellow outline
stylemd.StatusBadge(stylemd.StatusError)     // Red
```

#### Alert (CSS pattern)

Feedback message with colored left border.

```go
gui.Div(gui.Class("alert alert--error"))(gui.Text("Something went wrong"))
gui.Div(gui.Class("alert alert--success"))(gui.Text("Operation completed"))
gui.Div(gui.Class("alert alert--warning"))(gui.Text("Please check your input"))
gui.Div(gui.Class("alert alert--info"))(gui.Text("Here's some information"))
```

Variants: `alert--error`, `alert--success`, `alert--warning`, `alert--info`

#### Spinner (stylemd component)

Spinner with optional size.

```go
stylemd.Spinner(stylemd.SpinnerDefault)
stylemd.Spinner(stylemd.SpinnerSmall)
stylemd.Spinner(stylemd.SpinnerLarge)
```

| Size           | Description |
| -------------- | ----------- |
| SpinnerDefault | Medium      |
| SpinnerSmall   | Small       |
| SpinnerLarge   | Large       |

#### StatusDot

Small status indicator dot.

```go
stylemd.StatusDot(stylemd.StatusRunning)   // Green, pulsing
stylemd.StatusDot(stylemd.StatusStopped)   // Gray, static
stylemd.StatusDot(stylemd.StatusStarting)  // Yellow, pulsing
stylemd.StatusDot(stylemd.StatusError)     // Red, static
```

| Variant        | Color   | Animation |
| -------------- | ------- | --------- |
| StatusRunning  | Success | Pulse     |
| StatusStopped  | Muted   | None      |
| StatusStarting | Warning | Pulse     |
| StatusError    | Danger  | None      |

#### ErrorMessage / SuccessMessage

Inline feedback messages for forms.

```go
stylemd.ErrorMessage("Please enter a valid email address.")
stylemd.SuccessMessage("Settings saved successfully.")
```

#### EmptyState

Placeholder for empty content areas.

```go
stylemd.EmptyState("NO PROJECTS", "Get started by creating a new project.")
```

### Modal Components

#### Modal, ModalBody, ModalFooter (stylemd components)

Full-screen modal overlay with structured sections.

```go
stylemd.ModalBackdrop(
    stylemd.Modal("CONFIRM ACTION", handleClose,
        stylemd.ModalBody(
            gui.P()(gui.Text("Are you sure you want to proceed?")),
        ),
        stylemd.ModalFooter(
            stylemd.Button(stylemd.ButtonProps{OnClick: handleClose},
                gui.Text("CANCEL"),
            ),
            stylemd.Button(stylemd.ButtonProps{Variant: stylemd.ButtonPrimary, OnClick: handleConfirm},
                gui.Text("CONFIRM"),
            ),
        ),
    ),
)
```

### Navigation Components

#### Navbar (stylemd component)

Top navigation bar with brand and actions.

```go
stylemd.Navbar(stylemd.NavbarProps{
    Brand: "MY APP",
    Nav:   navLinks,
    Actions: actionButtons,
})
```

#### NavLink (stylemd component)

Sidebar navigation link with optional icon.

```go
stylemd.NavLink(stylemd.NavLinkProps{
    Icon:    "icon-chat",
    Label:   "Chat",
    Active:  true,
    OnClick: handleClick,
})
```

#### TabBar (stylemd component)

Horizontal tab navigation.

```go
stylemd.TabBar([]stylemd.TabBarTab{
    {Label: "General", Active: true, OnClick: handleGeneral},
    {Label: "Advanced", OnClick: handleAdvanced},
})
```

#### BottomTabBar (stylemd component)

Mobile bottom tab bar.

```go
stylemd.BottomTabBar([]stylemd.BottomTabItem{
    {Icon: "icon-home", Label: "Home", Active: true, OnClick: handleHome},
    {Icon: "icon-settings", Label: "Settings", OnClick: handleSettings},
})
```

#### Sidebar (stylemd component)

Side navigation panel.

```go
stylemd.Sidebar(stylemd.SidebarProps{
    Header: sidebarHeader,
    Footer: sidebarFooter,
}, sidebarContent...)
```

### Layout Components

#### AppShell (stylemd component)

Full-page layout with navbar, sidebar, and main content area.

```go
stylemd.AppShell(navbar,
    stylemd.AppShellBody(
        sidebar,
        stylemd.AppShellMain(content...),
    ),
)
```

#### Container (CSS pattern)

Centered content container with max-width.

```go
gui.Div(gui.Class("container container--md"))(
    gui.Text("Centered content"),
)
```

Sizes: `container--sm` (384px), `container--md` (448px), `container--lg` (512px), `container--xl` (640px)

#### Divider (CSS pattern)

Horizontal rule with configurable spacing.

```go
gui.Hr(gui.Class("divider"))()
gui.Hr(gui.Class("divider divider--sm"))()
gui.Hr(gui.Class("divider divider--lg"))()
```

### Data Display Components

#### DataTable (stylemd component)

Table for structured data with headers and rows.

```go
stylemd.DataTable(
    []string{"Name", "Status", "CPU", "Memory"},
    [][]gui.Node{
        {gui.Text("node-01"), stylemd.StatusBadge(stylemd.StatusRunning), gui.Text("23%"), gui.Text("2.1 GB")},
        {gui.Text("node-02"), stylemd.StatusBadge(stylemd.StatusRunning), gui.Text("67%"), gui.Text("3.8 GB")},
    },
)
```

#### DiffViewer (stylemd component)

Code diff display with add/remove highlighting.

```go
stylemd.DiffViewer(diffContent)
```

#### CodeBlock / InlineCode (CSS patterns)

Display code and technical values.

```go
// Block code
gui.Pre(gui.Class("code-block"))(
    gui.Code()(gui.Text("eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...")),
)

// Inline code
gui.Code(gui.Class("inline-code"))(gui.Text("--verbose"))
```

### List Components

#### ConversationItem, InstanceCard, ServiceRow, RunnerRow

Pre-built list item components for common patterns. See the Go source files for detailed props.

```go
stylemd.ConversationItem(stylemd.ConversationItemProps{
    Title:    "Refactor auth module",
    Subtitle: "claude-3.5-sonnet",
    Active:   true,
    OnClick:  handleClick,
})

stylemd.ServiceRow(stylemd.ServiceRowProps{
    Name:   "postgres",
    Image:  "postgres:16-alpine",
    Status: stylemd.StatusRunning,
    Ports:  "5432:5432",
})
```

### Overlay Components

#### SearchOverlay (stylemd component)

Full-screen search overlay with tabs and results.

```go
stylemd.SearchOverlay(tabs, searchInput, results...)
```

#### ContextMenu (stylemd component)

Positioned context menu.

```go
stylemd.ContextMenu(x, y, []stylemd.ContextMenuItem{
    {Label: "Edit", OnClick: handleEdit},
    {Label: "Delete", Danger: true, OnClick: handleDelete},
})
```

### Page Components

#### LoginPage (stylemd component)

Pre-built login page layout.

```go
stylemd.LoginPage(stylemd.LoginPageProps{
    Title:      "Welcome Back",
    OnSubmit:   handleLogin,
    OnRegister: handleRegister,
})
```

#### SetupWizard (stylemd component)

Multi-step setup wizard with progress indicator.

```go
stylemd.SetupWizard(steps, currentStep, content)
```

---

## Layout Patterns

### Centered Card Layout (Login, Forms)

```go
func LoginPage() gui.Node {
    return gui.Div(gui.Class("flex items-center justify-center"),
        gui.Style("min-height: 100vh; background: var(--bg)"),
    )(
        gui.Div(gui.Style("max-width: 400px; width: 100%"))(
            gui.Div(gui.Class("card"))(
                gui.Div(gui.Class("card__header card__header--dark"))(
                    gui.H2(gui.Style("margin: 0"))(gui.Text("SIGN IN")),
                ),
                gui.Div(gui.Class("card__body"))(
                    stylemd.FormGroup("USERNAME",
                        stylemd.TextInput(stylemd.TextInputProps{ID: "username"}),
                    ),
                    stylemd.FormGroup("PASSWORD",
                        stylemd.TextInput(stylemd.TextInputProps{ID: "password", Type: "password"}),
                    ),
                    stylemd.Button(stylemd.ButtonProps{Variant: stylemd.ButtonPrimary},
                        gui.Text("SIGN IN"),
                    ),
                ),
            ),
        ),
    )
}
```

### Dashboard Header with Stats

```go
gui.Div(gui.Class("display-panel"),
    gui.Style("display: flex; justify-content: space-between; align-items: center; margin-bottom: var(--space-6)"),
)(
    gui.Div()(
        gui.Span(gui.Class("display-panel__label"))(gui.Text("SECTION NAME")),
        gui.Span(gui.Class("display-panel__value display-panel__value--highlight"),
            gui.Style("font-size: var(--text-xl)"),
        )(gui.Text("MAIN TITLE")),
    ),
    gui.Div(gui.Class("flex items-center gap-3"))(
        gui.Span(gui.Class("display-panel__label"),
            gui.Style("margin-bottom: 0"),
        )(gui.Text("TOTAL")),
        gui.Div(gui.Style("font-size: var(--text-2xl); font-weight: 700; color: var(--white); font-variant-numeric: tabular-nums"))(
            gui.Text(fmt.Sprintf("%02d", count)),
        ),
    ),
)
```

### Confirmation Modal

```go
func ConfirmDeleteModal(itemID string, onClose, onDelete func()) gui.Node {
    return stylemd.ModalBackdrop(
        stylemd.Modal("CONFIRM DELETE", onClose,
            stylemd.ModalBody(
                gui.P(gui.Style("margin-top: 0"))(gui.Text("Are you sure you want to delete this item?")),
                gui.Div(gui.Class("display-panel"),
                    gui.Style("padding: var(--space-3) var(--space-4)"),
                )(
                    gui.Code(gui.Style("font-family: var(--font-mono); color: var(--accent); font-size: var(--text-sm)"))(
                        gui.Text(itemID),
                    ),
                ),
                gui.Div(gui.Class("alert alert--error"),
                    gui.Style("margin-top: var(--space-4)"),
                )(gui.Text("This action cannot be undone.")),
            ),
            stylemd.ModalFooter(
                stylemd.Button(stylemd.ButtonProps{OnClick: onClose},
                    gui.Text("CANCEL"),
                ),
                stylemd.Button(stylemd.ButtonProps{Variant: stylemd.ButtonDanger, OnClick: onDelete},
                    gui.Text("DELETE"),
                ),
            ),
        ),
    )
}
```

### List Item with Status

```go
func ListItem(title, subtitle string, isActive bool) gui.Node {
    variant := stylemd.StatusStopped
    badgeClass := "badge badge--outline"
    badgeText := "INACTIVE"
    if isActive {
        variant = stylemd.StatusRunning
        badgeClass = "badge badge--success"
        badgeText = "ACTIVE"
    }

    return gui.Div(gui.Class("card"),
        gui.Style("padding: var(--space-4); display: flex; align-items: center; gap: var(--space-3)"),
    )(
        stylemd.StatusDot(variant),
        gui.Div(gui.Style("flex: 1"))(
            gui.Div(gui.Style("font-weight: 700; font-size: var(--text-base)"))(
                gui.Text(title),
            ),
            gui.Div(gui.Style("font-size: var(--text-sm); color: var(--dark)"))(
                gui.Text(subtitle),
            ),
        ),
        gui.Span(gui.Class(badgeClass))(gui.Text(badgeText)),
        stylemd.Button(stylemd.ButtonProps{Size: stylemd.ButtonSmall, Variant: stylemd.ButtonGhost},
            gui.Text("EDIT"),
        ),
    )
}
```

### App Shell with Header/Footer

```go
func AppLayout(content gui.Node) gui.Node {
    return stylemd.AppShell(
        stylemd.Navbar(stylemd.NavbarProps{
            Brand: "MY APP",
            Nav: []gui.Node{
                stylemd.NavLink(stylemd.NavLinkProps{Label: "Dashboard", Active: true}),
                stylemd.NavLink(stylemd.NavLinkProps{Label: "Settings"}),
            },
        }),
        stylemd.AppShellBody(
            stylemd.Sidebar(stylemd.SidebarProps{}, sidebarNav...),
            stylemd.AppShellMain(content),
        ),
    )
}
```

---

## Do's and Don'ts

### DO

- **Use design tokens** - Always use CSS variables (`var(--*)`) instead of hardcoded values
- **Keep text uppercase** for labels, badges, and headings
- **Use wide letter-spacing** with uppercase text (`--tracking-wider`)
- **Apply hard shadows** - No blur, just solid offset shadows
- **Use orange sparingly** - It's for emphasis, not decoration
- **Maintain high contrast** - Black/white with gray backgrounds
- **Use display panels for important data** - Dark background, white/orange text
- **Use tabular-nums** for numbers that update or align
- **Import styles first** - Include `styles.css` in your entry point
- **Compose components using `stylemd` package functions** for modals, forms, navigation
- **Use semantic variants** - `danger` for destructive, `primary` for main CTA

### DON'T

- **Don't use soft shadows** - No blur radius, no spread
- **Don't use rounded corners larger than 8px** - Keep it industrial
- **Don't mix fonts** - Space Mono only
- **Don't use gradients** - Flat colors only
- **Don't overuse orange** - Reserve for CTAs, highlights, and active states
- **Don't skip the border** - Cards and inputs need their 2px black border
- **Don't use sentence case for labels** - Always uppercase
- **Don't use colors outside the palette** - Stick to the defined tokens
- **Don't add decorative elements** - Every element should be functional
- **Don't use lowercase for buttons** - Buttons are always uppercase

### Component Selection Guide

| Need               | Use                                                                   |
| ------------------ | --------------------------------------------------------------------- |
| Primary CTA        | `stylemd.Button(stylemd.ButtonProps{Variant: stylemd.ButtonPrimary})` |
| Secondary action   | `stylemd.Button(stylemd.ButtonProps{})`                               |
| Destructive action | `stylemd.Button(stylemd.ButtonProps{Variant: stylemd.ButtonDanger})`  |
| Text-only button   | `stylemd.Button(stylemd.ButtonProps{Variant: stylemd.ButtonGhost})`   |
| Display key data   | `.display-panel` with `__label` and `__value`                        |
| Form container     | `.card` with `.card__header` and `.card__body`                       |
| Modal dialog       | `stylemd.Modal()` inside `stylemd.ModalBackdrop()`                   |
| Error message      | `.alert.alert--error` or `stylemd.ErrorMessage()`                    |
| Success message    | `.alert.alert--success` or `stylemd.SuccessMessage()`                |
| Status indicator   | `stylemd.StatusDot()` with appropriate variant                       |
| Tag/label          | `.badge` with variant class                                          |
| Long code/keys     | `.code-block` with `<pre><code>`                                     |
| Inline code        | `.inline-code` with `<code>`                                         |
| Form field wrapper | `stylemd.FormGroup("LABEL", input)`                                  |
| Loading state      | `stylemd.Spinner()` with optional size                               |
| Page wrapper       | `.container` with size variant                                       |

---

## Code Examples

### Complete Login Form

```go
package main

import (
    gui "github.com/readmedotmd/gui.md"
    stylemd "github.com/readmedotmd/core.md"
)

func LoginPage(onSubmit func(), errMsg string) gui.Node {
    return gui.Div(gui.Class("flex items-center justify-center"),
        gui.Style("min-height: 100vh; background: var(--bg)"),
    )(
        gui.Div(gui.Style("max-width: 400px; width: 100%"))(
            gui.Div(gui.Class("card"))(
                gui.Div(gui.Class("card__header card__header--dark"))(
                    gui.H2(gui.Style("margin: 0; font-size: var(--text-lg)"))(
                        gui.Text("SIGN IN"),
                    ),
                ),
                gui.Div(gui.Class("card__body"))(
                    // Error alert (conditional)
                    gui.If(errMsg != "",
                        gui.Div(gui.Class("alert alert--error"),
                            gui.Style("margin-bottom: var(--space-4)"),
                        )(gui.Text(errMsg)),
                    ),

                    stylemd.FormGroup("USERNAME",
                        stylemd.TextInput(stylemd.TextInputProps{
                            ID:          "username",
                            Placeholder: "Enter username",
                        }),
                    ),
                    stylemd.FormGroup("PASSWORD",
                        stylemd.TextInput(stylemd.TextInputProps{
                            ID:          "password",
                            Type:        "password",
                            Placeholder: "Enter password",
                        }),
                    ),
                    stylemd.Button(stylemd.ButtonProps{
                        Variant: stylemd.ButtonPrimary,
                        OnClick: onSubmit,
                    },
                        gui.Text("SIGN IN"),
                    ),

                    gui.Hr(gui.Class("divider"))(),

                    gui.P(gui.Style("text-align: center; font-size: var(--text-sm); color: var(--dark); margin: 0"))(
                        gui.Text("Don't have an account? "),
                        gui.A(gui.Attr("href", "/register"),
                            gui.Style("color: var(--accent)"),
                        )(gui.Text("REGISTER")),
                    ),
                ),
            ),
        ),
    )
}
```

### Data Table with Actions

```go
package main

import (
    gui "github.com/readmedotmd/gui.md"
    stylemd "github.com/readmedotmd/core.md"
)

type Item struct {
    ID          string
    Name        string
    Description string
    Active      bool
}

func DataList(items []Item, loading bool, errMsg string, onDelete func(string)) gui.Node {
    if loading {
        return gui.Div(gui.Class("card"),
            gui.Style("display: flex; justify-content: center; padding: var(--space-12)"),
        )(
            stylemd.Spinner(stylemd.SpinnerLarge),
        )
    }

    if errMsg != "" {
        return gui.Div(gui.Class("alert alert--error"))(gui.Text(errMsg))
    }

    if len(items) == 0 {
        return stylemd.EmptyState("NO ITEMS FOUND", "Nothing to display yet.")
    }

    children := make([]gui.Node, len(items))
    for i, item := range items {
        children[i] = dataListItem(item, onDelete)
    }

    return gui.Div(gui.Class("flex-col gap-3"))(children...)
}

func dataListItem(item Item, onDelete func(string)) gui.Node {
    variant := stylemd.StatusStopped
    if item.Active {
        variant = stylemd.StatusRunning
    }

    return gui.Div(gui.Class("card"),
        gui.Style("display: flex; align-items: center; padding: var(--space-4); gap: var(--space-4)"),
    )(
        stylemd.StatusDot(variant),
        gui.Div(gui.Style("flex: 1; min-width: 0"))(
            gui.Div(gui.Style("font-weight: 700; margin-bottom: var(--space-1)"))(
                gui.Text(item.Name),
            ),
            gui.Div(gui.Style("font-size: var(--text-sm); color: var(--dark); overflow: hidden; text-overflow: ellipsis"))(
                gui.Text(item.Description),
            ),
        ),
        gui.Span(gui.Class(badgeClass(item.Active)))(
            gui.Text(badgeText(item.Active)),
        ),
        stylemd.Button(stylemd.ButtonProps{
            Size:    stylemd.ButtonSmall,
            Variant: stylemd.ButtonDanger,
            OnClick: func() { onDelete(item.ID) },
        },
            gui.Text("DELETE"),
        ),
    )
}
```

### Settings Panel

```go
package main

import (
    gui "github.com/readmedotmd/gui.md"
    stylemd "github.com/readmedotmd/core.md"
)

func SettingsPanel(username string, verified bool) gui.Node {
    return gui.Div(gui.Class("card"))(
        gui.Div(gui.Class("card__header card__header--dark"))(
            gui.H2(gui.Style("margin: 0; font-size: var(--text-md)"))(
                gui.Text("ACCOUNT SETTINGS"),
            ),
        ),
        gui.Div(gui.Class("card__body"))(
            // User info display panel
            gui.Div(gui.Class("display-panel"),
                gui.Style("margin-bottom: var(--space-6); display: flex; justify-content: space-between; align-items: flex-start"),
            )(
                gui.Div()(
                    gui.Span(gui.Class("display-panel__label"))(gui.Text("USERNAME")),
                    gui.Span(gui.Class("display-panel__value display-panel__value--highlight"))(gui.Text(username)),
                ),
                gui.Span(gui.Class("badge badge--success"))(gui.Text("VERIFIED")),
            ),

            // Settings list
            settingsRow("EMAIL NOTIFICATIONS", "Receive updates via email",
                stylemd.Button(stylemd.ButtonProps{Size: stylemd.ButtonSmall},
                    gui.Text("CONFIGURE"),
                ),
            ),
            settingsRow("TWO-FACTOR AUTH", "Extra security for your account",
                stylemd.Button(stylemd.ButtonProps{Size: stylemd.ButtonSmall, Variant: stylemd.ButtonPrimary},
                    gui.Text("ENABLE"),
                ),
            ),
            settingsRow("DELETE ACCOUNT", "Permanently remove your account",
                stylemd.Button(stylemd.ButtonProps{Size: stylemd.ButtonSmall, Variant: stylemd.ButtonDanger},
                    gui.Text("DELETE"),
                ),
            ),
        ),
    )
}

func settingsRow(title, description string, action gui.Node) gui.Node {
    return gui.Div(gui.Style("display: flex; justify-content: space-between; align-items: center; padding: var(--space-4) 0; border-bottom: 1px solid var(--bg)"))(
        gui.Div()(
            gui.Div(gui.Style("font-weight: 700"))(gui.Text(title)),
            gui.Div(gui.Style("font-size: var(--text-sm); color: var(--dark)"))(
                gui.Text(description),
            ),
        ),
        action,
    )
}
```

---

## Utility Classes

The stylesheet includes utility classes for common patterns:

### Text Utilities

- `.text-xs` through `.text-2xl` - Font sizes
- `.text-uppercase` - Uppercase with wider letter-spacing
- `.text-muted` - Muted color
- `.text-accent`, `.text-danger`, `.text-success`, `.text-warning`, `.text-blue` - Text colors

### Layout Utilities

- `.flex`, `.inline-flex`, `.grid`, `.block` - Display
- `.flex-col` - Column direction
- `.items-center`, `.items-start`, `.items-end` - Align items
- `.justify-center`, `.justify-between`, `.justify-end` - Justify content
- `.gap-1` through `.gap-6` - Gap spacing

### Spacing Utilities

- `.p-0` through `.p-6` - Padding
- `.m-0`, `.m-auto` - Margin
- `.mb-2`, `.mb-4`, `.mb-6` - Margin bottom

### Animation Utilities

- `.animate-spin` - Rotating spinner
- `.animate-blink` - Blinking effect
- `.animate-pulse` - Pulsing opacity
- `.animate-fade-in` - Fade in
- `.animate-slide-up` - Slide up with fade

### Special Utilities

- `.tabular-nums` - Tabular number alignment
- `.break-all` - Word break for long strings
- `.truncate` - Text overflow ellipsis
- `.sr-only` - Screen-reader only (visually hidden)

---

## Z-Index Layers

Use these for proper stacking:

| Token          | Value | Usage               |
| -------------- | ----- | ------------------- |
| `--z-base`     | 0     | Default content     |
| `--z-dropdown` | 100   | Dropdowns, popovers |
| `--z-sticky`   | 200   | Sticky headers      |
| `--z-overlay`  | 300   | Overlay backgrounds |
| `--z-modal`    | 400   | Modal dialogs       |
| `--z-tooltip`  | 500   | Tooltips            |

---

## Transitions

| Token                | Value     | Usage                        |
| -------------------- | --------- | ---------------------------- |
| `--transition-fast`  | 0.1s ease | Hover states, quick feedback |
| `--transition-normal`| 0.2s ease | Most interactions            |
| `--transition-slow`  | 0.3s ease | Page transitions, modals     |
