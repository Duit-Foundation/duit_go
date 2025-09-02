package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_painting"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AnimatedScaleAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*animations.ImplicitAnimatable
	Scale         duit_utils.Tristate[float32] `json:"scale,omitempty"`
	Alignmen      duit_props.Alignment     `json:"alignment,omitempty"`
	FilterQuality duit_painting.FilterQuality  `json:"filterQuality,omitempty"`
}

func (r *AnimatedScaleAttributes) Validate() error {
	if r.ImplicitAnimatable != nil {
		if err := r.ImplicitAnimatable.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("implicitAnimatable property is required on implicit animated widgets")
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	var scaleV = r.Scale

	if scaleV == nil {
		return errors.New("scale property is required")
	} else {
		if *scaleV <= 0  {
			return errors.New("scale must be greater than 0 and positive")
		}
	}

	return nil
}
