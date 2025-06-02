package duit_attributes

import (
	"encoding/json"
	
	animations "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"
)

type OffstageAttributes struct {
	animations.AnimatedPropertyOwner
	ValueReferenceHolder

	Offstage	*bool	`json:"offstage,omitempty"`
}

func (attributes *OffstageAttributes) MarshalJSON() ([]byte, error) {

	if attributes.Offstage == nil {
		var bPtr = duit_utils.BoolPtr(true)
		attributes.Offstage = bPtr
	}

	return json.Marshal(*attributes)
}