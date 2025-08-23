package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestAnimatedRotationAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AnimatedRotationAttributes{
		Turns: duit_utils.Float32Value(3.14),
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

func TestAnimatedRotationAttributes_Validate_NilTurns(t *testing.T) {
	attrs := &duit_attributes.AnimatedRotationAttributes{
		Turns: nil,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil turns")
	}

	expected := "turns property is required"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestAnimatedRotationAttributes_Validate_DifferentTurnsValues(t *testing.T) {
	testCases := []struct {
		name  string
		turns duit_utils.Tristate[float32]
	}{
		{
			name:  "Zero turns",
			turns: duit_utils.Float32Value(0.0),
		},
		{
			name:  "Positive turns",
			turns: duit_utils.Float32Value(1.5),
		},
		{
			name:  "Negative turns",
			turns: duit_utils.Float32Value(-2.5),
		},
		{
			name:  "Full rotation",
			turns: duit_utils.Float32Value(1.0),
		},
		{
			name:  "Multiple rotations",
			turns: duit_utils.Float32Value(3.14159),
		},
		{
			name:  "Small fractional turns",
			turns: duit_utils.Float32Value(0.25),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedRotationAttributes{
				Turns: tc.turns,
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

// JSON Serialization tests for Tristate fields

func TestAnimatedRotationAttributes_MarshalJSON_TurnsField(t *testing.T) {
	attrs := &duit_attributes.AnimatedRotationAttributes{
		Turns: duit_utils.Float32Value(3.14),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"turns":3.14`) {
		t.Fatalf("expected JSON to contain '\"turns\":3.14', got: %s", jsonStr)
	}
}

func TestAnimatedRotationAttributes_MarshalJSON_NilTurns(t *testing.T) {
	attrs := &duit_attributes.AnimatedRotationAttributes{
		// Turns is nil (default state)
		// All embedded fields are also nil, so the struct serializes as empty object
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with nil turns, got:", err)
	}

	jsonStr := string(jsonData)
	
	// When all fields (including embedded ones) are nil, the struct serializes as empty object
	expected := `{}`
	if jsonStr != expected {
		t.Fatalf("expected JSON to be '%s' for all nil fields, got: %s", expected, jsonStr)
	}
}



func TestAnimatedRotationAttributes_MarshalJSON_DifferentTurnsValues(t *testing.T) {
	testCases := []struct {
		name     string
		turns    duit_utils.Tristate[float32]
		expected string
	}{
		{
			name:     "Zero turns",
			turns:    duit_utils.Float32Value(0.0),
			expected: `"turns":0`,
		},
		{
			name:     "Positive turns",
			turns:    duit_utils.Float32Value(1.5),
			expected: `"turns":1.5`,
		},
		{
			name:     "Negative turns",
			turns:    duit_utils.Float32Value(-2.5),
			expected: `"turns":-2.5`,
		},
		{
			name:     "Full rotation",
			turns:    duit_utils.Float32Value(1.0),
			expected: `"turns":1`,
		},
		{
			name:     "Pi turns",
			turns:    duit_utils.Float32Value(3.14159),
			expected: `"turns":3.14159`,
		},
		{
			name:     "Quarter turn",
			turns:    duit_utils.Float32Value(0.25),
			expected: `"turns":0.25`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedRotationAttributes{
				Turns: tc.turns,
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

func TestAnimatedRotationAttributes_Validate_NilImplicitAnimatable(t *testing.T) {
	attrs := &duit_attributes.AnimatedRotationAttributes{
		Turns: duit_utils.Float32Value(3.14),
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