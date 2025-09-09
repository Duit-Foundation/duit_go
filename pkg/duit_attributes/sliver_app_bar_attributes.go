package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// SliverAppBarAttributes defines attributes for SliverAppBar widget.
// See: https://api.flutter.dev/flutter/material/SliverAppBar-class.html
type SliverAppBarAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	*SliverProps                      `gen:"wrap"`
	TitleTextStyle                    *duit_props.TextStyle        `json:"titleTextStyle,omitempty"`
	ToolbarTextStyle                  *duit_props.TextStyle        `json:"toolbarTextStyle,omitempty"`
	ShadowColor                       *duit_props.Color            `json:"shadowColor,omitempty"`
	BackgroundColor                   *duit_props.Color            `json:"backgroundColor,omitempty"`
	ForegroundColor                   *duit_props.Color            `json:"foregroundColor,omitempty"`
	SurfaceTintColor                  *duit_props.Color            `json:"surfaceTintColor,omitempty"`
	ActionsPadding                    *duit_props.EdgeInsets       `json:"actionsPadding,omitempty"`
	Shape                             *duit_props.ShapeBorder      `json:"shape,omitempty"`
	ClipBehavior                      duit_props.Clip              `json:"clipBehavior,omitempty"`
	TitleSpacing                      float32                      `json:"titleSpacing,omitempty"`
	Elevation                         float32                      `json:"elevation,omitempty"`
	ToolbarHeight                     float32                      `json:"toolbarHeight,omitempty"`
	LeadingWidth                      float32                      `json:"leadingWidth,omitempty"`
	ScrolledUnderElevation            float32                      `json:"scrolledUnderElevation,omitempty"`
	BottomOpacity                     duit_utils.Tristate[float32] `json:"bottomOpacity,omitempty"`
	ToolbarOpacity                    duit_utils.Tristate[float32] `json:"toolbarOpacity,omitempty"`
	ExpandedHeight                    float32                      `json:"expandedHeight,omitempty"`
	CollapsedHeight                   float32                      `json:"collapsedHeight,omitempty"`
	StretchTriggerOffset              float32                      `json:"stretchTriggerOffset,omitempty"`
	ForceElevated                     duit_utils.Tristate[bool]    `json:"forceElevated,omitempty"`
	UseDefaultSemanticsOrder          duit_utils.Tristate[bool]    `json:"useDefaultSemanticsOrder,omitempty"`
	ForceMaterialTransparency         duit_utils.Tristate[bool]    `json:"forceMaterialTransparency,omitempty"`
	CenterTitle                       duit_utils.Tristate[bool]    `json:"centerTitle,omitempty"`
	AutomaticallyImplyLeading         duit_utils.Tristate[bool]    `json:"automaticallyImplyLeading,omitempty"`
	ExcludeHeaderSemantics            duit_utils.Tristate[bool]    `json:"excludeHeaderSemantics,omitempty"`
	Primary                           duit_utils.Tristate[bool]    `json:"primary,omitempty"`
	Floating                          duit_utils.Tristate[bool]    `json:"floating,omitempty"`
	Pinned                            duit_utils.Tristate[bool]    `json:"pinned,omitempty"`
	Snap                              duit_utils.Tristate[bool]    `json:"snap,omitempty"`
	Stretch                           duit_utils.Tristate[bool]    `json:"stretch,omitempty"`
}

//bindgen:exclude
func (r *SliverAppBarAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.BottomOpacity != nil {
		if err := duit_utils.RangeValidator(*r.BottomOpacity, 0, 1); err != nil {
			return err
		}
	}

	if r.ToolbarOpacity != nil {
		if err := duit_utils.RangeValidator(*r.ToolbarOpacity, 0, 1); err != nil {
			return err
		}
	}

	return nil
}

// NewSliverAppBarAttributes creates a new SliverAppBarAttributes instance.
func NewSliverAppBarAttributes() *SliverAppBarAttributes {
	return &SliverAppBarAttributes{}
}

// SetTitleTextStyle sets the title text style for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetTitleTextStyle(titleTextStyle *duit_props.TextStyle) *SliverAppBarAttributes {
	r.TitleTextStyle = titleTextStyle
	return r
}

// SetToolbarTextStyle sets the toolbar text style for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetToolbarTextStyle(toolbarTextStyle *duit_props.TextStyle) *SliverAppBarAttributes {
	r.ToolbarTextStyle = toolbarTextStyle
	return r
}

// SetShadowColor sets the shadow color for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetShadowColor(shadowColor *duit_props.Color) *SliverAppBarAttributes {
	r.ShadowColor = shadowColor
	return r
}

// SetBackgroundColor sets the background color for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetBackgroundColor(backgroundColor *duit_props.Color) *SliverAppBarAttributes {
	r.BackgroundColor = backgroundColor
	return r
}

// SetForegroundColor sets the foreground color for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetForegroundColor(foregroundColor *duit_props.Color) *SliverAppBarAttributes {
	r.ForegroundColor = foregroundColor
	return r
}

