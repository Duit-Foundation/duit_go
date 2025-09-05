package duit_props

// TextSpan represents a span of text with styling
type TextSpan struct {
	Text     string      `json:"text,omitempty"`
	Style    *TextStyle  `json:"style,omitempty"`
	SpellOut bool        `json:"spellOut,omitempty"`
	Children []*TextSpan `json:"children,omitempty"`
}
