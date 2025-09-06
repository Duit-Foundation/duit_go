package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestAbsorbPointerAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AbsorbPointerAttributes{
		Absorbing: duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestAbsorbPointerAttributes_Validate_WithoutAbsorbing(t *testing.T) {
	attrs := &duit_attributes.AbsorbPointerAttributes{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error when absorbing is not set, got:", err)
	}
}

// Tests for Tristate[bool] property serialization
func TestAbsorbPointerAttributes_Tristate_JSON_True(t *testing.T) {
	attrs := &duit_attributes.AbsorbPointerAttributes{
		Absorbing: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"absorbing":true`) {
		t.Fatalf("expected JSON to contain '\"absorbing\":true', got: %s", jsonStr)
	}
}

func TestAbsorbPointerAttributes_Tristate_JSON_False(t *testing.T) {
	attrs := &duit_attributes.AbsorbPointerAttributes{
		Absorbing: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"absorbing":false`) {
		t.Fatalf("expected JSON to contain '\"absorbing\":false', got: %s", jsonStr)
	}
}

func TestAbsorbPointerAttributes_Tristate_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.AbsorbPointerAttributes{
		Absorbing: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	// When Tristate is nil/unset, it should be omitted from JSON due to omitempty tag
	if strings.Contains(jsonStr, `"absorbing"`) {
		t.Fatalf("expected JSON to not contain 'absorbing' field when nil, got: %s", jsonStr)
	}
}

func TestAbsorbPointerAttributes_Tristate_Unmarshal_True(t *testing.T) {
	jsonStr := `{"absorbing": true}`
	
	var attrs duit_attributes.AbsorbPointerAttributes
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.Absorbing == nil {
		t.Fatal("expected absorbing to have a value")
	}

	if !*attrs.Absorbing {
		t.Fatal("expected absorbing value to be true")
	}
}

func TestAbsorbPointerAttributes_Tristate_Unmarshal_False(t *testing.T) {
	jsonStr := `{"absorbing": false}`
	
	var attrs duit_attributes.AbsorbPointerAttributes
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.Absorbing == nil {
		t.Fatal("expected absorbing to have a value")
	}

	if *attrs.Absorbing {
		t.Fatal("expected absorbing value to be false")
	}
}

func TestAbsorbPointerAttributes_Tristate_Unmarshal_Omitted(t *testing.T) {
	jsonStr := `{}`
	
	var attrs duit_attributes.AbsorbPointerAttributes
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.Absorbing != nil {
		t.Fatal("expected absorbing to not have a value when omitted")
	}
}
