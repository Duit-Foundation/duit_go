package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	SliverList(
		&duit_attributes.SliverList{
			Type:                   duit_attributes.SliverListCommon,
		},
		"id",
		false,
		children...,
	)
*/
func SliverList[T duit_attributes.SliverListAttributes](attributes *T, id string, controlled bool, children ...*duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SliverList, id, "", attributes, nil, controlled, 2, nil, children...)
} 