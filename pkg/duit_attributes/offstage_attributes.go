package duit_attributes

import (
	"encoding/json"
	
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type OffstageAttributes struct {
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