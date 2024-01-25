package duit_attributes

import (
	"encoding/json"
	"errors"
)

type OpacityAttributes struct {
	Opacity float32 `json:"opacity"`
}

func (attributes *OpacityAttributes) MarshalJSON() ([]byte, error) {

	if attributes.Opacity >= 0 && attributes.Opacity <= 1 {
		return nil, errors.New("opacity must be between 0 and 1")
	}

	return json.Marshal(attributes)
}
