package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_builder"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type SliverGridConstructor uint

const (
	SliverGridCommon SliverGridConstructor = iota
	SliverGridCount
	SliverGridBuilder
	SliverGridExtent
)

type DefaultSliverGridAttributes struct {
	Constructor            duit_utils.Tristate[SliverGridConstructor] `json:"constructor"`
	AddAutomaticKeepAlives *bool                                      `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries   *bool                                      `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes     *bool                                      `json:"addSemanticIndexes,omitempty"`
	CacheExtent            float32                                    `json:"cacheExtent,omitempty"`
	SemanticChildCount     int                                        `json:"semanticChildCount,omitempty"`
	ClipBehavior           duit_props.Clip                             `json:"clipBehavior,omitempty"`
}

func (r *DefaultSliverGridAttributes) Validate() error {
	if r.Constructor == nil {
		return errors.New("constructor is required")
	}

	return nil
}

type SliverGridCommonAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*DefaultSliverGridAttributes
	SliverGridDelegateKey     string         `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions map[string]any `json:"sliverGridDelegateOptions,omitempty"`
}

func (r *SliverGridCommonAttributes) Validate() error {
	// Validate embedded components
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.DefaultSliverGridAttributes == nil {
		return errors.New("defaultSliverGridAttributes is required")
	} else {
		if err := r.DefaultSliverGridAttributes.Validate(); err != nil {
			return err
		}

		// Validate constructor
		if *r.DefaultSliverGridAttributes.Constructor != SliverGridCommon {
			return errors.New("constructor must be SliverGridCommon")
		}
	}

	// Validate grid delegate configuration
	if len(r.SliverGridDelegateKey) == 0 {
		return errors.New("SliverGridDelegateKey is required")
	}

	return nil
}

type SliverGridCountAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*DefaultSliverGridAttributes
	CrossAxisCount   int     `json:"crossAxisCount"`
	MainAxisSpacing  float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio float32 `json:"childAspectRatio,omitempty"`
}

func (r *SliverGridCountAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.DefaultSliverGridAttributes == nil {
		return errors.New("defaultSliverGridAttributes is required")
	} else {
		if err := r.DefaultSliverGridAttributes.Validate(); err != nil {
			return err
		}

		// Validate constructor
		if *r.DefaultSliverGridAttributes.Constructor != SliverGridCount {
			return errors.New("constructor must be SliverGridCount")
		}
	}

	if r.CrossAxisCount == 0 {
		return errors.New("crossAxisCount is required and must be greater than 0")
	}

	return nil
}

type SliverGridBuilderAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*duit_builder.Builder
	*DefaultSliverGridAttributes
	SliverGridDelegateKey     string                 `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions map[string]interface{} `json:"sliverGridDelegateOptions,omitempty"`
}

func (r *SliverGridBuilderAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.DefaultSliverGridAttributes == nil {
		return errors.New("defaultSliverGridAttributes is required")
	} else {
		if err := r.DefaultSliverGridAttributes.Validate(); err != nil {
			return err
		}

		// Validate constructor
		if *r.DefaultSliverGridAttributes.Constructor != SliverGridBuilder {
			return errors.New("constructor must be SliverGridBuilder")
		}
	}

	if r.Builder == nil {
		return errors.New("builder is required")
	}

	if len(r.SliverGridDelegateKey) == 0 {
		return errors.New("sliverGridDelegateKey is required")
	}

	return nil
}

type SliverGridExtentAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*DefaultSliverGridAttributes
	MaxCrossAxisExtent float32 `json:"maxCrossAxisExtent"`
	MainAxisSpacing    float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing   float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio   float32 `json:"childAspectRatio,omitempty"`
}

func (r *SliverGridExtentAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}
	if r.DefaultSliverGridAttributes == nil {
		return errors.New("defaultSliverGridAttributes is required")
	} else {
		if err := r.DefaultSliverGridAttributes.Validate(); err != nil {
			return err
		}

		// Validate constructor
		if *r.DefaultSliverGridAttributes.Constructor != SliverGridExtent {
			return errors.New("constructor must be SliverGridExtent")
		}
	}

	if r.MaxCrossAxisExtent == 0 {
		return errors.New("MaxCrossAxisExtent is required and must be greater than 0")
	}

	return nil
}

type SliverGridAttributes interface {
	SliverGridCommonAttributes | SliverGridCountAttributes | SliverGridBuilderAttributes | SliverGridExtentAttributes
}
