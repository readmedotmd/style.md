package industrialmd

import (
	coremd "github.com/readmedotmd/style.md/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	GitPanelProps = coremd.GitPanelProps
	GitFileProps  = coremd.GitFileProps
	SkillCard     = coremd.SkillCard
	TerminalTab   = coremd.TerminalTab
)

// ServicesPanel renders a themed services panel.
func ServicesPanel(title string, headerActions []gui.Node, services ...gui.Node) gui.Node {
	return theme.ServicesPanel(title, headerActions, services...)
}

// RunnerPanel renders a themed runner panel.
func RunnerPanel(title string, runners ...gui.Node) gui.Node {
	return theme.RunnerPanel(title, runners...)
}

// GitPanel renders a themed git panel.
func GitPanel(props GitPanelProps, headerActions []gui.Node, content gui.Node) gui.Node {
	return theme.GitPanel(props, headerActions, content)
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

// GitSectionHeader renders a themed git section header.
func GitSectionHeader(label string, staged bool) gui.Node {
	return theme.GitSectionHeader(label, staged)
}

// GitFileList renders a themed git file list.
func GitFileList(children ...gui.Node) gui.Node {
	return theme.GitFileList(children...)
}

// GitFile renders a themed git file entry.
func GitFile(props GitFileProps) gui.Node {
	return theme.GitFile(props)
}

// GitCommitArea renders a themed commit area.
func GitCommitArea(input gui.Node, actions ...gui.Node) gui.Node {
	return theme.GitCommitArea(input, actions...)
}

// DiffCommentButton renders a themed diff comment button.
func DiffCommentButton(onClick func()) gui.Node {
	return theme.DiffCommentButton(onClick)
}

// DiffInlineComment renders a themed diff inline comment.
func DiffInlineComment(input gui.Node) gui.Node {
	return theme.DiffInlineComment(input)
}

// ServiceActionButton renders a themed service action button.
func ServiceActionButton(icon, variant string, onClick func()) gui.Node {
	return theme.ServiceActionButton(icon, variant, onClick)
}

// RunnerPanelEmpty renders a themed runner panel empty state.
func RunnerPanelEmpty(message string) gui.Node {
	return theme.RunnerPanelEmpty(message)
}
