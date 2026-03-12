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

**core.md** provides 60+ UI components that render semantic HTML with `data-*` attributes for state. Components include just enough CSS for usability — no visual opinions.

Build your UI with core.md, then apply any theme on top with a single `<link>` tag.

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
| **Primitives**  | Stack, HStack, Grid, Center, Spacer, Card, Badge, MicroBadge, Divider, Heading, Paragraph, CodeBlock, InlineCode, Link, Image, UnorderedList, OrderedList, Quote, Muted, Mono, Truncate, SrOnly | `primitives.go` |
| **Buttons**     | Button (primary, danger, toolbar, send, cancel, success-outline, warning-outline, info-outline, ghost; medium, small) | `button.go` |
| **Forms**       | FormGroup, TextInput, TextArea, CodeTextArea, SelectInput, Checkbox, FeatureRow, VariableRow, KVRow, ErrorMessage, SuccessMessage | `form.go` |
| **Input**       | ChatInput, AutocompletePopup, MessageQueue, SearchInputField | `input.go` |
| **Display**     | MessageBubble, RichText, ThinkingIndicator, ThinkingCollapsible, ToolBadge, Indicator, QuestionPrompt, StatusBadge, StatusDot, LabelBadge, UsageBadge, DiffViewer, DataTable, EmptyState, ClusterStatsBar | `display.go` |
| **Lists**       | List, ListItem, SectionHeader, ConversationItem, InstanceCard, ServiceRow, RunnerRow, FileTree | `list.go` |
| **Navigation**  | NavLink, TabBar, BottomTabBar | `navigation.go` |
| **Overlay**     | SearchOverlay, ContextMenu, BottomSheet, Modal, ModalBackdrop | `overlay.go` |
| **Panels**      | Panel, PanelHeader, ServicesPanel, RunnerPanel, GitPanel, SkillsPanel, TerminalPanel, FileBrowser | `panel.go` |
| **Layout**      | AppShell, Navbar, Sidebar, DragHandle | `layout.go` |
| **Pages**       | LoginPage, SetupWizard, DashboardPage, SettingsCard | `page.go` |
| **Utility**     | Spinner, Icon, ImagePreview | `utility.go` |

## Data Attributes

Components use `data-*` attributes for state, which CSS themes can target:

