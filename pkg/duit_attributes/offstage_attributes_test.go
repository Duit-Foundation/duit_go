package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestOffstageAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.OffstageAttributes{
		Offstage: duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestOffstageAttributes_Validate_WithoutOffstage(t *testing.T) {
	attrs := &duit_attributes.OffstageAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when offstage is not set")
	}
}

// Tests for Tristate[bool] property serialization
func TestOffstageAttributes_Tristate_JSON_True(t *testing.T) {
	attrs := &duit_attributes.OffstageAttributes{
		Offstage: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"offstage":true`) {
		t.Fatalf("expected JSON to contain '\"offstage\":true', got: %s", jsonStr)
	}
}

func TestOffstageAttributes_Tristate_JSON_False(t *testing.T) {
	attrs := &duit_attributes.OffstageAttributes{
		Offstage: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"offstage":false`) {
		t.Fatalf("expected JSON to contain '\"offstage\":false', got: %s", jsonStr)
	}
}

func TestOffstageAttributes_Tristate_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.OffstageAttributes{
		Offstage: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	// When Tristate is nil/unset, it should be omitted from JSON due to omitempty tag
	if strings.Contains(jsonStr, `"offstage"`) {
		t.Fatalf("expected JSON to not contain 'offstage' field when nil, got: %s", jsonStr)
	}
}