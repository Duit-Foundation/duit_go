package duit_attributes

import "errors"

type ValueRef struct {
	AttributeKey string `json:"attributeKey"`
	ObjectKey    string `json:"objectKey"`
	DefaultValue any    `json:"defaultValue,omitempty"`
}

type ValueReferenceHolder struct {
	Refs []ValueRef `json:"refs,omitempty"`
}

func (r *ValueReferenceHolder) Validate() error {
	if r == nil {
		return nil
	}

	if len(r.Refs) == 0 {
		return errors.New("refs property must not be empty")
	}
	return nil

}

func Ref(attributeKey string, objectKey string, defaultValue any) ValueRef {
	return ValueRef{AttributeKey: attributeKey, ObjectKey: objectKey, DefaultValue: defaultValue}
}
