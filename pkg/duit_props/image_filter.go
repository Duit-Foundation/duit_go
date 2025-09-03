package duit_props

type BlurImageFilter struct {
	Type     uint8    `json:"type"`
	SigmaX   float32  `json:"sigmaX,omitempty"`
	SigmaY   float32  `json:"sigmaY,omitempty"`
	TileMode TileMode `json:"tileMode,omitempty"`
}

func NewBlurImageFilter(sigmaX, sigmaY float32, tileMode TileMode) *BlurImageFilter {
	m := &BlurImageFilter{
		Type:   0,
		SigmaX: sigmaX,
		SigmaY: sigmaY,
	}

	if tileMode != "" {
		m.TileMode = tileMode
	}

	return m
}

type ComposeImageFilter struct {
	Type  uint8 `json:"type"`
	Outer any   `json:"outer,omitempty"`
	Inner any   `json:"inner,omitempty"`
}

func NewComposeImageFilter(outer, inner any) *ComposeImageFilter {

	switch outer.(type) {
	case *BlurImageFilter:
	case *MatrixImageFilter:
	case *ComposeImageFilter:
	case *DilateImageFilter:
	case *ErodeImageFilter:
	case nil:
		break
	default:
		panic("unknown outer filter model")
	}

	switch inner.(type) {
	case *BlurImageFilter:
	case *MatrixImageFilter:
	case *ComposeImageFilter:
	case *DilateImageFilter:
	case *ErodeImageFilter:
	case nil:
		break
	default:
		panic("unknown outer filter model")
	}

	return &ComposeImageFilter{
		Type:  1,
		Outer: outer,
		Inner: inner,
	}
}

type DilateImageFilter struct {
	Type    uint8   `json:"type"`
	RadiusX float32 `json:"radiusX"`
	RaduisY float32 `json:"radiusY"`
}

func NewDilateImageFilter(radiusX, radiusY float32) *DilateImageFilter {
	return &DilateImageFilter{
		Type:    2,
		RadiusX: radiusX,
		RaduisY: radiusY,
	}
}

type ErodeImageFilter struct {
	Type    uint8   `json:"type"`
	RadiusX float32 `json:"radiusX"`
	RaduisY float32 `json:"radiusY"`
}

func NewErodeImageFilter(radiusX, radiusY float32) *ErodeImageFilter {
	return &ErodeImageFilter{
		Type:    3,
		RadiusX: radiusX,
		RaduisY: radiusY,
	}
}

type MatrixImageFilter struct {
	Type          uint8         `json:"type"`
	Matrix        [16]float64   `json:"matrix4"`
	FilterQuality FilterQuality `json:"filterQuality,omitempty"`
}

func NewMatrixImageFilter(matrix [16]float64, filterQuality FilterQuality) *MatrixImageFilter {
	m := &MatrixImageFilter{
		Type:          4,
		Matrix:        matrix,
		FilterQuality: filterQuality,
	}

	if filterQuality != "" {
		m.FilterQuality = filterQuality
	}

	return m
}

type ImageFilter interface {
	BlurImageFilter | MatrixImageFilter | DilateImageFilter | ErodeImageFilter | ComposeImageFilter
}
