package coremd

import (
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestButton_BaseClass(t *testing.T) {
	t.Run("with_custom_class", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{Class: "my-btn"}, gui.Text("Click")))
		screen.Assert(t).
			HasElement("button").
			HTMLContains(`class="btn my-btn"`).
			TextVisible("Click")
	})
	t.Run("empty_class_renders_base", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{}, gui.Text("OK")))
		screen.Assert(t).
			HasElement("button").
			HTMLContains(`class="btn"`).
			TextVisible("OK")
	})
}

func TestButton_Variants(t *testing.T) {
	t.Run("primary", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{Variant: ButtonPrimary}, gui.Text("Go")))
		screen.Assert(t).
			HTMLContains(`class="btn"`).
			HTMLContains(`data-variant="primary"`)
	})
	t.Run("danger", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{Variant: ButtonDanger}, gui.Text("Del")))
		screen.Assert(t).HTMLContains(`data-variant="danger"`)
	})
	t.Run("toolbar", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{Variant: ButtonToolbar}, gui.Text("T")))
		screen.Assert(t).HTMLContains(`data-variant="toolbar"`)
	})
	t.Run("ghost", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{Variant: ButtonGhost}))
		screen.Assert(t).HTMLContains(`data-variant="ghost"`)
	})
	t.Run("default_no_variant_attr", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{}, gui.Text("X")))
		screen.Assert(t).HTMLNotContains("data-variant")
	})
}

func TestButton_Icon(t *testing.T) {
	t.Run("icon_prepended_to_children", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{Icon: "icon-star"}, gui.Text("Star")))
		screen.Assert(t).
			HasElement("i").
			HTMLContains(`class="icon-star"`).
			TextVisible("Star")
	})
	t.Run("icon_only", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{
			Variant: ButtonGhost,
			Icon:    "icon-close",
		}))
		screen.Assert(t).
			HasElement("i").
			HTMLContains(`class="icon-close"`).
			HTMLContains(`data-variant="ghost"`)
	})
	t.Run("no_icon", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{}, gui.Text("Plain")))
		screen.Assert(t).HasNoElement("i")
	})
}

func TestButton_AriaLabel(t *testing.T) {
	t.Run("with_aria_label", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{AriaLabel: "Close"}))
		screen.Assert(t).HTMLContains(`aria-label="Close"`)
	})
	t.Run("without_aria_label", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{}))
		screen.Assert(t).HTMLNotContains("aria-label")
	})
}

func TestButton_Size(t *testing.T) {
	t.Run("small", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{Size: ButtonSmall}, gui.Text("S")))
		screen.Assert(t).HTMLContains(`data-size="small"`)
	})
	t.Run("default_no_size_attr", func(t *testing.T) {
		screen := guitesting.Render(Button(ButtonProps{}, gui.Text("M")))
		screen.Assert(t).HTMLNotContains("data-size")
	})
}

func TestButton_Disabled(t *testing.T) {
	screen := guitesting.Render(Button(ButtonProps{Disabled: true}, gui.Text("No")))
	screen.Assert(t).HTMLContains("disabled")
}

func TestButton_OnClick(t *testing.T) {
	clicked := false
	screen := guitesting.Render(Button(ButtonProps{OnClick: func() { clicked = true }}, gui.Text("Go")))
	btn := screen.GetByRole("button")
	screen.Click(btn)
	if !clicked {
		t.Error("expected onClick to fire")
	}
}
