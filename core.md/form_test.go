package coremd

import (
	"testing"

	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestNumberInput(t *testing.T) {
	t.Run("renders_number_type", func(t *testing.T) {
		s := guitesting.Render(NumberInput(TextInputProps{
			Placeholder: "Count",
			Value:       "5",
			Min:         "0",
			Max:         "100",
			Step:        "1",
		}))
		a := s.Assert(t)
		a.HasElement("input")
		a.HTMLContains(`type="number"`)
		a.HTMLContains(`min="0"`)
		a.HTMLContains(`max="100"`)
		a.HTMLContains(`step="1"`)
		a.HTMLContains(`value="5"`)
	})
	t.Run("minimal", func(t *testing.T) {
		s := guitesting.Render(NumberInput(TextInputProps{}))
		a := s.Assert(t)
		a.HTMLContains(`type="number"`)
		a.HTMLNotContains("min=")
		a.HTMLNotContains("max=")
		a.HTMLNotContains("step=")
	})
}

func TestTextInput_MinMaxStep(t *testing.T) {
	s := guitesting.Render(TextInput(TextInputProps{
		Type: "number",
		Min:  "1",
		Max:  "10",
		Step: "0.5",
		Name: "quantity",
	}))
	a := s.Assert(t)
	a.HTMLContains(`min="1"`)
	a.HTMLContains(`max="10"`)
	a.HTMLContains(`step="0.5"`)
	a.HTMLContains(`name="quantity"`)
}

func TestCheckbox_IDNameValue(t *testing.T) {
	t.Run("with_id_name_value", func(t *testing.T) {
		s := guitesting.Render(Checkbox(CheckboxProps{
			Label:   "Enable",
			ID:      "cb-enable",
			Name:    "enable",
			Value:   "yes",
			Checked: true,
		}))
		a := s.Assert(t)
		a.HTMLContains(`id="cb-enable"`)
		a.HTMLContains(`name="enable"`)
		a.HTMLContains(`value="yes"`)
		a.HTMLContains(`checked`)
		a.TextVisible("Enable")
	})
	t.Run("without_optional_fields", func(t *testing.T) {
		s := guitesting.Render(Checkbox(CheckboxProps{Label: "Toggle"}))
		a := s.Assert(t)
		a.HTMLNotContains("id=")
		a.HTMLNotContains("name=")
		a.HTMLNotContains(`value=`)
		a.TextVisible("Toggle")
	})
}

func TestEditableVariableRow(t *testing.T) {
	t.Run("full_props", func(t *testing.T) {
		removed := false
		toggled := false
		s := guitesting.Render(EditableVariableRow(EditableVariableRowProps{
			Class:       "evr",
			Key:         "API_KEY",
			Value:       "secret",
			Passthrough: true,
			OnToggle:    func() { toggled = true },
			OnRemove:    func() { removed = true },
		}))
		a := s.Assert(t)
		a.HTMLContains(`class="evr"`)
		a.HTMLContains(`data-editable-var-row`)
		a.HTMLContains(`data-passthrough="true"`)
		a.HTMLContains(`value="API_KEY"`)
		a.HTMLContains(`value="secret"`)
		a.TextVisible("Passthrough")
		a.TextVisible("Remove")
		// Verify both inputs are present
		inputs := s.QueryAllByTag("input")
		if len(inputs) < 2 {
			t.Errorf("expected at least 2 inputs, got %d", len(inputs))
		}
		_ = removed
		_ = toggled
	})
	t.Run("minimal", func(t *testing.T) {
		s := guitesting.Render(EditableVariableRow(EditableVariableRowProps{}))
		a := s.Assert(t)
		a.HTMLContains(`data-editable-var-row`)
		a.HTMLNotContains("data-passthrough")
		a.HTMLNotContains("Remove")
		a.HTMLNotContains("Passthrough")
	})
}

func TestSchemaField(t *testing.T) {
	t.Run("all_fields", func(t *testing.T) {
		s := guitesting.Render(SchemaField(SchemaFieldProps{
			Class:       "sf",
			Name:        "port",
			Type:        "integer",
			Description: "The port to listen on",
		}))
		a := s.Assert(t)
		a.HTMLContains(`class="sf"`)
		a.HTMLContains(`data-schema-field`)
		a.HTMLContains(`data-schema-field-name`)
		a.HTMLContains(`data-schema-field-type`)
		a.HTMLContains(`data-schema-field-desc`)
		a.TextVisible("port")
		a.TextVisible("integer")
		a.TextVisible("The port to listen on")
	})
	t.Run("empty_class", func(t *testing.T) {
		s := guitesting.Render(SchemaField(SchemaFieldProps{Name: "x", Type: "y", Description: "z"}))
		a := s.Assert(t)
		a.HTMLContains(`data-schema-field`)
		a.HTMLNotContains("class=")
	})
}

func TestPasswordField(t *testing.T) {
	t.Run("hidden_password", func(t *testing.T) {
		s := guitesting.Render(PasswordField(PasswordFieldProps{
			Class:       "pf",
			Placeholder: "Enter password",
			Value:       "secret",
			Visible:     false,
		}))
		a := s.Assert(t)
		a.HTMLContains(`class="password-field pf"`)
		a.HTMLContains(`data-password-field`)
		a.HTMLNotContains(`data-visible`)
		a.HTMLContains(`type="password"`)
		a.HTMLContains(`data-password-toggle`)
		a.HTMLContains(`class="icon-eye"`)
	})

	t.Run("visible_password", func(t *testing.T) {
		s := guitesting.Render(PasswordField(PasswordFieldProps{
			Visible: true,
		}))
		a := s.Assert(t)
		a.HTMLContains(`data-visible="true"`)
		a.HTMLContains(`type="text"`)
		a.HTMLContains(`class="icon-eye-off"`)
	})

	t.Run("toggle_click", func(t *testing.T) {
		toggled := false
		s := guitesting.Render(PasswordField(PasswordFieldProps{
			OnToggle: func() { toggled = true },
		}))
		btns := s.QueryAllByTag("button")
		for _, b := range btns {
			s.Click(b)
			break
		}
		if !toggled {
			t.Error("expected OnToggle to fire")
		}
	})
}

func TestSecretField(t *testing.T) {
	t.Run("full_props", func(t *testing.T) {
		copied := false
		removed := false
		s := guitesting.Render(SecretField(SecretFieldProps{
			Class:    "sf",
			KeyName:  "API_KEY",
			Scope:    "agent",
			OnCopy:   func() { copied = true },
			OnRemove: func() { removed = true },
		}))
		a := s.Assert(t)
		a.HTMLContains(`class="secret-field sf"`)
		a.HTMLContains(`data-secret-field`)
		a.HTMLContains(`data-secret-key`)
		a.HTMLContains(`data-secret-value`)
		a.HTMLContains(`data-secret-scope`)
		a.HTMLContains(`data-secret-copy`)
		a.HTMLContains(`data-secret-remove`)
		a.TextVisible("API_KEY")
		a.TextVisible("agent")

		removeBtn := s.GetByText("\u00d7")
		s.Click(removeBtn)
		if !removed {
			t.Error("expected OnRemove to fire")
		}
		_ = copied
	})

	t.Run("minimal", func(t *testing.T) {
		s := guitesting.Render(SecretField(SecretFieldProps{
			KeyName: "TOKEN",
		}))
		a := s.Assert(t)
		a.HTMLContains(`data-secret-field`)
		a.TextVisible("TOKEN")
		a.HTMLNotContains("data-secret-scope")
		a.HTMLNotContains("data-secret-copy")
		a.HTMLNotContains("data-secret-remove")
	})
}
