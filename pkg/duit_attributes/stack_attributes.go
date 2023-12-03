package duit_attributes

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_alignment"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_clip"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_text_properties"
)

type StackAttributes struct {
	TextDirection duit_text_properties.TextDirection `json:"textDirection,omitempty"`
	ClipBehavior  duit_clip.Clip                     `json:"clipBehavior,omitempty"`
	Alignment     duit_alignment.Alignment           `json:"alignment,omitempty"`
	Fit           duit_flex.StackFit                 `json:"fit,omitempty"`
}
