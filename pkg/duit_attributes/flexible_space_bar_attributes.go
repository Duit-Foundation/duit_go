package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// FlexibleSpaceBarAttributes defines attributes for FlexibleSpaceBar widget.
// See: https://api.flutter.dev/flutter/material/FlexibleSpaceBar-class.html
type FlexibleSpaceBarAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	TitlePadding          *duit_props.EdgeInsets    `json:"titlePadding,omitempty"`
	CollapseMode          duit_props.CollapseMode   `json:"collapseMode,omitempty"`
	StretchModes          []duit_props.StretchMode  `json:"stretchModes,omitempty"`
	CenterTitle           duit_utils.Tristate[bool] `json:"centerTitle,omitempty"`
	ExpandedTitleScale    float32                   `json:"expandedTitleScale,omitempty"`
}

//bindgen:exclude
func (r *FlexibleSpaceBarAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	return nil
}

// NewFlexibleSpaceBarAttributes creates a new FlexibleSpaceBarAttributes instance.
func NewFlexibleSpaceBarAttributes() *FlexibleSpaceBarAttributes {
	return &FlexibleSpaceBarAttributes{}
}

// SetTitlePadding sets the title padding for the flexible space bar widget.
func (r *FlexibleSpaceBarAttributes) SetTitlePadding(titlePadding *duit_props.EdgeInsets) *FlexibleSpaceBarAttributes {
	r.TitlePadding = titlePadding
	return r
}

// SetCollapseMode sets the collapse mode for the flexible space bar widget.
func (r *FlexibleSpaceBarAttributes) SetCollapseMode(collapseMode duit_props.CollapseMode) *FlexibleSpaceBarAttributes {
	r.CollapseMode = collapseMode
	return r
}

// SetStretchModes sets the stretch modes for the flexible space bar widget.
func (r *FlexibleSpaceBarAttributes) SetStretchModes(stretchModes []duit_props.StretchMode) *FlexibleSpaceBarAttributes {
	r.StretchModes = stretchModes
	return r
}

// AddStretchMode adds a single stretch mode to the flexible space bar widget.
func (r *FlexibleSpaceBarAttributes) AddStretchMode(stretchMode duit_props.StretchMode) *FlexibleSpaceBarAttributes {
	r.StretchModes = append(r.StretchModes, stretchMode)
	return r
}

// SetCenterTitle sets the center title for the flexible space bar widget.
func (r *FlexibleSpaceBarAttributes) SetCenterTitle(centerTitle duit_utils.Tristate[bool]) *FlexibleSpaceBarAttributes {
	r.CenterTitle = centerTitle
	return r
}

// SetExpandedTitleScale sets the expanded title scale for the flexible space bar widget.
func (r *FlexibleSpaceBarAttributes) SetExpandedTitleScale(expandedTitleScale float32) *FlexibleSpaceBarAttributes {
	r.ExpandedTitleScale = expandedTitleScale
	return r
}
