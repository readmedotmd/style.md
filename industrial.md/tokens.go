package industrialmd

import (
	coremd "github.com/readmedotmd/style.md/core.md"
)

// CSS class name constants for the industrial design system.
// These map directly to rules in styles.css.
const (
	// Button
	ClassBtn        = "btn"
	ClassBtnPrimary = "btn--primary"
	ClassBtnDanger  = "btn--danger"
	ClassBtnSmall   = "btn--small"
	ClassBtnToolbar = "btn--toolbar"

	// Layout
	ClassAppShell     = "app-shell"
	ClassAppShellBody = "app-shell__body"
	ClassAppShellMain = "app-shell__main"

	ClassNavbar      = "navbar"
	ClassNavbarBrand = "navbar__brand"
	ClassNavbarNav   = "navbar__nav"
	ClassNavbarStats = "navbar__stats"

	ClassSidebar       = "sidebar"
	ClassSidebarOpen   = "sidebar--open"
	ClassSidebarHeader = "sidebar__header"
	ClassSidebarTitle  = "sidebar__title"
	ClassSidebarList   = "sidebar__list"

	ClassPanel         = "panel"
	ClassPanelExpanded = "panel--expanded"
	ClassPanelHeader   = "panel__header"
	ClassPanelTitle    = "panel__title"
	ClassPanelActions  = "panel__actions"
	ClassPanelBody     = "panel__body"

	ClassModalBackdrop    = "modal-backdrop"
	ClassModal            = "modal"
	ClassModalHeader      = "modal__header"
	ClassModalHeaderTitle = "modal__header-title"
	ClassModalBody        = "modal__body"
	ClassModalFooter      = "modal__footer"

	ClassDragHandle = "drag-handle"

	ClassDashboardLayout = "dashboard"
	ClassSidebarCol      = "sidebar-col"
	ClassSidebarColOpen  = "sidebar-col--open"
	ClassSidebarOverlay  = "sidebar-overlay"
	ClassCenterCol       = "center-col"
	ClassChatArea        = "chat-area"
	ClassChatHeader      = "chat-header"
	ClassMsgList         = "message-list"
	ClassChatInputArea   = "chat-input-area"
	ClassChatInputRow    = "chat-input-row"
	ClassChatInputWrap   = "chat-input-wrap"
	ClassChatInputWrapEx = "chat-input-wrap--expanded"

	// Navigation
	ClassNavLink       = "nav-link"
	ClassNavLinkActive = "nav-link--active"
	ClassNavLinkIcon   = "nav-link__icon"

	ClassTabBar          = "tab-bar"
	ClassTabBarTab       = "tab-bar__tab"
	ClassTabBarTabActive = "tab-bar__tab--active"

	ClassBottomTabBar           = "bottom-tab-bar"
	ClassBottomTabBarItem       = "bottom-tab-bar__item"
	ClassBottomTabBarItemActive = "bottom-tab-bar__item--active"
	ClassBottomTabBarIcon       = "bottom-tab-bar__icon"

	ClassChatBackBtn   = "chat-back-btn"
	ClassHamburgerBtn  = "hamburger-btn"
	ClassChatToolbar   = "chat-toolbar"
	ClassToolbarBtn    = "chat-toolbar-btn"
	ClassToolbarBtnDgr = "chat-toolbar-btn--danger"

	// Data Display
	ClassMessage          = "message"
	ClassMessageUser      = "message--user"
	ClassMessageAssistant = "message--assistant"
	ClassMessageStreaming  = "message--streaming"
	ClassMessageBubble    = "message__bubble"

	ClassThinking          = "thinking"
	ClassThinkingCollapsib = "thinking--collapsible"

	ClassToolBadge        = "tool-badge"
	ClassToolBadgeSpinner = "tool-badge__spinner"

	ClassQuestionPrompt           = "question-prompt"
	ClassQuestionPromptText       = "question-prompt__text"
	ClassQuestionPromptOptions    = "question-prompt__options"
	ClassQuestionPromptOption     = "question-prompt__option"
	ClassQuestionPromptOptionLbl  = "question-prompt__option-label"
	ClassQuestionPromptOptionDesc = "question-prompt__option-desc"

	ClassStatusBadge = "status-badge"
	ClassStatusDot   = "status-dot"

	ClassLabelBadge = "label-badge"

	ClassUsageBadge          = "usage-badge"
	ClassUsageBadgeSeparator = "usage-badge__separator"

	ClassDiffViewer     = "diff-viewer"
	ClassDiffViewerLine = "diff-viewer__line"
	ClassDiffLineAdd    = "diff-viewer__line--add"
	ClassDiffLineRemove = "diff-viewer__line--remove"
	ClassDiffLineHeader = "diff-viewer__line--header"

	ClassDataTable = "data-table"

	ClassEmptyState     = "empty-state"
	ClassEmptyStateHead = "empty-state__heading"
	ClassEmptyStateDesc = "empty-state__description"

	ClassClusterStats      = "cluster-stats"
	ClassClusterStatsItem  = "cluster-stats__item"
	ClassClusterStatsValue = "cluster-stats__value"

	ClassMessageContent     = "message-content"
	ClassMessageContentUser = "message-content--user"
	ClassWorkingIndicator   = "working-indicator"
	ClassChatStatusBadge    = "chat-header-status"
	ClassThinkingHistory    = "thinking-history"
	ClassChatError          = "chat-error"
	ClassAcceptPlanBar      = "accept-plan-bar"

	// List
	ClassConversationItem       = "conversation-item"
	ClassConversationItemActive = "conversation-item--active"
	ClassConversationItemTitle  = "conversation-item__title"
	ClassConversationItemMeta   = "conversation-item__meta"

	ClassInstanceCard        = "instance-card"
	ClassInstanceCardActive  = "instance-card--active"
	ClassInstanceCardHeader  = "instance-card__header"
	ClassInstanceCardName    = "instance-card__name"
	ClassInstanceCardDone    = "instance-card__done-label"
	ClassInstanceCardRepo    = "instance-card__repo"
	ClassInstanceCardFooter  = "instance-card__footer"
	ClassInstanceCardSpinner = "instance-card__spinner"

	ClassInstanceListHeader  = "instance-list__header"
	ClassInstanceListTitle   = "instance-list__title"
	ClassInstanceListActions = "instance-list__actions"
	ClassInstanceListBody    = "instance-list__body"

	ClassServiceRow        = "service-row"
	ClassServiceRowStatus  = "service-row__status"
	ClassServiceRowInfo    = "service-row__info"
	ClassServiceRowName    = "service-row__name"
	ClassServiceRowImage   = "service-row__image"
	ClassServiceRowPorts   = "service-row__ports"
	ClassServiceRowActions = "service-row__actions"

	ClassRunnerRow        = "runner-row"
	ClassRunnerRowHeader  = "runner-row__header"
	ClassRunnerRowInfo    = "runner-row__info"
	ClassRunnerRowName    = "runner-row__name"
	ClassRunnerRowDesc    = "runner-row__description"
	ClassRunnerRowCount   = "runner-row__count"
	ClassRunnerRowActions = "runner-row__actions"
	ClassRunnerRowProcess = "runner-row__process"
	ClassRunnerRowProcTtl = "runner-row__process-title"

	ClassFileTree     = "file-tree"
	ClassFileTreeItem = "file-tree__item"
	ClassFileTreeDir  = "file-tree__item--dir"

	// Form
	ClassFormGroup      = "form-group"
	ClassFormGroupLabel = "form-group__label"

	ClassTextInput      = "text-input"
	ClassTextInputError = "text-input--error"

	ClassTextarea        = "textarea"
	ClassTextareaAutoGrw = "textarea--auto-grow"
	ClassTextareaFixed   = "textarea--fixed"

	ClassSelect = "select"

	ClassCheckbox      = "checkbox"
	ClassCheckboxInput = "checkbox__input"
	ClassCheckboxLabel = "checkbox__label"

	ClassFeatureRow     = "feature-row"
	ClassFeatureRowInfo = "feature-row__info"
	ClassFeatureRowName = "feature-row__name"
	ClassFeatureRowDesc = "feature-row__description"

	ClassVariableRow        = "variable-row"
	ClassVariableRowKey     = "variable-row__key"
	ClassVariableRowValue   = "variable-row__value"
	ClassVariableRowActions = "variable-row__actions"

	ClassErrorMessage   = "error-message"
	ClassSuccessMessage = "success-message"

	// Input
	ClassChatInput         = "chat-input"
	ClassChatInputWrapper  = "chat-input__wrapper"
	ClassChatInputTextarea = "chat-input__textarea"
	ClassChatInputToolbar  = "chat-input__toolbar"
	ClassChatInputSpacer   = "chat-input__toolbar-spacer"

	ClassAutocomplete         = "autocomplete"
	ClassAutocompleteItem     = "autocomplete__item"
	ClassAutocompleteSelected = "autocomplete__item--selected"
	ClassAutocompleteIcon     = "autocomplete__icon"
	ClassAutocompleteLabel    = "autocomplete__label"
	ClassAutocompleteDetail   = "autocomplete__detail"
	ClassAutocompleteSnippet  = "autocomplete__snippet"

	ClassMessageQueue        = "message-queue"
	ClassMessageQueueItem    = "message-queue__item"
	ClassMessageQueuePreview = "message-queue__preview"
	ClassMessageQueueImgTag  = "message-queue__image-tag"
	ClassMessageQueueActions = "message-queue__actions"

	ClassSearchInput = "search-input"

	ClassPastePreview    = "paste-preview"
	ClassExpandBtn       = "expand-btn"
	ClassAttachBtn       = "attach-btn"
	ClassSendBtn         = "send-btn"
	ClassCancelBtn       = "cancel-btn"
	ClassModeBtn         = "mode-btn"
	ClassMsgQueueBar     = "message-queue-bar"
	ClassQueuedItem      = "queued-item"
	ClassAutocompleteHdr = "ac-header"

	// Overlay
	ClassSearchOverlay        = "search-overlay"
	ClassSearchOverlayPanel   = "search-overlay__panel"
	ClassSearchOverlayResults = "search-overlay__results"

	ClassContextMenuBackdrop = "context-menu-backdrop"
	ClassContextMenu         = "context-menu"
	ClassContextMenuItem     = "context-menu__item"
	ClassContextMenuDanger   = "context-menu__item--danger"

	ClassBottomSheet         = "bottom-sheet-overlay"
	ClassSearchCard          = "search-card"
	ClassSearchResult        = "search-result"
	ClassSearchResultContent = "search-result search-result-content"
	ClassSearchSnippet       = "search-snippet"

	// Panel compositions
	ClassTerminalPanel       = "terminal-panel"
	ClassTerminalPanelTabs   = "terminal-panel__tabs"
	ClassTerminalPanelTab    = "terminal-panel__tab"
	ClassTerminalPanelActive = "terminal-panel__tab--active"
	ClassTerminalPanelClose  = "terminal-panel__tab-close"
	ClassTerminalPanelAdd    = "terminal-panel__add"
	ClassTerminalPanelBody   = "terminal-panel__body"

	ClassSkillCard     = "skill-card"
	ClassSkillCardName = "skill-card__name"
	ClassSkillCardDesc = "skill-card__description"

	ClassGitPanel          = "git-panel"
	ClassGitSectionHdr     = "git-section-header"
	ClassGitFileList       = "git-file-list"
	ClassGitFile           = "git-file"
	ClassGitFileSelected   = "git-file--selected"
	ClassGitCommitArea     = "git-commit-area"
	ClassDiffCommentBtn    = "diff-comment-btn"
	ClassDiffInlineComment = "diff-inline-comment"
	ClassSvcActionBtn      = "svc-action-btn"
	ClassRunnerEmpty       = "run-panel-empty"

	// Page
	ClassLoginPage      = "login-page"
	ClassLoginCard      = "login-card"
	ClassLoginCardTitle = "login-card__title"

	ClassSetupWizardSteps     = "setup-wizard__steps"
	ClassSetupWizardStep      = "setup-wizard__step"
	ClassSetupWizardActive    = "setup-wizard__step--active"
	ClassSetupWizardCompleted = "setup-wizard__step--completed"
	ClassSetupWizardNumber    = "setup-wizard__step-number"
	ClassSetupWizardConnector = "setup-wizard__step-connector"
	ClassSetupWizardBody      = "setup-wizard__body"

	ClassSettingsPage        = "settings-page"
	ClassSettingsCardFull    = "settings-card"
	ClassSettingsSection     = "settings-section-group"
	ClassSettingsSubsection  = "settings-subsection"
	ClassSettingsForm        = "settings-env-form"
	ClassSettingsFormActions  = "settings-env-form-actions"
	ClassSettingsFormHelp     = "settings-env-form-help"
	ClassSettingsCodeInput   = "settings-code-input"
	ClassSettingsEnvRow      = "settings-env-row"
	ClassSettingsFieldError  = "settings-field-error"
	ClassSettingsSchemaTable = "settings-schema"
	ClassAdminPage           = "users-page"
	ClassClusterPage         = "cluster-page"
	ClassClusterSummaryCard  = "cluster-summary-card"
	ClassClusterSummaryRow   = "cluster-summary"

	// Utility
	ClassSpinner      = "spinner"
	ClassSpinnerSmall = "spinner--small"
	ClassSpinnerLarge = "spinner--large"
	ClassIcon         = "icon"
	ClassAppShellFull = "app"
)

// classes joins multiple CSS class names into a single string.
func classes(names ...string) string {
	return coremd.Classes(names...)
}

// classIf returns the base class with the conditional class appended if condition is true.
func classIf(base string, condition bool, conditional string) string {
	return coremd.ClassIf(base, condition, conditional)
}
