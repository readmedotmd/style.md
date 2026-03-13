<p align="center">
  <img src="../design/core-banner.png" width="900" alt="core.md" />
</p>

<p align="center">
  Headless UI components for Go with minimal base styles, built on gui.md.
</p>

<p align="center">

```
go get github.com/readmedotmd/style.md/core.md
```

</p>

---

## What is core.md?

**core.md** provides 110+ UI components that render semantic HTML with `data-*` attributes for state. Components include just enough CSS for usability — no visual opinions.

Build your UI with core.md, then apply any theme by loading its CSS and setting `data-theme` on `<html>`.

```go
import (
    gui "github.com/readmedotmd/gui.md"
    coremd "github.com/readmedotmd/style.md/core.md"
)

func App() gui.Node {
    return coremd.Stack("lg",
        coremd.Heading(1, "", gui.Text("Dashboard")),
        coremd.HStack("md",
            coremd.Badge("", coremd.BadgeSuccess, "Online"),
            coremd.Muted("3 services running"),
        ),
        coremd.Card(coremd.CardProps{},
            coremd.DataTable("", []string{"Name", "Status"}, [][]gui.Node{
                {gui.Text("api"), coremd.StatusBadge("", coremd.StatusRunning, "Running")},
                {gui.Text("worker"), coremd.StatusBadge("", coremd.StatusPending, "Pending")},
            }),
        ),
        coremd.Button(coremd.ButtonProps{Variant: "primary"}, gui.Text("Deploy")),
    )
}
```

```html
<link rel="stylesheet" href="core.md/styles.css">
```

## Primitives

Layout and content primitives so your UI never needs external CSS.

### Layout

```go
// Vertical stack with medium gap
coremd.Stack("md", child1, child2, child3)

// Horizontal stack with spacer
coremd.HStack("md", left, coremd.Spacer(), right)

// 3-column grid
coremd.Grid(coremd.GridProps{Cols: "3"}, col1, col2, col3)

// Centered content
coremd.Center("", content)
```

```html
<!-- Or use data attributes directly in HTML -->
<div data-stack="md">...</div>
<div data-hstack="lg">...</div>
<div data-grid="3">...</div>
<div data-center>...</div>
```

Gap values: `xs` (4px), `sm` (8px), `md` (16px), `lg` (24px), `xl` (32px), `none` (0).

### Cards & Badges

```go
coremd.Card(coremd.CardProps{}, content)
coremd.Card(coremd.CardProps{Variant: "surface"}, content)

coremd.Badge("", coremd.BadgeAccent, "New")
coremd.Badge("", coremd.BadgeDanger, "Critical")
```

### Typography

```go
coremd.Heading(1, "", gui.Text("Title"))
coremd.Paragraph("", gui.Text("Body text."))
coremd.CodeBlock("", "fmt.Println(\"hello\")")
coremd.InlineCode("go build")
coremd.Muted("Secondary text")
coremd.Quote("", gui.Text("Important note."))
```

### Links, Images, Lists

```go
coremd.Link(coremd.LinkProps{Href: "/docs"}, gui.Text("Documentation"))
coremd.Image(coremd.ImageProps{Src: "photo.jpg", Alt: "Photo", Rounded: true})
coremd.Image(coremd.ImageProps{Src: "avatar.jpg", Avatar: true})

coremd.UnorderedList("",
    coremd.ListItem(gui.Text("First")),
    coremd.ListItem(gui.Text("Second")),
)
```

### Utilities

```go
coremd.Divider("")           // Horizontal rule
coremd.Truncate("", text)    // Ellipsis overflow
coremd.Mono("monospace")     // Monospace font
coremd.SrOnly("screen only") // Screen reader only
```

## Components

