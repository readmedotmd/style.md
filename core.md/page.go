package coremd

import (
	"fmt"

	gui "github.com/readmedotmd/gui.md"
)

// LoginPage renders a centered login page with a styled card container.
//
// CSS classes:
//   - .login-page: outer wrapper (flex centering)
//   - .login-card: card container (surface background, border, padding)
//   - .login-title: h2 title
//   - .login-error: error message
//   - .login-form: form content wrapper (stack with gap)
func LoginPage(class, title string, form gui.Node, errorMsg string) gui.Node {
	cardChildren := []gui.Node{
		gui.H2(gui.Class("login-title"))(gui.Text(title)),
	}
	if errorMsg != "" {
		cardChildren = append(cardChildren, gui.Div(gui.Class("login-error"))(gui.Text(errorMsg)))
	}
	if form != nil {
		cardChildren = append(cardChildren, gui.Form(gui.Class("login-form"))(form))
	}
	return gui.Div(collectAttrs(optClass(joinClass("login-page", class)))...)(
		gui.Div(gui.Class("login-card"))(cardChildren...),
	)
}

// SetupStep represents a single step in a setup wizard.
type SetupStep struct {
	Label     string
	Active    bool
	Completed bool
}

// SetupWizard renders a multi-step wizard with a step indicator.
//
// Data attributes on steps:
//   - data-active: "true" (on the active step)
//   - data-completed: "true" (on completed steps)
func SetupWizard(class string, steps []SetupStep, content gui.Node) gui.Node {
	stepNodes := make([]gui.Node, 0, len(steps)*2-1)
	for i, step := range steps {
		stepAttrs := []gui.Attr{}
		if step.Active {
			stepAttrs = append(stepAttrs, dataAttr("active", "true"))
		}
		if step.Completed {
			stepAttrs = append(stepAttrs, dataAttr("completed", "true"))
		}
		numberText := fmt.Sprintf("%d", i+1)
		if step.Completed {
			numberText = "\u2713"
		}
		stepNodes = append(stepNodes, gui.Div(stepAttrs...)(
			gui.Span()(gui.Text(numberText)),
			gui.Text(step.Label),
		))
		if i < len(steps)-1 {
			stepNodes = append(stepNodes, gui.Div()())
		}
	}
	return gui.Div(collectAttrs(optClass(class))...)(
		gui.Div()(stepNodes...),
		gui.Div()(content),
	)
}

// SettingsCard renders a settings section card with a header and body.
//
// Data attributes:
//   - data-settings-card-header: on the header div
//   - data-settings-card-body: on the body div
func SettingsCard(class, title string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("settings-card", class)))...)(
		gui.Div(dataAttr("settings-card-header", ""))(gui.Text(title)),
		gui.Div(dataAttr("settings-card-body", ""))(children...),
	)
}

// ─── New Page Components ───

// SettingsPage renders a centered, padded settings page container.
func SettingsPage(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("settings-page", class)))...)(children...)
}

// SettingsLayout renders a settings page with sidebar navigation + content area.
// On mobile (<768px), the sidebar collapses to horizontal wrapped pills.
//
// CSS classes:
//   - .settings-layout: outer flex container
//   - .settings-sidebar: left nav column
//   - .settings-content: right content area
func SettingsLayout(class string, sidebar gui.Node, content gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("settings-layout", class)))...)(
		gui.Div(gui.Class("settings-sidebar"))(sidebar),
		gui.Div(gui.Class("settings-content"))(content),
	)
}

// SettingsSidebarSection renders a section title in the settings sidebar.
//
// CSS classes:
//   - .settings-sidebar-section-title
func SettingsSidebarSection(title string) gui.Node {
	return gui.Div(gui.Class("settings-sidebar-section-title"))(gui.Text(title))
}

// SettingsSidebarItem renders a clickable nav item in the settings sidebar.
//
// CSS classes:
//   - .settings-sidebar-item
//   - .settings-sidebar-item.active (when active)
//
// Data attributes:
//   - data-active: "true" (when active)
func SettingsSidebarItem(class string, icon string, label string, active bool, onClick func()) gui.Node {
	cls := "settings-sidebar-item"
	if active {
		cls += " active"
	}
	attrs := collectAttrs(optClass(joinClass(cls, class)))
	if active {
		attrs = append(attrs, dataAttr("active", "true"))
	}
	if onClick != nil {
		attrs = append(attrs, gui.OnClick(onClick))
	}
	children := []gui.Node{}
	if icon != "" {
		children = append(children, gui.I(gui.Class(icon))())
	}
	children = append(children, gui.Span()(gui.Text(label)))
	return gui.Button(attrs...)(children...)
}

