package coremd

import (
	"strings"

	gui "github.com/readmedotmd/gui.md"
)

// dataAttr sets a data-* attribute on an element.
func dataAttr(name, value string) gui.Attr {
	return gui.Attr_("data-"+name, value)
}

// optClass returns a gui.Class attr if cls is non-empty, otherwise nil.
func optClass(cls string) gui.Attr {
	if cls == "" {
		return nil
	}
	return gui.Class(cls)
}

// collectAttrs gathers non-nil attributes into a slice.
func collectAttrs(attrs ...gui.Attr) []gui.Attr {
	result := make([]gui.Attr, 0, len(attrs))
	for _, a := range attrs {
		if a != nil {
			result = append(result, a)
		}
	}
	return result
}

// boolStr returns "true" or "false" as a string.
func boolStr(v bool) string {
	if v {
		return "true"
	}
	return "false"
}

// Classes joins multiple CSS class names into a single string.
func Classes(names ...string) string {
	return strings.Join(names, " ")
}

// ClassIf returns the base class with the conditional class appended if condition is true.
func ClassIf(base string, condition bool, conditional string) string {
	if condition {
		return base + " " + conditional
	}
	return base
}
