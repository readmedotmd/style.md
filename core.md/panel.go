package coremd

import (
	gui "github.com/readmedotmd/gui.md"
)

// ServicesPanel renders a panel containing service rows.
func ServicesPanel(title string, headerActions []gui.Node, services ...gui.Node) gui.Node {
	return Panel(PanelProps{Title: title}, headerActions, services...)
}

// RunnerPanel renders a panel containing runner rows.
func RunnerPanel(title string, runners ...gui.Node) gui.Node {
	return Panel(PanelProps{Title: title}, nil, runners...)
}

// GitPanelProps configures the GitPanel component.
type GitPanelProps struct {
	Class     string
	Branch    string
	Expanded  bool
	ActiveTab string
	Tabs      []TabBarTab
	OnClose   func()
	OnExpand  func()
	OnRefresh func()
}

// GitPanel renders a git panel with branch, tabs, header actions, and content.
//
// Data attributes:
//   - data-expanded: "true" (when expanded)
func GitPanel(props GitPanelProps, headerActions []gui.Node, content gui.Node) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.Expanded {
		attrs = append(attrs, dataAttr("expanded", "true"))
	}

	// Header: branch name + action buttons
	headerChildren := []gui.Node{
		gui.Span()(gui.Text(props.Branch)),
	}
	actions := []gui.Node{}
	if props.OnRefresh != nil {
		actions = append(actions, gui.Button(gui.OnClick(props.OnRefresh))(gui.Text("\u21bb")))
	}
	if props.OnExpand != nil {
		actions = append(actions, gui.Button(gui.OnClick(props.OnExpand))(gui.Text("\u2922")))
	}
	if props.OnClose != nil {
		actions = append(actions, gui.Button(gui.OnClick(props.OnClose))(gui.Text("\u00d7")))
	}
	actions = append(actions, headerActions...)
	if len(actions) > 0 {
		headerChildren = append(headerChildren, gui.Div()(actions...))
	}

	panelChildren := []gui.Node{
		gui.Div()(headerChildren...),
	}
	if len(props.Tabs) > 0 {
		panelChildren = append(panelChildren, TabBar("", props.Tabs))
	}
	if content != nil {
		panelChildren = append(panelChildren, gui.Div()(content))
	}

	return gui.Div(attrs...)(panelChildren...)
}

// SkillCard represents a skill in the skills panel.
type SkillCard struct {
	Name        string
	Description string
	OnClick     func()
}

// SkillsPanel renders a panel containing skill cards.
func SkillsPanel(class string, skills []SkillCard) gui.Node {
	children := make([]gui.Node, len(skills))
	for i, skill := range skills {
		cardAttrs := []gui.Attr{}
		if skill.OnClick != nil {
			cardAttrs = append(cardAttrs, gui.OnClick(skill.OnClick))
		}
		children[i] = gui.Div(cardAttrs...)(
			gui.Div()(gui.Text(skill.Name)),
			gui.Div()(gui.Text(skill.Description)),
		)
	}
	return Panel(PanelProps{Title: "Skills"}, nil,
		gui.Div()(children...),
	)
}

// TerminalTab represents a single tab in the terminal panel.
type TerminalTab struct {
	Title   string
	Active  bool
	OnClick func()
	OnClose func()
}

// TerminalPanel renders a terminal panel with tabs and content.
//
// Data attributes on tabs:
//   - data-active: "true" (on the active tab)
func TerminalPanel(class string, tabs []TerminalTab, onAddTab func(), terminalContent gui.Node) gui.Node {
	tabNodes := make([]gui.Node, len(tabs))
	for i, tab := range tabs {
		tabAttrs := []gui.Attr{}
		if tab.Active {
			tabAttrs = append(tabAttrs, dataAttr("active", "true"))
		}
		tabChildren := []gui.Node{gui.Text(tab.Title)}
		if tab.OnClose != nil {
			tabChildren = append(tabChildren, gui.Button(gui.OnClick(tab.OnClose))(gui.Text("x")))
		}
		if tab.OnClick != nil {
			tabAttrs = append(tabAttrs, gui.OnClick(tab.OnClick))
		}
		tabNodes[i] = gui.Div(tabAttrs...)(tabChildren...)
	}

	addBtn := gui.Button(gui.OnClick(onAddTab))(gui.Text("+"))
	tabBarChildren := append(tabNodes, addBtn)

	return gui.Div(collectAttrs(optClass(joinClass("terminal-panel", class)))...)(
		gui.Div()(tabBarChildren...),
		gui.Div()(terminalContent),
	)
}

