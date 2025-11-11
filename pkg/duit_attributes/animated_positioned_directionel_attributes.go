package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AnimatedPositionedDirectionalAttributes struct {
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*duit_props.ImplicitAnimatable    `gen:"wrap"`
	*ValueReferenceHolder             `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Start                             duit_utils.Tristate[float32] `json:"start,omitempty"`
	End                               duit_utils.Tristate[float32] `json:"end,omitempty"`
	Top                               duit_utils.Tristate[float32] `json:"top,omitempty"`
	Bottom                            duit_utils.Tristate[float32] `json:"bottom,omitempty"`
	Width                             duit_utils.Tristate[float32] `json:"Width,omitempty"`
	Height                            duit_utils.Tristate[float32] `json:"Height,omitempty"`
}

func NewAnimatedPositionedDirectionalAttributes() *AnimatedPositionedDirectionalAttributes {
	return &AnimatedPositionedDirectionalAttributes{}
}

// SetStart sets the start property and returns the AnimatedPositionedDirectionalAttributes instance for method chaining.
func (r *AnimatedPositionedDirectionalAttributes) SetStart(start float32) *AnimatedPositionedDirectionalAttributes {
	r.Start = duit_utils.Float32Value(start)
	return r
}

// SetEnd sets the end property and returns the AnimatedPositionedDirectionalAttributes instance for method chaining.
func (r *AnimatedPositionedDirectionalAttributes) SetEnd(end float32) *AnimatedPositionedDirectionalAttributes {
	r.End = duit_utils.Float32Value(end)
	return r
}

// SetWidth sets the width property and returns the AnimatedPositionedDirectionalAttributes instance for method chaining.
func (r *AnimatedPositionedDirectionalAttributes) SetWidth(width float32) *AnimatedPositionedDirectionalAttributes {
	r.Width = duit_utils.Float32Value(width)
	return r
}

// SetHeight sets the height property and returns the AnimatedPositionedDirectionalAttributes instance for method chaining.
func (r *AnimatedPositionedDirectionalAttributes) SetHeight(height float32) *AnimatedPositionedDirectionalAttributes {
	r.Height = duit_utils.Float32Value(height)
	return r
}

// SetTop sets the left property and returns the AnimatedPositionedDirectionalAttributes instance for method chaining.
func (r *AnimatedPositionedDirectionalAttributes) SetTop(top float32) *AnimatedPositionedDirectionalAttributes {
	r.Top = duit_utils.Float32Value(top)
	return r
}

// SetBottom sets the bottom property and returns the AnimatedPositionedDirectionalAttributes instance for method chaining.
func (r *AnimatedPositionedDirectionalAttributes) SetBottom(bottom float32) *AnimatedPositionedDirectionalAttributes {
	r.Bottom = duit_utils.Float32Value(bottom)
	return r
}

//bindgen:exclude
func (r *AnimatedPositionedDirectionalAttributes) Validate() error {
	if r.ImplicitAnimatable != nil {
		if err := r.ImplicitAnimatable.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("implicitAnimatable property is required on implicit animated widgets")
	}

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
