package duit_props

import "encoding/json"

// ImageFilter represents a filter that can be applied to an image.
// It is used to create a filter effect on an image.
//
// This is a Go implementation of Flutter's ImageFilter class.
// For more information, see: https://api.flutter.dev/flutter/dart-ui/ImageFilter-class.html
type ImageFilter struct {
	data any
}

func (r *ImageFilter) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }

type blurImageFilter struct {
	Type     uint8    `json:"type"`
	SigmaX   float32  `json:"sigmaX,omitempty"`
	SigmaY   float32  `json:"sigmaY,omitempty"`
	TileMode TileMode `json:"tileMode,omitempty"`
}

// https://api.flutter.dev/flutter/dart-ui/ImageFilter/ImageFilter.blur.html
func NewBlurImageFilter(sigmaX, sigmaY float32, tileMode TileMode) *ImageFilter {
	m := &blurImageFilter{
		Type:   0,
		SigmaX: sigmaX,
		SigmaY: sigmaY,
	}

	if len(tileMode) > 0 {
		m.TileMode = tileMode
	}

	return &ImageFilter{data: m}
}

type composeImageFilter struct {
	Type  uint8 `json:"type"`
	Outer any   `json:"outer,omitempty"`
	Inner any   `json:"inner,omitempty"`
}

// https://api.flutter.dev/flutter/dart-ui/ImageFilter/ImageFilter.compose.html
func NewComposeImageFilter(outer, inner any) *ImageFilter {

	switch outer.(type) {
	case *blurImageFilter:
	case *matrixImageFilter:
	case *composeImageFilter:
	case *dilateImageFilter:
	case *erodeImageFilter:
	case nil:
		break
	default:
		panic("unknown outer filter model")
	}

	switch inner.(type) {
	case *blurImageFilter:
	case *matrixImageFilter:
	case *composeImageFilter:
	case *dilateImageFilter:
	case *erodeImageFilter:
	case nil:
		break
	default:
		panic("unknown outer filter model")
	}

	return &ImageFilter{data: &composeImageFilter{
		Type:  1,
		Outer: outer,
		Inner: inner,
	}}
}

type dilateImageFilter struct {
	Type    uint8   `json:"type"`
	RadiusX float32 `json:"radiusX"`
	RaduisY float32 `json:"radiusY"`
}

// https://api.flutter.dev/flutter/dart-ui/ImageFilter/ImageFilter.dilate.html
func NewDilateImageFilter(radiusX, radiusY float32) *ImageFilter {
	return &ImageFilter{data: &dilateImageFilter{
		Type:    2,
		RadiusX: radiusX,
		RaduisY: radiusY,
	}}
}

type erodeImageFilter struct {
	Type    uint8   `json:"type"`
	RadiusX float32 `json:"radiusX"`
	RaduisY float32 `json:"radiusY"`
}

// https://api.flutter.dev/flutter/dart-ui/ImageFilter/ImageFilter.erode.html
func NewErodeImageFilter(radiusX, radiusY float32) *ImageFilter {
	return &ImageFilter{data: &erodeImageFilter{
		Type:    3,
		RadiusX: radiusX,
		RaduisY: radiusY,
	}}
}

type matrixImageFilter struct {
	Type          uint8         `json:"type"`
	Matrix        [16]float64   `json:"matrix4"`
	FilterQuality FilterQuality `json:"filterQuality,omitempty"`
}

// https://api.flutter.dev/flutter/dart-ui/ImageFilter/ImageFilter.matrix.html
func NewMatrixImageFilter(matrix [16]float64, filterQuality FilterQuality) *ImageFilter {
	m := &matrixImageFilter{
		Type:          4,
		Matrix:        matrix,
		FilterQuality: filterQuality,
	}

	if len(filterQuality) > 0 {
		m.FilterQuality = filterQuality
	}

	return &ImageFilter{data: m}
}
