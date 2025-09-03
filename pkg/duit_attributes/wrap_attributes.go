package duit_attributes

import (
	"errors"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type WrapAttributes struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	TextDirection      duit_props.TextDirection `json:"textDirection,omitempty"`
	VerticalDirection  duit_props.VerticalDirection        `json:"verticalDirection,omitempty"`
	Alignment          duit_props.MainAxisAlignment   `json:"alignment,omitempty"`
	RunAlignment       duit_props.MainAxisAlignment   `json:"runAlignment,omitempty"`
	CrossAxisAlignment duit_props.CrossAxisAlignment `json:"crossAxisAlignment,omitempty"`
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

	if r.CrossAxisAlignment == duit_props.CrossAxisAlignmentBaseline || r.CrossAxisAlignment == duit_props.CrossAxisAlignmentStretch {
		return errors.New("CrossAxisAlignment property cannot have 'baseline' or 'stretch' values for curret kind of attribute")
	}

	return nil
}
