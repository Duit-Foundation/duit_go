package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
)

// RemoteSubtreeAttributes defines attributes for RemoteSubtree widget.
// See: https://api.flutter.dev/flutter/widgets/RemoteSubtree-class.html
type RemoteSubtreeAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	DownloadPath          string                          `json:"downloadPath"`
	Data                  any                             `json:"data,omitempty"`
	Meta                  any                             `json:"meta,omitempty"`
	Dependencies          []*duit_action.ActionDependency `json:"dependencies,omitempty"`
}

//bindgen:exclude
func (r *RemoteSubtreeAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if len(r.DownloadPath) == 0 {
		return errors.New("downloadPath property is required")
	}

	return nil
}

// NewRemoteSubtreeAttributes creates a new RemoteSubtreeAttributes instance.
func NewRemoteSubtreeAttributes() *RemoteSubtreeAttributes {
	return &RemoteSubtreeAttributes{}
}

// SetDownloadPath sets the download path for the remote subtree widget.
func (r *RemoteSubtreeAttributes) SetDownloadPath(downloadPath string) *RemoteSubtreeAttributes {
	r.DownloadPath = downloadPath
	return r
}

// SetData sets the data for the remote subtree widget.
func (r *RemoteSubtreeAttributes) SetData(data any) *RemoteSubtreeAttributes {
	r.Data = data
	return r
}

// SetMeta sets the meta for the remote subtree widget.
func (r *RemoteSubtreeAttributes) SetMeta(meta any) *RemoteSubtreeAttributes {
	r.Meta = meta
	return r
}

// SetDependencies sets the dependencies for the remote subtree widget.
func (r *RemoteSubtreeAttributes) SetDependencies(dependencies []*duit_action.ActionDependency) *RemoteSubtreeAttributes {
	r.Dependencies = dependencies
	return r
}

// AddDependency adds a single dependency to the remote subtree widget.
func (r *RemoteSubtreeAttributes) AddDependency(dependency *duit_action.ActionDependency) *RemoteSubtreeAttributes {
	r.Dependencies = append(r.Dependencies, dependency)
	return r
}
