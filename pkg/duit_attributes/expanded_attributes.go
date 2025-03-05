package duit_attributes

type ExpandedAttributes struct {
	ValueReferenceHolder
	ThemeConsumer
	Flex uint16 `json:"flex,omitempty"`
}
