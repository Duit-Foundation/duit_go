package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type ClipOvalAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	ClipBehavior          duit_props.Clip `json:"clipBehavior,omitempty"`
}

//bindgen:exclude
func (r *ClipOvalAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	return nil
}

// NewClipOvalAttributes creates a new instance of ClipOvalAttributes.
func NewClipOvalAttributes() *ClipOvalAttributes {
	return &ClipOvalAttributes{}
}

// SetClipBehavior sets the clipping behavior for the ClipOval widget.
func (r *ClipOvalAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *ClipOvalAttributes {
	r.ClipBehavior = clipBehavior
	return r
}
