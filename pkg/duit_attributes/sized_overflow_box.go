package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"

type SizedOverflowBoxAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Size                              *duit_props.Size     `json:"size,omitempty"`
	Alignment                         duit_props.Alignment `json:"alignment,omitempty"`
}

func NewSizedOverflowBoxAttributes() *SizedOverflowBoxAttributes {
	return &SizedOverflowBoxAttributes{}
}

func (r *SizedOverflowBoxAttributes) SetSize(size *duit_props.Size) *SizedOverflowBoxAttributes {
	r.Size = size
	return r
}

func (r *SizedOverflowBoxAttributes) SetAlignment(alignment duit_props.Alignment) *SizedOverflowBoxAttributes {
	r.Alignment = alignment
	return r
}

func (r *SizedOverflowBoxAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}
