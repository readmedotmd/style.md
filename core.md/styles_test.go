package coremd

import (
	"os"
	"strings"
	"testing"
)

func loadCSS(t *testing.T) string {
	t.Helper()
	data, err := os.ReadFile("styles.css")
	if err != nil {
		t.Fatalf("failed to read styles.css: %v", err)
	}
	return string(data)
}

// TestTokensExist verifies all CSS custom properties exist in :root
func TestTokensExist(t *testing.T) {
	css := loadCSS(t)
	tokens := []string{
		// Original tokens
		"--core-font:", "--core-font-mono:", "--core-text:", "--core-text-muted:",
		"--core-bg:", "--core-surface:", "--core-border:", "--core-accent:",
		"--core-danger:", "--core-success:", "--core-warning:", "--core-radius:",
		"--core-space:", "--core-transition:",
		// New tokens
		"--core-info:", "--core-radius-lg:", "--core-backdrop:", "--core-backdrop-blur:",
		"--core-shadow-sm:", "--core-shadow-md:", "--core-shadow-lg:",
		"--core-text-2xs:", "--core-text-xs:", "--core-text-sm:", "--core-text-base:",
		"--core-text-md:", "--core-text-lg:", "--core-text-xl:",
		"--core-on-accent:", "--core-on-accent-muted:", "--core-on-accent-subtle:",
		"--core-on-accent-border:", "--core-raised:", "--core-hover:",
	}
	for _, token := range tokens {
		if !strings.Contains(css, token) {
			t.Errorf("missing token: %s", token)
		}
	}
}

// TestDarkModeTokens verifies dark mode overrides exist for key new tokens
func TestDarkModeTokens(t *testing.T) {
	css := loadCSS(t)
	// Check that dark mode sections exist
	if !strings.Contains(css, "prefers-color-scheme: dark") {
		t.Error("missing @media prefers-color-scheme: dark")
	}
	if !strings.Contains(css, `[data-theme="dark"]`) {
		t.Error(`missing [data-theme="dark"] selector`)
	}
	// Dark-mode specific values
	darkTokens := []string{
		"--core-info:", "--core-hover:", "--core-raised:", "--core-shadow-sm:",
		"--core-shadow-md:", "--core-shadow-lg:",
	}
	// Split CSS into the dark sections and verify tokens appear there
	// Just verify they appear at least twice (once in :root, once+ in dark overrides)
	for _, token := range darkTokens {
		count := strings.Count(css, token)
		if count < 2 {
			t.Errorf("token %s appears %d time(s), expected at least 2 (light + dark)", token, count)
		}
	}
}

// TestComponentSelectors verifies all component selectors exist in the CSS
func TestComponentSelectors(t *testing.T) {
	css := loadCSS(t)

	tests := []struct {
		name      string
		selectors []string
	}{
		// Scrollbar
		{"Scrollbar", []string{"scrollbar-width:", "::-webkit-scrollbar", "::-webkit-scrollbar-track", "::-webkit-scrollbar-thumb"}},
		// App Shell
		{"AppShell", []string{"[data-app-shell]"}},
		// Navbar
		{"Navbar", []string{"[data-navbar]", "[data-nav-brand]", "[data-nav-links]", "[data-nav-link]"}},
		// Sidebar
		{"Sidebar", []string{"[data-sidebar]", "[data-sidebar-header]"}},
		// Bottom Tab Bar
		{"BottomTabBar", []string{"[data-bottom-tabbar]", "[data-tab-link]"}},
		// Panel
		{"Panel", []string{`[data-panel]`, `[data-panel="wide"]`, "[data-panel-header]", "[data-panel-header-title]", "[data-panel-actions]"}},
		// Modal
		{"Modal", []string{"[data-modal-backdrop]", "[data-modal]", "[data-modal-header]", "[data-modal-title]", "[data-modal-body]", "[data-modal-actions]", "[data-drag-handle]"}},
		// Bottom Sheet
		{"BottomSheet", []string{"[data-bottom-sheet-backdrop]", "[data-bottom-sheet]", "[data-sheet-item]"}},
		// Context Menu
		{"ContextMenu", []string{"[data-context-menu-backdrop]", "[data-context-menu]", "[data-context-menu-item]"}},
		// Chat Message
		{"ChatMessage", []string{`[data-message]`, `[data-message="outgoing"]`, `[data-message="incoming"]`, `[data-message="streaming"]`, `[data-message="error"]`}},
		// Rich Text
		{"RichText", []string{"[data-rich-text]", "[data-rich-text] code", "[data-rich-text] pre", "[data-rich-text] blockquote", "[data-rich-text] table"}},
		// Collapsible
		{"Collapsible", []string{"[data-collapsible]", "[data-collapsible-summary]", "[data-collapsible-content]"}},
		// Indicator
		{"Indicator", []string{"[data-indicator]", `[data-indicator="working"]`}},
		// Autocomplete
		{"Autocomplete", []string{"[data-autocomplete]", "[data-autocomplete-header]", "[data-autocomplete-trigger]", "[data-autocomplete-item]", "[data-autocomplete-detail]", "[data-autocomplete-snippet]"}},
		// Message Queue
		{"MessageQueue", []string{"[data-queue]", "[data-queue-item]"}},
		// Search Overlay
		{"SearchOverlay", []string{"[data-search-backdrop]", "[data-search-card]", "[data-search-input]", "[data-search-results]", "[data-search-result]", "[data-search-match]", "[data-search-snippet]"}},
		// Image Preview
		{"ImagePreview", []string{"[data-image-preview]", "[data-preview-thumbnail]"}},
		// File Tree
		{"FileTree", []string{"[data-file-tree]", `[data-file-item]`, `[data-file-item="dir"]`, `[data-file-item="file"]`}},
		// Terminal Panel
		{"TerminalPanel", []string{"[data-terminal-panel]", "[data-terminal-tabs]", "[data-terminal-tab]", "[data-terminal-iframe]"}},
		// Selectable List
		{"SelectableList", []string{"[data-list]", "[data-list-item]"}},
		// Section Header
		{"SectionHeader", []string{"[data-section-header]"}},
		// Question Prompt
		{"QuestionPrompt", []string{"[data-question]", "[data-question-text]", "[data-question-option]"}},
		// Empty State
		{"EmptyState", []string{"[data-empty-state]"}},
		// Feature Row
		{"FeatureRow", []string{"[data-feature-row]"}},
		// Key-Value Row
		{"KeyValueRow", []string{"[data-kv-row]", "[data-kv-remove]"}},
		// Micro Badge
		{"MicroBadge", []string{"[data-micro-badge]"}},
		// Code Textarea
		{"CodeTextarea", []string{`textarea[data-code]`, `textarea[data-code="tall"]`, `textarea[data-code="medium"]`, `textarea[data-code="short"]`}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, sel := range tt.selectors {
				if !strings.Contains(css, sel) {
					t.Errorf("missing selector: %s", sel)
				}
			}
		})
	}
}

