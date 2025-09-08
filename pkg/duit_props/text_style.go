package duit_props

// TextStyle represents the style of text in a widget.
// It is used to control the style of text in a widget.
//
// This is a Go implementation of Flutter's TextStyle class.
// For more information, see: https://api.flutter.dev/flutter/painting/TextStyle-class.html
type TextStyle struct {
	Color               *Color              `json:"color,omitempty"`
	DecorationColor     *Color              `json:"decorationColor,omitempty"`
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

// NewTextStyle creates a new instance of TextStyle
func NewTextStyle() *TextStyle {
	return &TextStyle{}
}

// SetColor sets the color of the text
func (r *TextStyle) SetColor(color *Color) *TextStyle {
	r.Color = color
	return r
}

// SetDecorationColor sets the color of the decoration
func (r *TextStyle) SetDecorationColor(color *Color) *TextStyle {
	r.DecorationColor = color
	return r
}

// SetFontFamily sets the font family of the text
func (r *TextStyle) SetFontFamily(fontFamily string) *TextStyle {
	r.FontFamily = fontFamily
	return r
}

// SetFontWeight sets the font weight of the text
func (r *TextStyle) SetFontWeight(fontWeight FontWeight) *TextStyle {
	r.FontWeight = fontWeight
	return r
}

// SetFontSize sets the font size of the text
func (r *TextStyle) SetFontSize(fontSize float32) *TextStyle {
	r.FontSize = fontSize
	return r
}

// SetLetterSpacing sets the letter spacing of the text
func (r *TextStyle) SetLetterSpacing(letterSpacing float32) *TextStyle {
	r.LetterSpacing = letterSpacing
	return r
}

// SetWordSpacing sets the word spacing of the text
func (r *TextStyle) SetWordSpacing(wordSpacing float32) *TextStyle {
	r.WordSpacing = wordSpacing
	return r
}

// SetHeight sets the height of the text
func (r *TextStyle) SetHeight(height float32) *TextStyle {
	r.Height = height
	return r
}

// SetDecoration sets the decoration of the text
func (r *TextStyle) SetDecoration(decoration TextDecoration) *TextStyle {
	r.Decoration = decoration
	return r
}

// SetDecorationStyle sets the decoration style of the text
func (r *TextStyle) SetDecorationStyle(decorationStyle TextDecorationStyle) *TextStyle {
	r.DecorationStyle = decorationStyle
	return r
}

// SetDecorationThickness sets the decoration thickness of the text
func (r *TextStyle) SetDecorationThickness(decorationThickness float32) *TextStyle {
	r.DecorationThickness = decorationThickness
	return r
}

func (r *TextStyle) Validate() error {
	if r.Color != nil {
		if err := r.Color.Validate(); err != nil {
			return err
		}
	}

	if r.DecorationColor != nil {
		if err := r.DecorationColor.Validate(); err != nil {
			return err
		}
	}

	return nil
}