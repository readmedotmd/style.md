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
//
// Data attributes:
//   - data-form-group
func FormGroup(props FormGroupProps, children ...gui.Node) gui.Node {
	all := []gui.Node{
		gui.Label()(gui.Text(props.Label)),
	}
	all = append(all, children...)
	return gui.Div(collectAttrs(optClass(props.Class), dataAttr("form-group", ""))...)(all...)
}

// TextInputProps configures the TextInput component.
type TextInputProps struct {
	Class       string
	Placeholder string
	Value       string
	Type        string
	ID          string
	Name        string
	Error       bool
	Min         string // min attribute (for type="number")
	Max         string // max attribute (for type="number")
	Step        string // step attribute (for type="number")
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
	if props.Name != "" {
		attrs = append(attrs, gui.Name(props.Name))
	}
	if props.Min != "" {
		attrs = append(attrs, gui.Attr_("min", props.Min))
	}
	if props.Max != "" {
		attrs = append(attrs, gui.Attr_("max", props.Max))
	}
	if props.Step != "" {
		attrs = append(attrs, gui.Attr_("step", props.Step))
	}
	if props.OnInput != nil {
		attrs = append(attrs, gui.On("input", props.OnInput))
	}
	return gui.Input(attrs...)()
}

// NumberInput renders a number input field. It is a convenience wrapper around
// TextInput with Type set to "number".
//
// Data attributes:
//   - data-error: "true" (present only when Error is true)
func NumberInput(props TextInputProps) gui.Node {
	props.Type = "number"
	return TextInput(props)
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
	attrs := collectAttrs(optClass(joinClass("text-area", props.Class)))
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
	ID       string
	Name     string
	Value    string
	OnChange func()
}

