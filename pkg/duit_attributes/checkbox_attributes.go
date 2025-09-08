package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type CheckboxAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	Value         duit_utils.Tristate[bool]    `json:"value"`
	Autofocus     duit_utils.Tristate[bool]    `json:"autofocus,omitempty"`
	Tristate      duit_utils.Tristate[bool]    `json:"tristate,omitempty"`
	IsError       duit_utils.Tristate[bool]    `json:"isError,omitempty"`
	SplashRadius  float32                      `json:"splashRadius,omitempty"`
	SemanticLabel string                       `json:"semanticLabel,omitempty"`
	Side          *duit_props.BorderSide       `json:"side,omitempty"`
	VisualDensity *duit_props.VisualDensity    `json:"visualDensity,omitempty"`
	CheckColor    *duit_props.Color            `json:"checkColor,omitempty"`
	HoverColor    *duit_props.Color            `json:"hoverColor,omitempty"`
	ActiveColor   *duit_props.Color            `json:"activeColor,omitempty"`
	FocusColor    *duit_props.Color            `json:"focusColor,omitempty"`
	FillColor     *duit_props.WidgetStateColor `json:"fillColor,omitempty"`
	OverlayColor  *duit_props.WidgetStateColor `json:"overlayColor,omitempty"`
}

func (r *CheckboxAttributes) Validate() error {
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
