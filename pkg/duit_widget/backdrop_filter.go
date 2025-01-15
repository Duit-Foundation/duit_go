package duit_widget

import (
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_painting"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func BackdropFilter[T duit_painting.ImageFilter](attributes *duit_attributes.BackdropFilterAttributes[T], id string, controlled bool, child *duit_core.DuitElementModel) *duit_core.DuitElementModel {
	return new(duit_core.DuitElementModel).CreateElement(duit_core.BackdropFilter, id, "", attributes, nil, controlled, 1, nil, child)
}
