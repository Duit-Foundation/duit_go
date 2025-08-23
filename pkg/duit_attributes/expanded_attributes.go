package duit_attributes

import "errors"

type ExpandedAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	Flex uint16 `json:"flex,omitempty"`
}

func (r *ExpandedAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Flex < 1 {
		return errors.New("flex must be greater than 0")
	}

	return nil
}
