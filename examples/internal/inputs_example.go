package internal

import (
	"fmt"

	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_decorations"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_material"
	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func InputsExample() []byte {

	var holder duit_core.DuitView
	builder := holder.Builder()

	root := builder.CreateRootOfExactType(duit_core.Column, &duit_attributes.FlexAttributes{}, "")

	root.AddChildren([]*duit_core.DuitElementModel{
		duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Text input 1"}, "", false),
		duit_widget.SizedBox(&duit_attributes.SizedBoxAttributes{
			Height: 12,
		}, "", false, nil),
		duit_widget.Padding[duit_edge_insets.EdgeInsetsLTRB](&duit_attributes.PaddingAttributes[duit_edge_insets.EdgeInsetsLTRB]{
			Padding: [4]float32{12, 0, 0, 12},
		}, "", false, duit_widget.TextField[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString](&duit_attributes.TextFieldAttributes[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString]{
			Decoration: &duit_decoration.InputDecoration[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString]{
				Border: &duit_decoration.InputBorder[duit_color.ColorString]{
					Type: duit_decoration.Outline,
					Options: &duit_decoration.InputBorderOptions[duit_color.ColorString]{
						BorderSide: &duit_decoration.BorderSide[duit_color.ColorString]{
							Width: 3,
						},
					},
				},
			},
		}, "textInput1", duit_core.HttpAction("/textInput1change", []*duit_core.ActionDependency{
			{Id: "textInput1", Target: "value"},
		}, &duit_core.HttpActionMetainfo{Method: "GET"}))),
		duit_widget.SizedBox(&duit_attributes.SizedBoxAttributes{
			Height: 24,
		}, "", false, nil),
		duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Text input 2"}, "", false),
		duit_widget.SizedBox(&duit_attributes.SizedBoxAttributes{
			Height: 12,
		}, "", false, nil),
		duit_widget.Padding[duit_edge_insets.EdgeInsetsSymmentric](&duit_attributes.PaddingAttributes[duit_edge_insets.EdgeInsetsSymmentric]{
			Padding: [2]float32{0, 12},
		}, "", false, nil).AddChild(
			duit_widget.TextField[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString](&duit_attributes.TextFieldAttributes[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString]{
				Decoration: &duit_decoration.InputDecoration[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString]{
					LabelText: "TextFieldUiElement",
					Border: &duit_decoration.InputBorder[duit_color.ColorString]{
						Type: duit_decoration.Outline,
						Options: &duit_decoration.InputBorderOptions[duit_color.ColorString]{
							BorderSide: &duit_decoration.BorderSide[duit_color.ColorString]{
								Width: 3,
								Color: "#f57b42",
							},
						},
					},
				},
			}, "textInput2", duit_core.HttpAction("/textInput2change", []*duit_core.ActionDependency{
				{Id: "textInput2", Target: "value"},
			}, &duit_core.HttpActionMetainfo{Method: "POST"})),
		),
		duit_widget.Row(&duit_attributes.FlexAttributes{}, "", false, nil).AddChildren([]*duit_core.DuitElementModel{
			duit_widget.CheckBox[duit_color.ColorString](&duit_attributes.CheckboxAttributes[duit_color.ColorString]{
				Value: true,
				FillColor: &duit_material.MaterialStateProperty[duit_color.ColorString]{
					Selected: "#4287f5",
				},
			}, "checkbox1", nil),
			duit_widget.SizedBox(&duit_attributes.SizedBoxAttributes{
				Height: 12,
			}, "", false, nil),
			duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Checkbox"}, "", false),
		}),
		duit_widget.SizedBox(&duit_attributes.SizedBoxAttributes{
			Height: 24,
		}, "", false, nil),
		duit_widget.ElevatedButton(&duit_attributes.ElevatedButtonAttributes{}, "button1", duit_core.HttpAction("/apply", []*duit_core.ActionDependency{
			{Id: "textInput1", Target: "value1"},
			{Id: "textInput2", Target: "value2"},
			{Id: "checkbox1", Target: "checkbox"},
			{Id: "radioGroup1", Target: "radio"},
		}, &duit_core.HttpActionMetainfo{Method: "POST"}), nil).AddChild(
			duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "ElevatedButton"}, "", false),
		),
		duit_widget.Slider[duit_color.ColorString](
			&duit_attributes.SliderAttributes[duit_color.ColorString]{
				Value:                12,
				Min:                  0,
				Max:                  100,
				SecondaryTrackValue:  50,
				Divisions:            5,
				SecondaryActiveColor: "#eb4034",
				OnChangeEnd:          duit_core.HttpAction("/slider", []*duit_core.ActionDependency{}, &duit_core.HttpActionMetainfo{Method: "GET"}),
			}, "",
		),
		duit_widget.Padding[duit_edge_insets.EdgeInsetsAll](
			&duit_attributes.PaddingAttributes[duit_edge_insets.EdgeInsetsAll]{
				Padding: 12,
			}, "", false, nil,
		).AddChild(
			duit_widget.Container[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString](
				&duit_attributes.ContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_color.ColorString]{
					Decoration: &duit_decoration.BoxDecoration[duit_color.ColorString]{
						Color:        "#f5f5f5",
						BorderRadius: 12,
						Boder:        &duit_decoration.BorderSide[duit_color.ColorString]{Width: 1.25, Color: "#f57b42"},
					},
				}, "", false, nil,
			).AddChild(
				duit_widget.RadioGroupContext[string](
					&duit_attributes.RadioGroupContextAttributes[string]{
						GroupValue: "1",
					}, "radioGroup1", nil, nil,
				).AddChild(
					duit_widget.Column(
						&duit_attributes.FlexAttributes{}, "", false, nil,
					).AddChildren([]*duit_core.DuitElementModel{
						duit_widget.Row(
							&duit_attributes.FlexAttributes{}, "", false, nil,
						).AddChildren([]*duit_core.DuitElementModel{
							duit_widget.Radio[string, duit_color.ColorString](
								&duit_attributes.RadioAttributes[string, duit_color.ColorString]{
									Value:        "1",
									SplashRadius: 12,
									Toggleable:   true,
								}, "", false,
							),
							duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Toggle radio 1"}, "", false),
						}),
						duit_widget.Row(
							&duit_attributes.FlexAttributes{}, "", false, nil,
						).AddChildren([]*duit_core.DuitElementModel{
							duit_widget.Radio[string, duit_color.ColorString](
								&duit_attributes.RadioAttributes[string, duit_color.ColorString]{
									Value:        "2",
									SplashRadius: 24,
									Toggleable:   true,
								}, "", false,
							),
							duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Toggle radio 2"}, "", false),
						}),
						duit_widget.Row(
							&duit_attributes.FlexAttributes{}, "", false, nil).AddChildren([]*duit_core.DuitElementModel{
							duit_widget.Radio[string, duit_color.ColorString](
								&duit_attributes.RadioAttributes[string, duit_color.ColorString]{
									Value:        "3",
									SplashRadius: 36,
									Toggleable:   true,
								}, "", false,
							),
							duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Toggle radio 3"}, "", false),
						}),
					}),
				),
			),
		),
		duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Have a nice day"}, "text1", true),
		duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Have a nice day"}, "text2", true),
		duit_widget.Text[duit_color.ColorString](&duit_attributes.TextAttributes[duit_color.ColorString]{Data: "Have a nice day"}, "text3", true),
	})

	if value, err := builder.Build(); err != nil {
		fmt.Println(err.Error())
		return []byte{}
	} else {
		return value
	}
}
