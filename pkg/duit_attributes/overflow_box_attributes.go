package duit_attributes

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_alignment"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
)

type OverflowBoxAttributes struct {
	MinWidth  float32                  `json:"minWidth,omitempty"`
	MaxWidth  float32                  `json:"maxWidth,omitempty"`
	MinHeight float32                  `json:"minHeight,omitempty"`
	MaxHeight float32                  `json:"maxHeight,omitempty"`
	Alignment duit_alignment.Alignment `json:"alignment,omitempty"`
	Fit       duit_flex.OverflowBoxFit `json:"fit,omitempty"`
}
