package duit_transform

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_alignment"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_painting"
)

type TransfromType string

const (
	Rotate    TransfromType = "rotate"
	Scale     TransfromType = "scale"
	Flip      TransfromType = "flip"
	Translate TransfromType = "translate"
)

type TransformBase struct {
	Origin            *duit_flex.Offset           `json:"origin,omitempty"`
	Alignment         duit_alignment.Alignment    `json:"alignment,omitempty"`
	TransformHitTests bool                        `json:"transformHitTests,omitempty"`
	FilterQuality     duit_painting.FilterQuality `json:"filterQuality,omitempty"`
}

type ScaleTransfrom struct {
	*TransformBase
	Scale  float32 `json:"scale,omitempty"`
	ScaleX float32 `json:"scaleX,omitempty"`
	ScaleY float32 `json:"scaleY,omitempty"`
}

type RotateTransfrom struct {
	*TransformBase
	Angle float32 `json:"angle,omitempty"`
}

type TranslateTransfrom struct {
	*TransformBase
	Offset *duit_flex.Offset `json:"offset,omitempty"`
}

type FlipTransfrom struct {
	*TransformBase
	FlipX bool `json:"flipX,omitempty"`
	FlipY bool `json:"flipY,omitempty"`
}

type Transform interface {
	ScaleTransfrom | TranslateTransfrom | FlipTransfrom | RotateTransfrom
}
