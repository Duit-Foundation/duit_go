package duit_action

import (
	"errors"
)

func CheckActionType(action any) error {
	switch action.(type) {
	case *RemoteAction, *LocalAction, *ScriptAction, nil:
		return nil
	default:
		return errors.New("invalid action type. Must be instance of duit_core.Action or nil")	
	}
}