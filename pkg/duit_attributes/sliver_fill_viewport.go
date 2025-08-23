package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SliverFillViewportAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	ViewportFraction       float32                       `json:"viewportFraction,omitempty"`
	PadEnds                duit_utils.Tristate[bool]     `json:"padEnds,omitempty"`
	IsBuilderDelegate      duit_utils.Tristate[bool]     `json:"isBuilderDelegate,omitempty"`
	AddAutomaticKeepAlives duit_utils.Tristate[bool]     `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries   duit_utils.Tristate[bool]     `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes     duit_utils.Tristate[bool]     `json:"addSemanticIndexes,omitempty"`
	ChildCount             uint                          `json:"childCount,omitempty"`
	ChildObjects           []*duit_core.DuitElementModel `json:"childObjects,omitempty"`
}

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
