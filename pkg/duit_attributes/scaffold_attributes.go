package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// ScaffoldAttributes defines attributes for Scaffold widget.
// See: https://api.flutter.dev/flutter/material/Scaffold-class.html
type ScaffoldAttributes struct {
	*ValueReferenceHolder        `gen:"wrap"`
	*ThemeConsumer               `gen:"wrap"`
	PersistentFooterAlignment    duit_props.Alignment      `json:"persistentFooterAlignment,omitempty"`
	Primary                      duit_utils.Tristate[bool] `json:"primary,omitempty"`
	ExtendBody                   duit_utils.Tristate[bool] `json:"extendBody,omitempty"`
	ResizeToAvoidBottomInset     duit_utils.Tristate[bool] `json:"resizeToAvoidBottomInset,omitempty"`
	BackgroundColor              *duit_props.Color         `json:"backgroundColor,omitempty"`
	RestorationId                string                    `json:"restorationId,omitempty"`
	FloatingActionButtonLocation duit_props.FabLocation    `json:"floatingActionButtonLocation,omitempty"`
}

//bindgen:exclude
func (r *ScaffoldAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}

// NewScaffoldAttributes creates a new ScaffoldAttributes instance.
func NewScaffoldAttributes() *ScaffoldAttributes {
	return &ScaffoldAttributes{}
}

// SetPersistentFooterAlignment sets the persistent footer alignment for the scaffold widget.
func (r *ScaffoldAttributes) SetPersistentFooterAlignment(persistentFooterAlignment duit_props.Alignment) *ScaffoldAttributes {
	r.PersistentFooterAlignment = persistentFooterAlignment
	return r
}

// SetPrimary sets the primary for the scaffold widget.
func (r *ScaffoldAttributes) SetPrimary(primary duit_utils.Tristate[bool]) *ScaffoldAttributes {
	r.Primary = primary
	return r
}

// SetExtendBody sets the extend body for the scaffold widget.
func (r *ScaffoldAttributes) SetExtendBody(extendBody duit_utils.Tristate[bool]) *ScaffoldAttributes {
	r.ExtendBody = extendBody
	return r
}

// SetResizeToAvoidBottomInset sets the resize to avoid bottom inset for the scaffold widget.
func (r *ScaffoldAttributes) SetResizeToAvoidBottomInset(resizeToAvoidBottomInset duit_utils.Tristate[bool]) *ScaffoldAttributes {
	r.ResizeToAvoidBottomInset = resizeToAvoidBottomInset
	return r
}

// SetBackgroundColor sets the background color for the scaffold widget.
func (r *ScaffoldAttributes) SetBackgroundColor(backgroundColor *duit_props.Color) *ScaffoldAttributes {
	r.BackgroundColor = backgroundColor
	return r
}

// SetRestorationId sets the restoration id for the scaffold widget.
func (r *ScaffoldAttributes) SetRestorationId(restorationId string) *ScaffoldAttributes {
	r.RestorationId = restorationId
	return r
}

// SetFloatingActionButtonLocation sets the floating action button location for the scaffold widget.
func (r *ScaffoldAttributes) SetFloatingActionButtonLocation(floatingActionButtonLocation duit_props.FabLocation) *ScaffoldAttributes {
	r.FloatingActionButtonLocation = floatingActionButtonLocation
	return r
}
