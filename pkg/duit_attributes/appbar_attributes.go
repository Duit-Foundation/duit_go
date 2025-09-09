package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// AppBarAttributes defines attributes for AppBar widget.
// See: https://api.flutter.dev/flutter/material/AppBar-class.html
type AppBarAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
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
	ForceMaterialTransparency         duit_utils.Tristate[bool]    `json:"forceMaterialTransparency,omitempty"`
	CenterTitle                       duit_utils.Tristate[bool]    `json:"centerTitle,omitempty"`
	AutomaticallyImplyLeading         duit_utils.Tristate[bool]    `json:"automaticallyImplyLeading,omitempty"`
	ExcludeHeaderSemantics            duit_utils.Tristate[bool]    `json:"excludeHeaderSemantics,omitempty"`
	Primary                           duit_utils.Tristate[bool]    `json:"primary,omitempty"`
}

//bindgen:exclude
func (r *AppBarAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
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

	if r.ActionsPadding != nil {
		if err := r.ActionsPadding.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// NewAppBarAttributes creates a new AppBarAttributes instance.
func NewAppBarAttributes() *AppBarAttributes {
	return &AppBarAttributes{}
}

// SetTitleTextStyle sets the title text style for the app bar.
func (r *AppBarAttributes) SetTitleTextStyle(titleTextStyle *duit_props.TextStyle) *AppBarAttributes {
	r.TitleTextStyle = titleTextStyle
	return r
}

// SetToolbarTextStyle sets the toolbar text style for the app bar.
func (r *AppBarAttributes) SetToolbarTextStyle(toolbarTextStyle *duit_props.TextStyle) *AppBarAttributes {
	r.ToolbarTextStyle = toolbarTextStyle
	return r
}

// SetShadowColor sets the shadow color for the app bar.
func (r *AppBarAttributes) SetShadowColor(shadowColor *duit_props.Color) *AppBarAttributes {
	r.ShadowColor = shadowColor
	return r
}

// SetBackgroundColor sets the background color for the app bar.
func (r *AppBarAttributes) SetBackgroundColor(backgroundColor *duit_props.Color) *AppBarAttributes {
	r.BackgroundColor = backgroundColor
	return r
}

// SetForegroundColor sets the foreground color for the app bar.
func (r *AppBarAttributes) SetForegroundColor(foregroundColor *duit_props.Color) *AppBarAttributes {
	r.ForegroundColor = foregroundColor
	return r
}

// SetSurfaceTintColor sets the surface tint color for the app bar.
func (r *AppBarAttributes) SetSurfaceTintColor(surfaceTintColor *duit_props.Color) *AppBarAttributes {
	r.SurfaceTintColor = surfaceTintColor
	return r
}

// SetActionsPadding sets the actions padding for the app bar.
func (r *AppBarAttributes) SetActionsPadding(actionsPadding *duit_props.EdgeInsets) *AppBarAttributes {
	r.ActionsPadding = actionsPadding
	return r
}

// SetShape sets the shape for the app bar.
func (r *AppBarAttributes) SetShape(shape *duit_props.ShapeBorder) *AppBarAttributes {
	r.Shape = shape
	return r
}

// SetClipBehavior sets the clip behavior for the app bar.
func (r *AppBarAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *AppBarAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetTitleSpacing sets the title spacing for the app bar.
func (r *AppBarAttributes) SetTitleSpacing(titleSpacing float32) *AppBarAttributes {
	r.TitleSpacing = titleSpacing
	return r
}

// SetElevation sets the elevation for the app bar.
func (r *AppBarAttributes) SetElevation(elevation float32) *AppBarAttributes {
	r.Elevation = elevation
	return r
}

// SetToolbarHeight sets the toolbar height for the app bar.
func (r *AppBarAttributes) SetToolbarHeight(toolbarHeight float32) *AppBarAttributes {
	r.ToolbarHeight = toolbarHeight
	return r
}

// SetLeadingWidth sets the leading width for the app bar.
func (r *AppBarAttributes) SetLeadingWidth(leadingWidth float32) *AppBarAttributes {
	r.LeadingWidth = leadingWidth
	return r
}

// SetScrolledUnderElevation sets the scrolled under elevation for the app bar.
func (r *AppBarAttributes) SetScrolledUnderElevation(scrolledUnderElevation float32) *AppBarAttributes {
	r.ScrolledUnderElevation = scrolledUnderElevation
	return r
}

// SetBottomOpacity sets the bottom opacity for the app bar.
func (r *AppBarAttributes) SetBottomOpacity(bottomOpacity duit_utils.Tristate[float32]) *AppBarAttributes {
	r.BottomOpacity = bottomOpacity
	return r
}

// SetToolbarOpacity sets the toolbar opacity for the app bar.
func (r *AppBarAttributes) SetToolbarOpacity(toolbarOpacity duit_utils.Tristate[float32]) *AppBarAttributes {
	r.ToolbarOpacity = toolbarOpacity
	return r
}

// SetForceMaterialTransparency sets the force material transparency for the app bar.
func (r *AppBarAttributes) SetForceMaterialTransparency(forceMaterialTransparency duit_utils.Tristate[bool]) *AppBarAttributes {
	r.ForceMaterialTransparency = forceMaterialTransparency
	return r
}

// SetCenterTitle sets the center title for the app bar.
func (r *AppBarAttributes) SetCenterTitle(centerTitle duit_utils.Tristate[bool]) *AppBarAttributes {
	r.CenterTitle = centerTitle
	return r
}

// SetAutomaticallyImplyLeading sets the automatically imply leading for the app bar.
func (r *AppBarAttributes) SetAutomaticallyImplyLeading(automaticallyImplyLeading duit_utils.Tristate[bool]) *AppBarAttributes {
	r.AutomaticallyImplyLeading = automaticallyImplyLeading
	return r
}

// SetExcludeHeaderSemantics sets the exclude header semantics for the app bar.
func (r *AppBarAttributes) SetExcludeHeaderSemantics(excludeHeaderSemantics duit_utils.Tristate[bool]) *AppBarAttributes {
	r.ExcludeHeaderSemantics = excludeHeaderSemantics
	return r
}

// SetPrimary sets the primary for the app bar.
func (r *AppBarAttributes) SetPrimary(primary duit_utils.Tristate[bool]) *AppBarAttributes {
	r.Primary = primary
	return r
}
