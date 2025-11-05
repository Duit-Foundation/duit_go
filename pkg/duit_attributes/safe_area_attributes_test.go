package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestSafeAreaAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SafeAreaAttributes{
		Left:  duit_utils.BoolValue(true),
		Top:   duit_utils.BoolValue(false),
		Right: duit_utils.BoolValue(true),
		Bottom: duit_utils.BoolValue(false),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSafeAreaAttributes_Validate_WithAllProperties(t *testing.T) {
	minimum := duit_props.NewEdgeInsetsLTRB(10.0, 20.0, 30.0, 40.0)
	
	attrs := &duit_attributes.SafeAreaAttributes{
		Left:                      duit_utils.BoolValue(true),
		Top:                       duit_utils.BoolValue(true),
		Right:                     duit_utils.BoolValue(true),
		Bottom:                    duit_utils.BoolValue(true),
		Minimum:                   minimum,
		MaintainBottomViewPadding: duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for Left Tristate[bool] property serialization
func TestSafeAreaAttributes_Left_JSON_True(t *testing.T) {
	attrs := &duit_attributes.SafeAreaAttributes{
		Left: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"left":true`) {
		t.Fatalf("expected JSON to contain '\"left\":true', got: %s", jsonStr)
	}
}

func TestSafeAreaAttributes_Left_JSON_False(t *testing.T) {
	attrs := &duit_attributes.SafeAreaAttributes{
		Left: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"left":false`) {
		t.Fatalf("expected JSON to contain '\"left\":false', got: %s", jsonStr)
	}
}

func TestSafeAreaAttributes_Left_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.SafeAreaAttributes{
		Left: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"left"`) {
		t.Fatalf("expected JSON to not contain 'left' field when nil, got: %s", jsonStr)
	}
}

// Tests for Top Tristate[bool] property serialization
func TestSafeAreaAttributes_Top_JSON_True(t *testing.T) {
	attrs := &duit_attributes.SafeAreaAttributes{
		Top: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"top":true`) {
		t.Fatalf("expected JSON to contain '\"top\":true', got: %s", jsonStr)
	}
}

func TestSafeAreaAttributes_Top_JSON_False(t *testing.T) {
	attrs := &duit_attributes.SafeAreaAttributes{
		Top: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"top":false`) {
		t.Fatalf("expected JSON to contain '\"top\":false', got: %s", jsonStr)
	}
}

// Tests for MaintainBottomViewPadding Tristate[bool] property serialization
func TestSafeAreaAttributes_MaintainBottomViewPadding_JSON_True(t *testing.T) {
	attrs := &duit_attributes.SafeAreaAttributes{
		MaintainBottomViewPadding: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"maintainBottomViewPadding":true`) {
		t.Fatalf("expected JSON to contain '\"maintainBottomViewPadding\":true', got: %s", jsonStr)
	}
}

func TestSafeAreaAttributes_MaintainBottomViewPadding_JSON_False(t *testing.T) {
	attrs := &duit_attributes.SafeAreaAttributes{
		MaintainBottomViewPadding: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"maintainBottomViewPadding":false`) {
		t.Fatalf("expected JSON to contain '\"maintainBottomViewPadding\":false', got: %s", jsonStr)
	}
}