// SetSurfaceTintColor sets the surface tint color for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetSurfaceTintColor(surfaceTintColor *duit_props.Color) *SliverAppBarAttributes {
	r.SurfaceTintColor = surfaceTintColor
	return r
}

// SetActionsPadding sets the actions padding for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetActionsPadding(actionsPadding *duit_props.EdgeInsets) *SliverAppBarAttributes {
	r.ActionsPadding = actionsPadding
	return r
}

// SetShape sets the shape for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetShape(shape *duit_props.ShapeBorder) *SliverAppBarAttributes {
	r.Shape = shape
	return r
}

// SetClipBehavior sets the clip behavior for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *SliverAppBarAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetTitleSpacing sets the title spacing for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetTitleSpacing(titleSpacing float32) *SliverAppBarAttributes {
	r.TitleSpacing = titleSpacing
	return r
}

// SetElevation sets the elevation for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetElevation(elevation float32) *SliverAppBarAttributes {
	r.Elevation = elevation
	return r
}

// SetToolbarHeight sets the toolbar height for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetToolbarHeight(toolbarHeight float32) *SliverAppBarAttributes {
	r.ToolbarHeight = toolbarHeight
	return r
}

// SetLeadingWidth sets the leading width for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetLeadingWidth(leadingWidth float32) *SliverAppBarAttributes {
	r.LeadingWidth = leadingWidth
	return r
}

// SetScrolledUnderElevation sets the scrolled under elevation for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetScrolledUnderElevation(scrolledUnderElevation float32) *SliverAppBarAttributes {
	r.ScrolledUnderElevation = scrolledUnderElevation
	return r
}

// SetBottomOpacity sets the bottom opacity for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetBottomOpacity(bottomOpacity duit_utils.Tristate[float32]) *SliverAppBarAttributes {
	r.BottomOpacity = bottomOpacity
	return r
}

// SetToolbarOpacity sets the toolbar opacity for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetToolbarOpacity(toolbarOpacity duit_utils.Tristate[float32]) *SliverAppBarAttributes {
	r.ToolbarOpacity = toolbarOpacity
	return r
}

// SetExpandedHeight sets the expanded height for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetExpandedHeight(expandedHeight float32) *SliverAppBarAttributes {
	r.ExpandedHeight = expandedHeight
	return r
}

// SetCollapsedHeight sets the collapsed height for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetCollapsedHeight(collapsedHeight float32) *SliverAppBarAttributes {
	r.CollapsedHeight = collapsedHeight
	return r
}

// SetStretchTriggerOffset sets the stretch trigger offset for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetStretchTriggerOffset(stretchTriggerOffset float32) *SliverAppBarAttributes {
	r.StretchTriggerOffset = stretchTriggerOffset
	return r
}

// SetForceElevated sets the force elevated for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetForceElevated(forceElevated duit_utils.Tristate[bool]) *SliverAppBarAttributes {
	r.ForceElevated = forceElevated
	return r
}

// SetUseDefaultSemanticsOrder sets the use default semantics order for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetUseDefaultSemanticsOrder(useDefaultSemanticsOrder duit_utils.Tristate[bool]) *SliverAppBarAttributes {
	r.UseDefaultSemanticsOrder = useDefaultSemanticsOrder
	return r
}

// SetForceMaterialTransparency sets the force material transparency for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetForceMaterialTransparency(forceMaterialTransparency duit_utils.Tristate[bool]) *SliverAppBarAttributes {
	r.ForceMaterialTransparency = forceMaterialTransparency
	return r
}

// SetCenterTitle sets the center title for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetCenterTitle(centerTitle duit_utils.Tristate[bool]) *SliverAppBarAttributes {
	r.CenterTitle = centerTitle
	return r
}

// SetAutomaticallyImplyLeading sets the automatically imply leading for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetAutomaticallyImplyLeading(automaticallyImplyLeading duit_utils.Tristate[bool]) *SliverAppBarAttributes {
	r.AutomaticallyImplyLeading = automaticallyImplyLeading
	return r
}

// SetExcludeHeaderSemantics sets the exclude header semantics for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetExcludeHeaderSemantics(excludeHeaderSemantics duit_utils.Tristate[bool]) *SliverAppBarAttributes {
	r.ExcludeHeaderSemantics = excludeHeaderSemantics
	return r
}

// SetPrimary sets the primary for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetPrimary(primary duit_utils.Tristate[bool]) *SliverAppBarAttributes {
	r.Primary = primary
	return r
}

// SetFloating sets the floating for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetFloating(floating duit_utils.Tristate[bool]) *SliverAppBarAttributes {
	r.Floating = floating
	return r
}

// SetPinned sets the pinned for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetPinned(pinned duit_utils.Tristate[bool]) *SliverAppBarAttributes {
	r.Pinned = pinned
	return r
}

// SetSnap sets the snap for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetSnap(snap duit_utils.Tristate[bool]) *SliverAppBarAttributes {
	r.Snap = snap
	return r
}

// SetStretch sets the stretch for the sliver app bar widget.
func (r *SliverAppBarAttributes) SetStretch(stretch duit_utils.Tristate[bool]) *SliverAppBarAttributes {
	r.Stretch = stretch
	return r
}
