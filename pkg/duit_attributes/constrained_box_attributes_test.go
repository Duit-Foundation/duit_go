package duit_attributes_test

import (
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
)

func TestConstrainedBoxAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.ConstrainedBoxAttributes{
		MinWidth:  100.0,
		MaxWidth:  200.0,
		MinHeight: 100.0,
		MaxHeight: 200.0,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}