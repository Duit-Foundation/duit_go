package duit_props

import (
	"encoding/json"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type TransformType = string

const (
	TransformTypeRotate    TransformType = "rotate"
	TransformTypeScale     TransformType = "scale"
	TransformTypeFlip      TransformType = "flip"
	TransformTypeTranslate TransformType = "translate"
)

// Transform represents a transformation that can be applied to a widget.
// It is used to create a transformation effect on a widget.
//
// This is a Go implementation of Flutter's Transform class.
// For more information, see: https://api.flutter.dev/flutter/widgets/Transform-class.html
type Transform struct {
	data any
}

func (r *Transform) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }

// SetOrigin sets the origin of the transformation.
func (r *Transform) SetOrigin(origin *Offset) *Transform {

	switch r.data.(type) {
	case *scaleTransform:
		r.data.(*scaleTransform).Origin = origin
	case *rotateTransform:
		r.data.(*rotateTransform).Origin = origin
	case *translateTransform:
		r.data.(*translateTransform).Origin = origin
	case *flipTransform:
		r.data.(*flipTransform).Origin = origin
	}
	return r
}

// SetAlignment sets the alignment of the transformation.
func (r *Transform) SetAlignment(alignment Alignment) *Transform {
	switch r.data.(type) {
	case *scaleTransform:
		r.data.(*scaleTransform).Alignment = alignment
	case *rotateTransform:
		r.data.(*rotateTransform).Alignment = alignment
	case *translateTransform:
		r.data.(*translateTransform).Alignment = alignment
	case *flipTransform:
		r.data.(*flipTransform).Alignment = alignment
	}
	return r
}

// SetTransformHitTests sets the transform hit tests of the transformation.
func (r *Transform) SetTransformHitTests(transformHitTests bool) *Transform {
	switch r.data.(type) {
	case *scaleTransform:
		r.data.(*scaleTransform).TransformHitTests = duit_utils.BoolValue(transformHitTests)
	case *rotateTransform:
		r.data.(*rotateTransform).TransformHitTests = duit_utils.BoolValue(transformHitTests)
	case *translateTransform:
		r.data.(*translateTransform).TransformHitTests = duit_utils.BoolValue(transformHitTests)
	case *flipTransform:
		r.data.(*flipTransform).TransformHitTests = duit_utils.BoolValue(transformHitTests)
	}
	return r
}

// SetFilterQuality sets the filter quality of the transformation.
func (r *Transform) SetFilterQuality(filterQuality FilterQuality) *Transform {
	switch r.data.(type) {
	case *scaleTransform:
		r.data.(*scaleTransform).FilterQuality = filterQuality
	case *rotateTransform:
		r.data.(*rotateTransform).FilterQuality = filterQuality
	case *translateTransform:
		r.data.(*translateTransform).FilterQuality = filterQuality
	case *flipTransform:
		r.data.(*flipTransform).FilterQuality = filterQuality
	}
	return r
}

type scaleTransform struct {
	Origin            *Offset                   `json:"origin,omitempty"`
	Alignment         Alignment                 `json:"alignment,omitempty"`
	TransformHitTests duit_utils.Tristate[bool] `json:"transformHitTests,omitempty"`
	FilterQuality     FilterQuality             `json:"filterQuality,omitempty"`
	Scale             float32                   `json:"scale,omitempty"`
	ScaleX            float32                   `json:"scaleX,omitempty"`
	ScaleY            float32                   `json:"scaleY,omitempty"`
}

// https://api.flutter.dev/flutter/widgets/Transform/Transform.scale.html
func NewScaleTransform() *Transform {
	return &Transform{data: &scaleTransform{}}
}

// SetScale sets the scale of the transformation.
func (r *Transform) SetScale(scale float32) *Transform {
	r.data.(*scaleTransform).Scale = scale
	return r
}

// SetScaleX sets the scale x of the transformation.
func (r *Transform) SetScaleX(scaleX float32) *Transform {
	r.data.(*scaleTransform).ScaleX = scaleX
	return r
}

// SetScaleY sets the scale y of the transformation.
func (r *Transform) SetScaleY(scaleY float32) *Transform {
	r.data.(*scaleTransform).ScaleY = scaleY
	return r
}

type rotateTransform struct {
	Origin            *Offset                   `json:"origin,omitempty"`
	Alignment         Alignment                 `json:"alignment,omitempty"`
	TransformHitTests duit_utils.Tristate[bool] `json:"transformHitTests,omitempty"`
	FilterQuality     FilterQuality             `json:"filterQuality,omitempty"`
	Angle             float32                   `json:"angle,omitempty"`
}

// https://api.flutter.dev/flutter/widgets/Transform/Transform.rotate.html
func NewRotateTransform() *Transform {
	return &Transform{data: &rotateTransform{}}
}

// SetAngle sets the angle of the transformation.
func (r *Transform) SetAngle(angle float32) *Transform {
	r.data.(*rotateTransform).Angle = angle
	return r
}

type translateTransform struct {
	Origin            *Offset                   `json:"origin,omitempty"`
	Alignment         Alignment                 `json:"alignment,omitempty"`
	TransformHitTests duit_utils.Tristate[bool] `json:"transformHitTests,omitempty"`
	FilterQuality     FilterQuality             `json:"filterQuality,omitempty"`
	Offset            *Offset                   `json:"offset,omitempty"`
}

// https://api.flutter.dev/flutter/widgets/Transform/Transform.translate.html
func NewTranslateTransform() *Transform {
	return &Transform{data: &translateTransform{}}
}

// SetOffset sets the offset of the transformation.
func (r *Transform) SetOffset(offset *Offset) *Transform {
	r.data.(*translateTransform).Offset = offset
	return r
}

type flipTransform struct {
	Origin            *Offset                   `json:"origin,omitempty"`
	Alignment         Alignment                 `json:"alignment,omitempty"`
	TransformHitTests duit_utils.Tristate[bool] `json:"transformHitTests,omitempty"`
	FilterQuality     FilterQuality             `json:"filterQuality,omitempty"`
	FlipX             bool                      `json:"flipX,omitempty"`
	FlipY             bool                      `json:"flipY,omitempty"`
}

// https://api.flutter.dev/flutter/widgets/Transform/Transform.flip.html
func NewFlipTransform() *Transform {
	return &Transform{data: &flipTransform{}}
}

// SetFlipX sets the flip x of the transformation.
func (r *Transform) SetFlipX(flipX bool) *Transform {
	r.data.(*flipTransform).FlipX = flipX
	return r
}

// SetFlipY sets the flip y of the transformation.
func (r *Transform) SetFlipY(flipY bool) *Transform {
	r.data.(*flipTransform).FlipY = flipY
	return r
}
