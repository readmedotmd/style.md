package coremd

import (
	gui "github.com/readmedotmd/gui.md"
)

// MessageBubbleProps configures the MessageBubble component.
type MessageBubbleProps struct {
	Class     string
	Role      string // "user" or "assistant"
	Streaming bool
}

// MessageBubble renders a chat message bubble.
//
// Data attributes:
//   - data-role: "user" or "assistant"
//   - data-streaming: "true" (when streaming)
func MessageBubble(props MessageBubbleProps, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.Role != "" {
		attrs = append(attrs, dataAttr("role", props.Role))
	}
	if props.Streaming {
		attrs = append(attrs, dataAttr("streaming", "true"))
	}
	return gui.Div(attrs...)(
		gui.Div()(children...),
	)
}

// ThinkingIndicator renders a spinner with a text label.
func ThinkingIndicator(class, label string) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(
		Spinner(SpinnerProps{Size: SpinnerSmall}),
		gui.Text(label),
	)
}

// ThinkingCollapsible renders a collapsible thinking section.
func ThinkingCollapsible(class, label string, children ...gui.Node) gui.Node {
	inner := []gui.Node{gui.Summary()(gui.Text(label))}
	inner = append(inner, children...)
	return gui.Details(collectAttrs(optClass(class))...)(inner...)
}

// ToolBadge renders a small pill badge for a tool invocation.
func ToolBadge(class, name string) gui.Node {
	return gui.Span(collectAttrs(optClass(class))...)(
		gui.Span()(),
		gui.Text(name),
	)
}

// QuestionPromptOption represents a selectable option in a QuestionPrompt.
type QuestionPromptOption struct {
	Label       string
	Description string
	OnClick     func()
}

// QuestionPrompt renders a prompt with selectable options.
func QuestionPrompt(class, question string, options []QuestionPromptOption) gui.Node {
	optNodes := make([]gui.Node, len(options))
	for i, opt := range options {
		optChildren := []gui.Node{
			gui.Div()(gui.Text(opt.Label)),
		}
		if opt.Description != "" {
			optChildren = append(optChildren, gui.Div()(gui.Text(opt.Description)))
		}
		btnAttrs := []gui.Attr{}
		if opt.OnClick != nil {
			btnAttrs = append(btnAttrs, gui.OnClick(opt.OnClick))
		}
		optNodes[i] = gui.Button(btnAttrs...)(optChildren...)
	}
	return gui.Div(collectAttrs(optClass(class))...)(
		gui.Div()(gui.Text(question)),
		gui.Div()(optNodes...),
	)
}

// StatusBadgeStatus represents the status of a badge or dot.
type StatusBadgeStatus string

const (
	StatusRunning  StatusBadgeStatus = "running"
	StatusStopped  StatusBadgeStatus = "stopped"
	StatusStarting StatusBadgeStatus = "starting"
	StatusPending  StatusBadgeStatus = "pending"
	StatusError    StatusBadgeStatus = "error"
)

// StatusBadge renders a colored status pill.
//
// Data attributes:
//   - data-status: "running", "stopped", "starting", "pending", "error"
func StatusBadge(class string, status StatusBadgeStatus, label string) gui.Node {
	attrs := collectAttrs(optClass(class))
	attrs = append(attrs, dataAttr("status", string(status)))
	return gui.Span(attrs...)(gui.Text(label))
}

// StatusDot renders a small status dot indicator.
//
// Data attributes:
//   - data-status: "running", "stopped", "starting", "pending", "error"
func StatusDot(class string, status StatusBadgeStatus) gui.Node {
	attrs := collectAttrs(optClass(class))
	attrs = append(attrs, dataAttr("status", string(status)))
	return gui.Span(attrs...)()
}

// LabelBadge renders a small label with an optional icon.
func LabelBadge(class, icon, text string) gui.Node {
	children := []gui.Node{}
	if icon != "" {
		children = append(children, gui.I(gui.Class(icon))())
	}
	children = append(children, gui.Text(text))
	return gui.Span(collectAttrs(optClass(class))...)(children...)
}

// UsageBadge renders CPU and memory usage indicators.
func UsageBadge(class, cpu, memory string) gui.Node {
	return gui.Span(collectAttrs(optClass(class))...)(
		gui.Text(cpu),
		gui.Span()(gui.Text("|")),
		gui.Text(memory),
	)
}

// DiffLine represents a single line in a diff.
type DiffLine struct {
	Type    string // "add", "remove", "header", "context"
	Content string
}

// DiffViewer renders a code diff view.
//
// Data attributes on lines:
//   - data-diff: "add", "remove", "header", "context"
func DiffViewer(class string, lines []DiffLine) gui.Node {
	children := make([]gui.Node, len(lines))
	for i, line := range lines {
		children[i] = gui.Div(dataAttr("diff", line.Type))(gui.Text(line.Content))
	}
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// DataTable renders a simple data table.
func DataTable(class string, columns []string, rows [][]gui.Node) gui.Node {
	ths := make([]gui.Node, len(columns))
	for i, col := range columns {
		ths[i] = gui.Th()(gui.Text(col))
	}
	thead := gui.Thead()(gui.Tr()(ths...))

	trs := make([]gui.Node, len(rows))
	for i, row := range rows {
		tds := make([]gui.Node, len(row))
		for j, cell := range row {
			tds[j] = gui.Td()(cell)
		}
		trs[i] = gui.Tr()(tds...)
	}
	tbody := gui.Tbody()(trs...)

	return gui.Table(collectAttrs(optClass(class))...)(thead, tbody)
}

// EmptyState renders an empty state placeholder.
func EmptyState(class, heading, description string) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(
		gui.Div()(gui.Text(heading)),
		gui.Div()(gui.Text(description)),
	)
}

// ClusterStat represents a single stat item in the ClusterStatsBar.
type ClusterStat struct {
	Icon  string
	Label string
	Value string
}

// ClusterStatsBar renders a row of cluster statistics.
func ClusterStatsBar(class string, stats []ClusterStat, onClick func()) gui.Node {
	children := make([]gui.Node, len(stats))
	for i, stat := range stats {
		itemChildren := []gui.Node{}
		if stat.Icon != "" {
			itemChildren = append(itemChildren, gui.I(gui.Class(stat.Icon))())
		}
		itemChildren = append(itemChildren,
			gui.Text(stat.Label+": "),
			gui.Span()(gui.Text(stat.Value)),
		)
		children[i] = gui.Div()(itemChildren...)
	}
	attrs := collectAttrs(optClass(class))
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	return gui.Div(attrs...)(children...)
}
