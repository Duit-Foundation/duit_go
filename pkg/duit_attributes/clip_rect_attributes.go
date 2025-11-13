package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type ClipRectAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	ClipBehavior          duit_props.Clip `json:"color,omitempty"`
}

//bindgen:exclude
func (r *ClipRectAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	return nil
}

// NewClipRectAttributes creates a new instance of ClipRectAttributes.
func NewClipRectAttributes() *ClipRectAttributes {
	return  &ClipRectAttributes{}
}

// SetClipBehavior sets the clipping behavior for the ClipRect widget.
func (r *ClipRectAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *ClipRectAttributes {
	r.ClipBehavior = clipBehavior
	return r
}
