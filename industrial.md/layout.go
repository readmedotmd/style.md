package industrialmd

import (
	coremd "github.com/readmedotmd/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	NavbarProps  = coremd.NavbarProps
	SidebarProps = coremd.SidebarProps
	PanelProps   = coremd.PanelProps
)

// AppShell wraps content in the themed app-shell container.
func AppShell(children ...gui.Node) gui.Node {
	return theme.AppShell(children...)
}

// AppShellBody wraps the themed flex body area.
func AppShellBody(children ...gui.Node) gui.Node {
	return theme.AppShellBody(children...)
}

// AppShellMain wraps the themed main scrollable content area.
func AppShellMain(children ...gui.Node) gui.Node {
	return theme.AppShellMain(children...)
}

// Navbar renders a themed top navigation bar.
func Navbar(props NavbarProps, links ...gui.Node) gui.Node {
	return theme.Navbar(props, links...)
}

// Sidebar renders a themed side navigation panel.
func Sidebar(props SidebarProps, header gui.Node, children ...gui.Node) gui.Node {
	return theme.Sidebar(props, header, children...)
}

// SidebarHeader renders a themed sidebar header.
func SidebarHeader(title string, actions ...gui.Node) gui.Node {
	return theme.SidebarHeader(title, actions...)
}

// Panel renders a themed content panel.
func Panel(props PanelProps, actions []gui.Node, children ...gui.Node) gui.Node {
	return theme.Panel(props, actions, children...)
}

// ModalBackdrop renders a themed modal backdrop.
func ModalBackdrop(children ...gui.Node) gui.Node {
	return theme.ModalBackdrop(children...)
}

// Modal renders a themed modal dialog.
func Modal(title string, children ...gui.Node) gui.Node {
	return theme.Modal(title, children...)
}

// ModalBody wraps content in a themed modal body.
func ModalBody(children ...gui.Node) gui.Node {
	return theme.ModalBody(children...)
}

// ModalFooter renders a themed modal footer.
func ModalFooter(children ...gui.Node) gui.Node {
	return theme.ModalFooter(children...)
}

// DragHandle renders a themed drag handle indicator.
func DragHandle() gui.Node {
	return theme.DragHandle()
}
