package coremd

import (
	"strings"
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestNavLink(t *testing.T) {
	t.Run("active_with_icon", func(t *testing.T) {
		clicked := false
		screen := guitesting.Render(NavLink(NavLinkProps{
			Class:   "nl",
			Icon:    "icon-home",
			Label:   "Home",
			Active:  true,
			OnClick: func() { clicked = true },
		}))
		screen.Assert(t).
			HTMLContains(`class="nav-link nl"`).
			HTMLContains(`data-active="true"`).
			HasElement("button").
			TextVisible("Home")
		ref := screen.QueryAllByTag("button")[0]
		screen.Click(ref)
		if !clicked {
			t.Error("expected onclick to fire")
		}
	})
	t.Run("inactive_no_icon", func(t *testing.T) {
		screen := guitesting.Render(NavLink(NavLinkProps{Label: "About"}))
		screen.Assert(t).
			HTMLNotContains("data-active").
			TextVisible("About")
	})
}

func TestTabBar(t *testing.T) {
	t.Run("renders_tabs_with_active", func(t *testing.T) {
		tabs := []TabBarTab{
			{Label: "Tab1", Active: true},
			{Label: "Tab2"},
		}
		screen := guitesting.Render(TabBar("tb", tabs))
		screen.Assert(t).
			HTMLContains(`class="tab-bar tb"`).
			HTMLContains(`data-active="true"`).
			TextVisible("Tab1").
			TextVisible("Tab2")
		// Only one tab should be active
		html := screen.HTML()
		if strings.Count(html, `data-active="true"`) != 1 {
			t.Errorf("expected exactly 1 active tab, got HTML: %s", html)
		}
	})
	t.Run("tab_buttons_have_tab_bar_item_class", func(t *testing.T) {
		tabs := []TabBarTab{{Label: "A"}, {Label: "B"}}
		screen := guitesting.Render(TabBar("", tabs))
		screen.Assert(t).HTMLContains(`class="tab-bar-item"`)
		html := screen.HTML()
		if strings.Count(html, "tab-bar-item") != 2 {
			t.Errorf("expected 2 tab-bar-item classes, got HTML: %s", html)
		}
	})
	t.Run("empty_class_renders_base", func(t *testing.T) {
		tabs := []TabBarTab{{Label: "X"}}
		screen := guitesting.Render(TabBar("", tabs))
		screen.Assert(t).HTMLContains(`class="tab-bar"`)
	})
}

func TestNavbar_NavLinksWrapper(t *testing.T) {
	t.Run("links_wrapped_in_nav_links_div", func(t *testing.T) {
		link := gui.Span()(gui.Text("Home"))
		screen := guitesting.Render(Navbar(NavbarProps{Brand: "App"}, link))
		screen.Assert(t).
			HTMLContains(`class="nav-links"`).
			TextVisible("Home")
	})
}

func TestBottomTabBar(t *testing.T) {
	t.Run("renders_items_with_active", func(t *testing.T) {
		items := []BottomTabItem{
			{Icon: "icon-chat", Label: "Chat", Active: true},
			{Icon: "icon-settings", Label: "Settings"},
		}
		screen := guitesting.Render(BottomTabBar("btb", items))
		screen.Assert(t).
			HTMLContains(`class="bottom-tab-bar btb"`).
			HTMLContains(`data-active="true"`).
			TextVisible("Chat").
			TextVisible("Settings")
	})
}

func TestChatBackButton(t *testing.T) {
	t.Run("renders_back_arrow", func(t *testing.T) {
		clicked := false
		screen := guitesting.Render(ChatBackButton("cbb", func() { clicked = true }))
		screen.Assert(t).
			HTMLContains(`class="chat-back-btn cbb"`).
			HasElement("button").
			TextVisible("\u2190")
		ref := screen.QueryAllByTag("button")[0]
		screen.Click(ref)
		if !clicked {
			t.Error("expected onclick to fire")
		}
	})
	t.Run("nil_onclick", func(t *testing.T) {
		screen := guitesting.Render(ChatBackButton("", nil))
		screen.Assert(t).HasElement("button")
	})
}

func TestHamburgerButton(t *testing.T) {
	t.Run("renders_hamburger", func(t *testing.T) {
		screen := guitesting.Render(HamburgerButton("hb", func() {}))
		screen.Assert(t).
			HTMLContains(`class="hb"`).
			HasElement("button").
			TextVisible("\u2630")
	})
}

func TestChatToolbar(t *testing.T) {
	t.Run("with_desktop_and_mobile", func(t *testing.T) {
		desktop := gui.Text("desktop-tools")
		mobile := gui.Text("mobile-trigger")
		screen := guitesting.Render(ChatToolbar("ct", desktop, mobile))
		screen.Assert(t).
			HTMLContains(`class="chat-toolbar ct"`).
			TextVisible("desktop-tools").
			TextVisible("mobile-trigger")
	})
	t.Run("nil_children", func(t *testing.T) {
		screen := guitesting.Render(ChatToolbar("", nil, nil))
		html := screen.HTML()
		// Should render outer div with no child divs
		if strings.Count(html, "<div>") > 1 {
			t.Errorf("expected no child divs for nil nodes, got: %s", html)
		}
	})
}

func TestToolbarButton(t *testing.T) {
	t.Run("danger_with_icon_and_label", func(t *testing.T) {
		screen := guitesting.Render(ToolbarButton("tbb", "icon-trash", "Delete", true, func() {}))
		screen.Assert(t).
			HTMLContains(`class="chat-toolbar-btn chat-toolbar-btn-danger tbb"`).
			HTMLContains(`data-danger="true"`).
			HasElement("button").
			TextVisible("Delete")
		iElems := screen.QueryAllByTag("i")
		if len(iElems) == 0 {
			t.Error("expected icon element")
		}
	})
	t.Run("not_danger_no_icon", func(t *testing.T) {
		screen := guitesting.Render(ToolbarButton("", "", "Save", false, nil))
		screen.Assert(t).
			HTMLContains(`class="chat-toolbar-btn"`).
			HTMLNotContains("data-danger").
			HTMLNotContains("chat-toolbar-btn-danger").
			TextVisible("Save")
	})
}
