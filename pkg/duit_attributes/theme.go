package duit_attributes

import "errors"

type ThemeOverrideRule = uint8

const (
	ThemeOverlay ThemeOverrideRule = iota
	ThemePriority
)

type ThemeConsumer struct {
	Theme        string            `json:"theme,omitempty"`
	IgnoreTheme  bool              `json:"ignoreTheme,omitempty"`
	OverrideRule ThemeOverrideRule `json:"overrideRule,omitempty"`
}

// SetTheme sets the theme property
func (t *ThemeConsumer) SetTheme(theme string) {
	t.Theme = theme
}

// SetIgnoreTheme sets the ignoreTheme property
func (t *ThemeConsumer) SetIgnoreTheme(ignoreTheme bool) {
	t.IgnoreTheme = ignoreTheme
}

// SetOverrideRule sets the overrideRule property
func (t *ThemeConsumer) SetOverrideRule(overrideRule ThemeOverrideRule) {
	t.OverrideRule = overrideRule
}

//bindgen:exclude
func (t *ThemeConsumer) Validate() error {
	if t == nil {
		return nil
	}

	if t.Theme == "" && !t.IgnoreTheme {
		return errors.New("theme property is required unless ignoreTheme is true")
	}

	return nil
}
