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

// GitPanelTab represents the active tab in the git panel.
type GitPanelTab string

const (
	GitPanelDiff   GitPanelTab = "diff"
	GitPanelStatus GitPanelTab = "status"
	GitPanelReview GitPanelTab = "review"
)

// GitPanel renders a git panel with tabs and content.
func GitPanel(activeTab GitPanelTab, tabs []TabBarTab, content gui.Node) gui.Node {
	return Panel(PanelProps{Title: "Git"}, nil,
		TabBar("", tabs),
		content,
	)
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

	return gui.Div(collectAttrs(optClass(class))...)(
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
