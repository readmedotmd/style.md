package industrialmd

import (
	"strings"
	"testing"

	gui "github.com/readmedotmd/gui.md"
	guitesting "github.com/readmedotmd/gui.md/testing"
)

func TestThemedWrappers(t *testing.T) {
	noop := func() {}

	cases := []struct {
		name  string
		node  gui.Node
		class string
	}{
		// Layout
		{"AppShell", AppShell(gui.Text("x")), "app-shell"},
		{"AppShellBody", AppShellBody(gui.Text("x")), "app-shell__body"},
		{"AppShellMain", AppShellMain(gui.Text("x")), "app-shell__main"},
		{"Navbar", Navbar(NavbarProps{Brand: "b"}, gui.Text("link")), "navbar"},
		{"Sidebar", Sidebar(SidebarProps{}, gui.Text("hdr"), gui.Text("x")), "sidebar"},
		{"SidebarHeader", SidebarHeader("Title"), "sidebar__header"},
		{"Panel", Panel(PanelProps{Title: "p"}, nil, gui.Text("x")), "panel"},
		{"ModalBackdrop", ModalBackdrop(gui.Text("x")), "modal-backdrop"},
		{"Modal", Modal("title", gui.Text("x")), "modal"},
		{"ModalBody", ModalBody(gui.Text("x")), "modal__body"},
		{"ModalFooter", ModalFooter(gui.Text("x")), "modal__footer"},
		{"DragHandle", DragHandle(), "drag-handle"},
		{"DashboardLayout", DashboardLayout(gui.Text("x")), "dashboard"},
		{"SidebarColumn", SidebarColumn(false, gui.Text("x")), "sidebar-col"},
		{"SidebarOverlay", SidebarOverlay(noop), "sidebar-overlay"},
		{"CenterColumn", CenterColumn(gui.Text("x")), "center-col"},
		{"ChatArea", ChatArea(gui.Text("x")), "chat-area"},
		{"ChatHeader", ChatHeader(gui.Text("t"), gui.Text("tb")), "chat-header"},
		{"MessageList", MessageList(gui.Text("x")), "message-list"},
		{"ChatInputArea", ChatInputArea(gui.Text("x")), "chat-input-area"},
		{"ChatInputRow", ChatInputRow(gui.Text("x")), "chat-input-row"},
		{"ChatInputWrap", ChatInputWrap(false, gui.Text("x")), "chat-input-wrap"},

		// Display
		{"MessageBubble", MessageBubble(MessageBubbleProps{}, gui.Text("x")), "message"},
		{"ThinkingIndicator", ThinkingIndicator("thinking"), "thinking"},
		{"ThinkingCollapsible", ThinkingCollapsible("label", gui.Text("x")), "thinking--collapsible"},
		{"ToolBadge", ToolBadge("tool"), "tool-badge"},
		{"QuestionPrompt", QuestionPrompt("q?", []QuestionPromptOption{{Label: "a", OnClick: noop}}), "question-prompt"},
		{"StatusBadge", StatusBadge(StatusRunning, "ok"), "status-badge"},
		{"StatusDot", StatusDot(StatusRunning), "status-dot"},
		{"LabelBadge", LabelBadge("icon", "text"), "label-badge"},
		{"UsageBadge", UsageBadge("10%", "20%", noop), "usage-badge"},
		{"DiffViewer", DiffViewer([]DiffLine{{Content: "line"}}), "diff-viewer"},
		{"DataTable", DataTable([]string{"col"}, [][]gui.Node{{gui.Text("v")}}), "data-table"},
		{"EmptyState", EmptyState("title", "desc"), "empty-state"},
		{"ClusterStatsBar", ClusterStatsBar([]ClusterStat{{Label: "cpu", Value: "1"}}, noop), "cluster-stats"},
		{"MessageContent", MessageContent("assistant", gui.Text("x")), "message-content"},
		{"WorkingIndicator", WorkingIndicator("working"), "working-indicator"},
		{"ChatStatusBadge", ChatStatusBadge("status"), "chat-header-status"},
		{"ThinkingHistory", ThinkingHistory("summary", gui.Text("x")), "thinking-history"},
		{"ChatError", ChatError("err"), "chat-error"},
		{"AcceptPlanBar", AcceptPlanBar(noop), "accept-plan-bar"},

		// Navigation
		{"ChatBackButton", ChatBackButton(noop), "chat-back-btn"},
		{"HamburgerButton", HamburgerButton(noop), "hamburger-btn"},
		{"ChatToolbar", ChatToolbar(gui.Text("d"), gui.Text("m")), "chat-toolbar"},
		{"ToolbarButton", ToolbarButton("icon", "label", false, noop), "chat-toolbar-btn"},

		// Input
		{"PastePreview", PastePreview([]PastePreviewItem{{Src: "f"}}), "paste-preview"},
		{"ExpandButton", ExpandButton(false, noop), "expand-btn"},
		{"AttachButton", AttachButton(noop), "attach-btn"},
		{"SendButton", SendButton("send", noop), "send-btn"},
		{"CancelButton", CancelButton("cancel", noop), "cancel-btn"},
		{"ModeButton", ModeButton("mode", noop), "mode-btn"},
		{"MessageQueueBar", MessageQueueBar(gui.Text("x")), "message-queue-bar"},
		{"QueuedItem", QueuedItem("text", false, noop, noop), "queued-item"},
		{"AutocompleteHeader", AutocompleteHeader("/", "commands"), "ac-header"},

		// Overlay
		{"BottomSheet", BottomSheet([]BottomSheetItem{{Label: "x", OnClick: noop}}), "bottom-sheet-overlay"},
		{"SearchOverlayCard", SearchOverlayCard(gui.Text("t"), gui.Text("i"), gui.Text("r")), "search-card"},
		{"SearchResult", SearchResult("icon", "path", "text", noop), "search-result"},
		{"SearchResultContent", SearchResultContent("path", gui.Text("snip"), noop), "search-result-content"},
		{"SearchSnippet", SearchSnippet([]SearchSnippetLine{{Text: "x"}}), "search-snippet"},

		// Panel
		{"GitPanel", GitPanel(GitPanelProps{Branch: "main"}, nil, gui.Text("x")), "git-panel"},
		{"GitSectionHeader", GitSectionHeader("staged", true), "git-section-header"},
		{"GitFileList", GitFileList(gui.Text("x")), "git-file-list"},
		{"GitFile", GitFile(GitFileProps{Path: "f.go", State: "M"}), "git-file"},
		{"GitCommitArea", GitCommitArea(gui.Text("input")), "git-commit-area"},
		{"DiffCommentButton", DiffCommentButton(noop), "diff-comment-btn"},
		{"DiffInlineComment", DiffInlineComment(gui.Text("x")), "diff-inline-comment"},
		{"ServiceActionButton", ServiceActionButton("icon", "start", noop), "svc-action-btn"},
		{"RunnerPanelEmpty", RunnerPanelEmpty("no runners"), "run-panel-empty"},

		// Page
		{"SettingsPage", SettingsPage(gui.Text("x")), "settings-page"},
		{"SettingsCardFull", SettingsCardFull("icon", "title", gui.Text("x")), "settings-card"},
		{"SettingsSection", SettingsSection("icon", "title", "desc", gui.Text("x")), "settings-section-group"},
		{"SettingsSubsection", SettingsSubsection("icon", "title", "desc", gui.Text("x")), "settings-subsection"},
		{"SettingsForm", SettingsForm(gui.Text("title"), gui.Text("x")), "settings-env-form"},
		{"SettingsFormActions", SettingsFormActions(gui.Text("x")), "settings-env-form-actions"},
		{"SettingsFormHelp", SettingsFormHelp(gui.Text("x")), "settings-env-form-help"},
		{"SettingsCodeInput", SettingsCodeInput(SettingsCodeInputProps{Value: "code"}), "settings-code-input"},
		{"SettingsEnvRow", SettingsEnvRow("ENV_VAR", nil, nil), "settings-env-row"},
		{"SettingsFieldError", SettingsFieldError("error"), "settings-field-error"},
		{"SettingsSchemaTable", SettingsSchemaTable([]SettingsSchemaRow{{Type: "string", Description: "a field"}}), "settings-schema"},
		{"AdminPage", AdminPage(gui.Text("x")), "users-page"},
		{"ClusterPage", ClusterPage(gui.Text("x")), "cluster-page"},
		{"ClusterSummaryCard", ClusterSummaryCard("icon", "42", "nodes"), "cluster-summary-card"},
		{"ClusterSummaryRow", ClusterSummaryRow(gui.Text("x")), "cluster-summary"},

		// Utility
		{"AppShellFull", AppShellFull(false, gui.Text("x")), "app"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			screen := guitesting.Render(tc.node)
			html := screen.HTML()
			if !strings.Contains(html, tc.class) {
				t.Errorf("expected HTML to contain class %q, got:\n%s", tc.class, html)
			}
		})
	}
}
