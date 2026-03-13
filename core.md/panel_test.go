package coremd

import (
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

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
			HTMLContains(`class="git-panel git-panel-expanded gp"`).
			HTMLContains(`data-expanded="true"`).
			HTMLContains("data-side-panel").
			HTMLContains("data-header").
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
			HTMLContains("data-side-panel").
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
			HTMLContains(`class="terminal-panel tp"`).
			HTMLContains("data-side-panel").
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

func TestGitSectionHeader(t *testing.T) {
	t.Run("staged", func(t *testing.T) {
		screen := guitesting.Render(GitSectionHeader("gsh", "Staged Changes", true))
		screen.Assert(t).
			HTMLContains(`class="git-section-header git-staged-header gsh"`).
			HTMLContains(`data-staged="true"`).
			HTMLContains("data-header").
			TextVisible("Staged Changes")
	})
	t.Run("unstaged", func(t *testing.T) {
		screen := guitesting.Render(GitSectionHeader("", "Unstaged", false))
		screen.Assert(t).
			HTMLNotContains("data-staged").
			HTMLContains("data-header").
			TextVisible("Unstaged")
	})
}

func TestGitFileList(t *testing.T) {
	t.Run("renders_children", func(t *testing.T) {
		screen := guitesting.Render(GitFileList("gfl", gui.Text("f1"), gui.Text("f2")))
		screen.Assert(t).
			HTMLContains(`class="git-file-list gfl"`).
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
			HTMLContains(`class="git-file gf"`).
			HTMLContains("data-list-item").
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
			HTMLContains("data-list-item").
			HTMLNotContains("data-staged").
			HTMLNotContains("data-selected").
			TextVisible("f.go")
	})
}

func TestGitCommitArea(t *testing.T) {
	t.Run("with_input_and_actions", func(t *testing.T) {
		screen := guitesting.Render(GitCommitArea("gca", gui.Text("textarea"), gui.Text("commit"), gui.Text("cancel")))
		screen.Assert(t).
			HTMLContains(`class="git-commit-area gca"`).
			TextVisible("textarea").
			TextVisible("commit").
			TextVisible("cancel")
	})
	t.Run("nil_input_no_actions", func(t *testing.T) {
		screen := guitesting.Render(GitCommitArea("", nil))
		screen.Assert(t).HasElement("div")
	})
}
