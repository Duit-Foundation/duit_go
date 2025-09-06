package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type AnimatedContainerAttributes struct {
	*ValueReferenceHolder
	*duit_props.ImplicitAnimatable
	*ThemeConsumer
	Width                float32                    `json:"width,omitempty"`
	Height               float32                    `json:"height,omitempty"`
	Color                *duit_props.Color          `json:"color,omitempty"`
	ClipBehavior         duit_props.Clip            `json:"clipBehavior,omitempty"`
	Decoration           *duit_props.BoxDecoration  `json:"decoration,omitempty"`
	ForegroundDecoration *duit_props.BoxDecoration  `json:"foregroundDecoration,omitempty"`
	Padding              *duit_props.EdgeInsets     `json:"padding,omitempty"`
	Margin               *duit_props.EdgeInsets     `json:"margin,omitempty"`
	Alignment            duit_props.Alignment       `json:"alignment,omitempty"`
	TransformAlignment   duit_props.Alignment       `json:"transformAlignment,omitempty"`
	Constraints          *duit_props.BoxConstraints `json:"constraints,omitempty"`
}

// Делать валидацию внутренних свойств, где это требуется
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

	return nil
}
