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
		a.HTMLContains(`class="markdown-content prose"`)
		a.HTMLContains(`data-rich-text`)
		a.TextVisible("Hello world")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(MarkdownContent(""))
		s.Assert(t).HasElement("div").HTMLContains(`class="markdown-content"`).HTMLContains(`data-rich-text`)
	})
}

func TestSectionHeader(t *testing.T) {
	t.Run("with_title_and_actions", func(t *testing.T) {
		action := gui.Span()(gui.Text("Edit"))
		s := guitesting.Render(SectionHeader("sec-hdr", "Settings", action))
		a := s.Assert(t)
		a.HTMLContains(`class="section-header sec-hdr"`)
		a.HTMLContains(`data-section-header`)
		a.TextVisible("Settings")
		a.TextVisible("Edit")
	})
	t.Run("title_only_no_actions", func(t *testing.T) {
		s := guitesting.Render(SectionHeader("", "Title"))
		a := s.Assert(t)
		a.HTMLContains(`data-section-header`)
		a.TextVisible("Title")
		a.HTMLContains(`class="section-header"`)
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
		a.HTMLContains(`class="collapsible collapse"`)
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
		a.HTMLContains(`class="animate anim-cls"`)
		a.HTMLContains(`data-animate="fade-in"`)
		a.TextVisible("content")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(Animate("", "spin"))
		s.Assert(t).HTMLContains(`class="animate"`).HTMLContains(`data-animate="spin"`)
	})
}

func TestStack(t *testing.T) {
	t.Run("default_gap", func(t *testing.T) {
		s := guitesting.Render(Stack("", gui.Text("child")))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`data-stack="md"`)
		a.TextVisible("child")
	})
	t.Run("custom_gap", func(t *testing.T) {
		s := guitesting.Render(Stack("lg", gui.Text("a"), gui.Text("b")))
		a := s.Assert(t)
		a.HTMLContains(`data-stack="lg"`)
		a.TextVisible("a")
		a.TextVisible("b")
	})
}

func TestHStack(t *testing.T) {
	t.Run("default_gap", func(t *testing.T) {
		s := guitesting.Render(HStack("", gui.Text("child")))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`data-hstack="md"`)
		a.TextVisible("child")
	})
	t.Run("custom_gap", func(t *testing.T) {
		s := guitesting.Render(HStack("sm", gui.Text("left"), gui.Text("right")))
		a := s.Assert(t)
		a.HTMLContains(`data-hstack="sm"`)
	})
}

func TestGrid(t *testing.T) {
	t.Run("defaults", func(t *testing.T) {
		s := guitesting.Render(Grid(GridProps{}, gui.Text("cell")))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`class="grid"`)
		a.HTMLContains(`data-grid="2"`)
		a.TextVisible("cell")
	})
	t.Run("with_all_props", func(t *testing.T) {
		s := guitesting.Render(Grid(GridProps{
			Class:   "my-grid",
			Cols:    "3",
			Align:   "center",
			Justify: "space-between",
		}, gui.Text("cell")))
		a := s.Assert(t)
		a.HTMLContains(`class="grid my-grid"`)
		a.HTMLContains(`data-grid="3"`)
		a.HTMLContains(`data-align="center"`)
		a.HTMLContains(`data-justify="space-between"`)
	})
}

func TestCenter(t *testing.T) {
	t.Run("with_class", func(t *testing.T) {
		s := guitesting.Render(Center("centered", gui.Text("hello")))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`class="center centered"`)
		a.HTMLContains(`data-center`)
		a.TextVisible("hello")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(Center("", gui.Text("x")))
		a := s.Assert(t)
		a.HTMLContains(`class="center"`)
		a.HTMLContains(`data-center`)
	})
}

func TestSpacer(t *testing.T) {
	s := guitesting.Render(Spacer())
	a := s.Assert(t)
	a.HasElement("div")
	a.HTMLContains(`data-spacer`)
}

func TestCard(t *testing.T) {
	t.Run("default_variant", func(t *testing.T) {
		s := guitesting.Render(Card(CardProps{}, gui.Text("content")))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`class="card"`)
		a.HTMLContains(`data-card="true"`)
		a.TextVisible("content")
	})
	t.Run("with_variant_and_class", func(t *testing.T) {
		s := guitesting.Render(Card(CardProps{Class: "my-card", Variant: "surface"}, gui.Text("body")))
		a := s.Assert(t)
		a.HTMLContains(`class="card my-card"`)
		a.HTMLContains(`data-card="surface"`)
	})
}

