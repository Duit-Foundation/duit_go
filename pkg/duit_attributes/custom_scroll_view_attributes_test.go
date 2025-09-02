package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_gestures"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestCustomScrollViewAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.CustomScrollViewAttributes{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestCustomScrollViewAttributes_Validate_WithAllFields(t *testing.T) {
	attrs := &duit_attributes.CustomScrollViewAttributes{
		Physics:                 duit_gestures.BouncingScrollPhysics,
		Reverse:                 duit_utils.BoolValue(true),
		Primary:                 duit_utils.BoolValue(false),
		ShrinkWrap:              duit_utils.BoolValue(true),
		ScrollDirection:         duit_flex.Vertical,
		Anchor:                  0.5,
		SemantickChildCount:     10,
		ClipBehavior:            duit_props.ClipAntiAlias,
		RestorationId:           "scroll_view_1",
		DragStarnBehavior:       duit_gestures.Start,
		KeyboardDismissBehavior: duit_gestures.Manual,
		HitTestBehavior:         duit_gestures.Opaque,
		CacheExtent:             250.0,
		Center:                  "center_key",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all fields, got:", err)
	}
}

// Tests for Tristate[bool] fields JSON serialization

func TestCustomScrollViewAttributes_Reverse_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CustomScrollViewAttributes{
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

func TestCustomScrollViewAttributes_Reverse_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CustomScrollViewAttributes{
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

func TestCustomScrollViewAttributes_Reverse_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.CustomScrollViewAttributes{
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

func TestCustomScrollViewAttributes_Primary_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CustomScrollViewAttributes{
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

func TestCustomScrollViewAttributes_Primary_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CustomScrollViewAttributes{
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

func TestCustomScrollViewAttributes_Primary_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.CustomScrollViewAttributes{
		Primary: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"primary"`) {
		t.Fatalf("expected JSON to not contain 'primary' field when nil, got: %s", jsonStr)
	}
}

func TestCustomScrollViewAttributes_ShrinkWrap_JSON_True(t *testing.T) {
	attrs := &duit_attributes.CustomScrollViewAttributes{
		ShrinkWrap: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"shrinkWrap":true`) {
		t.Fatalf("expected JSON to contain '\"shrinkWrap\":true', got: %s", jsonStr)
	}
}

func TestCustomScrollViewAttributes_ShrinkWrap_JSON_False(t *testing.T) {
	attrs := &duit_attributes.CustomScrollViewAttributes{
		ShrinkWrap: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"shrinkWrap":false`) {
		t.Fatalf("expected JSON to contain '\"shrinkWrap\":false', got: %s", jsonStr)
	}
}

func TestCustomScrollViewAttributes_ShrinkWrap_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.CustomScrollViewAttributes{
		ShrinkWrap: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"shrinkWrap"`) {
		t.Fatalf("expected JSON to not contain 'shrinkWrap' field when nil, got: %s", jsonStr)
	}
}
