package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestFlexibleSpaceBarAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.FlexibleSpaceBarAttributes[duit_props.EdgeInsetsAll]{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

// Tests for CenterTitle Tristate[bool] property serialization
func TestFlexibleSpaceBarAttributes_CenterTitle_JSON_True(t *testing.T) {
	attrs := &duit_attributes.FlexibleSpaceBarAttributes[duit_props.EdgeInsetsAll]{
		CenterTitle: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"centerTitle":true`) {
		t.Fatalf("expected JSON to contain '\"centerTitle\":true', got: %s", jsonStr)
	}
}

func TestFlexibleSpaceBarAttributes_CenterTitle_JSON_False(t *testing.T) {
	attrs := &duit_attributes.FlexibleSpaceBarAttributes[duit_props.EdgeInsetsAll]{
		CenterTitle: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"centerTitle":false`) {
		t.Fatalf("expected JSON to contain '\"centerTitle\":false', got: %s", jsonStr)
	}
}

func TestFlexibleSpaceBarAttributes_CenterTitle_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.FlexibleSpaceBarAttributes[duit_props.EdgeInsetsAll]{
		CenterTitle: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"centerTitle"`) {
		t.Fatalf("expected JSON to not contain 'centerTitle' field when nil, got: %s", jsonStr)
	}
}