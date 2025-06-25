package duit_attributes

type SliverFillRemainingAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	SliverProps
	HasScrollBody *bool `json:"hasScrollBody,omitempty"`
	FillOverscroll *bool `json:"fillOverscroll,omitempty"`
}
