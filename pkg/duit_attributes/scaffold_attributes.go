package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_material"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
)

type ScaffoldAttributes[TColor duit_color.Color] struct {
	ValueReferenceHolder
	ThemeConsumer
	AppBar                    *duit_core.DuitElementModel  `json:"appBar,omitempty"`
	BottomNavigationBar       *duit_core.DuitElementModel  `json:"bottomNavigationBar,omitempty"`
	BottomSheet               *duit_core.DuitElementModel  `json:"bottomSheet,omitempty"`
	FloatingActionButton      *duit_core.DuitElementModel  `json:"floatingActionButton,omitempty"`
	PersistentFooterButtons   []duit_core.DuitElementModel `json:"persistentFooterButtons,omitempty"`
	PersistentFooterAlignment duit_alignment.Alignment     `json:"persistentFooterAlignment,omitempty"`
	Primary                   *bool                        `json:"primary,omitempty"`
	ExtendBody                *bool                        `json:"extendBody,omitempty"`
	ResizeToAvoidBottomInset  *bool                        `json:"resizeToAvoidBottomInset,omitempty"`
	BackgroundColor           TColor                       `json:"backgroundColor,omitempty"`
	RestorationId             *string                      `json:"restorationId,omitempty"`
	FloatingActionButtonLocation duit_material.FabLocation `json:"floatingActionButtonLocation,omitempty"`
}
