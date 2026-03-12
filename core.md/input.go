package coremd

import (
	gui "github.com/readmedotmd/gui.md"
)

// ChatInputProps configures the ChatInput component.
type ChatInputProps struct {
	Class     string
	OnSend    func(string)
	OnCancel  func()
	Streaming bool
}

// ChatInput renders the chat message input area.
//
// Data attributes:
//   - data-streaming: "true" (when streaming)
func ChatInput(props ChatInputProps) gui.Node {
	attrs := collectAttrs(optClass(joinClass("chat-input", props.Class)))
	if props.Streaming {
		attrs = append(attrs, dataAttr("streaming", "true"))
	}

	toolbarChildren := []gui.Node{
		gui.Div()(),
	}
	if props.Streaming && props.OnCancel != nil {
		toolbarChildren = append(toolbarChildren, Button(ButtonProps{
			Variant: ButtonDanger,
			Size:    ButtonSmall,
			OnClick: props.OnCancel,
		}, gui.Text("Cancel")))
	} else {
		toolbarChildren = append(toolbarChildren, Button(ButtonProps{
			Variant: ButtonPrimary,
			Size:    ButtonSmall,
		}, gui.Text("Send")))
	}

	return gui.Div(attrs...)(
		gui.Div()(
			gui.Textarea(gui.Placeholder("Type a message..."))(),
			gui.Div()(toolbarChildren...),
		),
	)
}

// AutocompleteItem represents a single item in the autocomplete popup.
type AutocompleteItem struct {
	Icon     string
	Label    string
	Detail   string
	Snippet  string
	OnSelect func()
}

// AutocompletePopup renders a popup list of autocomplete suggestions.
//
// Data attributes on items:
//   - data-selected: "true" (on the currently selected item)
func AutocompletePopup(class string, items []AutocompleteItem, selectedIndex int) gui.Node {
	children := make([]gui.Node, len(items))
	for i, item := range items {
		itemAttrs := []gui.Attr{}
		if i == selectedIndex {
			itemAttrs = append(itemAttrs, dataAttr("selected", "true"))
		}
		if item.OnSelect != nil {
			itemAttrs = append(itemAttrs, gui.OnClick(item.OnSelect))
		}
		itemChildren := []gui.Node{}
		if item.Icon != "" {
			itemChildren = append(itemChildren, gui.Span()(gui.I(gui.Class(item.Icon))()))
		}
		itemChildren = append(itemChildren, gui.Span()(gui.Text(item.Label)))
		if item.Detail != "" {
			itemChildren = append(itemChildren, gui.Span()(gui.Text(item.Detail)))
		}
		if item.Snippet != "" {
			itemChildren = append(itemChildren, gui.Span()(gui.Text(item.Snippet)))
		}
		children[i] = gui.Div(itemAttrs...)(itemChildren...)
	}
	return gui.Div(collectAttrs(optClass(joinClass("autocomplete-popup", class)))...)(children...)
}

// MessageQueueItem represents a queued message.
type MessageQueueItem struct {
	Preview  string
	HasImage bool
	OnSend   func()
	OnRemove func()
}

// MessageQueue renders a list of queued messages.
func MessageQueue(class string, items []MessageQueueItem) gui.Node {
	children := make([]gui.Node, len(items))
	for i, item := range items {
		itemChildren := []gui.Node{
			gui.Span()(gui.Text(item.Preview)),
		}
		if item.HasImage {
			itemChildren = append(itemChildren, gui.Span(dataAttr("tag", "image"))(gui.Text("IMG")))
		}
		actionNodes := []gui.Node{}
		if item.OnSend != nil {
			actionNodes = append(actionNodes, Button(ButtonProps{Size: ButtonSmall, OnClick: item.OnSend}, gui.Text("Send")))
		}
		if item.OnRemove != nil {
			actionNodes = append(actionNodes, Button(ButtonProps{Size: ButtonSmall, OnClick: item.OnRemove}, gui.Text("Remove")))
		}
		if len(actionNodes) > 0 {
			itemChildren = append(itemChildren, gui.Div()(actionNodes...))
		}
		children[i] = gui.Div()(itemChildren...)
	}
	return gui.Div(collectAttrs(optClass(joinClass("message-queue", class)))...)(children...)
}

