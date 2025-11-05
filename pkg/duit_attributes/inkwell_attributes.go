package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// InkwellAttributes represents attributes for InkWell widget.
// https://api.flutter.dev/flutter/material/InkWell-class.html
type InkwellAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	OnTap                any                          `json:"onTap,omitempty"`
	OnDoubleTap          any                          `json:"onDoubleTap,omitempty"`
	OnLongPress          any                          `json:"onLongPress,omitempty"`
	OnTapDown            any                          `json:"onTapDown,omitempty"`
	OnTapUp              any                          `json:"onTapUp,omitempty"`
	OnTapCancel          any                          `json:"onTapCancel,omitempty"`
	OnSecondaryTapDown   any                          `json:"onSecondaryTapDown,omitempty"`
	OnSecondaryTapCancel any                          `json:"onSecondaryTapCancel,omitempty"`
	OnSecondaryTap       any                          `json:"onSecondaryTap,omitempty"`
	OnSecondaryTapUp     any                          `json:"onSecondaryTapUp,omitempty"`
	FocusColor           *duit_props.Color            `json:"focusColor,omitempty"`
	HoverColor           *duit_props.Color            `json:"hoverColor,omitempty"`
	HighlightColor       *duit_props.Color            `json:"highlightColor,omitempty"`
	SplashColor          *duit_props.Color            `json:"splashColor,omitempty"`
	OverlayColor         *duit_props.WidgetStateColor `json:"overlayColor,omitempty"`
	Radius               float32                      `json:"radius,omitempty"`
	BorderRadius         *duit_props.BorderRadius     `json:"borderRadius,omitempty"`
	CustomBorder         *duit_props.ShapeBorder      `json:"customBorder,omitempty"`
	EnableFeedback       duit_utils.Tristate[bool]    `json:"enableFeedback,omitempty"`
	ExcludeFromSemantics duit_utils.Tristate[bool]    `json:"excludeFromSemantics,omitempty"`
	Autofocus            duit_utils.Tristate[bool]    `json:"autofocus,omitempty"`
	CanRequestFocus      duit_utils.Tristate[bool]    `json:"canRequestFocus,omitempty"`
	HoverDuration        uint                         `json:"hoverDuration,omitempty"`
}

// NewInkwellAttributes creates a new instance of InkwellAttributes.
func NewInkwellAttributes() *InkwellAttributes {
	return &InkwellAttributes{}
}

// SetOnTap sets the onTap property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetOnTap(onTap any) *InkwellAttributes {
	r.OnTap = onTap
	return r
}

// SetOnDoubleTap sets the onDoubleTap property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetOnDoubleTap(onDoubleTap any) *InkwellAttributes {
	r.OnDoubleTap = onDoubleTap
	return r
}

// SetOnLongPress sets the onLongPress property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetOnLongPress(onLongPress any) *InkwellAttributes {
	r.OnLongPress = onLongPress
	return r
}

// SetOnTapDown sets the onTapDown property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetOnTapDown(onTapDown any) *InkwellAttributes {
	r.OnTapDown = onTapDown
	return r
}

// SetOnTapUp sets the onTapUp property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetOnTapUp(onTapUp any) *InkwellAttributes {
	r.OnTapUp = onTapUp
	return r
}

// SetOnTapCancel sets the onTapCancel property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetOnTapCancel(onTapCancel any) *InkwellAttributes {
	r.OnTapCancel = onTapCancel
	return r
}

// SetOnSecondaryTapDown sets the onSecondaryTapDown property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetOnSecondaryTapDown(onSecondaryTapDown any) *InkwellAttributes {
	r.OnSecondaryTapDown = onSecondaryTapDown
	return r
}

// SetOnSecondaryTapCancel sets the onSecondaryTapCancel property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetOnSecondaryTapCancel(onSecondaryTapCancel any) *InkwellAttributes {
	r.OnSecondaryTapCancel = onSecondaryTapCancel
	return r
}

// SetOnSecondaryTap sets the onSecondaryTap property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetOnSecondaryTap(onSecondaryTap any) *InkwellAttributes {
	r.OnSecondaryTap = onSecondaryTap
	return r
}

// SetOnSecondaryTapUp sets the onSecondaryTapUp property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetOnSecondaryTapUp(onSecondaryTapUp any) *InkwellAttributes {
	r.OnSecondaryTapUp = onSecondaryTapUp
	return r
}

// SetFocusColor sets the focusColor property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetFocusColor(focusColor *duit_props.Color) *InkwellAttributes {
	r.FocusColor = focusColor
	return r
}

// SetHoverColor sets the hoverColor property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetHoverColor(hoverColor *duit_props.Color) *InkwellAttributes {
	r.HoverColor = hoverColor
	return r
}

// SetHighlightColor sets the highlightColor property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetHighlightColor(highlightColor *duit_props.Color) *InkwellAttributes {
	r.HighlightColor = highlightColor
	return r
}

// SetSplashColor sets the splashColor property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetSplashColor(splashColor *duit_props.Color) *InkwellAttributes {
	r.SplashColor = splashColor
	return r
}

// SetOverlayColor sets the overlayColor property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetOverlayColor(overlayColor *duit_props.WidgetStateColor) *InkwellAttributes {
	r.OverlayColor = overlayColor
	return r
}

// SetRadius sets the radius property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetRadius(radius float32) *InkwellAttributes {
	r.Radius = radius
	return r
}

// SetBorderRadius sets the borderRadius property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetBorderRadius(borderRadius *duit_props.BorderRadius) *InkwellAttributes {
	r.BorderRadius = borderRadius
	return r
}

// SetCustomBorder sets the customBorder property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetCustomBorder(customBorder *duit_props.ShapeBorder) *InkwellAttributes {
	r.CustomBorder = customBorder
	return r
}

// SetEnableFeedback sets the enableFeedback property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetEnableFeedback(enableFeedback bool) *InkwellAttributes {
	r.EnableFeedback = duit_utils.BoolValue(enableFeedback)
	return r
}

// SetExcludeFromSemantics sets the excludeFromSemantics property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetExcludeFromSemantics(excludeFromSemantics bool) *InkwellAttributes {
	r.ExcludeFromSemantics = duit_utils.BoolValue(excludeFromSemantics)
	return r
}

// SetAutofocus sets the autofocus property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetAutofocus(autofocus bool) *InkwellAttributes {
	r.Autofocus = duit_utils.BoolValue(autofocus)
	return r
}

// SetCanRequestFocus sets the canRequestFocus property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetCanRequestFocus(canRequestFocus bool) *InkwellAttributes {
	r.CanRequestFocus = duit_utils.BoolValue(canRequestFocus)
	return r
}

// SetHoverDuration sets the hoverDuration property and returns the InkwellAttributes instance for method chaining.
func (r *InkwellAttributes) SetHoverDuration(hoverDuration uint) *InkwellAttributes {
	r.HoverDuration = hoverDuration
	return r
}

//bindgen:exclude
func (r *InkwellAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	callbacks := []any{
		r.OnTap,
		r.OnDoubleTap,
		r.OnLongPress,
		r.OnTapDown,
		r.OnTapUp,
		r.OnTapCancel,
		r.OnSecondaryTapDown,
		r.OnSecondaryTapCancel,
		r.OnSecondaryTap,
		r.OnSecondaryTapUp,
	}

	for _, callback := range callbacks {
		if err := duit_action.CheckActionType(callback); err != nil {
			return err
		}
	}

	return nil
}
