package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// WrapAttributes represents attributes for Wrap widget.
// https://api.flutter.dev/flutter/widgets/Wrap-class.html
type WrapAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	TextDirection                     duit_props.TextDirection      `json:"textDirection,omitempty"`
	VerticalDirection                 duit_props.VerticalDirection  `json:"verticalDirection,omitempty"`
	Alignment                         duit_props.MainAxisAlignment  `json:"alignment,omitempty"`
	RunAlignment                      duit_props.MainAxisAlignment  `json:"runAlignment,omitempty"`
	CrossAxisAlignment                duit_props.CrossAxisAlignment `json:"crossAxisAlignment,omitempty"`
	Spacing                           float32                       `json:"spacing,omitempty"`
	RunSpacing                        float32                       `json:"runSpacing,omitempty"`
	ClipBehavior                      duit_props.Clip               `json:"clipBehavior,omitempty"`
	Direction                         duit_props.Axis               `json:"direction,omitempty"`
}

// NewWrapAttributes creates a new instance of WrapAttributes.
func NewWrapAttributes() *WrapAttributes {
	return &WrapAttributes{}
}

// SetTextDirection sets the textDirection property and returns the WrapAttributes instance for method chaining.
func (r *WrapAttributes) SetTextDirection(textDirection duit_props.TextDirection) *WrapAttributes {
	r.TextDirection = textDirection
	return r
}

// SetVerticalDirection sets the verticalDirection property and returns the WrapAttributes instance for method chaining.
func (r *WrapAttributes) SetVerticalDirection(verticalDirection duit_props.VerticalDirection) *WrapAttributes {
	r.VerticalDirection = verticalDirection
	return r
}

// SetAlignment sets the alignment property and returns the WrapAttributes instance for method chaining.
func (r *WrapAttributes) SetAlignment(alignment duit_props.MainAxisAlignment) *WrapAttributes {
	r.Alignment = alignment
	return r
}

// SetRunAlignment sets the runAlignment property and returns the WrapAttributes instance for method chaining.
func (r *WrapAttributes) SetRunAlignment(runAlignment duit_props.MainAxisAlignment) *WrapAttributes {
	r.RunAlignment = runAlignment
	return r
}

// SetCrossAxisAlignment sets the crossAxisAlignment property and returns the WrapAttributes instance for method chaining.
func (r *WrapAttributes) SetCrossAxisAlignment(crossAxisAlignment duit_props.CrossAxisAlignment) *WrapAttributes {
	r.CrossAxisAlignment = crossAxisAlignment
	return r
}

// SetSpacing sets the spacing property and returns the WrapAttributes instance for method chaining.
func (r *WrapAttributes) SetSpacing(spacing float32) *WrapAttributes {
	r.Spacing = spacing
	return r
}

// SetRunSpacing sets the runSpacing property and returns the WrapAttributes instance for method chaining.
func (r *WrapAttributes) SetRunSpacing(runSpacing float32) *WrapAttributes {
	r.RunSpacing = runSpacing
	return r
}

// SetClipBehavior sets the clipBehavior property and returns the WrapAttributes instance for method chaining.
func (r *WrapAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *WrapAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetDirection sets the direction property and returns the WrapAttributes instance for method chaining.
func (r *WrapAttributes) SetDirection(direction duit_props.Axis) *WrapAttributes {
	r.Direction = direction
	return r
}

//bindgen:exclude
func (r *WrapAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.CrossAxisAlignment == duit_props.CrossAxisAlignmentBaseline || r.CrossAxisAlignment == duit_props.CrossAxisAlignmentStretch {
		return errors.New("CrossAxisAlignment property cannot have 'baseline' or 'stretch' values for curret kind of attribute")
	}

	return nil
}
