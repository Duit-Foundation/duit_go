package duit_attributes

type ThemeOverrideRule uint8

const (
	ThemeOverlay ThemeOverrideRule = iota
	ThemePriority
)

type ThemeConsumer struct {
	Theme string `json:"theme,omitempty"`
	IgnoreTheme bool `json:"ignoreTheme,omitempty"`
	OverrideRule ThemeOverrideRule `json:"overrideRule,omitempty"`
}
