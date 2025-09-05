package duit_props

// InputBorderOptions represents options for input border
type InputBorderOptions struct {
	GapPadding   float32     `json:"gapPadding,omitempty"`
	BorderRadius float32     `json:"borderRadius,omitempty"`
	BorderSide   *BorderSide `json:"borderSide,omitempty"`
}

// InputBorder represents an input border
type InputBorder struct {
	Type    Border              `json:"type"`
	Options *InputBorderOptions `json:"options,omitempty"`
}
