package devboxmd

import (
	coremd "github.com/readmedotmd/style.md/core.md"
)

// CSS class name constants for the Devbox design system.
// These map directly to rules in styles.css.
const (
	// Button
	ClassBtn        = "dbx-btn"
	ClassBtnPrimary = "dbx-btn--primary"
	ClassBtnDanger  = "dbx-btn--danger"
	ClassBtnSmall   = "dbx-btn--small"
	ClassBtnToolbar = "dbx-btn--toolbar"

	// Layout
	ClassAppShell     = "dbx-app-shell"
	ClassAppShellBody = "dbx-app-shell__body"
	ClassAppShellMain = "dbx-app-shell__main"

	ClassNavbar      = "dbx-navbar"
	ClassNavbarBrand = "dbx-navbar__brand"
	ClassNavbarNav   = "dbx-navbar__nav"
	ClassNavbarStats = "dbx-navbar__stats"

	ClassSidebar       = "dbx-sidebar"
	ClassSidebarOpen   = "dbx-sidebar--open"
	ClassSidebarHeader = "dbx-sidebar__header"
	ClassSidebarTitle  = "dbx-sidebar__title"
	ClassSidebarList   = "dbx-sidebar__list"

	ClassPanel         = "dbx-panel"
	ClassPanelExpanded = "dbx-panel--expanded"
	ClassPanelHeader   = "dbx-panel__header"
	ClassPanelTitle    = "dbx-panel__title"
	ClassPanelActions  = "dbx-panel__actions"
	ClassPanelBody     = "dbx-panel__body"

	ClassModalBackdrop    = "dbx-modal-backdrop"
	ClassModal            = "dbx-modal"
	ClassModalHeader      = "dbx-modal__header"
	ClassModalHeaderTitle = "dbx-modal__header-title"
	ClassModalBody        = "dbx-modal__body"
	ClassModalFooter      = "dbx-modal__footer"

	ClassDragHandle = "dbx-drag-handle"

	// Navigation
	ClassNavLink       = "dbx-nav-link"
	ClassNavLinkActive = "dbx-nav-link--active"
	ClassNavLinkIcon   = "dbx-nav-link__icon"

	ClassTabBar          = "dbx-tab-bar"
	ClassTabBarTab       = "dbx-tab-bar__tab"
	ClassTabBarTabActive = "dbx-tab-bar__tab--active"

	ClassBottomTabBar           = "dbx-bottom-tab-bar"
	ClassBottomTabBarItem       = "dbx-bottom-tab-bar__item"
	ClassBottomTabBarItemActive = "dbx-bottom-tab-bar__item--active"
	ClassBottomTabBarIcon       = "dbx-bottom-tab-bar__icon"

	// Data Display
	ClassMessage          = "dbx-message"
	ClassMessageUser      = "dbx-message--user"
	ClassMessageAssistant = "dbx-message--assistant"
	ClassMessageStreaming = "dbx-message--streaming"
	ClassMessageBubble    = "dbx-message__bubble"

	ClassThinking          = "dbx-thinking"
	ClassThinkingCollapsib = "dbx-thinking--collapsible"

	ClassToolBadge        = "dbx-tool-badge"
	ClassToolBadgeSpinner = "dbx-tool-badge__spinner"

	ClassQuestionPrompt           = "dbx-question-prompt"
	ClassQuestionPromptText       = "dbx-question-prompt__text"
	ClassQuestionPromptOptions    = "dbx-question-prompt__options"
	ClassQuestionPromptOption     = "dbx-question-prompt__option"
	ClassQuestionPromptOptionLbl  = "dbx-question-prompt__option-label"
	ClassQuestionPromptOptionDesc = "dbx-question-prompt__option-desc"

	ClassStatusBadge = "dbx-status-badge"
	ClassStatusDot   = "dbx-status-dot"

	ClassLabelBadge = "dbx-label-badge"

	ClassUsageBadge          = "dbx-usage-badge"
	ClassUsageBadgeSeparator = "dbx-usage-badge__separator"

	ClassDiffViewer     = "dbx-diff-viewer"
	ClassDiffViewerLine = "dbx-diff-viewer__line"
	ClassDiffLineAdd    = "dbx-diff-viewer__line--add"
	ClassDiffLineRemove = "dbx-diff-viewer__line--remove"
	ClassDiffLineHeader = "dbx-diff-viewer__line--header"

	ClassDataTable = "dbx-data-table"

	ClassEmptyState     = "dbx-empty-state"
	ClassEmptyStateHead = "dbx-empty-state__heading"
	ClassEmptyStateDesc = "dbx-empty-state__description"

	ClassClusterStats      = "dbx-cluster-stats"
	ClassClusterStatsItem  = "dbx-cluster-stats__item"
	ClassClusterStatsValue = "dbx-cluster-stats__value"

	// List
	ClassConversationItem       = "dbx-conversation-item"
	ClassConversationItemActive = "dbx-conversation-item--active"
	ClassConversationItemTitle  = "dbx-conversation-item__title"
	ClassConversationItemMeta   = "dbx-conversation-item__meta"

	ClassInstanceCard        = "dbx-instance-card"
	ClassInstanceCardActive  = "dbx-instance-card--active"
	ClassInstanceCardHeader  = "dbx-instance-card__header"
	ClassInstanceCardName    = "dbx-instance-card__name"
	ClassInstanceCardDone    = "dbx-instance-card__done-label"
	ClassInstanceCardRepo    = "dbx-instance-card__repo"
	ClassInstanceCardFooter  = "dbx-instance-card__footer"
	ClassInstanceCardSpinner = "dbx-instance-card__spinner"

	ClassInstanceListHeader  = "dbx-instance-list__header"
	ClassInstanceListTitle   = "dbx-instance-list__title"
	ClassInstanceListActions = "dbx-instance-list__actions"
	ClassInstanceListBody    = "dbx-instance-list__body"

	ClassServiceRow        = "dbx-service-row"
	ClassServiceRowStatus  = "dbx-service-row__status"
	ClassServiceRowInfo    = "dbx-service-row__info"
	ClassServiceRowName    = "dbx-service-row__name"
	ClassServiceRowImage   = "dbx-service-row__image"
	ClassServiceRowPorts   = "dbx-service-row__ports"
	ClassServiceRowActions = "dbx-service-row__actions"

	ClassRunnerRow        = "dbx-runner-row"
	ClassRunnerRowHeader  = "dbx-runner-row__header"
	ClassRunnerRowInfo    = "dbx-runner-row__info"
	ClassRunnerRowName    = "dbx-runner-row__name"
	ClassRunnerRowDesc    = "dbx-runner-row__description"
	ClassRunnerRowCount   = "dbx-runner-row__count"
	ClassRunnerRowActions = "dbx-runner-row__actions"
	ClassRunnerRowProcess = "dbx-runner-row__process"
	ClassRunnerRowProcTtl = "dbx-runner-row__process-title"

	ClassFileTree     = "dbx-file-tree"
	ClassFileTreeItem = "dbx-file-tree__item"
	ClassFileTreeDir  = "dbx-file-tree__item--dir"

	// Form
	ClassFormGroup      = "dbx-form-group"
	ClassFormGroupLabel = "dbx-form-group__label"

	ClassTextInput      = "dbx-text-input"
	ClassTextInputError = "dbx-text-input--error"

	ClassTextarea        = "dbx-textarea"
	ClassTextareaAutoGrw = "dbx-textarea--auto-grow"
	ClassTextareaFixed   = "dbx-textarea--fixed"

	ClassSelect = "dbx-select"

	ClassCheckbox      = "dbx-checkbox"
	ClassCheckboxInput = "dbx-checkbox__input"
	ClassCheckboxLabel = "dbx-checkbox__label"

	ClassFeatureRow     = "dbx-feature-row"
	ClassFeatureRowInfo = "dbx-feature-row__info"
	ClassFeatureRowName = "dbx-feature-row__name"
	ClassFeatureRowDesc = "dbx-feature-row__description"

	ClassVariableRow        = "dbx-variable-row"
	ClassVariableRowKey     = "dbx-variable-row__key"
	ClassVariableRowValue   = "dbx-variable-row__value"
	ClassVariableRowActions = "dbx-variable-row__actions"

	ClassErrorMessage   = "dbx-error-message"
	ClassSuccessMessage = "dbx-success-message"

	// Input
	ClassChatInput         = "dbx-chat-input"
	ClassChatInputWrapper  = "dbx-chat-input__wrapper"
	ClassChatInputTextarea = "dbx-chat-input__textarea"
	ClassChatInputToolbar  = "dbx-chat-input__toolbar"
	ClassChatInputSpacer   = "dbx-chat-input__toolbar-spacer"

	ClassAutocomplete         = "dbx-autocomplete"
	ClassAutocompleteItem     = "dbx-autocomplete__item"
	ClassAutocompleteSelected = "dbx-autocomplete__item--selected"
	ClassAutocompleteIcon     = "dbx-autocomplete__icon"
	ClassAutocompleteLabel    = "dbx-autocomplete__label"
	ClassAutocompleteDetail   = "dbx-autocomplete__detail"
	ClassAutocompleteSnippet  = "dbx-autocomplete__snippet"

	ClassMessageQueue        = "dbx-message-queue"
	ClassMessageQueueItem    = "dbx-message-queue__item"
	ClassMessageQueuePreview = "dbx-message-queue__preview"
	ClassMessageQueueImgTag  = "dbx-message-queue__image-tag"
	ClassMessageQueueActions = "dbx-message-queue__actions"

	ClassSearchInput = "dbx-search-input"

	// Overlay
	ClassSearchOverlay        = "dbx-search-overlay"
	ClassSearchOverlayPanel   = "dbx-search-overlay__panel"
	ClassSearchOverlayResults = "dbx-search-overlay__results"

	ClassContextMenuBackdrop = "dbx-context-menu-backdrop"
	ClassContextMenu         = "dbx-context-menu"
	ClassContextMenuItem     = "dbx-context-menu__item"
	ClassContextMenuDanger   = "dbx-context-menu__item--danger"

	// Panel compositions
	ClassTerminalPanel       = "dbx-terminal-panel"
	ClassTerminalPanelTabs   = "dbx-terminal-panel__tabs"
	ClassTerminalPanelTab    = "dbx-terminal-panel__tab"
	ClassTerminalPanelActive = "dbx-terminal-panel__tab--active"
	ClassTerminalPanelClose  = "dbx-terminal-panel__tab-close"
	ClassTerminalPanelAdd    = "dbx-terminal-panel__add"
	ClassTerminalPanelBody   = "dbx-terminal-panel__body"

	ClassSkillCard     = "dbx-skill-card"
	ClassSkillCardName = "dbx-skill-card__name"
	ClassSkillCardDesc = "dbx-skill-card__description"

	// Page
	ClassLoginPage      = "dbx-login-page"
	ClassLoginCard      = "dbx-login-card"
	ClassLoginCardTitle = "dbx-login-card__title"

	ClassSetupWizardSteps     = "dbx-setup-wizard__steps"
	ClassSetupWizardStep      = "dbx-setup-wizard__step"
	ClassSetupWizardActive    = "dbx-setup-wizard__step--active"
	ClassSetupWizardCompleted = "dbx-setup-wizard__step--completed"
	ClassSetupWizardNumber    = "dbx-setup-wizard__step-number"
	ClassSetupWizardConnector = "dbx-setup-wizard__step-connector"
	ClassSetupWizardBody      = "dbx-setup-wizard__body"

	// Utility
	ClassSpinner      = "dbx-spinner"
	ClassSpinnerSmall = "dbx-spinner--small"
	ClassSpinnerLarge = "dbx-spinner--large"
	ClassIcon         = "dbx-icon"
)

// classes joins multiple CSS class names into a single string.
func classes(names ...string) string {
	return coremd.Classes(names...)
}

// classIf returns the base class with the conditional class appended if condition is true.
func classIf(base string, condition bool, conditional string) string {
	return coremd.ClassIf(base, condition, conditional)
}
