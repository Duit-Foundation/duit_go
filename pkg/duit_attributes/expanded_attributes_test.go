package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
)

func TestExpandedAttributes_Validate_InvalidFlex_Zero(t *testing.T) {
	attrs := &duit_attributes.ExpandedAttributes{
		Flex: 0,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for flex value 0, got nil")
	}

	expectedError := "flex must be greater than 0"
	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf("expected error to contain '%s', got: %s", expectedError, err.Error())
	}
}

func TestExpandedAttributes_Validate_ValidFlex_One(t *testing.T) {
	attrs := &duit_attributes.ExpandedAttributes{
		Flex: 1,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for flex value 1, got:", err)
	}
}

// Edge case tests

func TestExpandedAttributes_JSON_EmptyStruct(t *testing.T) {
	attrs := &duit_attributes.ExpandedAttributes{}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with empty struct, got:", err)
	}

	jsonStr := string(jsonData)
	// Empty struct should result in empty JSON object or minimal content due to omitempty
	if strings.Contains(jsonStr, `"flex"`) {
		t.Fatalf("expected JSON to not contain 'flex' field for empty struct, got: %s", jsonStr)
	}
}
