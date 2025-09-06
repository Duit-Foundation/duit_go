package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type PrimitiveValue interface {
	string | int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64 | bool
}

type RadioAttributes[TValue PrimitiveValue] struct {
	*ValueReferenceHolder
	*duit_props.AnimatedPropertyOwner
	*ThemeConsumer
	Value                 duit_utils.Tristate[TValue]                         `json:"value,omitempty"`
	Toggleable            bool                                                `json:"toggleable,omitempty"`
	Autofocus             bool                                                `json:"autofocus,omitempty"`
	ActiveColor           *duit_props.Color                                   `json:"activeColor,omitempty"`
	FocusColor            *duit_props.Color                                   `json:"focusColor,omitempty"`
	HoverColor            *duit_props.Color                                   `json:"hoverColor,omitempty"`
	FillColor             duit_props.MaterialStateProperty[*duit_props.Color] `json:"fillColor,omitempty"`
	OverlayColor          duit_props.MaterialStateProperty[*duit_props.Color] `json:"overlayColor,omitempty"`
	SplashRadius          float32                                             `json:"splashRadius,omitempty"`
	VisualDensity         *duit_props.VisualDensity                           `json:"visualDensity,omitempty"`
	MaterialTapTargetSize duit_props.MaterialTapTargetSize                    `json:"materialTapTargetSize,omitempty"`
}

func (r *RadioAttributes[TValue]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Value == nil {
		return errors.New("value property is required")
	}

	return nil
}

type RadioGroupContextAttributes[TValue PrimitiveValue] struct {
	*ValueReferenceHolder
	GroupValue duit_utils.Tristate[TValue] `json:"groupValue,omitempty"`
}

func (r *RadioGroupContextAttributes[TValue]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.GroupValue == nil {
		return errors.New("groupValue property is required")
	}

	return nil
}
