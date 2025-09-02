package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestSliderAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliderAttributes[duit_action.RemoteAction, duit_props.ColorString]{
		Value: duit_utils.Float32Value(50.0),
		Min:   0.0,
		Max:   100.0,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliderAttributes_Validate_MissingValue(t *testing.T) {
	attrs := &duit_attributes.SliderAttributes[duit_action.RemoteAction, duit_props.ColorString]{
		Value: nil,
		Min:   0.0,
		Max:   100.0,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing value")
	}

	if !strings.Contains(err.Error(), "value property is required") {
		t.Fatalf("expected error message about required value, got: %s", err.Error())
	}
}

func TestSliderAttributes_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.SliderAttributes[duit_action.RemoteAction, duit_props.ColorString]{
		Value:                duit_utils.Float32Value(75.0),
		Min:                  0.0,
		Max:                  100.0,
		Divisions:            10,
		SecondaryTrackValue:  25.0,
		ActiveColor:          "#FF0000",
		InactiveColor:        "#CCCCCC",
		ThumbColor:           "#0000FF",
		SecondaryActiveColor: "#00FF00",
		Autofocus:            duit_utils.BoolValue(true),
		Label:                "Test Slider",
		AllowedInteraction:   "tap",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for Value Tristate[float32] property serialization
func TestSliderAttributes_Value_JSON_WithValue(t *testing.T) {
	attrs := &duit_attributes.SliderAttributes[duit_action.RemoteAction, duit_props.ColorString]{
		Value: duit_utils.Float32Value(60.5),
		Min:   0.0,
		Max:   100.0,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"value":60.5`) {
		t.Fatalf("expected JSON to contain '\"value\":60.5', got: %s", jsonStr)
	}
}

func TestSliderAttributes_Value_JSON_ZeroValue(t *testing.T) {
	attrs := &duit_attributes.SliderAttributes[duit_action.RemoteAction, duit_props.ColorString]{
		Value: duit_utils.Float32Value(0.0),
		Min:   0.0,
		Max:   100.0,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"value":0`) {
		t.Fatalf("expected JSON to contain '\"value\":0', got: %s", jsonStr)
	}
}

func TestSliderAttributes_Value_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.SliderAttributes[duit_action.RemoteAction, duit_props.ColorString]{
		Value: duit_utils.Nillable[float32](),
		Min:   0.0,
		Max:   100.0,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"value"`) {
		t.Fatalf("expected JSON to not contain 'value' field when nil, got: %s", jsonStr)
	}
}

// Tests for Autofocus Tristate[bool] property serialization
func TestSliderAttributes_Autofocus_JSON_True(t *testing.T) {
	attrs := &duit_attributes.SliderAttributes[duit_action.RemoteAction, duit_props.ColorString]{
		Value:     duit_utils.Float32Value(50.0),
		Min:       0.0,
		Max:       100.0,
		Autofocus: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"autofocus":true`) {
		t.Fatalf("expected JSON to contain '\"autofocus\":true', got: %s", jsonStr)
	}
}

func TestSliderAttributes_Autofocus_JSON_False(t *testing.T) {
	attrs := &duit_attributes.SliderAttributes[duit_action.RemoteAction, duit_props.ColorString]{
		Value:     duit_utils.Float32Value(50.0),
		Min:       0.0,
		Max:       100.0,
		Autofocus: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"autofocus":false`) {
		t.Fatalf("expected JSON to contain '\"autofocus\":false', got: %s", jsonStr)
	}
}
