package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SliverSafeAreaAttributes defines attributes for SliverSafeArea widget.
// See: https://api.flutter.dev/flutter/widgets/SliverSafeArea-class.html
type SliverSafeAreaAttributes struct {
	*ValueReferenceHolder     `gen:"wrap"`
	*ThemeConsumer            `gen:"wrap"`
	*SliverProps              `gen:"wrap"`
	Left                      duit_utils.Tristate[bool] `json:"left,omitempty"`
	Top                       duit_utils.Tristate[bool] `json:"top,omitempty"`
	Right                     duit_utils.Tristate[bool] `json:"right,omitempty"`
	Bottom                    duit_utils.Tristate[bool] `json:"bottom,omitempty"`
	Minimum                   *duit_props.EdgeInsets    `json:"minimum,omitempty"`
	MaintainBottomViewPadding duit_utils.Tristate[bool] `json:"maintainBottomViewPadding,omitempty"`
}

//bindgen:exclude
func (r *SliverSafeAreaAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}

// NewSliverSafeAreaAttributes creates a new SliverSafeAreaAttributes instance.
func NewSliverSafeAreaAttributes() *SliverSafeAreaAttributes {
	return &SliverSafeAreaAttributes{}
}

// SetLeft sets the left for the sliver safe area widget.
func (r *SliverSafeAreaAttributes) SetLeft(left duit_utils.Tristate[bool]) *SliverSafeAreaAttributes {
	r.Left = left
	return r
}

// SetTop sets the top for the sliver safe area widget.
func (r *SliverSafeAreaAttributes) SetTop(top duit_utils.Tristate[bool]) *SliverSafeAreaAttributes {
	r.Top = top
	return r
}

// SetRight sets the right for the sliver safe area widget.
func (r *SliverSafeAreaAttributes) SetRight(right duit_utils.Tristate[bool]) *SliverSafeAreaAttributes {
	r.Right = right
	return r
}

// SetBottom sets the bottom for the sliver safe area widget.
func (r *SliverSafeAreaAttributes) SetBottom(bottom duit_utils.Tristate[bool]) *SliverSafeAreaAttributes {
	r.Bottom = bottom
	return r
}

// SetMinimum sets the minimum for the sliver safe area widget.
func (r *SliverSafeAreaAttributes) SetMinimum(minimum *duit_props.EdgeInsets) *SliverSafeAreaAttributes {
	r.Minimum = minimum
	return r
}

// SetMaintainBottomViewPadding sets the maintain bottom view padding for the sliver safe area widget.
func (r *SliverSafeAreaAttributes) SetMaintainBottomViewPadding(maintainBottomViewPadding duit_utils.Tristate[bool]) *SliverSafeAreaAttributes {
	r.MaintainBottomViewPadding = maintainBottomViewPadding
	return r
}
