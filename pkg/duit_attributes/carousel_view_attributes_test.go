package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	duit_decorations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestCarouselViewAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestCarouselViewAttributes_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		Padding:          duit_props.EdgeInsetsAll(8.0),
		BackgroundColor:  "#FF0000",
		Shape:            nil,
		OverlayColor:     duit_props.MaterialStateProperty[duit_props.ColorString]{},
		Elevation:        4.0,
		ShrinkExtent:     100.0,
		ItemExtent:       200.0,
		ScrollDirection:  duit_flex.Horizontal,
		EnableSplash:     duit_utils.BoolValue(true),
		Reverse:          duit_utils.BoolValue(false),
		ItemSnapping:     duit_utils.BoolValue(true),
		ConsumeMaxWeight: duit_utils.BoolValue(false),
		FlexWeights:      []uint{1, 2, 3},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for EnableSplash Tristate[bool] property serialization
func TestCarouselViewAttributes_EnableSplash_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		EnableSplash: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"enableSplash":true`) {
		t.Fatalf("expected JSON to contain '\"enableSplash\":true', got: %s", jsonStr)
	}
}

func TestCarouselViewAttributes_EnableSplash_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		EnableSplash: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"enableSplash":false`) {
		t.Fatalf("expected JSON to contain '\"enableSplash\":false', got: %s", jsonStr)
	}
}

func TestCarouselViewAttributes_EnableSplash_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		EnableSplash: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"enableSplash"`) {
		t.Fatalf("expected JSON to not contain 'enableSplash' field when nil, got: %s", jsonStr)
	}
}

// Tests for Reverse Tristate[bool] property serialization
func TestCarouselViewAttributes_Reverse_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
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

func TestCarouselViewAttributes_Reverse_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
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

func TestCarouselViewAttributes_Reverse_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
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

// Tests for ItemSnapping Tristate[bool] property serialization
func TestCarouselViewAttributes_ItemSnapping_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ItemSnapping: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"itemSnapping":true`) {
		t.Fatalf("expected JSON to contain '\"itemSnapping\":true', got: %s", jsonStr)
	}
}

func TestCarouselViewAttributes_ItemSnapping_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ItemSnapping: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"itemSnapping":false`) {
		t.Fatalf("expected JSON to contain '\"itemSnapping\":false', got: %s", jsonStr)
	}
}

func TestCarouselViewAttributes_ItemSnapping_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ItemSnapping: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"itemSnapping"`) {
		t.Fatalf("expected JSON to not contain 'itemSnapping' field when nil, got: %s", jsonStr)
	}
}

// Tests for ConsumeMaxWeight Tristate[bool] property serialization
func TestCarouselViewAttributes_ConsumeMaxWeight_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ConsumeMaxWeight: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"consumeMaxWeight":true`) {
		t.Fatalf("expected JSON to contain '\"consumeMaxWeight\":true', got: %s", jsonStr)
	}
}

func TestCarouselViewAttributes_ConsumeMaxWeight_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ConsumeMaxWeight: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"consumeMaxWeight":false`) {
		t.Fatalf("expected JSON to contain '\"consumeMaxWeight\":false', got: %s", jsonStr)
	}
}

func TestCarouselViewAttributes_ConsumeMaxWeight_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.CarouselViewAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ConsumeMaxWeight: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"consumeMaxWeight"`) {
		t.Fatalf("expected JSON to not contain 'consumeMaxWeight' field when nil, got: %s", jsonStr)
	}
}
