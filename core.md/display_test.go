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
	a.HTMLContains(`class="message my-bubble"`)
	a.TextVisible("hello")
}

func TestMessageBubble_NoStreaming(t *testing.T) {
	s := guitesting.Render(MessageBubble(MessageBubbleProps{Role: "user"}))
	s.Assert(t).HTMLContains(`data-role="user"`).HTMLNotContains(`data-streaming`)
}

func TestQuestionPrompt(t *testing.T) {
	clicked := false
	opts := []QuestionPromptOption{
		{Label: "Yes", Description: "Confirm", OnClick: func() { clicked = true }},
		{Label: "No"},
	}
	s := guitesting.Render(QuestionPrompt("qp-cls", "Continue?", opts))
	a := s.Assert(t)
	a.HTMLContains(`class="question-prompt qp-cls"`)
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
			s.Assert(t).HTMLContains(`data-status="` + string(st) + `"`).TextVisible("lbl").HTMLContains(`class="status-badge sb"`)
		})
	}
}

func TestStatusDot(t *testing.T) {
	s := guitesting.Render(StatusDot("sd", StatusError))
	s.Assert(t).HTMLContains(`data-status="error"`).HTMLContains(`class="status-dot sd"`).HasElement("span")
}

func TestLabelBadge_WithIcon(t *testing.T) {
	s := guitesting.Render(LabelBadge("lb", "icon-star", "Star"))
	a := s.Assert(t)
	a.HTMLContains(`class="label-badge lb"`)
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
	a.HTMLContains(`class="usage-badge ub"`)

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
	a.HTMLContains(`class="diff-viewer dv"`)
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
	a.HTMLContains(`class="data-table dt"`)
	a.TextVisible("Name")
	a.TextVisible("Val")
	a.TextVisible("a")
	a.TextVisible("2")
	a.ElementCount("tr", 3) // 1 header + 2 body
}

func TestEmptyState(t *testing.T) {
	s := guitesting.Render(EmptyState("es", "No items", "Try adding one"))
	a := s.Assert(t)
	a.HTMLContains(`class="empty-state es"`)
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
	a.HTMLContains(`class="cluster-stats-bar cs"`)
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

func TestActionTag(t *testing.T) {
	t.Run("with_click", func(t *testing.T) {
		clicked := false
		s := guitesting.Render(ActionTag("at-cls", "browser", func() { clicked = true }))
		a := s.Assert(t)
		a.HasElement("button")
		a.HTMLContains(`class="action-tag at-cls"`)
		a.HTMLContains(`data-action-tag`)
		a.TextVisible("browser")

		btn := s.GetByRole("button")
		s.Click(btn)
		if !clicked {
			t.Error("expected onClick to fire")
		}
	})

	t.Run("no_click", func(t *testing.T) {
		s := guitesting.Render(ActionTag("", "vscode", nil))
		a := s.Assert(t)
		a.HTMLContains(`class="action-tag"`)
		a.HTMLContains(`data-action-tag`)
		a.TextVisible("vscode")
		html := s.HTML()
		if strings.Contains(html, "onclick") {
			t.Errorf("expected no onclick, got: %s", html)
		}
	})
}

func TestSystemStats(t *testing.T) {
	t.Run("with_items", func(t *testing.T) {
		items := []SystemStatItem{
			{Icon: "icon-cpu", Label: "CPU", Value: "42%"},
			{Label: "Mem", Value: "880MB"},
			{Value: "9.0G / 14.8G"},
		}
		s := guitesting.Render(SystemStats("ss", items))
		a := s.Assert(t)
		a.HTMLContains(`class="system-stats ss"`)
		a.HTMLContains(`data-system-stats`)
		a.TextVisible("42%")
		a.TextVisible("880MB")
		a.TextVisible("9.0G / 14.8G")
		a.HasElement("i") // icon-cpu
	})

	t.Run("empty", func(t *testing.T) {
		s := guitesting.Render(SystemStats("", []SystemStatItem{}))
		a := s.Assert(t)
		a.HTMLContains(`class="system-stats"`)
		a.HTMLContains(`data-system-stats`)
	})
}

func TestDiffPanel(t *testing.T) {
	t.Run("with_file_path_and_content", func(t *testing.T) {
		s := guitesting.Render(DiffPanel(DiffPanelProps{
			Class:    "dp",
			FilePath: "core.md/styles.css",
			Language: "css",
		}, gui.Text("+new line")))
		a := s.Assert(t)
		a.HTMLContains(`class="diff-panel dp"`)
		a.HTMLContains(`data-diff-panel`)
		a.HTMLContains(`data-lang="css"`)
		a.HTMLContains(`data-diff-panel-header`)
		a.HTMLContains(`data-diff-panel-body`)
		a.TextVisible("core.md/styles.css")
		a.TextVisible("+new line")
	})

	t.Run("no_file_path_no_lang", func(t *testing.T) {
		s := guitesting.Render(DiffPanel(DiffPanelProps{}, gui.Text("content")))
		a := s.Assert(t)
		a.HTMLContains(`class="diff-panel"`)
		a.HTMLContains(`data-diff-panel`)
		a.HTMLNotContains("data-lang")
		a.HTMLNotContains("data-diff-panel-header")
		a.HTMLContains(`data-diff-panel-body`)
		a.TextVisible("content")
	})
}

func TestStatChip(t *testing.T) {
	t.Run("with_icon", func(t *testing.T) {
		s := guitesting.Render(StatChip("sc", "icon-var", "4"))
		a := s.Assert(t)
		a.HasElement("span")
		a.HTMLContains(`class="stat-chip sc"`)
		a.HTMLContains(`data-stat-chip`)
		a.HTMLContains(`class="icon-var"`)
		a.TextVisible("4")
	})

	t.Run("no_icon", func(t *testing.T) {
		s := guitesting.Render(StatChip("", "", "12"))
		a := s.Assert(t)
		a.HTMLContains(`class="stat-chip"`)
		a.HTMLContains(`data-stat-chip`)
		a.HasNoElement("i")
		a.TextVisible("12")
	})
}

func TestVariableChip(t *testing.T) {
	t.Run("with_click_and_icon", func(t *testing.T) {
		clicked := false
		s := guitesting.Render(VariableChip(VariableChipProps{
			Class:   "vc",
			Icon:    "icon-lock",
			Label:   "API_KEY",
			OnClick: func() { clicked = true },
		}))
		a := s.Assert(t)
		a.HasElement("button")
		a.HTMLContains(`class="variable-chip vc"`)
		a.HTMLContains(`data-variable-chip`)
		a.HTMLContains(`class="icon-lock"`)
		a.TextVisible("API_KEY")

		btn := s.GetByRole("button")
		s.Click(btn)
		if !clicked {
			t.Error("expected onClick to fire")
		}
	})

	t.Run("no_icon_no_click", func(t *testing.T) {
		s := guitesting.Render(VariableChip(VariableChipProps{Label: "TOKEN"}))
		a := s.Assert(t)
		a.HTMLContains(`class="variable-chip"`)
		a.HTMLContains(`data-variable-chip`)
		a.HasNoElement("i")
		a.TextVisible("TOKEN")
	})
}
