package duit_attributes

type FlexAttributes struct {
	MainAxisAlignment  MainAxisAlignment  `json:"mainAxisAlignment,omitempty"`
	MainAxisSize       MainAxisSize       `json:"mainAxisSize,omitempty"`
	CrossAxisAlignment CrossAxisAlignment `json:"crossAxisAlignment,omitempty"`
	TextDirection      TextDirection      `json:"textDirection,omitempty"`
	VerticalDirection  VerticalDirection  `json:"verticalDirection,omitempty"`
	ClipBehavior       Clip               `json:"clipBehavior,omitempty"`
}
