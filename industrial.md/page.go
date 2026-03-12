package industrialmd

import (
	coremd "github.com/readmedotmd/style.md/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	SetupStep           = coremd.SetupStep
	SettingsCodeInputProps = coremd.SettingsCodeInputProps
	SettingsSchemaRow   = coremd.SettingsSchemaRow
)

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

// SettingsPage renders a themed settings page container.
func SettingsPage(children ...gui.Node) gui.Node {
	return theme.SettingsPage(children...)
}

// SettingsCardFull renders a themed settings card with colored header.
func SettingsCardFull(icon, title string, children ...gui.Node) gui.Node {
	return theme.SettingsCardFull(icon, title, children...)
}

// SettingsSection renders a themed settings section.
func SettingsSection(icon, title, description string, children ...gui.Node) gui.Node {
	return theme.SettingsSection(icon, title, description, children...)
}

// SettingsSubsection renders a themed settings subsection.
func SettingsSubsection(icon, title, description string, children ...gui.Node) gui.Node {
	return theme.SettingsSubsection(icon, title, description, children...)
}

// SettingsForm renders a themed settings form.
func SettingsForm(title gui.Node, children ...gui.Node) gui.Node {
	return theme.SettingsForm(title, children...)
}

// SettingsFormActions renders a themed settings form actions row.
func SettingsFormActions(children ...gui.Node) gui.Node {
	return theme.SettingsFormActions(children...)
}

// SettingsFormHelp renders a themed settings form help text.
func SettingsFormHelp(children ...gui.Node) gui.Node {
	return theme.SettingsFormHelp(children...)
}

// SettingsCodeInput renders a themed code input.
func SettingsCodeInput(props SettingsCodeInputProps) gui.Node {
	return theme.SettingsCodeInput(props)
}

// SettingsEnvRow renders a themed environment row.
func SettingsEnvRow(name string, badges []gui.Node, actions []gui.Node) gui.Node {
	return theme.SettingsEnvRow(name, badges, actions)
}

// SettingsFieldError renders a themed field error.
func SettingsFieldError(message string) gui.Node {
	return theme.SettingsFieldError(message)
}

// SettingsSchemaTable renders a themed schema table.
func SettingsSchemaTable(rows []SettingsSchemaRow) gui.Node {
	return theme.SettingsSchemaTable(rows)
}

// AdminPage renders a themed admin page.
func AdminPage(children ...gui.Node) gui.Node {
	return theme.AdminPage(children...)
}

// ClusterPage renders a themed cluster page.
func ClusterPage(children ...gui.Node) gui.Node {
	return theme.ClusterPage(children...)
}

// ClusterSummaryCard renders a themed cluster summary card.
func ClusterSummaryCard(icon, value, label string) gui.Node {
	return theme.ClusterSummaryCard(icon, value, label)
}

// ClusterSummaryRow renders a themed cluster summary row.
func ClusterSummaryRow(children ...gui.Node) gui.Node {
	return theme.ClusterSummaryRow(children...)
}
