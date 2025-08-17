package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type AnimatedPhysicalModelAttributes struct {
	*ValueReferenceHolder
	*duit_animations.ImplicitAnimatable
	*ThemeConsumer
	Elevation          duit_utils.Tristate[float32]    `json:"elevation"`
	Color              *duit_color.ColorV2             `json:"color,omitempty"`
	ShadowColor        *duit_color.ColorV2             `json:"shadowColor,omitempty"`
	ClipBehavior       duit_clip.Clip                  `json:"clipBehavior,omitempty"`
	BorderRadius       *duit_decoration.BorderRadiusV2 `json:"borderRadius,omitempty"`
	Shape              duit_decoration.BoxShape        `json:"shape,omitempty"`
	AnimateColor       duit_utils.Tristate[bool]       `json:"animateColor,omitempty"`
	AnimateShadowColor duit_utils.Tristate[bool]       `json:"animateShadowColor,omitempty"`
}

func (a *AnimatedPhysicalModelAttributes) Validate() error {
	if err := a.ImplicitAnimatable.Validate(); err != nil {
		return err
	}

	if err := a.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := a.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if a.Elevation != nil {
		if *a.Elevation < 0 {
			return errors.New("elevation must be greater than 0")
		}
	} else {
		return errors.New("elevation property is required")
	}

	if a.Color != nil {
		if err := a.Color.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("color property is required")
	}

	if a.ShadowColor != nil {
		if err := a.ShadowColor.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("shadowColor property is required")
	}

	if a.BorderRadius != nil {
		if err := a.BorderRadius.Validate(); err != nil {
			return err
		}
	}

	return nil
}
