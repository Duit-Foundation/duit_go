package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_material"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestInkwellAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestInkwellAttributes_Validate_WithAllFields(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "inkwell_tapped",
	}

	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		OnTap:                remoteAction,
		OnDoubleTap:          remoteAction,
		OnLongPress:          remoteAction,
		OnTapDown:            remoteAction,
		OnTapUp:              remoteAction,
		OnTapCancel:          remoteAction,
		OnSecondaryTapDown:   remoteAction,
		OnSecondaryTapCancel: remoteAction,
		OnSecondaryTap:       remoteAction,
		OnSecondaryTapUp:     remoteAction,
		FocusColor:           "#FF0000",
		HoverColor:           "#00FF00",
		HighlightColor:       "#0000FF",
		SplashColor:          "#FFFF00",
		Radius:               10.5,
		HoverDuration:        300,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all fields, got:", err)
	}
}

func TestInkwellAttributes_Validate_WithMaterialStateProperty(t *testing.T) {
	overlayColor := &duit_material.MaterialStateProperty[duit_color.ColorString]{
		Hovered: "#FF0000",
	}

	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		OverlayColor: overlayColor,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with MaterialStateProperty, got:", err)
	}
}

func TestInkwellAttributes_Validate_WithBorderRadius(t *testing.T) {
	borderRadius := &duit_decoration.BorderRadius{
		TopLeft:     &duit_decoration.Radius{Radius: 10.0},
		TopRight:    &duit_decoration.Radius{Radius: 15.0},
		BottomLeft:  &duit_decoration.Radius{Radius: 20.0},
		BottomRight: &duit_decoration.Radius{Radius: 25.0},
	}

	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		BorderRadius: borderRadius,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with BorderRadius, got:", err)
	}
}

func TestInkwellAttributes_Validate_WithCustomBorder(t *testing.T) {
	customBorder := &duit_decoration.RoundedRectangleBorder[duit_color.ColorString]{
		Side: &duit_decoration.BorderSide[duit_color.ColorString]{
			Color: "#FF0000",
			Width: 2.0,
		},
		BorderRadius: &duit_decoration.BorderRadius{
			TopLeft: &duit_decoration.Radius{Radius: 8.0},
		},
	}

	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		CustomBorder: customBorder,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with CustomBorder, got:", err)
	}
}

// Tests for EnableFeedback Tristate[bool] property serialization
func TestInkwellAttributes_EnableFeedback_JSON_True(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		EnableFeedback: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"enableFeedback":true`) {
		t.Fatalf("expected JSON to contain '\"enableFeedback\":true', got: %s", jsonStr)
	}
}

func TestInkwellAttributes_EnableFeedback_JSON_False(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		EnableFeedback: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"enableFeedback":false`) {
		t.Fatalf("expected JSON to contain '\"enableFeedback\":false', got: %s", jsonStr)
	}
}

func TestInkwellAttributes_EnableFeedback_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		EnableFeedback: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"enableFeedback"`) {
		t.Fatalf("expected JSON to not contain 'enableFeedback' field when nil, got: %s", jsonStr)
	}
}

// Tests for ExcludeFromSemantics Tristate[bool] property serialization
func TestInkwellAttributes_ExcludeFromSemantics_JSON_True(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
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

func TestInkwellAttributes_ExcludeFromSemantics_JSON_False(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
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

func TestInkwellAttributes_ExcludeFromSemantics_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
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

// Tests for Autofocus Tristate[bool] property serialization
func TestInkwellAttributes_Autofocus_JSON_True(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		Autofocus: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"autofocus":true`) {
		t.Fatalf("expected JSON to contain '\"autofocus\":true', got: %s", jsonStr)
	}
}

func TestInkwellAttributes_Autofocus_JSON_False(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		Autofocus: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"autofocus":false`) {
		t.Fatalf("expected JSON to contain '\"autofocus\":false', got: %s", jsonStr)
	}
}

func TestInkwellAttributes_Autofocus_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		Autofocus: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"autofocus"`) {
		t.Fatalf("expected JSON to not contain 'autofocus' field when nil, got: %s", jsonStr)
	}
}

// Tests for CanRequestFocus Tristate[bool] property serialization
func TestInkwellAttributes_CanRequestFocus_JSON_True(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		CanRequestFocus: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"canRequestFocus":true`) {
		t.Fatalf("expected JSON to contain '\"canRequestFocus\":true', got: %s", jsonStr)
	}
}

func TestInkwellAttributes_CanRequestFocus_JSON_False(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		CanRequestFocus: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"canRequestFocus":false`) {
		t.Fatalf("expected JSON to contain '\"canRequestFocus\":false', got: %s", jsonStr)
	}
}

func TestInkwellAttributes_CanRequestFocus_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		CanRequestFocus: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"canRequestFocus"`) {
		t.Fatalf("expected JSON to not contain 'canRequestFocus' field when nil, got: %s", jsonStr)
	}
}

// Tests for other fields JSON serialization
func TestInkwellAttributes_OnTap_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "test_tap",
	}

	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		OnTap: remoteAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onTap"`) {
		t.Fatalf("expected JSON to contain 'onTap' field, got: %s", jsonStr)
	}
}

func TestInkwellAttributes_FocusColor_JSON(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		FocusColor: "#FF0000",
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"focusColor":"#FF0000"`) {
		t.Fatalf("expected JSON to contain '\"focusColor\":\"#FF0000\"', got: %s", jsonStr)
	}
}

func TestInkwellAttributes_Radius_JSON(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		Radius: 15.5,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"radius":15.5`) {
		t.Fatalf("expected JSON to contain '\"radius\":15.5', got: %s", jsonStr)
	}
}

func TestInkwellAttributes_HoverDuration_JSON(t *testing.T) {
	attrs := &duit_attributes.InkwellAttributes[duit_color.ColorString, duit_action.RemoteAction, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		HoverDuration: 500,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"hoverDuration":500`) {
		t.Fatalf("expected JSON to contain '\"hoverDuration\":500', got: %s", jsonStr)
	}
}
