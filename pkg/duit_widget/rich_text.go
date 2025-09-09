package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func RichText(attributes *duit_attributes.RichTextAttributes, id string, controlled bool) *duit_core.DuitElementModel {
	checkAttributes(attributes)
	return new(duit_core.DuitElementModel).CreateElement(duit_core.RichText, id, "", attributes, nil, controlled, 0, nil)
}
