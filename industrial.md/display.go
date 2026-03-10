package industrialmd

import (
	coremd "github.com/readmedotmd/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	MessageBubbleProps   = coremd.MessageBubbleProps
	QuestionPromptOption = coremd.QuestionPromptOption
	StatusBadgeStatus    = coremd.StatusBadgeStatus
	DiffLine             = coremd.DiffLine
	ClusterStat          = coremd.ClusterStat
)

// Re-export status constants.
const (
	StatusRunning  = coremd.StatusRunning
	StatusStopped  = coremd.StatusStopped
	StatusStarting = coremd.StatusStarting
	StatusPending  = coremd.StatusPending
	StatusError    = coremd.StatusError
)

// MessageBubble renders a themed chat message bubble.
func MessageBubble(props MessageBubbleProps, children ...gui.Node) gui.Node {
	return theme.MessageBubble(props, children...)
}

// ThinkingIndicator renders a themed spinner with a text label.
func ThinkingIndicator(label string) gui.Node {
	return theme.ThinkingIndicator(label)
}

// ThinkingCollapsible renders a themed collapsible thinking section.
func ThinkingCollapsible(label string, children ...gui.Node) gui.Node {
	return theme.ThinkingCollapsible(label, children...)
}

// ToolBadge renders a themed tool badge pill.
func ToolBadge(name string) gui.Node {
	return theme.ToolBadge(name)
}

// QuestionPrompt renders a themed prompt with selectable options.
func QuestionPrompt(question string, options []QuestionPromptOption) gui.Node {
	return theme.QuestionPrompt(question, options)
}

// StatusBadge renders a themed colored status pill.
func StatusBadge(status StatusBadgeStatus, label string) gui.Node {
	return theme.StatusBadge(status, label)
}

// StatusDot renders a themed small colored status dot.
func StatusDot(status StatusBadgeStatus) gui.Node {
	return theme.StatusDot(status)
}

// LabelBadge renders a themed small label with an icon.
func LabelBadge(icon, text string) gui.Node {
	return theme.LabelBadge(icon, text)
}

// UsageBadge renders themed CPU and memory usage indicators.
func UsageBadge(cpu, memory string) gui.Node {
	return theme.UsageBadge(cpu, memory)
}

// DiffViewer renders a themed code diff view.
func DiffViewer(lines []DiffLine) gui.Node {
	return theme.DiffViewer(lines)
}

// DataTable renders a themed data table.
func DataTable(columns []string, rows [][]gui.Node) gui.Node {
	return theme.DataTable(columns, rows)
}

// EmptyState renders a themed empty state placeholder.
func EmptyState(heading, description string) gui.Node {
	return theme.EmptyState(heading, description)
}

// ClusterStatsBar renders a themed row of cluster statistics.
func ClusterStatsBar(stats []ClusterStat, onClick func()) gui.Node {
	return theme.ClusterStatsBar(stats, onClick)
}
