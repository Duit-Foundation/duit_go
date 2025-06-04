package duit_attributes

import (
	"encoding/json"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_utils"
)

type AbsorbPointerAttributes struct {
	ValueReferenceHolder
	Absorbing *bool `json:"absorbing,omitempty"`
}

func (attributes *AbsorbPointerAttributes) MarshalJSON() ([]byte, error) {
	if attributes.Absorbing == nil {
		var bPtr = duit_utils.BoolPtr(true)
		attributes.Absorbing = bPtr
	}

	return json.Marshal(*attributes)
}