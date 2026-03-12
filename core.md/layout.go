package coremd

import (
	gui "github.com/readmedotmd/gui.md"
)

// AppShell wraps content in the top-level app-shell container.
func AppShell(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("app", class)))...)(children...)
}

// AppShellBody wraps the flex body area of the app shell.
func AppShellBody(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("app-shell-body", class)))...)(children...)
}

// AppShellMain wraps the main scrollable content area.
func AppShellMain(class string, children ...gui.Node) gui.Node {
	return gui.Main(collectAttrs(optClass(joinClass("app-shell-main", class)))...)(children...)
}

// NavbarProps configures the Navbar component.
type NavbarProps struct {
	Class string
	Brand string
	Stats gui.Node
}

// Navbar renders a top navigation bar.
func Navbar(props NavbarProps, links ...gui.Node) gui.Node {
	children := []gui.Node{
		gui.Span()(gui.Text(props.Brand)),
		gui.Div(gui.Class("nav-links"))(links...),
	}
	if props.Stats != nil {
		children = append(children, gui.Div()(props.Stats))
	}
	return gui.Nav(collectAttrs(optClass(joinClass("navbar", props.Class)))...)(children...)
}

// SidebarProps configures the Sidebar component.
type SidebarProps struct {
	Class string
	Open  bool
}

// Sidebar renders a side navigation panel.
//
// Data attributes:
//   - data-open: "true" (when open)
func Sidebar(props SidebarProps, header gui.Node, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(joinClass("sidebar", props.Class)))
	if props.Open {
		attrs = append(attrs, dataAttr("open", "true"))
	}
	items := []gui.Node{}
	if header != nil {
		items = append(items, header)
	}
	items = append(items, gui.Div()(children...))
	return gui.Aside(attrs...)(items...)
}

// SidebarHeader renders the header area of a sidebar.
func SidebarHeader(class, title string, actions ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("sidebar-header", class)))...)(
		gui.Span()(gui.Text(title)),
		gui.Div()(actions...),
	)
}

// PanelProps configures the Panel component.
type PanelProps struct {
	Class    string
	Title    string
	Expanded bool
}

// Panel renders a content panel with a header and body.
//
// Data attributes:
//   - data-expanded: "true" (when expanded)
func Panel(props PanelProps, actions []gui.Node, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(joinClass("panel", props.Class)))
	if props.Expanded {
		attrs = append(attrs, dataAttr("expanded", "true"))
	}
	headerChildren := []gui.Node{
		gui.Span()(gui.Text(props.Title)),
	}
	if len(actions) > 0 {
		headerChildren = append(headerChildren, gui.Div()(actions...))
	}
	panelChildren := []gui.Node{
		gui.Div()(headerChildren...),
		gui.Div()(children...),
	}
	return gui.Div(attrs...)(panelChildren...)
}

// ModalBackdrop renders a full-screen backdrop for modals.
func ModalBackdrop(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("modal-backdrop", class)))...)(children...)
}

// Modal renders a modal dialog with a title header.
func Modal(class, title string, children ...gui.Node) gui.Node {
	header := gui.Div()(
		gui.Span()(gui.Text(title)),
	)
	all := []gui.Node{header}
	all = append(all, children...)
	return gui.Div(collectAttrs(optClass(joinClass("modal", class)))...)(all...)
}

// ModalBody wraps content in the modal body area.
func ModalBody(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("modal-body", class)))...)(children...)
}

// ModalFooter renders the bottom action area of a modal.
func ModalFooter(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("modal-footer", class)))...)(children...)
}

// DragHandle renders a drag handle indicator.
func DragHandle(class string) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("drag-handle", class)))...)(
		gui.Div()(),
	)
}

// ─── Dashboard Layout Components ───

// DashboardLayout renders a flex row container for sidebar + center column + right panels.
func DashboardLayout(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("dashboard-layout", class)))...)(children...)
}

// SidebarColumn renders a fixed-width sidebar column wrapper.
//
// Data attributes:
//   - data-open: "true" (when open)
func SidebarColumn(class string, open bool, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(joinClass("sidebar-col", class)))
	if open {
		attrs = append(attrs, dataAttr("open", "true"))
	}
	return gui.Div(attrs...)(children...)
}

// SidebarOverlay renders a semi-transparent overlay behind the sidebar on tablet.
func SidebarOverlay(class string, onClick func()) gui.Node {
	attrs := collectAttrs(optClass(joinClass("sidebar-overlay", class)))
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	return gui.Div(attrs...)()
}

// CenterColumn renders a flex:1 center column that stacks chat area + terminal panel.
func CenterColumn(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("center-col", class)))...)(children...)
}

// ─── Chat Layout Components ───

// ChatArea renders a flex column container for chat header + message list + input area.
func ChatArea(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("chat-area", class)))...)(children...)
}

