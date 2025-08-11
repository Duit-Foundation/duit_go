package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
)

/*
Example:
	SliverGrid(
		&SliverGridCountAttributes{
			CrossAxisCount: 2,
			MainAxisSpacing: 8.0,
			CrossAxisSpacing: 8.0,
			ChildAspectRatio: 1.0,
		},
		"id",
		false,
		children,
	)
*/
func SliverGrid[T duit_attributes.SliverGridAttributes](attributes *T, id string, controlled bool, children []*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverGrid, id, "", attributes, nil, controlled, 2, nil, children...)
} 