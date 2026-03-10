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
