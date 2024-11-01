package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_decorations"
)

type DecoratedBoxAttributes[T duit_color.Color] struct {
	ValueReferenceHolder
	Decoration *duit_decoration.BoxDecoration[T] `json:"decoration,omitempty"`
}