| Attribute | Values | Used by |
|-----------|--------|---------|
| `data-variant` | `primary`, `danger`, `toolbar`, `send`, `cancel`, `success-outline`, `warning-outline`, `info-outline`, `ghost` | Button |
| `data-size` | `small`, `large` | Button, Spinner |
| `data-active` | `true` | NavLink, TabBar, NavLink, ListItem, TerminalTab |
| `data-status` | `running`, `stopped`, `starting`, `pending`, `error` | StatusBadge, StatusDot |
| `data-stack` | `xs`, `sm`, `md`, `lg`, `xl`, `none` | Stack |
| `data-hstack` | `xs`, `sm`, `md`, `lg`, `xl`, `none` | HStack |
| `data-grid` | `1`-`6` | Grid |
| `data-card` | `true`, `surface`, `flush` | Card |
| `data-badge` | `true`, `accent`, `success`, `danger`, `warning` | Badge |
| `data-error` | `true` | TextInput |
| `data-streaming` | `true` | MessageBubble, ChatInput |
| `data-open` | `true` | Sidebar |
| `data-app-shell` | (presence) | AppShell |
| `data-navbar` | (presence) | Navbar |
| `data-nav-brand` | (presence) | Navbar brand text |
| `data-nav-links` | (presence) | Navbar link container |
| `data-nav-link` | (presence) | Individual nav link |
| `data-sidebar` | (presence) | Sidebar |
| `data-sidebar-header` | (presence) | Sidebar header |
| `data-bottom-tabbar` | (presence) | BottomTabBar (mobile) |
| `data-tab-link` | (presence) | Tab link in BottomTabBar |
| `data-panel` | (presence), `wide` | Panel |
| `data-panel-header` | (presence) | Panel header row |
| `data-panel-header-title` | (presence) | Panel header title |
| `data-panel-actions` | (presence) | Panel header action buttons |
| `data-modal-backdrop` | (presence) | Modal backdrop overlay |
| `data-modal` | (presence) | Modal dialog container |
| `data-modal-header` | (presence) | Modal header |
| `data-modal-title` | (presence) | Modal title text |
| `data-modal-body` | (presence) | Modal body content |
| `data-modal-actions` | (presence) | Modal footer actions |
| `data-drag-handle` | (presence) | Mobile drag indicator |
| `data-bottom-sheet-backdrop` | (presence) | Bottom sheet backdrop |
| `data-bottom-sheet` | (presence) | Bottom sheet container |
| `data-sheet-item` | (presence) | Bottom sheet action item |
| `data-context-menu-backdrop` | (presence) | Context menu backdrop |
| `data-context-menu` | (presence) | Context menu container |
| `data-context-menu-item` | (presence) | Context menu item |
| `data-message` | `outgoing`, `incoming`, `streaming`, `error` | MessageBubble |
| `data-rich-text` | (presence) | Rich text container |
| `data-collapsible` | (presence) | Collapsible container |
| `data-collapsible-summary` | (presence) | Collapsible toggle |
| `data-collapsible-content` | (presence) | Collapsible body |
| `data-indicator` | (presence), `working` | Indicator pill |
| `data-autocomplete` | (presence) | Autocomplete popup |
| `data-autocomplete-item` | (presence) | Autocomplete option |
| `data-queue` | (presence) | Message queue |
| `data-queue-item` | (presence) | Message queue item |
| `data-search-backdrop` | (presence) | Search overlay backdrop |
| `data-search-card` | (presence) | Search overlay card |
| `data-search-input` | (presence) | Search text input |
| `data-search-results` | (presence) | Search results container |
| `data-search-result` | (presence) | Individual search result |
| `data-image-preview` | (presence) | Image preview container |
| `data-preview-thumbnail` | (presence) | Preview thumbnail image |
| `data-file-tree` | (presence) | File tree container |
| `data-file-item` | `dir`, `file` | File tree entry |
| `data-terminal-panel` | (presence) | Terminal panel container |
| `data-terminal-tabs` | (presence) | Terminal tab bar |
| `data-terminal-tab` | (presence) | Terminal tab button |
| `data-terminal-iframe` | (presence) | Terminal iframe |
| `data-list` | (presence) | Selectable list |
| `data-list-item` | (presence) | List item |
| `data-section-header` | (presence), `success`, `warning` | Section header label |
| `data-question` | (presence) | Question prompt |
| `data-question-text` | (presence) | Question text |
| `data-question-option` | (presence) | Question option button |
| `data-empty-state` | (presence) | Empty state placeholder |
| `data-feature-row` | (presence) | Feature/setting row |
| `data-kv-row` | (presence) | Key-value input row |
| `data-kv-remove` | (presence) | KV row remove button |
| `data-micro-badge` | (presence) | Micro badge |
| `data-code` | (presence), `tall`, `medium`, `short` | Code textarea |
| `data-animate` | `pulse`, `appear` | Animation utility |
| `data-hide` | `mobile`, `desktop` | Responsive visibility |
| `data-hidden` | (presence) | Hide element |
| `data-visible` | `false` | Hide element |
| `data-danger` | `true` | Danger state modifier |

## CSS Custom Properties

Override these tokens to customize the base styles:

```css
:root {
  /* Typography */
  --core-font:       system-ui, sans-serif;
  --core-font-mono:  ui-monospace, monospace;

  /* Colors */
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

  /* On-accent (content rendered on accent backgrounds) */
  --core-on-accent:        #ffffff;
  --core-on-accent-muted:  rgba(255, 255, 255, 0.8);
  --core-on-accent-subtle: rgba(255, 255, 255, 0.13);
  --core-on-accent-border: rgba(255, 255, 255, 0.2);

  /* Surfaces */
  --core-raised:     var(--core-surface);
  --core-hover:      rgba(0, 0, 0, 0.04);

  /* Spacing & Radius */
  --core-radius:     6px;
  --core-radius-lg:  calc(var(--core-radius) * 2);
  --core-space:      8px;
  --core-transition: 150ms ease;

  /* Overlays */
  --core-backdrop:      rgba(0, 0, 0, 0.5);
  --core-backdrop-blur: 2px;

  /* Shadows */
  --core-shadow-sm: 0 1px 2px rgba(0, 0, 0, 0.05);
  --core-shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.07), 0 2px 4px -2px rgba(0, 0, 0, 0.05);
  --core-shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.08), 0 4px 6px -4px rgba(0, 0, 0, 0.04);

  /* Font size scale */
  --core-text-2xs:  0.65rem;
  --core-text-xs:   0.72rem;
  --core-text-sm:   0.82rem;
  --core-text-base: 0.875rem;
  --core-text-md:   0.95rem;
  --core-text-lg:   1.05rem;
  --core-text-xl:   1.25rem;
}
```

Dark mode: `prefers-color-scheme: dark` or `data-theme="dark"` on `<html>`.

## Button Variants

Buttons support the following `data-variant` values:

