package duit_attributes

import (
	"encoding/json"
	"errors"

	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_clip"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_cross_axis"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_main_axis"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_text_properties"
)

type WrapAttributes struct {
	TextDirection      duit_text_properties.TextDirection `json:"textDirection,omitempty"`
	VerticalDirection  duit_flex.VerticalDirection        `json:"verticalDirection,omitempty"`
	Alignment          duit_main_axis.MainAxisAlignment   `json:"alignment,omitempty"`
	RunAlignment       duit_main_axis.MainAxisAlignment   `json:"runAlignment,omitempty"`
	CrossAxisAlignment duit_cross_axis.CrossAxisAlignment `json:"crossAxisAlignment,omitempty"`
	Spacing            float32                            `json:"spacing,omitempty"`
	RunSpacing         float32                            `json:"runSpacing,omitempty"`
	ClipBehavior       duit_clip.Clip                     `json:"clipBehavior,omitempty"`
	Direction          duit_flex.Axis                     `json:"direction,omitempty"`
}

func (attrs *WrapAttributes) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(attrs)

	if err != nil {
		return nil, err
	}

	if attrs.CrossAxisAlignment == duit_cross_axis.Baseline || attrs.CrossAxisAlignment == duit_cross_axis.Stretch {
		return nil, errors.New("CrossAxisAlignment property cannot have 'baseline' or 'stretch' values for curret kind of attribute")
	}

	return data, nil
}