// Checkbox renders an unstyled checkbox with a label.
func Checkbox(props CheckboxProps) gui.Node {
	inputAttrs := []gui.Attr{gui.Type("checkbox")}
	if props.Checked {
		inputAttrs = append(inputAttrs, gui.Checked(true))
	}
	if props.ID != "" {
		inputAttrs = append(inputAttrs, gui.Id(props.ID))
	}
	if props.Name != "" {
		inputAttrs = append(inputAttrs, gui.Name(props.Name))
	}
	if props.Value != "" {
		inputAttrs = append(inputAttrs, gui.Value(props.Value))
	}
	if props.OnChange != nil {
		inputAttrs = append(inputAttrs, gui.OnClick(props.OnChange))
	}
	return gui.Label(collectAttrs(optClass(joinClass("checkbox", props.Class)))...)(
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
//
// Data attributes:
//   - data-feature-info: on the info container
//   - data-feature-name: on the name div
//   - data-feature-desc: on the description div
func FeatureRow(props FeatureRowProps) gui.Node {
	inputAttrs := []gui.Attr{gui.Type("checkbox")}
	if props.Checked {
		inputAttrs = append(inputAttrs, gui.Checked(true))
	}
	if props.OnChange != nil {
		inputAttrs = append(inputAttrs, gui.OnClick(props.OnChange))
	}
	return gui.Div(collectAttrs(optClass(joinClass("feature-row", props.Class)))...)(
		gui.Div(dataAttr("feature-info", ""))(
			gui.Div(dataAttr("feature-name", ""))(gui.Text(props.Name)),
			gui.Div(dataAttr("feature-desc", ""))(gui.Text(props.Description)),
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
	attrs := collectAttrs(optClass(joinClass("variable-row", props.Class)))
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
	attrs := collectAttrs(optClass(joinClass("error-message", class)))
	attrs = append(attrs, gui.Attr_("role", "alert"))
	return gui.Div(attrs...)(gui.Text(text))
}

// SuccessMessage renders a success message.
func SuccessMessage(class, text string) gui.Node {
	attrs := collectAttrs(optClass(joinClass("success-message", class)))
	attrs = append(attrs, gui.Attr_("role", "status"))
	return gui.Div(attrs...)(gui.Text(text))
}

// EditableVariableRowProps configures the EditableVariableRow component.
type EditableVariableRowProps struct {
	Class       string
	Key         string
	Value       string
	Passthrough bool
	OnKeyInput  func(gui.Event)
	OnValInput  func(gui.Event)
	OnToggle    func()
	OnRemove    func()
}

// EditableVariableRow renders an editable key/value row with an optional
// passthrough checkbox and delete button. Useful for environment variables,
// configuration maps, and similar key-value editors.
//
// Data attributes:
//   - data-editable-var-row
//   - data-passthrough: "true" (when Passthrough is true)
func EditableVariableRow(props EditableVariableRowProps) gui.Node {
	attrs := collectAttrs(optClass(props.Class), dataAttr("editable-var-row", ""))
	if props.Passthrough {
		attrs = append(attrs, dataAttr("passthrough", "true"))
	}
	children := []gui.Node{
		TextInput(TextInputProps{Placeholder: "Key", Value: props.Key, OnInput: props.OnKeyInput}),
		TextInput(TextInputProps{Placeholder: "Value", Value: props.Value, OnInput: props.OnValInput}),
	}
	if props.OnToggle != nil {
		children = append(children, Checkbox(CheckboxProps{
			Label:    "Passthrough",
			Checked:  props.Passthrough,
			OnChange: props.OnToggle,
		}))
	}
	if props.OnRemove != nil {
		children = append(children, Button(ButtonProps{
			Size:    ButtonSmall,
			Variant: ButtonDanger,
			OnClick: props.OnRemove,
		}, gui.Text("Remove")))
	}
	return gui.Div(attrs...)(children...)
}

// PasswordFieldProps configures the PasswordField component.
type PasswordFieldProps struct {
	Class       string
	Placeholder string
	Value       string
	ID          string
	Name        string
	Visible     bool
	OnInput     func(gui.Event)
	OnToggle    func()
}

// PasswordField renders a password input with a visibility toggle button.
//
// Data attributes:
//   - data-password-field
//   - data-visible: "true" (when password is visible)
func PasswordField(props PasswordFieldProps) gui.Node {
	attrs := collectAttrs(optClass(joinClass("password-field", props.Class)), dataAttr("password-field", ""))
	if props.Visible {
		attrs = append(attrs, dataAttr("visible", "true"))
	}
	inputType := "password"
	if props.Visible {
		inputType = "text"
	}
	input := TextInput(TextInputProps{
		Placeholder: props.Placeholder,
		Value:       props.Value,
		Type:        inputType,
		ID:          props.ID,
		Name:        props.Name,
		OnInput:     props.OnInput,
	})
	toggleAttrs := []gui.Attr{dataAttr("password-toggle", "")}
	if props.OnToggle != nil {
		toggleAttrs = append(toggleAttrs, gui.OnClick(props.OnToggle))
	}
	icon := "icon-eye"
	if props.Visible {
		icon = "icon-eye-off"
	}
	toggle := gui.Button(toggleAttrs...)(gui.I(gui.Class(icon))())
	return gui.Div(attrs...)(input, toggle)
}

// SecretFieldProps configures the SecretField component.
type SecretFieldProps struct {
	Class    string
	KeyName  string
	Value    string
	Scope    string
	OnCopy   func()
	OnRemove func()
}

// SecretField renders a masked secret value row with key name, masked value, copy button, scope badge, and remove button.
//
// Data attributes:
//   - data-secret-field
func SecretField(props SecretFieldProps) gui.Node {
	attrs := collectAttrs(optClass(joinClass("secret-field", props.Class)), dataAttr("secret-field", ""))

	children := []gui.Node{
		gui.Span(dataAttr("secret-key", ""))(gui.Text(props.KeyName)),
	}

	valueChildren := []gui.Node{gui.Span()(gui.Text("\u2022\u2022\u2022\u2022\u2022\u2022\u2022\u2022"))}
	if props.OnCopy != nil {
		valueChildren = append(valueChildren, gui.Button(gui.OnClick(props.OnCopy), dataAttr("secret-copy", ""))(gui.I(gui.Class("icon-copy"))()))
	}
	children = append(children, gui.Div(dataAttr("secret-value", ""))(valueChildren...))

	if props.Scope != "" {
		children = append(children, gui.Span(dataAttr("secret-scope", ""))(gui.Text(props.Scope)))
	}
	if props.OnRemove != nil {
		children = append(children, gui.Button(gui.OnClick(props.OnRemove), dataAttr("secret-remove", ""))(gui.Text("\u00d7")))
	}

	return gui.Div(attrs...)(children...)
}

// SchemaFieldProps configures the SchemaField component.
type SchemaFieldProps struct {
	Class       string
	Name        string
	Type        string
	Description string
}

// SchemaField renders a documentation entry for a single schema field,
// displaying the field name, its type, and a description.
//
// Data attributes:
//   - data-schema-field
func SchemaField(props SchemaFieldProps) gui.Node {
	attrs := collectAttrs(optClass(props.Class), dataAttr("schema-field", ""))
	return gui.Div(attrs...)(
		gui.Span(dataAttr("schema-field-name", ""))(gui.Text(props.Name)),
		gui.Span(dataAttr("schema-field-type", ""))(gui.Text(props.Type)),
		gui.Span(dataAttr("schema-field-desc", ""))(gui.Text(props.Description)),
	)
}
