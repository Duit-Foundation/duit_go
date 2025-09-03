package duit_props



type ButtonStyle[TColor Color, TInsets EdgeInsets, TShape ShapeBorder[TColor]] struct {
	TextStyle         MaterialStateProperty[TextStyle[TColor]]       `json:"textStyle,omitempty"`
	BackgroundColor   MaterialStateProperty[TColor]                  `json:"backgroundColor,omitempty"`
	ForegroundColor   MaterialStateProperty[TColor]                  `json:"foregroundColor,omitempty"`
	OverlayColor      MaterialStateProperty[TColor]                  `json:"overlayColor,omitempty"`
	ShadowColor       MaterialStateProperty[TColor]                  `json:"shadowColor,omitempty"`
	IconColor         MaterialStateProperty[TColor]                  `json:"iconColor,omitempty"`
	Elevation         MaterialStateProperty[float32]                 `json:"elevation,omitempty"`
	IconSize          MaterialStateProperty[float32]                 `json:"iconSize,omitempty"`
	Padding           MaterialStateProperty[TInsets]                 `json:"padding,omitempty"`
	MinimumSize       MaterialStateProperty[Size]          `json:"minimumSize,omitempty"`
	MaximumSize       MaterialStateProperty[Size]          `json:"maximumSize,omitempty"`
	Side              MaterialStateProperty[BorderSide[TColor]]      `json:"side,omitempty"`
	Shape             MaterialStateProperty[TShape]                  `json:"shape,omitempty"`
	VisualDensity     MaterialStateProperty[VisualDensity] `json:"visualDensity,omitempty"`
	TapTargetSize     MaterialTapTargetSize                          `json:"materialTapTargetSize,omitempty"`
	AnimationDuration uint                                           `json:"animationDuration,omitempty"`
}
