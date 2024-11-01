package duit_attributes

import "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_alignment"

type AlignAttributes struct {
	ValueReferenceHolder
	Alignment    duit_alignment.Alignment `json:"alignment,omitempty"`
	WidthFactor  float32                  `json:"widthFactor,omitempty"`
	HeightFactor float32                  `json:"heightFactor,omitempty"`
}
