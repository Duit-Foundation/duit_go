package duit_attributes

import (
	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SafeAreaAttributes[TInsets duit_edge_insets.EdgeInsets] struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	Left                      duit_utils.Tristate[bool] `json:"left,omitempty"`
	Top                       duit_utils.Tristate[bool] `json:"top,omitempty"`
	Right                     duit_utils.Tristate[bool] `json:"right,omitempty"`
	Bottom                    duit_utils.Tristate[bool] `json:"bottom,omitempty"`
	Minimum                   TInsets                   `json:"minimum,omitempty"`
	MaintainBottomViewPadding duit_utils.Tristate[bool] `json:"maintainBottomViewPadding,omitempty"`
}

func (r *SafeAreaAttributes[TInsets]) Validate() error {
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
