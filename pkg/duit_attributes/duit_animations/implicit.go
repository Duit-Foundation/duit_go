package duit_animations

type ImplicitAnimatable struct {
	Duration uint   `json:"duration"`
	Curve    Curves `json:"curve,omitempty"`
	OnEnd    any    `json:"onEnd,omitempty"`
}
