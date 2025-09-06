package duit_props

import (
	"encoding/json"
	"errors"
)

type Color struct {
	data any
}

func NewColorString(value string) *Color {
	c := Color{data: value}
	return &c
}

func NewColorRGBO(rgb [3]uint8, opacity float32) *Color {
	c := Color{data: [4]uint8{rgb[0], rgb[1], rgb[2], uint8(opacity * 255)}}
	return &c
}

func NewColorRGBA(value [4]uint8) *Color {
	c := Color{data: value}
	return &c
}

func (c *Color) Validate() error {
	if c == nil {
		return nil
	}

	switch c.data.(type) {
	case string:
		str := c.data.(string)
		if len(str) != 7 {
			return errors.New("invalid color string")
		}
		if str[0] != '#' {
			return errors.New("invalid color string. Must start with #")
		}
		return nil
	case [4]uint8:
		return nil
	default:
		return errors.New("invalid color type")
	}
}

func (r *Color) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }
