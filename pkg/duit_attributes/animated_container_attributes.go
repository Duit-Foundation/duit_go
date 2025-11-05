package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// AnimatedContainerAttributes represents attributes for AnimatedContainer widget.
// https://api.flutter.dev/flutter/widgets/AnimatedContainer-class.html
type AnimatedContainerAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	Width                          float32                    `json:"width,omitempty"`
	Height                         float32                    `json:"height,omitempty"`
	Color                          *duit_props.Color          `json:"color,omitempty"`
	ClipBehavior                   duit_props.Clip            `json:"clipBehavior,omitempty"`
	Decoration                     *duit_props.BoxDecoration  `json:"decoration,omitempty"`
	ForegroundDecoration           *duit_props.BoxDecoration  `json:"foregroundDecoration,omitempty"`
	Padding                        *duit_props.EdgeInsets     `json:"padding,omitempty"`
	Margin                         *duit_props.EdgeInsets     `json:"margin,omitempty"`
	Alignment                      duit_props.Alignment       `json:"alignment,omitempty"`
	TransformAlignment             duit_props.Alignment       `json:"transformAlignment,omitempty"`
	Constraints                    *duit_props.BoxConstraints `json:"constraints,omitempty"`
}

// NewAnimatedContainerAttributes creates a new instance of AnimatedContainerAttributes.
func NewAnimatedContainerAttributes() *AnimatedContainerAttributes {
	return &AnimatedContainerAttributes{}
}

// SetWidth sets the width property and returns the AnimatedContainerAttributes instance for method chaining.
func (r *AnimatedContainerAttributes) SetWidth(width float32) *AnimatedContainerAttributes {
	r.Width = width
	return r
}

// SetHeight sets the height property and returns the AnimatedContainerAttributes instance for method chaining.
func (r *AnimatedContainerAttributes) SetHeight(height float32) *AnimatedContainerAttributes {
	r.Height = height
	return r
}

// SetColor sets the color property and returns the AnimatedContainerAttributes instance for method chaining.
func (r *AnimatedContainerAttributes) SetColor(color *duit_props.Color) *AnimatedContainerAttributes {
	r.Color = color
	return r
}

// SetClipBehavior sets the clipBehavior property and returns the AnimatedContainerAttributes instance for method chaining.
func (r *AnimatedContainerAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *AnimatedContainerAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetDecoration sets the decoration property and returns the AnimatedContainerAttributes instance for method chaining.
func (r *AnimatedContainerAttributes) SetDecoration(decoration *duit_props.BoxDecoration) *AnimatedContainerAttributes {
	r.Decoration = decoration
	return r
}

// SetForegroundDecoration sets the foregroundDecoration property and returns the AnimatedContainerAttributes instance for method chaining.
func (r *AnimatedContainerAttributes) SetForegroundDecoration(foregroundDecoration *duit_props.BoxDecoration) *AnimatedContainerAttributes {
	r.ForegroundDecoration = foregroundDecoration
	return r
}

// SetPadding sets the padding property and returns the AnimatedContainerAttributes instance for method chaining.
func (r *AnimatedContainerAttributes) SetPadding(padding *duit_props.EdgeInsets) *AnimatedContainerAttributes {
	r.Padding = padding
	return r
}

// SetMargin sets the margin property and returns the AnimatedContainerAttributes instance for method chaining.
func (r *AnimatedContainerAttributes) SetMargin(margin *duit_props.EdgeInsets) *AnimatedContainerAttributes {
	r.Margin = margin
	return r
}

// SetAlignment sets the alignment property and returns the AnimatedContainerAttributes instance for method chaining.
func (r *AnimatedContainerAttributes) SetAlignment(alignment duit_props.Alignment) *AnimatedContainerAttributes {
	r.Alignment = alignment
	return r
}

// SetTransformAlignment sets the transformAlignment property and returns the AnimatedContainerAttributes instance for method chaining.
func (r *AnimatedContainerAttributes) SetTransformAlignment(transformAlignment duit_props.Alignment) *AnimatedContainerAttributes {
	r.TransformAlignment = transformAlignment
	return r
}

// SetConstraints sets the constraints property and returns the AnimatedContainerAttributes instance for method chaining.
func (r *AnimatedContainerAttributes) SetConstraints(constraints *duit_props.BoxConstraints) *AnimatedContainerAttributes {
	r.Constraints = constraints
	return r
}

//bindgen:exclude
func (r *AnimatedContainerAttributes) Validate() error {
	if r.ImplicitAnimatable != nil {
		if err := r.ImplicitAnimatable.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("implicitAnimatable property is required on implicit animated widgets")
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Color != nil {
		if err := r.Color.Validate(); err != nil {
			return err
		}
	}

	if r.Padding != nil {
		if err := r.Padding.Validate(); err != nil {
			return err
		}
	}

	if r.Margin != nil {
		if err := r.Margin.Validate(); err != nil {
			return err
		}
	}

	if r.Constraints != nil {
		if err := r.Constraints.Validate(); err != nil {
			return err
		}
	}

	return nil
}
