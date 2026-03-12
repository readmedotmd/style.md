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
//   - data-active: "true" (when active)
func NavLink(props NavLinkProps) gui.Node {
	attrs := collectAttrs(optClass(joinClass("nav-link", props.Class)))
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

// ─── New Navigation Components ───

// ChatBackButton renders a mobile-only back arrow button in the chat header.
func ChatBackButton(class string, onClick func()) gui.Node {
	attrs := collectAttrs(optClass(joinClass("chat-back-btn", class)))
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	return gui.Button(attrs...)(gui.Text("\u2190"))
}

// HamburgerButton renders a mobile menu toggle button.
func HamburgerButton(class string, onClick func()) gui.Node {
	attrs := collectAttrs(optClass(class))
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	return gui.Button(attrs...)(gui.Text("\u2630"))
}

// ChatToolbar renders a container that shows desktop toolbar on large screens,
// mobile trigger on small screens.
func ChatToolbar(class string, desktop gui.Node, mobile gui.Node) gui.Node {
	children := []gui.Node{}
	if desktop != nil {
		children = append(children, gui.Div()(desktop))
	}
	if mobile != nil {
		children = append(children, gui.Div()(mobile))
	}
	return gui.Div(collectAttrs(optClass(joinClass("chat-toolbar", class)))...)(children...)
}

// ToolbarButton renders a small outlined button used in the chat toolbar.
//
// Data attributes:
//   - data-danger: "true" (when danger is true)
func ToolbarButton(class, icon, label string, danger bool, onClick func()) gui.Node {
	base := ClassIf("chat-toolbar-btn", danger, "chat-toolbar-btn-danger")
	attrs := collectAttrs(optClass(joinClass(base, class)))
	if danger {
		attrs = append(attrs, dataAttr("danger", "true"))
	}
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	children := []gui.Node{}
	if icon != "" {
		children = append(children, gui.I(gui.Class(icon))())
	}
	if label != "" {
		children = append(children, gui.Text(label))
	}
	return gui.Button(attrs...)(children...)
}
