package coremd

import (
	"strings"
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

// TestSimpleContainers tests components that are simple div/element wrappers
// with optional class and children: AppShell, AppShellBody, ModalBackdrop,
// ModalBody, ModalFooter, DashboardLayout.
func TestSimpleContainers(t *testing.T) {
	type testCase struct {
		name   string
		base   string
		render func(class string, children ...gui.Node) gui.Node
		tag    string
	}
	cases := []testCase{
		{"AppShell", "app", AppShell, "div"},
		{"AppShellBody", "app-shell-body", AppShellBody, "div"},
		{"ModalBackdrop", "modal-backdrop", ModalBackdrop, "div"},
		{"ModalBody", "modal-body", ModalBody, "div"},
		{"ModalFooter", "modal-footer", ModalFooter, "div"},
		{"DashboardLayout", "dashboard-layout", DashboardLayout, "div"},
	}
	for _, tc := range cases {
		t.Run(tc.name+"/with_class_and_children", func(t *testing.T) {
			screen := guitesting.Render(tc.render("my-class", gui.Text("hello")))
			screen.Assert(t).
				HasElement(tc.tag).
				HTMLContains(`class="`+tc.base+` my-class"`).
				TextVisible("hello")
		})
		t.Run(tc.name+"/empty_class", func(t *testing.T) {
			screen := guitesting.Render(tc.render(""))
			screen.Assert(t).HasElement(tc.tag).HTMLContains(`class="`+tc.base+`"`)
		})
	}
}

func TestAppShellMain(t *testing.T) {
	t.Run("renders_main_element_with_class", func(t *testing.T) {
		screen := guitesting.Render(AppShellMain("content", gui.Text("body")))
		screen.Assert(t).
			HasElement("main").
			HTMLContains(`class="app-shell-main content"`).
			TextVisible("body")
	})
	t.Run("no_class", func(t *testing.T) {
		screen := guitesting.Render(AppShellMain(""))
		screen.Assert(t).HasElement("main").HTMLContains(`class="app-shell-main"`)
	})
}

func TestNavbar(t *testing.T) {
	t.Run("renders_nav_with_brand_and_links", func(t *testing.T) {
		link := gui.Span()(gui.Text("Home"))
		screen := guitesting.Render(Navbar(NavbarProps{
			Class: "topnav",
			Brand: "MyApp",
		}, link))
		screen.Assert(t).
			HasElement("nav").
			HTMLContains(`class="navbar topnav"`).
			TextVisible("MyApp").
			TextVisible("Home")
	})
	t.Run("with_stats", func(t *testing.T) {
		screen := guitesting.Render(Navbar(NavbarProps{
			Brand: "App",
			Stats: gui.Text("3 online"),
		}))
		screen.Assert(t).TextVisible("3 online")
	})
	t.Run("without_stats", func(t *testing.T) {
		screen := guitesting.Render(Navbar(NavbarProps{Brand: "App"}))
		html := screen.HTML()
		// Nav should have exactly 2 children (brand span + links div), not 3
		screen.Assert(t).TextVisible("App")
		// Verify no extra div for stats by counting divs
		if strings.Count(html, "<div>") > 1 {
			t.Errorf("expected no stats div, got HTML: %s", html)
		}
	})
}

func TestSidebar(t *testing.T) {
	t.Run("renders_aside_when_open", func(t *testing.T) {
		header := gui.Span()(gui.Text("Header"))
		screen := guitesting.Render(Sidebar(SidebarProps{
			Class: "side",
			Open:  true,
		}, header, gui.Text("content")))
		screen.Assert(t).
			HasElement("aside").
			HTMLContains(`class="sidebar side"`).
			HTMLContains(`data-open="true"`).
			TextVisible("Header").
			TextVisible("content")
	})
	t.Run("closed_has_no_data_open", func(t *testing.T) {
		screen := guitesting.Render(Sidebar(SidebarProps{Open: false}, nil))
		screen.Assert(t).
			HasElement("aside").
			HTMLNotContains("data-open")
	})
	t.Run("nil_header_omitted", func(t *testing.T) {
		screen := guitesting.Render(Sidebar(SidebarProps{}, nil, gui.Text("body")))
		screen.Assert(t).TextVisible("body")
	})
}

func TestSidebarHeader(t *testing.T) {
	t.Run("renders_title_and_actions", func(t *testing.T) {
		action := gui.Span()(gui.Text("X"))
		screen := guitesting.Render(SidebarHeader("hdr", "Files", action))
		screen.Assert(t).
			HTMLContains(`class="sidebar-header hdr"`).
			TextVisible("Files").
			TextVisible("X")
	})
}

func TestPanel(t *testing.T) {
	t.Run("expanded_with_title_and_actions", func(t *testing.T) {
		action := gui.Span()(gui.Text("toggle"))
		screen := guitesting.Render(Panel(PanelProps{
			Class:    "pnl",
			Title:    "Details",
			Expanded: true,
		}, []gui.Node{action}, gui.Text("body")))
		screen.Assert(t).
			HTMLContains(`class="panel pnl"`).
			HTMLContains(`data-expanded="true"`).
			TextVisible("Details").
			TextVisible("toggle").
			TextVisible("body")
	})
	t.Run("collapsed_no_data_expanded", func(t *testing.T) {
		screen := guitesting.Render(Panel(PanelProps{Title: "T"}, nil))
		screen.Assert(t).
			HTMLNotContains("data-expanded").
			TextVisible("T")
	})
	t.Run("no_actions_div_when_empty", func(t *testing.T) {
		screen := guitesting.Render(Panel(PanelProps{Title: "T"}, []gui.Node{}))
		// Header should contain only the title span, no actions div
		screen.Assert(t).TextVisible("T")
	})
}

func TestModal(t *testing.T) {
	t.Run("renders_title_and_children", func(t *testing.T) {
		screen := guitesting.Render(Modal("dlg", "Confirm", gui.Text("Are you sure?")))
		screen.Assert(t).
			HTMLContains(`class="modal dlg"`).
			TextVisible("Confirm").
			TextVisible("Are you sure?")
	})
	t.Run("empty_class", func(t *testing.T) {
		screen := guitesting.Render(Modal("", "Title"))
		screen.Assert(t).HTMLContains(`class="modal"`).TextVisible("Title")
	})
}

func TestSidebarColumn(t *testing.T) {
	t.Run("open", func(t *testing.T) {
		screen := guitesting.Render(SidebarColumn("col", true, gui.Text("nav")))
		screen.Assert(t).
			HTMLContains(`class="sidebar-col col"`).
			HTMLContains(`data-open="true"`).
			TextVisible("nav")
	})
	t.Run("closed", func(t *testing.T) {
		screen := guitesting.Render(SidebarColumn("col", false))
		screen.Assert(t).HTMLNotContains("data-open")
	})
}

func TestChatHeader(t *testing.T) {
	t.Run("with_title_and_toolbar", func(t *testing.T) {
		title := gui.Span()(gui.Text("Chat"))
		toolbar := gui.Span()(gui.Text("Settings"))
		screen := guitesting.Render(ChatHeader("hdr", title, toolbar))
		screen.Assert(t).
			HTMLContains(`class="chat-header hdr"`).
			TextVisible("Chat").
			TextVisible("Settings")
	})
	t.Run("nil_title_and_toolbar", func(t *testing.T) {
		screen := guitesting.Render(ChatHeader("hdr", nil, nil))
		screen.Assert(t).HasElement("div").HTMLContains(`class="chat-header hdr"`)
	})
	t.Run("only_title", func(t *testing.T) {
		title := gui.Span()(gui.Text("Room"))
		screen := guitesting.Render(ChatHeader("", title, nil))
		screen.Assert(t).HTMLContains(`class="chat-header"`).TextVisible("Room")
	})
}

func TestBox(t *testing.T) {
	t.Run("all_props", func(t *testing.T) {
		s := guitesting.Render(Box(BoxProps{
			Class:   "my-box",
			Pad:     "lg",
			Bg:      "surface",
			Border:  true,
			Flex:    true,
			Rounded: true,
		}, gui.Text("inside")))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`class="box my-box"`)
		a.HTMLContains(`data-box`)
		a.HTMLContains(`data-pad="lg"`)
		a.HTMLContains(`data-bg="surface"`)
		a.HTMLContains(`data-box-border="true"`)
		a.HTMLContains(`data-box-flex="true"`)
		a.HTMLContains(`data-box-rounded="true"`)
		a.TextVisible("inside")
	})
	t.Run("minimal_props", func(t *testing.T) {
		s := guitesting.Render(Box(BoxProps{}))
		a := s.Assert(t)
		a.HTMLContains(`data-box`)
		a.HTMLNotContains("data-pad")
		a.HTMLNotContains("data-bg")
		a.HTMLNotContains("data-box-border")
		a.HTMLNotContains("data-box-flex")
		a.HTMLNotContains("data-box-rounded")
		a.HTMLContains(`class="box"`)
	})
}

