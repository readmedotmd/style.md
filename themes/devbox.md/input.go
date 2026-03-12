package devboxmd

import (
	coremd "github.com/readmedotmd/style.md/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	ChatInputProps   = coremd.ChatInputProps
	AutocompleteItem = coremd.AutocompleteItem
	MessageQueueItem = coremd.MessageQueueItem
	PastePreviewItem = coremd.PastePreviewItem
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

// PastePreview renders a themed paste preview row.
func PastePreview(items []PastePreviewItem) gui.Node {
	return theme.PastePreview(items)
}

// ExpandButton renders a themed expand/collapse button.
func ExpandButton(expanded bool, onToggle func()) gui.Node {
	return theme.ExpandButton(expanded, onToggle)
}

// AttachButton renders a themed attach button.
func AttachButton(onAttach func()) gui.Node {
	return theme.AttachButton(onAttach)
}

// SendButton renders a themed send button.
func SendButton(label string, onClick func()) gui.Node {
	return theme.SendButton(label, onClick)
}

// CancelButton renders a themed cancel button.
func CancelButton(label string, onClick func()) gui.Node {
	return theme.CancelButton(label, onClick)
}

// ModeButton renders a themed mode toggle button.
func ModeButton(mode string, onClick func()) gui.Node {
	return theme.ModeButton(mode, onClick)
}

// MessageQueueBar renders a themed queue bar.
func MessageQueueBar(children ...gui.Node) gui.Node {
	return theme.MessageQueueBar(children...)
}

// QueuedItem renders a themed queued message row.
func QueuedItem(text string, hasImage bool, onSend func(), onRemove func()) gui.Node {
	return theme.QueuedItem(text, hasImage, onSend, onRemove)
}

// AutocompleteHeader renders a themed autocomplete header.
func AutocompleteHeader(trigger, label string) gui.Node {
	return theme.AutocompleteHeader(trigger, label)
}
