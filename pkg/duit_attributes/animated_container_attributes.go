package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
)

type AnimatedContainerAttributes[TInsets duit_props.EdgeInsets, TColor duit_props.Color] struct {
	*ValueReferenceHolder
	*animations.ImplicitAnimatable
	*ThemeConsumer
	Width                float32                                `json:"width,omitempty"`
	Height               float32                                `json:"height,omitempty"`
	Color                TColor                                 `json:"color,omitempty"`
	ClipBehavior         duit_props.Clip                         `json:"clipBehavior,omitempty"`
	Decoration           *duit_props.BoxDecoration[TColor] `json:"decoration,omitempty"`
	ForegroundDecoration *duit_props.BoxDecoration[TColor] `json:"foregroundDecoration,omitempty"`
	Padding              TInsets                                `json:"padding,omitempty"`
	Margin               TInsets                                `json:"margin,omitempty"`
	Alignment            duit_props.Alignment                   `json:"alignment,omitempty"`
	TransformAlignment   duit_props.Alignment                   `json:"transformAlignment,omitempty"`
	Constraints          *duit_flex.BoxConstraints              `json:"constraints,omitempty"`
}

// Делать валидацию внутренних свойств, где это требуется
func (r *AnimatedContainerAttributes[TInsets, TColor]) Validate() error {
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
