package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestCheckboxAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
		Value: duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestCheckboxAttributes_Validate_MissingValue(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
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

func TestCheckboxAttributes_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
		Value:         duit_utils.BoolValue(true),
		Autofocus:     duit_utils.BoolValue(false),
		Tristate:      duit_utils.BoolValue(true),
		IsError:       duit_utils.BoolValue(false),
		SplashRadius:  10.5,
		SemanticLabel: "Test checkbox",
		CheckColor:    "#00FF00",
		HoverColor:    "#0000FF",
		ActiveColor:   "#FF0000",
		FocusColor:    "#FFFF00",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for Value Tristate[bool] property serialization
func TestCheckboxAttributes_Value_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
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

func TestCheckboxAttributes_Value_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
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

// Tests for Autofocus Tristate[bool] property serialization
func TestCheckboxAttributes_Autofocus_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
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

func TestCheckboxAttributes_Autofocus_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
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

func TestCheckboxAttributes_Autofocus_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
		Value:     duit_utils.BoolValue(true),
		Autofocus: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"autofocus"`) {
		t.Fatalf("expected JSON to not contain 'autofocus' field when nil, got: %s", jsonStr)
	}
}

// Tests for Tristate Tristate[bool] property serialization
func TestCheckboxAttributes_Tristate_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
		Value:    duit_utils.BoolValue(true),
		Tristate: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"tristate":true`) {
		t.Fatalf("expected JSON to contain '\"tristate\":true', got: %s", jsonStr)
	}
}

func TestCheckboxAttributes_Tristate_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
		Value:    duit_utils.BoolValue(true),
		Tristate: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"tristate":false`) {
		t.Fatalf("expected JSON to contain '\"tristate\":false', got: %s", jsonStr)
	}
}

func TestCheckboxAttributes_Tristate_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
		Value:    duit_utils.BoolValue(true),
		Tristate: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"tristate"`) {
		t.Fatalf("expected JSON to not contain 'tristate' field when nil, got: %s", jsonStr)
	}
}

// Tests for IsError Tristate[bool] property serialization
func TestCheckboxAttributes_IsError_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
		Value:   duit_utils.BoolValue(true),
		IsError: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"isError":true`) {
		t.Fatalf("expected JSON to contain '\"isError\":true', got: %s", jsonStr)
	}
}

func TestCheckboxAttributes_IsError_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
		Value:   duit_utils.BoolValue(true),
		IsError: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"isError":false`) {
		t.Fatalf("expected JSON to contain '\"isError\":false', got: %s", jsonStr)
	}
}

func TestCheckboxAttributes_IsError_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.CheckboxAttributes[duit_props.ColorString]{
		Value:   duit_utils.BoolValue(true),
		IsError: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"isError"`) {
		t.Fatalf("expected JSON to not contain 'isError' field when nil, got: %s", jsonStr)
	}
}
