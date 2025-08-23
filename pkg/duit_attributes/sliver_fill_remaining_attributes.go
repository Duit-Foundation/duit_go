package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"

type SliverFillRemainingAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*SliverProps
	HasScrollBody  duit_utils.Tristate[bool] `json:"hasScrollBody,omitempty"`
	FillOverscroll duit_utils.Tristate[bool] `json:"fillOverscroll,omitempty"`
}

func (r *SliverFillRemainingAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}
