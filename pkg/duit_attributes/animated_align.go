package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// AnimatedAlignAttributes represents attributes for AnimatedAlign widget.
// See: https://api.flutter.dev/flutter/widgets/AnimatedAlign-class.html
type AnimatedAlignAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	Alignment                      duit_props.Alignment `json:"alignment,omitempty"`
	WidthFactor                    float32              `json:"widthFactor,omitempty"`
	HeightFactor                   float32              `json:"heightFactor,omitempty"`
}

//bindgen:exclude
func (r *AnimatedAlignAttributes) Validate() error {
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

	if len(r.Alignment) == 0 {
		return errors.New("alignment property is required")
	}

	return nil
}

// NewAnimatedAlignAttributes creates a new instance of AnimatedAlignAttributes.
func NewAnimatedAlignAttributes() *AnimatedAlignAttributes {
	return &AnimatedAlignAttributes{}
}

// SetAlignment sets the alignment property and returns the AnimatedAlignAttributes instance for method chaining.
func (r *AnimatedAlignAttributes) SetAlignment(alignment duit_props.Alignment) *AnimatedAlignAttributes {
	r.Alignment = alignment
	return r
}

// SetWidthFactor sets the widthFactor property and returns the AnimatedAlignAttributes instance for method chaining.
func (r *AnimatedAlignAttributes) SetWidthFactor(widthFactor float32) *AnimatedAlignAttributes {
	r.WidthFactor = widthFactor
	return r
}

// SetHeightFactor sets the heightFactor property and returns the AnimatedAlignAttributes instance for method chaining.
func (r *AnimatedAlignAttributes) SetHeightFactor(heightFactor float32) *AnimatedAlignAttributes {
	r.HeightFactor = heightFactor
	return r
}
