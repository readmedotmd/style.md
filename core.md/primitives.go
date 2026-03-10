package coremd

import (
	gui "github.com/readmedotmd/gui.md"
)

// ─── Layout Primitives ───

// Stack renders a vertical flex container (column direction).
// Gap values: "xs" (4px), "sm" (8px), "md" (16px), "lg" (24px), "xl" (32px), "none" (0).
//
// Data attributes:
//   - data-stack: gap size
func Stack(gap string, children ...gui.Node) gui.Node {
	if gap == "" {
		gap = "md"
	}
	return gui.Div(dataAttr("stack", gap))(children...)
}

// HStack renders a horizontal flex container (row direction).
// Gap values are the same as Stack.
//
// Data attributes:
//   - data-hstack: gap size
func HStack(gap string, children ...gui.Node) gui.Node {
	if gap == "" {
		gap = "md"
	}
	return gui.Div(dataAttr("hstack", gap))(children...)
}

// GridProps configures the Grid component.
type GridProps struct {
	Class   string
	Cols    string // "1"-"6", number of equal columns
	Gap     string // gap size (uses --core-space by default)
	Align   string // align-items value
	Justify string // justify-content value
}

// Grid renders a CSS grid container.
//
// Data attributes:
//   - data-grid: column count
//   - data-align: alignment
//   - data-justify: justification
func Grid(props GridProps, children ...gui.Node) gui.Node {
	cols := props.Cols
	if cols == "" {
		cols = "2"
	}
	attrs := collectAttrs(optClass(props.Class), dataAttr("grid", cols))
	if props.Align != "" {
		attrs = append(attrs, dataAttr("align", props.Align))
	}
	if props.Justify != "" {
		attrs = append(attrs, dataAttr("justify", props.Justify))
	}
	return gui.Div(attrs...)(children...)
}

// Center renders a flex container that centers its children both axes.
//
// Data attributes:
//   - data-center
func Center(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class), dataAttr("center", ""))...)(children...)
}

// Spacer renders an empty flex spacer element.
//
// Data attributes:
//   - data-spacer
func Spacer() gui.Node {
	return gui.Div(dataAttr("spacer", ""))()
}

// ─── Card ───

// CardProps configures the Card component.
type CardProps struct {
	Class   string
	Variant string // "", "surface", "flush"
}

// Card renders a bordered container.
//
// Data attributes:
//   - data-card: variant
func Card(props CardProps, children ...gui.Node) gui.Node {
	v := props.Variant
	if v == "" {
		v = "true"
	}
	return gui.Div(collectAttrs(optClass(props.Class), dataAttr("card", v))...)(children...)
}

// ─── Badge ───

// BadgeVariant identifies badge color schemes.
type BadgeVariant string

const (
	BadgeDefault BadgeVariant = ""
	BadgeAccent  BadgeVariant = "accent"
	BadgeSuccess BadgeVariant = "success"
	BadgeDanger  BadgeVariant = "danger"
	BadgeWarning BadgeVariant = "warning"
)

// Badge renders a small pill label.
//
// Data attributes:
//   - data-badge: variant
func Badge(class string, variant BadgeVariant, text string) gui.Node {
	v := string(variant)
	if v == "" {
		v = "true"
	}
	return gui.Span(collectAttrs(optClass(class), dataAttr("badge", v))...)(gui.Text(text))
}

// ─── Divider ───

// Divider renders a horizontal rule separator.
func Divider(class string) gui.Node {
	return gui.Hr(collectAttrs(optClass(class))...)()
}

// ─── Typography ───

// Heading renders an h1-h6 element. Level must be 1-6 (defaults to 2).
func Heading(level int, class string, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(class))
	switch level {
	case 1:
		return gui.H1(attrs...)(children...)
	case 3:
		return gui.H3(attrs...)(children...)
	case 4:
		return gui.H4(attrs...)(children...)
	case 5:
		return gui.H5(attrs...)(children...)
	case 6:
		return gui.H6(attrs...)(children...)
	default:
		return gui.H2(attrs...)(children...)
	}
}