| Variant | Appearance |
|---------|------------|
| (none) | Default — bordered, neutral background |
| `primary` | Solid accent background, white text |
| `danger` | Solid danger background, white text |
| `toolbar` | Transparent, no border; border appears on hover |
| `send` | Solid accent background, uses `--core-on-accent` text |
| `cancel` | Transparent with danger border/text; fills on hover |
| `success-outline` | Transparent with success border/text |
| `warning-outline` | Transparent with warning border/text |
| `info-outline` | Transparent with info border/text |
| `ghost` | Fully transparent, muted text; surface background on hover |

```go
coremd.Button(coremd.ButtonProps{Variant: "primary"}, gui.Text("Save"))
coremd.Button(coremd.ButtonProps{Variant: "send"}, gui.Text("Send"))
coremd.Button(coremd.ButtonProps{Variant: "cancel"}, gui.Text("Stop"))
coremd.Button(coremd.ButtonProps{Variant: "ghost"}, gui.Text("Dismiss"))
coremd.Button(coremd.ButtonProps{Variant: "success-outline"}, gui.Text("Approve"))
coremd.Button(coremd.ButtonProps{Variant: "warning-outline"}, gui.Text("Warn"))
coremd.Button(coremd.ButtonProps{Variant: "info-outline"}, gui.Text("Info"))
```

```html
<button data-variant="send">Send</button>
<button data-variant="cancel">Stop</button>
<button data-variant="ghost">Dismiss</button>
<button data-variant="success-outline">Approve</button>
```

## App Shell & Navigation

Full-page layout with navbar, sidebar, and mobile bottom tab bar.

```go
coremd.AppShell(
    coremd.Navbar(
        coremd.NavBrand("MyApp"),
        coremd.NavLinks(
            coremd.NavLink(coremd.NavLinkProps{Href: "/", Active: true}, gui.Text("Home")),
            coremd.NavLink(coremd.NavLinkProps{Href: "/settings"}, gui.Text("Settings")),
        ),
        rightActions,
    ),
    coremd.HStack("none",
        coremd.Sidebar(
            coremd.SidebarHeader(gui.Text("Conversations"), newBtn),
            listContent,
        ),
        mainContent,
    ),
    coremd.BottomTabBar(
        coremd.TabLink(coremd.TabLinkProps{Href: "/", Active: true}, icon, gui.Text("Home")),
        coremd.TabLink(coremd.TabLinkProps{Href: "/settings"}, icon, gui.Text("Settings")),
    ),
)
```

```html
<div data-app-shell>
  <nav data-navbar>
    <span data-nav-brand>MyApp</span>
    <div data-nav-links>
      <a data-nav-link data-active href="/">Home</a>
      <a data-nav-link href="/settings">Settings</a>
    </div>
  </nav>
  <div data-hstack="none" style="flex:1; overflow:hidden;">
    <aside data-sidebar>
      <div data-sidebar-header><h3>Chats</h3></div>
      <!-- list items -->
    </aside>
    <main style="flex:1;">...</main>
  </div>
  <nav data-bottom-tabbar>
    <a data-tab-link data-active href="/">Home</a>
    <a data-tab-link href="/settings">Settings</a>
  </nav>
</div>
```

The sidebar hides on mobile and shows as an overlay on tablets. The bottom tab bar is visible only on mobile.

## Panels & Overlays

### Panel

A right-side detail panel with header, actions, and mobile full-screen fallback.

```go
coremd.Panel(coremd.PanelProps{},
    coremd.PanelHeader("Services",
        coremd.Button(coremd.ButtonProps{Variant: "toolbar"}, gui.Text("+")),
    ),
    content,
)
```

```html
<div data-panel>
  <div data-panel-header>
    <span data-panel-header-title>Services</span>
    <div data-panel-actions>
      <button data-variant="toolbar">+</button>
    </div>
  </div>
  <!-- content -->
</div>
```

Use `data-panel="wide"` for a centered overlay panel.

### Modal

Centered dialog with backdrop blur. On mobile, slides up as a bottom sheet with a drag handle.

```go
coremd.ModalBackdrop(
    coremd.Modal(
        coremd.DragHandle(),
        coremd.ModalHeader(gui.Text("Confirm"), closeBtn),
        coremd.ModalBody(gui.Text("Are you sure?")),
        coremd.ModalActions(cancelBtn, confirmBtn),
    ),
)
```

```html
<div data-modal-backdrop>
  <div data-modal>
    <div data-drag-handle></div>
    <div data-modal-header>
      <span data-modal-title>Confirm</span>
      <button data-variant="ghost">X</button>
    </div>
    <div data-modal-body>Are you sure?</div>
    <div data-modal-actions>
      <button>Cancel</button>
      <button data-variant="danger">Delete</button>
    </div>
  </div>
</div>
```

