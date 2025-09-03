package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_cross_axis"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_main_axis"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type WrapAttributes struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	TextDirection      duit_props.TextDirection `json:"textDirection,omitempty"`
	VerticalDirection  duit_props.VerticalDirection        `json:"verticalDirection,omitempty"`
	Alignment          duit_main_axis.MainAxisAlignment   `json:"alignment,omitempty"`
	RunAlignment       duit_main_axis.MainAxisAlignment   `json:"runAlignment,omitempty"`
	CrossAxisAlignment duit_cross_axis.CrossAxisAlignment `json:"crossAxisAlignment,omitempty"`
	Spacing            float32                            `json:"spacing,omitempty"`
	RunSpacing         float32                            `json:"runSpacing,omitempty"`
	ClipBehavior       duit_props.Clip                     `json:"clipBehavior,omitempty"`
	Direction          duit_props.Axis                     `json:"direction,omitempty"`
}

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

	if r.CrossAxisAlignment == duit_cross_axis.Baseline || r.CrossAxisAlignment == duit_cross_axis.Stretch {
		return errors.New("CrossAxisAlignment property cannot have 'baseline' or 'stretch' values for curret kind of attribute")
	}

	return nil
}
