package duit_attributes

type ValueRef struct {
	AttributeKey string `json:"attributeKey"`
	ObjectKey    string `json:"objectKey"`
}

type ValueReferenceHolder struct {
	Refs []*ValueRef `json:"refs,omitempty"`
}

func Ref(attributeKey string, objectKey string) *ValueRef {
	return &ValueRef{AttributeKey: attributeKey, ObjectKey: objectKey}
}
