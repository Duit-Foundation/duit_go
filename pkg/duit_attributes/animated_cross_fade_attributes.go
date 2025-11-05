package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// AnimatedCrossFadeAttributes represents attributes for AnimatedCrossFade widget.
// https://api.flutter.dev/flutter/widgets/AnimatedCrossFade-class.html
type AnimatedCrossFadeAttributes struct {
	*ValueReferenceHolder          `gen:"wrap"`
	*duit_props.ImplicitAnimatable `gen:"wrap"`
	*ThemeConsumer                 `gen:"wrap"`
	ReverseDuration                uint                      `json:"reverseDuration,omitempty"`
	FirstCurve                     duit_props.Curves         `json:"firstCurve,omitempty"`
	SecondCurve                    duit_props.Curves         `json:"secondCurve,omitempty"`
	SizeCurve                      duit_props.Curves         `json:"sizeCurve,omitempty"`
	ExcludeBottomFocus             duit_utils.Tristate[bool] `json:"excludeBottomFocus,omitempty"`
	Alignment                      duit_props.Alignment      `json:"alignment,omitempty"`
	// Where 0 - CrossFadeState.showFirst and 1 - CrossFadeState.showSecond
	CrossFadeState uint8 `json:"crossFadeState"`
}

// NewAnimatedCrossFadeAttributes creates a new instance of AnimatedCrossFadeAttributes.
func NewAnimatedCrossFadeAttributes() *AnimatedCrossFadeAttributes {
	return &AnimatedCrossFadeAttributes{}
}

// SetReverseDuration sets the reverseDuration property and returns the AnimatedCrossFadeAttributes instance for method chaining.
func (r *AnimatedCrossFadeAttributes) SetReverseDuration(reverseDuration uint) *AnimatedCrossFadeAttributes {
	r.ReverseDuration = reverseDuration
	return r
}

// SetFirstCurve sets the firstCurve property and returns the AnimatedCrossFadeAttributes instance for method chaining.
func (r *AnimatedCrossFadeAttributes) SetFirstCurve(firstCurve duit_props.Curves) *AnimatedCrossFadeAttributes {
	r.FirstCurve = firstCurve
	return r
}

// SetSecondCurve sets the secondCurve property and returns the AnimatedCrossFadeAttributes instance for method chaining.
func (r *AnimatedCrossFadeAttributes) SetSecondCurve(secondCurve duit_props.Curves) *AnimatedCrossFadeAttributes {
	r.SecondCurve = secondCurve
	return r
}

// SetSizeCurve sets the sizeCurve property and returns the AnimatedCrossFadeAttributes instance for method chaining.
func (r *AnimatedCrossFadeAttributes) SetSizeCurve(sizeCurve duit_props.Curves) *AnimatedCrossFadeAttributes {
	r.SizeCurve = sizeCurve
	return r
}

// SetExcludeBottomFocus sets the excludeBottomFocus property and returns the AnimatedCrossFadeAttributes instance for method chaining.
func (r *AnimatedCrossFadeAttributes) SetExcludeBottomFocus(excludeBottomFocus bool) *AnimatedCrossFadeAttributes {
	r.ExcludeBottomFocus = duit_utils.BoolValue(excludeBottomFocus)
	return r
}

// SetAlignment sets the alignment property and returns the AnimatedCrossFadeAttributes instance for method chaining.
func (r *AnimatedCrossFadeAttributes) SetAlignment(alignment duit_props.Alignment) *AnimatedCrossFadeAttributes {
	r.Alignment = alignment
	return r
}

// SetCrossFadeState sets the crossFadeState property and returns the AnimatedCrossFadeAttributes instance for method chaining.
func (r *AnimatedCrossFadeAttributes) SetCrossFadeState(crossFadeState uint8) *AnimatedCrossFadeAttributes {
	r.CrossFadeState = crossFadeState
	return r
}

//bindgen:exclude
func (r *AnimatedCrossFadeAttributes) Validate() error {
	if r.ImplicitAnimatable != nil {
		if err := r.ImplicitAnimatable.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("implicitAnimatable property is required on implicit animated widgets")
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.CrossFadeState != 0 && r.CrossFadeState != 1 {
		return errors.New("crossFadeState must be 0 or 1")
	}

	return nil
}
