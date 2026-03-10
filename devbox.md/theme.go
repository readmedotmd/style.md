package devboxmd

import (
	coremd "github.com/readmedotmd/core.md"
)

var theme = coremd.Themed{ClassMap: coremd.ClassMap{
	// Button
	Btn:        ClassBtn,
	BtnPrimary: ClassBtnPrimary,
	BtnDanger:  ClassBtnDanger,
	BtnSmall:   ClassBtnSmall,
	BtnToolbar: ClassBtnToolbar,

	// Layout
	AppShell:      ClassAppShell,
	AppShellBody:  ClassAppShellBody,
	AppShellMain:  ClassAppShellMain,
	Navbar:        ClassNavbar,
	Sidebar:       ClassSidebar,
	SidebarOpen:   ClassSidebarOpen,
	SidebarHeader: ClassSidebarHeader,
	Panel:         ClassPanel,
	PanelExpanded: ClassPanelExpanded,
	ModalBackdrop: ClassModalBackdrop,
	Modal:         ClassModal,
	ModalBody:     ClassModalBody,
	ModalFooter:   ClassModalFooter,
	DragHandle:    ClassDragHandle,

	// Navigation
	NavLink:       ClassNavLink,
	NavLinkActive: ClassNavLinkActive,
	TabBar:        ClassTabBar,
	BottomTabBar:  ClassBottomTabBar,

	// Data Display
	Message:           ClassMessage,
	MessageUser:       ClassMessageUser,
	MessageAssistant:  ClassMessageAssistant,
	MessageStreaming:  ClassMessageStreaming,
	Thinking:          ClassThinking,
	ThinkingCollapsib: ClassThinkingCollapsib,
	ToolBadge:         ClassToolBadge,
	QuestionPrompt:    ClassQuestionPrompt,
	StatusBadge:       ClassStatusBadge,
	StatusBadgePrefix: "dbx-status-badge--",
	StatusDot:         ClassStatusDot,
	StatusDotPrefix:   "dbx-status-dot--",
	LabelBadge:        ClassLabelBadge,
	UsageBadge:        ClassUsageBadge,
	DiffViewer:        ClassDiffViewer,
	DataTable:         ClassDataTable,
	EmptyState:        ClassEmptyState,
	ClusterStats:      ClassClusterStats,

	// List
	ConversationItem:       ClassConversationItem,
	ConversationItemActive: ClassConversationItemActive,
	InstanceCard:           ClassInstanceCard,
	InstanceCardActive:     ClassInstanceCardActive,
	ServiceRow:             ClassServiceRow,
	RunnerRow:              ClassRunnerRow,
	FileTree:               ClassFileTree,

	// Form
	FormGroup:       ClassFormGroup,
	TextInput:       ClassTextInput,
	TextInputError:  ClassTextInputError,
	Textarea:        ClassTextarea,
	TextareaAutoGrw: ClassTextareaAutoGrw,
	TextareaFixed:   ClassTextareaFixed,
	Select:          ClassSelect,
	Checkbox:        ClassCheckbox,
	FeatureRow:      ClassFeatureRow,
	VariableRow:     ClassVariableRow,
	ErrorMessage:    ClassErrorMessage,
	SuccessMessage:  ClassSuccessMessage,

	// Input
	ChatInput:    ClassChatInput,
	Autocomplete: ClassAutocomplete,
	MessageQueue: ClassMessageQueue,
	SearchInput:  ClassSearchInput,

	// Overlay
	SearchOverlay: ClassSearchOverlay,
	ContextMenu:   ClassContextMenu,

	// Panel compositions
	SkillCard:     ClassSkillCard,
	TerminalPanel: ClassTerminalPanel,

	// Page
	LoginPage: ClassLoginPage,

	// Utility
	Spinner:      ClassSpinner,
	SpinnerSmall: ClassSpinnerSmall,
	SpinnerLarge: ClassSpinnerLarge,
	Icon:         ClassIcon,
}}
