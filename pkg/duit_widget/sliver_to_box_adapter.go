package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	SliverToBoxAdapter(
		"id",
		nil, //child
	)
*/
func SliverToBoxAdapter(id string, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverToBoxAdapter, id, "", nil, nil, false, 1, nil, child)
}
