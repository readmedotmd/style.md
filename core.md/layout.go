package coremd

import (
	gui "github.com/readmedotmd/gui.md"
)

// AppShell wraps content in the top-level app-shell container.
func AppShell(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// AppShellBody wraps the flex body area of the app shell.
func AppShellBody(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// AppShellMain wraps the main scrollable content area.
func AppShellMain(class string, children ...gui.Node) gui.Node {
	return gui.Main(collectAttrs(optClass(class))...)(children...)
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
		gui.Div()(links...),
	}
	if props.Stats != nil {
		children = append(children, gui.Div()(props.Stats))
	}
	return gui.Nav(collectAttrs(optClass(props.Class))...)(children...)
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
	attrs := collectAttrs(optClass(props.Class))
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
	return gui.Div(collectAttrs(optClass(class))...)(
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
	attrs := collectAttrs(optClass(props.Class))
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
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// Modal renders a modal dialog with a title header.
func Modal(class, title string, children ...gui.Node) gui.Node {
	header := gui.Div()(
		gui.Span()(gui.Text(title)),
	)
	all := []gui.Node{header}
	all = append(all, children...)
	return gui.Div(collectAttrs(optClass(class))...)(all...)
}

// ModalBody wraps content in the modal body area.
func ModalBody(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// ModalFooter renders the bottom action area of a modal.
func ModalFooter(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// DragHandle renders a drag handle indicator.
func DragHandle(class string) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(
		gui.Div()(),
	)
}

// ─── Dashboard Layout Components ───

// DashboardLayout renders a flex row container for sidebar + center column + right panels.
func DashboardLayout(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// SidebarColumn renders a fixed-width sidebar column wrapper.
//
// Data attributes:
//   - data-open: "true" (when open)
func SidebarColumn(class string, open bool, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(class))
	if open {
		attrs = append(attrs, dataAttr("open", "true"))
	}
	return gui.Div(attrs...)(children...)
}

// SidebarOverlay renders a semi-transparent overlay behind the sidebar on tablet.
func SidebarOverlay(class string, onClick func()) gui.Node {
	attrs := collectAttrs(optClass(class))
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	return gui.Div(attrs...)()
}

// CenterColumn renders a flex:1 center column that stacks chat area + terminal panel.
func CenterColumn(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// ─── Chat Layout Components ───

// ChatArea renders a flex column container for chat header + message list + input area.
func ChatArea(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
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
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// MessageList renders a scrollable flex column for messages with gap.
func MessageList(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// ChatInputArea renders a bottom-pinned input area container.
func ChatInputArea(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// ChatInputRow renders a horizontal row holding textarea + send/cancel/mode buttons.
func ChatInputRow(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// ChatInputWrap renders a wrapper around the textarea + expand button.
//
// Data attributes:
//   - data-expanded: "true" (when expanded)
func ChatInputWrap(class string, expanded bool, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(class))
	if expanded {
		attrs = append(attrs, dataAttr("expanded", "true"))
	}
	return gui.Div(attrs...)(children...)
}
