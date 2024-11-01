package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func RichText[T duit_color.Color](attributes *duit_attributes.RichTextAttributes[T], id string, controlled bool) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.RichText, id, "", attributes, nil, controlled, 0, nil)
}
