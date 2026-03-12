package coremd

import (
	"strings"
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestLoginPage(t *testing.T) {
	t.Run("with_form_and_error", func(t *testing.T) {
		form := gui.Span()(gui.Text("form"))
		screen := guitesting.Render(LoginPage("lp", "Sign In", form, "bad creds"))
		screen.Assert(t).
			HTMLContains(`class="lp"`).
			TextVisible("Sign In").
			TextVisible("form").
			TextVisible("bad creds")
	})
	t.Run("no_error_no_form", func(t *testing.T) {
		screen := guitesting.Render(LoginPage("", "Login", nil, ""))
		screen.Assert(t).TextVisible("Login")
		html := screen.HTML()
		if strings.Contains(html, "bad creds") {
			t.Errorf("expected no error, got: %s", html)
		}
	})
}

func TestSetupWizard(t *testing.T) {
	t.Run("renders_steps_and_content", func(t *testing.T) {
		steps := []SetupStep{
			{Label: "Account", Active: true},
			{Label: "Config", Completed: true},
			{Label: "Done"},
		}
		screen := guitesting.Render(SetupWizard("sw", steps, gui.Text("form")))
		screen.Assert(t).
			HTMLContains(`class="sw"`).
			HTMLContains(`data-active="true"`).
			HTMLContains(`data-completed="true"`).
			TextVisible("Account").
			TextVisible("Config").
			TextVisible("form")
		// Completed step shows checkmark
		html := screen.HTML()
		if !strings.Contains(html, "\u2713") {
			t.Errorf("expected checkmark for completed step, got: %s", html)
		}
	})
}

func TestDashboardPage(t *testing.T) {
	t.Run("renders_heading_and_description", func(t *testing.T) {
		screen := guitesting.Render(DashboardPage("dp", "Welcome", "Overview"))
		screen.Assert(t).
			HTMLContains(`class="dp"`).
			HasElement("h1").
			HasElement("p").
			TextVisible("Welcome").
			TextVisible("Overview")
	})
}

func TestSettingsCard(t *testing.T) {
	t.Run("renders_title_and_children", func(t *testing.T) {
		screen := guitesting.Render(SettingsCard("sc", "General", gui.Text("opt1")))
		screen.Assert(t).
			HTMLContains(`class="settings-card sc"`).
			HTMLContains(`data-settings-card-header`).
			HTMLContains(`data-settings-card-body`).
			TextVisible("General").
			TextVisible("opt1")
	})
}

func TestSettingsPage(t *testing.T) {
	t.Run("renders_children", func(t *testing.T) {
		screen := guitesting.Render(SettingsPage("sp", gui.Text("content")))
		screen.Assert(t).HTMLContains(`class="settings-page sp"`).TextVisible("content")
	})
	t.Run("empty_class", func(t *testing.T) {
		screen := guitesting.Render(SettingsPage(""))
		screen.Assert(t).HasElement("div").HTMLContains(`class="settings-page"`)
	})
}

func TestSettingsCardFull(t *testing.T) {
	t.Run("with_icon_and_children", func(t *testing.T) {
		screen := guitesting.Render(SettingsCardFull("scf", "icon-gear", "Prefs", gui.Text("body")))
		screen.Assert(t).
			HTMLContains(`class="scf"`).
			TextVisible("Prefs").
			TextVisible("body")
		// Icon should be present
		iElems := screen.QueryAllByTag("i")
		if len(iElems) == 0 {
			t.Error("expected icon element")
		}
	})
	t.Run("no_icon", func(t *testing.T) {
		screen := guitesting.Render(SettingsCardFull("", "", "Title", gui.Text("c")))
		screen.Assert(t).TextVisible("Title").TextVisible("c")
	})
}

func TestSettingsSection(t *testing.T) {
	t.Run("with_all_fields", func(t *testing.T) {
		screen := guitesting.Render(SettingsSection("ss", "icon-lock", "Security", "Manage access", gui.Text("child")))
		screen.Assert(t).
			HTMLContains(`class="settings-section-group ss"`).
			TextVisible("Security").
			TextVisible("Manage access").
			TextVisible("child")
	})
	t.Run("no_icon_no_desc", func(t *testing.T) {
		screen := guitesting.Render(SettingsSection("", "", "Title", ""))
		screen.Assert(t).TextVisible("Title")
	})
}

func TestSettingsSubsection(t *testing.T) {
	t.Run("renders_with_body_wrapper", func(t *testing.T) {
		screen := guitesting.Render(SettingsSubsection("sub", "icon-key", "API Keys", "Manage keys", gui.Text("list")))
		screen.Assert(t).
			HTMLContains(`class="settings-subsection sub"`).
			HTMLContains(`data-settings-subsection-header`).
			HTMLContains(`data-settings-subsection-body`).
			TextVisible("API Keys").
			TextVisible("Manage keys").
			TextVisible("list")
	})
}

func TestSettingsForm(t *testing.T) {
	t.Run("with_title", func(t *testing.T) {
		title := gui.Span()(gui.Text("Form Title"))
		screen := guitesting.Render(SettingsForm("sf", title, gui.Text("fields")))
		screen.Assert(t).
			HTMLContains(`class="settings-form sf"`).
			TextVisible("Form Title").
			TextVisible("fields")
	})
	t.Run("nil_title", func(t *testing.T) {
		screen := guitesting.Render(SettingsForm("", nil, gui.Text("f")))
		screen.Assert(t).TextVisible("f")
	})
}

func TestSettingsFormActions(t *testing.T) {
	t.Run("renders_children", func(t *testing.T) {
		screen := guitesting.Render(SettingsFormActions("sfa", gui.Text("Save")))
		screen.Assert(t).HTMLContains(`class="sfa"`).TextVisible("Save")
	})
}

func TestSettingsFormHelp(t *testing.T) {
	t.Run("renders_children", func(t *testing.T) {
		screen := guitesting.Render(SettingsFormHelp("sfh", gui.Text("Help text")))
		screen.Assert(t).HTMLContains(`class="settings-form-help sfh"`).TextVisible("Help text")
	})
}

func TestSettingsCodeInput(t *testing.T) {
	t.Run("renders_textarea", func(t *testing.T) {
		screen := guitesting.Render(SettingsCodeInput(SettingsCodeInputProps{
			Class:       "sci",
			Value:       "key: val",
			Placeholder: "Enter YAML",
			Rows:        10,
			ID:          "code-input",
		}))
		screen.Assert(t).
			HTMLContains(`class="settings-code-input sci"`).
			HasElement("textarea").
			HTMLContains(`placeholder="Enter YAML"`).
			HTMLContains(`id="code-input"`).
			HTMLContains("rows=").
			TextVisible("key: val")
	})
	t.Run("minimal", func(t *testing.T) {
		screen := guitesting.Render(SettingsCodeInput(SettingsCodeInputProps{}))
		screen.Assert(t).HasElement("textarea")
	})
}

func TestSettingsEnvRow(t *testing.T) {
	t.Run("with_badges_and_actions", func(t *testing.T) {
		badges := []gui.Node{gui.Span()(gui.Text("prod"))}
		actions := []gui.Node{gui.Span()(gui.Text("edit"))}
		screen := guitesting.Render(SettingsEnvRow("ser", "staging", badges, actions))
		screen.Assert(t).
			HTMLContains(`class="settings-env-row ser"`).
			HTMLContains(`data-settings-env-name`).
			HTMLContains(`data-settings-env-badges`).
			HTMLContains(`data-settings-env-actions`).
			TextVisible("staging").
			TextVisible("prod").
			TextVisible("edit")
	})
	t.Run("no_badges_no_actions", func(t *testing.T) {
		screen := guitesting.Render(SettingsEnvRow("", "dev", nil, nil))
		screen.Assert(t).
			HTMLContains(`data-settings-env-name`).
			TextVisible("dev")
	})
}

func TestSettingsFieldError(t *testing.T) {
	t.Run("renders_message", func(t *testing.T) {
		screen := guitesting.Render(SettingsFieldError("sfe", "Required"))
		screen.Assert(t).HTMLContains(`class="sfe"`).TextVisible("Required")
	})
}

func TestSettingsSchemaTable(t *testing.T) {
	t.Run("renders_rows", func(t *testing.T) {
		rows := []SettingsSchemaRow{
			{Type: "string", Description: "The name"},
			{Type: "int", Description: "The count"},
		}
		screen := guitesting.Render(SettingsSchemaTable("sst", rows))
		screen.Assert(t).
			HTMLContains(`class="sst"`).
			TextVisible("string").
			TextVisible("The name").
			TextVisible("int").
			TextVisible("The count")
	})
}

func TestAdminPage(t *testing.T) {
	t.Run("renders_children", func(t *testing.T) {
		screen := guitesting.Render(AdminPage("ap", gui.Text("admin")))
		screen.Assert(t).HTMLContains(`class="ap"`).TextVisible("admin")
	})
}

func TestClusterPage(t *testing.T) {
	t.Run("renders_children", func(t *testing.T) {
		screen := guitesting.Render(ClusterPage("cp", gui.Text("cluster")))
		screen.Assert(t).HTMLContains(`class="cluster-page cp"`).TextVisible("cluster")
	})
}

func TestClusterSummaryCard(t *testing.T) {
	t.Run("renders_icon_value_label", func(t *testing.T) {
		screen := guitesting.Render(ClusterSummaryCard("csc", "icon-cpu", "42", "Nodes"))
		screen.Assert(t).
			HTMLContains(`class="cluster-summary-card csc"`).
			TextVisible("42").
			TextVisible("Nodes")
		iElems := screen.QueryAllByTag("i")
		if len(iElems) == 0 {
			t.Error("expected icon element")
		}
	})
	t.Run("no_icon", func(t *testing.T) {
		screen := guitesting.Render(ClusterSummaryCard("", "", "10", "Pods"))
		screen.Assert(t).TextVisible("10").TextVisible("Pods")
	})
}

func TestClusterSummaryRow(t *testing.T) {
	t.Run("renders_children", func(t *testing.T) {
		screen := guitesting.Render(ClusterSummaryRow("csr", gui.Text("card1"), gui.Text("card2")))
		screen.Assert(t).
			HTMLContains(`class="cluster-summary csr"`).
			TextVisible("card1").
			TextVisible("card2")
	})
}
