package duit_attributes

import (
	duit_decoration "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_decorations"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_material"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"
)

type CheckboxAttributes[T duit_utility_styles.Color] struct {
	Value         bool                                   `json:"value"`
	Autofocus     bool                                   `json:"autofocus,omitempty"`
	Tristate      bool                                   `json:"tristate,omitempty"`
	IsError       bool                                   `json:"isError,omitempty"`
	SplashRadius  float32                                `json:"splashRadius,omitempty"`
	SemanticLabel string                                 `json:"semanticLabel,omitempty"`
	Side          duit_decoration.BorderSide[T]          `json:"side,omitempty"`
	VisualDensity duit_flex.VisualDensity                `json:"visualDensity,omitempty"`
	CheckColor    T                                      `json:"checkColor,omitempty"`
	HoverColor    T                                      `json:"hoverColor,omitempty"`
	ActiveColor   T                                      `json:"activeColor,omitempty"`
	FocusColor    T                                      `json:"focusColor,omitempty"`
	FillColor     duit_material.MaterialStateProperty[T] `json:"fillColor,omitempty"`
	OverlayColor  duit_material.MaterialStateProperty[T] `json:"overlayColor,omitempty"`
}
