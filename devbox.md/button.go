package devboxmd

import (
	coremd "github.com/readmedotmd/style.md/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types for convenience.
type (
	ButtonVariant = coremd.ButtonVariant
	ButtonSize    = coremd.ButtonSize
	ButtonProps   = coremd.ButtonProps
)

// Re-export core constants.
const (
	ButtonDefault = coremd.ButtonDefault
	ButtonPrimary = coremd.ButtonPrimary
	ButtonDanger  = coremd.ButtonDanger
	ButtonToolbar = coremd.ButtonToolbar
	ButtonMedium  = coremd.ButtonMedium
	ButtonSmall   = coremd.ButtonSmall
)

// Button renders a themed button with Devbox design classes.
func Button(props ButtonProps, children ...gui.Node) gui.Node {
	return theme.Button(props, children...)
}