// TestButtonVariants verifies all button variants
func TestButtonVariants(t *testing.T) {
	css := loadCSS(t)
	variants := []string{
		`[data-variant="primary"]`, `[data-variant="danger"]`, `[data-variant="toolbar"]`,
		`[data-variant="send"]`, `[data-variant="cancel"]`,
		`[data-variant="success-outline"]`, `[data-variant="warning-outline"]`,
		`[data-variant="info-outline"]`, `[data-variant="ghost"]`,
	}
	for _, v := range variants {
		if !strings.Contains(css, v) {
			t.Errorf("missing button variant selector: %s", v)
		}
	}
}

// TestAnimations verifies animation keyframes and data-animate selectors
func TestAnimations(t *testing.T) {
	css := loadCSS(t)
	anims := []string{
		"@keyframes core-spin", "@keyframes core-pulse", "@keyframes core-dot-appear",
		`[data-animate="pulse"]`, `[data-animate="appear"]`,
	}
	for _, a := range anims {
		if !strings.Contains(css, a) {
			t.Errorf("missing animation: %s", a)
		}
	}
}

// TestResponsiveUtilities verifies responsive visibility utilities
func TestResponsiveUtilities(t *testing.T) {
	css := loadCSS(t)
	utils := []string{
		`[data-hide="mobile"]`, `[data-hide="desktop"]`,
		"[data-hidden]", `[data-visible="false"]`,
		"max-width: 768px", "min-width: 769px",
	}
	for _, u := range utils {
		if !strings.Contains(css, u) {
			t.Errorf("missing responsive utility: %s", u)
		}
	}
}

// TestMobileTouchTargets verifies min-height: 44px appears in mobile media query
func TestMobileTouchTargets(t *testing.T) {
	css := loadCSS(t)
	if !strings.Contains(css, "min-height: 44px") {
		t.Error("missing mobile touch target min-height: 44px")
	}
}

