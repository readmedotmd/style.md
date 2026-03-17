package coremd

import (
	"fmt"

	gui "github.com/readmedotmd/gui.md"
)

// SidebarGroupProps configures the SidebarGroup accordion component.
type SidebarGroupProps struct {
	Class    string
	Name     string
	Open     bool
	Count    int
	OnToggle func()
}

// SidebarGroup renders a collapsible group header for sidebar lists.
// When Open is true, children are rendered below the header.
//
// CSS classes:
//   - .sidebar-group: outer container
//   - .sidebar-group-header: clickable header row
//   - .sidebar-group-name: group name label (flex: 1)
//   - .sidebar-group-count: item count pill
//   - .sidebar-group-body: children container (only rendered when open)
//
// Data attributes:
//   - data-open: "true" (when expanded)
func SidebarGroup(props SidebarGroupProps, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(joinClass("sidebar-group", props.Class)))
	if props.Open {
		attrs = append(attrs, dataAttr("open", "true"))
	}

	chevron := "▸"
	if props.Open {
		chevron = "▾"
	}

	headerAttrs := []gui.Attr{gui.Class("sidebar-group-header")}
	if props.OnToggle != nil {
		headerAttrs = append(headerAttrs, gui.OnClick(props.OnToggle))
	}

	headerChildren := []gui.Node{
		gui.Span()(gui.Text(chevron)),
		gui.Span(gui.Class("sidebar-group-name"))(gui.Text(props.Name)),
	}
	if props.Count > 0 {
		headerChildren = append(headerChildren,
			gui.Span(gui.Class("sidebar-group-count"))(gui.Text(fmt.Sprintf("%d", props.Count))),
		)
	}

	groupChildren := []gui.Node{
		gui.Div(headerAttrs...)(headerChildren...),
	}
	if props.Open {
		groupChildren = append(groupChildren, gui.Div(gui.Class("sidebar-group-body"))(children...))
	}

	return gui.Div(attrs...)(groupChildren...)
}
