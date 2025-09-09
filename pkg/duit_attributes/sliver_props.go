package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"

type SliverProps struct {
	NeedsBoxAdapter duit_utils.Tristate[bool] `json:"needsBoxAdapter,omitempty"`
}

func (r *SliverProps) SetNeedsBoxAdapter(needsBoxAdapter bool) {
	r.NeedsBoxAdapter = duit_utils.BoolValue(needsBoxAdapter)
}
