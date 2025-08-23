package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestIgnorePointerAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.IgnorePointerAttributes{
		Ignoring: duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestIgnorePointerAttributes_Validate_WithoutIgnoring(t *testing.T) {
	attrs := &duit_attributes.IgnorePointerAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when ignoring property is not set")
	}
}

func TestIgnorePointerAttributes_Validate_IgnoringTrue(t *testing.T) {
	attrs := &duit_attributes.IgnorePointerAttributes{
		Ignoring: duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error when ignoring is set to true, got:", err)
	}
}

func TestIgnorePointerAttributes_Validate_IgnoringFalse(t *testing.T) {
	attrs := &duit_attributes.IgnorePointerAttributes{
		Ignoring: duit_utils.BoolValue(false),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error when ignoring is set to false, got:", err)
	}
}

// Tests for Tristate[bool] property serialization
func TestIgnorePointerAttributes_Tristate_JSON_True(t *testing.T) {
	attrs := &duit_attributes.IgnorePointerAttributes{
		Ignoring: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"ignoring":true`) {
		t.Fatalf("expected JSON to contain '\"ignoring\":true', got: %s", jsonStr)
	}
}

func TestIgnorePointerAttributes_Tristate_JSON_False(t *testing.T) {
	attrs := &duit_attributes.IgnorePointerAttributes{
		Ignoring: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"ignoring":false`) {
		t.Fatalf("expected JSON to contain '\"ignoring\":false', got: %s", jsonStr)
	}
}

func TestIgnorePointerAttributes_Tristate_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.IgnorePointerAttributes{
		Ignoring: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	// When Tristate is nil/unset, it should be omitted from JSON due to omitempty tag
	if strings.Contains(jsonStr, `"ignoring"`) {
		t.Fatalf("expected JSON to not contain 'ignoring' field when nil, got: %s", jsonStr)
	}
}