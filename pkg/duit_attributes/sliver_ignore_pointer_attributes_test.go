package duit_attributes_test

import (
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestSliverIgnorePointerAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverIgnorePointerAttributes{
		Ignoring: duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverIgnorePointerAttributes_Validate_WithoutIgnoring(t *testing.T) {
	attrs := &duit_attributes.SliverIgnorePointerAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when ignoring is not set")
	}
}

func TestSliverIgnorePointerAttributes_Validate_BooleanValues(t *testing.T) {
	testCases := []struct {
		name     string
		ignoring bool
	}{
		{"Ignoring true", true},
		{"Ignoring false", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.SliverIgnorePointerAttributes{
				Ignoring: duit_utils.BoolValue(tc.ignoring),
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for ignoring %t, got: %v", tc.ignoring, err)
			}
		})
	}
}
