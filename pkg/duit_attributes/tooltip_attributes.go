package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// TooltipAttributes defines attributes for Tooltip widget.
// See: https://api.flutter.dev/flutter/material/Tooltip-class.html
type TooltipAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	Message               string                        `json:"message,omitempty"`
	RichMessage           *duit_props.TextSpan          `json:"richMessage,omitempty"`
	Padding               *duit_props.EdgeInsets        `json:"padding,omitempty"`
	Margin                *duit_props.EdgeInsets        `json:"margin,omitempty"`
	VerticalOffset        float32                       `json:"verticalOffset,omitempty"`
	Height                float32                       `json:"height,omitempty"`
	PreferBelow           duit_utils.Tristate[bool]     `json:"preferBelow,omitempty"`
	ExcludeFromSemantics  duit_utils.Tristate[bool]     `json:"excludeFromSemantics,omitempty"`
	Decoration            *duit_props.BoxDecoration     `json:"decoration,omitempty"`
	TextStyle             *duit_props.TextStyle         `json:"textStyle,omitempty"`
	TextAlign             duit_props.TextAlign          `json:"textAlign,omitempty"`
	WaitDuration          uint                          `json:"waitDuration,omitempty"`
	ShowDuration          uint                          `json:"showDuration,omitempty"`
	ExitDuration          uint                          `json:"exitDuration,omitempty"`
	EnableFeedback        duit_utils.Tristate[bool]     `json:"enableFeedback,omitempty"`
	TriggerMode           duit_props.TooltipTriggerMode `json:"triggerMode,omitempty"`
	EnableTapToDismiss    duit_utils.Tristate[bool]     `json:"enableTapToDismiss,omitempty"`
	// Constraints           *duit_props.BoxConstraints    `json:"constraints,omitempty"`
	// IgnorePointer         duit_utils.Tristate[bool] `json:"ignorePointer,omitempty"`
}

// NewTooltipAttributes creates a new TooltipAttributes instance.
func NewTooltipAttributes() *TooltipAttributes {
	return &TooltipAttributes{}
}

// SetMessage sets the message property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetMessage(message string) *TooltipAttributes {
	r.Message = message
	return r
}

// SetPadding sets the padding property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetPadding(padding *duit_props.EdgeInsets) *TooltipAttributes {
	r.Padding = padding
	return r
}

// SetMargin sets the margin property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetMargin(margin *duit_props.EdgeInsets) *TooltipAttributes {
	r.Margin = margin
	return r
}

// SetVerticalOffset sets the verticalOffset property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetVerticalOffset(verticalOffset float32) *TooltipAttributes {
	r.VerticalOffset = verticalOffset
	return r
}

// SetPreferBelow sets the preferBelow property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetPreferBelow(preferBelow bool) *TooltipAttributes {
	r.PreferBelow = duit_utils.BoolValue(preferBelow)
	return r
}

// SetExcludeFromSemantics sets the excludeFromSemantics property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetExcludeFromSemantics(excludeFromSemantics bool) *TooltipAttributes {
	r.ExcludeFromSemantics = duit_utils.BoolValue(excludeFromSemantics)
	return r
}

// SetDecoration sets the decoration property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetDecoration(decoration *duit_props.BoxDecoration) *TooltipAttributes {
	r.Decoration = decoration
	return r
}

// SetTextStyle sets the textStyle property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetTextStyle(textStyle *duit_props.TextStyle) *TooltipAttributes {
	r.TextStyle = textStyle
	return r
}

// SetTextAlign sets the textAlign property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetTextAlign(textAlign duit_props.TextAlign) *TooltipAttributes {
	r.TextAlign = textAlign
	return r
}

// SetWaitDuration sets the waitDuration property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetWaitDuration(waitDuration uint) *TooltipAttributes {
	r.WaitDuration = waitDuration
	return r
}

// SetShowDuration sets the showDuration property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetShowDuration(showDuration uint) *TooltipAttributes {
	r.ShowDuration = showDuration
	return r
}

// SetExitDuration sets the exitDuration property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetExitDuration(exitDuration uint) *TooltipAttributes {
	r.ExitDuration = exitDuration
	return r
}

// SetEnableFeedback sets the enableFeedback property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetEnableFeedback(enableFeedback bool) *TooltipAttributes {
	r.EnableFeedback = duit_utils.BoolValue(enableFeedback)
	return r
}

// SetEnableTapToDismiss sets the enableTapToDismiss property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetEnableTapToDismiss(enableTapToDismiss bool) *TooltipAttributes {
	r.EnableTapToDismiss = duit_utils.BoolValue(enableTapToDismiss)
	return r
}

// SetTriggerMode sets the triggerMode property and returns TooltipAttributes for method chaining.
func (r *TooltipAttributes) SetTriggerMode(triggerMode duit_props.TooltipTriggerMode) *TooltipAttributes {
	r.TriggerMode = triggerMode
	return r
}

// SetIgnorePointer sets the ignorePointer property and returns TooltipAttributes for method chaining.
// func (r *TooltipAttributes) SetIgnorePointer(ignorePointer bool) *TooltipAttributes {
// 	r.IgnorePointer = duit_utils.BoolValue(ignorePointer)
// 	return r
// }

// SetConstraints sets the constraints property and returns TooltipAttributes for method chaining.
// func (r *TooltipAttributes) SetConstraints(constraints *duit_props.BoxConstraints) *TooltipAttributes {
// 	r.Constraints = constraints
// 	return r
// }

//bindgen:exclude
func (r *TooltipAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.Message == "" {
		return errors.New("message property is required")
	}

	// if r.Constraints != nil {
	// 	if err := r.Constraints.Validate(); err != nil {
	// 		return err
	// 	}
	// }

	if r.Padding != nil {
		if err := r.Padding.Validate(); err != nil {
			return err
		}
	}

	if r.Margin != nil {
		if err := r.Margin.Validate(); err != nil {
			return err
		}
	}

	if r.Decoration != nil {
		if err := r.Decoration.Validate(); err != nil {
			return err
		}
	}

	if r.TextStyle != nil {
		if err := r.TextStyle.Validate(); err != nil {
			return err
		}
	}

	return nil
}
