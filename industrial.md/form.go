package industrialmd

import (
	coremd "github.com/readmedotmd/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	TextInputProps  = coremd.TextInputProps
	TextareaProps   = coremd.TextareaProps
	SelectProps     = coremd.SelectProps
	CheckboxProps   = coremd.CheckboxProps
	FeatureRowProps = coremd.FeatureRowProps
	VariableRowProps = coremd.VariableRowProps
)

// FormGroup renders a themed form group.
func FormGroup(label string, children ...gui.Node) gui.Node {
	return theme.FormGroup(label, children...)
}

// TextInput renders a themed text input field.
func TextInput(props TextInputProps) gui.Node {
	return theme.TextInput(props)
}

// TextArea renders a themed textarea.
func TextArea(props TextareaProps) gui.Node {
	return theme.TextArea(props)
}

// SelectInput renders a themed select dropdown.
func SelectInput(props SelectProps, options ...gui.Node) gui.Node {
	return theme.SelectInput(props, options...)
}

// SelectOption renders a single option element (pass-through).
func SelectOption(value, label string, selected bool) gui.Node {
	return coremd.SelectOption(value, label, selected)
}

// Checkbox renders a themed checkbox with a label.
func Checkbox(props CheckboxProps) gui.Node {
	return theme.Checkbox(props)
}

// FeatureRow renders a themed feature toggle row.
func FeatureRow(props FeatureRowProps) gui.Node {
	return theme.FeatureRow(props)
}

// VariableRow renders a themed key-value variable row.
func VariableRow(props VariableRowProps) gui.Node {
	return theme.VariableRow(props)
}

// ErrorMessage renders a themed error message.
func ErrorMessage(text string) gui.Node {
	return theme.ErrorMessage(text)
}

// SuccessMessage renders a themed success message.
func SuccessMessage(text string) gui.Node {
	return theme.SuccessMessage(text)
}