### Bottom Sheet

Mobile action menu that slides up from the bottom.

```go
coremd.BottomSheetBackdrop(
    coremd.BottomSheet(
        coremd.DragHandle(),
        coremd.SheetItem(nil, gui.Text("Edit")),
        coremd.SheetItem(nil, gui.Text("Delete"), coremd.SheetDanger),
    ),
)
```

```html
<div data-bottom-sheet-backdrop>
  <div data-bottom-sheet>
    <div data-drag-handle></div>
    <button data-sheet-item>Edit</button>
    <button data-sheet-item data-danger>Delete</button>
  </div>
</div>
```

### Context Menu

A positioned dropdown triggered by right-click or long-press.

```go
coremd.ContextMenuBackdrop(closeHandler,
    coremd.ContextMenu(x, y,
        coremd.ContextMenuItem(nil, gui.Text("Rename")),
        coremd.ContextMenuItem(nil, gui.Text("Delete"), coremd.ContextDanger),
    ),
)
```

```html
<div data-context-menu-backdrop>
  <div data-context-menu style="top:100px; left:200px;">
    <button data-context-menu-item>Rename</button>
    <button data-context-menu-item data-danger>Delete</button>
  </div>
</div>
```

### Search Overlay

Full-screen search with input, results list, and keyboard navigation.

```go
coremd.SearchOverlay(
    coremd.SearchCard(
        coremd.SearchInput("Search files..."),
        coremd.SearchResults(
            coremd.SearchResult(icon, gui.Text("main.go")),
            coremd.SearchResult(icon, gui.Text("utils.go")),
        ),
    ),
)
```

```html
<div data-search-backdrop>
  <div data-search-card>
    <input data-search-input placeholder="Search files..." />
    <div data-search-results>
      <div data-search-result>main.go</div>
      <div data-search-result>utils.go</div>
    </div>
  </div>
</div>
```

## Chat & Messages

### Message Bubble

Renders chat messages with directional styling.

```go
coremd.MessageBubble("", coremd.MessageOutgoing, gui.Text("Hello!"))
coremd.MessageBubble("", coremd.MessageIncoming, richContent)
coremd.MessageBubble("", coremd.MessageError, gui.Text("Failed to send."))
```

```html
<div data-message="outgoing">Hello!</div>
<div data-message="incoming">...</div>
<div data-message="streaming">...</div>
<div data-message="error">Failed to send.</div>
```

Outgoing messages use `--core-on-accent` tokens so links and code remain readable on accent backgrounds.

### Rich Text

Container for rendered Markdown/HTML content with proper spacing and code styling.

```go
coremd.RichText("", renderedMarkdown)
```

```html
<div data-rich-text>
  <p>Some text with <code>inline code</code>.</p>
  <pre><code>fmt.Println("hello")</code></pre>
</div>
```

### Collapsible

Expandable content block, used for thinking indicators and long outputs.

```go
coremd.Collapsible("",
    coremd.CollapsibleSummary(gui.Text("Thinking...")),
    coremd.CollapsibleContent(gui.Text("reasoning steps...")),
)
```

```html
<details data-collapsible>
  <summary data-collapsible-summary>Thinking...</summary>
  <div data-collapsible-content>reasoning steps...</div>
</details>
```

## Lists & Content

### Selectable List

Scrollable list with active/hover states and optional drag reordering.

```go
coremd.List("",
    coremd.SectionHeader("", "Recent"),
    coremd.ListItem(coremd.ListItemProps{Active: true}, gui.Text("Chat 1")),
    coremd.ListItem(coremd.ListItemProps{}, gui.Text("Chat 2")),
)
```

```html
<div data-list>
  <div data-section-header>Recent</div>
  <div data-list-item data-active>Chat 1</div>
  <div data-list-item>Chat 2</div>
</div>
```

Section headers support `success` and `warning` color variants.

### File Tree

Monospace file browser with directory/file distinction.

```go
coremd.FileTree("",
    coremd.FileItem("dir", "src/"),
    coremd.FileItem("file", "main.go"),
)
```

```html
<div data-file-tree>
  <div data-file-item="dir">src/</div>
  <div data-file-item="file">main.go</div>
</div>
```

### Empty State

Centered placeholder for empty views.

```go
coremd.EmptyState("", gui.Text("No conversations yet"), gui.Text("Start a new chat to begin."))
```

```html
<div data-empty-state>
  <h3>No conversations yet</h3>
  <p>Start a new chat to begin.</p>
</div>
```

