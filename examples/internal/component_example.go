package internal

import (
	"fmt"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_decorations"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_main_axis"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_painting"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_text_properties"
	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func ComponentTemplate() *duit_core.DuitElementModel {
	layout := duit_widget.DecoratedBox[duit_color.ColorString](
		&duit_attributes.DecoratedBoxAttributes[duit_color.ColorString]{
			Decoration: &duit_decoration.BoxDecoration[duit_color.ColorString]{
				Color:        "#32a852",
				BorderRadius: 8.0,
				BoxShadow: []duit_decoration.BoxShadow[duit_color.ColorString]{
					{
						Color: "#DCDCDC",
						Offset: duit_flex.Offset{
							Dx: 2.0,
							Dy: 2.0,
						},
						BlurRadius:   2.0,
						SpreadRadius: 2.0,
					},
				},
			},
		}, "", false, nil).AddChild(duit_widget.Padding[duit_edge_insets.EdgeInsetsAll](
		&duit_attributes.PaddingAttributes[duit_edge_insets.EdgeInsetsAll]{
			Padding: 16.0,
		}, "", false, nil)).AddChild(duit_widget.Row(
		&duit_attributes.FlexAttributes{}, "", false, nil)).AddChildren([]*duit_core.DuitElementModel{
		duit_widget.Image[duit_color.ColorString](
			&duit_attributes.ImageAttributes[duit_color.ColorString]{
				Type: duit_painting.Network,
				Fit:  duit_flex.Contain,
				Src:  "",
				ValueReferenceHolder: duit_attributes.ValueReferenceHolder{
					Refs: []*duit_attributes.ValueRef{
						duit_attributes.Ref("src", "image"),
					},
				},
			}, "", false),
		duit_widget.SizedBox(&duit_attributes.SizedBoxAttributes{
			Width: 16.0,
		}, "", false, nil),
		duit_widget.Column(&duit_attributes.FlexAttributes{
			MainAxisAlignment: duit_main_axis.SpaceEvenly,
		}, "", false, nil).AddChildren([]*duit_core.DuitElementModel{
			duit_widget.Text[duit_color.ColorString](
				&duit_attributes.TextAttributes[duit_color.ColorString]{
					Style: &duit_text_properties.TextStyle[duit_color.ColorString]{
						FontWeight: duit_text_properties.W900,
						FontSize:   24,
						Color:      "#192e1e",
					},
					ValueReferenceHolder: duit_attributes.ValueReferenceHolder{
						Refs: []*duit_attributes.ValueRef{
							duit_attributes.Ref("data", "primary"),
						},
					},
				}, "", false,
			),
			duit_widget.Text[duit_color.ColorString](
				&duit_attributes.TextAttributes[duit_color.ColorString]{
					Style: &duit_text_properties.TextStyle[duit_color.ColorString]{
						FontWeight: duit_text_properties.W500,
						FontSize:   16,
						Color:      "#192e1e",
					},
					ValueReferenceHolder: duit_attributes.ValueReferenceHolder{
						Refs: []*duit_attributes.ValueRef{
							duit_attributes.Ref("data", "description"),
						},
					},
				}, "", false,
			),
			duit_widget.Text[duit_color.ColorString](
				&duit_attributes.TextAttributes[duit_color.ColorString]{
					Style: &duit_text_properties.TextStyle[duit_color.ColorString]{
						FontWeight: duit_text_properties.W400,
						FontSize:   16,
						Color:      "#192e1e",
					},
					ValueReferenceHolder: duit_attributes.ValueReferenceHolder{
						Refs: []*duit_attributes.ValueRef{
							duit_attributes.Ref("data", "cost"),
						},
					},
				}, "", false,
			),
		}),
	})

	return layout
}

type dataStruct struct {
	Primary     string `json:"primary"`
	Description string `json:"description"`
	Cost        string `json:"cost"`
	Image       string `json:"image"`
}

func ComponentExample() []byte {
	var holder duit_core.DuitView
	builder := holder.Builder()

	root := builder.CreateRootOfExactType(duit_core.Column, nil, "")

	root.AddChildren(
		[]*duit_core.DuitElementModel{
			duit_widget.SizedBox(&duit_attributes.SizedBoxAttributes{
				Height: 24.0,
			}, "", false, nil),
			duit_widget.Component(&dataStruct{
				Primary:     "Лабутены",
				Description: "Топовые и дорогие",
				Cost:        "$10000",
				Image:       "https://resizer.mail.ru/p/b0685008-5021-5715-a506-e73621c958f5/AQAFGR6McYd06E2eQ8J-GzknNO9eDHZcCIpgrre-K3SwrZ0QnmOBQVPyD6yWisBlYJYZ38YEM768jMY1M4m4NYbQkmM.jpg",
			}, "shoes_card", "shoes_card1"),
			duit_widget.SizedBox(&duit_attributes.SizedBoxAttributes{
				Height: 24.0,
			}, "", false, nil),
			duit_widget.Component(&dataStruct{
				Primary:     "Найки",
				Description: "Прямо с рынка",
				Cost:        "$10",
				Image:       "https://ae04.alicdn.com/kf/A1255f8c200964840b940a3c7537ab567g.jpeg",
			}, "shoes_card", "shoes_card2"),
			duit_widget.SizedBox(&duit_attributes.SizedBoxAttributes{
				Height: 24.0,
			}, "", false, nil),
			duit_widget.Component(&dataStruct{
				Primary:     "Туфельки",
				Description: "Очень крутые!",
				Cost:        "$150",
				Image:       "https://static.lichi.com/product/44640/9d56d4435e5bc814b979c2ebcffdb9b7.jpg?v=0_44640.0&resize=size-middle",
			}, "shoes_card", "shoes_card3"),
			duit_widget.SizedBox(&duit_attributes.SizedBoxAttributes{
				Height: 24.0,
			}, "", false, nil),
			duit_widget.ElevatedButton(&duit_attributes.ElevatedButtonAttributes{}, "button1", duit_core.HttpAction("/send_component_update", nil, &duit_core.HttpActionMetainfo{Method: "GET"}), nil),
		},
	)

	if value, err := builder.Build(); err != nil {
		fmt.Println(err.Error())
		return []byte{}
	} else {
		return value
	}
}
