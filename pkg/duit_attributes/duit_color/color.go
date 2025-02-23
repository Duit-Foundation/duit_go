package duit_color

import "fmt"

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
	ColorString |  *ColorRGBO
}
