package coremd

import (
	gui "github.com/readmedotmd/gui.md"
)

// ClassMap holds every CSS class name that varies between themes.
// Theme packages populate this struct with their own class constants.
type ClassMap struct {
	// Button
	Btn        string
	BtnPrimary string
	BtnDanger  string
	BtnSmall   string
	BtnToolbar string

	// Layout
	AppShell     string
	AppShellBody string
	AppShellMain string

	Navbar string

	Sidebar       string
	SidebarOpen   string
	SidebarHeader string

	Panel         string
	PanelExpanded string

	ModalBackdrop string
	Modal         string
	ModalBody     string
	ModalFooter   string

	DragHandle string

	// Navigation
	NavLink       string
	NavLinkActive string

	TabBar string

	BottomTabBar string

	// Data Display
	Message          string
	MessageUser      string
	MessageAssistant string
	MessageStreaming  string

	Thinking          string
	ThinkingCollapsib string

	ToolBadge string

	QuestionPrompt string

	StatusBadge       string
	StatusBadgePrefix string // e.g. "dbx-status-badge--" or "status-badge--"
	StatusDot         string
	StatusDotPrefix   string // e.g. "dbx-status-dot--" or "status-dot--"

	LabelBadge string

	UsageBadge string

	DiffViewer string

	DataTable string

	EmptyState string

	ClusterStats string

	// List
	ConversationItem       string
	ConversationItemActive string

	InstanceCard       string
	InstanceCardActive string

	ServiceRow string

	RunnerRow string

	FileTree string

	// Form
	FormGroup string

	TextInput      string
	TextInputError string

	Textarea        string
	TextareaAutoGrw string
	TextareaFixed   string

	Select string

	Checkbox string

	FeatureRow string

	VariableRow string

	ErrorMessage   string
	SuccessMessage string

	// Input
	ChatInput    string
	Autocomplete string
	MessageQueue string
	SearchInput  string

	// Overlay
	SearchOverlay string
	ContextMenu   string

	// Panel compositions
	SkillCard     string
	TerminalPanel string

	// Page
	LoginPage string

	// Utility
	Spinner      string
	SpinnerSmall string
	SpinnerLarge string
	Icon         string
}

// Themed wraps a ClassMap and provides methods for every themed component.
// Theme packages create a package-level instance populated with their class constants.
type Themed struct {
	ClassMap
}

// ─── Button ───

// Button renders a themed button.
func (t *Themed) Button(props ButtonProps, children ...gui.Node) gui.Node {
	cls := t.Btn
	switch props.Variant {
	case ButtonPrimary:
		cls += " " + t.BtnPrimary
	case ButtonDanger:
		cls += " " + t.BtnDanger
	case ButtonToolbar:
		cls += " " + t.BtnToolbar
	}
	if props.Size == ButtonSmall {
		cls += " " + t.BtnSmall
	}
	props.Class = cls
	return Button(props, children...)
}

// ─── Layout ───

// AppShell wraps content in the themed app-shell container.
func (t *Themed) AppShell(children ...gui.Node) gui.Node {
	return AppShell(t.ClassMap.AppShell, children...)
}

// AppShellBody wraps the themed flex body area.
func (t *Themed) AppShellBody(children ...gui.Node) gui.Node {
	return AppShellBody(t.ClassMap.AppShellBody, children...)
}

// AppShellMain wraps the themed main scrollable content area.
func (t *Themed) AppShellMain(children ...gui.Node) gui.Node {
	return AppShellMain(t.ClassMap.AppShellMain, children...)
}

// Navbar renders a themed top navigation bar.
func (t *Themed) Navbar(props NavbarProps, links ...gui.Node) gui.Node {
	props.Class = t.ClassMap.Navbar
	return Navbar(props, links...)
}

// Sidebar renders a themed side navigation panel.
func (t *Themed) Sidebar(props SidebarProps, header gui.Node, children ...gui.Node) gui.Node {
	props.Class = ClassIf(t.ClassMap.Sidebar, props.Open, t.SidebarOpen)
	return Sidebar(props, header, children...)
}

// SidebarHeader renders a themed sidebar header.
func (t *Themed) SidebarHeader(title string, actions ...gui.Node) gui.Node {
	return SidebarHeader(t.ClassMap.SidebarHeader, title, actions...)
}

