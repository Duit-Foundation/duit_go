package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestDecoratedBoxAttributes_Validate_MissingDecoration(t *testing.T) {
	attrs := &duit_attributes.DecoratedBoxAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing decoration, got nil")
	}

	expectedError := "decoration property is required"
	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf("expected error to contain '%s', got: %s", expectedError, err.Error())
	}
}

func TestDecoratedBoxAttributes_Validate_ValidDecoration(t *testing.T) {
	attrs := &duit_attributes.DecoratedBoxAttributes{
		Decoration: &duit_props.BoxDecoration{
			Color: duit_props.NewColorString("#FF0000"),
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid decoration, got:", err)
	}
}

func TestDecoratedBoxAttributes_Validate_WithAllDecorationFields(t *testing.T) {
	attrs := &duit_attributes.DecoratedBoxAttributes{
		Decoration: &duit_props.BoxDecoration{
			Color:        duit_props.NewColorString("#FF0000"),
			BorderRadius: duit_props.NewBorderRadiusAll(duit_props.NewRadiusCircular(10.0)),
			Shape:        duit_props.BoxShapeRectangle,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid decoration with all fields, got:", err)
	}
}

func TestDecoratedBoxAttributes_JSON_WithDecoration(t *testing.T) {
	attrs := &duit_attributes.DecoratedBoxAttributes{
		Decoration: &duit_props.BoxDecoration{
			Color:        duit_props.NewColorString("#FF0000"),
			BorderRadius: duit_props.NewBorderRadiusAll(duit_props.NewRadiusCircular(10.0)),
		},	
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"decoration"`) {
		t.Fatalf("expected JSON to contain 'decoration' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"color":"#FF0000"`) {
		t.Fatalf("expected JSON to contain color value, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"borderRadius":{"radius":10}`) {
		t.Fatalf("expected JSON to contain borderRadius value, got: %s", jsonStr)
	}
}

func TestDecoratedBoxAttributes_JSON_NilDecoration(t *testing.T) {
	attrs := &duit_attributes.DecoratedBoxAttributes{
		Decoration: nil,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with nil decoration, got:", err)
	}

	jsonStr := string(jsonData)
	// Check that omitempty works - nil decoration should not be present
	if strings.Contains(jsonStr, `"decoration"`) {
		t.Fatalf("expected JSON to NOT contain 'decoration' field for nil value, got: %s", jsonStr)
	}
}

func TestDecoratedBoxAttributes_WithRGBOColor_JSON(t *testing.T) {
	rgboColor := duit_props.NewColorRGBO([3]uint8{255, 0, 0}, 1.0)

	attrs := &duit_attributes.DecoratedBoxAttributes{
		Decoration: &duit_props.BoxDecoration{
			Color:        rgboColor,
			BorderRadius: duit_props.NewBorderRadiusAll(duit_props.NewRadiusCircular(5.0)),
		},
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with RGBO color, got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"decoration"`) {
		t.Fatalf("expected JSON to contain 'decoration' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"color":[255,0,0,255]`) {
		t.Fatalf("expected JSON to contain RGBO color array, got: %s", jsonStr)
	}
}

func TestDecoratedBoxAttributes_WithRGBOColor_Validate(t *testing.T) {
	rgboColor := duit_props.NewColorRGBO([3]uint8{128, 64, 192}, 0.8);

	attrs := &duit_attributes.DecoratedBoxAttributes{
		Decoration: &duit_props.BoxDecoration{
			Color: rgboColor,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid RGBO color decoration, got:", err)
	}
}
