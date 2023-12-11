package duit_attributes

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_decorations"
)

type DecoratedBoxAttributes[T duit_color.Color] struct {
	Decoration *duit_decoration.BoxDecoration[T] `json:"decoration,omitempty"`
}
