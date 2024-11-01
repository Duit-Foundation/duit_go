package duit_attributes

import "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"

type ColoredBoxAttributes[T duit_color.Color] struct {
	ValueReferenceHolder
	Color T `json:"color,omitempty"`
}
