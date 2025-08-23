package duit_color

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Hex string
type ColorString string

// Array of 4 uint8, last item of array will be assigned as opacity
type ColorRGBO struct {
	R uint8   `json:"r"`
	G uint8   `json:"g"`
	B uint8   `json:"b"`
	O float32 `json:"o"`
}

func (color *ColorRGBO) MarshalJSON() ([]byte, error) {
	if color != nil {
		return []byte(fmt.Sprintf(`[%d,%d,%d,%f]`, color.R, color.G, color.B, color.O)), nil
	}

	return []byte{}, nil
}

type Color interface {
	ColorString | *ColorRGBO
}

type ColorV2 struct {
	data any
}

func (c ColorV2) FromString(value string) *ColorV2 {
	c.data = value
	return &c
}

func (c ColorV2) FromRGBO(value ColorRGBO) *ColorV2 {
	c.data = [4]uint8{value.R, value.G, value.B, uint8(value.O * 255)}
	return &c
}

func (c ColorV2) FromRGBA(value [4]uint8) *ColorV2 {
	c.data = value
	return &c
}

func (c *ColorV2) Validate() error {
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
		//TODO: добавть валидацию i <= 0 || i >= 255
		return nil
	default:
		return errors.New("invalid color type")
	}
}

func (r *ColorV2) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }
