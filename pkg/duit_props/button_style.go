package duit_props

// ButtonStyle defines the visual properties of buttons in Flutter.
// It allows customization of properties such as background color, text style,
// button shape, sizes, and other visual characteristics for different widget states.
//
// Flutter documentation: https://api.flutter.dev/flutter/material/ButtonStyle-class.html
type ButtonStyle struct {
	TextStyle         *WidgetStateTextStyle     `json:"textStyle,omitempty"`
	BackgroundColor   *WidgetStateColor         `json:"backgroundColor,omitempty"`
	ForegroundColor   *WidgetStateColor         `json:"foregroundColor,omitempty"`
	OverlayColor      *WidgetStateColor         `json:"overlayColor,omitempty"`
	ShadowColor       *WidgetStateColor         `json:"shadowColor,omitempty"`
	IconColor         *WidgetStateColor         `json:"iconColor,omitempty"`
	Elevation         *WidgetStateFloat32       `json:"elevation,omitempty"`
	IconSize          *WidgetStateFloat32       `json:"iconSize,omitempty"`
	Padding           *WidgetStateEdgeInsets    `json:"padding,omitempty"`
	MinimumSize       *WidgetStateSize          `json:"minimumSize,omitempty"`
	MaximumSize       *WidgetStateSize          `json:"maximumSize,omitempty"`
	Side              *WidgetStateBorderSide    `json:"side,omitempty"`
	Shape             *WidgetStateShapeBorder   `json:"shape,omitempty"`
	VisualDensity     *WidgetStateVisualDensity `json:"visualDensity,omitempty"`
	TapTargetSize     MaterialTapTargetSize     `json:"materialTapTargetSize,omitempty"`
	AnimationDuration uint                      `json:"animationDuration,omitempty"`
}

// NewButtonStyle creates a new empty ButtonStyle structure with default values
func NewButtonStyle() *ButtonStyle {
	return &ButtonStyle{}
}

// SetTextStyle sets the text style for the button in different widget states
func (r *ButtonStyle) SetTextStyle(textStyle *WidgetStateTextStyle) *ButtonStyle {
	r.TextStyle = textStyle
	return r
}

// SetBackgroundColor sets the background color for the button in different widget states
func (r *ButtonStyle) SetBackgroundColor(backgroundColor *WidgetStateColor) *ButtonStyle {
	r.BackgroundColor = backgroundColor
	return r
}

// SetForegroundColor sets the foreground color for the button in different widget states
func (r *ButtonStyle) SetForegroundColor(foregroundColor *WidgetStateColor) *ButtonStyle {
	r.ForegroundColor = foregroundColor
	return r
}

// SetOverlayColor sets the overlay color for the button in different widget states
func (r *ButtonStyle) SetOverlayColor(overlayColor *WidgetStateColor) *ButtonStyle {
	r.OverlayColor = overlayColor
	return r
}

// SetShadowColor sets the shadow color for the button in different widget states
func (r *ButtonStyle) SetShadowColor(shadowColor *WidgetStateColor) *ButtonStyle {
	r.ShadowColor = shadowColor
	return r
}

// SetIconColor sets the icon color for the button in different widget states
func (r *ButtonStyle) SetIconColor(iconColor *WidgetStateColor) *ButtonStyle {
	r.IconColor = iconColor
	return r
}

// SetElevation sets the elevation for the button in different widget states
func (r *ButtonStyle) SetElevation(elevation *WidgetStateFloat32) *ButtonStyle {
	r.Elevation = elevation
	return r
}

// SetIconSize sets the icon size for the button in different widget states
func (r *ButtonStyle) SetIconSize(iconSize *WidgetStateFloat32) *ButtonStyle {
	r.IconSize = iconSize
	return r
}

// SetPadding sets the padding for the button in different widget states
func (r *ButtonStyle) SetPadding(padding *WidgetStateEdgeInsets) *ButtonStyle {
	r.Padding = padding
	return r
}

// SetMinimumSize sets the minimum size for the button in different widget states
func (r *ButtonStyle) SetMinimumSize(minimumSize *WidgetStateSize) *ButtonStyle {
	r.MinimumSize = minimumSize
	return r
}

// SetMaximumSize sets the maximum size for the button in different widget states
func (r *ButtonStyle) SetMaximumSize(maximumSize *WidgetStateSize) *ButtonStyle {
	r.MaximumSize = maximumSize
	return r
}

// SetSide sets the border side for the button in different widget states
func (r *ButtonStyle) SetSide(side *WidgetStateBorderSide) *ButtonStyle {
	r.Side = side
	return r
}

// SetShape sets the shape for the button in different widget states
func (r *ButtonStyle) SetShape(shape *WidgetStateShapeBorder) *ButtonStyle {
	r.Shape = shape
	return r
}

// SetVisualDensity sets the visual density for the button in different widget states
func (r *ButtonStyle) SetVisualDensity(visualDensity *WidgetStateVisualDensity) *ButtonStyle {
	r.VisualDensity = visualDensity
	return r
}

// SetTapTargetSize sets the tap target size for the button
func (r *ButtonStyle) SetTapTargetSize(tapTargetSize MaterialTapTargetSize) *ButtonStyle {
	r.TapTargetSize = tapTargetSize
	return r
}

// SetAnimationDuration sets the animation duration for the button
func (r *ButtonStyle) SetAnimationDuration(animationDuration uint) *ButtonStyle {
	r.AnimationDuration = animationDuration
	return r
}
