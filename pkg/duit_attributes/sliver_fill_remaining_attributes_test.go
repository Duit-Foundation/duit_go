package duit_attributes_test

import (
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestSliverFillRemainingAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverFillRemainingAttributes{
		HasScrollBody:  duit_utils.BoolValue(true),
		FillOverscroll: duit_utils.BoolValue(false),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}
