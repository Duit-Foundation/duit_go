package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestPaddingAttributes_Validate_ValidAttributes(t *testing.T) {
	padding := &duit_props.EdgeInsetsV2{}
	padding = padding.LTRB(10.0, 20.0, 15.0, 25.0)

	attrs := &duit_attributes.PaddingAttributes{
		Padding: padding,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestPaddingAttributes_Validate_WithoutPadding(t *testing.T) {
	attrs := &duit_attributes.PaddingAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when padding is not set")
	}
}

func TestPaddingAttributes_Validate_WithZeroPadding(t *testing.T) {
	padding := &duit_props.EdgeInsetsV2{}
	padding = padding.LTRB(0.0, 0.0, 0.0, 0.0)

	attrs := &duit_attributes.PaddingAttributes{
		Padding: padding,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for zero padding, got:", err)
	}
}

func TestPaddingAttributes_Validate_WithPartialPadding(t *testing.T) {
	padding := &duit_props.EdgeInsetsV2{}
	padding = padding.LTRB(10.0, 20.0, 0.0, 0.0)

	attrs := &duit_attributes.PaddingAttributes{
		Padding: padding,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for partial padding, got:", err)
	}
}

func TestPaddingAttributes_MarshalJSON_WithPadding(t *testing.T) {
	padding := &duit_props.EdgeInsetsV2{}
	padding = padding.LTRB(10.0, 20.0, 15.0, 25.0)

	attrs := &duit_attributes.PaddingAttributes{
		Padding: padding,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"padding"`) {
		t.Fatalf("expected JSON to contain 'padding' field, got: %s", jsonStr)
	}
}

func TestPaddingAttributes_MarshalJSON_WithoutPadding(t *testing.T) {
	attrs := &duit_attributes.PaddingAttributes{}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	// When padding is nil, it should be omitted from JSON due to omitempty tag
	if strings.Contains(jsonStr, `"padding"`) {
		t.Fatalf("expected JSON to not contain 'padding' field when nil, got: %s", jsonStr)
	}
}