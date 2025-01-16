package duit_attributes

import (
	"encoding/json"
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_utils"
)

type AnimatedOpacityAttributes struct {
	ValueReferenceHolder
	animations.ImplicitAnimatable
	Opacity float32 `json:"opacity"`
}

func (attributes *AnimatedOpacityAttributes) MarshalJSON() ([]byte, error) {
	err := duit_utils.CheckActionType(attributes.ImplicitAnimatable.OnEnd)

	if err != nil {
		return nil, err
	}

	if attributes.Opacity >= 0 && attributes.Opacity <= 1 {
		return json.Marshal(*attributes)
	}

	return nil, errors.New("opacity must be between 0 and 1")
}