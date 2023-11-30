package duit_attributes

type TextStyle struct {
	Color         string  `json:"color,omitempty"`
	FontFamily    string  `json:"fontFamily,omitempty"`
	FontWeight    int     `json:"fontWeight,omitempty"`
	FontSize      float32 `json:"fontSize,omitempty"`
	LetterSpacing float32 `json:"letterSpacing,omitempty"`
	WordSpacing   float32 `json:"wordSpacing,omitempty"`
	Height        float32 `json:"height,omitempty"`
}