// ChatHeader renders a fixed-height header bar with title + toolbar buttons.
func ChatHeader(class string, title gui.Node, toolbar gui.Node) gui.Node {
	children := []gui.Node{}
	if title != nil {
		children = append(children, title)
	}
	if toolbar != nil {
		children = append(children, toolbar)
	}
	return gui.Div(collectAttrs(optClass(joinClass("chat-header", class)))...)(children...)
}

// MessageList renders a scrollable flex column for messages with gap.
func MessageList(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("message-list", class)))...)(children...)
}

// ChatInputArea renders a bottom-pinned input area container.
func ChatInputArea(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("chat-input-area", class)))...)(children...)
}

// ChatInputRow renders a horizontal row holding textarea + send/cancel/mode buttons.
func ChatInputRow(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("chat-input-row", class)))...)(children...)
}

// ChatInputWrap renders a wrapper around the textarea + expand button.
//
// Data attributes:
//   - data-expanded: "true" (when expanded)
func ChatInputWrap(class string, expanded bool, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(joinClass("chat-input-wrap", class)))
	if expanded {
		attrs = append(attrs, dataAttr("expanded", "true"))
	}
	return gui.Div(attrs...)(children...)
}

// ─── Generic Containers ───

// BoxProps configures the Box component.
type BoxProps struct {
	Class   string
	Pad     string // padding size: "xs","sm","md","lg","xl"
	Bg      string // background variant: "surface","accent","muted"
	Border  bool   // whether to show border
	Flex    bool   // display:flex
	Rounded bool   // border-radius
}

// Box renders a generic container with configurable padding, background, border, and flex.
//
// Data attributes:
//   - data-box
//   - data-pad: size
//   - data-bg: variant
//   - data-box-border: "true"
//   - data-box-flex: "true"
//   - data-box-rounded: "true"
func Box(props BoxProps, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(joinClass("box", props.Class)), dataAttr("box", ""))
	if props.Pad != "" {
		attrs = append(attrs, dataAttr("pad", props.Pad))
	}
	if props.Bg != "" {
		attrs = append(attrs, dataAttr("bg", props.Bg))
	}
	if props.Border {
		attrs = append(attrs, dataAttr("box-border", "true"))
	}
	if props.Flex {
		attrs = append(attrs, dataAttr("box-flex", "true"))
	}
	if props.Rounded {
		attrs = append(attrs, dataAttr("box-rounded", "true"))
	}
	return gui.Div(attrs...)(children...)
}

// ScrollArea renders a scrollable flex:1 container.
//
// Data attributes:
//   - data-scroll-area
func ScrollArea(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("scroll-area", class)), dataAttr("scroll-area", ""))...)(children...)
}

// SplitLayoutProps configures the SplitLayout component.
type SplitLayoutProps struct {
	Class    string
	Sidebar  string // width of sidebar column, e.g. "260px"
	Panel    string // width of right panel column, e.g. "320px"
}

// SplitLayout renders a three-column layout: fixed sidebar | flex center | fixed panel.
//
// Data attributes:
//   - data-split-layout
func SplitLayout(props SplitLayoutProps, sidebar, center, panel gui.Node) gui.Node {
	attrs := collectAttrs(optClass(joinClass("split-layout", props.Class)), dataAttr("split-layout", ""))
	children := []gui.Node{}
	if sidebar != nil {
		sideAttrs := []gui.Attr{}
		if props.Sidebar != "" {
			sideAttrs = append(sideAttrs, gui.Style("width:"+props.Sidebar+";flex-shrink:0"))
		}
		children = append(children, gui.Div(sideAttrs...)(sidebar))
	}
	if center != nil {
		children = append(children, gui.Div(gui.Style("flex:1;min-width:0"))(center))
	}
	if panel != nil {
		panelAttrs := []gui.Attr{}
		if props.Panel != "" {
			panelAttrs = append(panelAttrs, gui.Style("width:"+props.Panel+";flex-shrink:0"))
		}
		children = append(children, gui.Div(panelAttrs...)(panel))
	}
	return gui.Div(attrs...)(children...)
}

// Backdrop renders a semi-transparent overlay that fills its container.
//
// Data attributes:
//   - data-backdrop
func Backdrop(class string, onClick func()) gui.Node {
	attrs := collectAttrs(optClass(joinClass("backdrop", class)), dataAttr("backdrop", ""))
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	return gui.Div(attrs...)()
}

// IconButton renders a compact icon-only button (ghost variant).
//
// Data attributes:
//   - data-icon-button
func IconButton(class, icon, ariaLabel string, onClick func()) gui.Node {
	attrs := collectAttrs(optClass(joinClass("icon-button", class)), dataAttr("icon-button", ""))
	if ariaLabel != "" {
		attrs = append(attrs, gui.Attr_("aria-label", ariaLabel))
	}
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	return gui.Button(attrs...)(gui.I(gui.Class(icon))())
}

// Toolbar renders a horizontal row of small action buttons.
//
// Data attributes:
//   - data-toolbar
func Toolbar(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("toolbar", class)), dataAttr("toolbar", ""))...)(children...)
}
