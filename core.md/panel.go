package coremd

import (
	gui "github.com/readmedotmd/gui.md"
)

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
// CSS classes:
//   - .git-panel: outer container
//   - .git-panel-header: header row
//   - .git-header-actions: action buttons wrapper
//   - .git-tab-bar: tab bar (via TabBar)
//
// Data attributes:
//   - data-expanded: "true" (when expanded)
//   - data-side-panel: on the outer container
//   - data-header: on the header row
func GitPanel(props GitPanelProps, headerActions []gui.Node, content gui.Node) gui.Node {
	base := "git-panel"
	if props.Expanded {
		base = "git-panel git-panel-expanded"
	}
	attrs := collectAttrs(optClass(joinClass(base, props.Class)))
	attrs = append(attrs, dataAttr("side-panel", ""))
	if props.Expanded {
		attrs = append(attrs, dataAttr("expanded", "true"))
	}

	// Header: branch name + action buttons
	headerChildren := []gui.Node{
		gui.Span()(gui.Text(props.Branch)),
	}
	actions := []gui.Node{}
	if props.OnRefresh != nil {
		actions = append(actions, gui.Button(gui.Class("git-header-btn"), gui.OnClick(props.OnRefresh))(gui.Text("\u21bb")))
	}
	if props.OnExpand != nil {
		actions = append(actions, gui.Button(gui.Class("git-header-btn"), gui.OnClick(props.OnExpand))(gui.Text("\u2922")))
	}
	if props.OnClose != nil {
		actions = append(actions, gui.Button(gui.Class("git-header-btn"), gui.OnClick(props.OnClose))(gui.Text("\u00d7")))
	}
	actions = append(actions, headerActions...)
	if len(actions) > 0 {
		headerChildren = append(headerChildren, gui.Div(gui.Class("git-header-actions"))(actions...))
	}

	panelChildren := []gui.Node{
		gui.Div(gui.Class("git-panel-header"), dataAttr("header", ""))(headerChildren...),
	}
	if len(props.Tabs) > 0 {
		panelChildren = append(panelChildren, TabBar("git-tab-bar", props.Tabs))
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
// Data attributes:
//   - data-terminal-panel: on the outer container
//   - data-side-panel: on the outer container
//   - data-terminal-tabs: on the tab bar
//   - data-terminal-tab: on each tab
//   - data-active: "true" (on the active tab)
func TerminalPanel(class string, tabs []TerminalTab, onAddTab func(), terminalContent gui.Node) gui.Node {
	tabNodes := make([]gui.Node, len(tabs))
	for i, tab := range tabs {
		tabAttrs := []gui.Attr{dataAttr("terminal-tab", "")}
		if tab.Active {
			tabAttrs = append(tabAttrs, dataAttr("active", "true"))
		}
		tabChildren := []gui.Node{gui.Span()(gui.Text(tab.Title))}
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

	panelAttrs := collectAttrs(optClass(joinClass("terminal-panel", class)))
	panelAttrs = append(panelAttrs, dataAttr("terminal-panel", ""))
	panelAttrs = append(panelAttrs, dataAttr("side-panel", ""))

	return gui.Div(panelAttrs...)(
		gui.Div(dataAttr("terminal-tabs", ""))(tabBarChildren...),
		gui.Div()(terminalContent),
	)
}

// ─── New Panel Components ───

// GitSectionHeader renders a "Staged Changes" / "Unstaged Changes" label.
//
// CSS classes:
//   - .git-section-header
//   - .git-staged-header (when staged) or .git-unstaged-header
//
// Data attributes:
//   - data-staged: "true" (when staged is true)
//   - data-header: on the element
func GitSectionHeader(class, label string, staged bool) gui.Node {
	stateClass := "git-unstaged-header"
	if staged {
		stateClass = "git-staged-header"
	}
	attrs := collectAttrs(optClass(joinClass("git-section-header "+stateClass, class)))
	attrs = append(attrs, dataAttr("header", ""))
	if staged {
		attrs = append(attrs, dataAttr("staged", "true"))
	}
	return gui.Div(attrs...)(gui.Text(label))
}

// GitFileList renders a scrollable list of git files.
//
// CSS classes:
//   - .git-file-list
func GitFileList(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("git-file-list", class)))...)(children...)
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
//   - data-list-item: on the element
func GitFile(props GitFileProps) gui.Node {
	attrs := collectAttrs(optClass(joinClass("git-file", props.Class)))
	attrs = append(attrs, dataAttr("list-item", ""))
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
		gui.Span(gui.Class("git-file-state"))(gui.Text(props.State)),
		gui.Span(gui.Class("git-file-path"))(gui.Text(props.Path)),
	}
	if props.Desc != "" {
		children = append(children, gui.Span(gui.Class("git-file-desc"))(gui.Text(props.Desc)))
	}
	return gui.Div(attrs...)(children...)
}

// GitCommitArea renders the commit message form area with textarea + buttons.
//
// CSS classes:
//   - .git-commit-area
func GitCommitArea(class string, input gui.Node, actions ...gui.Node) gui.Node {
	children := []gui.Node{}
	if input != nil {
		children = append(children, input)
	}
	if len(actions) > 0 {
		children = append(children, gui.Div()(actions...))
	}
	return gui.Div(collectAttrs(optClass(joinClass("git-commit-area", class)))...)(children...)
}
