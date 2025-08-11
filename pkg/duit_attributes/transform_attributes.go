package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_transform"
)

type TransfromAttributes[T duit_transform.Transform] struct {
	ValueReferenceHolder
	animations.AnimatedPropertyOwner
	ThemeConsumer
	Type duit_transform.TransfromType `json:"type"`
	Data *T                           `json:"data"`
}