### Question Prompt

Interactive question block with selectable options.

```go
coremd.QuestionPrompt("",
    coremd.QuestionText(gui.Text("Allow file access?")),
    coremd.QuestionOption(handler, gui.Text("Yes"), gui.Text("Grant read/write access")),
    coremd.QuestionOption(handler, gui.Text("No"), gui.Text("Deny access")),
)
```

```html
<div data-question>
  <div data-question-text>Allow file access?</div>
  <button data-question-option>
    <strong>Yes</strong>
    <small>Grant read/write access</small>
  </button>
</div>
```

### Terminal Panel

Tabbed iframe container for embedded terminals or previews.

```go
coremd.TerminalPanel("",
    coremd.TerminalTabs(
        coremd.TerminalTab(coremd.TerminalTabProps{Active: true}, gui.Text("Terminal 1")),
        coremd.TerminalTab(coremd.TerminalTabProps{}, gui.Text("Terminal 2")),
    ),
    coremd.TerminalIframe("/terminal/1"),
)
```

```html
<div data-terminal-panel>
  <div data-terminal-tabs>
    <button data-terminal-tab data-active>Terminal 1</button>
    <button data-terminal-tab>Terminal 2</button>
  </div>
  <iframe data-terminal-iframe src="/terminal/1"></iframe>
</div>
```

## Forms & Settings

### Feature Row

A checkbox row with label and description, for settings pages.

```go
coremd.FeatureRow(coremd.FeatureRowProps{Checked: true},
    "Auto-save", "Automatically save changes every 30 seconds",
)
```

```html
<div data-feature-row>
  <input type="checkbox" checked />
  <div>
    <div data-label>Auto-save</div>
    <div data-desc>Automatically save changes every 30 seconds</div>
  </div>
</div>
```

### Key-Value Row

Inline key/value input pair with a remove button, for environment variables.

```go
coremd.KVRow("",
    coremd.TextInput(coremd.InputProps{Placeholder: "KEY"}),
    coremd.TextInput(coremd.InputProps{Placeholder: "VALUE"}),
    coremd.KVRemoveButton(removeHandler),
)
```

```html
<div data-kv-row>
  <input type="text" placeholder="KEY" />
  <input type="text" placeholder="VALUE" />
  <button data-kv-remove>x</button>
</div>
```

### Code Textarea

Monospace textarea for code/config editing with height presets.

```go
coremd.CodeTextArea(coremd.CodeTextAreaProps{Size: "tall"}, "package main\n...")
```

```html
<textarea data-code="tall">package main</textarea>
<textarea data-code="medium">...</textarea>
<textarea data-code="short">...</textarea>
```

## Utilities

### Animation

```html
<div data-animate="pulse">Pulsing content</div>
<div data-animate="appear">Appearing element</div>
```

### Responsive Visibility

```html
<div data-hide="mobile">Hidden on small screens</div>
<div data-hide="desktop">Hidden on large screens</div>
<div data-hidden>Always hidden</div>
<div data-visible="false">Also hidden</div>
```

### Micro Badge

A smaller, more compact badge for inline labels.

```html
<span data-micro-badge>beta</span>
```

### Indicator Pill

Monospace status pill, optionally animated.

```go
coremd.Indicator("", gui.Text("processing..."))
coremd.Indicator("working", gui.Text("thinking..."))
```

```html
<span data-indicator>idle</span>
<span data-indicator="working">thinking...</span>
```

### Image Preview

Thumbnail preview strip for attached images.

```html
<div data-image-preview>
  <img data-preview-thumbnail src="photo.jpg" />
</div>
```

### Autocomplete Popup

Dropdown suggestion list for input fields.

```html
<div data-autocomplete>
  <div data-autocomplete-item>Option 1</div>
  <div data-autocomplete-item data-selected>Option 2</div>
</div>
```

### Message Queue

Inline notification strip below an input, auto-hides when empty.

```html
<div data-queue>
  <div data-queue-item>File uploaded: main.go</div>
</div>
```

## Theming

core.md is designed to be themed. A theme is a CSS file that overrides `--core-*` properties and adds styles to `data-*` selectors:

```html
<link rel="stylesheet" href="core.md/styles.css">
<link rel="stylesheet" href="industrial.md/theme.css">
```

See [industrial.md](../industrial.md) for a complete theme example.

## Showcase

<p align="center">
  <img src="../design/screenshots/core-showcase-light.png" width="900" alt="core.md showcase" />
</p>

---

<p align="center">
  <strong>core.md</strong> is part of the <a href="https://github.com/readmedotmd">readme.md</a> project.
</p>
