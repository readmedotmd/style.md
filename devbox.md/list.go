package devboxmd

import (
	coremd "github.com/readmedotmd/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	ConversationItemProps = coremd.ConversationItemProps
	InstanceCardProps     = coremd.InstanceCardProps
	InstanceListProps     = coremd.InstanceListProps
	ServiceRowProps       = coremd.ServiceRowProps
	RunnerProcess         = coremd.RunnerProcess
	RunnerRowProps        = coremd.RunnerRowProps
	FileTreeItem          = coremd.FileTreeItem
)

// ConversationItem renders a themed conversation list entry.
func ConversationItem(props ConversationItemProps) gui.Node {
	return theme.ConversationItem(props)
}

// InstanceCard renders a themed instance card.
func InstanceCard(props InstanceCardProps) gui.Node {
	return theme.InstanceCard(props)
}

// InstanceList renders a themed instance list.
func InstanceList(props InstanceListProps, actions []gui.Node, children ...gui.Node) gui.Node {
	return theme.InstanceList(props, actions, children...)
}

// ServiceRow renders a themed service row.
func ServiceRow(props ServiceRowProps, actions ...gui.Node) gui.Node {
	return theme.ServiceRow(props, actions...)
}

// RunnerRow renders a themed runner row.
func RunnerRow(props RunnerRowProps, actions ...gui.Node) gui.Node {
	return theme.RunnerRow(props, actions...)
}

// FileTree renders a themed file tree.
func FileTree(items []FileTreeItem) gui.Node {
	return theme.FileTree(items)
}
