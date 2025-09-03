package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

func Image[T duit_props.Color](attributes *duit_attributes.ImageAttributes[T], id string, controlled bool) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.Image, id, "", attributes, nil, controlled, 0, nil)
}
