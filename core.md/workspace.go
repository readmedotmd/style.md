package coremd

import gui "github.com/readmedotmd/gui.md"

// WorkspaceStack renders an ordered list of components in a workspace with
// reorder (up/down) and remove controls. Used in workspace editing UI.
//
// CSS classes:
//   - .settings-workspace-stack: outer container
//   - .settings-workspace-stack-header: header row
//   - .settings-workspace-stack-list: item list
//   - .settings-workspace-stack-item: each item row
//   - .settings-workspace-stack-label: item label (name + type)
//   - .settings-workspace-stack-actions: action buttons wrapper
//   - .settings-workspace-add: add button at bottom
type WorkspaceStackItem struct {
	Label      string
	Type       string
	OnMoveUp   func()
	OnMoveDown func()
	OnRemove   func()
}

type WorkspaceStackProps struct {
	Class    string
	Title    string
	Items    []WorkspaceStackItem
	OnAdd    func()
	AddLabel string // defaults to "+ Add Component"
}

func WorkspaceStack(props WorkspaceStackProps) gui.Node {
	headerChildren := []gui.Node{gui.Span()(gui.Text(props.Title))}

	itemNodes := make([]gui.Node, len(props.Items))
	for i, item := range props.Items {
		actions := []gui.Node{}
		if item.OnMoveUp != nil {
			actions = append(actions, gui.Button(gui.OnClick(item.OnMoveUp))(gui.Text("↑")))
		}
		if item.OnMoveDown != nil {
			actions = append(actions, gui.Button(gui.OnClick(item.OnMoveDown))(gui.Text("↓")))
		}
		if item.OnRemove != nil {
			actions = append(actions, gui.Button(gui.OnClick(item.OnRemove))(gui.Text("×")))
		}
		labelText := item.Label
		if item.Type != "" {
			labelText += " (" + item.Type + ")"
		}
		itemNodes[i] = gui.Div(gui.Class("settings-workspace-stack-item"))(
			gui.Span(gui.Class("settings-workspace-stack-label"))(gui.Text(labelText)),
			gui.Div(gui.Class("settings-workspace-stack-actions"))(actions...),
		)
	}

	children := []gui.Node{
		gui.Div(gui.Class("settings-workspace-stack-header"))(headerChildren...),
		gui.Div(gui.Class("settings-workspace-stack-list"))(itemNodes...),
	}

	if props.OnAdd != nil {
		addLabel := props.AddLabel
		if addLabel == "" {
			addLabel = "+ Add Component"
		}
		children = append(children, gui.Button(gui.Class("settings-workspace-add"), gui.OnClick(props.OnAdd))(gui.Text(addLabel)))
	}

	return gui.Div(collectAttrs(optClass(joinClass("settings-workspace-stack", props.Class)))...)(children...)
}
