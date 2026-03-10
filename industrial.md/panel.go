package industrialmd

import (
	coremd "github.com/readmedotmd/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	GitPanelTab = coremd.GitPanelTab
	SkillCard   = coremd.SkillCard
	TerminalTab = coremd.TerminalTab
)

// Re-export git panel tab constants.
const (
	GitPanelDiff   = coremd.GitPanelDiff
	GitPanelStatus = coremd.GitPanelStatus
	GitPanelReview = coremd.GitPanelReview
)

// ServicesPanel renders a themed services panel.
func ServicesPanel(title string, headerActions []gui.Node, services ...gui.Node) gui.Node {
	return theme.ServicesPanel(title, headerActions, services...)
}

// RunnerPanel renders a themed runner panel.
func RunnerPanel(title string, runners ...gui.Node) gui.Node {
	return theme.RunnerPanel(title, runners...)
}

// GitPanel renders a themed git panel with tabs.
func GitPanel(activeTab GitPanelTab, tabs []TabBarTab, content gui.Node) gui.Node {
	return theme.GitPanel(activeTab, tabs, content)
}

// SkillsPanel renders a themed skills panel.
func SkillsPanel(skills []SkillCard) gui.Node {
	return theme.SkillsPanel(skills)
}

// TerminalPanel renders a themed terminal panel.
func TerminalPanel(tabs []TerminalTab, onAddTab func(), terminalContent gui.Node) gui.Node {
	return theme.TerminalPanel(tabs, onAddTab, terminalContent)
}

// FileBrowser renders a themed file browser.
func FileBrowser(heading string, items []FileTreeItem) gui.Node {
	return theme.FileBrowser(heading, items)
}
