package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestTextFieldAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.TextFieldAttributes[duit_props.EdgeInsetsLTRB, duit_props.ColorString]{
		Autocorrect:       duit_utils.BoolValue(true),
		EnableSuggestions: duit_utils.BoolValue(false),
		Expands:           duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestTextFieldAttributes_Validate_WithAllProperties(t *testing.T) {
	style := &duit_props.TextStyle[duit_props.ColorString]{}
	
	attrs := &duit_attributes.TextFieldAttributes[duit_props.EdgeInsetsLTRB, duit_props.ColorString]{
		Style:              style,
		TextAlign:          duit_props.TextAlignStart,
		TextDirection:      duit_props.TextDirectionLtr,
		Autocorrect:        duit_utils.BoolValue(true),
		EnableSuggestions:  duit_utils.BoolValue(true),
		Expands:            duit_utils.BoolValue(false),
		ReadOnly:           duit_utils.BoolValue(false),
		ShowCursor:         duit_utils.BoolValue(true),
		Enabled:            duit_utils.BoolValue(true),
		ObscureText:        duit_utils.BoolValue(false),
		Autofocus:          duit_utils.BoolValue(true),
		ObscuringCharacter: "*",
		MaxLines:           10,
		MinLines:           1,
		MaxLength:          100,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for Autocorrect Tristate[bool] property serialization
func TestTextFieldAttributes_Autocorrect_JSON_True(t *testing.T) {
	attrs := &duit_attributes.TextFieldAttributes[duit_props.EdgeInsetsLTRB, duit_props.ColorString]{
		Autocorrect: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"autocorrect":true`) {
		t.Fatalf("expected JSON to contain '\"autocorrect\":true', got: %s", jsonStr)
	}
}

func TestTextFieldAttributes_Autocorrect_JSON_False(t *testing.T) {
	attrs := &duit_attributes.TextFieldAttributes[duit_props.EdgeInsetsLTRB, duit_props.ColorString]{
		Autocorrect: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"autocorrect":false`) {
		t.Fatalf("expected JSON to contain '\"autocorrect\":false', got: %s", jsonStr)
	}
}

func TestTextFieldAttributes_Autocorrect_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.TextFieldAttributes[duit_props.EdgeInsetsLTRB, duit_props.ColorString]{
		Autocorrect: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"autocorrect"`) {
		t.Fatalf("expected JSON to not contain 'autocorrect' field when nil, got: %s", jsonStr)
	}
}

// Tests for EnableSuggestions Tristate[bool] property serialization
func TestTextFieldAttributes_EnableSuggestions_JSON_True(t *testing.T) {
	attrs := &duit_attributes.TextFieldAttributes[duit_props.EdgeInsetsLTRB, duit_props.ColorString]{
		EnableSuggestions: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"enableSuggestions":true`) {
		t.Fatalf("expected JSON to contain '\"enableSuggestions\":true', got: %s", jsonStr)
	}
}

func TestTextFieldAttributes_EnableSuggestions_JSON_False(t *testing.T) {
	attrs := &duit_attributes.TextFieldAttributes[duit_props.EdgeInsetsLTRB, duit_props.ColorString]{
		EnableSuggestions: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"enableSuggestions":false`) {
		t.Fatalf("expected JSON to contain '\"enableSuggestions\":false', got: %s", jsonStr)
	}
}

// Tests for Autofocus Tristate[bool] property serialization
func TestTextFieldAttributes_Autofocus_JSON_True(t *testing.T) {
	attrs := &duit_attributes.TextFieldAttributes[duit_props.EdgeInsetsLTRB, duit_props.ColorString]{
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

func TestTextFieldAttributes_Autofocus_JSON_False(t *testing.T) {
	attrs := &duit_attributes.TextFieldAttributes[duit_props.EdgeInsetsLTRB, duit_props.ColorString]{
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
