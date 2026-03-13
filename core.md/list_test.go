package coremd

import (
	"strings"
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestDevboxCard(t *testing.T) {
	t.Run("active_with_all_props", func(t *testing.T) {
		clicked := false
		s := guitesting.Render(DevboxCard(DevboxCardProps{
			Class:   "db",
			Name:    "my-devbox",
			URL:     "https://example.com",
			Status:  StatusRunning,
			Active:  true,
			OnClick: func() { clicked = true },
		},
			ActionTag("", "browser", nil),
			ActionTag("", "vscode", nil),
		))
		a := s.Assert(t)
		a.HTMLContains(`class="devbox-card db"`)
		a.HTMLContains(`data-active="true"`)
		a.TextVisible("my-devbox")
		a.TextVisible("https://example.com")
		a.TextVisible("running")
		a.TextVisible("browser")
		a.TextVisible("vscode")

		root := s.QueryAllByTag("div")[0]
		s.Click(root)
		if !clicked {
			t.Error("expected onClick to fire")
		}
	})

	t.Run("inactive_minimal", func(t *testing.T) {
		s := guitesting.Render(DevboxCard(DevboxCardProps{
			Name:   "test",
			Status: StatusStopped,
		}))
		a := s.Assert(t)
		a.HTMLContains(`class="devbox-card"`)
		a.HTMLNotContains("data-active")
		a.TextVisible("test")
		a.TextVisible("stopped")
	})

	t.Run("no_url_no_tags", func(t *testing.T) {
		s := guitesting.Render(DevboxCard(DevboxCardProps{
			Name:   "bare",
			Status: StatusPending,
		}))
		a := s.Assert(t)
		a.TextVisible("bare")
		a.TextVisible("pending")
		html := s.HTML()
		// Should have header div but no URL or tags divs (only 1 child div)
		if strings.Count(html, "https://") > 0 {
			t.Errorf("expected no URL, got: %s", html)
		}
	})
}

func TestConversationItem(t *testing.T) {
	t.Run("active_with_meta", func(t *testing.T) {
		clicked := false
		s := guitesting.Render(ConversationItem(ConversationItemProps{
			Class:   "ci",
			Title:   "Chat 1",
			Meta:    "2h ago",
			Active:  true,
			OnClick: func() { clicked = true },
		}))
		a := s.Assert(t)
		a.HTMLContains(`class="conv-item ci"`)
		a.HTMLContains(`data-active="true"`)
		a.TextVisible("Chat 1")
		a.TextVisible("2h ago")

		root := s.QueryAllByTag("div")[0]
		s.Click(root)
		if !clicked {
			t.Error("expected onClick to fire")
		}
	})

	t.Run("inactive_no_meta", func(t *testing.T) {
		s := guitesting.Render(ConversationItem(ConversationItemProps{Title: "Chat 2"}))
		a := s.Assert(t)
		a.HTMLNotContains("data-active")
		a.TextVisible("Chat 2")
	})
}

func TestInstanceCard(t *testing.T) {
	t.Run("active_working", func(t *testing.T) {
		clicked := false
		s := guitesting.Render(InstanceCard(InstanceCardProps{
			Class:     "ic",
			Name:      "inst-1",
			Repo:      "github.com/example",
			Status:    StatusRunning,
			Working:   true,
			DoneLabel: "Ready",
			Active:    true,
			Labels:    []gui.Node{gui.Span()(gui.Text("label1"))},
			OnClick:   func() { clicked = true },
		}))
		a := s.Assert(t)
		a.HTMLContains(`data-active="true"`)
		a.HTMLContains(`data-working="true"`)
		a.TextVisible("inst-1")
		a.TextVisible("github.com/example")
		a.TextVisible("Ready")
		a.TextVisible("label1")

		root := s.QueryAllByTag("div")[0]
		s.Click(root)
		if !clicked {
			t.Error("expected onClick to fire")
		}
	})

	t.Run("minimal", func(t *testing.T) {
		s := guitesting.Render(InstanceCard(InstanceCardProps{
			Name:   "inst-2",
			Status: StatusStopped,
		}))
		a := s.Assert(t)
		a.HTMLNotContains("data-active")
		a.HTMLNotContains("data-working")
		a.TextVisible("inst-2")
	})
}

func TestFileTree(t *testing.T) {
	t.Run("with_files_and_dirs", func(t *testing.T) {
		clicked := false
		items := []FileTreeItem{
			{Name: "src", IsDir: true, OnClick: func() { clicked = true }},
			{Name: "main.go", IsDir: false},
		}
		s := guitesting.Render(FileTree("ft", items))
		a := s.Assert(t)
		a.TextVisible("src")
		a.TextVisible("main.go")
		a.HTMLContains(`data-dir="true"`)

		ref := s.GetByText("src")
		s.Click(ref)
		if !clicked {
			t.Error("expected onClick to fire on dir")
		}
	})
}

func TestEnvironmentCard(t *testing.T) {
	t.Run("full_props", func(t *testing.T) {
		stats := []gui.Node{gui.Span()(gui.Text("2"))}
		tags := []gui.Node{gui.Span()(gui.Text("Browser"))}
		actions := []gui.Node{gui.Button()(gui.Text("Edit"))}
		s := guitesting.Render(EnvironmentCard(EnvironmentCardProps{
			Class: "ec",
			Name:  "b3-mono",
		}, stats, tags, actions))
		a := s.Assert(t)
		a.HTMLContains(`class="env-card ec"`)
		a.HTMLContains(`data-env-card`)
		a.HTMLContains(`data-env-card-name`)
		a.HTMLContains(`data-env-card-info`)
		a.HTMLContains(`data-env-card-stats`)
		a.HTMLContains(`data-env-card-tags`)
		a.HTMLContains(`data-env-card-actions`)
		a.TextVisible("b3-mono")
		a.TextVisible("2")
		a.TextVisible("Browser")
		a.TextVisible("Edit")
	})

	t.Run("minimal", func(t *testing.T) {
		s := guitesting.Render(EnvironmentCard(EnvironmentCardProps{Name: "test"}, nil, nil, nil))
		a := s.Assert(t)
		a.HTMLContains(`data-env-card`)
		a.TextVisible("test")
		a.HTMLNotContains("data-env-card-stats")
		a.HTMLNotContains("data-env-card-tags")
		a.HTMLNotContains("data-env-card-actions")
	})
}
