package internal

import (
	"fmt"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_text_properties"
	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func RichTextExample() []byte {
	var holder duit_core.DuitView
	builder := holder.Builder()

	child := duit_widget.SizedBoxUiElement(&duit_attributes.SizedBoxAttributes{
		Width:  100,
		Height: 100,
	}, "", false)

	child.AddChild(
		duit_widget.RichTextUiElement[duit_color.ColorString](&duit_attributes.RichTextAttributes[duit_color.ColorString]{
			TextSpan: &duit_text_properties.TextSpan[duit_color.ColorString]{
				Children: []*duit_text_properties.TextSpan[duit_color.ColorString]{
					{
						Text: "Hello, ",
						Style: &duit_text_properties.TextStyle[duit_color.ColorString]{
							Color: "#32a852",
						},
					},
					{
						Text: "world!",
						Style: &duit_text_properties.TextStyle[duit_color.ColorString]{
							Color: "#192e1e",
						},
					},
					{
						Style: &duit_text_properties.TextStyle[duit_color.ColorString]{
							FontWeight:    duit_text_properties.W900,
							FontSize:      24,
							LetterSpacing: 3,
							Color:         "#c43777",
						},
						Text: "Duit GO",
					},
					{
						Text: "Already here!",
						Style: &duit_text_properties.TextStyle[duit_color.ColorString]{
							Color: "#c47937",
						},
					},
				},
			},
		}, "", false),
	)

	builder.RootFrom(child)

	if value, err := builder.Build(); err != nil {
		fmt.Println(err.Error())
		return []byte{}
	} else {
		return value
	}
}
