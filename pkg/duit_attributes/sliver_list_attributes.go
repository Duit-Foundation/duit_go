package duit_attributes

import (
	"encoding/json"
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

type SliverListBase struct {
	Type                   duit_utils.Tristate[SliverListKind] `json:"type"`
	AddAutomaticKeepAlives duit_utils.Tristate[bool]           `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries   duit_utils.Tristate[bool]           `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes     duit_utils.Tristate[bool]           `json:"addSemanticIndexes,omitempty"`
}

func (r *SliverListBase) SetAddAutomaticKeepAlives(addAutomaticKeepAlives bool) {
	r.AddAutomaticKeepAlives = duit_utils.TristateFrom[bool](addAutomaticKeepAlives)
}

func (r *SliverListBase) SetAddRepaintBoundaries(addRepaintBoundaries bool) {
	r.AddRepaintBoundaries = duit_utils.TristateFrom[bool](addRepaintBoundaries)
}

func (r *SliverListBase) SetAddSemanticIndexes(addSemanticIndexes bool) {
	r.AddSemanticIndexes = duit_utils.TristateFrom[bool](addSemanticIndexes)
}

func (r *SliverListBase) Validate() error {
	if r.Type == nil {
		return errors.New("type is required")
	}

	return nil
}

type SliverListCommonAttributes struct {
	*SliverListBase
}

func NewSliverList(data *SliverListCommonAttributes) *SliverListAttributes {
	return &SliverListAttributes{data: data}
}

type SliverListBuilderAttributes struct {
	*SliverListBase
	*Builder
}

func NewSliverListBuilderAttributes(data *SliverListBuilderAttributes) *SliverListAttributes {
	return &SliverListAttributes{data: data}
}

func (r *SliverListBuilderAttributes) Validate() error {
	if r.SliverListBase == nil {
		return errors.New("sliverList is required")
	} else {
		if *r.SliverListBase.Type != SliverListBuilder {
			return errors.New("type must be SliverListBuilder")
		}
	}

	if r.Builder == nil {
		return errors.New("builder is required")
	}

	return nil
}

type SliverListSeparatedAttributes struct {
	*SliverListBase
	*Builder
	Separator *duit_core.DuitElementModel `json:"separator"`
}

func NewSliverListSeparatedAttributes(data *SliverListSeparatedAttributes) *SliverListAttributes {
	return &SliverListAttributes{data: data}
}

func (r *SliverListSeparatedAttributes) SetSeparator(separator *duit_core.DuitElementModel) *SliverListSeparatedAttributes {
	r.Separator = separator
	return r
}

func (r *SliverListSeparatedAttributes) Validate() error {
	if r.SliverListBase == nil {
		return errors.New("sliverList is required")
	} else {
		if *r.SliverListBase.Type != SliverListSeparated {
			return errors.New("type must be SliverListBuilder")
		}
	}

	if r.Separator == nil {
		return errors.New("separator property is required")
	}

	return nil
}

type SliverListAttributes struct {
	data any
}

func NewSliverListAttributes(data any) *SliverListAttributes {
	switch data := data.(type) {
	case *SliverListCommonAttributes:
		return NewSliverList(data)
	case *SliverListBuilderAttributes:
		return NewSliverListBuilderAttributes(data)
	case *SliverListSeparatedAttributes:
		return NewSliverListSeparatedAttributes(data)
	}
	return nil
}

func (r *SliverListAttributes) Validate() error {
	switch data := r.data.(type) {
	case *SliverListCommonAttributes:
		return data.Validate()
	case *SliverListBuilderAttributes:
		return data.Validate()
	case *SliverListSeparatedAttributes:
		return data.Validate()
	}
	return nil
}

func (r *SliverListAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.data)
}
