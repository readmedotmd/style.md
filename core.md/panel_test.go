package coremd

import (
	"strings"
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestServicesPanel(t *testing.T) {
	t.Run("delegates_to_panel", func(t *testing.T) {
		action := gui.Span()(gui.Text("act"))
		screen := guitesting.Render(ServicesPanel("Services", []gui.Node{action}, gui.Text("svc1")))
		screen.Assert(t).
			TextVisible("Services").
			TextVisible("act").
			TextVisible("svc1")
	})
}

func TestRunnerPanel(t *testing.T) {
	t.Run("renders_title_and_runners", func(t *testing.T) {
		screen := guitesting.Render(RunnerPanel("Runners", gui.Text("r1")))
		screen.Assert(t).TextVisible("Runners").TextVisible("r1")
	})
}

func TestGitPanel(t *testing.T) {
	t.Run("expanded_with_branch_and_buttons", func(t *testing.T) {
		refreshed := false
		screen := guitesting.Render(GitPanel(GitPanelProps{
			Class:    "gp",
			Branch:   "main",
			Expanded: true,
			OnRefresh: func() { refreshed = true },
			OnExpand:  func() {},
			OnClose:   func() {},
		}, nil, gui.Text("diff")))
		screen.Assert(t).
			HTMLContains(`class="gp"`).
			HTMLContains(`data-expanded="true"`).
			TextVisible("main").
			TextVisible("diff")
		// Should have refresh, expand, close buttons
		buttons := screen.QueryAllByTag("button")
		if len(buttons) < 3 {
			t.Errorf("expected at least 3 buttons, got %d", len(buttons))
		}
		screen.Click(buttons[0])
		if !refreshed {
			t.Error("expected refresh onclick to fire")
		}
	})
	t.Run("collapsed_no_data_expanded", func(t *testing.T) {
		screen := guitesting.Render(GitPanel(GitPanelProps{Branch: "dev"}, nil, nil))
		screen.Assert(t).
			HTMLNotContains("data-expanded").
			TextVisible("dev")
	})
	t.Run("with_tabs", func(t *testing.T) {
		tabs := []TabBarTab{{Label: "Changes", Active: true}, {Label: "History"}}
		screen := guitesting.Render(GitPanel(GitPanelProps{Branch: "main", Tabs: tabs}, nil, nil))
		screen.Assert(t).TextVisible("Changes").TextVisible("History")
	})
}

func TestSkillsPanel(t *testing.T) {
	t.Run("renders_skill_cards", func(t *testing.T) {
		clicked := false
		skills := []SkillCard{
			{Name: "Deploy", Description: "Deploy to prod", OnClick: func() { clicked = true }},
			{Name: "Test", Description: "Run tests"},
		}
		screen := guitesting.Render(SkillsPanel("sp", skills))
		screen.Assert(t).
			TextVisible("Deploy").
			TextVisible("Deploy to prod").
			TextVisible("Test")
		// Find the skill card div by its text and verify onclick
		ref := screen.QueryByText("Deploy to prod")
		if ref != nil {
			el := ref.Element()
			if _, ok := el.Props["onclick"].(func()); ok {
				screen.Click(ref)
				if !clicked {
					t.Error("expected clicked to be true")
				}
			}
		}
	})
}

func TestTerminalPanel(t *testing.T) {
	t.Run("renders_tabs_and_content", func(t *testing.T) {
		tabs := []TerminalTab{
			{Title: "bash", Active: true, OnClose: func() {}},
			{Title: "node", Active: false},
		}
		screen := guitesting.Render(TerminalPanel("tp", tabs, func() {}, gui.Text("$ ls")))
		screen.Assert(t).
			HTMLContains(`class="tp"`).
			TextVisible("bash").
			TextVisible("node").
			TextVisible("$ ls").
			HTMLContains(`data-active="true"`)
		// Should have close button (x) for bash and add button (+)
		buttons := screen.QueryAllByTag("button")
		if len(buttons) < 2 {
			t.Errorf("expected at least 2 buttons, got %d", len(buttons))
		}
	})
}

func TestFileBrowser(t *testing.T) {
	t.Run("renders_heading_and_items", func(t *testing.T) {
		items := []FileTreeItem{
			{Name: "src", IsDir: true},
			{Name: "main.go", IsDir: false},
		}
		screen := guitesting.Render(FileBrowser("Explorer", items))
		screen.Assert(t).
			TextVisible("Explorer").
			TextVisible("src").
			TextVisible("main.go")
	})
}

func TestGitSectionHeader(t *testing.T) {
	t.Run("staged", func(t *testing.T) {
		screen := guitesting.Render(GitSectionHeader("gsh", "Staged Changes", true))
		screen.Assert(t).
			HTMLContains(`class="gsh"`).
			HTMLContains(`data-staged="true"`).
			TextVisible("Staged Changes")
	})
	t.Run("unstaged", func(t *testing.T) {
		screen := guitesting.Render(GitSectionHeader("", "Unstaged", false))
		screen.Assert(t).
			HTMLNotContains("data-staged").
			TextVisible("Unstaged")
	})
}

func TestGitFileList(t *testing.T) {
	t.Run("renders_children", func(t *testing.T) {
		screen := guitesting.Render(GitFileList("gfl", gui.Text("f1"), gui.Text("f2")))
		screen.Assert(t).
			HTMLContains(`class="gfl"`).
			TextVisible("f1").
			TextVisible("f2")
	})
}

func TestGitFile(t *testing.T) {
	t.Run("full_props", func(t *testing.T) {
		screen := guitesting.Render(GitFile(GitFileProps{
			Class:    "gf",
			Path:     "main.go",
			State:    "M",
			Staged:   true,
			Selected: true,
			Desc:     "modified",
		}))
		screen.Assert(t).
			HTMLContains(`class="gf"`).
			HTMLContains(`data-state="M"`).
			HTMLContains(`data-staged="true"`).
			HTMLContains(`data-selected="true"`).
			TextVisible("main.go").
			TextVisible("M").
			TextVisible("modified")
	})
	t.Run("minimal", func(t *testing.T) {
		screen := guitesting.Render(GitFile(GitFileProps{Path: "f.go", State: "A"}))
		screen.Assert(t).
			HTMLContains(`data-state="A"`).
			HTMLNotContains("data-staged").
			HTMLNotContains("data-selected").
			TextVisible("f.go")
	})
}

func TestGitCommitArea(t *testing.T) {
	t.Run("with_input_and_actions", func(t *testing.T) {
		screen := guitesting.Render(GitCommitArea("gca", gui.Text("textarea"), gui.Text("commit"), gui.Text("cancel")))
		screen.Assert(t).
			HTMLContains(`class="gca"`).
			TextVisible("textarea").
			TextVisible("commit").
			TextVisible("cancel")
	})
	t.Run("nil_input_no_actions", func(t *testing.T) {
		screen := guitesting.Render(GitCommitArea("", nil))
		screen.Assert(t).HasElement("div")
	})
}

func TestDiffCommentButton(t *testing.T) {
	t.Run("renders_plus_button", func(t *testing.T) {
		clicked := false
		screen := guitesting.Render(DiffCommentButton("dcb", func() { clicked = true }))
		screen.Assert(t).
			HTMLContains(`class="dcb"`).
			HasElement("button").
			TextVisible("+")
		ref := screen.QueryAllByTag("button")[0]
		screen.Click(ref)
		if !clicked {
			t.Error("expected onclick to fire")
		}
	})
}

func TestDiffInlineComment(t *testing.T) {
	t.Run("with_input", func(t *testing.T) {
		screen := guitesting.Render(DiffInlineComment("dic", gui.Text("comment")))
		screen.Assert(t).
			HTMLContains(`class="dic"`).
			TextVisible("comment")
	})
	t.Run("nil_input", func(t *testing.T) {
		screen := guitesting.Render(DiffInlineComment("", nil))
		screen.Assert(t).HasElement("div")
	})
}

func TestServiceActionButton(t *testing.T) {
	t.Run("with_variant_and_icon", func(t *testing.T) {
		screen := guitesting.Render(ServiceActionButton("sab", "icon-play", "start", func() {}))
		screen.Assert(t).
			HTMLContains(`class="sab"`).
			HTMLContains(`data-variant="start"`).
			HasElement("button")
	})
	t.Run("no_variant_no_icon", func(t *testing.T) {
		screen := guitesting.Render(ServiceActionButton("", "", "", nil))
		html := screen.HTML()
		if strings.Contains(html, "data-variant") {
			t.Errorf("expected no data-variant, got: %s", html)
		}
	})
}

func TestRunnerPanelEmpty(t *testing.T) {
	t.Run("renders_message", func(t *testing.T) {
		screen := guitesting.Render(RunnerPanelEmpty("rpe", "No runners"))
		screen.Assert(t).
			HTMLContains(`class="rpe"`).
			TextVisible("No runners")
	})
}
