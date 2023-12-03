package internal

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes"
	"github.com/lesleysin/duit_go/pkg/duit_core"
	"github.com/lesleysin/duit_go/pkg/duit_widget"
)

func Form1() []byte {
	var holder duit_core.DuitView
	builder := holder.Builder()

	child := duit_widget.SizedBoxUiElement(&duit_attributes.SizedBoxAttributes{
		Width:  100,
		Height: 100,
	}, "", false, nil).AddChild(duit_widget.CenterUiElement(nil, "", false, nil).AddChild(duit_widget.TextUiElement(&duit_attributes.TextAttributes{
		Data: "Form 1",
	}, "", false, nil)))

	root := builder.CreateRootOfExactType(duit_core.Column, nil, "")

	root.AddChild(child)
	root.AddChild(duit_widget.SizedBoxUiElement(&duit_attributes.SizedBoxAttributes{
		Height: 50,
	}, "", false, nil))
	root.AddChild(child)
	if value, err := builder.Build(); err != nil {
		panic(err)
	} else {
		return value
	}
}
