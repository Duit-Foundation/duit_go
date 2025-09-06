package duit_props

type BorderSide struct {
	Color *Color      `json:"color,omitempty"`
	Width float32     `json:"width,omitempty"`
	Style BorderStyle `json:"style,omitempty"`
}
