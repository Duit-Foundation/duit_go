package duit_props

type TransformType string

const (
	TransformTypeRotate    TransformType = "rotate"
	TransformTypeScale     TransformType = "scale"
	TransformTypeFlip      TransformType = "flip"
	TransformTypeTranslate TransformType = "translate"
)

type TransformBase struct {
	Origin            *Offset       `json:"origin,omitempty"`
	Alignment         Alignment     `json:"alignment,omitempty"`
	TransformHitTests bool          `json:"transformHitTests,omitempty"`
	FilterQuality     FilterQuality `json:"filterQuality,omitempty"`
}

type ScaleTransform struct {
	*TransformBase
	Scale  float32 `json:"scale,omitempty"`
	ScaleX float32 `json:"scaleX,omitempty"`
	ScaleY float32 `json:"scaleY,omitempty"`
}

type RotateTransform struct {
	*TransformBase
	Angle float32 `json:"angle,omitempty"`
}

type TranslateTransform struct {
	*TransformBase
	Offset *Offset `json:"offset,omitempty"`
}

type FlipTransform struct {
	*TransformBase
	FlipX bool `json:"flipX,omitempty"`
	FlipY bool `json:"flipY,omitempty"`
}

type Transform interface {
	ScaleTransform | TranslateTransform | FlipTransform | RotateTransform
}
