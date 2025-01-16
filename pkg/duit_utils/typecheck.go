package duit_utils

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func CheckActionType(action any) error {
	switch action.(type) {
	case *duit_core.RemoteAction:
	case *duit_core.LocalAction:
	case *duit_core.ScriptAction:
	case nil:
		return nil
	default:
		return errors.New("invalid action type. Must be instance of duit_core.Action or nil")	
	}

	return errors.New("invalid action type. Must be instance of duit_core.Action or nil")
}