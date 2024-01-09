package internal

import (
	"fmt"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_transform"
	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func TransformExample() []byte {
	var holder duit_core.DuitView
	builder := holder.Builder()

	root := builder.CreateRootOfExactType(duit_core.Column, nil, "")

	root.AddChildren([]*duit_core.DuitElementModel{
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "Rotate",
		}, "", false),
		duit_widget.TransformUiElement[duit_transform.RotateTransfrom](&duit_attributes.TransfromAttributes[duit_transform.RotateTransfrom]{
			Type: duit_transform.Rotate,
			Data: &duit_transform.RotateTransfrom{Angle: 45},
		}, "", false).AddChild(duit_widget.ContainerUiElement[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString](&duit_attributes.ContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString]{
			Color:  "#4287f5",
			Width:  50,
			Height: 50,
		}, "", false)),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "Translate",
		}, "", false),
		duit_widget.TransformUiElement[duit_transform.TranslateTransfrom](&duit_attributes.TransfromAttributes[duit_transform.TranslateTransfrom]{
			Type: duit_transform.Translate,
			Data: &duit_transform.TranslateTransfrom{Offset: &duit_flex.Offset{Dx: 100, Dy: 10}},
		}, "", false).AddChild(duit_widget.ContainerUiElement[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString](&duit_attributes.ContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString]{
			Color:  "#4287f5",
			Width:  50,
			Height: 50,
		}, "", false)),
		duit_widget.TextUiElement[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{
			Data: "Scale",
		}, "", false),
		duit_widget.TransformUiElement[duit_transform.ScaleTransfrom](&duit_attributes.TransfromAttributes[duit_transform.ScaleTransfrom]{
			Type: duit_transform.Scale,
			Data: &duit_transform.ScaleTransfrom{Scale: 2},
		}, "", false).AddChild(duit_widget.ContainerUiElement[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString](&duit_attributes.ContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString]{
			Color:  "#4287f5",
			Width:  50,
			Height: 50,
		}, "", false)),
	})

	if value, err := builder.Build(); err != nil {
		fmt.Println(err.Error())
		return []byte{}
	} else {
		return value
	}
}