// Panel renders a themed content panel.
func (t *Themed) Panel(props PanelProps, actions []gui.Node, children ...gui.Node) gui.Node {
	props.Class = ClassIf(t.ClassMap.Panel, props.Expanded, t.PanelExpanded)
	return Panel(props, actions, children...)
}

// ModalBackdrop renders a themed modal backdrop.
func (t *Themed) ModalBackdrop(children ...gui.Node) gui.Node {
	return ModalBackdrop(t.ClassMap.ModalBackdrop, children...)
}

// Modal renders a themed modal dialog.
func (t *Themed) Modal(title string, children ...gui.Node) gui.Node {
	return Modal(t.ClassMap.Modal, title, children...)
}

// ModalBody wraps content in a themed modal body.
func (t *Themed) ModalBody(children ...gui.Node) gui.Node {
	return ModalBody(t.ClassMap.ModalBody, children...)
}

// ModalFooter renders a themed modal footer.
func (t *Themed) ModalFooter(children ...gui.Node) gui.Node {
	return ModalFooter(t.ClassMap.ModalFooter, children...)
}

// DragHandle renders a themed drag handle indicator.
func (t *Themed) DragHandle() gui.Node {
	return DragHandle(t.ClassMap.DragHandle)
}

// ─── Navigation ───

// NavLink renders a themed navigation link.
func (t *Themed) NavLink(props NavLinkProps) gui.Node {
	props.Class = ClassIf(t.ClassMap.NavLink, props.Active, t.NavLinkActive)
	return NavLink(props)
}

// TabBar renders a themed horizontal tab bar.
func (t *Themed) TabBar(tabs []TabBarTab) gui.Node {
	return TabBar(t.ClassMap.TabBar, tabs)
}

// BottomTabBar renders a themed mobile bottom tab bar.
func (t *Themed) BottomTabBar(items []BottomTabItem) gui.Node {
	return BottomTabBar(t.ClassMap.BottomTabBar, items)
}

// ─── Data Display ───

// MessageBubble renders a themed chat message bubble.
func (t *Themed) MessageBubble(props MessageBubbleProps, children ...gui.Node) gui.Node {
	cls := t.Message
	if props.Role == "user" {
		cls += " " + t.MessageUser
	} else {
		cls += " " + t.MessageAssistant
	}
	if props.Streaming {
		cls += " " + t.MessageStreaming
	}
	props.Class = cls
	return MessageBubble(props, children...)
}

// ThinkingIndicator renders a themed spinner with a text label.
func (t *Themed) ThinkingIndicator(label string) gui.Node {
	return ThinkingIndicator(t.Thinking, label)
}

// ThinkingCollapsible renders a themed collapsible thinking section.
func (t *Themed) ThinkingCollapsible(label string, children ...gui.Node) gui.Node {
	return ThinkingCollapsible(Classes(t.Thinking, t.ThinkingCollapsib), label, children...)
}

// ToolBadge renders a themed tool badge pill.
func (t *Themed) ToolBadge(name string) gui.Node {
	return ToolBadge(t.ClassMap.ToolBadge, name)
}

// QuestionPrompt renders a themed prompt with selectable options.
func (t *Themed) QuestionPrompt(question string, options []QuestionPromptOption) gui.Node {
	return QuestionPrompt(t.ClassMap.QuestionPrompt, question, options)
}

// StatusBadge renders a themed colored status pill.
func (t *Themed) StatusBadge(status StatusBadgeStatus, label string) gui.Node {
	cls := t.ClassMap.StatusBadge + " " + t.StatusBadgePrefix + string(status)
	return StatusBadge(cls, status, label)
}

// StatusDot renders a themed small colored status dot.
func (t *Themed) StatusDot(status StatusBadgeStatus) gui.Node {
	cls := t.ClassMap.StatusDot + " " + t.StatusDotPrefix + string(status)
	return StatusDot(cls, status)
}

// LabelBadge renders a themed small label with an icon.
func (t *Themed) LabelBadge(icon, text string) gui.Node {
	return LabelBadge(t.ClassMap.LabelBadge, icon, text)
}

// UsageBadge renders themed CPU and memory usage indicators.
func (t *Themed) UsageBadge(cpu, memory string) gui.Node {
	return UsageBadge(t.ClassMap.UsageBadge, cpu, memory)
}

