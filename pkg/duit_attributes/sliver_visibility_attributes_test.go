package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestSliverVisibilityAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverVisibilityAttributes{
		Visible:               duit_utils.BoolValue(true),
		MaintainState:         duit_utils.BoolValue(false),
		MaintainAnimation:     duit_utils.BoolValue(true),
		MaintainSize:          duit_utils.BoolValue(false),
		MaintainSemantics:     duit_utils.BoolValue(true),
		MaintainInteractivity: duit_utils.BoolValue(false),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverVisibilityAttributes_Validate_WithoutVisible(t *testing.T) {
	attrs := &duit_attributes.SliverVisibilityAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when visible is not set")
	}
}

func TestSliverVisibilityAttributes_Validate_BooleanValues(t *testing.T) {
	testCases := []struct {
		name    string
		visible bool
	}{
		{"Visible true", true},
		{"Visible false", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.SliverVisibilityAttributes{
				Visible: duit_utils.BoolValue(tc.visible),
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for visible %t, got: %v", tc.visible, err)
			}
		})
	}
}

func TestSliverVisibilityAttributes_Validate_WithReplacementSliver(t *testing.T) {
	replacementSliver := &duit_core.DuitElementModel{}

	attrs := &duit_attributes.SliverVisibilityAttributes{
		Visible:           duit_utils.BoolValue(false),
		ReplacementSliver: replacementSliver,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error with replacement sliver, got:", err)
	}
}
