package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// TextAttributes represents attributes for Text widget.
// https://api.flutter.dev/flutter/widgets/Text-class.html
type TextAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Data                              string                    `json:"data"`
	SemanticsLabel                    string                    `json:"semanticsLabel,omitempty"`
	TextAlign                         duit_props.TextAlign      `json:"textAlign,omitempty"`
	TextDirection                     duit_props.TextDirection  `json:"textDirection,omitempty"`
	Overflow                          duit_props.TextOverflow   `json:"overflow,omitempty"`
	SoftWrap                          duit_utils.Tristate[bool] `json:"softWrap,omitempty"`
	MaxLines                          int                       `json:"maxLines,omitempty"`
	Style                             *duit_props.TextStyle     `json:"style,omitempty"`
}

// NewTextAttributes creates a new instance of TextAttributes.
func NewTextAttributes() *TextAttributes {
	return &TextAttributes{}
}

// SetData sets the data property and returns the TextAttributes instance for method chaining.
func (r *TextAttributes) SetData(data string) *TextAttributes {
	r.Data = data
	return r
}

// SetSemanticsLabel sets the semanticsLabel property and returns the TextAttributes instance for method chaining.
func (r *TextAttributes) SetSemanticsLabel(semanticsLabel string) *TextAttributes {
	r.SemanticsLabel = semanticsLabel
	return r
}

// SetTextAlign sets the textAlign property and returns the TextAttributes instance for method chaining.
func (r *TextAttributes) SetTextAlign(textAlign duit_props.TextAlign) *TextAttributes {
	r.TextAlign = textAlign
	return r
}

// SetTextDirection sets the textDirection property and returns the TextAttributes instance for method chaining.
func (r *TextAttributes) SetTextDirection(textDirection duit_props.TextDirection) *TextAttributes {
	r.TextDirection = textDirection
	return r
}

// SetOverflow sets the overflow property and returns the TextAttributes instance for method chaining.
func (r *TextAttributes) SetOverflow(overflow duit_props.TextOverflow) *TextAttributes {
	r.Overflow = overflow
	return r
}

// SetSoftWrap sets the softWrap property and returns the TextAttributes instance for method chaining.
func (r *TextAttributes) SetSoftWrap(softWrap bool) *TextAttributes {
	r.SoftWrap = duit_utils.BoolValue(softWrap)
	return r
}

// SetMaxLines sets the maxLines property and returns the TextAttributes instance for method chaining.
func (r *TextAttributes) SetMaxLines(maxLines int) *TextAttributes {
	r.MaxLines = maxLines
	return r
}

// SetStyle sets the style property and returns the TextAttributes instance for method chaining.
func (r *TextAttributes) SetStyle(style *duit_props.TextStyle) *TextAttributes {
	r.Style = style
	return r
}

//bindgen:exclude
func (r *TextAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if len(r.Data) == 0 {
		return errors.New("data property is required")
	}

	return nil
}