func TestBadge(t *testing.T) {
	t.Run("default_variant", func(t *testing.T) {
		s := guitesting.Render(Badge("", BadgeDefault, "New"))
		a := s.Assert(t)
		a.HasElement("span")
		a.HTMLContains(`class="badge"`)
		a.HTMLContains(`data-badge="true"`)
		a.TextVisible("New")
	})
	t.Run("accent_variant_with_class", func(t *testing.T) {
		s := guitesting.Render(Badge("my-badge", BadgeAccent, "Hot"))
		a := s.Assert(t)
		a.HTMLContains(`class="badge my-badge"`)
		a.HTMLContains(`data-badge="accent"`)
		a.TextVisible("Hot")
	})
	t.Run("danger_variant", func(t *testing.T) {
		s := guitesting.Render(Badge("", BadgeDanger, "Error"))
		a := s.Assert(t)
		a.HTMLContains(`data-badge="danger"`)
		a.TextVisible("Error")
	})
}

func TestDivider(t *testing.T) {
	t.Run("with_class", func(t *testing.T) {
		s := guitesting.Render(Divider("sep"))
		a := s.Assert(t)
		a.HasElement("hr")
		a.HTMLContains(`class="divider sep"`)
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(Divider(""))
		a := s.Assert(t)
		a.HasElement("hr")
		a.HTMLContains(`class="divider"`)
	})
}

func TestHeading(t *testing.T) {
	t.Run("h1", func(t *testing.T) {
		s := guitesting.Render(Heading(1, "", gui.Text("Title")))
		a := s.Assert(t)
		a.HasElement("h1")
		a.HTMLContains(`class="heading"`)
		a.TextVisible("Title")
	})
	t.Run("h2_default", func(t *testing.T) {
		s := guitesting.Render(Heading(2, "big", gui.Text("Sub")))
		a := s.Assert(t)
		a.HasElement("h2")
		a.HTMLContains(`class="heading big"`)
	})
	t.Run("h3", func(t *testing.T) {
		s := guitesting.Render(Heading(3, "", gui.Text("H3")))
		s.Assert(t).HasElement("h3")
	})
	t.Run("h4", func(t *testing.T) {
		s := guitesting.Render(Heading(4, "", gui.Text("H4")))
		s.Assert(t).HasElement("h4")
	})
	t.Run("h5", func(t *testing.T) {
		s := guitesting.Render(Heading(5, "", gui.Text("H5")))
		s.Assert(t).HasElement("h5")
	})
	t.Run("h6", func(t *testing.T) {
		s := guitesting.Render(Heading(6, "", gui.Text("H6")))
		s.Assert(t).HasElement("h6")
	})
	t.Run("invalid_defaults_to_h2", func(t *testing.T) {
		s := guitesting.Render(Heading(99, "", gui.Text("Default")))
		s.Assert(t).HasElement("h2")
	})
}

func TestParagraph(t *testing.T) {
	t.Run("with_class", func(t *testing.T) {
		s := guitesting.Render(Paragraph("intro", gui.Text("Hello")))
		a := s.Assert(t)
		a.HasElement("p")
		a.HTMLContains(`class="paragraph intro"`)
		a.TextVisible("Hello")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(Paragraph("", gui.Text("text")))
		s.Assert(t).HTMLContains(`class="paragraph"`)
	})
}

func TestCodeBlock(t *testing.T) {
	t.Run("with_content", func(t *testing.T) {
		s := guitesting.Render(CodeBlock("go", "fmt.Println()"))
		a := s.Assert(t)
		a.HasElement("pre")
		a.HasElement("code")
		a.HTMLContains(`class="code-block go"`)
		a.TextVisible("fmt.Println()")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(CodeBlock("", "x := 1"))
		s.Assert(t).HTMLContains(`class="code-block"`)
	})
}

func TestInlineCode(t *testing.T) {
	s := guitesting.Render(InlineCode("var x"))
	a := s.Assert(t)
	a.HasElement("code")
	a.TextVisible("var x")
}

func TestMuted(t *testing.T) {
	s := guitesting.Render(Muted("secondary text"))
	a := s.Assert(t)
	a.HasElement("span")
	a.HTMLContains(`data-muted`)
	a.TextVisible("secondary text")
}

func TestMono(t *testing.T) {
	s := guitesting.Render(Mono("monospace text"))
	a := s.Assert(t)
	a.HasElement("span")
	a.HTMLContains(`data-mono`)
	a.TextVisible("monospace text")
}

