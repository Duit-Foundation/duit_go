package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"

type Builder struct {
	ChildObjects              []*duit_core.DuitElementModel `json:"childObjects,omitempty"`
	ScrollEndReachedThreshold float32                       `json:"scrollEndReachedThreshold,omitempty"`
}
