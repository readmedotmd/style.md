package devboxmd

import (
	coremd "github.com/readmedotmd/style.md/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	ContextMenuItem   = coremd.ContextMenuItem
	BottomSheetItem   = coremd.BottomSheetItem
	SearchSnippetLine = coremd.SearchSnippetLine
)

// SearchOverlay renders a themed full-screen search overlay.
func SearchOverlay(tabs []TabBarTab, input gui.Node, results ...gui.Node) gui.Node {
	return theme.SearchOverlay(tabs, input, results...)
}

// ContextMenu renders a themed positioned context menu.
func ContextMenu(x, y int, items []ContextMenuItem) gui.Node {
	return theme.ContextMenu(x, y, items)
}

// BottomSheet renders a themed mobile bottom sheet.
func BottomSheet(items []BottomSheetItem) gui.Node {
	return theme.BottomSheet(items)
}

// SearchOverlayCard renders a themed search overlay card.
func SearchOverlayCard(tabs gui.Node, input gui.Node, results gui.Node) gui.Node {
	return theme.SearchOverlayCard(tabs, input, results)
}

// SearchResult renders a themed search result row.
func SearchResult(icon, path, text string, onAdd func()) gui.Node {
	return theme.SearchResult(icon, path, text, onAdd)
}

// SearchResultContent renders a themed content search result.
func SearchResultContent(path string, snippet gui.Node, onAdd func()) gui.Node {
	return theme.SearchResultContent(path, snippet, onAdd)
}

// SearchSnippet renders a themed code snippet.
func SearchSnippet(lines []SearchSnippetLine) gui.Node {
	return theme.SearchSnippet(lines)
}