// SearchInputField renders a search input field.
func SearchInputField(class, placeholder string, onInput func(gui.Event)) gui.Node {
	attrs := collectAttrs(optClass(joinClass("search-input-field", class)))
	attrs = append(attrs, gui.Type("text"))
	if placeholder != "" {
		attrs = append(attrs, gui.Placeholder(placeholder))
	}
	if onInput != nil {
		attrs = append(attrs, gui.On("input", onInput))
	}
	return gui.Input(attrs...)()
}

// ─── New Input Components ───

// PastePreviewItem represents a pasted image thumbnail.
type PastePreviewItem struct {
	Src      string
	OnRemove func()
}

// PastePreview renders a row of pasted image thumbnails with remove buttons.
func PastePreview(class string, items []PastePreviewItem) gui.Node {
	children := make([]gui.Node, len(items))
	for i, item := range items {
		removeAttrs := []gui.Attr{}
		if item.OnRemove != nil {
			removeAttrs = append(removeAttrs, gui.OnClick(item.OnRemove))
		}
		children[i] = gui.Div()(
			gui.Img(gui.Src(item.Src), gui.Alt("paste preview"))(),
			gui.Button(removeAttrs...)(gui.Text("\u00d7")),
		)
	}
	return gui.Div(collectAttrs(optClass(joinClass("paste-preview", class)))...)(children...)
}

// ExpandButton renders a small icon button to expand/collapse the textarea.
//
// Data attributes:
//   - data-expanded: "true" (when expanded)
func ExpandButton(class string, expanded bool, onToggle func()) gui.Node {
	attrs := collectAttrs(optClass(joinClass("expand-btn", class)))
	if expanded {
		attrs = append(attrs, dataAttr("expanded", "true"))
	}
	if onToggle != nil {
		attrs = append(attrs, gui.OnClick(onToggle))
	}
	return gui.Button(attrs...)(gui.Text("\u2922"))
}

// AttachButton renders a button to attach files/images.
func AttachButton(class string, onAttach func()) gui.Node {
	attrs := collectAttrs(optClass(joinClass("attach-btn", class)))
	if onAttach != nil {
		attrs = append(attrs, gui.OnClick(onAttach))
	}
	return gui.Button(attrs...)(gui.Text("📎"))
}

// SendButton renders a primary-colored send button.
func SendButton(class, label string, onClick func()) gui.Node {
	attrs := collectAttrs(optClass(class))
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	return gui.Button(attrs...)(gui.Text(label))
}

// CancelButton renders a danger-outlined cancel button.
func CancelButton(class, label string, onClick func()) gui.Node {
	attrs := collectAttrs(optClass(class))
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	return gui.Button(attrs...)(gui.Text(label))
}

// ModeButton renders a plan/act toggle button.
//
// Data attributes:
//   - data-mode: "act" or "plan"
func ModeButton(class, mode string, onClick func()) gui.Node {
	attrs := collectAttrs(optClass(joinClass("mode-btn", class)))
	if mode != "" {
		attrs = append(attrs, dataAttr("mode", mode))
	}
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	label := "Act"
	if mode == "plan" {
		label = "Plan"
	}
	return gui.Button(attrs...)(gui.Text(label))
}

// MessageQueueBar renders a bar above the chat input showing queued messages.
func MessageQueueBar(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("message-queue", class)))...)(children...)
}

// QueuedItem renders a single queued message row with send/remove actions.
//
// Data attributes:
//   - data-has-image: "true" (when hasImage is true)
func QueuedItem(class, text string, hasImage bool, onSend func(), onRemove func()) gui.Node {
	attrs := collectAttrs(optClass(joinClass("queued-item", class)))
	if hasImage {
		attrs = append(attrs, dataAttr("has-image", "true"))
	}
	children := []gui.Node{
		gui.Span()(gui.Text(text)),
	}
	actions := []gui.Node{}
	if onSend != nil {
		actions = append(actions, gui.Button(gui.OnClick(onSend))(gui.Text("Send")))
	}
	if onRemove != nil {
		actions = append(actions, gui.Button(gui.OnClick(onRemove))(gui.Text("\u00d7")))
	}
	if len(actions) > 0 {
		children = append(children, gui.Div()(actions...))
	}
	return gui.Div(attrs...)(children...)
}

// AutocompleteHeader renders a header row showing the trigger character + label.
func AutocompleteHeader(class, trigger, label string) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(
		gui.Span()(gui.Text(trigger)),
		gui.Span()(gui.Text(label)),
	)
}
