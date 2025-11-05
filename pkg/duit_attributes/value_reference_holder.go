package duit_attributes

import "errors"

type ValueRef struct {
	AttributeKey string `json:"attributeKey"`
	ObjectKey    string `json:"objectKey"`
	DefaultValue any    `json:"defaultValue,omitempty"`
}

func Ref(attributeKey string, objectKey string, defaultValue any) ValueRef {
	return ValueRef{AttributeKey: attributeKey, ObjectKey: objectKey, DefaultValue: defaultValue}
}

type ValueReferenceHolder struct {
	Refs []ValueRef `json:"refs,omitempty"`
}

// SetRefs sets the refs property
func (r *ValueReferenceHolder) SetRefs(refs []ValueRef) {
	r.Refs = refs
}

// AddRef adds a reference
func (r *ValueReferenceHolder) AddRef(ref ValueRef) {
	r.Refs = append(r.Refs, ref)
}

//bindgen:exclude
func (r *ValueReferenceHolder) Validate() error {
	if r == nil {
		return nil
	}

	if len(r.Refs) == 0 {
		return errors.New("refs property must not be empty")
	}
	return nil

}
