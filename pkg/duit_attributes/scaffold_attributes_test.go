package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_material"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestScaffoldAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.ScaffoldAttributes[duit_props.ColorString]{
		Primary:                  duit_utils.BoolValue(true),
		ExtendBody:               duit_utils.BoolValue(false),
		ResizeToAvoidBottomInset: duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestScaffoldAttributes_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.ScaffoldAttributes[duit_props.ColorString]{
		Primary:                      duit_utils.BoolValue(true),
		ExtendBody:                   duit_utils.BoolValue(true),
		ResizeToAvoidBottomInset:     duit_utils.BoolValue(true),
		BackgroundColor:              "#FFFFFF",
		RestorationId:                "test_restoration_id",
		PersistentFooterAlignment:    "center",
		FloatingActionButtonLocation: duit_material.EndFloat,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for Primary Tristate[bool] property serialization
func TestScaffoldAttributes_Primary_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ScaffoldAttributes[duit_props.ColorString]{
		Primary: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"primary":true`) {
		t.Fatalf("expected JSON to contain '\"primary\":true', got: %s", jsonStr)
	}
}

func TestScaffoldAttributes_Primary_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ScaffoldAttributes[duit_props.ColorString]{
		Primary: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"primary":false`) {
		t.Fatalf("expected JSON to contain '\"primary\":false', got: %s", jsonStr)
	}
}

func TestScaffoldAttributes_Primary_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ScaffoldAttributes[duit_props.ColorString]{
		Primary: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"primary"`) {
		t.Fatalf("expected JSON to not contain 'primary' field when nil, got: %s", jsonStr)
	}
}

// Tests for ExtendBody Tristate[bool] property serialization
func TestScaffoldAttributes_ExtendBody_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ScaffoldAttributes[duit_props.ColorString]{
		ExtendBody: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"extendBody":true`) {
		t.Fatalf("expected JSON to contain '\"extendBody\":true', got: %s", jsonStr)
	}
}

func TestScaffoldAttributes_ExtendBody_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ScaffoldAttributes[duit_props.ColorString]{
		ExtendBody: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"extendBody":false`) {
		t.Fatalf("expected JSON to contain '\"extendBody\":false', got: %s", jsonStr)
	}
}

// Tests for ResizeToAvoidBottomInset Tristate[bool] property serialization
func TestScaffoldAttributes_ResizeToAvoidBottomInset_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ScaffoldAttributes[duit_props.ColorString]{
		ResizeToAvoidBottomInset: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"resizeToAvoidBottomInset":true`) {
		t.Fatalf("expected JSON to contain '\"resizeToAvoidBottomInset\":true', got: %s", jsonStr)
	}
}

func TestScaffoldAttributes_ResizeToAvoidBottomInset_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ScaffoldAttributes[duit_props.ColorString]{
		ResizeToAvoidBottomInset: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"resizeToAvoidBottomInset":false`) {
		t.Fatalf("expected JSON to contain '\"resizeToAvoidBottomInset\":false', got: %s", jsonStr)
	}
}
