package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_material"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type ScaffoldAttributes[TColor duit_color.Color] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	PersistentFooterAlignment    duit_alignment.Alignment      `json:"persistentFooterAlignment,omitempty"`
	Primary                      duit_utils.Tristate[bool]     `json:"primary,omitempty"`
	ExtendBody                   duit_utils.Tristate[bool]     `json:"extendBody,omitempty"`
	ResizeToAvoidBottomInset     duit_utils.Tristate[bool]     `json:"resizeToAvoidBottomInset,omitempty"`
	BackgroundColor              TColor                        `json:"backgroundColor,omitempty"`
	RestorationId                string                        `json:"restorationId,omitempty"`
	FloatingActionButtonLocation duit_material.FabLocation     `json:"floatingActionButtonLocation,omitempty"`
}

func (r *ScaffoldAttributes[TColor]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}