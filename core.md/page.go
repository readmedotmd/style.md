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
