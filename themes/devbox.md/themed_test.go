package devboxmd

import (
	"strings"
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestThemedComponents(t *testing.T) {
	tests := []struct {
		name  string
		node  gui.Node
		class string
	}{
		// Layout
		{"AppShell", AppShell(), "dbx-app-shell"},
		{"AppShellBody", AppShellBody(), "dbx-app-shell__body"},
		{"AppShellMain", AppShellMain(), "dbx-app-shell__main"},
		{"Navbar", Navbar(NavbarProps{}), "dbx-navbar"},
		{"Sidebar", Sidebar(SidebarProps{}, nil), "dbx-sidebar"},
		{"SidebarHeader", SidebarHeader("T"), "dbx-sidebar__header"},
		{"Panel", Panel(PanelProps{}, nil), "dbx-panel"},
		{"ModalBackdrop", ModalBackdrop(), "dbx-modal-backdrop"},
		{"Modal", Modal("T"), "dbx-modal"},
		{"ModalBody", ModalBody(), "dbx-modal__body"},
		{"ModalFooter", ModalFooter(), "dbx-modal__footer"},
		{"DragHandle", DragHandle(), "dbx-drag-handle"},
		{"DashboardLayout", DashboardLayout(), "dbx-dashboard"},
		{"SidebarColumn", SidebarColumn(false), "dbx-sidebar-col"},
		{"SidebarOverlay", SidebarOverlay(nil), "dbx-sidebar-overlay"},
		{"CenterColumn", CenterColumn(), "dbx-center-col"},
		{"ChatArea", ChatArea(), "dbx-chat-area"},
		{"ChatHeader", ChatHeader(nil, nil), "dbx-chat-header"},
		{"MessageList", MessageList(), "dbx-message-list"},
		{"ChatInputArea", ChatInputArea(), "dbx-chat-input-area"},
		{"ChatInputRow", ChatInputRow(), "dbx-chat-input-row"},
		{"ChatInputWrap", ChatInputWrap(false), "dbx-chat-input-wrap"},

		// Display
		{"MessageBubble", MessageBubble(MessageBubbleProps{}), "dbx-message"},
		{"ThinkingIndicator", ThinkingIndicator("x"), "dbx-thinking"},
		{"ThinkingCollapsible", ThinkingCollapsible("x"), "dbx-thinking--collapsible"},
		{"ToolBadge", ToolBadge("x"), "dbx-tool-badge"},
		{"QuestionPrompt", QuestionPrompt("q", nil), "dbx-question-prompt"},
		{"StatusBadge", StatusBadge(StatusRunning, "x"), "dbx-status-badge"},
		{"StatusDot", StatusDot(StatusRunning), "dbx-status-dot"},
		{"LabelBadge", LabelBadge("", "x"), "dbx-label-badge"},
		{"UsageBadge", UsageBadge("c", "m", nil), "dbx-usage-badge"},
		{"DiffViewer", DiffViewer(nil), "dbx-diff-viewer"},
		{"DataTable", DataTable(nil, nil), "dbx-data-table"},
		{"EmptyState", EmptyState("h", "d"), "dbx-empty-state"},
		{"ClusterStatsBar", ClusterStatsBar(nil, nil), "dbx-cluster-stats"},
		{"MessageContent", MessageContent("user"), "dbx-message-content"},
		{"WorkingIndicator", WorkingIndicator("x"), "dbx-working-indicator"},
		{"ChatStatusBadge", ChatStatusBadge("x"), "dbx-chat-header-status"},
		{"ThinkingHistory", ThinkingHistory("s", nil), "dbx-thinking-history"},
		{"ChatError", ChatError("e"), "dbx-chat-error"},
		{"AcceptPlanBar", AcceptPlanBar(nil), "dbx-accept-plan-bar"},

		// Navigation
		{"ChatBackButton", ChatBackButton(nil), "dbx-chat-back-btn"},
		{"HamburgerButton", HamburgerButton(nil), "dbx-hamburger-btn"},
		{"ChatToolbar", ChatToolbar(nil, nil), "dbx-chat-toolbar"},
		{"ToolbarButton", ToolbarButton("", "x", false, nil), "dbx-chat-toolbar-btn"},

		// Input
		{"PastePreview", PastePreview(nil), "dbx-paste-preview"},
		{"ExpandButton", ExpandButton(false, nil), "dbx-expand-btn"},
		{"AttachButton", AttachButton(nil), "dbx-attach-btn"},
		{"SendButton", SendButton("S", nil), "dbx-send-btn"},
		{"CancelButton", CancelButton("C", nil), "dbx-cancel-btn"},
		{"ModeButton", ModeButton("act", nil), "dbx-mode-btn"},
		{"MessageQueueBar", MessageQueueBar(), "dbx-message-queue-bar"},
		{"QueuedItem", QueuedItem("t", false, nil, nil), "dbx-queued-item"},
		{"AutocompleteHeader", AutocompleteHeader("/", "cmds"), "dbx-ac-header"},

		// Overlay
		{"BottomSheet", BottomSheet(nil), "dbx-bottom-sheet-overlay"},
		{"SearchOverlayCard", SearchOverlayCard(nil, nil, nil), "dbx-search-card"},
		{"SearchResult", SearchResult("", "p", "t", nil), "dbx-search-result"},
		{"SearchResultContent", SearchResultContent("p", nil, nil), "dbx-search-result dbx-search-result-content"},
		{"SearchSnippet", SearchSnippet(nil), "dbx-search-snippet"},

		// Panel
		{"GitPanel", GitPanel(GitPanelProps{}, nil, nil), "dbx-git-panel"},
		{"GitSectionHeader", GitSectionHeader("l", false), "dbx-git-section-header"},
		{"GitFileList", GitFileList(), "dbx-git-file-list"},
		{"GitFile", GitFile(GitFileProps{}), "dbx-git-file"},
		{"GitCommitArea", GitCommitArea(nil), "dbx-git-commit-area"},
		{"DiffCommentButton", DiffCommentButton(nil), "dbx-diff-comment-btn"},
		{"DiffInlineComment", DiffInlineComment(nil), "dbx-diff-inline-comment"},
		{"ServiceActionButton", ServiceActionButton("", "stop", nil), "dbx-svc-action-btn"},
		{"RunnerPanelEmpty", RunnerPanelEmpty("msg"), "dbx-run-panel-empty"},

		// Page
		{"SettingsPage", SettingsPage(), "dbx-settings-page"},
		{"SettingsCardFull", SettingsCardFull("", "T"), "dbx-settings-card"},
		{"SettingsSection", SettingsSection("", "T", ""), "dbx-settings-section-group"},
		{"SettingsSubsection", SettingsSubsection("", "T", ""), "dbx-settings-subsection"},
		{"SettingsForm", SettingsForm(nil), "dbx-settings-env-form"},
		{"SettingsFormActions", SettingsFormActions(), "dbx-settings-env-form-actions"},
		{"SettingsFormHelp", SettingsFormHelp(), "dbx-settings-env-form-help"},
		{"SettingsCodeInput", SettingsCodeInput(SettingsCodeInputProps{}), "dbx-settings-code-input"},
		{"SettingsEnvRow", SettingsEnvRow("n", nil, nil), "dbx-settings-env-row"},
		{"SettingsFieldError", SettingsFieldError("e"), "dbx-settings-field-error"},
		{"SettingsSchemaTable", SettingsSchemaTable(nil), "dbx-settings-schema"},
		{"AdminPage", AdminPage(), "dbx-users-page"},
		{"ClusterPage", ClusterPage(), "dbx-cluster-page"},
		{"ClusterSummaryCard", ClusterSummaryCard("", "v", "l"), "dbx-cluster-summary-card"},
		{"ClusterSummaryRow", ClusterSummaryRow(), "dbx-cluster-summary"},

		// Utility
		{"AppShellFull", AppShellFull(false), "dbx-app"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			screen := guitesting.Render(tt.node)
			html := screen.HTML()
			// For class strings that contain multiple space-separated classes,
			// verify each individual class is present in the rendered HTML.
			for _, cls := range strings.Fields(tt.class) {
				if !strings.Contains(html, cls) {
					t.Errorf("expected class %q in rendered HTML:\n%s", cls, html)
				}
			}
		})
	}
}
