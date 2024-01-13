package duit_attributes

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_transform"
)

type TransfromAttributes[T duit_transform.Transform] struct {
	ValueReferenceHolder
	Type duit_transform.TransfromType `json:"type"`
	Data *T                           `json:"data"`
}
