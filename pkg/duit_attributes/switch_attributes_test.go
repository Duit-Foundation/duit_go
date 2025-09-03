package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestSwitchAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SwitchAttributes[duit_props.ColorString]{
		Value: duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSwitchAttributes_Validate_MissingValue(t *testing.T) {
	attrs := &duit_attributes.SwitchAttributes[duit_props.ColorString]{
		Value: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing value")
	}

	if !strings.Contains(err.Error(), "value property is required") {
		t.Fatalf("expected error message about required value, got: %s", err.Error())
	}
}

func TestSwitchAttributes_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.SwitchAttributes[duit_props.ColorString]{
		Value:                 duit_utils.BoolValue(false),
		ActiveColor:           "#FF0000",
		FocusColor:            "#00FF00",
		HoverColor:            "#0000FF",
		ActiveTrackColor:      "#FFFF00",
		InactiveTrackColor:    "#FF00FF",
		SplashRadius:          20.0,
		MaterialTapTargetSize: "padded",
		Autofocus:             duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for Value Tristate[bool] property serialization
func TestSwitchAttributes_Value_JSON_True(t *testing.T) {
	attrs := &duit_attributes.SwitchAttributes[duit_props.ColorString]{
		Value: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"value":true`) {
		t.Fatalf("expected JSON to contain '\"value\":true', got: %s", jsonStr)
	}
}

func TestSwitchAttributes_Value_JSON_False(t *testing.T) {
	attrs := &duit_attributes.SwitchAttributes[duit_props.ColorString]{
		Value: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"value":false`) {
		t.Fatalf("expected JSON to contain '\"value\":false', got: %s", jsonStr)
	}
}

func TestSwitchAttributes_Value_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.SwitchAttributes[duit_props.ColorString]{
		Value: duit_utils.Nillable[bool](),
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
func TestSwitchAttributes_Autofocus_JSON_True(t *testing.T) {
	attrs := &duit_attributes.SwitchAttributes[duit_props.ColorString]{
		Value:     duit_utils.BoolValue(true),
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

func TestSwitchAttributes_Autofocus_JSON_False(t *testing.T) {
	attrs := &duit_attributes.SwitchAttributes[duit_props.ColorString]{
		Value:     duit_utils.BoolValue(true),
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
