package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestClipRectAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.ClipRectAttributes{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestClipRectAttributes_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.ClipRectAttributes{
		ClipBehavior: duit_props.ClipAntiAlias,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

func TestClipRectAttributes_SetClipBehavior(t *testing.T) {
	attrs := duit_attributes.NewClipRectAttributes()
	clipBehavior := duit_props.ClipHardEdge

	result := attrs.SetClipBehavior(clipBehavior)

	if attrs.ClipBehavior != clipBehavior {
		t.Fatalf("expected ClipBehavior to be %s, got: %s", clipBehavior, attrs.ClipBehavior)
	}

	if result != attrs {
		t.Fatal("expected SetClipBehavior to return the same instance for method chaining")
	}
}

func TestClipRectAttributes_ClipBehavior_JSON_WithValue(t *testing.T) {
	attrs := &duit_attributes.ClipRectAttributes{
		ClipBehavior: duit_props.ClipAntiAlias,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"antiAlias"`) {
		t.Fatalf("expected JSON to contain clipBehavior value, got: %s", jsonStr)
	}
}

func TestClipRectAttributes_ClipBehavior_JSON_EmptyValue(t *testing.T) {
	attrs := &duit_attributes.ClipRectAttributes{
		ClipBehavior: "",
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	// With omitempty tag, empty string should be omitted
	if strings.Contains(jsonStr, `"color"`) {
		t.Fatalf("expected JSON to not contain 'color' field when empty, got: %s", jsonStr)
	}
}

