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
	attrs := collectAttrs(optClass(props.Class))
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
	return gui.Div(collectAttrs(optClass(class))...)(children...)
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
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// SearchInputField renders a search input field.
func SearchInputField(class, placeholder string, onInput func(gui.Event)) gui.Node {
	attrs := collectAttrs(optClass(class))
	attrs = append(attrs, gui.Type("text"))
	if placeholder != "" {
		attrs = append(attrs, gui.Placeholder(placeholder))
	}
	if onInput != nil {
		attrs = append(attrs, gui.On("input", onInput))
	}
	return gui.Input(attrs...)()
}
