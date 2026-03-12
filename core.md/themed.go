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

	DashboardLayout string
	SidebarCol      string
	SidebarColOpen  string
	SidebarOverlay  string
	CenterCol       string
	ChatArea        string
	ChatHeaderCls   string
	MsgList         string
	ChatInputArea   string
	ChatInputRow    string
	ChatInputWrap   string
	ChatInputWrapEx string

	// Navigation
	NavLink       string
	NavLinkActive string

	TabBar string

	BottomTabBar string

	ChatBackBtn    string
	HamburgerBtn   string
	ChatToolbar    string
	ToolbarBtn     string
	ToolbarBtnDgr  string

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

	MessageContent     string
	MessageContentUser string
	WorkingIndicator   string
	ChatStatusBadge    string
	ThinkingHistory    string
	ChatError          string
	AcceptPlanBar      string

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
	ChatInput        string
	Autocomplete     string
	MessageQueue     string
	SearchInput      string
	PastePreview     string
	ExpandBtn        string
	AttachBtn        string
	SendBtn          string
	CancelBtn        string
	ModeBtn          string
	MsgQueueBar      string
	QueuedItem       string
	AutocompleteHdr  string

	// Overlay
	SearchOverlay       string
	ContextMenu         string
	BottomSheet         string
	SearchCard          string
	SearchResult        string
	SearchResultContent string
	SearchSnippet       string

	// Panel compositions
	SkillCard        string
	TerminalPanel    string
	GitPanelCls      string
	GitSectionHdr    string
	GitFileList      string
	GitFile          string
	GitFileSelected  string
	GitCommitArea    string
	DiffCommentBtn   string
	DiffInlineComment string
	SvcActionBtn     string
	RunnerEmpty      string

	// Page
	LoginPage          string
	SettingsPage       string
	SettingsCardFull   string
	SettingsSection    string
	SettingsSubsection string
	SettingsForm       string
	SettingsFormActions string
	SettingsFormHelp   string
	SettingsCodeInput  string
	SettingsEnvRow     string
	SettingsFieldError string
	SettingsSchemaTable string
	AdminPage          string
	ClusterPageCls     string
	ClusterSummaryCard string
	ClusterSummaryRow  string

	// Utility
	Spinner      string
	SpinnerSmall string
	SpinnerLarge string
	Icon         string
	AppShellFull string
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

// DashboardLayout renders a themed flex row container.
func (t *Themed) DashboardLayout(children ...gui.Node) gui.Node {
	return DashboardLayout(t.ClassMap.DashboardLayout, children...)
}

// SidebarColumn renders a themed sidebar column wrapper.
func (t *Themed) SidebarColumn(open bool, children ...gui.Node) gui.Node {
	cls := ClassIf(t.ClassMap.SidebarCol, open, t.SidebarColOpen)
	return SidebarColumn(cls, open, children...)
}

// SidebarOverlay renders a themed semi-transparent overlay.
func (t *Themed) SidebarOverlay(onClick func()) gui.Node {
	return SidebarOverlay(t.ClassMap.SidebarOverlay, onClick)
}

// CenterColumn renders a themed flex:1 center column.
func (t *Themed) CenterColumn(children ...gui.Node) gui.Node {
	return CenterColumn(t.ClassMap.CenterCol, children...)
}

// ChatArea renders a themed chat area container.
func (t *Themed) ChatArea(children ...gui.Node) gui.Node {
	return ChatArea(t.ClassMap.ChatArea, children...)
}

// ChatHeader renders a themed chat header bar.
func (t *Themed) ChatHeader(title gui.Node, toolbar gui.Node) gui.Node {
	return ChatHeader(t.ClassMap.ChatHeaderCls, title, toolbar)
}

// MessageList renders a themed scrollable message list.
func (t *Themed) MessageList(children ...gui.Node) gui.Node {
	return MessageList(t.ClassMap.MsgList, children...)
}

// ChatInputArea renders a themed bottom-pinned input area.
func (t *Themed) ChatInputArea(children ...gui.Node) gui.Node {
	return ChatInputArea(t.ClassMap.ChatInputArea, children...)
}

