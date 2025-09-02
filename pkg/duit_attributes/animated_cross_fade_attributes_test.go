package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestAnimatedCrossFadeAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AnimatedCrossFadeAttributes{
		CrossFadeState: 0, // showFirst
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

func TestAnimatedCrossFadeAttributes_Validate_ValidCrossFadeStates(t *testing.T) {
	testCases := []struct {
		name  string
		state uint8
	}{
		{
			name:  "ShowFirst state (0)",
			state: 0,
		},
		{
			name:  "ShowSecond state (1)",
			state: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedCrossFadeAttributes{
				CrossFadeState: tc.state,
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

func TestAnimatedCrossFadeAttributes_Validate_InvalidCrossFadeState(t *testing.T) {
	invalidStates := []uint8{2, 3, 255}

	for _, state := range invalidStates {
		t.Run("Invalid state", func(t *testing.T) {
			attrs := &duit_attributes.AnimatedCrossFadeAttributes{
				CrossFadeState: state,
				ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
					Duration: 300,
					Curve:    duit_animations.Ease,
				},
			}

			err := attrs.Validate()
			if err == nil {
				t.Fatalf("expected error for invalid CrossFadeState %d", state)
			}

			expected := "crossFadeState must be 0 or 1"
			if err.Error() != expected {
				t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
			}
		})
	}
}

func TestAnimatedCrossFadeAttributes_Validate_WithOptionalFields(t *testing.T) {
	attrs := &duit_attributes.AnimatedCrossFadeAttributes{
		CrossFadeState:     0,
		ReverseDuration:    200,
		FirstCurve:         duit_animations.EaseIn,
		SecondCurve:        duit_animations.EaseOut,
		SizeCurve:          duit_animations.Linear,
		ExcludeBottomFocus: duit_utils.BoolValue(true),
		Alignment:          duit_props.AlignmentCenter,
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

func TestAnimatedCrossFadeAttributes_MarshalJSON_ExcludeBottomFocus(t *testing.T) {
	attrs := &duit_attributes.AnimatedCrossFadeAttributes{
		CrossFadeState:     0,
		ExcludeBottomFocus: duit_utils.BoolValue(true),
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
	if !strings.Contains(jsonStr, `"excludeBottomFocus":true`) {
		t.Fatalf("expected JSON to contain '\"excludeBottomFocus\":true', got: %s", jsonStr)
	}
}

func TestAnimatedCrossFadeAttributes_MarshalJSON_NilExcludeBottomFocus(t *testing.T) {
	attrs := &duit_attributes.AnimatedCrossFadeAttributes{
		CrossFadeState:     0,
		ExcludeBottomFocus: nil,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with nil ExcludeBottomFocus, got:", err)
	}

	jsonStr := string(jsonData)
	// Check that omitempty works - nil fields should not be present
	if strings.Contains(jsonStr, `"excludeBottomFocus"`) {
		t.Fatalf("expected JSON to NOT contain 'excludeBottomFocus' field for nil value, got: %s", jsonStr)
	}
}

// Tests for ImplicitAnimatable embedded struct

func TestAnimatedCrossFadeAttributes_Validate_NilImplicitAnimatable(t *testing.T) {
	attrs := &duit_attributes.AnimatedCrossFadeAttributes{
		CrossFadeState:     0,
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

func TestAnimatedCrossFadeAttributes_MarshalJSON_WithAllFields(t *testing.T) {
	attrs := &duit_attributes.AnimatedCrossFadeAttributes{
		CrossFadeState:     1,
		ReverseDuration:    200,
		FirstCurve:         duit_animations.EaseIn,
		SecondCurve:        duit_animations.EaseOut,
		SizeCurve:          duit_animations.Linear,
		ExcludeBottomFocus: duit_utils.BoolValue(false),
		Alignment:          duit_props.AlignmentTopCenter,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 400,
			Curve:    duit_animations.BounceIn,
		},
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with all fields, got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"reverseDuration":200`) {
		t.Fatalf("expected JSON to contain '\"reverseDuration\":200', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"firstCurve":"easeIn"`) {
		t.Fatalf("expected JSON to contain '\"firstCurve\":\"easeIn\"', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"secondCurve":"easeOut"`) {
		t.Fatalf("expected JSON to contain '\"secondCurve\":\"easeOut\"', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"sizeCurve":"linear"`) {
		t.Fatalf("expected JSON to contain '\"sizeCurve\":\"linear\"', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"excludeBottomFocus":false`) {
		t.Fatalf("expected JSON to contain '\"excludeBottomFocus\":false', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"alignment":"topCenter"`) {
		t.Fatalf("expected JSON to contain '\"alignment\":\"topCenter\"', got: %s", jsonStr)
	}
}
