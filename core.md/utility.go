package coremd

import (
	gui "github.com/readmedotmd/gui.md"
)

// SpinnerSize controls the size of a spinner.
type SpinnerSize string

const (
	SpinnerDefault SpinnerSize = ""
	SpinnerSmall   SpinnerSize = "small"
	SpinnerLarge   SpinnerSize = "large"
)

// SpinnerProps configures the Spinner component.
type SpinnerProps struct {
	Class string
	Size  SpinnerSize
}

// Spinner renders a loading spinner element.
//
// Data attributes:
//   - data-size: "small" or "large" (omitted when default)
func Spinner(props SpinnerProps) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.Size != "" {
		attrs = append(attrs, dataAttr("size", string(props.Size)))
	}
	return gui.Div(attrs...)()
}

// Icon renders an icon element with the given CSS class.
func Icon(class, iconClass string) gui.Node {
	attrs := collectAttrs(optClass(class))
	return gui.I(attrs...)(gui.I(gui.Class(iconClass))())
}

// AppShellFull renders a top-level app shell with safe area handling, 100vh height,
// and optional scrollable mode.
//
// Data attributes:
//   - data-scrollable: "true" (when scrollable is true)
func AppShellFull(class string, scrollable bool, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(class))
	if scrollable {
		attrs = append(attrs, dataAttr("scrollable", "true"))
	}
	return gui.Div(attrs...)(children...)
}
