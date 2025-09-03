package duit_attributes

import (
	"errors"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type CheckboxAttributes[T duit_props.Color] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	Value         duit_utils.Tristate[bool]               `json:"value"`
	Autofocus     duit_utils.Tristate[bool]               `json:"autofocus,omitempty"`
	Tristate      duit_utils.Tristate[bool]               `json:"tristate,omitempty"`
	IsError       duit_utils.Tristate[bool]               `json:"isError,omitempty"`
	SplashRadius  float32                                 `json:"splashRadius,omitempty"`
	SemanticLabel string                                  `json:"semanticLabel,omitempty"`
	Side          *duit_props.BorderSide[T]          `json:"side,omitempty"`
	VisualDensity *duit_props.VisualDensity                `json:"visualDensity,omitempty"`
	CheckColor    T                                       `json:"checkColor,omitempty"`
	HoverColor    T                                       `json:"hoverColor,omitempty"`
	ActiveColor   T                                       `json:"activeColor,omitempty"`
	FocusColor    T                                       `json:"focusColor,omitempty"`
	FillColor     *duit_props.MaterialStateProperty[T] `json:"fillColor,omitempty"`
	OverlayColor  *duit_props.MaterialStateProperty[T] `json:"overlayColor,omitempty"`
}

func (r *CheckboxAttributes[T]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.Value == nil {
		return errors.New("value property is required")
	}

	return nil
}
