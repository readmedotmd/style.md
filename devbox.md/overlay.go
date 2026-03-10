package devboxmd

import (
	coremd "github.com/readmedotmd/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type ContextMenuItem = coremd.ContextMenuItem

// SearchOverlay renders a themed full-screen search overlay.
func SearchOverlay(tabs []TabBarTab, input gui.Node, results ...gui.Node) gui.Node {
	return theme.SearchOverlay(tabs, input, results...)
}

// ContextMenu renders a themed positioned context menu.
func ContextMenu(x, y int, items []ContextMenuItem) gui.Node {
	return theme.ContextMenu(x, y, items)
}
