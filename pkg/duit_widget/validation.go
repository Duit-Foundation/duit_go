package duit_widget

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"

func checkAttributes(attributes duit_attributes.Validatable) {
	if attributes == nil {
		panic("attributes must not be nil")
	}

	if err := attributes.Validate(); err != nil {
		panic(err)
	}
}