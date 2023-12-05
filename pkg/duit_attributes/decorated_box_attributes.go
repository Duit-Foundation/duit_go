package duit_attributes

import (
	duit_decoration "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_decorations"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"
)

type DecoratedBoxAttributes[T duit_utility_styles.Color] struct {
	Decoration duit_decoration.BoxDecoration[T] `json:"decoration,omitempty"`
}
