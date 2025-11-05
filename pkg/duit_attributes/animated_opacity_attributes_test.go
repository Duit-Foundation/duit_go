package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestAnimatedOpacityAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AnimatedOpacityAttributes{
		Opacity: duit_utils.Float32Value(0.5),
		ImplicitAnimatable: &duit_props.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_props.CurveEase,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestAnimatedOpacityAttributes_Validate_ValidOpacityValues(t *testing.T) {
	testCases := []struct {
		name    string
		opacity float32
	}{
		{
			name:    "Minimum opacity (0.0)",
			opacity: 0.0,
		},
		{
			name:    "Maximum opacity (1.0)",
			opacity: 1.0,
		},
		{
			name:    "Middle opacity (0.5)",
			opacity: 0.5,
		},
		{
			name:    "Low opacity (0.1)",
			opacity: 0.1,
		},
		{
			name:    "High opacity (0.9)",
			opacity: 0.9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedOpacityAttributes{
				Opacity: duit_utils.Float32Value(tc.opacity),
				ImplicitAnimatable: &duit_props.ImplicitAnimatable{
					Duration: 300,
					Curve:    duit_props.CurveEase,
				},
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for %s, got: %v", tc.name, err)
			}
		})
	}
}

func TestAnimatedOpacityAttributes_Validate_InvalidOpacityValues(t *testing.T) {
	testCases := []struct {
		name    string
		opacity float32
	}{
		{
			name:    "Negative opacity",
			opacity: -0.1,
		},
		{
			name:    "Opacity greater than 1",
			opacity: 1.1,
		},
		{
			name:    "Very negative opacity",
			opacity: -1.0,
		},
		{
			name:    "Very high opacity",
			opacity: 2.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedOpacityAttributes{
				Opacity: duit_utils.Float32Value(tc.opacity),
				ImplicitAnimatable: &duit_props.ImplicitAnimatable{
					Duration: 300,
					Curve:    duit_props.CurveEase,
				},
			}

			err := attrs.Validate()
			if err == nil {
				t.Fatalf("expected error for invalid opacity %f", tc.opacity)
			}
		})
	}
}

func TestAnimatedOpacityAttributes_Validate_NilOpacity(t *testing.T) {
	attrs := &duit_attributes.AnimatedOpacityAttributes{
		Opacity: nil,
		ImplicitAnimatable: &duit_props.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_props.CurveEase,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for nil opacity (optional field), got:", err)
	}
}

// JSON Serialization tests for Tristate fields

func TestAnimatedOpacityAttributes_MarshalJSON_OpacityField(t *testing.T) {
	attrs := &duit_attributes.AnimatedOpacityAttributes{
		Opacity: duit_utils.Float32Value(0.75),
		ImplicitAnimatable: &duit_props.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_props.CurveEase,
		},
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"opacity":0.75`) {
		t.Fatalf("expected JSON to contain '\"opacity\":0.75', got: %s", jsonStr)
	}
}

func TestAnimatedOpacityAttributes_MarshalJSON_NilOpacity(t *testing.T) {
	attrs := &duit_attributes.AnimatedOpacityAttributes{
		Opacity: nil,
		ImplicitAnimatable: &duit_props.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_props.CurveEase,
		},
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with nil opacity, got:", err)
	}

	jsonStr := string(jsonData)
	// Check that omitempty works - nil fields should not be present
	if strings.Contains(jsonStr, `"opacity"`) {
		t.Fatalf("expected JSON to NOT contain 'opacity' field for nil value, got: %s", jsonStr)
	}
}

func TestAnimatedOpacityAttributes_MarshalJSON_DifferentOpacityValues(t *testing.T) {
	testCases := []struct {
		name     string
		opacity  duit_utils.Tristate[float32]
		expected string
	}{
		{
			name:     "Zero opacity",
			opacity:  duit_utils.Float32Value(0.0),
			expected: `"opacity":0`,
		},
		{
			name:     "Full opacity",
			opacity:  duit_utils.Float32Value(1.0),
			expected: `"opacity":1`,
		},
		{
			name:     "Half opacity",
			opacity:  duit_utils.Float32Value(0.5),
			expected: `"opacity":0.5`,
		},
		{
			name:     "Quarter opacity",
			opacity:  duit_utils.Float32Value(0.25),
			expected: `"opacity":0.25`,
		},
		{
			name:     "High precision opacity",
			opacity:  duit_utils.Float32Value(0.123),
			expected: `"opacity":0.123`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedOpacityAttributes{
				Opacity: tc.opacity,
				ImplicitAnimatable: &duit_props.ImplicitAnimatable{
					Duration: 300,
					Curve:    duit_props.CurveEase,
				},
			}

			jsonData, err := json.Marshal(attrs)
			if err != nil {
				t.Fatalf("expected no error for MarshalJSON() in %s, got: %v", tc.name, err)
			}

			jsonStr := string(jsonData)
			if !strings.Contains(jsonStr, tc.expected) {
				t.Fatalf("expected JSON to contain '%s' in %s, got: %s", tc.expected, tc.name, jsonStr)
			}
		})
	}
}

// Tests for ImplicitAnimatable embedded struct

func TestAnimatedOpacityAttributes_Validate_NilImplicitAnimatable(t *testing.T) {
	attrs := &duit_attributes.AnimatedOpacityAttributes{
		Opacity:            duit_utils.Float32Value(0.5),
		ImplicitAnimatable: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil ImplicitAnimatable")
	}

	expected := "implicitAnimatable property is required on implicit animated widgets"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}


