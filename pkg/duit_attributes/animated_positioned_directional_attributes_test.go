package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestAnimatedPositionedDirectionalAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedDirectionalAttributes{
		Start:   duit_utils.Float32Value(10.0),
		Top:     duit_utils.Float32Value(20.0),
		Bottom:  duit_utils.Float32Value(40.0),
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

func TestAnimatedPositionedDirectionalAttributes_Validate_DifferentPositionCombinations(t *testing.T) {
	testCases := []struct {
		name  string
		attrs *duit_attributes.AnimatedPositionedDirectionalAttributes
	}{
		{
			name: "Only Start and Top",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Start: duit_utils.Float32Value(10.0),
				Top:   duit_utils.Float32Value(20.0),
			},
		},
		{
			name: "Only End and Bottom",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				End:    duit_utils.Float32Value(30.0),
				Bottom: duit_utils.Float32Value(40.0),
			},
		},
		{
			name: "All positions set",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Start:  duit_utils.Float32Value(10.0),
				End:    duit_utils.Float32Value(30.0),
				Top:    duit_utils.Float32Value(20.0),
				Bottom: duit_utils.Float32Value(40.0),
			},
		},
		{
			name: "No positions set",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{},
		},
		{
			name: "Only Start set",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Start: duit_utils.Float32Value(15.0),
			},
		},
		{
			name: "Only End set",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				End: duit_utils.Float32Value(35.0),
			},
		},
		{
			name: "Only Top set",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Top: duit_utils.Float32Value(25.0),
			},
		},
		{
			name: "Only Bottom set",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Bottom: duit_utils.Float32Value(45.0),
			},
		},
		{
			name: "Width and Height set",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Width:  duit_utils.Float32Value(100.0),
				Height: duit_utils.Float32Value(200.0),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.attrs.ImplicitAnimatable = &duit_props.ImplicitAnimatable{
				Duration: 300,
				Curve:    duit_props.CurveEase,
			}
			err := tc.attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for %s, got: %v", tc.name, err)
			}
		})
	}
}

func TestAnimatedPositionedDirectionalAttributes_Validate_NilImplicitAnimatable(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedDirectionalAttributes{
		Start:             duit_utils.Float32Value(10.0),
		Top:               duit_utils.Float32Value(20.0),
		ImplicitAnimatable: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil ImplicitAnimatable")
	}
}

// JSON Serialization tests for Tristate fields

func TestAnimatedPositionedDirectionalAttributes_MarshalJSON_StartField(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedDirectionalAttributes{
		Start: duit_utils.Float32Value(10.5),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"start":10.5`) {
		t.Fatalf("expected JSON to contain '\"start\":10.5', got: %s", jsonStr)
	}
}

func TestAnimatedPositionedDirectionalAttributes_MarshalJSON_EndField(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedDirectionalAttributes{
		End: duit_utils.Float32Value(30.5),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"end":30.5`) {
		t.Fatalf("expected JSON to contain '\"end\":30.5', got: %s", jsonStr)
	}
}

