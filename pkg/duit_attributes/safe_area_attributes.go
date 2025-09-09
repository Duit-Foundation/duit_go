package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SafeAreaAttributes defines attributes for SafeArea widget.
// See: https://api.flutter.dev/flutter/widgets/SafeArea-class.html
type SafeAreaAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Left                              duit_utils.Tristate[bool] `json:"left,omitempty"`
	Top                               duit_utils.Tristate[bool] `json:"top,omitempty"`
	Right                             duit_utils.Tristate[bool] `json:"right,omitempty"`
	Bottom                            duit_utils.Tristate[bool] `json:"bottom,omitempty"`
	Minimum                           *duit_props.EdgeInsets    `json:"minimum,omitempty"`
	MaintainBottomViewPadding         duit_utils.Tristate[bool] `json:"maintainBottomViewPadding,omitempty"`
}

//bindgen:exclude
func (r *SafeAreaAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}

// NewSafeAreaAttributes creates a new SafeAreaAttributes instance.
func NewSafeAreaAttributes() *SafeAreaAttributes {
	return &SafeAreaAttributes{}
}

// SetLeft sets the left for the safe area widget.
func (r *SafeAreaAttributes) SetLeft(left duit_utils.Tristate[bool]) *SafeAreaAttributes {
	r.Left = left
	return r
}

// SetTop sets the top for the safe area widget.
func (r *SafeAreaAttributes) SetTop(top duit_utils.Tristate[bool]) *SafeAreaAttributes {
	r.Top = top
	return r
}

// SetRight sets the right for the safe area widget.
func (r *SafeAreaAttributes) SetRight(right duit_utils.Tristate[bool]) *SafeAreaAttributes {
	r.Right = right
	return r
}

// SetBottom sets the bottom for the safe area widget.
func (r *SafeAreaAttributes) SetBottom(bottom duit_utils.Tristate[bool]) *SafeAreaAttributes {
	r.Bottom = bottom
	return r
}

// SetMinimum sets the minimum for the safe area widget.
func (r *SafeAreaAttributes) SetMinimum(minimum *duit_props.EdgeInsets) *SafeAreaAttributes {
	r.Minimum = minimum
	return r
}

// SetMaintainBottomViewPadding sets the maintain bottom view padding for the safe area widget.
func (r *SafeAreaAttributes) SetMaintainBottomViewPadding(maintainBottomViewPadding duit_utils.Tristate[bool]) *SafeAreaAttributes {
	r.MaintainBottomViewPadding = maintainBottomViewPadding
	return r
}
