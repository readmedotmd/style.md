package coremd

import (
	"strings"
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestMessageBubble(t *testing.T) {
	s := guitesting.Render(MessageBubble(MessageBubbleProps{
		Class: "my-bubble", Role: "assistant", Streaming: true,
	}, gui.Text("hello")))
	a := s.Assert(t)
	a.HTMLContains(`data-role="assistant"`)
	a.HTMLContains(`data-streaming="true"`)
	a.HTMLContains(`class="my-bubble"`)
	a.TextVisible("hello")
}

func TestMessageBubble_NoStreaming(t *testing.T) {
	s := guitesting.Render(MessageBubble(MessageBubbleProps{Role: "user"}))
	s.Assert(t).HTMLContains(`data-role="user"`).HTMLNotContains(`data-streaming`)
}

func TestThinkingIndicator(t *testing.T) {
	s := guitesting.Render(ThinkingIndicator("ti-cls", "Thinking..."))
	a := s.Assert(t)
	a.HTMLContains(`class="ti-cls"`)
	a.TextVisible("Thinking...")
	a.HasElement("div")
}

func TestThinkingCollapsible(t *testing.T) {
	s := guitesting.Render(ThinkingCollapsible("tc-cls", "Details", gui.Text("body")))
	a := s.Assert(t)
	a.HasElement("details")
	a.HasElement("summary")
	a.HTMLContains(`class="tc-cls"`)
	a.TextVisible("Details")
	a.TextVisible("body")
}

func TestToolBadge(t *testing.T) {
	s := guitesting.Render(ToolBadge("tb-cls", "grep"))
	a := s.Assert(t)
	a.HasElement("span")
	a.HTMLContains(`class="tb-cls"`)
	a.TextVisible("grep")
}

func TestQuestionPrompt(t *testing.T) {
	clicked := false
	opts := []QuestionPromptOption{
		{Label: "Yes", Description: "Confirm", OnClick: func() { clicked = true }},
		{Label: "No"},
	}
	s := guitesting.Render(QuestionPrompt("qp-cls", "Continue?", opts))
	a := s.Assert(t)
	a.HTMLContains(`class="qp-cls"`)
	a.TextVisible("Continue?")
	a.TextVisible("Yes")
	a.TextVisible("Confirm")
	a.TextVisible("No")
	a.HasElement("button")

	btn := s.GetByText("Yes")
	// walk up to button
	buttons := s.QueryAllByRole("button")
	for _, b := range buttons {
		if strings.Contains(b.Text(), "Yes") {
			btn = b
			break
		}
	}
	s.Click(btn)
	if !clicked {
		t.Error("expected onClick to fire")
	}
}

func TestStatusBadge(t *testing.T) {
	for _, st := range []StatusBadgeStatus{StatusRunning, StatusStopped, StatusStarting, StatusPending, StatusError} {
		t.Run(string(st), func(t *testing.T) {
			s := guitesting.Render(StatusBadge("sb", st, "lbl"))
			s.Assert(t).HTMLContains(`data-status="` + string(st) + `"`).TextVisible("lbl").HTMLContains(`class="sb"`)
		})
	}
}

func TestStatusDot(t *testing.T) {
	s := guitesting.Render(StatusDot("sd", StatusError))
	s.Assert(t).HTMLContains(`data-status="error"`).HTMLContains(`class="sd"`).HasElement("span")
}

func TestLabelBadge_WithIcon(t *testing.T) {
	s := guitesting.Render(LabelBadge("lb", "icon-star", "Star"))
	a := s.Assert(t)
	a.HTMLContains(`class="lb"`)
	a.HasElement("i")
	a.TextVisible("Star")
}

func TestLabelBadge_NoIcon(t *testing.T) {
	s := guitesting.Render(LabelBadge("", "", "Plain"))
	s.Assert(t).HasNoElement("i").TextVisible("Plain")
}

func TestUsageBadge_WithClick(t *testing.T) {
	clicked := false
	s := guitesting.Render(UsageBadge("ub", "50%", "2G", func() { clicked = true }))
	a := s.Assert(t)
	a.HasElement("button")
	a.TextVisible("50%")
	a.TextVisible("2G")
	a.HTMLContains(`class="ub"`)

	btn := s.GetByRole("button")
	s.Click(btn)
	if !clicked {
		t.Error("expected onClick to fire")
	}
}

func TestUsageBadge_NoClick(t *testing.T) {
	s := guitesting.Render(UsageBadge("", "10%", "1G", nil))
	s.Assert(t).HasNoElement("button").TextVisible("10%").TextVisible("1G")
}

func TestDiffViewer(t *testing.T) {
	lines := []DiffLine{
		{Type: "add", Content: "+new line"},
		{Type: "remove", Content: "-old line"},
		{Type: "context", Content: " ctx"},
		{Type: "header", Content: "@@ hdr @@"},
	}
	s := guitesting.Render(DiffViewer("dv", lines))
	a := s.Assert(t)
	a.HTMLContains(`class="dv"`)
	a.HTMLContains(`data-diff="add"`)
	a.HTMLContains(`data-diff="remove"`)
	a.HTMLContains(`data-diff="context"`)
	a.HTMLContains(`data-diff="header"`)
	a.TextVisible("+new line")
	a.TextVisible("-old line")
}

func TestDataTable(t *testing.T) {
	rows := [][]gui.Node{
		{gui.Text("a"), gui.Text("1")},
		{gui.Text("b"), gui.Text("2")},
	}
	s := guitesting.Render(DataTable("dt", []string{"Name", "Val"}, rows))
	a := s.Assert(t)
	a.HasElement("table")
	a.HasElement("thead")
	a.HasElement("tbody")
	a.HasElement("th")
	a.HasElement("td")
	a.HTMLContains(`class="dt"`)
	a.TextVisible("Name")
	a.TextVisible("Val")
	a.TextVisible("a")
	a.TextVisible("2")
	a.ElementCount("tr", 3) // 1 header + 2 body
}

func TestEmptyState(t *testing.T) {
	s := guitesting.Render(EmptyState("es", "No items", "Try adding one"))
	a := s.Assert(t)
	a.HTMLContains(`class="es"`)
	a.TextVisible("No items")
	a.TextVisible("Try adding one")
}

func TestClusterStatsBar(t *testing.T) {
	stats := []ClusterStat{
		{Icon: "icon-cpu", Label: "CPU", Value: "80%"},
		{Label: "Mem", Value: "4G"},
	}
	s := guitesting.Render(ClusterStatsBar("cs", stats, nil))
	a := s.Assert(t)
	a.HTMLContains(`class="cs"`)
	a.TextVisible("CPU: ")
	a.TextVisible("80%")
	a.TextVisible("Mem: ")
	a.TextVisible("4G")
	a.HasElement("i") // icon-cpu

	// With onClick
	clicked := false
	s2 := guitesting.Render(ClusterStatsBar("", stats, func() { clicked = true }))
	root := s2.QueryAllByTag("div")
	// the outer div should have onclick
	s2.Click(root[0])
	if !clicked {
		t.Error("expected onClick to fire")
	}
}

func TestMessageContent(t *testing.T) {
	s := guitesting.Render(MessageContent("mc", "user", gui.Text("hi")))
	a := s.Assert(t)
	a.HTMLContains(`data-role="user"`)
	a.HTMLContains(`class="mc"`)
	a.TextVisible("hi")
}

func TestMessageContent_NoRole(t *testing.T) {
	s := guitesting.Render(MessageContent("", "", gui.Text("x")))
	s.Assert(t).HTMLNotContains(`data-role`).TextVisible("x")
}

func TestWorkingIndicator(t *testing.T) {
	s := guitesting.Render(WorkingIndicator("wi", "Working..."))
	a := s.Assert(t)
	a.HTMLContains(`class="wi"`)
	a.TextVisible("Working...")
	a.HasElement("span")
}

func TestChatStatusBadge(t *testing.T) {
	s := guitesting.Render(ChatStatusBadge("csb", "Streaming"))
	a := s.Assert(t)
	a.HTMLContains(`class="csb"`)
	a.TextVisible("Streaming")
	a.HasElement("span")
}

func TestThinkingHistory(t *testing.T) {
	s := guitesting.Render(ThinkingHistory("th", "Past thinking", gui.Text("details")))
	a := s.Assert(t)
	a.HasElement("details")
	a.HasElement("summary")
	a.HTMLContains(`class="th"`)
	a.TextVisible("Past thinking")
	a.TextVisible("details")
}

func TestThinkingHistory_NilContent(t *testing.T) {
	s := guitesting.Render(ThinkingHistory("", "Summary", nil))
	a := s.Assert(t)
	a.HasElement("details")
	a.TextVisible("Summary")
}

func TestChatError(t *testing.T) {
	s := guitesting.Render(ChatError("ce", "Something went wrong"))
	a := s.Assert(t)
	a.HTMLContains(`class="ce"`)
	a.TextVisible("Something went wrong")
	a.HasElement("div")
}

func TestAcceptPlanBar(t *testing.T) {
	accepted := false
	s := guitesting.Render(AcceptPlanBar("ap", func() { accepted = true }))
	a := s.Assert(t)
	a.HTMLContains(`class="ap"`)
	a.TextVisible("Accept")
	a.HasElement("button")

	btn := s.GetByRole("button")
	s.Click(btn)
	if !accepted {
		t.Error("expected onAccept to fire")
	}
}

func TestAcceptPlanBar_NilOnAccept(t *testing.T) {
	s := guitesting.Render(AcceptPlanBar("", nil))
	a := s.Assert(t)
	a.HasElement("button")
	a.TextVisible("Accept")
	// button should exist but have no onclick — just verify no panic on render
}
