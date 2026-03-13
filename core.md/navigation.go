package coremd

import (
	gui "github.com/readmedotmd/gui.md"
)

// NavLinkProps configures the NavLink component.
type NavLinkProps struct {
	Class   string
	Icon    string
	Label   string
	Active  bool
	OnClick func()
}

// NavLink renders a navigation link button.
//
// Data attributes:
//   - data-list-item
//   - data-active: "true" (when active)
func NavLink(props NavLinkProps) gui.Node {
	attrs := collectAttrs(optClass(joinClass("nav-link", props.Class)), dataAttr("list-item", ""))
	if props.Active {
		attrs = append(attrs, dataAttr("active", "true"))
	}
	if props.OnClick != nil {
		attrs = append(attrs, gui.OnClick(props.OnClick))
	}
	children := []gui.Node{}
	if props.Icon != "" {
		children = append(children, gui.Span()(gui.I(gui.Class(props.Icon))()))
	}
	children = append(children, gui.Text(props.Label))
	return gui.Button(attrs...)(children...)
}

// TabBarTab represents a single tab in a TabBar.
type TabBarTab struct {
	Label   string
	Active  bool
	OnClick func()
}

// TabBar renders a horizontal tab bar.
//
// Data attributes on tabs:
//   - data-active: "true" (on the active tab)
func TabBar(class string, tabs []TabBarTab) gui.Node {
	children := make([]gui.Node, len(tabs))
	for i, tab := range tabs {
		tabAttrs := []gui.Attr{}
		if tab.Active {
			tabAttrs = append(tabAttrs, dataAttr("active", "true"))
		}
		if tab.OnClick != nil {
			tabAttrs = append(tabAttrs, gui.OnClick(tab.OnClick))
		}
		tabAttrs = append(tabAttrs, gui.Class("tab-bar-item"))
		children[i] = gui.Button(tabAttrs...)(gui.Text(tab.Label))
	}
	return gui.Div(collectAttrs(optClass(joinClass("tab-bar", class)))...)(children...)
}

// BottomTabItem represents a single item in a BottomTabBar.
type BottomTabItem struct {
	Icon    string
	Label   string
	Active  bool
	OnClick func()
}

// BottomTabBar renders a mobile bottom tab bar.
//
// Data attributes on items:
//   - data-active: "true" (on the active item)
func BottomTabBar(class string, items []BottomTabItem) gui.Node {
	children := make([]gui.Node, len(items))
	for i, item := range items {
		itemAttrs := []gui.Attr{}
		if item.Active {
			itemAttrs = append(itemAttrs, dataAttr("active", "true"))
		}
		if item.OnClick != nil {
			itemAttrs = append(itemAttrs, gui.OnClick(item.OnClick))
		}
		children[i] = gui.Button(itemAttrs...)(
			gui.Span()(gui.I(gui.Class(item.Icon))()),
			gui.Span()(gui.Text(item.Label)),
		)
	}
	return gui.Div(collectAttrs(optClass(joinClass("bottom-tab-bar", class)))...)(children...)
}

