package internal

import (
	"fmt"
	"os"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_painting"

	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func ImageExample() []byte {

	var holder duit_core.DuitView
	builder := holder.Builder()

	root := builder.CreateRootOfExactType(duit_core.Column, nil, "")

	file, err := os.ReadFile("assets/vlad_ten.png")

	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}

	root.AddChildren([]*duit_core.DuitElementModel{
		duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Netword image"}, "", false),
		duit_widget.Image(&duit_attributes.ImageAttributes[duit_color.ColorString]{
			Type:   duit_painting.Network,
			Src:    "https://via.assets.so/game.png?id=1&q=95&w=360&h=360&fit=fill",
			Width:  350,
			Height: 350,
		}, "", false),
		duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Asset image"}, "", false),
		duit_widget.Image(&duit_attributes.ImageAttributes[duit_color.ColorString]{
			Type:   duit_painting.Asset,
			Src:    "assets/betboom.jpg",
			Height: 350,
		}, "", false),
		duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Memory image"}, "", false),
		duit_widget.Image(&duit_attributes.ImageAttributes[duit_color.ColorString]{
			Type: duit_painting.Memory,
			Src:  "",
			ByteData: &duit_painting.ImageBuffer{
				Type: duit_painting.ByteBuffer,
				Data: file,
			},
		}, "", false),
		duit_widget.Empty(),
	})

	if value, err := builder.Build(); err != nil {
		fmt.Println(err.Error())
		return []byte{}
	} else {
		return value
	}
}
