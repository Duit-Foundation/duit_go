package duit_attributes

type StackAttributes struct {
	TextDirection TextDirection `json:"textDirection,omitempty"`
	ClipBehavior  Clip          `json:"clipBehavior,omitempty"`
	Alignment     Alignment     `json:"alignment,omitempty"`
	Fit           StackFit      `json:"fit,omitempty"`
}
