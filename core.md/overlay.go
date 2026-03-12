package coremd

import (
	"fmt"

	gui "github.com/readmedotmd/gui.md"
)

// SearchOverlay renders a full-screen search overlay with tabs, an input, and results.
func SearchOverlay(class string, tabs []TabBarTab, input gui.Node, results ...gui.Node) gui.Node {
	panelChildren := []gui.Node{}
	if len(tabs) > 0 {
		panelChildren = append(panelChildren, TabBar("", tabs))
	}
	if input != nil {
		panelChildren = append(panelChildren, input)
	}
	panelChildren = append(panelChildren, gui.Div()(results...))

	return gui.Div(collectAttrs(optClass(class))...)(
		gui.Div()(panelChildren...),
	)
}

// ContextMenuItem represents a single item in a context menu.
type ContextMenuItem struct {
	Label   string
	Danger  bool
	OnClick func()
}

// ContextMenu renders a positioned context menu.
//
// Data attributes on items:
//   - data-danger: "true" (on danger items)
func ContextMenu(class string, x, y int, items []ContextMenuItem) gui.Node {
	menuItems := make([]gui.Node, len(items))
	for i, item := range items {
		btnAttrs := []gui.Attr{}
		if item.Danger {
			btnAttrs = append(btnAttrs, dataAttr("danger", "true"))
		}
		if item.OnClick != nil {
			btnAttrs = append(btnAttrs, gui.OnClick(item.OnClick))
		}
		menuItems[i] = gui.Button(btnAttrs...)(gui.Text(item.Label))
	}
	style := fmt.Sprintf("left: %dpx; top: %dpx;", x, y)
	attrs := collectAttrs(optClass(class))
	attrs = append(attrs, gui.Style(style))
	return gui.Div(attrs...)(menuItems...)
}

// ─── New Overlay Components ───

// BottomSheetItem represents a single action in a bottom sheet.
type BottomSheetItem struct {
	Icon    string
	Label   string
	Danger  bool
	OnClick func()
}

// BottomSheet renders a mobile action sheet with handle + list of action rows.
//
// Data attributes on items:
//   - data-danger: "true" (on danger items)
func BottomSheet(class string, items []BottomSheetItem) gui.Node {
	rows := make([]gui.Node, len(items))
	for i, item := range items {
		btnAttrs := []gui.Attr{}
		if item.Danger {
			btnAttrs = append(btnAttrs, dataAttr("danger", "true"))
		}
		if item.OnClick != nil {
			btnAttrs = append(btnAttrs, gui.OnClick(item.OnClick))
		}
		children := []gui.Node{}
		if item.Icon != "" {
			children = append(children, gui.I(gui.Class(item.Icon))())
		}
		children = append(children, gui.Text(item.Label))
		rows[i] = gui.Button(btnAttrs...)(children...)
	}
	return gui.Div(collectAttrs(optClass(class))...)(
		gui.Div()(), // handle
		gui.Div()(rows...),
	)
}

// SearchOverlayCard renders the centered modal card used in the search overlay.
func SearchOverlayCard(class string, tabs gui.Node, input gui.Node, results gui.Node) gui.Node {
	children := []gui.Node{}
	if tabs != nil {
		children = append(children, tabs)
	}
	if input != nil {
		children = append(children, input)
	}
	if results != nil {
		children = append(children, results)
	}
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// SearchResult renders a single search result row.
func SearchResult(class, icon, path, text string, onAdd func()) gui.Node {
	children := []gui.Node{}
	if icon != "" {
		children = append(children, gui.I(gui.Class(icon))())
	}
	children = append(children,
		gui.Span()(gui.Text(path)),
		gui.Span()(gui.Text(text)),
	)
	if onAdd != nil {
		children = append(children, gui.Button(gui.OnClick(onAdd))(gui.Text("+")))
	}
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// SearchResultContent renders a content search result with code snippet.
func SearchResultContent(class, path string, snippet gui.Node, onAdd func()) gui.Node {
	children := []gui.Node{
		gui.Span()(gui.Text(path)),
	}
	if snippet != nil {
		children = append(children, snippet)
	}
	if onAdd != nil {
		children = append(children, gui.Button(gui.OnClick(onAdd))(gui.Text("+")))
	}
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// SearchSnippetLine represents a single line in a search snippet.
type SearchSnippetLine struct {
	Text    string
	IsMatch bool
}

// SearchSnippet renders a code snippet with highlighted match lines.
//
// Data attributes on lines:
//   - data-match: "true" (on matching lines)
func SearchSnippet(class string, lines []SearchSnippetLine) gui.Node {
	children := make([]gui.Node, len(lines))
	for i, line := range lines {
		lineAttrs := []gui.Attr{}
		if line.IsMatch {
			lineAttrs = append(lineAttrs, dataAttr("match", "true"))
		}
		children[i] = gui.Div(lineAttrs...)(gui.Text(line.Text))
	}
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}
