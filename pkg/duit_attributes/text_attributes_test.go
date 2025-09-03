package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestTextAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.TextAttributes[duit_props.ColorString]{
		Data: "Hello World",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestTextAttributes_Validate_MissingData(t *testing.T) {
	attrs := &duit_attributes.TextAttributes[duit_props.ColorString]{
		Data: "",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing data")
	}

	if !strings.Contains(err.Error(), "data property is required") {
		t.Fatalf("expected error message about required data, got: %s", err.Error())
	}
}

func TestTextAttributes_Validate_WithAllProperties(t *testing.T) {
	style := &duit_props.TextStyle[duit_props.ColorString]{}
	
	attrs := &duit_attributes.TextAttributes[duit_props.ColorString]{
		Data:           "Test text with all properties",
		SemanticsLabel: "Test label",
		TextAlign:      duit_props.TextAlignCenter,
		TextDirection:  duit_props.TextDirectionLtr,
		Overflow:       duit_props.TextOverflowEllipsis,
		SoftWrap:       duit_utils.BoolValue(true),
		MaxLines:       5,
		Style:          style,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for SoftWrap Tristate[bool] property serialization
func TestTextAttributes_SoftWrap_JSON_True(t *testing.T) {
	attrs := &duit_attributes.TextAttributes[duit_props.ColorString]{
		Data:     "Test text",
		SoftWrap: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"softWrap":true`) {
		t.Fatalf("expected JSON to contain '\"softWrap\":true', got: %s", jsonStr)
	}
}

func TestTextAttributes_SoftWrap_JSON_False(t *testing.T) {
	attrs := &duit_attributes.TextAttributes[duit_props.ColorString]{
		Data:     "Test text",
		SoftWrap: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"softWrap":false`) {
		t.Fatalf("expected JSON to contain '\"softWrap\":false', got: %s", jsonStr)
	}
}

func TestTextAttributes_SoftWrap_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.TextAttributes[duit_props.ColorString]{
		Data:     "Test text",
		SoftWrap: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"softWrap"`) {
		t.Fatalf("expected JSON to not contain 'softWrap' field when nil, got: %s", jsonStr)
	}
}

// Tests for MaxLines property serialization
func TestTextAttributes_MaxLines_JSON_WithValue(t *testing.T) {
	attrs := &duit_attributes.TextAttributes[duit_props.ColorString]{
		Data:     "Test text",
		MaxLines: 10,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"maxLines":10`) {
		t.Fatalf("expected JSON to contain '\"maxLines\":10', got: %s", jsonStr)
	}
}

func TestTextAttributes_MaxLines_JSON_ZeroValue(t *testing.T) {
	attrs := &duit_attributes.TextAttributes[duit_props.ColorString]{
		Data:     "Test text",
		MaxLines: 0,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"maxLines"`) {
		t.Fatalf("expected JSON to not contain 'maxLines' field when zero value, got: %s", jsonStr)
	}
}
