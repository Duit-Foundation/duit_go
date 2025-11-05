package duit_action

const (
	Transport = iota
	Local
	Script
)

type ExecutionModifier string

const (
	ExecutionModifierThrottle ExecutionModifier = "throttle"
	ExecutionModifierDebounce ExecutionModifier = "debounce"
)

// HttpActionMetainfo represents metadata for an HTTP action.
type HttpActionMetainfo struct {
	Method string `json:"method"`
}

type ActionDependency struct {
	Id     string `json:"id"`
	Target string `json:"target"`
}

type ScriptData struct {
	SourceCode   string      `json:"sourceCode"`
	FunctionName string      `json:"functionName"`
	Meta         interface{} `json:"meta,omitempty"`
}

type ExecutionOptions struct {
	Modifier ExecutionModifier `json:"modifier"`
	Duration int               `json:"duration"`
}

type RemoteAction struct {
	ExecutionType    uint8               `json:"executionType"`
	Event            string              `json:"event"`
	DependsOn        []*ActionDependency `json:"dependsOn"`
	Meta             any                 `json:"meta,omitempty"`
	ExecutionOptions *ExecutionOptions   `json:"options,omitempty"`
}

// NewWebSocketAction creates a WebSocket action.
//
// It takes in an event string and dependsOn slice of strings as parameters.
// It returns a pointer to an Action struct.
func NewWebSocketAction(event string, dependsOn []*ActionDependency) *RemoteAction {
	return &RemoteAction{
		ExecutionType: Transport,
		Event:         event,
		DependsOn:     dependsOn,
	}
}

// NewHttpAction creates a new HTTP action.
//
// The function takes in the following parameters:
// - event: a string representing the event of the action.
// - dependsOn: a slice of strings representing the dependencies of the action.
// - metainfo: a pointer to an instance of HttpActionMetainfo representing the meta information of the action.
//
// The function returns a pointer to an instance of Action.
func NewHttpAction(event string, dependsOn []*ActionDependency, metainfo *HttpActionMetainfo) *RemoteAction {
	return &RemoteAction{
		ExecutionType: Transport,
		Event:         event,
		DependsOn:     dependsOn,
		Meta:          metainfo,
	}
}

type LocalAction struct {
	ExecutionType    uint8             `json:"executionType"`
	Payload          any               `json:"payload,omitempty"`
	ExecutionOptions *ExecutionOptions `json:"options,omitempty"`
}

// NewLocalExecutedtAction creates a new local action.
//
// The function takes in the following parameter:
// - payload: any type of data that will be passed to the action.
//
// The function returns a pointer to an instance of Action.
func NewLocalExecutedtAction(payload any) *LocalAction {
	return &LocalAction{
		ExecutionType: Local,
		Payload:       payload,
	}
}

type ScriptAction struct {
	ExecutionType    uint8               `json:"executionType"`
	Script           *ScriptData         `json:"script,omitempty"`
	Event            string              `json:"event"`
	DependsOn        []*ActionDependency `json:"dependsOn"`
	ExecutionOptions *ExecutionOptions   `json:"options,omitempty"`
}

// NewScriptAction creates a new script-based action.
//
// This function takes in the following parameters:
// - event: a string representing the event associated with the action.
// - dependsOn: a slice of ActionDependency pointers that the action depends on.
// - script: a pointer to ScriptData containing the script details for the action.
//
// It returns a pointer to an instance of ScriptAction.
func NewScriptAction(event string, dependsOn []*ActionDependency, script *ScriptData) *ScriptAction {
	return &ScriptAction{
		ExecutionType: Script,
		Event:         event,
		DependsOn:     dependsOn,
		Script:        script,
	}
}
