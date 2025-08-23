package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
)

func TestPositionedAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.PositionedAttributes{
		Left:   10.0,
		Top:    20.0,
		Right:  15.0,
		Bottom: 25.0,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}