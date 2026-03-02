package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:
	SkeletonBox(
		duit_attributes.NewSkeletonBoxAttributes().
			SetWidth(150).
			SetHeight(80).
			AddExtra("type", "SkeletonBox"),
		"id",
		false,
	)
*/
func SkeletonBox(attributes *duit_attributes.SkeletonBoxAttributes, id string, controlled bool) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.SkeletonBox, id, "", attributes, nil, controlled, 0, nil)
}
