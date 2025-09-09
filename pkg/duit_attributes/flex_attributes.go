package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// FlexAttributes defines attributes for Flex widget.
// See: https://api.flutter.dev/flutter/widgets/Flex-class.html
type FlexAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	MainAxisAlignment                 duit_props.MainAxisAlignment  `json:"mainAxisAlignment,omitempty"`
	MainAxisSize                      duit_props.MainAxisSize       `json:"mainAxisSize,omitempty"`
	CrossAxisAlignment                duit_props.CrossAxisAlignment `json:"crossAxisAlignment,omitempty"`
	TextDirection                     duit_props.TextDirection      `json:"textDirection,omitempty"`
	VerticalDirection                 duit_props.VerticalDirection  `json:"verticalDirection,omitempty"`
	ClipBehavior                      duit_props.Clip               `json:"clipBehavior,omitempty"`
}

//bindgen:exclude
func (r *FlexAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	return nil
}

// NewFlexAttributes creates a new FlexAttributes instance.
func NewFlexAttributes() *FlexAttributes {
	return &FlexAttributes{}
}

// SetMainAxisAlignment sets the main axis alignment for the flex widget.
func (r *FlexAttributes) SetMainAxisAlignment(mainAxisAlignment duit_props.MainAxisAlignment) *FlexAttributes {
	r.MainAxisAlignment = mainAxisAlignment
	return r
}

// SetMainAxisSize sets the main axis size for the flex widget.
func (r *FlexAttributes) SetMainAxisSize(mainAxisSize duit_props.MainAxisSize) *FlexAttributes {
	r.MainAxisSize = mainAxisSize
	return r
}

// SetCrossAxisAlignment sets the cross axis alignment for the flex widget.
func (r *FlexAttributes) SetCrossAxisAlignment(crossAxisAlignment duit_props.CrossAxisAlignment) *FlexAttributes {
	r.CrossAxisAlignment = crossAxisAlignment
	return r
}

// SetTextDirection sets the text direction for the flex widget.
func (r *FlexAttributes) SetTextDirection(textDirection duit_props.TextDirection) *FlexAttributes {
	r.TextDirection = textDirection
	return r
}

// SetVerticalDirection sets the vertical direction for the flex widget.
func (r *FlexAttributes) SetVerticalDirection(verticalDirection duit_props.VerticalDirection) *FlexAttributes {
	r.VerticalDirection = verticalDirection
	return r
}

// SetClipBehavior sets the clip behavior for the flex widget.
func (r *FlexAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *FlexAttributes {
	r.ClipBehavior = clipBehavior
	return r
}
