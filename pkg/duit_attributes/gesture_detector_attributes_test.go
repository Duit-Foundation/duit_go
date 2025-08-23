package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_gestures"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestGestureDetectorAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestGestureDetectorAttributes_Validate_WithValidActions(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "tap_event",
	}

	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
		OnTap:       remoteAction,
		OnDoubleTap: remoteAction,
		OnLongPress: remoteAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with actions, got:", err)
	}
}

func TestGestureDetectorAttributes_Validate_WithAllGestureActions(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "gesture_event",
	}

	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
		OnTap:                 remoteAction,
		OnTapDown:             remoteAction,
		OnTapUp:               remoteAction,
		OnTapCancel:           remoteAction,
		OnDoubleTap:           remoteAction,
		OnDoubleTapDown:       remoteAction,
		OnDoubleTapCancel:     remoteAction,
		OnLongPressDown:       remoteAction,
		OnLongPressCancel:     remoteAction,
		OnLongPress:           remoteAction,
		OnLongPressStart:      remoteAction,
		OnLongPressMoveUpdate: remoteAction,
		OnLongPressUp:         remoteAction,
		OnLongPressEnd:        remoteAction,
		OnPanStart:            remoteAction,
		OnPanDown:             remoteAction,
		OnPanUpdate:           remoteAction,
		OnPanEnd:              remoteAction,
		OnPanCancel:           remoteAction,
		ExcludeFromSemantics:  duit_utils.BoolValue(true),
		DragStarnBehavior:     duit_gestures.Start,
		Behavior:              duit_gestures.Opaque,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all actions, got:", err)
	}
}

// Tests for Tristate[bool] field ExcludeFromSemantics JSON serialization

func TestGestureDetectorAttributes_ExcludeFromSemantics_JSON_True(t *testing.T) {
	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
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

func TestGestureDetectorAttributes_ExcludeFromSemantics_JSON_False(t *testing.T) {
	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
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

func TestGestureDetectorAttributes_ExcludeFromSemantics_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
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

// Tests for gesture action fields JSON serialization

func TestGestureDetectorAttributes_OnTap_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "tap_clicked",
	}

	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
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
	if !strings.Contains(jsonStr, `"event":"tap_clicked"`) {
		t.Fatalf("expected JSON to contain event value, got: %s", jsonStr)
	}
}

func TestGestureDetectorAttributes_OnDoubleTap_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "double_tap_event",
	}

	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
		OnDoubleTap: remoteAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onDoubleTap"`) {
		t.Fatalf("expected JSON to contain 'onDoubleTap' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"event":"double_tap_event"`) {
		t.Fatalf("expected JSON to contain event value, got: %s", jsonStr)
	}
}

func TestGestureDetectorAttributes_OnLongPress_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "long_press_event",
	}

	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
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

func TestGestureDetectorAttributes_PanGestures_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "pan_event",
	}

	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
		OnPanStart:  remoteAction,
		OnPanUpdate: remoteAction,
		OnPanEnd:    remoteAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onPanStart"`) {
		t.Fatalf("expected JSON to contain 'onPanStart' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"onPanUpdate"`) {
		t.Fatalf("expected JSON to contain 'onPanUpdate' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"onPanEnd"`) {
		t.Fatalf("expected JSON to contain 'onPanEnd' field, got: %s", jsonStr)
	}
}

// Tests for enum fields JSON serialization

func TestGestureDetectorAttributes_DragStartBehavior_JSON(t *testing.T) {
	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
		DragStarnBehavior: duit_gestures.Down,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"dragStartBehavior":"down"`) {
		t.Fatalf("expected JSON to contain dragStartBehavior value, got: %s", jsonStr)
	}
}

func TestGestureDetectorAttributes_Behavior_JSON(t *testing.T) {
	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
		Behavior: duit_gestures.Translucent,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"behavior":"translucent"`) {
		t.Fatalf("expected JSON to contain behavior value, got: %s", jsonStr)
	}
}

func TestGestureDetectorAttributes_NilActions_JSON(t *testing.T) {
	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
		OnTap:       nil,
		OnDoubleTap: nil,
		OnLongPress: nil,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with nil actions, got:", err)
	}

	jsonStr := string(jsonData)
	// Check that omitempty works - nil fields should not be present
	if strings.Contains(jsonStr, `"onTap"`) {
		t.Fatalf("expected JSON to NOT contain 'onTap' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"onDoubleTap"`) {
		t.Fatalf("expected JSON to NOT contain 'onDoubleTap' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"onLongPress"`) {
		t.Fatalf("expected JSON to NOT contain 'onLongPress' field for nil value, got: %s", jsonStr)
	}
}

// Test with different action types

func TestGestureDetectorAttributes_WithLocalAction_JSON(t *testing.T) {
	localAction := &duit_action.LocalAction{
		ExecutionType: duit_action.Local,
		Payload:       map[string]interface{}{"gesture": "tap"},
	}

	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.LocalAction]{
		OnTap: localAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with LocalAction, got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onTap"`) {
		t.Fatalf("expected JSON to contain 'onTap' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"executionType":1`) {
		t.Fatalf("expected JSON to contain Local execution type, got: %s", jsonStr)
	}
}

func TestGestureDetectorAttributes_WithScriptAction_JSON(t *testing.T) {
	scriptAction := &duit_action.ScriptAction{
		ExecutionType: duit_action.Script,
		Event:         "script_gesture",
		Script: &duit_action.ScriptData{
			SourceCode:   "handleGesture()",
			FunctionName: "handleGesture",
		},
	}

	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.ScriptAction]{
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
	if !strings.Contains(jsonStr, `"sourceCode":"handleGesture()"`) {
		t.Fatalf("expected JSON to contain script source code, got: %s", jsonStr)
	}
}

// Test comprehensive gesture detector with all fields

func TestGestureDetectorAttributes_ComprehensiveTest_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "comprehensive_event",
	}

	attrs := &duit_attributes.GestureDetectorAttributes[duit_action.RemoteAction]{
		OnTap:                 remoteAction,
		OnTapDown:             remoteAction,
		OnDoubleTap:           remoteAction,
		OnLongPress:           remoteAction,
		OnPanStart:            remoteAction,
		ExcludeFromSemantics:  duit_utils.BoolValue(false),
		DragStarnBehavior:     duit_gestures.Start,
		Behavior:              duit_gestures.DeferToChild,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for comprehensive validation, got:", err)
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for comprehensive MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onTap"`) {
		t.Fatalf("expected JSON to contain 'onTap' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"excludeFromSemantics":false`) {
		t.Fatalf("expected JSON to contain excludeFromSemantics value, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"dragStartBehavior":"start"`) {
		t.Fatalf("expected JSON to contain dragStartBehavior value, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"behavior":"deferToChild"`) {
		t.Fatalf("expected JSON to contain behavior value, got: %s", jsonStr)
	}
}
