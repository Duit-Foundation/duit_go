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

// SliverListBase defines attributes for SliverListBase widget.
// See: https://api.flutter.dev/flutter/widgets/SliverList/SliverList.html
type SliverListBase struct {
	Type                   duit_utils.Tristate[SliverListKind] `json:"type"`
	AddAutomaticKeepAlives duit_utils.Tristate[bool]           `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries   duit_utils.Tristate[bool]           `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes     duit_utils.Tristate[bool]           `json:"addSemanticIndexes,omitempty"`
}

// SetAddAutomaticKeepAlives sets the addAutomaticKeepAlives for the sliver list base widget.
func (r *SliverListBase) SetAddAutomaticKeepAlives(addAutomaticKeepAlives bool) {
	r.AddAutomaticKeepAlives = duit_utils.TristateFrom[bool](addAutomaticKeepAlives)
}

// SetAddRepaintBoundaries sets the addRepaintBoundaries for the sliver list base widget.
func (r *SliverListBase) SetAddRepaintBoundaries(addRepaintBoundaries bool) {
	r.AddRepaintBoundaries = duit_utils.TristateFrom[bool](addRepaintBoundaries)
}

// SetAddSemanticIndexes sets the addSemanticIndexes for the sliver list base widget.
func (r *SliverListBase) SetAddSemanticIndexes(addSemanticIndexes bool) {
	r.AddSemanticIndexes = duit_utils.TristateFrom[bool](addSemanticIndexes)
}

//bindgen:exclude
func (r *SliverListBase) Validate() error {
	if r.Type == nil {
		return errors.New("type is required")
	}

	return nil
}

// SliverListCommonAttributes defines attributes for SliverListCommon widget.
// See: https://api.flutter.dev/flutter/widgets/SliverList/SliverList.html
type SliverListCommonAttributes struct {
	*SliverListBase `gen:"wrap"`
}

// NewSliverList creates a new SliverListCommonAttributes instance.
func NewSliverList(data *SliverListCommonAttributes) *SliverListAttributes {
	return &SliverListAttributes{data: data}
}

// SliverListBuilderAttributes defines attributes for SliverListBuilder widget.
// See: https://api.flutter.dev/flutter/widgets/SliverList/SliverList.builder.html
type SliverListBuilderAttributes struct {
	*SliverListBase `gen:"wrap"`
	*Builder        `gen:"wrap"`
}

// NewSliverListBuilderAttributes creates a new SliverListBuilderAttributes instance.
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

// SliverListSeparatedAttributes defines attributes for SliverListSeparated widget.
// See: https://api.flutter.dev/flutter/widgets/SliverList/SliverList.separated.html
type SliverListSeparatedAttributes struct {
	*SliverListBase `gen:"wrap"`
	*Builder        `gen:"wrap"`
	Separator       *duit_core.DuitElementModel `json:"separator"`
}

// NewSliverListSeparatedAttributes creates a new SliverListSeparatedAttributes instance.
func NewSliverListSeparatedAttributes(data *SliverListSeparatedAttributes) *SliverListAttributes {
	return &SliverListAttributes{data: data}
}

// SetSeparator sets the separator for the sliver list separated widget.
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

// SliverListAttributes defines attributes for SliverList widget.
// See: https://api.flutter.dev/flutter/widgets/SliverList/SliverList.html
type SliverListAttributes struct {
	data any
}

// NewSliverListAttributes creates a new SliverListAttributes instance.
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
