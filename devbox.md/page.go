package devboxmd

import (
	coremd "github.com/readmedotmd/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type SetupStep = coremd.SetupStep

// LoginPage renders a themed centered login page.
func LoginPage(title string, form gui.Node, errorMsg string) gui.Node {
	return theme.LoginPage(title, form, errorMsg)
}

// SetupWizard renders a themed multi-step wizard.
func SetupWizard(steps []SetupStep, content gui.Node) gui.Node {
	return theme.SetupWizard(steps, content)
}

// DashboardPage renders a themed dashboard page.
func DashboardPage(heading, description string) gui.Node {
	return theme.DashboardPage(heading, description)
}

// SettingsCard renders a themed settings section card.
func SettingsCard(title string, children ...gui.Node) gui.Node {
	return theme.SettingsCard(title, children...)
}
