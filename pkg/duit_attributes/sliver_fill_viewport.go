package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SliverFillViewportAttributes defines attributes for SliverFillViewport widget.
// See: https://api.flutter.dev/flutter/widgets/SliverFillViewport-class.html
type SliverFillViewportAttributes struct {
	*ValueReferenceHolder  `gen:"wrap"`
	*ThemeConsumer         `gen:"wrap"`
	ViewportFraction       float32                       `json:"viewportFraction,omitempty"`
	PadEnds                duit_utils.Tristate[bool]     `json:"padEnds,omitempty"`
	IsBuilderDelegate      duit_utils.Tristate[bool]     `json:"isBuilderDelegate,omitempty"`
	AddAutomaticKeepAlives duit_utils.Tristate[bool]     `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries   duit_utils.Tristate[bool]     `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes     duit_utils.Tristate[bool]     `json:"addSemanticIndexes,omitempty"`
	ChildCount             uint                          `json:"childCount,omitempty"`
	ChildObjects           []*duit_core.DuitElementModel `json:"childObjects,omitempty"`
}

//bindgen:exclude
func (r *SliverFillViewportAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.ViewportFraction < 0 {
		return errors.New("viewportFraction must be greater than 0")
	}

	if r.IsBuilderDelegate != nil {
		isBuilder := *r.IsBuilderDelegate

		if isBuilder {
			if len(r.ChildObjects) == 0 {
				return errors.New("childObjects must be non-empty when isBuilderDelegate is true")
			}
		} else {
			if len(r.ChildObjects) != 0 {
				return errors.New("childObjects must be empty when isBuilderDelegate is false")
			}
		}
	} else {
		if len(r.ChildObjects) != 0 {
			return errors.New("childObjects must be empty when isBuilderDelegate is false or nil")
		}
	}

	return nil
}

// NewSliverFillViewportAttributes creates a new SliverFillViewportAttributes instance.
func NewSliverFillViewportAttributes() *SliverFillViewportAttributes {
	return &SliverFillViewportAttributes{}
}

// SetViewportFraction sets the viewport fraction for the sliver fill viewport widget.
func (r *SliverFillViewportAttributes) SetViewportFraction(viewportFraction float32) *SliverFillViewportAttributes {
	r.ViewportFraction = viewportFraction
	return r
}

// SetPadEnds sets the pad ends for the sliver fill viewport widget.
func (r *SliverFillViewportAttributes) SetPadEnds(padEnds duit_utils.Tristate[bool]) *SliverFillViewportAttributes {
	r.PadEnds = padEnds
	return r
}

// SetIsBuilderDelegate sets the is builder delegate for the sliver fill viewport widget.
func (r *SliverFillViewportAttributes) SetIsBuilderDelegate(isBuilderDelegate duit_utils.Tristate[bool]) *SliverFillViewportAttributes {
	r.IsBuilderDelegate = isBuilderDelegate
	return r
}

// SetAddAutomaticKeepAlives sets the add automatic keep alives for the sliver fill viewport widget.
func (r *SliverFillViewportAttributes) SetAddAutomaticKeepAlives(addAutomaticKeepAlives duit_utils.Tristate[bool]) *SliverFillViewportAttributes {
	r.AddAutomaticKeepAlives = addAutomaticKeepAlives
	return r
}

// SetAddRepaintBoundaries sets the add repaint boundaries for the sliver fill viewport widget.
func (r *SliverFillViewportAttributes) SetAddRepaintBoundaries(addRepaintBoundaries duit_utils.Tristate[bool]) *SliverFillViewportAttributes {
	r.AddRepaintBoundaries = addRepaintBoundaries
	return r
}

// SetAddSemanticIndexes sets the add semantic indexes for the sliver fill viewport widget.
func (r *SliverFillViewportAttributes) SetAddSemanticIndexes(addSemanticIndexes duit_utils.Tristate[bool]) *SliverFillViewportAttributes {
	r.AddSemanticIndexes = addSemanticIndexes
	return r
}

// SetChildCount sets the child count for the sliver fill viewport widget.
func (r *SliverFillViewportAttributes) SetChildCount(childCount uint) *SliverFillViewportAttributes {
	r.ChildCount = childCount
	return r
}

// SetChildObjects sets the child objects for the sliver fill viewport widget.
func (r *SliverFillViewportAttributes) SetChildObjects(childObjects []*duit_core.DuitElementModel) *SliverFillViewportAttributes {
	r.ChildObjects = childObjects
	return r
}

// AddChildObject adds a single child object to the sliver fill viewport widget.
func (r *SliverFillViewportAttributes) AddChildObject(childObject *duit_core.DuitElementModel) *SliverFillViewportAttributes {
	r.ChildObjects = append(r.ChildObjects, childObject)
	return r
}
