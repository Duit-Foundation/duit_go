package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type ListKind uint8

const (
	ListCommon ListKind = iota
	ListBuilder
	ListSeparated
)

type ListViewBaseAttributes[TInsets duit_props.EdgeInsets] struct {
	Type                    duit_utils.Tristate[ListKind]                   `json:"type,omitempty"`
	Reverse                 duit_utils.Tristate[bool]                       `json:"reverse,omitempty"`
	Primary                 duit_utils.Tristate[bool]                       `json:"primary,omitempty"`
	ShrinkWrap              duit_utils.Tristate[bool]                       `json:"shrinkWrap,omitempty"`
	AddAutomaticKeepAlives  duit_utils.Tristate[bool]                       `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries    duit_utils.Tristate[bool]                       `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes      duit_utils.Tristate[bool]                       `json:"addSemanticIndexes,omitempty"`
	ScrollDirection         duit_props.Axis                                  `json:"scrollDirection,omitempty"`
	ScrollPhysics           duit_props.ScrollPhysics                     `json:"scrollPhysics,omitempty"`
	CacheExtent             float32                                         `json:"cacheExtent,omitempty"`
	Anchor                  float32                                         `json:"anchors,omitempty"`
	SemantickChildCount     int                                             `json:"semanticChildCount,omitempty"`
	Padding                 *TInsets                                        `json:"padding,omitempty"`
	ItemExtent              float32                                         `json:"itemExtent,omitempty"`
	ClipBehavior            duit_props.Clip                                  `json:"clipBehavior,omitempty"`
	RestorationId           string                                          `json:"restorationId,omitempty"`
	DragStarnBehavior       duit_props.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	KeyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
}

func (r *ListViewBaseAttributes[TInsets]) Validate() error {
	if r.Type != nil {

	} else {
		return errors.New("type is required")
	}

	return nil
}

type ListViewBuilderAttributes[TInsets duit_props.EdgeInsets] struct {
	*ListViewBaseAttributes[TInsets]
	*Builder
}

func (r *ListViewBuilderAttributes[TInsets]) Validate() error {
	if r.ListViewBaseAttributes != nil {
		if err := r.ListViewBaseAttributes.Validate(); err != nil {
			return err
		}

	} else {
		return errors.New("listView property is required")
	}

	if r.Builder == nil {
		return errors.New("builder property is required")
	}

	return nil
}

type ListViewSeparatedAttributes[TInsets duit_props.EdgeInsets] struct {
	*ListViewBuilderAttributes[TInsets]
	Separator *duit_core.DuitElementModel `json:"separator"`
}

func (r *ListViewSeparatedAttributes[TInsets]) Validate() error {
	if r.ListViewBuilderAttributes != nil {
		if err := r.ListViewBuilderAttributes.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("listViewBuilderAttributes property is required")
	}

	if r.Separator == nil {
		return errors.New("separator property is required")
	}

	return nil
}

type ListViewAttributes[TInsets duit_props.EdgeInsets] interface {
	ListViewBaseAttributes[TInsets] | ListViewBuilderAttributes[TInsets] | ListViewSeparatedAttributes[TInsets]
}
