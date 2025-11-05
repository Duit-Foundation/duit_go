package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// Tests for RadioAttributes
func TestRadioAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.RadioAttributes{
		Value: duit_utils.TristateFrom[any]("option1"),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestRadioAttributes_Validate_WithoutValue(t *testing.T) {
	attrs := &duit_attributes.RadioAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when value is not set")
	}
}

// Tests for Tristate[string] property serialization in RadioAttributes
func TestRadioAttributes_Tristate_JSON_ValidValue(t *testing.T) {
	attrs := &duit_attributes.RadioAttributes{
		Value: duit_utils.TristateFrom[any]("test_option"),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"value":"test_option"`) {
		t.Fatalf("expected JSON to contain '\"value\":\"test_option\"', got: %s", jsonStr)
	}
}

func TestRadioAttributes_Tristate_JSON_EmptyString(t *testing.T) {
	attrs := &duit_attributes.RadioAttributes{
		Value: duit_utils.TristateFrom[any](""),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"value":""`) {
		t.Fatalf("expected JSON to contain '\"value\":\"\"', got: %s", jsonStr)
	}
}

func TestRadioAttributes_Tristate_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.RadioAttributes{
		Value: duit_utils.TristateFrom[any](nil),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	// When Tristate is nil/unset, it should be omitted from JSON due to omitempty tag
	if strings.Contains(jsonStr, `"value"`) {
		t.Fatalf("expected JSON to not contain 'value' field when nil, got: %s", jsonStr)
	}
}

// Tests for RadioGroupContextAttributes
func TestRadioGroupContextAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.RadioGroupContextAttributes{
		GroupValue: duit_utils.TristateFrom[any]("selected_option"),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestRadioGroupContextAttributes_Validate_WithoutGroupValue(t *testing.T) {
	attrs := &duit_attributes.RadioGroupContextAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when groupValue is not set")
	}
}

// Tests for Tristate[string] property serialization in RadioGroupContextAttributes
func TestRadioGroupContextAttributes_Tristate_JSON_ValidValue(t *testing.T) {
	attrs := &duit_attributes.RadioGroupContextAttributes{
		GroupValue: duit_utils.TristateFrom[any]("group_selection"),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"groupValue":"group_selection"`) {
		t.Fatalf("expected JSON to contain '\"groupValue\":\"group_selection\"', got: %s", jsonStr)
	}
}

func TestRadioGroupContextAttributes_Tristate_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.RadioGroupContextAttributes{
		GroupValue: duit_utils.TristateFrom[any](nil),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	// When Tristate is nil/unset, it should be omitted from JSON due to omitempty tag
	if strings.Contains(jsonStr, `"groupValue"`) {
		t.Fatalf("expected JSON to not contain 'groupValue' field when nil, got: %s", jsonStr)
	}
}
