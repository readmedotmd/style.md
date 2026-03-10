package devboxmd

import (
	coremd "github.com/readmedotmd/core.md"
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
