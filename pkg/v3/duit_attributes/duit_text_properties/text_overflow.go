package duit_text_properties

type TextOverflow string

const (
	Clip     TextOverflow = "clip"
	Fade     TextOverflow = "fade"
	Ellipsis TextOverflow = "ellipsis"
	Visible  TextOverflow = "visible"
)
