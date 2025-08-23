package duit_decoration

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Radius struct {
	Radius float32 `json:"radius"`
}

type RadiusElliptical [2]float32

type RadiusCircular float32

type BorderRadius struct {
	TopLeft     *Radius `json:"topLeft,omitempty"`
	TopRight    *Radius `json:"topRight,omitempty"`
	BottomLeft  *Radius `json:"bottomLeft,omitempty"`
	BottomRight *Radius `json:"bottomRight,omitempty"`
}

type BorderRadiusV2 struct {
	data any
}

type borderRadiusAll struct {
	Radius *RadiusV2 `json:"radius"`
}

type borderRadiusOnly struct {
	TopLeft     *RadiusV2 `json:"topLeft,omitempty"`
	TopRight    *RadiusV2 `json:"topRight,omitempty"`
	BottomLeft  *RadiusV2 `json:"bottomLeft,omitempty"`
	BottomRight *RadiusV2 `json:"bottomRight,omitempty"`
}

type borderRadiusVertical struct {
	Top    *RadiusV2 `json:"top,omitempty"`
	Bottom *RadiusV2 `json:"bottom,omitempty"`
}

type borderRadiusHorizontal struct {
	Left  *RadiusV2 `json:"left,omitempty"`
	Right *RadiusV2 `json:"right,omitempty"`
}



func (r BorderRadiusV2) All(radius *RadiusV2) *BorderRadiusV2 {
	r.data = borderRadiusAll{Radius: radius}
	return &r
}

func (r BorderRadiusV2) Only(topLeft, topRight, bottomLeft, bottomRight *RadiusV2) *BorderRadiusV2 {
	r.data = borderRadiusOnly{
		TopLeft:     topLeft,
		TopRight:    topRight,
		BottomLeft:  bottomLeft,
		BottomRight: bottomRight,
	}
	return &r
}

func (r BorderRadiusV2) Vertical(top, bottom *RadiusV2) *BorderRadiusV2 {
	r.data = borderRadiusVertical{
		Top:    top,
		Bottom: bottom,
	}
	return &r
}

func (r BorderRadiusV2) Horizontal(left, right *RadiusV2) *BorderRadiusV2 {
	r.data = borderRadiusHorizontal{
		Left:  left,
		Right: right,
	}
	return &r
}

func (r *BorderRadiusV2) Validate() error {
	if r == nil {
		return nil
	}

	switch data := r.data.(type) {
	case borderRadiusAll:
		if data.Radius != nil {
			return data.Radius.Validate()
		}
	case borderRadiusOnly:
		if data.TopLeft != nil {
			if err := data.TopLeft.Validate(); err != nil {
				return err
			}
		}
		if data.TopRight != nil {
			if err := data.TopRight.Validate(); err != nil {
				return err
			}
		}
		if data.BottomLeft != nil {
			if err := data.BottomLeft.Validate(); err != nil {
				return err
			}
		}
		if data.BottomRight != nil {
			if err := data.BottomRight.Validate(); err != nil {
				return err
			}
		}
	case borderRadiusVertical:
		if data.Top != nil {
			if err := data.Top.Validate(); err != nil {
				return err
			}
		}
		if data.Bottom != nil {
			if err := data.Bottom.Validate(); err != nil {
				return err
			}
		}
	case borderRadiusHorizontal:
		if data.Left != nil {
			if err := data.Left.Validate(); err != nil {
				return err
			}
		}
		if data.Right != nil {
			if err := data.Right.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *BorderRadiusV2) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }

type RadiusV2 struct {
	data any
}

func (r RadiusV2) Elliptical(value [2]float32) *RadiusV2 {
	r.data = RadiusElliptical(value)
	return &r
}

func (r RadiusV2) Circular(value float32) *RadiusV2 {
	r.data = RadiusCircular(value)
	return &r
}

func (r *RadiusV2) Validate() error {
	if r == nil {
		return nil
	}

	switch r.data.(type) {
	case RadiusElliptical:
		for i, v := range r.data.(RadiusElliptical) {
			if v < 0 {
				return errors.New("radius value at index " + strconv.Itoa(i) + " must be greater than 0")
			}
		}

		return nil
	case RadiusCircular:
		if r.data.(RadiusCircular) < 0 {
			return errors.New("radius must be greater than 0")
		}

		return nil
	}

	return nil
}

func (r *RadiusV2) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }
