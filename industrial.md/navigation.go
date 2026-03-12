package industrialmd

import (
	coremd "github.com/readmedotmd/style.md/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	NavLinkProps  = coremd.NavLinkProps
	TabBarTab     = coremd.TabBarTab
	BottomTabItem = coremd.BottomTabItem
)

// NavLink renders a themed navigation link.
func NavLink(props NavLinkProps) gui.Node {
	return theme.NavLink(props)
}

// TabBar renders a themed horizontal tab bar.
func TabBar(tabs []TabBarTab) gui.Node {
	return theme.TabBar(tabs)
}

// BottomTabBar renders a themed mobile bottom tab bar.
func BottomTabBar(items []BottomTabItem) gui.Node {
	return theme.BottomTabBar(items)
}

// ChatBackButton renders a themed mobile back button.
func ChatBackButton(onClick func()) gui.Node {
	return theme.ChatBackButton(onClick)
}

// HamburgerButton renders a themed mobile menu button.
func HamburgerButton(onClick func()) gui.Node {
	return theme.HamburgerButton(onClick)
}

// ChatToolbar renders a themed chat toolbar.
func ChatToolbar(desktop gui.Node, mobile gui.Node) gui.Node {
	return theme.ChatToolbar(desktop, mobile)
}

// ToolbarButton renders a themed toolbar button.
func ToolbarButton(icon, label string, danger bool, onClick func()) gui.Node {
	return theme.ToolbarButton(icon, label, danger, onClick)
}
