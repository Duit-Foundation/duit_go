package internal

import (
	"fmt"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func GestureExample() []byte {
	var holder duit_core.DuitView
	builder := holder.Builder()

	root := builder.CreateRootOfExactType(duit_core.Column, nil, "")

	root.AddChildren([]*duit_core.DuitElementModel{
		duit_widget.GestureDetector(&duit_attributes.GestureDetectorAttributes{
			OnTap:       duit_core.HttpAction("/tap", nil, &duit_core.HttpActionMetainfo{Method: "GET"}),
			OnLongPress: duit_core.HttpAction("/longPress", nil, &duit_core.HttpActionMetainfo{Method: "GET"}),
		}, "", nil).AddChild(duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Tap/LongPress to fire event"}, "", false)),
	})

	if value, err := builder.Build(); err != nil {
		fmt.Println(err.Error())
		return []byte{}
	} else {
		return value
	}
}
