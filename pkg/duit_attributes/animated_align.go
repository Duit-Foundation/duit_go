package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
)

type AnimatedAlignAttributes struct {
	*ValueReferenceHolder
	*animations.ImplicitAnimatable
	*ThemeConsumer
	Alignment    duit_alignment.Alignment `json:"alignment,omitempty"`
	WidthFactor  float32                  `json:"widthFactor,omitempty"`
	HeightFactor float32                  `json:"heightFactor,omitempty"`
}

func (r *AnimatedAlignAttributes) Validate() error {
	if err := r.ImplicitAnimatable.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if len(r.Alignment) == 0 {
		return errors.New("alignment property is required")
	}

	return nil
}
