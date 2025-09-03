package duit_props

type BorderSide[T Color] struct {
	Color T           `json:"color,omitempty"`
	Width float32     `json:"width,omitempty"`
	Style BorderStyle `json:"style,omitempty"`
}
