package internal

import (
	"fmt"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_gestures"
	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func ScrollviewExample() []byte {
	var holder duit_core.DuitView
	builder := holder.Builder()

	builder.RootFrom(duit_widget.SingleChildScrollViewUiElement[duit_edge_insets.EdgeInsetsAll](
		&duit_attributes.SingleChildScrollViewAttributes[duit_edge_insets.EdgeInsetsAll]{
			Physics: duit_gestures.NeverScrollableScrollPhysics,
		},
		"",
		false,
	).AddChild(duit_widget.ColumnUiElement(
		nil, "", false,
	).AddChildren([]*duit_core.DuitElementModel{
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "List item",
		}, "", false),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "The end!",
		}, "", false),
	})))

	if value, err := builder.Build(); err != nil {
		fmt.Println(err.Error())
		return []byte{}
	} else {
		return value
	}
}
