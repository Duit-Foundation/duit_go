package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestAnimatedPositionedAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedAttributes{
		Left:   duit_utils.Float32Value(10.0),
		Top:    duit_utils.Float32Value(20.0),
		Right:  duit_utils.Float32Value(30.0),
		Bottom: duit_utils.Float32Value(40.0),
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

func TestAnimatedPositionedAttributes_Validate_DifferentPositionCombinations(t *testing.T) {
	testCases := []struct {
		name   string
		attrs  *duit_attributes.AnimatedPositionedAttributes
	}{
		{
			name: "Only Left and Top",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Left: duit_utils.Float32Value(10.0),
				Top:  duit_utils.Float32Value(20.0),
			},
		},
		{
			name: "Only Right and Bottom",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Right:  duit_utils.Float32Value(30.0),
				Bottom: duit_utils.Float32Value(40.0),
			},
		},
		{
			name: "All positions set",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Left:   duit_utils.Float32Value(10.0),
				Top:    duit_utils.Float32Value(20.0),
				Right:  duit_utils.Float32Value(30.0),
				Bottom: duit_utils.Float32Value(40.0),
			},
		},
		{
			name: "No positions set",
			attrs: &duit_attributes.AnimatedPositionedAttributes{},
		},
		{
			name: "Only Left set",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Left: duit_utils.Float32Value(15.0),
			},
		},
		{
			name: "Only Top set",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Top: duit_utils.Float32Value(25.0),
			},
		},
		{
			name: "Only Right set",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Right: duit_utils.Float32Value(35.0),
			},
		},
		{
			name: "Only Bottom set",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Bottom: duit_utils.Float32Value(45.0),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.attrs.ImplicitAnimatable = &duit_animations.ImplicitAnimatable{
				Duration: 300,
				Curve:    duit_animations.Ease,
			}
			err := tc.attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for %s, got: %v", tc.name, err)
			}
		})
	}
}

// JSON Serialization tests for Tristate fields

func TestAnimatedPositionedAttributes_MarshalJSON_LeftField(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedAttributes{
		Left: duit_utils.Float32Value(10.5),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"left":10.5`) {
		t.Fatalf("expected JSON to contain '\"left\":10.5', got: %s", jsonStr)
	}
}

func TestAnimatedPositionedAttributes_MarshalJSON_TopField(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedAttributes{
		Top: duit_utils.Float32Value(20.5),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"top":20.5`) {
		t.Fatalf("expected JSON to contain '\"top\":20.5', got: %s", jsonStr)
	}
}

func TestAnimatedPositionedAttributes_MarshalJSON_RightField(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedAttributes{
		Right: duit_utils.Float32Value(30.5),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"right":30.5`) {
		t.Fatalf("expected JSON to contain '\"right\":30.5', got: %s", jsonStr)
	}
}

func TestAnimatedPositionedAttributes_MarshalJSON_BottomField(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedAttributes{
		Bottom: duit_utils.Float32Value(40.5),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"bottom":40.5`) {
		t.Fatalf("expected JSON to contain '\"bottom\":40.5', got: %s", jsonStr)
	}
}

func TestAnimatedPositionedAttributes_MarshalJSON_NilFields(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedAttributes{
		// All fields are nil (default state)
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with nil fields, got:", err)
	}

	jsonStr := string(jsonData)
	
	// Check that omitempty works - nil fields should not be present
	if strings.Contains(jsonStr, `"left"`) {
		t.Fatalf("expected JSON to NOT contain 'left' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"top"`) {
		t.Fatalf("expected JSON to NOT contain 'top' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"right"`) {
		t.Fatalf("expected JSON to NOT contain 'right' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"bottom"`) {
		t.Fatalf("expected JSON to NOT contain 'bottom' field for nil value, got: %s", jsonStr)
	}
}

func TestAnimatedPositionedAttributes_MarshalJSON_AllFieldsCombinations(t *testing.T) {
	testCases := []struct {
		name     string
		attrs    *duit_attributes.AnimatedPositionedAttributes
		expected []string
		notExpected []string
	}{
		{
			name: "Left and Top only",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Left: duit_utils.Float32Value(10.0),
				Top:  duit_utils.Float32Value(20.0),
			},
			expected:    []string{`"left":10`, `"top":20`},
			notExpected: []string{`"right"`, `"bottom"`},
		},
		{
			name: "Right and Bottom only",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Right:  duit_utils.Float32Value(30.0),
				Bottom: duit_utils.Float32Value(40.0),
			},
			expected:    []string{`"right":30`, `"bottom":40`},
			notExpected: []string{`"left"`, `"top"`},
		},
		{
			name: "All fields set",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Left:   duit_utils.Float32Value(10.0),
				Top:    duit_utils.Float32Value(20.0),
				Right:  duit_utils.Float32Value(30.0),
				Bottom: duit_utils.Float32Value(40.0),
			},
			expected:    []string{`"left":10`, `"top":20`, `"right":30`, `"bottom":40`},
			notExpected: []string{},
		},
		{
			name: "Zero values",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Left:   duit_utils.Float32Value(0.0),
				Top:    duit_utils.Float32Value(0.0),
				Right:  duit_utils.Float32Value(0.0),
				Bottom: duit_utils.Float32Value(0.0),
			},
			expected:    []string{`"left":0`, `"top":0`, `"right":0`, `"bottom":0`},
			notExpected: []string{},
		},
		{
			name: "Negative values",
			attrs: &duit_attributes.AnimatedPositionedAttributes{
				Left:   duit_utils.Float32Value(-10.5),
				Top:    duit_utils.Float32Value(-20.5),
				Right:  duit_utils.Float32Value(-30.5),
				Bottom: duit_utils.Float32Value(-40.5),
			},
			expected:    []string{`"left":-10.5`, `"top":-20.5`, `"right":-30.5`, `"bottom":-40.5`},
			notExpected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			jsonData, err := json.Marshal(tc.attrs)
			if err != nil {
				t.Fatalf("expected no error for MarshalJSON() in %s, got: %v", tc.name, err)
			}

			jsonStr := string(jsonData)

			// Check expected strings are present
			for _, expected := range tc.expected {
				if !strings.Contains(jsonStr, expected) {
					t.Fatalf("expected JSON to contain '%s' in %s, got: %s", expected, tc.name, jsonStr)
				}
			}

			// Check not expected strings are absent
			for _, notExpected := range tc.notExpected {
				if strings.Contains(jsonStr, notExpected) {
					t.Fatalf("expected JSON to NOT contain '%s' in %s, got: %s", notExpected, tc.name, jsonStr)
				}
			}
		})
	}
}

// Tests for ImplicitAnimatable embedded struct

func TestAnimatedPositionedAttributes_Validate_NilImplicitAnimatable(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedAttributes{
		Left: duit_utils.Float32Value(10.0),
		Top:  duit_utils.Float32Value(20.0),
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