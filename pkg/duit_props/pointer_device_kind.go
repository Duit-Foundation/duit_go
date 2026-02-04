package duit_props

type PointerDeviceKind = string

const (
	PointerDeviceKindMouse          PointerDeviceKind = "mouse"
	PointerDeviceKindTouch          PointerDeviceKind = "touch"
	PointerDeviceKindTrackpad       PointerDeviceKind = "trackpad"
	PointerDeviceKindStylus         PointerDeviceKind = "stylus"
	PointerDeviceKindInvertedStylus PointerDeviceKind = "invertedStylus"
	PointerDeviceKindUnknown        PointerDeviceKind = "unknown"
)
