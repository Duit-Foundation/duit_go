package internal

import (
	"fmt"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_alignment"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_decorations"
	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func DecoratedBoxExample() []byte {
	var holder duit_core.DuitView
	builder := holder.Builder()

	child := duit_widget.SizedBoxUiElement(&duit_attributes.SizedBoxAttributes{
		Width:  100,
		Height: 100,
	}, "", false)

	child.AddChild(duit_widget.DecoratedBoxUiElement[duit_color.ColorString](
		&duit_attributes.DecoratedBoxAttributes[duit_color.ColorString]{
			Decoration: &duit_decoration.BoxDecoration[duit_color.ColorString]{
				Gradient: &duit_decoration.LinearGradient[duit_color.ColorString]{
					Colors:        []duit_color.ColorString{"#4287f5", "#f57b42"},
					RotationAngle: 25,
					Begin:         duit_alignment.TopLeft,
					End:           duit_alignment.BottomRight,
				},
				Boder: &duit_decoration.BorderSide[duit_color.ColorString]{
					Color: duit_color.ColorString("#f57b42"),
					Width: 3,
				},
				BorderRadius: 12,
			},
		},
		"",
		false).AddChild(
		duit_widget.CenterUiElement(&duit_attributes.CenterAttributes{}, "", false).AddChild(
			duit_widget.TextUiElement[duit_color.ColorString](
				&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Centred Text"},
				"",
				false),
		)))

	builder.RootFrom(child)

	if value, err := builder.Build(); err != nil {
		fmt.Println(err.Error())
		return []byte{}
	} else {
		return value
	}
}
