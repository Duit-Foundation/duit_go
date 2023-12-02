package duit_border

type Border string
type BorderStyle string

const (
	Outline   Border      = "outline"
	Underline Border      = "underline"
	Solid     BorderStyle = "solid"
	None      BorderStyle = "none"
)

type BorderSide struct {
	Color string      `json:"color,omitempty"`
	Width float32     `json:"width,omitempty"`
	Style BorderStyle `json:"style,omitempty"`
}

type InputBorder struct {
	Type         Border     `json:"type"`
	Options      BorderSide `json:"options,omitempty"`
	GapPadding   float32    `json:"gapPadding,omitempty"`
	BorderRadius float32    `json:"borderRadius,omitempty"`
}
