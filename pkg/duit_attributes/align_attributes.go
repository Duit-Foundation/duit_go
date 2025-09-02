package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type AlignAttributes struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	Alignment    duit_props.Alignment `json:"alignment,omitempty"`
	WidthFactor  float32              `json:"widthFactor,omitempty"`
	HeightFactor float32              `json:"heightFactor,omitempty"`
}

func (r *AlignAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if len(r.Alignment) == 0 {
		return errors.New("alignment property is required")
	}

	return nil
}
