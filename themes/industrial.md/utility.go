package industrialmd

import (
	coremd "github.com/readmedotmd/style.md/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type SpinnerSize = coremd.SpinnerSize

// Re-export spinner size constants.
const (
	SpinnerDefault = coremd.SpinnerDefault
	SpinnerSmall   = coremd.SpinnerSmall
	SpinnerLarge   = coremd.SpinnerLarge
)

// Spinner renders a themed loading spinner.
func Spinner(size SpinnerSize) gui.Node {
	return theme.Spinner(size)
}

// Icon renders a themed icon element.
func Icon(class string) gui.Node {
	return theme.Icon(class)
}

// AppShellFull renders a themed top-level app shell.
func AppShellFull(scrollable bool, children ...gui.Node) gui.Node {
	return theme.AppShellFull(scrollable, children...)
}
