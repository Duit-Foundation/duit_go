package internal

import (
	"fmt"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_alignment"
	duit_decoration "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_decorations"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"
	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func DecoratedBoxExample() []byte {
	var holder duit_core.DuitView
	builder := holder.Builder()

	child := duit_widget.SizedBoxUiElement(&duit_attributes.SizedBoxAttributes{
		Width:  100,
		Height: 100,
	}, "", false, nil)

	child.AddChild(duit_widget.DecoratedBoxUiElement[duit_utility_styles.ColorString](
		&duit_attributes.DecoratedBoxAttributes[duit_utility_styles.ColorString]{
			Decoration: duit_decoration.BoxDecoration[duit_utility_styles.ColorString]{
				Gradient: duit_decoration.LinearGradient[duit_utility_styles.ColorString]{
					Colors:        []duit_utility_styles.ColorString{"#4287f5", "#f57b42"},
					RotationAngle: 25,
					Begin:         duit_alignment.TopLeft,
					End:           duit_alignment.BottomRight,
				},
				Boder: duit_decoration.BorderSide[duit_utility_styles.ColorString]{
					Color: duit_utility_styles.ColorString("#f57b42"),
					Width: 3,
				},
				BorderRadius: 12,
			},
		},
		"",
		false,
		nil).AddChild(
		duit_widget.CenterUiElement(&duit_attributes.CenterAttributes{}, "", false, nil).AddChild(
			duit_widget.TextUiElement[duit_utility_styles.ColorString](
				&duit_attributes.TextAttributes[duit_utility_styles.ColorString]{Data: "Centred Text"},
				"",
				false,
				nil,
			),
		)))

	builder.RootFrom(child)

	if value, err := builder.Build(); err != nil {
		fmt.Println(err.Error())
		return []byte{}
	} else {
		return value
	}
}
