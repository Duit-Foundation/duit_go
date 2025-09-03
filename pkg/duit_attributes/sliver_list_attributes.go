package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SliverListKind uint8

const (
	SliverListCommon SliverListKind = iota
	SliverListBuilder
	SliverListSeparated
)

type SliverList struct {
	Type                   duit_utils.Tristate[SliverListKind] `json:"type"`
	AddAutomaticKeepAlives duit_utils.Tristate[bool]           `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries   duit_utils.Tristate[bool]           `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes     duit_utils.Tristate[bool]           `json:"addSemanticIndexes,omitempty"`
}

func (r *SliverList) Validate() error {
	if r.Type == nil {
		return errors.New("type is required")
	}

	return nil
}

type SliverListBuilderAttributes struct {
	*SliverList
	*Builder
}

func (r *SliverListBuilderAttributes) Validate() error {
	if r.SliverList == nil {
		return errors.New("sliverList is required")
	} else {
		if *r.SliverList.Type != SliverListBuilder {
			return errors.New("type must be SliverListBuilder")
		}
	}

	if r.Builder == nil {
		return errors.New("builder is required")
	}

	return nil
}

type SliverListSeparatedAttributes struct {
	*SliverListBuilderAttributes
	Separator *duit_core.DuitElementModel `json:"separator"`
}

func (r *SliverListSeparatedAttributes) Validate() error {
	if r.SliverListBuilderAttributes == nil {
		return errors.New("sliverList is required")
	} else {
		if *r.SliverList.Type != SliverListBuilder {
			return errors.New("type must be SliverListBuilder")
		}
	}

	if r.Separator == nil {
		return errors.New("separator property is required")
	}

	return nil
}

type SliverListAttributes interface {
	SliverList | SliverListBuilderAttributes | SliverListSeparatedAttributes
}
