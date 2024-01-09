package internal

import (
	"fmt"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_decorations"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"

	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func StackExample() []byte {

	var holder duit_core.DuitView
	builder := holder.Builder()

	root := builder.CreateRootOfExactType(duit_core.Stack, nil, "")

	root.AddChildren([]*duit_core.DuitElementModel{
		duit_widget.PositionedUiElement(&duit_attributes.PositionedAttributes{
			Top:  25,
			Left: 25,
		}, "", false).AddChild(
			duit_widget.ContainerUiElement[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString](&duit_attributes.ContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString]{
				Width:  50,
				Height: 50,
				Color:  "#9e2f2f",
			}, "", false),
		),
		duit_widget.PositionedUiElement(&duit_attributes.PositionedAttributes{
			Top:  50,
			Left: 50,
		}, "", false).AddChild(
			duit_widget.ContainerUiElement[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString](&duit_attributes.ContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString]{
				Width:  75,
				Height: 75,
				Decoration: &duit_decoration.BoxDecoration[duit_color.ColorString]{
					Color: "#32a852",
					Shape: duit_decoration.Circle,
					Boder: &duit_decoration.BorderSide[duit_color.ColorString]{
						Width: 2.5,
						Color: "#FFFFF",
					},
				},
			}, "", false),
		),
		duit_widget.PositionedUiElement(&duit_attributes.PositionedAttributes{
			Top:  75,
			Left: 75,
		}, "", false).AddChild(
			duit_widget.SizedBoxUiElement(&duit_attributes.SizedBoxAttributes{
				Width:  50,
				Height: 50,
			}, "", false).AddChild(
				duit_widget.ColoredBoxUiElement[duit_color.ColorRGBO](&duit_attributes.ColoredBoxAttributes[duit_color.ColorRGBO]{
					Color: duit_color.ColorRGBO{R: 255, G: 255, B: 255, O: 1},
				}, "", false),
			),
		),
	})

	if value, err := builder.Build(); err != nil {
		fmt.Println(err.Error())
		return []byte{}
	} else {
		return value
	}
}
