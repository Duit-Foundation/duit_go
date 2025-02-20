package duit_attributes

import "github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"

type RemoteSubtreeAttributes struct {
	ValueReferenceHolder
	DownloadPath string `json:"downloadPath"`
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
	Dependencies []*duit_core.ActionDependency `json:"dependencies,omitempty"`
}