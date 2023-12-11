package duit_core

type ActionDependency struct {
	Id     string `json:"id"`
	Target string `json:"target"`
}

// Event is a string field that represents the event associated with the action.
//
// DependsOn is a string slice field that represents the dependencies of the action.
//
// Meta is an interface{} field that represents optional additional metadata for the action.
type Action struct {
	Event     string              `json:"event"`
	DependsOn []*ActionDependency `json:"dependsOn"`
	Meta      interface{}         `json:"meta,omitempty"`
}
