package duit_material

type MaterialStateProperty[T any] struct {
	Disabled T `json:"disabled,omitempty"`
	Hovered  T `json:"hovered,omitempty"`
	Focused  T `json:"focused,omitempty"`
	Pressed  T `json:"pressed,omitempty"`
	Dragged  T `json:"dragged,omitempty"`
	Selected T `json:"selected,omitempty"`
	Error    T `json:"error,omitempty"`
}
