package duit_props

// WidgetStateProperty represents values that depend on a widget's "state".
// The state is encoded as a set of WidgetState values, like WidgetState.disabled,
// WidgetState.hovered, WidgetState.focused. This class is useful for properties
// that need to change based on the widget's current state.
//
// See: https://api.flutter.dev/flutter/material/WidgetStateProperty-class.html
type WidgetStateProperty[T any] struct {
	Disabled T `json:"disabled,omitempty"`
	Hovered  T `json:"hovered,omitempty"`
	Focused  T `json:"focused,omitempty"`
	Pressed  T `json:"pressed,omitempty"`
	Dragged  T `json:"dragged,omitempty"`
	Selected T `json:"selected,omitempty"`
	Error    T `json:"error,omitempty"`
}

// Predefined types

type WidgetStateColor = WidgetStateProperty[*Color]
type WidgetStateSize = WidgetStateProperty[Size]
type WidgetStateEdgeInsets = WidgetStateProperty[*EdgeInsets]
type WidgetStateFloat32 = WidgetStateProperty[float32]
type WidgetStateBorderSide = WidgetStateProperty[*BorderSide]
type WidgetStateShapeBorder = WidgetStateProperty[*ShapeBorder]
type WidgetStateTextStyle = WidgetStateProperty[TextStyle]
type WidgetStateVisualDensity = WidgetStateProperty[VisualDensity]

//Constructors

// NewWidgetStateProperty creates a new instance of WidgetStateProperty
func NewWidgetStateProperty[T any]() *WidgetStateProperty[T] {
	return &WidgetStateProperty[T]{}
}

// NewWidgetStateColor creates a new instance of WidgetStateColor
func NewWidgetStateColor() *WidgetStateColor {
	return &WidgetStateColor{}
}

// NewWidgetStateSize creates a new instance of WidgetStateSize
func NewWidgetStateSize() *WidgetStateSize {
	return &WidgetStateSize{}
}

// NewWidgetStateEdgeInsets creates a new instance of WidgetStateEdgeInsets
func NewWidgetStateEdgeInsets() *WidgetStateEdgeInsets {
	return &WidgetStateEdgeInsets{}
}

// NewWidgetStateFloat32 creates a new instance of WidgetStateFloat32
func NewWidgetStateFloat32() *WidgetStateFloat32 {
	return &WidgetStateFloat32{}
}

// NewWidgetStateBorderSide creates a new instance of WidgetStateBorderSide
func NewWidgetStateBorderSide() *WidgetStateBorderSide {
	return &WidgetStateBorderSide{}
}

// NewWidgetStateShapeBorder creates a new instance of WidgetStateShapeBorder
func NewWidgetStateShapeBorder() *WidgetStateShapeBorder {
	return &WidgetStateShapeBorder{}
}

// NewWidgetStateTextStyle creates a new instance of WidgetStateTextStyle
func NewWidgetStateTextStyle() *WidgetStateTextStyle {
	return &WidgetStateTextStyle{}
}

// NewWidgetStateVisualDensity creates a new instance of WidgetStateVisualDensity
func NewWidgetStateVisualDensity() *WidgetStateVisualDensity {
	return &WidgetStateVisualDensity{}
}

// Setters

// SetDisabled sets the value for the disabled state
func (w *WidgetStateProperty[T]) SetDisabled(value T) *WidgetStateProperty[T] {
	w.Disabled = value
	return w
}

// SetHovered sets the value for the hovered state
func (w *WidgetStateProperty[T]) SetHovered(value T) *WidgetStateProperty[T] {
	w.Hovered = value
	return w
}

// SetFocused sets the value for the focused state
func (w *WidgetStateProperty[T]) SetFocused(value T) *WidgetStateProperty[T] {
	w.Focused = value
	return w
}

// SetPressed sets the value for the pressed state
func (w *WidgetStateProperty[T]) SetPressed(value T) *WidgetStateProperty[T] {
	w.Pressed = value
	return w
}

// SetDragged sets the value for the dragged state
func (w *WidgetStateProperty[T]) SetDragged(value T) *WidgetStateProperty[T] {
	w.Dragged = value
	return w
}

// SetSelected sets the value for the selected state
func (w *WidgetStateProperty[T]) SetSelected(value T) *WidgetStateProperty[T] {
	w.Selected = value
	return w
}

// SetError sets the value for the error state
func (w *WidgetStateProperty[T]) SetError(value T) *WidgetStateProperty[T] {
	w.Error = value
	return w
}
