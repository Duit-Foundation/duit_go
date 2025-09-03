package duit_props

// InputBorderOptions represents options for input border
type InputBorderOptions[T Color] struct {
	GapPadding   float32        `json:"gapPadding,omitempty"`
	BorderRadius float32        `json:"borderRadius,omitempty"`
	BorderSide   *BorderSide[T] `json:"borderSide,omitempty"`
}

// InputBorder represents an input border
type InputBorder[T Color] struct {
	Type    Border                 `json:"type"`
	Options *InputBorderOptions[T] `json:"options,omitempty"`
}
