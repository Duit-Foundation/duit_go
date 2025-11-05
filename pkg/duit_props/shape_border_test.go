package duit_props_test

import (
	"encoding/json"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for NewRoundedRectangleBorder constructor
func TestNewRoundedRectangleBorder(t *testing.T) {
	borderRadius := duit_props.NewBorderRadiusAll(duit_props.NewRadiusCircular(10.0))
	borderSide := duit_props.NewBorderSide()
	border := duit_props.NewRoundedRectangleBorder(borderRadius, borderSide)

	if border == nil {
		t.Fatal("expected non-nil ShapeBorder")
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for valid rounded rectangle border, got:", err)
	}
}

func TestNewRoundedRectangleBorder_NilParams(t *testing.T) {
	border := duit_props.NewRoundedRectangleBorder(nil, nil)

	if border == nil {
		t.Fatal("expected non-nil ShapeBorder")
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for nil parameters, got:", err)
	}
}

// Tests for NewBeveledRectangleBorder constructor
func TestNewBeveledRectangleBorder(t *testing.T) {
	borderRadius := duit_props.NewBorderRadiusAll(duit_props.NewRadiusCircular(5.0))
	borderSide := duit_props.NewBorderSide().SetWidth(2.0)
	border := duit_props.NewBeveledRectangleBorder(borderRadius, borderSide)

	if border == nil {
		t.Fatal("expected non-nil ShapeBorder")
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for valid beveled rectangle border, got:", err)
	}
}

func TestNewBeveledRectangleBorder_NilParams(t *testing.T) {
	border := duit_props.NewBeveledRectangleBorder(nil, nil)

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for nil parameters, got:", err)
	}
}

// Tests for NewContinuousRectangleBorder constructor
func TestNewContinuousRectangleBorder(t *testing.T) {
	borderRadius := duit_props.NewBorderRadiusAll(duit_props.NewRadiusCircular(8.0))
	borderSide := duit_props.NewBorderSide().SetWidth(1.5)
	border := duit_props.NewContinuousRectangleBorder(borderRadius, borderSide)

	if border == nil {
		t.Fatal("expected non-nil ShapeBorder")
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for valid continuous rectangle border, got:", err)
	}
}

func TestNewContinuousRectangleBorder_NilParams(t *testing.T) {
	border := duit_props.NewContinuousRectangleBorder(nil, nil)

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for nil parameters, got:", err)
	}
}

// Tests for NewCircleBorder constructor
func TestNewCircleBorder(t *testing.T) {
	borderSide := duit_props.NewBorderSide().SetWidth(3.0)
	border := duit_props.NewCircleBorder(borderSide, 0.5)

	if border == nil {
		t.Fatal("expected non-nil ShapeBorder")
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for valid circle border, got:", err)
	}
}

func TestNewCircleBorder_NilBorderSide(t *testing.T) {
	border := duit_props.NewCircleBorder(nil, 1.0)

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for nil border side, got:", err)
	}
}

func TestNewCircleBorder_ZeroEccentricity(t *testing.T) {
	borderSide := duit_props.NewBorderSide()
	border := duit_props.NewCircleBorder(borderSide, 0.0)

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for zero eccentricity, got:", err)
	}
}

// Tests for NewStadiumBorder constructor
func TestNewStadiumBorder(t *testing.T) {
	borderSide := duit_props.NewBorderSide().SetWidth(2.5)
	border := duit_props.NewStadiumBorder(borderSide)

	if border == nil {
		t.Fatal("expected non-nil ShapeBorder")
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for valid stadium border, got:", err)
	}
}

func TestNewStadiumBorder_NilBorderSide(t *testing.T) {
	border := duit_props.NewStadiumBorder(nil)

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for nil border side, got:", err)
	}
}

// Tests for ShapeBorder.Validate() method with nil receiver
func TestShapeBorder_Validate_NilReceiver(t *testing.T) {
	var border *duit_props.ShapeBorder

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for nil ShapeBorder, got:", err)
	}
}

// Tests for ShapeBorder.MarshalJSON() method
func TestShapeBorder_MarshalJSON_RoundedRectangle(t *testing.T) {
	borderRadius := duit_props.NewBorderRadiusAll(duit_props.NewRadiusCircular(10.0))
	borderSide := duit_props.NewBorderSide().SetWidth(2.0)
	border := duit_props.NewRoundedRectangleBorder(borderRadius, borderSide)

	data, err := border.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling rounded rectangle border, got:", err)
	}

	// Verify that the JSON is valid
	var result interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("expected valid JSON, got unmarshal error: %v", err)
	}

	// Verify that JSON contains expected type
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("expected JSON object")
	}

	if resultMap["type"] != "RoundedRectangleBorder" {
		t.Fatalf("expected type 'RoundedRectangleBorder', got '%v'", resultMap["type"])
	}
}

func TestShapeBorder_MarshalJSON_Circle(t *testing.T) {
	borderSide := duit_props.NewBorderSide().SetWidth(1.0)
	border := duit_props.NewCircleBorder(borderSide, 0.8)

	data, err := border.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling circle border, got:", err)
	}

	// Verify that the JSON is valid
	var result interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("expected valid JSON, got unmarshal error: %v", err)
	}

	// Verify that JSON contains expected type
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("expected JSON object")
	}

	if resultMap["type"] != "CircleBorder" {
		t.Fatalf("expected type 'CircleBorder', got '%v'", resultMap["type"])
	}
}

func TestShapeBorder_MarshalJSON_Stadium(t *testing.T) {
	borderSide := duit_props.NewBorderSide()
	border := duit_props.NewStadiumBorder(borderSide)

	data, err := border.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling stadium border, got:", err)
	}

	// Verify that the JSON is valid
	var result interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("expected valid JSON, got unmarshal error: %v", err)
	}

	// Verify that JSON contains expected type
	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatal("expected JSON object")
	}

	if resultMap["type"] != "StadiumBorder" {
		t.Fatalf("expected type 'StadiumBorder', got '%v'", resultMap["type"])
	}
}

// Tests for all shape border types using table-driven tests
func TestShapeBorder_AllTypes(t *testing.T) {
	borderRadius := duit_props.NewBorderRadiusAll(duit_props.NewRadiusCircular(5.0))
	borderSide := duit_props.NewBorderSide().SetWidth(1.0)

	testCases := []struct {
		name   string
		border *duit_props.ShapeBorder
	}{
		{
			name:   "RoundedRectangleBorder",
			border: duit_props.NewRoundedRectangleBorder(borderRadius, borderSide),
		},
		{
			name:   "BeveledRectangleBorder",
			border: duit_props.NewBeveledRectangleBorder(borderRadius, borderSide),
		},
		{
			name:   "ContinuousRectangleBorder",
			border: duit_props.NewContinuousRectangleBorder(borderRadius, borderSide),
		},
		{
			name:   "CircleBorder",
			border: duit_props.NewCircleBorder(borderSide, 1.0),
		},
		{
			name:   "StadiumBorder",
			border: duit_props.NewStadiumBorder(borderSide),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.border.Validate()
			if err != nil {
				t.Fatalf("expected no error for valid %s, got: %v", tc.name, err)
			}

			data, err := tc.border.MarshalJSON()
			if err != nil {
				t.Fatalf("expected no error marshaling %s, got: %v", tc.name, err)
			}

			// Verify that the JSON is valid
			var result interface{}
			err = json.Unmarshal(data, &result)
			if err != nil {
				t.Fatalf("expected valid JSON for %s, got unmarshal error: %v", tc.name, err)
			}
		})
	}
}
