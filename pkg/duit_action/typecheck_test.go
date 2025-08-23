package duit_action_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
)

func TestCheckActionType_BasicCases(t *testing.T) {
	remoteAction := &duit_action.RemoteAction{
		ExecutionType: duit_action.Transport,
		Event:         "tap_event",
	}

	err := duit_action.CheckActionType(remoteAction)
	if err != nil {
		t.Fatal("expected no error, got:", err)
	}

	err = duit_action.CheckActionType(nil)
	if err != nil {
		t.Fatal("expected no error, got:", err)
	}

	err = duit_action.CheckActionType(1)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}

func TestCheckActionType_ValidTypes(t *testing.T) {
	validActions := []any{
		&duit_action.RemoteAction{
			ExecutionType: duit_action.Transport,
			Event:         "test_event",
		},
		&duit_action.LocalAction{
			ExecutionType: duit_action.Local,
			Payload:       "test_payload",
		},
		&duit_action.ScriptAction{
			ExecutionType: duit_action.Script,
			Script: &duit_action.ScriptData{
				SourceCode:   "console.log('test')",
				FunctionName: "testFunction",
			},
			Event: "script_event",
		},
		nil,
	}

	for i, action := range validActions {
		err := duit_action.CheckActionType(action)
		if err != nil {
			t.Fatalf("expected no error for valid action at index %d (%T), got: %s", i, action, err.Error())
		}
	}
}

func TestCheckActionType_InvalidTypes(t *testing.T) {
	invalidActions := []any{
		"invalid string",
		123,
		45.67,
		true,
		[]string{"test"},
		map[string]int{"key": 1},
		struct{ field string }{field: "test"},
		func() {},
		&struct{ field int }{field: 42},
	}

	for i, invalidAction := range invalidActions {
		err := duit_action.CheckActionType(invalidAction)
		if err == nil {
			t.Fatalf("expected error for invalid action at index %d (%T), but got nil", i, invalidAction)
		}
		
		expectedErrorMsg := "invalid action type. Must be instance of duit_core.Action or nil"
		if err.Error() != expectedErrorMsg {
			t.Fatalf("expected error message '%s' for invalid action at index %d, got: %s", expectedErrorMsg, i, err.Error())
		}
	}
}