package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
)

type AnimatedPaddingAttributes struct {
	*ValueReferenceHolder
	*animations.ImplicitAnimatable
	*ThemeConsumer
	Padding *duit_edge_insets.EdgeInsetsV2 `json:"padding,omitempty"`
}

func (a *AnimatedPaddingAttributes) Validate() error {
	if err := a.ImplicitAnimatable.Validate(); err != nil {
		return err
	}

	if err := a.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := a.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if a.Padding != nil {
		if err := a.Padding.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("padding property is required")
	}

	return nil
}
