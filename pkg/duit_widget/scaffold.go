package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func Scaffold[TColor duit_props.Color](
	attributes *duit_attributes.ScaffoldAttributes[TColor],
	id string,
	controlled bool,
	body,
	appBar,
	bottomNavigationBar,
	bottomSheet,
	floatingActionButton *duit_core.DuitElementModel,
	persistentFooterButtons []*duit_core.DuitElementModel,
) *duit_core.DuitElementModel {
	children := make([]*duit_core.DuitElementModel, 5)

	children[0] = body
	children[1] = appBar
	children[2] = floatingActionButton
	children[3] = bottomSheet
	children[4] = bottomNavigationBar

	if len(persistentFooterButtons) > 0 {
		children = append(children, persistentFooterButtons...)
	}

	if attributes == nil {
		panic("attributes must not be nil")
	}

	if err := attributes.Validate(); err != nil {
		panic(err)
	}

	return new(duit_core.DuitElementModel).CreateElement(duit_core.Scaffold, id, "", attributes, nil, controlled, 2, nil, children...)
}
