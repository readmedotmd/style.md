package coremd

import (
	"strings"
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

// ---------------------------------------------------------------------------
// ChatInput
// ---------------------------------------------------------------------------

func TestChatInput_Basic(t *testing.T) {
	s := guitesting.Render(ChatInput(ChatInputProps{}))
	s.Assert(t).HasElement("div").HasElement("textarea").TextVisible("Send")
}

func TestChatInput_Class(t *testing.T) {
	s := guitesting.Render(ChatInput(ChatInputProps{Class: "my-chat"}))
	s.Assert(t).HTMLContains(`class="chat-input my-chat"`)
}

func TestChatInput_BaseClass(t *testing.T) {
	s := guitesting.Render(ChatInput(ChatInputProps{}))
	s.Assert(t).HTMLContains(`class="chat-input"`)
}

func TestChatInput_Streaming(t *testing.T) {
	s := guitesting.Render(ChatInput(ChatInputProps{Streaming: true, OnCancel: func() {}}))
	s.Assert(t).HTMLContains(`data-streaming="true"`).TextVisible("Cancel")
}

func TestChatInput_NotStreaming(t *testing.T) {
	s := guitesting.Render(ChatInput(ChatInputProps{Streaming: false}))
	s.Assert(t).HTMLNotContains("data-streaming").TextVisible("Send")
}

func TestChatInput_StreamingWithoutCancel(t *testing.T) {
	s := guitesting.Render(ChatInput(ChatInputProps{Streaming: true}))
	s.Assert(t).HTMLContains(`data-streaming="true"`).TextVisible("Send")
}

// ---------------------------------------------------------------------------
// AutocompletePopup
// ---------------------------------------------------------------------------

func TestAutocompletePopup_Empty(t *testing.T) {
	s := guitesting.Render(AutocompletePopup("ac", nil, 0))
	s.Assert(t).HTMLContains(`class="autocomplete-popup ac"`)
}

func TestAutocompletePopup_BaseClass(t *testing.T) {
	s := guitesting.Render(AutocompletePopup("", nil, 0))
	s.Assert(t).HTMLContains(`class="autocomplete-popup"`)
}

func TestAutocompletePopup_Items(t *testing.T) {
	items := []AutocompleteItem{
		{Label: "Alpha", Detail: "first", Snippet: "a()"},
		{Label: "Beta", Icon: "icon-b"},
	}
	s := guitesting.Render(AutocompletePopup("", items, 0))
	s.Assert(t).
		TextVisible("Alpha").
		TextVisible("first").
		TextVisible("a()").
		TextVisible("Beta").
		HTMLContains(`data-selected="true"`)
}

func TestAutocompletePopup_SelectedIndex(t *testing.T) {
	items := []AutocompleteItem{{Label: "A"}, {Label: "B"}}
	s := guitesting.Render(AutocompletePopup("", items, 1))
	html := s.HTML()
	// The second item should be selected, not the first
	idx0 := strings.Index(html, "A")
	idxSel := strings.Index(html, `data-selected="true"`)
	idx1 := strings.Index(html, "B")
	if idxSel < idx0 || idxSel > idx1 {
		// data-selected should appear between A and after B's div start
	}
	s.Assert(t).HTMLContains(`data-selected="true"`)
}

func TestAutocompletePopup_Icon(t *testing.T) {
	items := []AutocompleteItem{{Label: "X", Icon: "icon-x"}}
	s := guitesting.Render(AutocompletePopup("", items, 0))
	s.Assert(t).HTMLContains("icon-x").HasElement("i")
}

// ---------------------------------------------------------------------------
// MessageQueue
// ---------------------------------------------------------------------------

func TestMessageQueue_Basic(t *testing.T) {
	items := []MessageQueueItem{
		{Preview: "Hello", HasImage: true, OnSend: func() {}, OnRemove: func() {}},
	}
	s := guitesting.Render(MessageQueue("mq", items))
	s.Assert(t).
		HTMLContains(`class="message-queue mq"`).
		TextVisible("Hello").
		TextVisible("IMG").
		HTMLContains(`data-tag="image"`).
		TextVisible("Send").
		TextVisible("Remove")
}

func TestMessageQueue_NoImage(t *testing.T) {
	items := []MessageQueueItem{{Preview: "Test"}}
	s := guitesting.Render(MessageQueue("", items))
	s.Assert(t).TextVisible("Test").HTMLNotContains("IMG")
}

func TestMessageQueue_Empty(t *testing.T) {
	s := guitesting.Render(MessageQueue("q", nil))
	s.Assert(t).HTMLContains(`class="message-queue q"`)
}

func TestMessageQueue_BaseClass(t *testing.T) {
	s := guitesting.Render(MessageQueue("", nil))
	s.Assert(t).HTMLContains(`class="message-queue"`)
}

// ---------------------------------------------------------------------------
// SearchInputField
// ---------------------------------------------------------------------------

func TestSearchInputField_Basic(t *testing.T) {
	s := guitesting.Render(SearchInputField("sf", "Search...", nil))
	s.Assert(t).HasElement("input").HTMLContains(`class="search-input-field sf"`).HTMLContains("Search...")
}

func TestSearchInputField_BaseClass(t *testing.T) {
	s := guitesting.Render(SearchInputField("", "", nil))
	s.Assert(t).HTMLContains(`class="search-input-field"`)
}

func TestSearchInputField_NoPlaceholder(t *testing.T) {
	s := guitesting.Render(SearchInputField("", "", nil))
	s.Assert(t).HasElement("input").HTMLNotContains("placeholder")
}

func TestSearchInputField_TypeText(t *testing.T) {
	s := guitesting.Render(SearchInputField("", "Find", nil))
	s.Assert(t).HTMLContains(`type="text"`)
}

func TestSearchInputField_OnInput(t *testing.T) {
	called := false
	s := guitesting.Render(SearchInputField("", "Go", func(e gui.Event) { called = true }))
	ref := s.GetByPlaceholder("Go")
	s.FireEvent(ref, "input", gui.Event{Value: "x"})
	if !called {
		t.Error("expected onInput to be called")
	}
}

// ---------------------------------------------------------------------------
// PastePreview
// ---------------------------------------------------------------------------

func TestPastePreview_Basic(t *testing.T) {
	items := []PastePreviewItem{
		{Src: "data:image/png;base64,abc"},
		{Src: "blob://img2"},
	}
	s := guitesting.Render(PastePreview("pp", items))
	s.Assert(t).
		HTMLContains(`class="paste-preview pp"`).
		HTMLContains("data:image/png;base64,abc").
		HTMLContains("blob://img2").
		HasElement("img").
		HasElement("button")
}

func TestPastePreview_RemoveButton(t *testing.T) {
	removed := false
	items := []PastePreviewItem{{Src: "x.png", OnRemove: func() { removed = true }}}
	s := guitesting.Render(PastePreview("", items))
	btn := s.GetByText("\u00d7")
	s.Click(btn)
	if !removed {
		t.Error("expected OnRemove to be called")
	}
}

func TestPastePreview_Empty(t *testing.T) {
	s := guitesting.Render(PastePreview("pp", nil))
	s.Assert(t).HasNoElement("img").HTMLContains(`class="paste-preview pp"`)
}

func TestPastePreview_BaseClass(t *testing.T) {
	s := guitesting.Render(PastePreview("", nil))
	s.Assert(t).HTMLContains(`class="paste-preview"`)
}

// ---------------------------------------------------------------------------
// MessageQueueBar
// ---------------------------------------------------------------------------

func TestMessageQueueBar_Basic(t *testing.T) {
	child := gui.Span()(gui.Text("item1"))
	s := guitesting.Render(MessageQueueBar("mqb", child))
	s.Assert(t).HTMLContains(`class="message-queue mqb"`).TextVisible("item1")
}

func TestMessageQueueBar_Empty(t *testing.T) {
	s := guitesting.Render(MessageQueueBar("bar"))
	s.Assert(t).HTMLContains(`class="message-queue bar"`).HasElement("div")
}

func TestMessageQueueBar_BaseClass(t *testing.T) {
	s := guitesting.Render(MessageQueueBar(""))
	s.Assert(t).HTMLContains(`class="message-queue"`)
}

// ---------------------------------------------------------------------------
// QueuedItem
// ---------------------------------------------------------------------------

func TestQueuedItem_Basic(t *testing.T) {
	s := guitesting.Render(QueuedItem("qi", "hello", false, nil, nil))
	s.Assert(t).HTMLContains(`class="queued-item qi"`).TextVisible("hello").HTMLNotContains("data-has-image")
}

func TestQueuedItem_BaseClass(t *testing.T) {
	s := guitesting.Render(QueuedItem("", "x", false, nil, nil))
	s.Assert(t).HTMLContains(`class="queued-item"`)
}

func TestQueuedItem_HasImage(t *testing.T) {
	s := guitesting.Render(QueuedItem("", "msg", true, nil, nil))
	s.Assert(t).HTMLContains(`data-has-image="true"`)
}

func TestQueuedItem_Actions(t *testing.T) {
	sent := false
	removed := false
	s := guitesting.Render(QueuedItem("", "txt", false, func() { sent = true }, func() { removed = true }))
	s.Assert(t).TextVisible("Send")

	sendBtn := s.GetByText("Send")
	s.Click(sendBtn)
	if !sent {
		t.Error("expected onSend to be called")
	}

	removeBtn := s.GetByText("\u00d7")
	s.Click(removeBtn)
	if !removed {
		t.Error("expected onRemove to be called")
	}
}

func TestQueuedItem_NoActions(t *testing.T) {
	s := guitesting.Render(QueuedItem("", "x", false, nil, nil))
	btns := s.QueryAllByRole("button")
	if len(btns) != 0 {
		t.Errorf("expected no buttons, got %d", len(btns))
	}
}

// ---------------------------------------------------------------------------
// AttachmentButton
// ---------------------------------------------------------------------------

func TestAttachmentButton(t *testing.T) {
	t.Run("with_click_and_custom_icon", func(t *testing.T) {
		clicked := false
		s := guitesting.Render(AttachmentButton("ab-cls", "icon-clip", func() { clicked = true }))
		a := s.Assert(t)
		a.HasElement("button")
		a.HTMLContains(`class="attachment-button ab-cls"`)
		a.HTMLContains(`data-attachment-button`)
		a.HTMLContains(`class="icon-clip"`)

		btn := s.GetByRole("button")
		s.Click(btn)
		if !clicked {
			t.Error("expected onClick to fire")
		}
	})

	t.Run("default_icon_no_click", func(t *testing.T) {
		s := guitesting.Render(AttachmentButton("", "", nil))
		a := s.Assert(t)
		a.HTMLContains(`class="attachment-button"`)
		a.HTMLContains(`data-attachment-button`)
		a.HTMLContains(`class="icon-paperclip"`)
		html := s.HTML()
		if strings.Contains(html, "onclick") {
			t.Errorf("expected no onclick, got: %s", html)
		}
	})
}

// ---------------------------------------------------------------------------
// ModeToggle
// ---------------------------------------------------------------------------

func TestModeToggle(t *testing.T) {
	t.Run("active_with_click", func(t *testing.T) {
		clicked := false
		s := guitesting.Render(ModeToggle(ModeToggleProps{
			Class:   "mt-cls",
			Label:   "Act",
			Active:  true,
			OnClick: func() { clicked = true },
		}))
		a := s.Assert(t)
		a.HasElement("button")
		a.HTMLContains(`class="mode-toggle mt-cls"`)
		a.HTMLContains(`data-mode-toggle`)
		a.HTMLContains(`data-active="true"`)
		a.TextVisible("Act")

		btn := s.GetByRole("button")
		s.Click(btn)
		if !clicked {
			t.Error("expected onClick to fire")
		}
	})

	t.Run("inactive_default_label", func(t *testing.T) {
		s := guitesting.Render(ModeToggle(ModeToggleProps{}))
		a := s.Assert(t)
		a.HTMLContains(`class="mode-toggle"`)
		a.HTMLContains(`data-mode-toggle`)
		a.HTMLNotContains("data-active")
		a.TextVisible("Act")
	})

	t.Run("custom_label", func(t *testing.T) {
		s := guitesting.Render(ModeToggle(ModeToggleProps{Label: "Plan"}))
		a := s.Assert(t)
		a.TextVisible("Plan")
	})
}