// DiffViewer renders a themed code diff view.
func (t *Themed) DiffViewer(lines []DiffLine) gui.Node {
	return DiffViewer(t.ClassMap.DiffViewer, lines)
}

// DataTable renders a themed data table.
func (t *Themed) DataTable(columns []string, rows [][]gui.Node) gui.Node {
	return DataTable(t.ClassMap.DataTable, columns, rows)
}

// EmptyState renders a themed empty state placeholder.
func (t *Themed) EmptyState(heading, description string) gui.Node {
	return EmptyState(t.ClassMap.EmptyState, heading, description)
}

// ClusterStatsBar renders a themed row of cluster statistics.
func (t *Themed) ClusterStatsBar(stats []ClusterStat, onClick func()) gui.Node {
	return ClusterStatsBar(t.ClassMap.ClusterStats, stats, onClick)
}

// ─── List ───

// ConversationItem renders a themed conversation list entry.
func (t *Themed) ConversationItem(props ConversationItemProps) gui.Node {
	props.Class = ClassIf(t.ClassMap.ConversationItem, props.Active, t.ConversationItemActive)
	return ConversationItem(props)
}

// InstanceCard renders a themed instance card.
func (t *Themed) InstanceCard(props InstanceCardProps) gui.Node {
	props.Class = ClassIf(t.ClassMap.InstanceCard, props.Active, t.InstanceCardActive)
	return InstanceCard(props)
}

// InstanceList renders a themed instance list.
func (t *Themed) InstanceList(props InstanceListProps, actions []gui.Node, children ...gui.Node) gui.Node {
	return InstanceList(props, actions, children...)
}

// ServiceRow renders a themed service row.
func (t *Themed) ServiceRow(props ServiceRowProps, actions ...gui.Node) gui.Node {
	props.Class = t.ClassMap.ServiceRow
	return ServiceRow(props, actions...)
}

// RunnerRow renders a themed runner row.
func (t *Themed) RunnerRow(props RunnerRowProps, actions ...gui.Node) gui.Node {
	props.Class = t.ClassMap.RunnerRow
	return RunnerRow(props, actions...)
}

// FileTree renders a themed file tree.
func (t *Themed) FileTree(items []FileTreeItem) gui.Node {
	return FileTree(t.ClassMap.FileTree, items)
}

// ─── Form ───

// FormGroup renders a themed form group.
func (t *Themed) FormGroup(label string, children ...gui.Node) gui.Node {
	return FormGroup(FormGroupProps{Class: t.ClassMap.FormGroup, Label: label}, children...)
}

// TextInput renders a themed text input field.
func (t *Themed) TextInput(props TextInputProps) gui.Node {
	props.Class = ClassIf(t.ClassMap.TextInput, props.Error, t.TextInputError)
	return TextInput(props)
}

// TextArea renders a themed textarea.
func (t *Themed) TextArea(props TextareaProps) gui.Node {
	cls := t.Textarea
	if props.AutoGrow {
		cls += " " + t.TextareaAutoGrw
	}
	if props.Fixed {
		cls += " " + t.TextareaFixed
	}
	props.Class = cls
	return TextArea(props)
}

// SelectInput renders a themed select dropdown.
func (t *Themed) SelectInput(props SelectProps, options ...gui.Node) gui.Node {
	props.Class = t.ClassMap.Select
	return SelectInput(props, options...)
}

// Checkbox renders a themed checkbox with a label.
func (t *Themed) Checkbox(props CheckboxProps) gui.Node {
	props.Class = t.ClassMap.Checkbox
	return Checkbox(props)
}

// FeatureRow renders a themed feature toggle row.
func (t *Themed) FeatureRow(props FeatureRowProps) gui.Node {
	props.Class = t.ClassMap.FeatureRow
	return FeatureRow(props)
}

// VariableRow renders a themed key-value variable row.
func (t *Themed) VariableRow(props VariableRowProps) gui.Node {
	props.Class = t.ClassMap.VariableRow
	return VariableRow(props)
}

// ErrorMessage renders a themed error message.
func (t *Themed) ErrorMessage(text string) gui.Node {
	return ErrorMessage(t.ClassMap.ErrorMessage, text)
}

// SuccessMessage renders a themed success message.
func (t *Themed) SuccessMessage(text string) gui.Node {
	return SuccessMessage(t.ClassMap.SuccessMessage, text)
}

