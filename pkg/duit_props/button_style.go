package duit_props

type ButtonStyle[TInsets EdgeInsets, TShape ShapeBorder] struct {
	TextStyle         *MaterialStateProperty[TextStyle]     `json:"textStyle,omitempty"`
	BackgroundColor   *MaterialStateProperty[Color]        `json:"backgroundColor,omitempty"`
	ForegroundColor   *MaterialStateProperty[Color]        `json:"foregroundColor,omitempty"`
	OverlayColor      *MaterialStateProperty[Color]        `json:"overlayColor,omitempty"`
	ShadowColor       *MaterialStateProperty[Color]        `json:"shadowColor,omitempty"`
	IconColor         *MaterialStateProperty[Color]        `json:"iconColor,omitempty"`
	Elevation         *MaterialStateProperty[float32]       `json:"elevation,omitempty"`
	IconSize          *MaterialStateProperty[float32]       `json:"iconSize,omitempty"`
	Padding           *MaterialStateProperty[TInsets]       `json:"padding,omitempty"`
	MinimumSize       *MaterialStateProperty[Size]          `json:"minimumSize,omitempty"`
	MaximumSize       *MaterialStateProperty[Size]          `json:"maximumSize,omitempty"`
	Side              *MaterialStateProperty[BorderSide]    `json:"side,omitempty"`
	Shape             *MaterialStateProperty[TShape]        `json:"shape,omitempty"`
	VisualDensity     *MaterialStateProperty[VisualDensity] `json:"visualDensity,omitempty"`
	TapTargetSize     MaterialTapTargetSize                 `json:"materialTapTargetSize,omitempty"`
	AnimationDuration uint                                  `json:"animationDuration,omitempty"`
}
