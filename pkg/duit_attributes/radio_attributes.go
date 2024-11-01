package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_material"
)

type PrimitiveValue interface {
	string | int | uint | uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64 | float32 | float64 | bool
}

type RadioAttributes[TValue PrimitiveValue, TColor duit_color.Color] struct {
	ValueReferenceHolder
	Value                 TValue                                      `json:"value"`
	Toggleable            bool                                        `json:"toggleable,omitempty"`
	Autofocus             bool                                        `json:"autofocus,omitempty"`
	ActiveColor           TColor                                      `json:"activeColor,omitempty"`
	FocusColor            TColor                                      `json:"focusColor,omitempty"`
	HoverColor            TColor                                      `json:"hoverColor,omitempty"`
	FillColor             duit_material.MaterialStateProperty[TColor] `json:"fillColor,omitempty"`
	OverlayColor          duit_material.MaterialStateProperty[TColor] `json:"overlayColor,omitempty"`
	SplashRadius          float32                                     `json:"splashRadius,omitempty"`
	VisualDensity         *duit_flex.VisualDensity                    `json:"visualDensity,omitempty"`
	MaterialTapTargetSize duit_material.MaterialTapTargetSize         `json:"materialTapTargetSize,omitempty"`
}

type RadioGroupContextAttributes[TValue PrimitiveValue] struct {
	ValueReferenceHolder
	GroupValue TValue `json:"groupValue"`
}
