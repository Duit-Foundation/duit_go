package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"

type UnconstrainedBoxAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	Alignment             duit_props.Alignment `json:"alignment,omitempty"`
	ConstrainedAxis       duit_props.Axis      `json:"constrainedAxis,omitempty"`
	ClipBehavior          duit_props.Clip      `json:"clipBehavior,omitempty"`
}

func NewUnconstrainedBoxAttributes() *UnconstrainedBoxAttributes {
	return &UnconstrainedBoxAttributes{}
}

func (r *UnconstrainedBoxAttributes) SetAlignment(alignment duit_props.Alignment) *UnconstrainedBoxAttributes {
	r.Alignment = alignment
	return r
}

func (r *UnconstrainedBoxAttributes) SetConstrainedAxis(constrainedAxis duit_props.Axis) *UnconstrainedBoxAttributes {
	r.ConstrainedAxis = constrainedAxis
	return r
}

func (r *UnconstrainedBoxAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *UnconstrainedBoxAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

func (r *UnconstrainedBoxAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}