// ChatInputRow renders a themed horizontal input row.
func (t *Themed) ChatInputRow(children ...gui.Node) gui.Node {
	return ChatInputRow(t.ClassMap.ChatInputRow, children...)
}

// ChatInputWrap renders a themed textarea wrapper.
func (t *Themed) ChatInputWrap(expanded bool, children ...gui.Node) gui.Node {
	cls := ClassIf(t.ClassMap.ChatInputWrap, expanded, t.ChatInputWrapEx)
	return ChatInputWrap(cls, expanded, children...)
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

// ChatBackButton renders a themed mobile back button.
func (t *Themed) ChatBackButton(onClick func()) gui.Node {
	return ChatBackButton(t.ClassMap.ChatBackBtn, onClick)
}

// HamburgerButton renders a themed mobile menu button.
func (t *Themed) HamburgerButton(onClick func()) gui.Node {
	return HamburgerButton(t.ClassMap.HamburgerBtn, onClick)
}

// ChatToolbar renders a themed chat toolbar.
func (t *Themed) ChatToolbar(desktop gui.Node, mobile gui.Node) gui.Node {
	return ChatToolbar(t.ClassMap.ChatToolbar, desktop, mobile)
}

// ToolbarButton renders a themed toolbar button.
func (t *Themed) ToolbarButton(icon, label string, danger bool, onClick func()) gui.Node {
	cls := ClassIf(t.ClassMap.ToolbarBtn, danger, t.ToolbarBtnDgr)
	return ToolbarButton(cls, icon, label, danger, onClick)
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
func (t *Themed) UsageBadge(cpu, memory string, onClick func()) gui.Node {
	return UsageBadge(t.ClassMap.UsageBadge, cpu, memory, onClick)
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

// MessageContent renders a themed message content wrapper.
func (t *Themed) MessageContent(role string, children ...gui.Node) gui.Node {
	cls := t.ClassMap.MessageContent
	if role == "user" && t.MessageContentUser != "" {
		cls += " " + t.MessageContentUser
	}
	return MessageContent(cls, role, children...)
}

// WorkingIndicator renders a themed pulsing working bar.
func (t *Themed) WorkingIndicator(label string) gui.Node {
	return WorkingIndicator(t.ClassMap.WorkingIndicator, label)
}

// ChatStatusBadge renders a themed streaming status badge.
func (t *Themed) ChatStatusBadge(label string) gui.Node {
	return ChatStatusBadge(t.ClassMap.ChatStatusBadge, label)
}

// ThinkingHistory renders a themed collapsible thinking history block.
func (t *Themed) ThinkingHistory(summary string, content gui.Node) gui.Node {
	return ThinkingHistory(t.ClassMap.ThinkingHistory, summary, content)
}

// ChatError renders a themed chat error message.
func (t *Themed) ChatError(message string) gui.Node {
	return ChatError(t.ClassMap.ChatError, message)
}

// AcceptPlanBar renders a themed accept plan bar.
func (t *Themed) AcceptPlanBar(onAccept func()) gui.Node {
	return AcceptPlanBar(t.ClassMap.AcceptPlanBar, onAccept)
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

// PastePreview renders a themed paste preview row.
func (t *Themed) PastePreview(items []PastePreviewItem) gui.Node {
	return PastePreview(t.ClassMap.PastePreview, items)
}

// ExpandButton renders a themed expand/collapse button.
func (t *Themed) ExpandButton(expanded bool, onToggle func()) gui.Node {
	return ExpandButton(t.ClassMap.ExpandBtn, expanded, onToggle)
}

// AttachButton renders a themed attach button.
func (t *Themed) AttachButton(onAttach func()) gui.Node {
	return AttachButton(t.ClassMap.AttachBtn, onAttach)
}

// SendButton renders a themed send button.
func (t *Themed) SendButton(label string, onClick func()) gui.Node {
	return SendButton(t.ClassMap.SendBtn, label, onClick)
}

// CancelButton renders a themed cancel button.
func (t *Themed) CancelButton(label string, onClick func()) gui.Node {
	return CancelButton(t.ClassMap.CancelBtn, label, onClick)
}

// ModeButton renders a themed mode toggle button.
func (t *Themed) ModeButton(mode string, onClick func()) gui.Node {
	return ModeButton(t.ClassMap.ModeBtn, mode, onClick)
}

// MessageQueueBar renders a themed queue bar.
func (t *Themed) MessageQueueBar(children ...gui.Node) gui.Node {
	return MessageQueueBar(t.ClassMap.MsgQueueBar, children...)
}

// QueuedItem renders a themed queued message row.
func (t *Themed) QueuedItem(text string, hasImage bool, onSend func(), onRemove func()) gui.Node {
	return QueuedItem(t.ClassMap.QueuedItem, text, hasImage, onSend, onRemove)
}

// AutocompleteHeader renders a themed autocomplete header.
func (t *Themed) AutocompleteHeader(trigger, label string) gui.Node {
	return AutocompleteHeader(t.ClassMap.AutocompleteHdr, trigger, label)
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

// BottomSheet renders a themed mobile bottom sheet.
func (t *Themed) BottomSheet(items []BottomSheetItem) gui.Node {
	return BottomSheet(t.ClassMap.BottomSheet, items)
}

// SearchOverlayCard renders a themed search overlay card.
func (t *Themed) SearchOverlayCard(tabs gui.Node, input gui.Node, results gui.Node) gui.Node {
	return SearchOverlayCard(t.ClassMap.SearchCard, tabs, input, results)
}

// SearchResult renders a themed search result row.
func (t *Themed) SearchResult(icon, path, text string, onAdd func()) gui.Node {
	return SearchResult(t.ClassMap.SearchResult, icon, path, text, onAdd)
}

// SearchResultContent renders a themed content search result.
func (t *Themed) SearchResultContent(path string, snippet gui.Node, onAdd func()) gui.Node {
	return SearchResultContent(t.ClassMap.SearchResultContent, path, snippet, onAdd)
}

// SearchSnippet renders a themed code snippet.
func (t *Themed) SearchSnippet(lines []SearchSnippetLine) gui.Node {
	return SearchSnippet(t.ClassMap.SearchSnippet, lines)
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

// GitPanel renders a themed git panel.
func (t *Themed) GitPanel(props GitPanelProps, headerActions []gui.Node, content gui.Node) gui.Node {
	props.Class = t.ClassMap.GitPanelCls
	return GitPanel(props, headerActions, content)
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

// GitSectionHeader renders a themed git section header.
func (t *Themed) GitSectionHeader(label string, staged bool) gui.Node {
	return GitSectionHeader(t.ClassMap.GitSectionHdr, label, staged)
}

// GitFileList renders a themed git file list.
func (t *Themed) GitFileList(children ...gui.Node) gui.Node {
	return GitFileList(t.ClassMap.GitFileList, children...)
}

// GitFile renders a themed git file entry.
func (t *Themed) GitFile(props GitFileProps) gui.Node {
	props.Class = ClassIf(t.ClassMap.GitFile, props.Selected, t.GitFileSelected)
	return GitFile(props)
}

// GitCommitArea renders a themed commit area.
func (t *Themed) GitCommitArea(input gui.Node, actions ...gui.Node) gui.Node {
	return GitCommitArea(t.ClassMap.GitCommitArea, input, actions...)
}

// DiffCommentButton renders a themed diff comment button.
func (t *Themed) DiffCommentButton(onClick func()) gui.Node {
	return DiffCommentButton(t.ClassMap.DiffCommentBtn, onClick)
}

// DiffInlineComment renders a themed diff inline comment.
func (t *Themed) DiffInlineComment(input gui.Node) gui.Node {
	return DiffInlineComment(t.ClassMap.DiffInlineComment, input)
}

// ServiceActionButton renders a themed service action button.
func (t *Themed) ServiceActionButton(icon, variant string, onClick func()) gui.Node {
	return ServiceActionButton(t.ClassMap.SvcActionBtn, icon, variant, onClick)
}

// RunnerPanelEmpty renders a themed runner panel empty state.
func (t *Themed) RunnerPanelEmpty(message string) gui.Node {
	return RunnerPanelEmpty(t.ClassMap.RunnerEmpty, message)
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

// SettingsPage renders a themed settings page container.
func (t *Themed) SettingsPage(children ...gui.Node) gui.Node {
	return SettingsPage(t.ClassMap.SettingsPage, children...)
}

// SettingsCardFull renders a themed settings card with colored header.
func (t *Themed) SettingsCardFull(icon, title string, children ...gui.Node) gui.Node {
	return SettingsCardFull(t.ClassMap.SettingsCardFull, icon, title, children...)
}

// SettingsSection renders a themed settings section.
func (t *Themed) SettingsSection(icon, title, description string, children ...gui.Node) gui.Node {
	return SettingsSection(t.ClassMap.SettingsSection, icon, title, description, children...)
}

// SettingsSubsection renders a themed settings subsection.
func (t *Themed) SettingsSubsection(icon, title, description string, children ...gui.Node) gui.Node {
	return SettingsSubsection(t.ClassMap.SettingsSubsection, icon, title, description, children...)
}

// SettingsForm renders a themed settings form.
func (t *Themed) SettingsForm(title gui.Node, children ...gui.Node) gui.Node {
	return SettingsForm(t.ClassMap.SettingsForm, title, children...)
}

// SettingsFormActions renders a themed settings form actions row.
func (t *Themed) SettingsFormActions(children ...gui.Node) gui.Node {
	return SettingsFormActions(t.ClassMap.SettingsFormActions, children...)
}

// SettingsFormHelp renders a themed settings form help text.
func (t *Themed) SettingsFormHelp(children ...gui.Node) gui.Node {
	return SettingsFormHelp(t.ClassMap.SettingsFormHelp, children...)
}

// SettingsCodeInput renders a themed code input.
func (t *Themed) SettingsCodeInput(props SettingsCodeInputProps) gui.Node {
	props.Class = t.ClassMap.SettingsCodeInput
	return SettingsCodeInput(props)
}

// SettingsEnvRow renders a themed environment row.
func (t *Themed) SettingsEnvRow(name string, badges []gui.Node, actions []gui.Node) gui.Node {
	return SettingsEnvRow(t.ClassMap.SettingsEnvRow, name, badges, actions)
}

// SettingsFieldError renders a themed field error.
func (t *Themed) SettingsFieldError(message string) gui.Node {
	return SettingsFieldError(t.ClassMap.SettingsFieldError, message)
}

// SettingsSchemaTable renders a themed schema table.
func (t *Themed) SettingsSchemaTable(rows []SettingsSchemaRow) gui.Node {
	return SettingsSchemaTable(t.ClassMap.SettingsSchemaTable, rows)
}

// AdminPage renders a themed admin page.
func (t *Themed) AdminPage(children ...gui.Node) gui.Node {
	return AdminPage(t.ClassMap.AdminPage, children...)
}

// ClusterPage renders a themed cluster page.
func (t *Themed) ClusterPage(children ...gui.Node) gui.Node {
	return ClusterPage(t.ClassMap.ClusterPageCls, children...)
}

// ClusterSummaryCard renders a themed cluster summary card.
func (t *Themed) ClusterSummaryCard(icon, value, label string) gui.Node {
	return ClusterSummaryCard(t.ClassMap.ClusterSummaryCard, icon, value, label)
}

// ClusterSummaryRow renders a themed cluster summary row.
func (t *Themed) ClusterSummaryRow(children ...gui.Node) gui.Node {
	return ClusterSummaryRow(t.ClassMap.ClusterSummaryRow, children...)
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

// AppShellFull renders a themed top-level app shell.
func (t *Themed) AppShellFull(scrollable bool, children ...gui.Node) gui.Node {
	return AppShellFull(t.ClassMap.AppShellFull, scrollable, children...)
}
