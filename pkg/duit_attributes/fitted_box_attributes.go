package duit_attributes

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_alignment"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_clip"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
)

type FittedBoxAttributes struct {
	ValueReferenceHolder
	Fit          duit_flex.BoxFit         `json:"fit,omitempty"`
	ClipBehavior duit_clip.Clip           `json:"clipBehavior,omitempty"`
	Alignment    duit_alignment.Alignment `json:"alignment,omitempty"`
}
