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
	ButtonGhost   ButtonVariant = "ghost"
)

// ButtonSize determines the size of a button.
type ButtonSize string

const (
	ButtonMedium ButtonSize = ""
	ButtonSmall  ButtonSize = "small"
)

// ButtonProps configures the Button component.
type ButtonProps struct {
	Class     string
	Variant   ButtonVariant
	Size      ButtonSize
	Icon      string // CSS icon class (e.g. "icon-settings"); renders an <i> before children
	AriaLabel string // aria-label for icon-only buttons
	Disabled  bool
	OnClick   func()
}

// Button renders a button element with data attributes for theming.
//
// When Icon is set, an <i class="icon"> element is prepended to children.
// Use Variant ButtonGhost for borderless icon-only buttons (replaces IconButton).
// Use Variant ButtonToolbar with Icon for icon+label toolbar buttons (replaces ToolbarButton).
//
// Data attributes:
//   - data-variant: "primary", "danger", "toolbar", "ghost" (omitted when default)
//   - data-size: "small" (omitted when medium/default)
func Button(props ButtonProps, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(joinClass("btn", props.Class)))
	if props.Variant != "" {
		attrs = append(attrs, dataAttr("variant", string(props.Variant)))
	}
	if props.Size != "" {
		attrs = append(attrs, dataAttr("size", string(props.Size)))
	}
	if props.AriaLabel != "" {
		attrs = append(attrs, gui.Attr_("aria-label", props.AriaLabel))
	}
	if props.Disabled {
		attrs = append(attrs, gui.Disabled(true))
	}
	if props.OnClick != nil {
		attrs = append(attrs, gui.OnClick(props.OnClick))
	}
	if props.Icon != "" {
		iconNode := gui.I(gui.Class(props.Icon))()
		children = append([]gui.Node{iconNode}, children...)
	}
	return gui.Button(attrs...)(children...)
}
