package duit_attributes_test

import (
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestSliverOffstageAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverOffstageAttributes{
		Offstage: duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverOffstageAttributes_Validate_WithoutOffstage(t *testing.T) {
	attrs := &duit_attributes.SliverOffstageAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when offstage is not set")
	}
}

func TestSliverOffstageAttributes_Validate_BooleanValues(t *testing.T) {
	testCases := []struct {
		name     string
		offstage bool
	}{
		{"Offstage true", true},
		{"Offstage false", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.SliverOffstageAttributes{
				Offstage: duit_utils.BoolValue(tc.offstage),
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for offstage %t, got: %v", tc.offstage, err)
			}
		})
	}
}
