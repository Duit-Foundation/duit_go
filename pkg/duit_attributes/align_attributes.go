package duit_attributes

import "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_alignment"

type AlignAttributes struct {
	Alignment    duit_alignment.Alignment `json:"alignment,omitempty"`
	WidthFactor  float32                  `json:"widthFactor,omitempty"`
	HeightFactor float32                  `json:"heightFactor,omitempty"`
}
