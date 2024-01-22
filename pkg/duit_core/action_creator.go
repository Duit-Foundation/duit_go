package duit_core

// HttpActionMetainfo represents metadata for an HTTP action.
type HttpActionMetainfo struct {
	Method string `json:"method"`
}

// CreateHttpAction creates a new HTTP action.
//
// The function takes in the following parameters:
// - event: a string representing the event of the action.
// - dependsOn: a slice of strings representing the dependencies of the action.
// - metainfo: a pointer to an instance of HttpActionMetainfo representing the meta information of the action.
//
// The function returns a pointer to an instance of Action.
func CreateHttpAction(event string, dependsOn []*ActionDependency, metainfo *HttpActionMetainfo) *Action {
	return &Action{
		ExecutionType: Transport,
		Event:         event,
		DependsOn:     dependsOn,
		Meta:          *metainfo,
	}
}

// CreateWebSocketAction creates a WebSocket action.
//
// It takes in an event string and dependsOn slice of strings as parameters.
// It returns a pointer to an Action struct.
func CreateWebSocketAction(event string, dependsOn []*ActionDependency) *Action {
	return &Action{
		ExecutionType: Transport,
		Event:         event,
		DependsOn:     dependsOn,
	}
}

// CreateWebSocketAction creates a WebSocket action.
//
// It takes in an event string and dependsOn slice of strings as parameters.
// It returns a pointer to an Action struct.
func CreateLocalExecutedtAction(payload *Event) *Action {
	return &Action{
		ExecutionType: Local,
		Event:         "local_event",
		DependsOn:     []*ActionDependency{},
		Payload:       payload,
	}
}
