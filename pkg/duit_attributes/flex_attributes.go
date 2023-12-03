package duit_attributes

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_clip"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_cross_axis"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_main_axis"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_text_properties"
)

type FlexAttributes struct {
	MainAxisAlignment  duit_main_axis.MainAxisAlignment   `json:"mainAxisAlignment,omitempty"`
	MainAxisSize       duit_main_axis.MainAxisSize        `json:"mainAxisSize,omitempty"`
	CrossAxisAlignment duit_cross_axis.CrossAxisAlignment `json:"crossAxisAlignment,omitempty"`
	TextDirection      duit_text_properties.TextDirection `json:"textDirection,omitempty"`
	VerticalDirection  duit_flex.VerticalDirection        `json:"verticalDirection,omitempty"`
	ClipBehavior       duit_clip.Clip                     `json:"clipBehavior,omitempty"`
}
