package duit_attributes_test

import (
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestSliverSafeAreaAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverSafeAreaAttributes[duit_props.EdgeInsetsAll]{
		Left:                      duit_utils.BoolValue(true),
		Top:                       duit_utils.BoolValue(false),
		Right:                     duit_utils.BoolValue(true),
		Bottom:                    duit_utils.BoolValue(false),
		MaintainBottomViewPadding: duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverSafeAreaAttributes_Validate_WithoutTristateFields(t *testing.T) {
	attrs := &duit_attributes.SliverSafeAreaAttributes[duit_props.EdgeInsetsAll]{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error when tristate fields are not set, got:", err)
	}
}
