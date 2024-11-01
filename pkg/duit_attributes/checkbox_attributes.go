package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/pkg/duit_attributes/duit_material"
)

type CheckboxAttributes[T duit_color.Color] struct {
	ValueReferenceHolder
	Value         bool                                    `json:"value"`
	Autofocus     bool                                    `json:"autofocus,omitempty"`
	Tristate      bool                                    `json:"tristate,omitempty"`
	IsError       bool                                    `json:"isError,omitempty"`
	SplashRadius  float32                                 `json:"splashRadius,omitempty"`
	SemanticLabel string                                  `json:"semanticLabel,omitempty"`
	Side          *duit_decoration.BorderSide[T]          `json:"side,omitempty"`
	VisualDensity *duit_flex.VisualDensity                `json:"visualDensity,omitempty"`
	CheckColor    T                                       `json:"checkColor,omitempty"`
	HoverColor    T                                       `json:"hoverColor,omitempty"`
	ActiveColor   T                                       `json:"activeColor,omitempty"`
	FocusColor    T                                       `json:"focusColor,omitempty"`
	FillColor     *duit_material.MaterialStateProperty[T] `json:"fillColor,omitempty"`
	OverlayColor  *duit_material.MaterialStateProperty[T] `json:"overlayColor,omitempty"`
}
