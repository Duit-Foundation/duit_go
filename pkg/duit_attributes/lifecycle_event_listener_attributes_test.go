package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
)

func TestLifecycleEventListenerAttributes_Validate_ValidAttributes_OnStateChanged(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "state_changed",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnStateChanged: remoteAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with OnStateChanged, got:", err)
	}
}

func TestLifecycleEventListenerAttributes_Validate_ValidAttributes_OnResumed(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "resumed",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnResumed: remoteAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with OnResumed, got:", err)
	}
}

func TestLifecycleEventListenerAttributes_Validate_ValidAttributes_OnInactive(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "inactive",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnInactive: remoteAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with OnInactive, got:", err)
	}
}

func TestLifecycleEventListenerAttributes_Validate_ValidAttributes_OnPaused(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "paused",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnPaused: remoteAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with OnPaused, got:", err)
	}
}

func TestLifecycleEventListenerAttributes_Validate_ValidAttributes_OnDetached(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "detached",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnDetached: remoteAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with OnDetached, got:", err)
	}
}

func TestLifecycleEventListenerAttributes_Validate_ValidAttributes_OnHidden(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "hidden",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnHidden: remoteAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with OnHidden, got:", err)
	}
}

func TestLifecycleEventListenerAttributes_Validate_ValidAttributes_MultipleCallbacks(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "lifecycle_event",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnStateChanged: remoteAction,
		OnResumed:      remoteAction,
		OnInactive:     remoteAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with multiple callbacks, got:", err)
	}
}

func TestLifecycleEventListenerAttributes_Validate_ValidAttributes_AllCallbacks(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "lifecycle_event",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnStateChanged: remoteAction,
		OnResumed:      remoteAction,
		OnInactive:     remoteAction,
		OnPaused:       remoteAction,
		OnDetached:     remoteAction,
		OnHidden:       remoteAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all callbacks, got:", err)
	}
}

func TestLifecycleEventListenerAttributes_Validate_InvalidAttributes_NoCallbacks(t *testing.T) {
	attrs := &duit_attributes.LifecycleEventListenerAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for attributes with no callbacks")
	}
}

func TestLifecycleEventListenerAttributes_Validate_InvalidAttributes_AllNil(t *testing.T) {
	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnStateChanged: nil,
		OnResumed:      nil,
		OnInactive:     nil,
		OnPaused:       nil,
		OnDetached:     nil,
		OnHidden:       nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for attributes with all nil callbacks")
	}
}

func TestLifecycleEventListenerAttributes_Validate_WithLocalAction(t *testing.T) {
	localAction := &duit_action.LocalAction{
		ExecutionType: duit_action.Transport,
		Payload:       "local_lifecycle",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnStateChanged: localAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with LocalAction, got:", err)
	}
}

func TestLifecycleEventListenerAttributes_Validate_WithScriptAction(t *testing.T) {
	scriptAction := &duit_action.ScriptAction{
		ExecutionType: duit_action.Transport,
		Event:         "script_lifecycle",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnStateChanged: scriptAction,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with ScriptAction, got:", err)
	}
}

// Tests for JSON serialization
func TestLifecycleEventListenerAttributes_OnStateChanged_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "test_state_changed",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnStateChanged: remoteAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onStateChanged"`) {
		t.Fatalf("expected JSON to contain 'onStateChanged' field, got: %s", jsonStr)
	}
}

func TestLifecycleEventListenerAttributes_OnResumed_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "test_resumed",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnResumed: remoteAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onResumed"`) {
		t.Fatalf("expected JSON to contain 'onResumed' field, got: %s", jsonStr)
	}
}

func TestLifecycleEventListenerAttributes_OnInactive_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "test_inactive",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnInactive: remoteAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onInactive"`) {
		t.Fatalf("expected JSON to contain 'onInactive' field, got: %s", jsonStr)
	}
}

func TestLifecycleEventListenerAttributes_OnPaused_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "test_paused",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnPaused: remoteAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onPaused"`) {
		t.Fatalf("expected JSON to contain 'onPaused' field, got: %s", jsonStr)
	}
}

func TestLifecycleEventListenerAttributes_OnDetached_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "test_detached",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnDetached: remoteAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onDetached"`) {
		t.Fatalf("expected JSON to contain 'onDetached' field, got: %s", jsonStr)
	}
}

func TestLifecycleEventListenerAttributes_OnHidden_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "test_hidden",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnHidden: remoteAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onHidden"`) {
		t.Fatalf("expected JSON to contain 'onHidden' field, got: %s", jsonStr)
	}
}

func TestLifecycleEventListenerAttributes_MultipleCallbacks_JSON(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "test_multiple",
	}

	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnStateChanged: remoteAction,
		OnResumed:      remoteAction,
		OnInactive:     remoteAction,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"onStateChanged"`) {
		t.Fatalf("expected JSON to contain 'onStateChanged' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"onResumed"`) {
		t.Fatalf("expected JSON to contain 'onResumed' field, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"onInactive"`) {
		t.Fatalf("expected JSON to contain 'onInactive' field, got: %s", jsonStr)
	}
}

func TestLifecycleEventListenerAttributes_NilCallbacks_JSON(t *testing.T) {
	attrs := &duit_attributes.LifecycleEventListenerAttributes{
		OnStateChanged: nil,
		OnResumed:      nil,
		OnInactive:     nil,
		OnPaused:       nil,
		OnDetached:     nil,
		OnHidden:       nil,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"onStateChanged"`) || strings.Contains(jsonStr, `"onResumed"`) || 
	   strings.Contains(jsonStr, `"onInactive"`) || strings.Contains(jsonStr, `"onPaused"`) || 
	   strings.Contains(jsonStr, `"onDetached"`) || strings.Contains(jsonStr, `"onHidden"`) {
		t.Fatalf("expected JSON to not contain any callback fields when nil, got: %s", jsonStr)
	}
}
