package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// RichTextAttributes defines attributes for RichText widget.
// See: https://api.flutter.dev/flutter/widgets/RichText-class.html
type RichTextAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	TextSpan              *duit_props.TextSpan      `json:"textSpan"`
	Style                 *duit_props.TextStyle     `json:"style,omitempty"`
	TextAlign             duit_props.TextAlign      `json:"textAlign,omitempty"`
	MaxLines              int                       `json:"maxLines,omitempty"`
	SoftWrap              bool                      `json:"softWrap,omitempty"`
	Overflow              duit_props.TextOverflow   `json:"overflow,omitempty"`
	SemanticsLabel        string                    `json:"semanticsLabel,omitempty"`
	SelectionColor        *duit_props.Color         `json:"selectionColor,omitempty"`
	TextDirection         duit_props.TextDirection  `json:"textDirection,omitempty"`
	TextScaler            *duit_props.TextScaler    `json:"textScaleFactor,omitempty"`
	TextWidthBasis        duit_props.TextWidthBasis `json:"textWidthBasis,omitempty"`
}

//bindgen:exclude
func (r *RichTextAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.TextSpan == nil {
		return errors.New("textSpan property is required")
	}

	return nil
}

// NewRichTextAttributes creates a new RichTextAttributes instance.
func NewRichTextAttributes() *RichTextAttributes {
	return &RichTextAttributes{}
}

// SetTextSpan sets the text span for the rich text widget.
func (r *RichTextAttributes) SetTextSpan(textSpan *duit_props.TextSpan) *RichTextAttributes {
	r.TextSpan = textSpan
	return r
}

// SetStyle sets the style for the rich text widget.
func (r *RichTextAttributes) SetStyle(style *duit_props.TextStyle) *RichTextAttributes {
	r.Style = style
	return r
}

// SetTextAlign sets the text align for the rich text widget.
func (r *RichTextAttributes) SetTextAlign(textAlign duit_props.TextAlign) *RichTextAttributes {
	r.TextAlign = textAlign
	return r
}

// SetMaxLines sets the max lines for the rich text widget.
func (r *RichTextAttributes) SetMaxLines(maxLines int) *RichTextAttributes {
	r.MaxLines = maxLines
	return r
}

// SetSoftWrap sets the soft wrap for the rich text widget.
func (r *RichTextAttributes) SetSoftWrap(softWrap bool) *RichTextAttributes {
	r.SoftWrap = softWrap
	return r
}

// SetOverflow sets the overflow for the rich text widget.
func (r *RichTextAttributes) SetOverflow(overflow duit_props.TextOverflow) *RichTextAttributes {
	r.Overflow = overflow
	return r
}

// SetSemanticsLabel sets the semantics label for the rich text widget.
func (r *RichTextAttributes) SetSemanticsLabel(semanticsLabel string) *RichTextAttributes {
	r.SemanticsLabel = semanticsLabel
	return r
}

// SetSelectionColor sets the selection color for the rich text widget.
func (r *RichTextAttributes) SetSelectionColor(selectionColor *duit_props.Color) *RichTextAttributes {
	r.SelectionColor = selectionColor
	return r
}

// SetTextDirection sets the text direction for the rich text widget.
func (r *RichTextAttributes) SetTextDirection(textDirection duit_props.TextDirection) *RichTextAttributes {
	r.TextDirection = textDirection
	return r
}

// SetTextScaler sets the text scaler for the rich text widget.
func (r *RichTextAttributes) SetTextScaler(textScaler *duit_props.TextScaler) *RichTextAttributes {
	r.TextScaler = textScaler
	return r
}

// SetTextWidthBasis sets the text width basis for the rich text widget.
func (r *RichTextAttributes) SetTextWidthBasis(textWidthBasis duit_props.TextWidthBasis) *RichTextAttributes {
	r.TextWidthBasis = textWidthBasis
	return r
}
