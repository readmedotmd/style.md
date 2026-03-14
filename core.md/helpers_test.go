package coremd

import "testing"

func TestClasses(t *testing.T) {
	t.Run("multiple_names", func(t *testing.T) {
		got := Classes("btn", "primary", "large")
		if got != "btn primary large" {
			t.Errorf("expected 'btn primary large', got %q", got)
		}
	})
	t.Run("single_name", func(t *testing.T) {
		got := Classes("btn")
		if got != "btn" {
			t.Errorf("expected 'btn', got %q", got)
		}
	})
	t.Run("empty", func(t *testing.T) {
		got := Classes()
		if got != "" {
			t.Errorf("expected empty string, got %q", got)
		}
	})
}

func TestClassIf(t *testing.T) {
	t.Run("condition_true", func(t *testing.T) {
		got := ClassIf("btn", true, "active")
		if got != "btn active" {
			t.Errorf("expected 'btn active', got %q", got)
		}
	})
	t.Run("condition_false", func(t *testing.T) {
		got := ClassIf("btn", false, "active")
		if got != "btn" {
			t.Errorf("expected 'btn', got %q", got)
		}
	})
}
