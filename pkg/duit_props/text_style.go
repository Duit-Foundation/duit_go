package duit_props

// TextStyle represents text styling properties
type TextStyle struct {
	Color               *Color               `json:"color,omitempty"`
	DecorationColor     *Color               `json:"decorationColor,omitempty"`
	FontFamily          string              `json:"fontFamily,omitempty"`
	FontWeight          FontWeight          `json:"fontWeight,omitempty"`
	FontSize            float32             `json:"fontSize,omitempty"`
	LetterSpacing       float32             `json:"letterSpacing,omitempty"`
	WordSpacing         float32             `json:"wordSpacing,omitempty"`
	Height              float32             `json:"height,omitempty"`
	Decoration          TextDecoration      `json:"decoration,omitempty"`
	DecorationStyle     TextDecorationStyle `json:"decorationStyle,omitempty"`
	DecorationThickness float32             `json:"decorationThickness,omitempty"`
}
