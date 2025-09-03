package duit_transform

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type TransfromType string

const (
	Rotate    TransfromType = "rotate"
	Scale     TransfromType = "scale"
	Flip      TransfromType = "flip"
	Translate TransfromType = "translate"
)

type TransformBase struct {
	Origin            *duit_props.Offset           `json:"origin,omitempty"`
	Alignment         duit_props.Alignment    `json:"alignment,omitempty"`
	TransformHitTests bool                        `json:"transformHitTests,omitempty"`
	FilterQuality     duit_props.FilterQuality `json:"filterQuality,omitempty"`
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
	Offset *duit_props.Offset `json:"offset,omitempty"`
}

type FlipTransfrom struct {
	*TransformBase
	FlipX bool `json:"flipX,omitempty"`
	FlipY bool `json:"flipY,omitempty"`
}

type Transform interface {
	ScaleTransfrom | TranslateTransfrom | FlipTransfrom | RotateTransfrom
}
