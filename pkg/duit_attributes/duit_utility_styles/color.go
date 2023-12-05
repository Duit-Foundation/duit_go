package duit_utility_styles

// Hex string
type ColorString string

// Array of 4 uint8, last item of array will be assigned as Alpha channel
type ColorARGB [4]uint8
type ColorRGBO [4]float32

type Color interface {
	ColorString | ColorARGB | ColorRGBO
}
