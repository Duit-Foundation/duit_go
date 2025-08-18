package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestAnimatedScaleAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AnimatedScaleAttributes{
		Scale: duit_utils.Float32Value(1.0),
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestAnimatedScaleAttributes_Validate_ValidScaleValues(t *testing.T) {
	testCases := []struct {
		name  string
		scale float32
	}{
		{
			name:  "Normal scale (1.0)",
			scale: 1.0,
		},
		{
			name:  "Small scale (0.5)",
			scale: 0.5,
		},
		{
			name:  "Large scale (2.0)",
			scale: 2.0,
		},
		{
			name:  "Tiny scale (0.1)",
			scale: 0.1,
		},
		{
			name:  "Very large scale (5.0)",
			scale: 5.0,
		},
		{
			name:  "Minimal positive scale (0.001)",
			scale: 0.001,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedScaleAttributes{
				Scale: duit_utils.Float32Value(tc.scale),
				ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
					Duration: 300,
					Curve:    duit_animations.Ease,
				},
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for %s, got: %v", tc.name, err)
			}
		})
	}
}

func TestAnimatedScaleAttributes_Validate_InvalidScaleValues(t *testing.T) {
	testCases := []struct {
		name  string
		scale float32
	}{
		{
			name:  "Negative scale",
			scale: -1.0,
		},
		{
			name:  "Very negative scale",
			scale: -0.5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedScaleAttributes{
				Scale: duit_utils.Float32Value(tc.scale),
				ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
					Duration: 300,
					Curve:    duit_animations.Ease,
				},
			}

			err := attrs.Validate()
			if err == nil {
				t.Fatalf("expected error for invalid scale %f", tc.scale)
			}

			expected := "scale must be greater than 0 and positive"
			if err.Error() != expected {
				t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
			}
		})
	}
}

func TestAnimatedScaleAttributes_Validate_NilScale(t *testing.T) {
	attrs := &duit_attributes.AnimatedScaleAttributes{
		Scale: nil,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil scale")
	}

	expected := "scale property is required"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestAnimatedScaleAttributes_Validate_WithOptionalFields(t *testing.T) {
	attrs := &duit_attributes.AnimatedScaleAttributes{
		Scale:         duit_utils.Float32Value(1.5),
		Alignmen:      duit_alignment.Center,
		FilterQuality: "high",
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for attributes with optional fields, got:", err)
	}
}

// JSON Serialization tests for Tristate fields

func TestAnimatedScaleAttributes_MarshalJSON_ScaleField(t *testing.T) {
	attrs := &duit_attributes.AnimatedScaleAttributes{
		Scale: duit_utils.Float32Value(2.5),
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"scale":2.5`) {
		t.Fatalf("expected JSON to contain '\"scale\":2.5', got: %s", jsonStr)
	}
}

func TestAnimatedScaleAttributes_MarshalJSON_NilScale(t *testing.T) {
	attrs := &duit_attributes.AnimatedScaleAttributes{
		Scale: nil,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with nil scale, got:", err)
	}

	jsonStr := string(jsonData)
	// Check that omitempty works - nil fields should not be present
	if strings.Contains(jsonStr, `"scale"`) {
		t.Fatalf("expected JSON to NOT contain 'scale' field for nil value, got: %s", jsonStr)
	}
}

func TestAnimatedScaleAttributes_MarshalJSON_DifferentScaleValues(t *testing.T) {
	testCases := []struct {
		name     string
		scale    duit_utils.Tristate[float32]
		expected string
	}{
		{
			name:     "Normal scale",
			scale:    duit_utils.Float32Value(1.0),
			expected: `"scale":1`,
		},
		{
			name:     "Half scale",
			scale:    duit_utils.Float32Value(0.5),
			expected: `"scale":0.5`,
		},
		{
			name:     "Double scale",
			scale:    duit_utils.Float32Value(2.0),
			expected: `"scale":2`,
		},
		{
			name:     "Tiny scale",
			scale:    duit_utils.Float32Value(0.1),
			expected: `"scale":0.1`,
		},
		{
			name:     "Large scale",
			scale:    duit_utils.Float32Value(5.0),
			expected: `"scale":5`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedScaleAttributes{
				Scale: tc.scale,
				ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
					Duration: 300,
					Curve:    duit_animations.Ease,
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

func TestAnimatedScaleAttributes_Validate_NilImplicitAnimatable(t *testing.T) {
	attrs := &duit_attributes.AnimatedScaleAttributes{
		Scale:              duit_utils.Float32Value(1.0),
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

func TestAnimatedScaleAttributes_MarshalJSON_WithAllFields(t *testing.T) {
	attrs := &duit_attributes.AnimatedScaleAttributes{
		Scale:         duit_utils.Float32Value(0.8),
		Alignmen:      duit_alignment.BottomRight,
		FilterQuality: "medium",
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 400,
			Curve:    duit_animations.BounceOut,
		},
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with all fields, got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"scale":0.8`) {
		t.Fatalf("expected JSON to contain '\"scale\":0.8', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"alignment":"bottomRight"`) {
		t.Fatalf("expected JSON to contain '\"alignment\":\"bottomRight\"', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"filterQuality":"medium"`) {
		t.Fatalf("expected JSON to contain '\"filterQuality\":\"medium\"', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"duration":400`) {
		t.Fatalf("expected JSON to contain '\"duration\":400', got: %s", jsonStr)
	}
}
