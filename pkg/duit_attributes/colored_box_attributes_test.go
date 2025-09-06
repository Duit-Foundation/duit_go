package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestColoredBoxAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.ColoredBoxAttributes{
		Color: duit_props.NewColorString("#FF0000"),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestColoredBoxAttributes_Validate_NilColor(t *testing.T) {
	attrs := &duit_attributes.ColoredBoxAttributes{
		Color: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil color, got nil")
	}
}

