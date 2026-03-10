package coremd

import (
	gui "github.com/readmedotmd/gui.md"
)

// ButtonVariant determines the semantic variant of a button.
type ButtonVariant string

const (
	ButtonDefault ButtonVariant = ""
	ButtonPrimary ButtonVariant = "primary"
	ButtonDanger  ButtonVariant = "danger"
	ButtonToolbar ButtonVariant = "toolbar"
)

// ButtonSize determines the size of a button.
type ButtonSize string

const (
	ButtonMedium ButtonSize = ""
	ButtonSmall  ButtonSize = "small"
)

// ButtonProps configures the Button component.
type ButtonProps struct {
	Class    string
	Variant  ButtonVariant
	Size     ButtonSize
	Disabled bool
	OnClick  func()
}

// Button renders an unstyled button element with data attributes for theming.
//
// Data attributes:
//   - data-variant: "primary", "danger", "toolbar" (omitted when default)
//   - data-size: "small" (omitted when medium/default)
func Button(props ButtonProps, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.Variant != "" {
		attrs = append(attrs, dataAttr("variant", string(props.Variant)))
	}
	if props.Size != "" {
		attrs = append(attrs, dataAttr("size", string(props.Size)))
	}
	if props.Disabled {
		attrs = append(attrs, gui.Disabled(true))
	}
	if props.OnClick != nil {
		attrs = append(attrs, gui.OnClick(props.OnClick))
	}
	return gui.Button(attrs...)(children...)
}
