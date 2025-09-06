package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type ScaffoldAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	PersistentFooterAlignment    duit_props.Alignment      `json:"persistentFooterAlignment,omitempty"`
	Primary                      duit_utils.Tristate[bool] `json:"primary,omitempty"`
	ExtendBody                   duit_utils.Tristate[bool] `json:"extendBody,omitempty"`
	ResizeToAvoidBottomInset     duit_utils.Tristate[bool] `json:"resizeToAvoidBottomInset,omitempty"`
	BackgroundColor              *duit_props.Color         `json:"backgroundColor,omitempty"`
	RestorationId                string                    `json:"restorationId,omitempty"`
	FloatingActionButtonLocation duit_props.FabLocation    `json:"floatingActionButtonLocation,omitempty"`
}

func (r *ScaffoldAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}
