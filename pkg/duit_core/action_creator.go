package duit_core

// HttpActionMetainfo represents metadata for an HTTP action.
type HttpActionMetainfo struct {
	Method string `json:"method"`
}

// HttpAction creates a new HTTP action.
//
// The function takes in the following parameters:
// - event: a string representing the event of the action.
// - dependsOn: a slice of strings representing the dependencies of the action.
// - metainfo: a pointer to an instance of HttpActionMetainfo representing the meta information of the action.
//
// The function returns a pointer to an instance of Action.
func HttpAction(event string, dependsOn []*ActionDependency, metainfo *HttpActionMetainfo) *RemoteAction {
	return &RemoteAction{
		ExecutionType: Transport,
		Event:         event,
		DependsOn:     dependsOn,
		Meta:          *metainfo,
	}
}

// WebSocketAction creates a WebSocket action.
//
// It takes in an event string and dependsOn slice of strings as parameters.
// It returns a pointer to an Action struct.
func WebSocketAction(event string, dependsOn []*ActionDependency) *RemoteAction {
	return &RemoteAction{
		ExecutionType: Transport,
		Event:         event,
		DependsOn:     dependsOn,
	}
}

func LocalExecutedtAction(payload any) *LocalAction {
	return &LocalAction{
		ExecutionType: Local,
		Payload:       payload,
	}
}

func ScriptAction(event string, dependsOn []*ActionDependency, script *ScriptData) *ScriptActionS {
	return &ScriptActionS{
		ExecutionType: Script,
		Event:         event,
		DependsOn:     dependsOn,
		Script:        script,
	}
}
