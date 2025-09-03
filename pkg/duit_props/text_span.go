package duit_props

// TextSpan represents a span of text with styling
// Note: Uses any type for Color to avoid import cycles
type TextSpan[T any] struct {
	Text     string         `json:"text,omitempty"`
	Style    *TextStyle[T]  `json:"style,omitempty"`
	SpellOut bool           `json:"spellOut,omitempty"`
	Children []*TextSpan[T] `json:"children,omitempty"`
}
