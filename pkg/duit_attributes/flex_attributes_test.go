package duit_attributes_test

import (
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
)

func TestFlexAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.FlexAttributes{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}