// FileBrowser renders a file browser panel with a heading and file tree items.
func FileBrowser(heading string, items []FileTreeItem) gui.Node {
	return Panel(PanelProps{Title: heading}, nil,
		FileTree("", items),
	)
}

// ─── New Panel Components ───

// GitSectionHeader renders a "Staged Changes" / "Unstaged Changes" label.
//
// Data attributes:
//   - data-staged: "true" (when staged is true)
func GitSectionHeader(class, label string, staged bool) gui.Node {
	attrs := collectAttrs(optClass(class))
	if staged {
		attrs = append(attrs, dataAttr("staged", "true"))
	}
	return gui.Div(attrs...)(gui.Text(label))
}

// GitFileList renders a scrollable list of git files.
func GitFileList(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// GitFileProps configures the GitFile component.
type GitFileProps struct {
	Class    string
	Path     string
	State    string // "M", "A", "D", "??"
	Staged   bool
	Selected bool
	Desc     string
	OnClick  func()
}

// GitFile renders a single git file entry.
//
// Data attributes:
//   - data-state: "M", "A", "D", "??"
//   - data-staged: "true" (when staged)
//   - data-selected: "true" (when selected)
func GitFile(props GitFileProps) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.State != "" {
		attrs = append(attrs, dataAttr("state", props.State))
	}
	if props.Staged {
		attrs = append(attrs, dataAttr("staged", "true"))
	}
	if props.Selected {
		attrs = append(attrs, dataAttr("selected", "true"))
	}
	if props.OnClick != nil {
		attrs = append(attrs, gui.OnClick(props.OnClick))
	}
	children := []gui.Node{
		gui.Span()(gui.Text(props.State)),
		gui.Span()(gui.Text(props.Path)),
	}
	if props.Desc != "" {
		children = append(children, gui.Span()(gui.Text(props.Desc)))
	}
	return gui.Div(attrs...)(children...)
}

// GitCommitArea renders the commit message form area with textarea + buttons.
func GitCommitArea(class string, input gui.Node, actions ...gui.Node) gui.Node {
	children := []gui.Node{}
	if input != nil {
		children = append(children, input)
	}
	if len(actions) > 0 {
		children = append(children, gui.Div()(actions...))
	}
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// DiffCommentButton renders a gutter button to add inline comment on a diff line.
func DiffCommentButton(class string, onClick func()) gui.Node {
	attrs := collectAttrs(optClass(class))
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	return gui.Button(attrs...)(gui.Text("+"))
}

// DiffInlineComment renders an inline comment input below a diff line.
func DiffInlineComment(class string, input gui.Node) gui.Node {
	children := []gui.Node{}
	if input != nil {
		children = append(children, input)
	}
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// ServiceActionButton renders a colored action button for services (start/stop/restart).
//
// Data attributes:
//   - data-variant: "start", "stop", "restart" or custom variant
func ServiceActionButton(class, icon, variant string, onClick func()) gui.Node {
	attrs := collectAttrs(optClass(class))
	if variant != "" {
		attrs = append(attrs, dataAttr("variant", variant))
	}
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	children := []gui.Node{}
	if icon != "" {
		children = append(children, gui.I(gui.Class(icon))())
	}
	return gui.Button(attrs...)(children...)
}

// RunnerPanelEmpty renders an empty state for the runner panel.
func RunnerPanelEmpty(class, message string) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(gui.Text(message))
}
