package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decorations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestCardAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.CardAttributes[duit_color.ColorString, duit_edge_insets.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_color.ColorString]]{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestCardAttributes_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.CardAttributes[duit_color.ColorString, duit_edge_insets.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_color.ColorString]]{
		Color:              "#FF0000",
		ShadowColor:        "#000000",
		Elevation:          4.0,
		BorderOnForeground: duit_utils.BoolValue(true),
		SemanticContainer:  duit_utils.BoolValue(false),
		Margin:             duit_edge_insets.EdgeInsetsAll(8.0),
		Shape:              nil,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for BorderOnForeground Tristate[bool] property serialization
func TestCardAttributes_BorderOnForeground_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CardAttributes[duit_color.ColorString, duit_edge_insets.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_color.ColorString]]{
		BorderOnForeground: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"borderOnForeground":true`) {
		t.Fatalf("expected JSON to contain '\"borderOnForeground\":true', got: %s", jsonStr)
	}
}

func TestCardAttributes_BorderOnForeground_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CardAttributes[duit_color.ColorString, duit_edge_insets.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_color.ColorString]]{
		BorderOnForeground: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"borderOnForeground":false`) {
		t.Fatalf("expected JSON to contain '\"borderOnForeground\":false', got: %s", jsonStr)
	}
}

// Tests for SemanticContainer Tristate[bool] property serialization
func TestCardAttributes_SemanticContainer_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CardAttributes[duit_color.ColorString, duit_edge_insets.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_color.ColorString]]{
		SemanticContainer: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"semanticContainer":true`) {
		t.Fatalf("expected JSON to contain '\"semanticContainer\":true', got: %s", jsonStr)
	}
}

func TestCardAttributes_SemanticContainer_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CardAttributes[duit_color.ColorString, duit_edge_insets.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_color.ColorString]]{
		SemanticContainer: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"semanticContainer":false`) {
		t.Fatalf("expected JSON to contain '\"semanticContainer\":false', got: %s", jsonStr)
	}
}

func TestCardAttributes_SemanticContainer_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.CardAttributes[duit_color.ColorString, duit_edge_insets.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_color.ColorString]]{
		SemanticContainer: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"semanticContainer"`) {
		t.Fatalf("expected JSON to not contain 'semanticContainer' field when nil, got: %s", jsonStr)
	}
}
