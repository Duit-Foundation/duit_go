package duit_attributes

type SliverOffstageAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	SliverProps
	Offstage *bool `json:"offstage,omitempty"`
} 