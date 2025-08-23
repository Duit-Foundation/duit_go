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

func (t *ThemeConsumer) Validate() error {
	if t == nil {
		return  nil
	}

	if t.Theme == "" && !t.IgnoreTheme {
		return errors.New("theme property is required unless ignoreTheme is true")
	}

	return nil
}