func TestScrollArea(t *testing.T) {
	t.Run("with_class_and_children", func(t *testing.T) {
		s := guitesting.Render(ScrollArea("scroll", gui.Text("scrollable")))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`class="scroll-area scroll"`)
		a.HTMLContains(`data-scroll-area`)
		a.TextVisible("scrollable")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(ScrollArea(""))
		s.Assert(t).HTMLContains(`class="scroll-area"`).HTMLContains(`data-scroll-area`)
	})
}

func TestSplitLayout(t *testing.T) {
	t.Run("three_columns", func(t *testing.T) {
		s := guitesting.Render(SplitLayout(SplitLayoutProps{
			Class:   "split",
			Sidebar: "260px",
			Panel:   "320px",
		},
			gui.Text("sidebar"),
			gui.Text("center"),
			gui.Text("panel"),
		))
		a := s.Assert(t)
		a.HTMLContains(`class="split-layout split"`)
		a.HTMLContains(`data-split-layout`)
		a.TextVisible("sidebar")
		a.TextVisible("center")
		a.TextVisible("panel")
		html := s.HTML()
		if !strings.Contains(html, "width:260px") {
			t.Errorf("expected sidebar width style, got: %s", html)
		}
		if !strings.Contains(html, "width:320px") {
			t.Errorf("expected panel width style, got: %s", html)
		}
	})
	t.Run("nil_sidebar_and_panel", func(t *testing.T) {
		s := guitesting.Render(SplitLayout(SplitLayoutProps{},
			nil,
			gui.Text("center only"),
			nil,
		))
		a := s.Assert(t)
		a.HTMLContains(`data-split-layout`)
		a.TextVisible("center only")
	})
}

