package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type GridConstructor uint

const (
	GridCommon GridConstructor = iota
	GridCount
	GridBuilder
	GridExtent
)

type DefautlGridViewAttributes[TInsets duit_props.EdgeInsets] struct {
	Constructor             duit_utils.Tristate[GridConstructor]            `json:"constructor"`
	ScrollDirection         duit_props.Axis                                  `json:"scrollDirection,omitempty"`
	Reverse                 duit_utils.Tristate[bool]                       `json:"reverse,omitempty"`
	Primary                 duit_utils.Tristate[bool]                       `json:"primary,omitempty"`
	ShrinkWrap              duit_utils.Tristate[bool]                       `json:"shrinkWrap,omitempty"`
	AddAutomaticKeepAlives  duit_utils.Tristate[bool]                       `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries    duit_utils.Tristate[bool]                       `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes      duit_utils.Tristate[bool]                       `json:"addSemanticIndexes,omitempty"`
	Padding                 TInsets                                         `json:"padding,omitempty"`
	CacheExtent             float32                                         `json:"cacheExtent,omitempty"`
	SemanticChildCount      int                                             `json:"semanticChildCount,omitempty"`
	DragStartBehavior       duit_props.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	ClipBehavior            duit_props.Clip                                  `json:"clipBehavior,omitempty"`
	KeyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
	Physics                 duit_props.ScrollPhysics                     `json:"physics,omitempty"`
	RestorationId           string                                          `json:"restorationId,omitempty"`
}

type GridViewCommonAttributes[TInsets duit_props.EdgeInsets] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*DefautlGridViewAttributes[TInsets]
	SliverGridDelegateKey     string         `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions map[string]any `json:"sliverGridDelegateOptions,omitempty"`
}

func (r *GridViewCommonAttributes[TInsets]) Validate() error {
	// Validate embedded components
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.DefautlGridViewAttributes != nil {
		if r.DefautlGridViewAttributes.Constructor == nil {
			return errors.New("DefautlGridViewAttributes.Constructor is required")
		}

		if *r.DefautlGridViewAttributes.Constructor != GridCommon {
			return errors.New("DefautlGridViewAttributes.Constructor must be GridCommon")
		}
	} else {
		return errors.New("DefautlGridViewAttributes is required")
	}

	// Validate grid delegate configuration
	if len(r.SliverGridDelegateKey) == 0 {
		return errors.New("SliverGridDelegateKey is required")
	}

	return nil
}

type GridViewCountAttributes[TInsets duit_props.EdgeInsets] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*DefautlGridViewAttributes[TInsets]
	CrossAxisCount   uint    `json:"crossAxisCount"`
	MainAxisSpacing  float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio float32 `json:"childAspectRatio,omitempty"`
}

func (r *GridViewCountAttributes[TInsets]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.DefautlGridViewAttributes != nil {
		if r.DefautlGridViewAttributes.Constructor == nil {
			return errors.New("DefautlGridViewAttributes.Constructor is required")
		}

		if *r.DefautlGridViewAttributes.Constructor != GridCount {
			return errors.New("DefautlGridViewAttributes.Constructor must be GridCount")
		}
	} else {
		return errors.New("DefautlGridViewAttributes is required")
	}

	if r.CrossAxisCount == 0 {
		return errors.New("crossAxisCount is required and must be greater than 0")
	}

	return nil
}

type GridViewBuilderAttributes[TInsets duit_props.EdgeInsets] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*Builder
	*DefautlGridViewAttributes[TInsets]
	SliverGridDelegateKey     string         `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions map[string]any `json:"sliverGridDelegateOptions,omitempty"`
}

func (r *GridViewBuilderAttributes[TInsets]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.DefautlGridViewAttributes != nil {
		if r.DefautlGridViewAttributes.Constructor == nil {
			return errors.New("DefautlGridViewAttributes.Constructor is required")
		}

		if *r.DefautlGridViewAttributes.Constructor != GridBuilder {
			return errors.New("DefautlGridViewAttributes.Constructor must be GridBuilder")
		}
	} else {
		return errors.New("DefautlGridViewAttributes is required")
	}

	if r.Builder == nil {
		return errors.New("builder is required")
	}

	if len(r.SliverGridDelegateKey) == 0 {
		return errors.New("sliverGridDelegateKey is required")
	}

	return nil
}

type GridViewExtentAttributes[TInsets duit_props.EdgeInsets] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*DefautlGridViewAttributes[TInsets]
	MaxCrossAxisExtent float32 `json:"maxCrossAxisExtent"`
	MainAxisSpacing    float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing   float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio   float32 `json:"childAspectRatio,omitempty"`
}

func (r *GridViewExtentAttributes[TInsets]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.DefautlGridViewAttributes != nil {
		if r.DefautlGridViewAttributes.Constructor == nil {
			return errors.New("DefautlGridViewAttributes.Constructor is required")
		}

		if *r.DefautlGridViewAttributes.Constructor != GridExtent {
			return errors.New("DefautlGridViewAttributes.Constructor must be GridExtent")
		}
	} else {
		return errors.New("DefautlGridViewAttributes is required")
	}

	if r.MaxCrossAxisExtent == 0 {
		return errors.New("MaxCrossAxisExtent is required and must be greater than 0")
	}

	return nil
}

type GridViewAttributes[TInsets duit_props.EdgeInsets] interface {
	GridViewCommonAttributes[TInsets] | GridViewCountAttributes[TInsets] | GridViewBuilderAttributes[TInsets] | GridViewExtentAttributes[TInsets]
}
