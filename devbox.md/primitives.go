package devboxmd

import (
	coremd "github.com/readmedotmd/core.md"
	gui "github.com/readmedotmd/gui.md"
)

// Re-export core types.
type (
	GridProps    = coremd.GridProps
	CardProps    = coremd.CardProps
	LinkProps    = coremd.LinkProps
	ImageProps   = coremd.ImageProps
	BadgeVariant = coremd.BadgeVariant
)

// Re-export badge variant constants.
const (
	BadgeDefault = coremd.BadgeDefault
	BadgeAccent  = coremd.BadgeAccent
	BadgeSuccess = coremd.BadgeSuccess
	BadgeDanger  = coremd.BadgeDanger
	BadgeWarning = coremd.BadgeWarning
)

// ─── Layout ───

// Stack renders a vertical flex container.
func Stack(gap string, children ...gui.Node) gui.Node {
	return coremd.Stack(gap, children...)
}

// HStack renders a horizontal flex container.
func HStack(gap string, children ...gui.Node) gui.Node {
	return coremd.HStack(gap, children...)
}

// Grid renders a CSS grid container.
func Grid(props GridProps, children ...gui.Node) gui.Node {
	return coremd.Grid(props, children...)
}

// Center renders a flex container that centers its children.
func Center(class string, children ...gui.Node) gui.Node {
	return coremd.Center(class, children...)
}

// Spacer renders a flex spacer.
func Spacer() gui.Node {
	return coremd.Spacer()
}

// ─── Card ───

// Card renders a bordered container.
func Card(props CardProps, children ...gui.Node) gui.Node {
	return coremd.Card(props, children...)
}

// ─── Badge ───

// Badge renders a small pill label.
func Badge(class string, variant BadgeVariant, text string) gui.Node {
	return coremd.Badge(class, variant, text)
}

// ─── Divider ───

// Divider renders a horizontal rule separator.
func Divider(class string) gui.Node {
	return coremd.Divider(class)
}

// ─── Typography ───

// Heading renders an h1-h6 element.
func Heading(level int, class string, children ...gui.Node) gui.Node {
	return coremd.Heading(level, class, children...)
}

// Paragraph renders a p element.
func Paragraph(class string, children ...gui.Node) gui.Node {
	return coremd.Paragraph(class, children...)
}

// CodeBlock renders a pre>code block.
func CodeBlock(class, content string) gui.Node {
	return coremd.CodeBlock(class, content)
}

// InlineCode renders an inline code element.
func InlineCode(text string) gui.Node {
	return coremd.InlineCode(text)
}

// Muted renders text in the muted color.
func Muted(text string) gui.Node {
	return coremd.Muted(text)
}

// Mono renders text in the monospace font.
func Mono(text string) gui.Node {
	return coremd.Mono(text)
}

// ─── Links ───

// Link renders an anchor element.
func Link(props LinkProps, children ...gui.Node) gui.Node {
	return coremd.Link(props, children...)
}

// ─── Images ───

// Image renders an img element.
func Image(props ImageProps) gui.Node {
	return coremd.Image(props)
}

// ─── Lists ───

// UnorderedList renders a ul with li children.
func UnorderedList(class string, items ...gui.Node) gui.Node {
	return coremd.UnorderedList(class, items...)
}

// OrderedList renders an ol with li children.
func OrderedList(class string, items ...gui.Node) gui.Node {
	return coremd.OrderedList(class, items...)
}

// ListItem renders a single li element.
func ListItem(children ...gui.Node) gui.Node {
	return coremd.ListItem(children...)
}

// ─── Blockquote ───

// Quote renders a blockquote element.
func Quote(class string, children ...gui.Node) gui.Node {
	return coremd.Quote(class, children...)
}

// ─── Utilities ───

// Truncate wraps content with text truncation.
func Truncate(class string, children ...gui.Node) gui.Node {
	return coremd.Truncate(class, children...)
}

// SrOnly renders content only visible to screen readers.
func SrOnly(text string) gui.Node {
	return coremd.SrOnly(text)
}
