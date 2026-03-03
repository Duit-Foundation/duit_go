package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

/*
Example:

	Dismissible(
		duit_attributes.NewDismissibleAttributes().SetDirection("horizontal"),
		"id",
		false,
		child,      // children[0] - required
		background, // children[1] - optional
		nil,        // children[2] - secondaryBackground, optional
	)

Multichild slots:
  - child (required)
  - background (optional)
  - secondaryBackground (optional)
*/
func Dismissible(attributes *duit_attributes.DismissibleAttributes, id string, controlled bool, child *duit_core.DuitElementModel, background *duit_core.DuitElementModel, secondaryBackground *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	children := []*duit_core.DuitElementModel{child, background, secondaryBackground}
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Dismissible, id, "", attributes, nil, controlled, 3, nil, children...)
}
