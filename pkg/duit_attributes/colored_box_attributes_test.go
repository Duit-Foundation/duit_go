package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestColoredBoxAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.ColoredBoxAttributes[duit_props.ColorString]{
		Color: duit_props.ColorString("#FF0000"),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}