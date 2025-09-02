package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_painting"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestImageAttributes_Validate_AssetType_Valid(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type: duit_painting.Asset,
		Src:  "assets/image.png",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid asset image attributes, got:", err)
	}
}

func TestImageAttributes_Validate_AssetType_MissingSrc(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type: duit_painting.Asset,
		Src:  "",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for asset type without src")
	}
}

func TestImageAttributes_Validate_NetworkType_Valid(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type: duit_painting.Network,
		Src:  "https://example.com/image.jpg",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid network image attributes, got:", err)
	}
}

func TestImageAttributes_Validate_NetworkType_MissingSrc(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type: duit_painting.Network,
		Src:  "",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for network type without src")
	}
}

func TestImageAttributes_Validate_MemoryType_Valid(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:     duit_painting.Memory,
		ByteData: &duit_painting.ImageBuffer{},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid memory image attributes, got:", err)
	}
}

func TestImageAttributes_Validate_MemoryType_MissingByteData(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:     duit_painting.Memory,
		ByteData: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for memory type without byteData")
	}
}

func TestImageAttributes_Validate_MissingType(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type: "",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing type")
	}
}

func TestImageAttributes_Validate_InvalidType(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type: "invalid_type",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for invalid type")
	}
}

func TestImageAttributes_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:                 duit_painting.Asset,
		Src:                  "assets/image.png",
		Width:                100.0,
		Height:               200.0,
		CacheWidth:           50,
		CacheHeight:          100,
		Scale:                1.5,
		Repeat:               duit_painting.NoRepeat,
		IsAntiAlias:          duit_utils.BoolValue(true),
		MatchTextDirection:   duit_utils.BoolValue(false),
		GaplessPlayback:      duit_utils.BoolValue(true),
		ExcludeFromSemantics: duit_utils.BoolValue(false),
		FilterQuality:        "high",
		ColorBlendMode:       duit_painting.SrcOver,
		Color:                "#FF0000",
		Fit:                  duit_flex.Contain,
		Alignment:            duit_props.AlignmentCenter,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for IsAntiAlias Tristate[bool] property serialization
func TestImageAttributes_IsAntiAlias_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:        duit_painting.Asset,
		Src:         "assets/image.png",
		IsAntiAlias: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"isAntiAlias":true`) {
		t.Fatalf("expected JSON to contain '\"isAntiAlias\":true', got: %s", jsonStr)
	}
}

func TestImageAttributes_IsAntiAlias_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:        duit_painting.Asset,
		Src:         "assets/image.png",
		IsAntiAlias: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"isAntiAlias":false`) {
		t.Fatalf("expected JSON to contain '\"isAntiAlias\":false', got: %s", jsonStr)
	}
}

func TestImageAttributes_IsAntiAlias_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:        duit_painting.Asset,
		Src:         "assets/image.png",
		IsAntiAlias: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"isAntiAlias"`) {
		t.Fatalf("expected JSON to not contain 'isAntiAlias' field when nil, got: %s", jsonStr)
	}
}

// Tests for MatchTextDirection Tristate[bool] property serialization
func TestImageAttributes_MatchTextDirection_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:               duit_painting.Asset,
		Src:                "assets/image.png",
		MatchTextDirection: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"matchTextDirection":true`) {
		t.Fatalf("expected JSON to contain '\"matchTextDirection\":true', got: %s", jsonStr)
	}
}

func TestImageAttributes_MatchTextDirection_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:               duit_painting.Asset,
		Src:                "assets/image.png",
		MatchTextDirection: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"matchTextDirection":false`) {
		t.Fatalf("expected JSON to contain '\"matchTextDirection\":false', got: %s", jsonStr)
	}
}

func TestImageAttributes_MatchTextDirection_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:               duit_painting.Asset,
		Src:                "assets/image.png",
		MatchTextDirection: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"matchTextDirection"`) {
		t.Fatalf("expected JSON to not contain 'matchTextDirection' field when nil, got: %s", jsonStr)
	}
}

// Tests for GaplessPlayback Tristate[bool] property serialization
func TestImageAttributes_GaplessPlayback_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:            duit_painting.Asset,
		Src:             "assets/image.png",
		GaplessPlayback: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"gaplessPlayback":true`) {
		t.Fatalf("expected JSON to contain '\"gaplessPlayback\":true', got: %s", jsonStr)
	}
}

func TestImageAttributes_GaplessPlayback_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:            duit_painting.Asset,
		Src:             "assets/image.png",
		GaplessPlayback: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"gaplessPlayback":false`) {
		t.Fatalf("expected JSON to contain '\"gaplessPlayback\":false', got: %s", jsonStr)
	}
}

func TestImageAttributes_GaplessPlayback_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:            duit_painting.Asset,
		Src:             "assets/image.png",
		GaplessPlayback: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"gaplessPlayback"`) {
		t.Fatalf("expected JSON to not contain 'gaplessPlayback' field when nil, got: %s", jsonStr)
	}
}

// Tests for ExcludeFromSemantics Tristate[bool] property serialization
func TestImageAttributes_ExcludeFromSemantics_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:                 duit_painting.Asset,
		Src:                  "assets/image.png",
		ExcludeFromSemantics: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"excludeFromSemantics":true`) {
		t.Fatalf("expected JSON to contain '\"excludeFromSemantics\":true', got: %s", jsonStr)
	}
}

func TestImageAttributes_ExcludeFromSemantics_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:                 duit_painting.Asset,
		Src:                  "assets/image.png",
		ExcludeFromSemantics: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"excludeFromSemantics":false`) {
		t.Fatalf("expected JSON to contain '\"excludeFromSemantics\":false', got: %s", jsonStr)
	}
}

func TestImageAttributes_ExcludeFromSemantics_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ImageAttributes[duit_props.ColorString]{
		Type:                 duit_painting.Asset,
		Src:                  "assets/image.png",
		ExcludeFromSemantics: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"excludeFromSemantics"`) {
		t.Fatalf("expected JSON to not contain 'excludeFromSemantics' field when nil, got: %s", jsonStr)
	}
}