// settingsBlock is a shared internal helper for settings card/section/subsection pattern:
// header(icon + title + optional description) + body(children).
func settingsBlock(class, icon, title, description string, headerAttr, bodyAttr gui.Attr, children ...gui.Node) gui.Node {
	headerChildren := []gui.Node{}
	if icon != "" {
		headerChildren = append(headerChildren, gui.I(gui.Class(icon))())
	}
	headerChildren = append(headerChildren, gui.Span()(gui.Text(title)))
	if description != "" {
		headerChildren = append(headerChildren, gui.Span()(gui.Text(description)))
	}
	return gui.Div(collectAttrs(optClass(class))...)(
		gui.Div(headerAttr, dataAttr("header", ""))(headerChildren...),
		gui.Div(bodyAttr)(children...),
	)
}

// SettingsCardFull renders a settings card with colored header (icon + title) + body.
//
// Data attributes:
//   - data-header: on the header div
//   - data-settings-card-header: on the header div
//   - data-settings-card-body: on the body div
func SettingsCardFull(class, icon, title string, children ...gui.Node) gui.Node {
	return settingsBlock(class, icon, title, "", dataAttr("settings-card-header", ""), dataAttr("settings-card-body", ""), children...)
}

// SettingsSection renders a section within a settings card with border-top separator.
func SettingsSection(class, icon, title, description string, children ...gui.Node) gui.Node {
	headerChildren := []gui.Node{}
	if icon != "" {
		headerChildren = append(headerChildren, gui.I(gui.Class(icon))())
	}
	headerChildren = append(headerChildren, gui.Span()(gui.Text(title)))
	if description != "" {
		headerChildren = append(headerChildren, gui.Span()(gui.Text(description)))
	}
	all := []gui.Node{gui.Div(dataAttr("header", ""))(headerChildren...)}
	all = append(all, children...)
	return gui.Div(collectAttrs(optClass(joinClass("settings-section-group", class)))...)(all...)
}

// SettingsSubsection renders a bordered subsection within a settings card.
//
// Data attributes:
//   - data-header: on the header div
//   - data-settings-subsection-header: on the header div
//   - data-settings-subsection-body: on the body wrapper div
func SettingsSubsection(class, icon, title, description string, children ...gui.Node) gui.Node {
	return settingsBlock(joinClass("settings-subsection", class), icon, title, description, dataAttr("settings-subsection-header", ""), dataAttr("settings-subsection-body", ""), children...)
}

// SettingsForm renders a form area within settings (raised bg, bordered).
func SettingsForm(class string, title gui.Node, children ...gui.Node) gui.Node {
	all := []gui.Node{}
	if title != nil {
		all = append(all, title)
	}
	all = append(all, children...)
	return gui.Div(collectAttrs(optClass(joinClass("settings-form", class)))...)(all...)
}

// SettingsCodeInputProps configures the SettingsCodeInput component.
type SettingsCodeInputProps struct {
	Class       string
	Value       string
	Placeholder string
	Rows        int
	ID          string
	OnInput     func(gui.Event)
}

// SettingsCodeInput renders a monospace code textarea for YAML/JSON editing.
func SettingsCodeInput(props SettingsCodeInputProps) gui.Node {
	attrs := collectAttrs(optClass(joinClass("settings-code-input", props.Class)))
	if props.Placeholder != "" {
		attrs = append(attrs, gui.Placeholder(props.Placeholder))
	}
	if props.ID != "" {
		attrs = append(attrs, gui.Id(props.ID))
	}
	if props.Rows > 0 {
		attrs = append(attrs, gui.Attr_("rows", props.Rows))
	}
	if props.OnInput != nil {
		attrs = append(attrs, gui.On("input", props.OnInput))
	}
	return gui.Textarea(attrs...)(gui.Text(props.Value))
}

// ClusterSummaryCard renders a summary stat card with icon + large number + label.
func ClusterSummaryCard(class, icon, value, label string) gui.Node {
	children := []gui.Node{}
	if icon != "" {
		children = append(children, gui.I(gui.Class(icon))())
	}
	children = append(children,
		gui.Div()(gui.Text(value)),
		gui.Div()(gui.Text(label)),
	)
	return gui.Div(collectAttrs(optClass(joinClass("cluster-summary-card", class)))...)(children...)
}

// ClusterSummaryRow renders a flex row of summary cards.
func ClusterSummaryRow(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(joinClass("cluster-summary", class)))...)(children...)
}