| Category        | Components | File |
|-----------------|------------|------|
| **Primitives**  | Stack, HStack, Grid, Center, Spacer, Card, Badge, Divider, Heading, Paragraph, CodeBlock, InlineCode, Link, Image, UnorderedList, OrderedList, Quote, Muted, Mono, Truncate, SrOnly, MarkdownContent, SectionHeader, Collapsible, Animate, HelpText | `primitives.go` |
| **Buttons**     | Button (primary, danger, toolbar; medium, small) | `button.go` |
| **Forms**       | FormGroup, TextInput, NumberInput, TextArea, SelectInput, Checkbox, FeatureRow, VariableRow, EditableVariableRow, PasswordField, SecretField, SchemaField, ErrorMessage, SuccessMessage | `form.go` |
| **Input**       | ChatInput, AutocompletePopup, MessageQueue, SearchInputField, PastePreview, MessageQueueBar, QueuedItem, AttachmentButton, ModeToggle | `input.go` |
| **Display**     | MessageBubble, QuestionPrompt, StatusBadge, StatusDot, LabelBadge, UsageBadge, DiffViewer, DataTable, EmptyState, ClusterStatsBar, MessageContent, ActionTag, SystemStats, DiffPanel, StatChip, VariableChip | `display.go` |
| **Lists**       | ConversationItem, InstanceCard, InstanceList, ServiceRow, RunnerRow, FileTree, DevboxCard, EnvironmentCard | `list.go` |
| **Navigation**  | NavLink, TabBar, BottomTabBar | `navigation.go` |
| **Overlay**     | SearchOverlay, ContextMenu, BottomSheet, SearchResult, SearchResultContent, SearchSnippet | `overlay.go` |
| **Panels**      | GitPanel, SkillsPanel, TerminalPanel, GitSectionHeader, GitFileList, GitFile, GitCommitArea | `panel.go` |
| **Layout**      | AppShell, Navbar, Sidebar, Panel, Modal, ModalBackdrop, DashboardLayout, SidebarColumn, ChatHeader, Box, ScrollArea, SplitLayout, Backdrop, IconButton, ToolbarButton, Toolbar, ResizeHandle | `layout.go` |
| **Pages**       | LoginPage, SetupWizard, SettingsCard, SettingsPage, SettingsCardFull, SettingsSection, SettingsSubsection, SettingsForm, SettingsCodeInput, ClusterSummaryCard, ClusterSummaryRow | `page.go` |
| **Utility**     | Spinner, Icon, AppShellFull | `utility.go` |

## CSS Base Classes

Every component renders a **semantic base class** by default, even when no custom class is provided. This guarantees CSS hooks are always present for styling. When you pass a custom `class`, it's appended after the base class.

```go
// Always renders class="btn"
coremd.Button(coremd.ButtonProps{}, gui.Text("OK"))

// Renders class="btn my-btn"
coremd.Button(coremd.ButtonProps{Class: "my-btn"}, gui.Text("OK"))
```

