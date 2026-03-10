package devboxmd

import (
	coremd "github.com/readmedotmd/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	ChatInputProps   = coremd.ChatInputProps
	AutocompleteItem = coremd.AutocompleteItem
	MessageQueueItem = coremd.MessageQueueItem
)

// ChatInput renders a themed chat message input area.
func ChatInput(props ChatInputProps) gui.Node {
	return theme.ChatInput(props)
}

// AutocompletePopup renders a themed autocomplete popup.
func AutocompletePopup(items []AutocompleteItem, selectedIndex int) gui.Node {
	return theme.AutocompletePopup(items, selectedIndex)
}

// MessageQueue renders a themed list of queued messages.
func MessageQueue(items []MessageQueueItem) gui.Node {
	return theme.MessageQueue(items)
}

// SearchInputField renders a themed search input field.
func SearchInputField(placeholder string, onInput func(gui.Event)) gui.Node {
	return theme.SearchInputField(placeholder, onInput)
}
