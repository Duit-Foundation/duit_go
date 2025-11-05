package duit_props

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"

// TextSpan represents a span of text with styling
//
// This is a Go implementation of Flutter's TextSpan class.
// For more information, see: https://api.flutter.dev/flutter/painting/TextSpan-class.html
type TextSpan struct {
	Text     string                    `json:"text,omitempty"`
	Style    *TextStyle                `json:"style,omitempty"`
	SpellOut duit_utils.Tristate[bool] `json:"spellOut,omitempty"`
	Children []*TextSpan               `json:"children,omitempty"`
}

// NewTextSpan creates a new instance of TextSpan
func NewTextSpan() *TextSpan {
	return &TextSpan{}
}

// SetText sets the text of the text span
func (r *TextSpan) SetText(text string) *TextSpan {
	r.Text = text
	return r
}

// SetStyle sets the style of the text span
func (r *TextSpan) SetStyle(style *TextStyle) *TextSpan {
	r.Style = style
	return r
}

// SetSpellOut sets the spell out of the text span
func (r *TextSpan) SetSpellOut(spellOut bool) *TextSpan {
	r.SpellOut = duit_utils.BoolValue(spellOut)
	return r
}

// SetChildren sets the children of the text span
func (r *TextSpan) SetChildren(children []*TextSpan) *TextSpan {
	r.Children = children
	return r
}

func (r *TextSpan) AddChild(child *TextSpan) *TextSpan {
	r.Children = append(r.Children, child)
	return r
}

func (r *TextSpan) Validate() error {
	if r.Style != nil {
		if err := r.Style.Validate(); err != nil {
			return err
		}
	}

	if r.Children != nil {
		for _, child := range r.Children {
			if err := child.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
