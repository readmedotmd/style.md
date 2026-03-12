package coremd

import (
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestMarkdownContent(t *testing.T) {
	t.Run("with_class_and_children", func(t *testing.T) {
		s := guitesting.Render(MarkdownContent("prose", gui.Text("Hello world")))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`class="prose"`)
		a.HTMLContains(`data-rich-text`)
		a.TextVisible("Hello world")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(MarkdownContent(""))
		s.Assert(t).HasElement("div").HTMLNotContains("class=").HTMLContains(`data-rich-text`)
	})
}

func TestSectionHeader(t *testing.T) {
	t.Run("with_title_and_actions", func(t *testing.T) {
		action := gui.Span()(gui.Text("Edit"))
		s := guitesting.Render(SectionHeader("sec-hdr", "Settings", action))
		a := s.Assert(t)
		a.HTMLContains(`class="sec-hdr"`)
		a.HTMLContains(`data-section-header`)
		a.TextVisible("Settings")
		a.TextVisible("Edit")
	})
	t.Run("title_only_no_actions", func(t *testing.T) {
		s := guitesting.Render(SectionHeader("", "Title"))
		a := s.Assert(t)
		a.HTMLContains(`data-section-header`)
		a.TextVisible("Title")
		a.HTMLNotContains("class=")
	})
}

func TestCollapsible(t *testing.T) {
	t.Run("open_with_summary_and_children", func(t *testing.T) {
		s := guitesting.Render(Collapsible(CollapsibleProps{
			Class:   "collapse",
			Open:    true,
			Summary: "More info",
		}, gui.Text("details here")))
		a := s.Assert(t)
		a.HasElement("details")
		a.HTMLContains(`class="collapse"`)
		a.HTMLContains(`data-collapsible`)
		a.HTMLContains(`data-open="true"`)
		a.TextVisible("More info")
		a.TextVisible("details here")
	})
	t.Run("closed_no_data_open", func(t *testing.T) {
		s := guitesting.Render(Collapsible(CollapsibleProps{
			Summary: "Toggle",
		}))
		a := s.Assert(t)
		a.HasElement("details")
		a.HTMLContains(`data-collapsible`)
		a.HTMLNotContains("data-open")
		a.TextVisible("Toggle")
	})
}

func TestAnimate(t *testing.T) {
	t.Run("with_animation_and_children", func(t *testing.T) {
		s := guitesting.Render(Animate("anim-cls", "fade-in", gui.Text("content")))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`class="anim-cls"`)
		a.HTMLContains(`data-animate="fade-in"`)
		a.TextVisible("content")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(Animate("", "spin"))
		s.Assert(t).HTMLNotContains("class=").HTMLContains(`data-animate="spin"`)
	})
}