func TestLink(t *testing.T) {
	t.Run("with_all_props", func(t *testing.T) {
		s := guitesting.Render(Link(LinkProps{
			Class:  "nav",
			Href:   "https://example.com",
			Target: "_blank",
		}, gui.Text("Click")))
		a := s.Assert(t)
		a.HasElement("a")
		a.HTMLContains(`class="link nav"`)
		a.HTMLContains(`href="https://example.com"`)
		a.HTMLContains(`target="_blank"`)
		a.HTMLContains(`rel="noopener noreferrer"`)
		a.TextVisible("Click")
	})
	t.Run("self_target_no_rel", func(t *testing.T) {
		s := guitesting.Render(Link(LinkProps{
			Href:   "/about",
			Target: "_self",
		}, gui.Text("About")))
		a := s.Assert(t)
		a.HTMLContains(`target="_self"`)
		a.HTMLNotContains("noopener")
	})
	t.Run("minimal", func(t *testing.T) {
		s := guitesting.Render(Link(LinkProps{}, gui.Text("bare")))
		a := s.Assert(t)
		a.HasElement("a")
		a.HTMLContains(`class="link"`)
		a.HTMLNotContains("href=")
		a.HTMLNotContains("target=")
	})
}

func TestImage(t *testing.T) {
	t.Run("with_all_props", func(t *testing.T) {
		s := guitesting.Render(Image(ImageProps{
			Class:   "pic",
			Src:     "photo.jpg",
			Alt:     "A photo",
			Rounded: true,
			Avatar:  true,
		}))
		a := s.Assert(t)
		a.HasElement("img")
		a.HTMLContains(`class="image pic"`)
		a.HTMLContains(`src="photo.jpg"`)
		a.HTMLContains(`alt="A photo"`)
		a.HTMLContains(`data-rounded`)
		a.HTMLContains(`data-avatar`)
	})
	t.Run("minimal", func(t *testing.T) {
		s := guitesting.Render(Image(ImageProps{}))
		a := s.Assert(t)
		a.HasElement("img")
		a.HTMLContains(`class="image"`)
		a.HTMLNotContains("src=")
		a.HTMLNotContains("alt=")
		a.HTMLNotContains("data-rounded")
		a.HTMLNotContains("data-avatar")
	})
}

func TestUnorderedList(t *testing.T) {
	s := guitesting.Render(UnorderedList("items",
		ListItem(gui.Text("one")),
		ListItem(gui.Text("two")),
	))
	a := s.Assert(t)
	a.HasElement("ul")
	a.HTMLContains(`class="unordered-list items"`)
	a.TextVisible("one")
	a.TextVisible("two")
}

func TestOrderedList(t *testing.T) {
	s := guitesting.Render(OrderedList("steps",
		ListItem(gui.Text("first")),
		ListItem(gui.Text("second")),
	))
	a := s.Assert(t)
	a.HasElement("ol")
	a.HTMLContains(`class="ordered-list steps"`)
	a.TextVisible("first")
	a.TextVisible("second")
}

func TestListItem(t *testing.T) {
	s := guitesting.Render(ListItem(gui.Text("item")))
	a := s.Assert(t)
	a.HasElement("li")
	a.TextVisible("item")
}

func TestQuote(t *testing.T) {
	t.Run("with_class", func(t *testing.T) {
		s := guitesting.Render(Quote("highlight", gui.Text("quoted text")))
		a := s.Assert(t)
		a.HasElement("blockquote")
		a.HTMLContains(`class="quote highlight"`)
		a.TextVisible("quoted text")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(Quote("", gui.Text("plain")))
		s.Assert(t).HTMLContains(`class="quote"`)
	})
}

func TestTruncate(t *testing.T) {
	t.Run("with_class", func(t *testing.T) {
		s := guitesting.Render(Truncate("trunc", gui.Text("long text")))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`class="truncate trunc"`)
		a.HTMLContains(`data-truncate`)
		a.TextVisible("long text")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(Truncate("", gui.Text("x")))
		s.Assert(t).HTMLContains(`class="truncate"`)
	})
}

func TestSrOnly(t *testing.T) {
	s := guitesting.Render(SrOnly("screen reader text"))
	a := s.Assert(t)
	a.HasElement("span")
	a.HTMLContains(`data-sr-only`)
	a.TextVisible("screen reader text")
}

func TestHelpText(t *testing.T) {
	t.Run("with_class_and_children", func(t *testing.T) {
		s := guitesting.Render(HelpText("ht", gui.Text("Create a token at "), gui.A()(gui.Text("github.com"))))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`class="help-text ht"`)
		a.HTMLContains(`data-help-text`)
		a.TextVisible("Create a token at")
		a.TextVisible("github.com")
	})

	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(HelpText("", gui.Text("info")))
		a := s.Assert(t)
		a.HTMLContains(`class="help-text"`)
		a.HTMLContains(`data-help-text`)
		a.TextVisible("info")
	})
}
