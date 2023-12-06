package duit_core

// Event is a string field that represents the event associated with the action.
//
// DependsOn is a string slice field that represents the dependencies of the action.
//
// Meta is an interface{} field that represents optional additional metadata for the action.
type Action struct {
	Event     string      `json:"event"`
	DependsOn []string    `json:"dependsOn"`
	Meta      interface{} `json:"meta,omitempty"`
}
