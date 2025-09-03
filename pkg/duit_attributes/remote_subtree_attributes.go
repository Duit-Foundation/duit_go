package duit_attributes

import (
	"errors"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
)

type RemoteSubtreeAttributes struct {
	*ValueReferenceHolder
	DownloadPath string                          `json:"downloadPath"`
	Data         any                             `json:"data,omitempty"`
	Meta         any                             `json:"meta,omitempty"`
	Dependencies []*duit_action.ActionDependency `json:"dependencies,omitempty"`
}

func (r *RemoteSubtreeAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if len(r.DownloadPath) == 0 {
		return errors.New("downloadPath property is required")
	}

	return nil
}