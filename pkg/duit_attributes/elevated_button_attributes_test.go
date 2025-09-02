package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	duit_decorations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_material"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestElevatedButtonAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.ElevatedButtonAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString], duit_action.RemoteAction]{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestElevatedButtonAttributes_Validate_WithAllFields(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "button_clicked",
	}

	buttonStyle := &duit_material.ButtonStyle[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		AnimationDuration: 200,
	}

	attrs := &duit_attributes.ElevatedButtonAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString], duit_action.RemoteAction]{
		Autofocus:    duit_utils.BoolValue(true),
		ClipBehavior: duit_props.ClipAntiAlias,
		Style:        buttonStyle,
		OnLongPress:  remoteAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all fields, got:", err)
	}
}

// Tests for Tristate[bool] field Autofocus JSON serialization

func TestElevatedButtonAttributes_Autofocus_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ElevatedButtonAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString], duit_action.RemoteAction]{
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

func TestElevatedButtonAttributes_Autofocus_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ElevatedButtonAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString], duit_action.RemoteAction]{
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

func TestElevatedButtonAttributes_Autofocus_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ElevatedButtonAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString], duit_action.RemoteAction]{
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

// Tests for other fields JSON serialization

func TestElevatedButtonAttributes_ClipBehavior_JSON(t *testing.T) {
	attrs := &duit_attributes.ElevatedButtonAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString], duit_action.RemoteAction]{
		ClipBehavior: duit_props.ClipAntiAlias,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"clipBehavior":"antiAlias"`) {
		t.Fatalf("expected JSON to contain clipBehavior value, got: %s", jsonStr)
	}
}

func TestElevatedButtonAttributes_Style_JSON(t *testing.T) {
	buttonStyle := &duit_material.ButtonStyle[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		AnimationDuration: 300,
	}

	attrs := &duit_attributes.ElevatedButtonAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString], duit_action.RemoteAction]{
		Style: buttonStyle,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"style"`) {
		t.Fatalf("expected JSON to contain 'style' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"animationDuration":300`) {
		t.Fatalf("expected JSON to contain animation duration, got: %s", jsonStr)
	}
}

func TestElevatedButtonAttributes_OnLongPress_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "long_press_event",
	}

	attrs := &duit_attributes.ElevatedButtonAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString], duit_action.RemoteAction]{
		OnLongPress: remoteAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onLongPress"`) {
		t.Fatalf("expected JSON to contain 'onLongPress' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"event":"long_press_event"`) {
		t.Fatalf("expected JSON to contain event value, got: %s", jsonStr)
	}
}

func TestElevatedButtonAttributes_NilFields_JSON(t *testing.T) {
	attrs := &duit_attributes.ElevatedButtonAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString], duit_action.RemoteAction]{
		Style:       nil,
		OnLongPress: nil,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with nil fields, got:", err)
	}

	jsonStr := string(jsonData)
	// Check that omitempty works - nil fields should not be present
	if strings.Contains(jsonStr, `"style"`) {
		t.Fatalf("expected JSON to NOT contain 'style' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"onLongPress"`) {
		t.Fatalf("expected JSON to NOT contain 'onLongPress' field for nil value, got: %s", jsonStr)
	}
}

// Test with different action types

func TestElevatedButtonAttributes_WithLocalAction_JSON(t *testing.T) {
	localAction := &duit_action.LocalAction{
		ExecutionType: duit_action.Local,
		Payload:       map[string]interface{}{"key": "value"},
	}

	attrs := &duit_attributes.ElevatedButtonAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString], duit_action.LocalAction]{
		OnLongPress: localAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with LocalAction, got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onLongPress"`) {
		t.Fatalf("expected JSON to contain 'onLongPress' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"executionType":1`) {
		t.Fatalf("expected JSON to contain Local execution type, got: %s", jsonStr)
	}
}

func TestElevatedButtonAttributes_WithScriptAction_JSON(t *testing.T) {
	scriptAction := &duit_action.ScriptAction{
		ExecutionType: duit_action.Script,
		Event:         "script_event",
		Script: &duit_action.ScriptData{
			SourceCode:   "console.log('test')",
			FunctionName: "testFunction",
		},
	}

	attrs := &duit_attributes.ElevatedButtonAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString], duit_action.ScriptAction]{
		OnLongPress: scriptAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with ScriptAction, got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onLongPress"`) {
		t.Fatalf("expected JSON to contain 'onLongPress' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"executionType":2`) {
		t.Fatalf("expected JSON to contain Script execution type, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"sourceCode":"console.log('test')"`) {
		t.Fatalf("expected JSON to contain script source code, got: %s", jsonStr)
	}
}