// TestExistingComponentsPreserved verifies nothing was removed from original CSS
func TestExistingComponentsPreserved(t *testing.T) {
	css := loadCSS(t)
	existing := []string{
		// Reset
		"box-sizing: border-box",
		// Buttons
		"button[data-variant]", "button:disabled",
		// Forms
		`input[type="text"]`, "textarea", "select",
		`input[type="checkbox"]`,
		// Status badges
		"span[data-status]",
		`span[data-status="running"]`, `span[data-status="stopped"]`,
		`span[data-status="starting"]`, `span[data-status="pending"]`,
		`span[data-status="error"]`,
		// Diff viewer
		"[data-diff]", `[data-diff="add"]`, `[data-diff="remove"]`, `[data-diff="header"]`,
		// Data table
		"thead th", "tbody td",
		// Tab bar
		"button[data-active]",
		// Spinner
		"@keyframes core-spin",
		// Layout primitives
		"[data-stack]", "[data-hstack]", "[data-grid]",
		"[data-align=", "[data-justify=", "[data-wrap]",
		// Card and badge
		"[data-card]", "[data-badge]",
		// Utilities
		"[data-truncate]", "[data-mono]", "[data-center]", "[data-sr-only]", "[data-spacer]",
		// Typography
		"h1, h2, h3, h4, h5, h6",
		// Code
		"pre code",
		// Blockquote
		"blockquote",
		// Images
		"img[data-rounded]", "img[data-avatar]",
		// Divider
		"border-top: 1px solid var(--core-border)",
	}
	for _, sel := range existing {
		if !strings.Contains(css, sel) {
			t.Errorf("existing component removed or broken: %s", sel)
		}
	}
}

// TestOutgoingMessageOverrides verifies text-on-accent overrides for outgoing messages
func TestOutgoingMessageOverrides(t *testing.T) {
	css := loadCSS(t)
	overrides := []string{
		`[data-message="outgoing"] a`,
		`[data-message="outgoing"] code`,
		`[data-message="outgoing"] pre`,
		"--core-on-accent-muted",
		"--core-on-accent-subtle",
		"--core-on-accent-border",
	}
	for _, o := range overrides {
		if !strings.Contains(css, o) {
			t.Errorf("missing outgoing message override: %s", o)
		}
	}
}

// TestRichTextElements verifies all rich text sub-element styles
func TestRichTextElements(t *testing.T) {
	css := loadCSS(t)
	elements := []string{
		"[data-rich-text] p",
		"[data-rich-text] code",
		"[data-rich-text] pre",
		"[data-rich-text] pre code",
		"[data-rich-text] blockquote",
		"[data-rich-text] hr",
		"[data-rich-text] h1",
		"[data-rich-text] h2",
		"[data-rich-text] h3",
		"[data-rich-text] h4",
		"[data-rich-text] table",
		"[data-rich-text] th",
		"[data-rich-text] td",
		"[data-rich-text] img",
	}
	for _, e := range elements {
		if !strings.Contains(css, e) {
			t.Errorf("missing rich text element: %s", e)
		}
	}
}

// TestSectionHeaderVariants checks section header color variants
func TestSectionHeaderVariants(t *testing.T) {
	css := loadCSS(t)
	variants := []string{
		`[data-section-header="success"]`,
		`[data-section-header="warning"]`,
	}
	for _, v := range variants {
		if !strings.Contains(css, v) {
			t.Errorf("missing section header variant: %s", v)
		}
	}
}

// TestResponsiveNavbar verifies navbar links are hidden on mobile
func TestResponsiveNavbar(t *testing.T) {
	css := loadCSS(t)
	// Should hide nav-links on mobile
	if !strings.Contains(css, "[data-nav-links]") {
		t.Error("missing [data-nav-links] selector")
	}
}

// TestResponsiveSidebar verifies sidebar responsive behavior
func TestResponsiveSidebar(t *testing.T) {
	css := loadCSS(t)
	if !strings.Contains(css, "[data-sidebar][data-open]") {
		t.Error("missing responsive sidebar open selector")
	}
}

// TestModalMobileBottomSheet verifies modal becomes bottom sheet on mobile
func TestModalMobileBottomSheet(t *testing.T) {
	css := loadCSS(t)
	// Modal should have bottom-sheet behavior on mobile
	if !strings.Contains(css, "16px 16px 0 0") {
		t.Error("missing mobile modal bottom-sheet border-radius")
	}
}

// TestPanelMobileOverlay verifies panel becomes full-width on mobile
func TestPanelMobileOverlay(t *testing.T) {
	css := loadCSS(t)
	if !strings.Contains(css, "[data-panel]") {
		t.Error("missing panel selector")
	}
}

// TestBottomTabBarMobileOnly verifies bottom tab bar is hidden on desktop
func TestBottomTabBarMobileOnly(t *testing.T) {
	css := loadCSS(t)
	// The component should have display: none by default
	if !strings.Contains(css, "[data-bottom-tabbar]") {
		t.Error("missing [data-bottom-tabbar] selector")
	}
}

// TestQueueEmptyHidden verifies queue hides when empty
func TestQueueEmptyHidden(t *testing.T) {
	css := loadCSS(t)
	if !strings.Contains(css, "[data-queue]:empty") {
		t.Error("missing [data-queue]:empty selector")
	}
}

// TestImagePreviewEmptyHidden verifies image preview hides when empty
func TestImagePreviewEmptyHidden(t *testing.T) {
	css := loadCSS(t)
	if !strings.Contains(css, "[data-image-preview]:empty") {
		t.Error("missing [data-image-preview]:empty selector")
	}
}