| Component | Base Class | File |
|-----------|-----------|------|
| AppShell | `app` | layout.go |
| AppShellBody | `app-shell-body` | layout.go |
| AppShellMain | `app-shell-main` | layout.go |
| Navbar | `navbar` | layout.go |
| Sidebar | `sidebar` | layout.go |
| SidebarHeader | `sidebar-header` | layout.go |
| Panel | `panel` | layout.go |
| ModalBackdrop | `modal-backdrop` | layout.go |
| Modal | `modal` | layout.go |
| ModalBody | `modal-body` | layout.go |
| ModalFooter | `modal-footer` | layout.go |
| DashboardLayout | `dashboard-layout` | layout.go |
| SidebarColumn | `sidebar-col` | layout.go |
| ChatHeader | `chat-header` | layout.go |
| Box | `box` | layout.go |
| ScrollArea | `scroll-area` | layout.go |
| SplitLayout | `split-layout` | layout.go |
| Backdrop | `backdrop` | layout.go |
| IconButton | `icon-button` | layout.go |
| Toolbar | `toolbar` | layout.go |
| ToolbarButton | `toolbar-button` | layout.go |
| ResizeHandle | `resize-handle` | layout.go |
| Button | `btn` | button.go |
| NavLink | `nav-link` | navigation.go |
| TabBar | `tab-bar` | navigation.go |
| BottomTabBar | `bottom-tab-bar` | navigation.go |
| MessageBubble | `message` | display.go |
| QuestionPrompt | `question-prompt` | display.go |
| StatusBadge | `status-badge` | display.go |
| StatusDot | `status-dot` | display.go |
| LabelBadge | `label-badge` | display.go |
| UsageBadge | `usage-badge` | display.go |
| DiffViewer | `diff-viewer` | display.go |
| DataTable | `data-table` | display.go |
| EmptyState | `empty-state` | display.go |
| ClusterStatsBar | `cluster-stats-bar` | display.go |
| ActionTag | `action-tag` | display.go |
| SystemStats | `system-stats` | display.go |
| DiffPanel | `diff-panel` | display.go |
| StatChip | `stat-chip` | display.go |
| VariableChip | `variable-chip` | display.go |
| ChatInput | `chat-input` | input.go |
| AutocompletePopup | `autocomplete-popup` | input.go |
| MessageQueue | `message-queue` | input.go |
| SearchInputField | `search-input-field` | input.go |
| PastePreview | `paste-preview` | input.go |
| MessageQueueBar | `message-queue` | input.go |
| QueuedItem | `queued-item` | input.go |
| AttachmentButton | `attachment-button` | input.go |
| ModeToggle | `mode-toggle` | input.go |
| Checkbox | `checkbox` | form.go |
| FeatureRow | `feature-row` | form.go |
| VariableRow | `variable-row` | form.go |
| ErrorMessage | `error-message` | form.go |
| SuccessMessage | `success-message` | form.go |
| PasswordField | `password-field` | form.go |
| SecretField | `secret-field` | form.go |
| TextArea | `text-area` | form.go |
| SearchOverlay | `search-overlay` | overlay.go |
| ContextMenu | `context-menu` | overlay.go |
| BottomSheet | `bottom-sheet` | overlay.go |
| ConversationItem | `conv-item` | list.go |
| DevboxCard | `devbox-card` | list.go |
| EnvironmentCard | `env-card` | list.go |
| GitPanel | `git-panel` | panel.go |
| TerminalPanel | `terminal-panel` | panel.go |
| GitSectionHeader | `git-section-header` | panel.go |
| GitFileList | `git-file-list` | panel.go |
| GitFile | `git-file` | panel.go |
| GitCommitArea | `git-commit-area` | panel.go |
| SettingsCard | `settings-card` | page.go |
| SettingsPage | `settings-page` | page.go |
| SettingsSection | `settings-section-group` | page.go |
| SettingsSubsection | `settings-subsection` | page.go |
| SettingsForm | `settings-form` | page.go |
| SettingsCodeInput | `settings-code-input` | page.go |
| ClusterSummaryCard | `cluster-summary-card` | page.go |
| ClusterSummaryRow | `cluster-summary` | page.go |
| Spinner | `spinner` | utility.go |
| Icon | `icon` | utility.go |
| AppShellFull | `app` | utility.go |
| Grid | `grid` | primitives.go |
| Center | `center` | primitives.go |
| Card | `card` | primitives.go |
| Badge | `badge` | primitives.go |
| Divider | `divider` | primitives.go |
| Heading | `heading` | primitives.go |
| Paragraph | `paragraph` | primitives.go |
| CodeBlock | `code-block` | primitives.go |
| Link | `link` | primitives.go |
| Image | `image` | primitives.go |
| UnorderedList | `unordered-list` | primitives.go |
| OrderedList | `ordered-list` | primitives.go |
| Quote | `quote` | primitives.go |
| Truncate | `truncate` | primitives.go |
| MarkdownContent | `markdown-content` | primitives.go |
| SectionHeader | `section-header` | primitives.go |
| Collapsible | `collapsible` | primitives.go |
| Animate | `animate` | primitives.go |
| HelpText | `help-text` | primitives.go |

Components **without** a base class (styled via element selectors or `data-*` attributes only): FormGroup, TextInput, SelectInput.

Internally, base classes are built with the `joinClass` helper:

```go
// joinClass("btn", "my-btn") → "btn my-btn"
// joinClass("btn", "")       → "btn"
```

## Data Attributes

Components use `data-*` attributes for state, which CSS themes can target:

