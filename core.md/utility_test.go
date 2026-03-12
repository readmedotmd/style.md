package coremd

import (
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestSpinner(t *testing.T) {
	t.Run("with_size", func(t *testing.T) {
		screen := guitesting.Render(Spinner(SpinnerProps{Class: "sp", Size: SpinnerSmall}))
		screen.Assert(t).
			HTMLContains(`class="sp"`).
			HTMLContains(`data-size="small"`).
			HasElement("div")
	})
	t.Run("default_size", func(t *testing.T) {
		screen := guitesting.Render(Spinner(SpinnerProps{}))
		screen.Assert(t).HasElement("div").HTMLNotContains("data-size")
	})
	t.Run("large", func(t *testing.T) {
		screen := guitesting.Render(Spinner(SpinnerProps{Size: SpinnerLarge}))
		screen.Assert(t).HTMLContains(`data-size="large"`)
	})
}

func TestIcon(t *testing.T) {
	t.Run("renders_nested_i_elements", func(t *testing.T) {
		screen := guitesting.Render(Icon("ic", "icon-star"))
		screen.Assert(t).
			HTMLContains(`class="ic"`).
			HasElement("i")
		iElems := screen.QueryAllByTag("i")
		if len(iElems) < 2 {
			t.Errorf("expected at least 2 i elements (outer+inner), got %d", len(iElems))
		}
	})
}

func TestAppShellFull(t *testing.T) {
	t.Run("scrollable", func(t *testing.T) {
		screen := guitesting.Render(AppShellFull("asf", true, gui.Text("app")))
		screen.Assert(t).
			HTMLContains(`class="asf"`).
			HTMLContains(`data-scrollable="true"`).
			TextVisible("app")
	})
	t.Run("not_scrollable", func(t *testing.T) {
		screen := guitesting.Render(AppShellFull("", false))
		screen.Assert(t).HasElement("div").HTMLNotContains("data-scrollable")
	})
}