func TestBackdrop(t *testing.T) {
	t.Run("with_onclick", func(t *testing.T) {
		clicked := false
		s := guitesting.Render(Backdrop("overlay", func() { clicked = true }))
		s.Assert(t).HTMLContains(`class="backdrop overlay"`).HTMLContains(`data-backdrop`)
		ref := s.QueryAllByTag("div")[0]
		s.Click(ref)
		if !clicked {
			t.Error("expected onclick to fire")
		}
	})
	t.Run("nil_onclick", func(t *testing.T) {
		s := guitesting.Render(Backdrop("", nil))
		a := s.Assert(t)
		a.HTMLContains(`data-backdrop`)
		a.HTMLContains(`class="backdrop"`)
		html := s.HTML()
		if strings.Contains(html, "onclick") {
			t.Errorf("expected no onclick, got: %s", html)
		}
	})
}

func TestIconButton(t *testing.T) {
	t.Run("with_all_props", func(t *testing.T) {
		clicked := false
		s := guitesting.Render(IconButton("ib-cls", "icon-close", "Close", func() { clicked = true }))
		a := s.Assert(t)
		a.HasElement("button")
		a.HTMLContains(`class="icon-button ib-cls"`)
		a.HTMLContains(`data-icon-button`)
		a.HTMLContains(`aria-label="Close"`)
		a.HTMLContains(`class="icon-close"`)
		ref := s.QueryAllByTag("button")[0]
		s.Click(ref)
		if !clicked {
			t.Error("expected onclick to fire")
		}
	})
	t.Run("nil_onclick_empty_label", func(t *testing.T) {
		s := guitesting.Render(IconButton("", "icon-menu", "", nil))
		a := s.Assert(t)
		a.HasElement("button")
		a.HTMLContains(`data-icon-button`)
		a.HTMLNotContains("aria-label")
		html := s.HTML()
		if strings.Contains(html, "onclick") {
			t.Errorf("expected no onclick, got: %s", html)
		}
	})
}

func TestToolbar(t *testing.T) {
	t.Run("with_class_and_children", func(t *testing.T) {
		s := guitesting.Render(Toolbar("tb", gui.Text("btn1"), gui.Text("btn2")))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`class="toolbar tb"`)
		a.HTMLContains(`data-toolbar`)
		a.TextVisible("btn1")
		a.TextVisible("btn2")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(Toolbar(""))
		s.Assert(t).HTMLContains(`class="toolbar"`).HTMLContains(`data-toolbar`)
	})
}

func TestToolbarButton(t *testing.T) {
	t.Run("with_icon_and_click", func(t *testing.T) {
		clicked := false
		s := guitesting.Render(ToolbarButton("tb-cls", "icon-terminal", "Terminal", func() { clicked = true }))
		a := s.Assert(t)
		a.HasElement("button")
		a.HTMLContains(`class="toolbar-button tb-cls"`)
		a.HTMLContains(`data-toolbar-button`)
		a.HTMLContains(`class="icon-terminal"`)
		a.TextVisible("Terminal")

		btn := s.GetByRole("button")
		s.Click(btn)
		if !clicked {
			t.Error("expected onClick to fire")
		}
	})

	t.Run("no_icon_no_click", func(t *testing.T) {
		s := guitesting.Render(ToolbarButton("", "", "Run", nil))
		a := s.Assert(t)
		a.HasElement("button")
		a.HTMLContains(`class="toolbar-button"`)
		a.HTMLContains(`data-toolbar-button`)
		a.TextVisible("Run")
		a.HasNoElement("i")
		html := s.HTML()
		if strings.Contains(html, "onclick") {
			t.Errorf("expected no onclick, got: %s", html)
		}
	})
}

func TestResizeHandle(t *testing.T) {
	t.Run("vertical", func(t *testing.T) {
		s := guitesting.Render(ResizeHandle("rh", ResizeVertical))
		a := s.Assert(t)
		a.HasElement("div")
		a.HTMLContains(`class="resize-handle rh"`)
		a.HTMLContains(`data-resize-handle`)
		a.HTMLContains(`data-direction="vertical"`)
	})

	t.Run("horizontal", func(t *testing.T) {
		s := guitesting.Render(ResizeHandle("", ResizeHorizontal))
		a := s.Assert(t)
		a.HTMLContains(`class="resize-handle"`)
		a.HTMLContains(`data-direction="horizontal"`)
	})

	t.Run("default_direction", func(t *testing.T) {
		s := guitesting.Render(ResizeHandle("", ""))
		a := s.Assert(t)
		a.HTMLContains(`data-direction="vertical"`)
	})
}