func TestAnimatedPositionedDirectionalAttributes_MarshalJSON_TopField(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedDirectionalAttributes{
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

func TestAnimatedPositionedDirectionalAttributes_MarshalJSON_BottomField(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedDirectionalAttributes{
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

func TestAnimatedPositionedDirectionalAttributes_MarshalJSON_WidthField(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedDirectionalAttributes{
		Width: duit_utils.Float32Value(100.5),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"Width":100.5`) {
		t.Fatalf("expected JSON to contain '\"Width\":100.5', got: %s", jsonStr)
	}
}

func TestAnimatedPositionedDirectionalAttributes_MarshalJSON_HeightField(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedDirectionalAttributes{
		Height: duit_utils.Float32Value(200.5),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"Height":200.5`) {
		t.Fatalf("expected JSON to contain '\"Height\":200.5', got: %s", jsonStr)
	}
}

func TestAnimatedPositionedDirectionalAttributes_MarshalJSON_NilFields(t *testing.T) {
	attrs := &duit_attributes.AnimatedPositionedDirectionalAttributes{
		// All fields are nil (default state)
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with nil fields, got:", err)
	}

	jsonStr := string(jsonData)

	// Check that omitempty works - nil fields should not be present
	if strings.Contains(jsonStr, `"start"`) {
		t.Fatalf("expected JSON to NOT contain 'start' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"end"`) {
		t.Fatalf("expected JSON to NOT contain 'end' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"top"`) {
		t.Fatalf("expected JSON to NOT contain 'top' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"bottom"`) {
		t.Fatalf("expected JSON to NOT contain 'bottom' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"Width"`) {
		t.Fatalf("expected JSON to NOT contain 'Width' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"Height"`) {
		t.Fatalf("expected JSON to NOT contain 'Height' field for nil value, got: %s", jsonStr)
	}
}

func TestAnimatedPositionedDirectionalAttributes_MarshalJSON_AllFieldsCombinations(t *testing.T) {
	testCases := []struct {
		name        string
		attrs       *duit_attributes.AnimatedPositionedDirectionalAttributes
		expected    []string
		notExpected []string
	}{
		{
			name: "Start and Top only",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Start: duit_utils.Float32Value(10.0),
				Top:   duit_utils.Float32Value(20.0),
			},
			expected:    []string{`"start":10`, `"top":20`},
			notExpected: []string{`"end"`, `"bottom"`, `"Width"`, `"Height"`},
		},
		{
			name: "End and Bottom only",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				End:    duit_utils.Float32Value(30.0),
				Bottom: duit_utils.Float32Value(40.0),
			},
			expected:    []string{`"end":30`, `"bottom":40`},
			notExpected: []string{`"start"`, `"top"`, `"Width"`, `"Height"`},
		},
		{
			name: "All position fields set",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Start:  duit_utils.Float32Value(10.0),
				End:    duit_utils.Float32Value(30.0),
				Top:    duit_utils.Float32Value(20.0),
				Bottom: duit_utils.Float32Value(40.0),
			},
			expected:    []string{`"start":10`, `"end":30`, `"top":20`, `"bottom":40`},
			notExpected: []string{`"Width"`, `"Height"`},
		},
		{
			name: "Width and Height only",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Width:  duit_utils.Float32Value(100.0),
				Height: duit_utils.Float32Value(200.0),
			},
			expected:    []string{`"Width":100`, `"Height":200`},
			notExpected: []string{`"start"`, `"end"`, `"top"`, `"bottom"`},
		},
		{
			name: "All fields set",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Start:  duit_utils.Float32Value(10.0),
				End:    duit_utils.Float32Value(30.0),
				Top:    duit_utils.Float32Value(20.0),
				Bottom: duit_utils.Float32Value(40.0),
				Width:  duit_utils.Float32Value(100.0),
				Height: duit_utils.Float32Value(200.0),
			},
			expected:    []string{`"start":10`, `"end":30`, `"top":20`, `"bottom":40`, `"Width":100`, `"Height":200`},
			notExpected: []string{},
		},
		{
			name: "Zero values",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Start:  duit_utils.Float32Value(0.0),
				End:    duit_utils.Float32Value(0.0),
				Top:    duit_utils.Float32Value(0.0),
				Bottom: duit_utils.Float32Value(0.0),
				Width:  duit_utils.Float32Value(0.0),
				Height: duit_utils.Float32Value(0.0),
			},
			expected:    []string{`"start":0`, `"end":0`, `"top":0`, `"bottom":0`, `"Width":0`, `"Height":0`},
			notExpected: []string{},
		},
		{
			name: "Negative values",
			attrs: &duit_attributes.AnimatedPositionedDirectionalAttributes{
				Start:  duit_utils.Float32Value(-10.5),
				End:    duit_utils.Float32Value(-30.5),
				Top:    duit_utils.Float32Value(-20.5),
				Bottom: duit_utils.Float32Value(-40.5),
				Width:  duit_utils.Float32Value(-100.5),
				Height: duit_utils.Float32Value(-200.5),
			},
			expected:    []string{`"start":-10.5`, `"end":-30.5`, `"top":-20.5`, `"bottom":-40.5`, `"Width":-100.5`, `"Height":-200.5`},
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