package duit_attributes

import (
	"encoding/json"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SkeletonBoxAttributes defines attributes for SkeletonBox widget.
// Placeholder widget displaying skeleton loading state. Supports width, height and arbitrary extra properties.
type SkeletonBoxAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	Width                 duit_utils.Tristate[float32] `json:"width,omitempty"`
	Height                duit_utils.Tristate[float32] `json:"height,omitempty"`
	Extra                 map[string]any                `json:"-"`
}

// NewSkeletonBoxAttributes creates a new SkeletonBoxAttributes instance.
func NewSkeletonBoxAttributes() *SkeletonBoxAttributes {
	return &SkeletonBoxAttributes{}
}

// SetWidth sets the width property (default 100 on Flutter side when omitted).
func (r *SkeletonBoxAttributes) SetWidth(width float32) *SkeletonBoxAttributes {
	r.Width = duit_utils.Float32Value(width)
	return r
}

// SetHeight sets the height property (default 100 on Flutter side when omitted).
func (r *SkeletonBoxAttributes) SetHeight(height float32) *SkeletonBoxAttributes {
	r.Height = duit_utils.Float32Value(height)
	return r
}

// SetExtra sets additional properties (e.g. "type" for customSkeletonBuilder). Merged into payload at serialization.
func (r *SkeletonBoxAttributes) SetExtra(extra map[string]any) *SkeletonBoxAttributes {
	r.Extra = extra
	return r
}

// AddExtra adds a key-value pair to extra properties.
func (r *SkeletonBoxAttributes) AddExtra(key string, value any) *SkeletonBoxAttributes {
	if r.Extra == nil {
		r.Extra = make(map[string]any)
	}
	r.Extra[key] = value
	return r
}

// MarshalJSON implements custom JSON marshaling to merge Extra into the output payload.
func (r *SkeletonBoxAttributes) MarshalJSON() ([]byte, error) {
	type skeletonBoxAttributesOmitExtra struct {
		*ValueReferenceHolder
		*ThemeConsumer
		Width  duit_utils.Tristate[float32] `json:"width,omitempty"`
		Height duit_utils.Tristate[float32] `json:"height,omitempty"`
	}
	base, err := json.Marshal(&skeletonBoxAttributesOmitExtra{
		ValueReferenceHolder: r.ValueReferenceHolder,
		ThemeConsumer:        r.ThemeConsumer,
		Width:                r.Width,
		Height:               r.Height,
	})
	if err != nil {
		return nil, err
	}
	var m map[string]any
	if err := json.Unmarshal(base, &m); err != nil {
		return nil, err
	}
	for k, v := range r.Extra {
		if v != nil {
			m[k] = v
		}
	}
	return json.Marshal(m)
}

//bindgen:exclude
func (r *SkeletonBoxAttributes) Validate() error {
	if r != nil && r.ValueReferenceHolder != nil {
		if err := r.ValueReferenceHolder.Validate(); err != nil {
			return err
		}
	}
	if r != nil && r.ThemeConsumer != nil {
		if err := r.ThemeConsumer.Validate(); err != nil {
			return err
		}
	}
	return nil
}
