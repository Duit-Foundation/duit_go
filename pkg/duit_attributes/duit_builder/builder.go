package duit_builder

import "github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"

type Builder struct {
	ChildObjects              []*duit_core.DuitElementModel `json:"childObjects,omitempty"`
	ScrollEndReachedThreshold float32                       `json:"scrollEndReachedThreshold,omitempty"`
	MergeStrategy             MergeStrategy                 `json:"mergeStrategy,omitempty"`
}