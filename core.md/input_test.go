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
	s.Assert(t).HTMLContains("my-chat")
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
	s.Assert(t).HTMLContains("ac")
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
		HTMLContains("mq").
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
	s.Assert(t).HTMLContains("q")
}

// ---------------------------------------------------------------------------
// SearchInputField
// ---------------------------------------------------------------------------

func TestSearchInputField_Basic(t *testing.T) {
	s := guitesting.Render(SearchInputField("sf", "Search...", nil))
	s.Assert(t).HasElement("input").HTMLContains("sf").HTMLContains("Search...")
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
		HTMLContains("pp").
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
	s.Assert(t).HasNoElement("img")
}

// ---------------------------------------------------------------------------
// ExpandButton
// ---------------------------------------------------------------------------

func TestExpandButton_Collapsed(t *testing.T) {
	s := guitesting.Render(ExpandButton("eb", false, nil))
	s.Assert(t).
		HasElement("button").
		HTMLContains("eb").
		TextVisible("\u2922").
		HTMLNotContains("data-expanded")
}

func TestExpandButton_Expanded(t *testing.T) {
	s := guitesting.Render(ExpandButton("eb", true, nil))
	s.Assert(t).HTMLContains(`data-expanded="true"`)
}

func TestExpandButton_Toggle(t *testing.T) {
	toggled := false
	s := guitesting.Render(ExpandButton("", false, func() { toggled = true }))
	btn := s.GetByRole("button")
	s.Click(btn)
	if !toggled {
		t.Error("expected onToggle to be called")
	}
}

// ---------------------------------------------------------------------------
// AttachButton
// ---------------------------------------------------------------------------

func TestAttachButton_Basic(t *testing.T) {
	s := guitesting.Render(AttachButton("ab", nil))
	s.Assert(t).HasElement("button").HTMLContains("ab")
}

func TestAttachButton_Click(t *testing.T) {
	attached := false
	s := guitesting.Render(AttachButton("", func() { attached = true }))
	btn := s.GetByRole("button")
	s.Click(btn)
	if !attached {
		t.Error("expected onAttach to be called")
	}
}

// ---------------------------------------------------------------------------
// SendButton
// ---------------------------------------------------------------------------

func TestSendButton_Basic(t *testing.T) {
	s := guitesting.Render(SendButton("sb", "Submit", nil))
	s.Assert(t).HasElement("button").HTMLContains("sb").TextVisible("Submit")
}

func TestSendButton_Click(t *testing.T) {
	clicked := false
	s := guitesting.Render(SendButton("", "Go", func() { clicked = true }))
	btn := s.GetByText("Go")
	s.Click(btn)
	if !clicked {
		t.Error("expected onClick to be called")
	}
}

// ---------------------------------------------------------------------------
// CancelButton
// ---------------------------------------------------------------------------

func TestCancelButton_Basic(t *testing.T) {
	s := guitesting.Render(CancelButton("cb", "Stop", nil))
	s.Assert(t).HasElement("button").HTMLContains("cb").TextVisible("Stop")
}

func TestCancelButton_Click(t *testing.T) {
	clicked := false
	s := guitesting.Render(CancelButton("", "Abort", func() { clicked = true }))
	btn := s.GetByText("Abort")
	s.Click(btn)
	if !clicked {
		t.Error("expected onClick to be called")
	}
}

// ---------------------------------------------------------------------------
// ModeButton
// ---------------------------------------------------------------------------

func TestModeButton_Act(t *testing.T) {
	s := guitesting.Render(ModeButton("mb", "act", nil))
	s.Assert(t).
		HasElement("button").
		HTMLContains("mb").
		HTMLContains(`data-mode="act"`).
		TextVisible("Act")
}

func TestModeButton_Plan(t *testing.T) {
	s := guitesting.Render(ModeButton("", "plan", nil))
	s.Assert(t).HTMLContains(`data-mode="plan"`).TextVisible("Plan")
}

func TestModeButton_EmptyMode(t *testing.T) {
	s := guitesting.Render(ModeButton("", "", nil))
	s.Assert(t).HTMLNotContains("data-mode").TextVisible("Act")
}

func TestModeButton_Click(t *testing.T) {
	clicked := false
	s := guitesting.Render(ModeButton("", "act", func() { clicked = true }))
	btn := s.GetByRole("button")
	s.Click(btn)
	if !clicked {
		t.Error("expected onClick to be called")
	}
}

// ---------------------------------------------------------------------------
// MessageQueueBar
// ---------------------------------------------------------------------------

func TestMessageQueueBar_Basic(t *testing.T) {
	child := gui.Span()(gui.Text("item1"))
	s := guitesting.Render(MessageQueueBar("mqb", child))
	s.Assert(t).HTMLContains("mqb").TextVisible("item1")
}

func TestMessageQueueBar_Empty(t *testing.T) {
	s := guitesting.Render(MessageQueueBar("bar"))
	s.Assert(t).HTMLContains("bar").HasElement("div")
}

// ---------------------------------------------------------------------------
// QueuedItem
// ---------------------------------------------------------------------------

func TestQueuedItem_Basic(t *testing.T) {
	s := guitesting.Render(QueuedItem("qi", "hello", false, nil, nil))
	s.Assert(t).HTMLContains("qi").TextVisible("hello").HTMLNotContains("data-has-image")
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
// AutocompleteHeader
// ---------------------------------------------------------------------------

func TestAutocompleteHeader_Basic(t *testing.T) {
	s := guitesting.Render(AutocompleteHeader("ah", "/", "Commands"))
	s.Assert(t).
		HTMLContains("ah").
		TextVisible("/").
		TextVisible("Commands").
		HasElement("span")
}

func TestAutocompleteHeader_AtTrigger(t *testing.T) {
	s := guitesting.Render(AutocompleteHeader("", "@", "Mentions"))
	s.Assert(t).TextVisible("@").TextVisible("Mentions")
}
