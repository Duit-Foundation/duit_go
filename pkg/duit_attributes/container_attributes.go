package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// ContainerAttributes represents attributes for Container widget.
// https://api.flutter.dev/flutter/widgets/Container-class.html
type ContainerAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Width                             float32                    `json:"width,omitempty"`
	Height                            float32                    `json:"height,omitempty"`
	Color                             *duit_props.Color          `json:"color,omitempty"`
	ClipBehavior                      duit_props.Clip            `json:"clipBehavior,omitempty"`
	Decoration                        *duit_props.BoxDecoration  `json:"decoration,omitempty"`
	ForegroundDecoration              *duit_props.BoxDecoration  `json:"foregroundDecoration,omitempty"`
	Padding                           *duit_props.EdgeInsets     `json:"padding,omitempty"`
	Margin                            *duit_props.EdgeInsets     `json:"margin,omitempty"`
	Alignment                         duit_props.Alignment       `json:"alignment,omitempty"`
	TransformAlignment                duit_props.Alignment       `json:"transformAlignment,omitempty"`
	Constraints                       *duit_props.BoxConstraints `json:"constraints,omitempty"`
}

// NewContainerAttributes creates a new instance of ContainerAttributes.
func NewContainerAttributes() *ContainerAttributes {
	return &ContainerAttributes{}
}

// SetWidth sets the width property and returns the ContainerAttributes instance for method chaining.
func (r *ContainerAttributes) SetWidth(width float32) *ContainerAttributes {
	r.Width = width
	return r
}

// SetHeight sets the height property and returns the ContainerAttributes instance for method chaining.
func (r *ContainerAttributes) SetHeight(height float32) *ContainerAttributes {
	r.Height = height
	return r
}

// SetColor sets the color property and returns the ContainerAttributes instance for method chaining.
func (r *ContainerAttributes) SetColor(color *duit_props.Color) *ContainerAttributes {
	r.Color = color
	return r
}

// SetClipBehavior sets the clipBehavior property and returns the ContainerAttributes instance for method chaining.
func (r *ContainerAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *ContainerAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetDecoration sets the decoration property and returns the ContainerAttributes instance for method chaining.
func (r *ContainerAttributes) SetDecoration(decoration *duit_props.BoxDecoration) *ContainerAttributes {
	r.Decoration = decoration
	return r
}

// SetForegroundDecoration sets the foregroundDecoration property and returns the ContainerAttributes instance for method chaining.
func (r *ContainerAttributes) SetForegroundDecoration(foregroundDecoration *duit_props.BoxDecoration) *ContainerAttributes {
	r.ForegroundDecoration = foregroundDecoration
	return r
}

// SetPadding sets the padding property and returns the ContainerAttributes instance for method chaining.
func (r *ContainerAttributes) SetPadding(padding *duit_props.EdgeInsets) *ContainerAttributes {
	r.Padding = padding
	return r
}

// SetMargin sets the margin property and returns the ContainerAttributes instance for method chaining.
func (r *ContainerAttributes) SetMargin(margin *duit_props.EdgeInsets) *ContainerAttributes {
	r.Margin = margin
	return r
}

// SetAlignment sets the alignment property and returns the ContainerAttributes instance for method chaining.
func (r *ContainerAttributes) SetAlignment(alignment duit_props.Alignment) *ContainerAttributes {
	r.Alignment = alignment
	return r
}

// SetTransformAlignment sets the transformAlignment property and returns the ContainerAttributes instance for method chaining.
func (r *ContainerAttributes) SetTransformAlignment(transformAlignment duit_props.Alignment) *ContainerAttributes {
	r.TransformAlignment = transformAlignment
	return r
}

// SetConstraints sets the constraints property and returns the ContainerAttributes instance for method chaining.
func (r *ContainerAttributes) SetConstraints(constraints *duit_props.BoxConstraints) *ContainerAttributes {
	r.Constraints = constraints
	return r
}

//bindgen:exclude
func (r *ContainerAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
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

	if r.Color != nil {
		if err := r.Color.Validate(); err != nil {
			return err
		}
	}
	
	return nil
}
