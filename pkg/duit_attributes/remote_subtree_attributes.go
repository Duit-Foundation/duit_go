package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"

type RemoteSubtreeAttributes struct {
	ValueReferenceHolder
	DownloadPath string `json:"downloadPath"`
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
	Dependencies []*duit_action.ActionDependency `json:"dependencies,omitempty"`
}