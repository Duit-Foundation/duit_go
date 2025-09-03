package duit_props

// TextStyle represents text styling properties
// Note: Uses any type for Color to avoid import cycles
type TextStyle[T any] struct {
	Color               T                    `json:"color,omitempty"`
	FontFamily          string               `json:"fontFamily,omitempty"`
	FontWeight          FontWeight           `json:"fontWeight,omitempty"`
	FontSize            float32              `json:"fontSize,omitempty"`
	LetterSpacing       float32              `json:"letterSpacing,omitempty"`
	WordSpacing         float32              `json:"wordSpacing,omitempty"`
	Height              float32              `json:"height,omitempty"`
	Decoration          TextDecoration       `json:"decoration,omitempty"`
	DecorationStyle     TextDecorationStyle  `json:"decorationStyle,omitempty"`
	DecorationColor     T                    `json:"decorationColor,omitempty"`
	DecorationThickness float32              `json:"decorationThickness,omitempty"`
}