| Attribute | Values | Used by |
|-----------|--------|---------|
| `data-variant` | `primary`, `danger`, `toolbar` | Button |
| `data-size` | `small`, `large` | Button, Spinner |
| `data-active` | `true` | NavLink, TabBar |
| `data-status` | `running`, `stopped`, `starting`, `pending`, `error` | StatusBadge, StatusDot |
| `data-stack` | `xs`, `sm`, `md`, `lg`, `xl` | Stack |
| `data-hstack` | `xs`, `sm`, `md`, `lg`, `xl` | HStack |
| `data-grid` | `1`-`6` | Grid |
| `data-card` | `true`, `surface`, `flush` | Card |
| `data-badge` | `true`, `accent`, `success`, `danger`, `warning` | Badge |
| `data-error` | `true` | TextInput |
| `data-streaming` | `true` | MessageBubble, ChatInput |
| `data-open` | `true` | Sidebar, SidebarColumn |
| `data-expanded` | `true` | Panel, GitPanel |
| `data-role` | `user`, `assistant` | MessageBubble, MessageContent |
| `data-has-image` | `true` | QueuedItem |
| `data-match` | `true` | SearchSnippet lines |
| `data-danger` | `true` | ContextMenu items, BottomSheet items |
| `data-staged` | `true` | GitSectionHeader, GitFile |
| `data-state` | `M`, `A`, `D`, `??` | GitFile |
| `data-selected` | `true` | AutocompletePopup items, GitFile |
| `data-diff` | `add`, `remove`, `header`, `context` | DiffViewer lines |
| `data-scrollable` | `true` | AppShellFull |
| `data-completed` | `true` | SetupWizard steps |
| `data-box` | | Box |
| `data-pad` | `xs`, `sm`, `md`, `lg`, `xl` | Box padding |
| `data-bg` | `surface`, `accent`, `muted` | Box background |
| `data-box-border` | `true` | Box border |
| `data-box-flex` | `true` | Box flex display |
| `data-box-rounded` | `true` | Box rounded corners |
| `data-scroll-area` | | ScrollArea |
| `data-split-layout` | | SplitLayout |
| `data-backdrop` | | Backdrop |
| `data-icon-button` | | IconButton |
| `data-toolbar` | | Toolbar |
| `data-rich-text` | | MarkdownContent |
| `data-section-header` | | SectionHeader |
| `data-collapsible` | | Collapsible |
| `data-animate` | `pulse`, `fade-in`, `spin` | Animate |
| `data-form-group` | | FormGroup |
| `data-feature-info` | | FeatureRow info container |
| `data-feature-name` | | FeatureRow name |
| `data-feature-desc` | | FeatureRow description |
| `data-settings-card-header` | | SettingsCard, SettingsCardFull header |
| `data-settings-card-body` | | SettingsCard, SettingsCardFull body |
| `data-editable-var-row` | | EditableVariableRow |
| `data-passthrough` | `true` | EditableVariableRow |
| `data-schema-field` | | SchemaField |
| `data-schema-field-name` | | SchemaField name span |
| `data-schema-field-type` | | SchemaField type span |
| `data-schema-field-desc` | | SchemaField description span |
| `data-header` | | Generic header primitive (SidebarHeader, ChatHeader, GitPanel, GitSectionHeader, SettingsSection, SettingsCardFull, SettingsSubsection) |
| `data-list-item` | | Generic list item primitive (NavLink, ConversationItem, GitFile) |
| `data-side-panel` | | Generic side panel primitive (GitPanel, TerminalPanel) |
| `data-terminal-panel` | | TerminalPanel |
| `data-terminal-tabs` | | TerminalPanel tab bar |
| `data-terminal-tab` | | TerminalPanel individual tab |
| `data-settings-subsection-header` | | SettingsSubsection header |
| `data-settings-subsection-body` | | SettingsSubsection body |
| `data-dir` | `true` | FileTree directory items |
| `data-working` | `true` | InstanceCard |
| `data-password-field` | | PasswordField |
| `data-password-toggle` | | PasswordField toggle button |
| `data-visible` | `true` | PasswordField (visible mode) |
| `data-secret-field` | | SecretField |
| `data-secret-key` | | SecretField key display |
| `data-secret-value` | | SecretField value display |
| `data-secret-scope` | | SecretField scope display |
| `data-secret-copy` | | SecretField copy button |
| `data-secret-remove` | | SecretField remove button |
| `data-env-card` | | EnvironmentCard |
| `data-stat-chip` | | StatChip |
| `data-variable-chip` | | VariableChip |
| `data-help-text` | | HelpText |

## CSS Custom Properties

Override these tokens to customize the base styles:

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
  --core-info:       #3b82f6;
  --core-radius:     6px;
  --core-space:      8px;
  --core-transition: 150ms ease;

  /* Layout widths */
  --core-width-sidebar:  300px;
  --core-width-panel:    300px;
  --core-width-git-panel: 420px;
  --core-width-modal:    480px;
  --core-width-search:   600px;
  --core-width-message:  680px;
  --core-width-content:  720px;
  --core-width-expanded: 960px;
  --core-width-login:    400px;
}
```

Dark mode: `prefers-color-scheme: dark` (auto-detected by themes) or `data-mode="dark"` on `<html>`.

## Theming

core.md is designed to be themed. A theme is a CSS file scoped under `[data-theme="..."]` that overrides `--core-*` properties and adds styles to `data-*` selectors. Switch themes by changing `data-theme` on `<html>`:

```html
<link rel="stylesheet" href="core.md/styles.css">
<link rel="stylesheet" href="themes/industrial.md/theme.css">
<link rel="stylesheet" href="themes/devbox.md/theme.css">
<html data-theme="industrial">
```

All theme CSS files can be loaded at once — only selectors matching the active `data-theme` apply. See [industrial.md](../themes/industrial.md) and [devbox.md](../themes/devbox.md) for complete theme examples.

## Showcase

<p align="center">
  <img src="../design/screenshots/core-showcase-light.png" width="900" alt="core.md showcase" />
</p>

---

<p align="center">
  <strong>core.md</strong> is part of the <a href="https://github.com/readmedotmd">readme.md</a> project.
</p>
