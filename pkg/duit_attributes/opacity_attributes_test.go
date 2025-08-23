package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestOpacityAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.OpacityAttributes{
		Opacity: duit_utils.Float32Value(0.5),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestOpacityAttributes_Validate_WithoutOpacity(t *testing.T) {
	attrs := &duit_attributes.OpacityAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when opacity is not set")
	}
}

func TestOpacityAttributes_Validate_ValidOpacityValues(t *testing.T) {
	testCases := []struct {
		name    string
		opacity float32
	}{
		{"Zero opacity", 0.0},
		{"Low opacity", 0.1},
		{"Half opacity", 0.5},
		{"High opacity", 0.9},
		{"Full opacity", 1.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.OpacityAttributes{
				Opacity: duit_utils.Float32Value(tc.opacity),
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for opacity %f, got: %v", tc.opacity, err)
			}
		})
	}
}

func TestOpacityAttributes_Validate_InvalidOpacityValues(t *testing.T) {
	testCases := []struct {
		name    string
		opacity float32
	}{
		{"Negative opacity", -0.1},
		{"Very high opacity", 2.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.OpacityAttributes{
				Opacity: duit_utils.Float32Value(tc.opacity),
			}

			err := attrs.Validate()
			if err == nil {
				t.Fatalf("expected error for opacity %f", tc.opacity)
			}
		})
	}
}