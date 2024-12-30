package duit_core

const (
	Transport = iota
	Local
	Script
)

type ActionDependency struct {
	Id     string `json:"id"`
	Target string `json:"target"`
}

type ScriptData struct {
	SourceCode   string      `json:"sourceCode"`
	FunctionName string      `json:"functionName"`
	Meta         interface{} `json:"meta,omitempty"`
}

// Event is a string field that represents the event associated with the action.
//
// DependsOn is a string slice field that represents the dependencies of the action.
//
// Meta is an interface{} field that represents optional additional metadata for the action.
type Action interface {
	RemoteAction | LocalAction | ScriptAction
}

type RemoteAction struct {
	ExecutionType uint8               `json:"executionType"`
	Event         string              `json:"event"`
	DependsOn     []*ActionDependency `json:"dependsOn"`
	Meta          any                 `json:"meta,omitempty"`
}

type LocalAction struct {
	ExecutionType uint8 `json:"executionType"`
	Payload       any   `json:"payload,omitempty"`
}

type ScriptAction struct {
	ExecutionType uint8               `json:"executionType"`
	Script        *ScriptData         `json:"script,omitempty"`
	Event         string              `json:"event"`
	DependsOn     []*ActionDependency `json:"dependsOn"`
}
