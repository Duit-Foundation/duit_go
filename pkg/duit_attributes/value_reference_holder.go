package duit_attributes

type ValueRef struct {
	AttributeKey string `json:"attributeKey"`
	ObjectKey    string `json:"objectKey"`
	DefaultValue any    `json:"defaultValue,omitempty"`
}

type ValueReferenceHolder struct {
	Refs []*ValueRef `json:"refs,omitempty"`
}

func Ref(attributeKey string, objectKey string, defaultValue any) *ValueRef {
	return &ValueRef{AttributeKey: attributeKey, ObjectKey: objectKey, DefaultValue: defaultValue}
}
