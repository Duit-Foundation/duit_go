package internal

import (
	"fmt"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_text_properties"
	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func WsExample() []byte {
	var holder duit_core.DuitView
	builder := holder.Builder()

	builder.RootFrom(duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
		Data: "WebSocket Example",
		Style: &duit_text_properties.TextStyle[duit_color.ColorString]{
			FontWeight: duit_text_properties.W900,
		},
	}, "", false))

	value, err := builder.Build()

	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}

	return value
}
