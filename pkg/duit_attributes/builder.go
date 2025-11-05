package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"

type Builder struct {
	ChildObjects              []*duit_core.DuitElementModel `json:"childObjects,omitempty"`
	ScrollEndReachedThreshold float32                       `json:"scrollEndReachedThreshold,omitempty"`
}

func (r *Builder) SetChildObjects(childObjects []*duit_core.DuitElementModel) *Builder {
	r.ChildObjects = childObjects
	return r
}

func (r *Builder) AddChildObject(childObject *duit_core.DuitElementModel) *Builder {
	r.ChildObjects = append(r.ChildObjects, childObject)
	return r
}

func (r *Builder) SetScrollEndReachedThreshold(scrollEndReachedThreshold float32) *Builder {
	r.ScrollEndReachedThreshold = scrollEndReachedThreshold
	return r
}