// Paragraph renders a p element.
func Paragraph(class string, children ...gui.Node) gui.Node {
	return gui.P(collectAttrs(optClass(class))...)(children...)
}

// CodeBlock renders a pre>code block for displaying code.
func CodeBlock(class, content string) gui.Node {
	return gui.Pre(collectAttrs(optClass(class))...)(
		gui.Code()(gui.Text(content)),
	)
}

// InlineCode renders an inline code element.
func InlineCode(text string) gui.Node {
	return gui.Code()(gui.Text(text))
}

// Muted renders text in the muted color.
//
// Data attributes:
//   - data-muted
func Muted(text string) gui.Node {
	return gui.Span(dataAttr("muted", ""))(gui.Text(text))
}

// Mono renders text in the monospace font.
//
// Data attributes:
//   - data-mono
func Mono(text string) gui.Node {
	return gui.Span(dataAttr("mono", ""))(gui.Text(text))
}

// ─── Links ───

// LinkProps configures the Link component.
type LinkProps struct {
	Class  string
	Href   string
	Target string // "_blank", "_self", etc.
}

// Link renders an anchor element.
func Link(props LinkProps, children ...gui.Node) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.Href != "" {
		attrs = append(attrs, gui.Attr_("href", props.Href))
	}
	if props.Target != "" {
		attrs = append(attrs, gui.Attr_("target", props.Target))
		if props.Target == "_blank" {
			attrs = append(attrs, gui.Attr_("rel", "noopener noreferrer"))
		}
	}
	return gui.A(attrs...)(children...)
}

// ─── Images ───

// ImageProps configures the Image component.
type ImageProps struct {
	Class   string
	Src     string
	Alt     string
	Rounded bool // data-rounded
	Avatar  bool // data-avatar (circular)
}

// Image renders an img element.
//
// Data attributes:
//   - data-rounded: rounded corners
//   - data-avatar: circular
func Image(props ImageProps) gui.Node {
	attrs := collectAttrs(optClass(props.Class))
	if props.Src != "" {
		attrs = append(attrs, gui.Attr_("src", props.Src))
	}
	if props.Alt != "" {
		attrs = append(attrs, gui.Attr_("alt", props.Alt))
	}
	if props.Rounded {
		attrs = append(attrs, dataAttr("rounded", ""))
	}
	if props.Avatar {
		attrs = append(attrs, dataAttr("avatar", ""))
	}
	return gui.Img(attrs...)()
}

// ─── Lists ───

// UnorderedList renders a ul with li children.
func UnorderedList(class string, items ...gui.Node) gui.Node {
	return gui.Ul(collectAttrs(optClass(class))...)(items...)
}

// OrderedList renders an ol with li children.
func OrderedList(class string, items ...gui.Node) gui.Node {
	return gui.Ol(collectAttrs(optClass(class))...)(items...)
}

// ListItem renders a single li element.
func ListItem(children ...gui.Node) gui.Node {
	return gui.Li()(children...)
}

// ─── Blockquote ───

// Quote renders a blockquote element.
func Quote(class string, children ...gui.Node) gui.Node {
	return gui.Blockquote(collectAttrs(optClass(class))...)(children...)
}

// ─── Truncate ───

// Truncate wraps content with text truncation (ellipsis).
//
// Data attributes:
//   - data-truncate
func Truncate(class string, children ...gui.Node) gui.Node {
	return gui.Div(collectAttrs(optClass(class), dataAttr("truncate", ""))...)(children...)
}

// ─── Screen Reader Only ───

// SrOnly renders content only visible to screen readers.
//
// Data attributes:
//   - data-sr-only
func SrOnly(text string) gui.Node {
	return gui.Span(dataAttr("sr-only", ""))(gui.Text(text))
}
