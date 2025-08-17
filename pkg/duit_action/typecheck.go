package duit_action

import (
	"errors"
)

func CheckActionType(action any) error {
	switch action.(type) {
	case *RemoteAction:
	case *LocalAction:
	case *ScriptAction:
	case nil:
		return nil
	default:
		return errors.New("invalid action type. Must be instance of duit_core.Action or nil")	
	}

	return errors.New("invalid action type. Must be instance of duit_core.Action or nil")
}