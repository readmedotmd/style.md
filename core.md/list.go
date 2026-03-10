package coremd

import (
	"fmt"

	gui "github.com/readmedotmd/gui.md"
)

// ConversationItemProps configures the ConversationItem component.
type ConversationItemProps struct {
	Class   string
	Title   string
	Meta    string
	Active  bool
	OnClick func()
}

// ConversationItem renders a sidebar conversation list entry.
//
// Data attributes:
//   - data-active: "true" (when active)
func ConversationItem(props ConversationItemProps) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.Active {
		attrs = append(attrs, dataAttr("active", "true"))
	}
	if props.OnClick != nil {
		attrs = append(attrs, gui.OnClick(props.OnClick))
	}
	children := []gui.Node{
		gui.Div()(gui.Text(props.Title)),
	}
	if props.Meta != "" {
		children = append(children, gui.Div()(gui.Text(props.Meta)))
	}
	return gui.Div(attrs...)(children...)
}

// InstanceCardProps configures the InstanceCard component.
type InstanceCardProps struct {
	Class     string
	Name      string
	Repo      string
	Status    StatusBadgeStatus
	Working   bool
	DoneLabel string
	Active    bool
	Labels    []gui.Node
	OnClick   func()
}

// InstanceCard renders a card representing an instance.
//
// Data attributes:
//   - data-active: "true" (when active)
//   - data-working: "true" (when working)
func InstanceCard(props InstanceCardProps) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.Active {
		attrs = append(attrs, dataAttr("active", "true"))
	}
	if props.Working {
		attrs = append(attrs, dataAttr("working", "true"))
	}
	if props.OnClick != nil {
		attrs = append(attrs, gui.OnClick(props.OnClick))
	}

	headerChildren := []gui.Node{
		StatusDot("", props.Status),
		gui.Span()(gui.Text(props.Name)),
	}
	if props.DoneLabel != "" {
		headerChildren = append(headerChildren, gui.Span()(gui.Text(props.DoneLabel)))
	}

	cardChildren := []gui.Node{
		gui.Div()(headerChildren...),
	}

	if props.Repo != "" {
		cardChildren = append(cardChildren, gui.Div()(gui.Text(props.Repo)))
	}

	footerChildren := []gui.Node{}
	if props.Working {
		footerChildren = append(footerChildren, gui.Span()())
	}
	footerChildren = append(footerChildren, props.Labels...)
	if len(footerChildren) > 0 {
		cardChildren = append(cardChildren, gui.Div()(footerChildren...))
	}

	return gui.Div(attrs...)(cardChildren...)
}

// InstanceListProps configures the InstanceList component.
type InstanceListProps struct {
	Class string
	Title string
}

// InstanceList renders a list of instances with a header.
func InstanceList(props InstanceListProps, actions []gui.Node, children ...gui.Node) gui.Node {
	header := gui.Div()(
		gui.Span()(gui.Text(props.Title)),
		gui.Div()(actions...),
	)
	return gui.Div(collectAttrs(optClass(props.Class))...)(
		header,
		gui.Div()(children...),
	)
}

// ServiceRowProps configures the ServiceRow component.
type ServiceRowProps struct {
	Class  string
	Name   string
	Image  string
	Status StatusBadgeStatus
	Ports  []string
}

// ServiceRow renders a row representing a service.
func ServiceRow(props ServiceRowProps, actions ...gui.Node) gui.Node {
	infoChildren := []gui.Node{
		gui.Div()(gui.Text(props.Name)),
	}
	if props.Image != "" {
		infoChildren = append(infoChildren, gui.Div()(gui.Text(props.Image)))
	}

	rowChildren := []gui.Node{
		gui.Div()(StatusDot("", props.Status)),
		gui.Div()(infoChildren...),
	}

	if len(props.Ports) > 0 {
		portNodes := make([]gui.Node, len(props.Ports))
		for i, p := range props.Ports {
			portNodes[i] = gui.Span()(gui.Text(p))
		}
		rowChildren = append(rowChildren, gui.Div()(portNodes...))
	}

	if len(actions) > 0 {
		rowChildren = append(rowChildren, gui.Div()(actions...))
	}

	return gui.Div(collectAttrs(optClass(props.Class))...)(rowChildren...)
}

// RunnerProcess represents a sub-process within a runner row.
type RunnerProcess struct {
	Title   string
	Actions []gui.Node
}

// RunnerRowProps configures the RunnerRow component.
type RunnerRowProps struct {
	Class        string
	Name         string
	Description  string
	ProcessCount int
	Processes    []RunnerProcess
}

// RunnerRow renders a row representing a runner with optional sub-processes.
func RunnerRow(props RunnerRowProps, actions ...gui.Node) gui.Node {
	infoChildren := []gui.Node{
		gui.Div()(gui.Text(props.Name)),
	}
	if props.Description != "" {
		infoChildren = append(infoChildren, gui.Div()(gui.Text(props.Description)))
	}

	headerChildren := []gui.Node{
		gui.Div()(infoChildren...),
	}
	if props.ProcessCount > 0 {
		headerChildren = append(headerChildren, gui.Span()(gui.Text(fmt.Sprintf("%d processes", props.ProcessCount))))
	}
	if len(actions) > 0 {
		headerChildren = append(headerChildren, gui.Div()(actions...))
	}

	rowChildren := []gui.Node{
		gui.Div()(headerChildren...),
	}

	for _, proc := range props.Processes {
		procChildren := []gui.Node{
			gui.Span()(gui.Text(proc.Title)),
		}
		procChildren = append(procChildren, proc.Actions...)
		rowChildren = append(rowChildren, gui.Div()(procChildren...))
	}

	return gui.Div(collectAttrs(optClass(props.Class))...)(rowChildren...)
}

// FileTreeItem represents a single file or directory in a file tree.
type FileTreeItem struct {
	Name    string
	IsDir   bool
	OnClick func()
}

// FileTree renders a list of files and directories.
//
// Data attributes on items:
//   - data-dir: "true" (when item is a directory)
func FileTree(class string, items []FileTreeItem) gui.Node {
	children := make([]gui.Node, len(items))
	for i, item := range items {
		itemAttrs := []gui.Attr{}
		if item.IsDir {
			itemAttrs = append(itemAttrs, dataAttr("dir", "true"))
		}
		if item.OnClick != nil {
			itemAttrs = append(itemAttrs, gui.OnClick(item.OnClick))
		}
		children[i] = gui.Div(itemAttrs...)(gui.Text(item.Name))
	}
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}
