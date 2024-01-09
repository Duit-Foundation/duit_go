package internal

import (
	"fmt"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func UpdateExample() []byte {
	var holder duit_core.DuitView
	builder := holder.Builder()

	root := builder.CreateRootOfExactType(duit_core.Column, nil, "")

	root.AddChildren([]*duit_core.DuitElementModel{
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Layout update example"}, "", false),
		duit_widget.ElevatedButtonUiElement(&duit_attributes.ElevatedButtonAttributes{}, "button1", duit_core.CreateHttpAction("/updateLayout", nil, &duit_core.HttpActionMetainfo{Method: "GET"})).AddChild(
			duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "ElevatedButton"}, "", false),
		),
	})

	if value, err := builder.Build(); err != nil {
		fmt.Println(err.Error())
		return []byte{}
	} else {
		return value
	}
}

func UpdatePayload() *duit_core.DuitElementModel {
	var holder duit_core.DuitView
	builder := holder.Builder()

	root := builder.CreateRootOfExactType(duit_core.Column, nil, "")

	root.AddChildren([]*duit_core.DuitElementModel{
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Force update!"}, "", false),
	})

	return root
}
