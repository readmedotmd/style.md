package coremd

import (
	"fmt"

	gui "github.com/readmedotmd/gui.md"
)

// LoginPage renders a centered login page.
func LoginPage(class, title string, form gui.Node, errorMsg string) gui.Node {
	cardChildren := []gui.Node{
		gui.Div()(gui.Text(title)),
	}
	if errorMsg != "" {
		cardChildren = append(cardChildren, ErrorMessage("", errorMsg))
	}
	if form != nil {
		cardChildren = append(cardChildren, form)
	}
	return gui.Div(collectAttrs(optClass(class))...)(
		gui.Div()(cardChildren...),
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

// DashboardPage renders a dashboard page with a heading and description.
func DashboardPage(class, heading, description string) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(
		gui.H1()(gui.Text(heading)),
		gui.P()(gui.Text(description)),
	)
}

// SettingsCard renders a settings section card.
func SettingsCard(class, title string, children ...gui.Node) gui.Node {
	all := []gui.Node{
		gui.H2()(gui.Text(title)),
	}
	all = append(all, children...)
	return gui.Div(collectAttrs(optClass(class))...)(
		gui.Div()(all...),
	)
}

// ─── New Page Components ───

// SettingsPage renders a centered, padded settings page container.
func SettingsPage(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// SettingsCardFull renders a settings card with colored header (icon + title) + body.
func SettingsCardFull(class, icon, title string, children ...gui.Node) gui.Node {
	headerChildren := []gui.Node{}
	if icon != "" {
		headerChildren = append(headerChildren, gui.I(gui.Class(icon))())
	}
	headerChildren = append(headerChildren, gui.Text(title))
	all := []gui.Node{
		gui.Div()(headerChildren...),
		gui.Div()(children...),
	}
	return gui.Div(collectAttrs(optClass(class))...)(all...)
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
	all := []gui.Node{gui.Div()(headerChildren...)}
	all = append(all, children...)
	return gui.Div(collectAttrs(optClass(class))...)(all...)
}

// SettingsSubsection renders a bordered subsection within a settings card.
func SettingsSubsection(class, icon, title, description string, children ...gui.Node) gui.Node {
	headerChildren := []gui.Node{}
	if icon != "" {
		headerChildren = append(headerChildren, gui.I(gui.Class(icon))())
	}
	headerChildren = append(headerChildren, gui.Span()(gui.Text(title)))
	if description != "" {
		headerChildren = append(headerChildren, gui.Span()(gui.Text(description)))
	}
	all := []gui.Node{gui.Div()(headerChildren...)}
	all = append(all, children...)
	return gui.Div(collectAttrs(optClass(class))...)(all...)
}

// SettingsForm renders a form area within settings (raised bg, bordered).
func SettingsForm(class string, title gui.Node, children ...gui.Node) gui.Node {
	all := []gui.Node{}
	if title != nil {
		all = append(all, title)
	}
	all = append(all, children...)
	return gui.Div(collectAttrs(optClass(class))...)(all...)
}

// SettingsFormActions renders a button row at the bottom of a settings form.
func SettingsFormActions(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// SettingsFormHelp renders a help text block within a settings form.
func SettingsFormHelp(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
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
	attrs := collectAttrs(optClass(props.Class))
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

// SettingsEnvRow renders a row in the environments list showing name + badges + actions.
func SettingsEnvRow(class, name string, badges []gui.Node, actions []gui.Node) gui.Node {
	children := []gui.Node{
		gui.Span()(gui.Text(name)),
	}
	if len(badges) > 0 {
		children = append(children, gui.Div()(badges...))
	}
	if len(actions) > 0 {
		children = append(children, gui.Div()(actions...))
	}
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// SettingsFieldError renders a small red field-level error text.
func SettingsFieldError(class, message string) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(gui.Text(message))
}

// SettingsSchemaRow represents a row in a schema documentation table.
type SettingsSchemaRow struct {
	Type        string
	Description string
}

// SettingsSchemaTable renders a documentation table for schema fields.
func SettingsSchemaTable(class string, rows []SettingsSchemaRow) gui.Node {
	children := make([]gui.Node, len(rows))
	for i, row := range rows {
		children[i] = gui.Div()(
			gui.Span()(gui.Text(row.Type)),
			gui.Span()(gui.Text(row.Description)),
		)
	}
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// AdminPage renders a wider padded admin page container.
func AdminPage(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// ClusterPage renders a cluster stats page container.
func ClusterPage(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
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
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}

// ClusterSummaryRow renders a flex row of summary cards.
func ClusterSummaryRow(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class))...)(children...)
}
