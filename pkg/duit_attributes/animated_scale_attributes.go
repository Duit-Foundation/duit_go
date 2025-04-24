package duit_attributes

import (
	"encoding/json"
	"errors"

	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_alignment"
	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_painting"
)

type AnimatedScaleAttributes struct  {
	ValueReferenceHolder
	animations.ImplicitAnimatable
	ThemeConsumer
	Scale float32 `json:"scale"`
	Alignmen duit_alignment.Alignment `json:"alignment,omitempty"`
	FilterQuality duit_painting.FilterQuality `json:"filterQuality,omitempty"`
}

func (s *AnimatedScaleAttributes) MarshalJSON() ([]byte, error) {
	if s.Scale == 0 {
		return nil, errors.New("scale must be greater than 0")
	}

	return json.Marshal(*s)
}