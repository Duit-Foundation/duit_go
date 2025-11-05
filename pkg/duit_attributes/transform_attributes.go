package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// TransformAttributes represents attributes for Transform widget.
// https://api.flutter.dev/flutter/widgets/Transform-class.html
type TransfromAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Type                              duit_props.TransformType `json:"type"`
	Data                              *duit_props.Transform    `json:"data"`
}

// NewTransformAttributes creates a new instance of TransformAttributes.
func NewTransformAttributes() *TransfromAttributes {
	return &TransfromAttributes{}
}

// SetType sets the type property and returns the TransformAttributes instance for method chaining.
func (r *TransfromAttributes) SetType(transformType duit_props.TransformType) *TransfromAttributes {
	r.Type = transformType
	return r
}

// SetData sets the data property and returns the TransformAttributes instance for method chaining.
func (r *TransfromAttributes) SetData(data *duit_props.Transform) *TransfromAttributes {
	r.Data = data
	return r
}

//bindgen:exclude
func (r *TransfromAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if len(r.Type) == 0 {
		return errors.New("type property is required")
	}

	if r.Data == nil {
		return errors.New("data property is required")
	}

	return nil
}
