package duit_attributes

import (
	"errors"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type SliverPaddingAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*duit_props.AnimatedPropertyOwner
	*SliverProps
	Padding *duit_props.EdgeInsetsV2 `json:"padding,omitempty"`
}

func (r *SliverPaddingAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Padding == nil {
		return errors.New("padding property is required")
	}

	return nil
}
