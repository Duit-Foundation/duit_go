package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestSingleChildScrollViewAttributes_Validate_ValidAttributes(t *testing.T) {	
	attrs := &duit_attributes.SingleChildScrollViewAttributes{
		Reverse: duit_utils.BoolValue(true),
		Primary: duit_utils.BoolValue(false),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSingleChildScrollViewAttributes_Validate_WithAllProperties(t *testing.T) {
	padding := duit_props.NewEdgeInsetsLTRB(10.0, 20.0, 30.0, 40.0)
	
	attrs := &duit_attributes.SingleChildScrollViewAttributes{
		ScrollDirection:         "vertical",
		Reverse:                 duit_utils.BoolValue(true),
		Primary:                 duit_utils.BoolValue(true),
		Padding:                 padding,
		RestorationId:           "test_restoration_id",
		ClipBehavior:            "antiAlias",
		DragStartBehavior:       "start",
		KeyboardDismissBehavior: "manual",
		Physics:                 "bouncing",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for Reverse Tristate[bool] property serialization
func TestSingleChildScrollViewAttributes_Reverse_JSON_True(t *testing.T) {
	attrs := &duit_attributes.SingleChildScrollViewAttributes{
		Reverse: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"reverse":true`) {
		t.Fatalf("expected JSON to contain '\"reverse\":true', got: %s", jsonStr)
	}
}

func TestSingleChildScrollViewAttributes_Reverse_JSON_False(t *testing.T) {
	attrs := &duit_attributes.SingleChildScrollViewAttributes{
		Reverse: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"reverse":false`) {
		t.Fatalf("expected JSON to contain '\"reverse\":false', got: %s", jsonStr)
	}
}

func TestSingleChildScrollViewAttributes_Reverse_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.SingleChildScrollViewAttributes{
		Reverse: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"reverse"`) {
		t.Fatalf("expected JSON to not contain 'reverse' field when nil, got: %s", jsonStr)
	}
}

// Tests for Primary Tristate[bool] property serialization
func TestSingleChildScrollViewAttributes_Primary_JSON_True(t *testing.T) {
	attrs := &duit_attributes.SingleChildScrollViewAttributes{
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

func TestSingleChildScrollViewAttributes_Primary_JSON_False(t *testing.T) {
	attrs := &duit_attributes.SingleChildScrollViewAttributes{
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
