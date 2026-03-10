package coremd

import (
	"strconv"

	gui "github.com/readmedotmd/gui.md"
)

// FormGroupProps configures the FormGroup component.
type FormGroupProps struct {
	Class string
	Label string
}

// FormGroup renders a labeled form field group.
func FormGroup(props FormGroupProps, children ...gui.Node) gui.Node {
	all := []gui.Node{
		gui.Label()(gui.Text(props.Label)),
	}
	all = append(all, children...)
	return gui.Div(collectAttrs(optClass(props.Class))...)(all...)
}

// TextInputProps configures the TextInput component.
type TextInputProps struct {
	Class       string
	Placeholder string
	Value       string
	Type        string
	ID          string
	Error       bool
	OnInput     func(gui.Event)
}

// TextInput renders an unstyled text input field.
//
// Data attributes:
//   - data-error: "true" (present only when Error is true)
func TextInput(props TextInputProps) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.Error {
		attrs = append(attrs, dataAttr("error", "true"))
	}
	if props.Placeholder != "" {
		attrs = append(attrs, gui.Placeholder(props.Placeholder))
	}
	if props.Value != "" {
		attrs = append(attrs, gui.Value(props.Value))
	}
	if props.Type != "" {
		attrs = append(attrs, gui.Type(props.Type))
	} else {
		attrs = append(attrs, gui.Type("text"))
	}
	if props.ID != "" {
		attrs = append(attrs, gui.Id(props.ID))
	}
	if props.OnInput != nil {
		attrs = append(attrs, gui.On("input", props.OnInput))
	}
	return gui.Input(attrs...)()
}

// TextareaProps configures the TextArea component.
type TextareaProps struct {
	Class       string
	Placeholder string
	Value       string
	ID          string
	AutoGrow    bool
	Fixed       bool
	Rows        int
	OnInput     func(gui.Event)
}

// TextArea renders an unstyled textarea element.
//
// Data attributes:
//   - data-auto-grow: "true" (when AutoGrow is true)
//   - data-fixed: "true" (when Fixed is true)
func TextArea(props TextareaProps) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.AutoGrow {
		attrs = append(attrs, dataAttr("auto-grow", "true"))
	}
	if props.Fixed {
		attrs = append(attrs, dataAttr("fixed", "true"))
	}
	if props.Placeholder != "" {
		attrs = append(attrs, gui.Placeholder(props.Placeholder))
	}
	if props.ID != "" {
		attrs = append(attrs, gui.Id(props.ID))
	}
	if props.Rows > 0 {
		attrs = append(attrs, gui.Attr_("rows", strconv.Itoa(props.Rows)))
	}
	if props.OnInput != nil {
		attrs = append(attrs, gui.On("input", props.OnInput))
	}
	return gui.Textarea(attrs...)(gui.Text(props.Value))
}

// SelectProps configures the SelectInput component.
type SelectProps struct {
	Class    string
	ID       string
	OnChange func(gui.Event)
}

// SelectInput renders an unstyled select dropdown.
func SelectInput(props SelectProps, options ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.ID != "" {
		attrs = append(attrs, gui.Id(props.ID))
	}
	if props.OnChange != nil {
		attrs = append(attrs, gui.On("change", props.OnChange))
	}
	return gui.Select(attrs...)(options...)
}

// SelectOption renders a single option element.
func SelectOption(value, label string, selected bool) gui.Node {
	props := []gui.Attr{gui.Value(value)}
	if selected {
		props = append(props, gui.Attr_("selected", "true"))
	}
	return gui.Option(props...)(gui.Text(label))
}

// CheckboxProps configures the Checkbox component.
type CheckboxProps struct {
	Class    string
	Label    string
	Checked  bool
	OnChange func()
}

// Checkbox renders an unstyled checkbox with a label.
func Checkbox(props CheckboxProps) gui.Node {
	inputAttrs := []gui.Attr{gui.Type("checkbox")}
	if props.Checked {
		inputAttrs = append(inputAttrs, gui.Checked(true))
	}
	if props.OnChange != nil {
		inputAttrs = append(inputAttrs, gui.OnClick(props.OnChange))
	}
	return gui.Label(collectAttrs(optClass(props.Class))...)(
		gui.Input(inputAttrs...)(),
		gui.Span()(gui.Text(props.Label)),
	)
}

// FeatureRowProps configures the FeatureRow component.
type FeatureRowProps struct {
	Class       string
	Name        string
	Description string
	Checked     bool
	OnChange    func()
}

// FeatureRow renders a feature toggle row with name, description, and checkbox.
func FeatureRow(props FeatureRowProps) gui.Node {
	inputAttrs := []gui.Attr{gui.Type("checkbox")}
	if props.Checked {
		inputAttrs = append(inputAttrs, gui.Checked(true))
	}
	if props.OnChange != nil {
		inputAttrs = append(inputAttrs, gui.OnClick(props.OnChange))
	}
	return gui.Div(collectAttrs(optClass(props.Class))...)(
		gui.Div()(
			gui.Div()(gui.Text(props.Name)),
			gui.Div()(gui.Text(props.Description)),
		),
		gui.Input(inputAttrs...)(),
	)
}

// VariableRowProps configures the VariableRow component.
type VariableRowProps struct {
	Class       string
	Key         string
	Value       string
	Masked      bool
	Passthrough bool
	OnRemove    func()
}

// VariableRow renders a key-value variable row with actions.
//
// Data attributes:
//   - data-masked: "true" (when value is masked)
func VariableRow(props VariableRowProps) gui.Node {
	displayValue := props.Value
	attrs := collectAttrs(optClass(props.Class))
	if props.Masked {
		displayValue = "••••••••"
		attrs = append(attrs, dataAttr("masked", "true"))
	}
	children := []gui.Node{
		gui.Span()(gui.Text(props.Key)),
		gui.Span()(gui.Text(displayValue)),
	}
	if props.OnRemove != nil {
		children = append(children, gui.Div()(
			Button(ButtonProps{Size: ButtonSmall, OnClick: props.OnRemove}, gui.Text("Remove")),
		))
	}
	return gui.Div(attrs...)(children...)
}

// ErrorMessage renders an error message.
func ErrorMessage(class, text string) gui.Node {
	attrs := collectAttrs(optClass(class))
	attrs = append(attrs, gui.Attr_("role", "alert"))
	return gui.Div(attrs...)(gui.Text(text))
}

// SuccessMessage renders a success message.
func SuccessMessage(class, text string) gui.Node {
	attrs := collectAttrs(optClass(class))
	attrs = append(attrs, gui.Attr_("role", "status"))
	return gui.Div(attrs...)(gui.Text(text))
}
