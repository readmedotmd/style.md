package coremd

import (
	"strings"
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestSearchOverlay(t *testing.T) {
	t.Run("with_tabs_input_results", func(t *testing.T) {
		tabs := []TabBarTab{{Label: "Files", Active: true}, {Label: "Code"}}
		input := gui.Span()(gui.Text("search-input"))
		screen := guitesting.Render(SearchOverlay("so", tabs, input, gui.Text("result1")))
		screen.Assert(t).
			HTMLContains(`class="so"`).
			TextVisible("Files").
			TextVisible("search-input").
			TextVisible("result1")
	})
	t.Run("no_tabs_no_input", func(t *testing.T) {
		screen := guitesting.Render(SearchOverlay("", nil, nil, gui.Text("r")))
		screen.Assert(t).HasElement("div").TextVisible("r")
	})
}

func TestContextMenu(t *testing.T) {
	t.Run("renders_items_with_position", func(t *testing.T) {
		items := []ContextMenuItem{
			{Label: "Copy", Danger: false},
			{Label: "Delete", Danger: true},
		}
		screen := guitesting.Render(ContextMenu("ctx", 100, 200, items))
		screen.Assert(t).
			HTMLContains(`class="ctx"`).
			HTMLContains("left: 100px").
			HTMLContains("top: 200px").
			TextVisible("Copy").
			TextVisible("Delete").
			HTMLContains(`data-danger="true"`)
	})
	t.Run("onclick_fires", func(t *testing.T) {
		clicked := false
		items := []ContextMenuItem{{Label: "Act", OnClick: func() { clicked = true }}}
		screen := guitesting.Render(ContextMenu("", 0, 0, items))
		ref := screen.QueryAllByTag("button")[0]
		screen.Click(ref)
		if !clicked {
			t.Error("expected onclick to fire")
		}
	})
}

func TestBottomSheet(t *testing.T) {
	t.Run("renders_handle_and_items", func(t *testing.T) {
		items := []BottomSheetItem{
			{Icon: "icon-share", Label: "Share"},
			{Label: "Delete", Danger: true},
		}
		screen := guitesting.Render(BottomSheet("bs", items))
		screen.Assert(t).
			HTMLContains(`class="bs"`).
			TextVisible("Share").
			TextVisible("Delete").
			HTMLContains(`data-danger="true"`)
		// handle div + items div = at least 2 child divs
		divs := screen.QueryAllByTag("div")
		if len(divs) < 2 {
			t.Errorf("expected at least 2 divs, got %d", len(divs))
		}
	})
}

func TestSearchOverlayCard(t *testing.T) {
	t.Run("with_all_children", func(t *testing.T) {
		screen := guitesting.Render(SearchOverlayCard("soc",
			gui.Text("tabs"), gui.Text("input"), gui.Text("results")))
		screen.Assert(t).
			HTMLContains(`class="soc"`).
			TextVisible("tabs").
			TextVisible("input").
			TextVisible("results")
	})
	t.Run("nil_children_omitted", func(t *testing.T) {
		screen := guitesting.Render(SearchOverlayCard("", nil, nil, nil))
		html := screen.HTML()
		// Should just be an empty div
		if !strings.Contains(html, "<div>") {
			t.Errorf("expected empty div, got: %s", html)
		}
	})
}

func TestSearchResult(t *testing.T) {
	t.Run("with_icon_and_add_button", func(t *testing.T) {
		clicked := false
		screen := guitesting.Render(SearchResult("sr", "icon-file", "/src/main.go", "func main", func() { clicked = true }))
		screen.Assert(t).
			HTMLContains(`class="sr"`).
			TextVisible("/src/main.go").
			TextVisible("func main").
			TextVisible("+")
		ref := screen.QueryAllByTag("button")[0]
		screen.Click(ref)
		if !clicked {
			t.Error("expected onAdd to fire")
		}
	})
	t.Run("no_icon_no_add", func(t *testing.T) {
		screen := guitesting.Render(SearchResult("", "", "path", "text", nil))
		screen.Assert(t).TextVisible("path").TextVisible("text")
		html := screen.HTML()
		if strings.Contains(html, "<button") {
			t.Errorf("expected no button, got: %s", html)
		}
	})
}

func TestSearchResultContent(t *testing.T) {
	t.Run("with_snippet_and_add", func(t *testing.T) {
		snippet := gui.Span()(gui.Text("code"))
		screen := guitesting.Render(SearchResultContent("src", "/file.go", snippet, func() {}))
		screen.Assert(t).
			HTMLContains(`class="src"`).
			TextVisible("/file.go").
			TextVisible("code").
			TextVisible("+")
	})
	t.Run("nil_snippet_no_add", func(t *testing.T) {
		screen := guitesting.Render(SearchResultContent("", "path", nil, nil))
		screen.Assert(t).TextVisible("path")
	})
}

func TestSearchSnippet(t *testing.T) {
	t.Run("renders_lines_with_match", func(t *testing.T) {
		lines := []SearchSnippetLine{
			{Text: "line1", IsMatch: false},
			{Text: "line2", IsMatch: true},
			{Text: "line3", IsMatch: false},
		}
		screen := guitesting.Render(SearchSnippet("snip", lines))
		screen.Assert(t).
			HTMLContains(`class="snip"`).
			TextVisible("line1").
			TextVisible("line2").
			HTMLContains(`data-match="true"`)
		// Only one line should have data-match
		html := screen.HTML()
		if strings.Count(html, `data-match="true"`) != 1 {
			t.Errorf("expected exactly 1 data-match, got HTML: %s", html)
		}
	})
}