// ─── Input ───

// ChatInput renders a themed chat message input area.
func (t *Themed) ChatInput(props ChatInputProps) gui.Node {
	props.Class = t.ClassMap.ChatInput
	return ChatInput(props)
}

// AutocompletePopup renders a themed autocomplete popup.
func (t *Themed) AutocompletePopup(items []AutocompleteItem, selectedIndex int) gui.Node {
	return AutocompletePopup(t.ClassMap.Autocomplete, items, selectedIndex)
}

// MessageQueue renders a themed list of queued messages.
func (t *Themed) MessageQueue(items []MessageQueueItem) gui.Node {
	return MessageQueue(t.ClassMap.MessageQueue, items)
}

// SearchInputField renders a themed search input field.
func (t *Themed) SearchInputField(placeholder string, onInput func(gui.Event)) gui.Node {
	return SearchInputField(t.ClassMap.SearchInput, placeholder, onInput)
}

// ─── Overlay ───

// SearchOverlay renders a themed full-screen search overlay.
func (t *Themed) SearchOverlay(tabs []TabBarTab, input gui.Node, results ...gui.Node) gui.Node {
	return SearchOverlay(t.ClassMap.SearchOverlay, tabs, input, results...)
}

// ContextMenu renders a themed positioned context menu.
func (t *Themed) ContextMenu(x, y int, items []ContextMenuItem) gui.Node {
	return ContextMenu(t.ClassMap.ContextMenu, x, y, items)
}

// ─── Panel Compositions ───

// ServicesPanel renders a themed services panel.
func (t *Themed) ServicesPanel(title string, headerActions []gui.Node, services ...gui.Node) gui.Node {
	return t.Panel(PanelProps{Title: title}, headerActions, services...)
}

// RunnerPanel renders a themed runner panel.
func (t *Themed) RunnerPanel(title string, runners ...gui.Node) gui.Node {
	return t.Panel(PanelProps{Title: title}, nil, runners...)
}

// GitPanel renders a themed git panel with tabs.
func (t *Themed) GitPanel(activeTab GitPanelTab, tabs []TabBarTab, content gui.Node) gui.Node {
	return t.Panel(PanelProps{Title: "Git"}, nil,
		t.TabBar(tabs),
		content,
	)
}

// SkillsPanel renders a themed skills panel.
func (t *Themed) SkillsPanel(skills []SkillCard) gui.Node {
	return SkillsPanel(t.ClassMap.SkillCard, skills)
}

// TerminalPanel renders a themed terminal panel.
func (t *Themed) TerminalPanel(tabs []TerminalTab, onAddTab func(), terminalContent gui.Node) gui.Node {
	return TerminalPanel(t.ClassMap.TerminalPanel, tabs, onAddTab, terminalContent)
}

// FileBrowser renders a themed file browser.
func (t *Themed) FileBrowser(heading string, items []FileTreeItem) gui.Node {
	return FileBrowser(heading, items)
}

// ─── Page ───

// LoginPage renders a themed centered login page.
func (t *Themed) LoginPage(title string, form gui.Node, errorMsg string) gui.Node {
	return LoginPage(t.ClassMap.LoginPage, title, form, errorMsg)
}

// SetupWizard renders a themed multi-step wizard.
func (t *Themed) SetupWizard(steps []SetupStep, content gui.Node) gui.Node {
	return SetupWizard("", steps, content)
}

// DashboardPage renders a themed dashboard page.
func (t *Themed) DashboardPage(heading, description string) gui.Node {
	return DashboardPage("p-5", heading, description)
}

// SettingsCard renders a themed settings section card.
func (t *Themed) SettingsCard(title string, children ...gui.Node) gui.Node {
	return SettingsCard("p-5", title, children...)
}

// ─── Utility ───

// Spinner renders a themed loading spinner.
func (t *Themed) Spinner(size SpinnerSize) gui.Node {
	cls := t.ClassMap.Spinner
	switch size {
	case SpinnerSmall:
		cls += " " + t.SpinnerSmall
	case SpinnerLarge:
		cls += " " + t.SpinnerLarge
	}
	return Spinner(SpinnerProps{Class: cls, Size: size})
}

// Icon renders a themed icon element.
func (t *Themed) Icon(class string) gui.Node {
	return Icon(t.ClassMap.Icon+" "+class, class)
}
