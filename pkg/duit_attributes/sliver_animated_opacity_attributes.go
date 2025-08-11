package duit_attributes

import (
	"encoding/json"
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
)

type SliverAnimatedpacityAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	animations.ImplicitAnimatable
	SliverProps
	Opacity float32 `json:"opacity,omitempty"`
}

func (attributes *SliverAnimatedpacityAttributes) MarshalJSON() ([]byte, error) {

	if attributes.Opacity >= 0 && attributes.Opacity <= 1 {
		return json.Marshal(*attributes)
	}

	return nil, errors.New("opacity must be between 0 and 1")
}