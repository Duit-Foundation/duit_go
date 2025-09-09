package duit_attributes

import (
	"encoding/json"
	"errors"

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

// DefaultSliverGridAttributes defines attributes for DefaultSliverGrid widget.
// See: https://api.flutter.dev/flutter/widgets/SliverGrid/SliverGrid.html
type DefaultSliverGridAttributes struct {
	Constructor            duit_utils.Tristate[SliverGridConstructor] `json:"constructor"`
	AddAutomaticKeepAlives *bool                                      `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries   *bool                                      `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes     *bool                                      `json:"addSemanticIndexes,omitempty"`
	CacheExtent            float32                                    `json:"cacheExtent,omitempty"`
	SemanticChildCount     int                                        `json:"semanticChildCount,omitempty"`
	ClipBehavior           duit_props.Clip                            `json:"clipBehavior,omitempty"`
}

// SetConstructor sets the constructor for the sliver grid widget.
func (r *DefaultSliverGridAttributes) SetConstructor(constructor SliverGridConstructor) {
	r.Constructor = duit_utils.TristateFrom[SliverGridConstructor](constructor)
}

// SetAddAutomaticKeepAlives sets the addAutomaticKeepAlives for the sliver grid widget.
func (r *DefaultSliverGridAttributes) SetAddAutomaticKeepAlives(addAutomaticKeepAlives bool) {
	r.AddAutomaticKeepAlives = &addAutomaticKeepAlives
}

// SetAddRepaintBoundaries sets the addRepaintBoundaries for the sliver grid widget.
func (r *DefaultSliverGridAttributes) SetAddRepaintBoundaries(addRepaintBoundaries bool) {
	r.AddRepaintBoundaries = &addRepaintBoundaries
}

// SetAddSemanticIndexes sets the addSemanticIndexes for the sliver grid widget.
func (r *DefaultSliverGridAttributes) SetAddSemanticIndexes(addSemanticIndexes bool) {
	r.AddSemanticIndexes = &addSemanticIndexes
}

// SetCacheExtent sets the cacheExtent for the sliver grid widget.
func (r *DefaultSliverGridAttributes) SetCacheExtent(cacheExtent float32) {
	r.CacheExtent = cacheExtent
}

// SetSemanticChildCount sets the semanticChildCount for the sliver grid widget.
func (r *DefaultSliverGridAttributes) SetSemanticChildCount(semanticChildCount int) {
	r.SemanticChildCount = semanticChildCount
}

// SetClipBehavior sets the clipBehavior for the sliver grid widget.
func (r *DefaultSliverGridAttributes) SetClipBehavior(clipBehavior duit_props.Clip) {
	r.ClipBehavior = clipBehavior
}

func (r *DefaultSliverGridAttributes) Validate() error {
	if r.Constructor == nil {
		return errors.New("constructor is required")
	}

	return nil
}

// SliverGridCommonAttributes defines attributes for SliverGridCommon widget.
// See: https://api.flutter.dev/flutter/widgets/SliverGrid/SliverGrid.html
type SliverGridCommonAttributes struct {
	*ValueReferenceHolder        `gen:"wrap"`
	*ThemeConsumer               `gen:"wrap"`
	*DefaultSliverGridAttributes `gen:"wrap"`
	SliverGridDelegateKey        string         `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions    map[string]any `json:"sliverGridDelegateOptions,omitempty"`
}

// NewSliverGridCommonAttributes creates a new SliverGridCommonAttributes instance.
func NewSliverGridCommonAttributes(data *SliverGridCommonAttributes) *SliverGridAttributes {
	data.DefaultSliverGridAttributes.Constructor = duit_utils.TristateFrom[SliverGridConstructor](SliverGridCommon)
	return &SliverGridAttributes{data: data}
}

// SetSliverGridDelegateKey sets the sliverGridDelegateKey for the sliver grid common widget.
func (r *SliverGridCommonAttributes) SetSliverGridDelegateKey(sliverGridDelegateKey string) *SliverGridCommonAttributes {
	r.SliverGridDelegateKey = sliverGridDelegateKey
	return r
}

// SetSliverGridDelegateOptions sets the sliverGridDelegateOptions for the sliver grid common widget.
func (r *SliverGridCommonAttributes) SetSliverGridDelegateOptions(sliverGridDelegateOptions map[string]any) *SliverGridCommonAttributes {
	r.SliverGridDelegateOptions = sliverGridDelegateOptions
	return r
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

// SliverGridCountAttributes defines attributes for SliverGridCount widget.
// See: https://api.flutter.dev/flutter/widgets/SliverGrid/SliverGrid.count.html
type SliverGridCountAttributes struct {
	*ValueReferenceHolder        `gen:"wrap"`
	*ThemeConsumer               `gen:"wrap"`
	*DefaultSliverGridAttributes `gen:"wrap"`
	CrossAxisCount               int     `json:"crossAxisCount"`
	MainAxisSpacing              float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing             float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio             float32 `json:"childAspectRatio,omitempty"`
}

// NewSliverGridCountAttributes creates a new SliverGridCountAttributes instance.
func NewSliverGridCountAttributes(data *SliverGridCountAttributes) *SliverGridAttributes {
	data.DefaultSliverGridAttributes.Constructor = duit_utils.TristateFrom[SliverGridConstructor](SliverGridCount)
	return &SliverGridAttributes{data: data}
}

// SetCrossAxisCount sets the crossAxisCount for the sliver grid count widget.
func (r *SliverGridCountAttributes) SetCrossAxisCount(crossAxisCount int) *SliverGridCountAttributes {
	r.CrossAxisCount = crossAxisCount
	return r
}

// SetMainAxisSpacing sets the mainAxisSpacing for the sliver grid count widget.
func (r *SliverGridCountAttributes) SetMainAxisSpacing(mainAxisSpacing float32) *SliverGridCountAttributes {
	r.MainAxisSpacing = mainAxisSpacing
	return r
}

// SetCrossAxisSpacing sets the crossAxisSpacing for the sliver grid count widget.
func (r *SliverGridCountAttributes) SetCrossAxisSpacing(crossAxisSpacing float32) *SliverGridCountAttributes {
	r.CrossAxisSpacing = crossAxisSpacing
	return r
}

// SetChildAspectRatio sets the childAspectRatio for the sliver grid count widget.
func (r *SliverGridCountAttributes) SetChildAspectRatio(childAspectRatio float32) *SliverGridCountAttributes {
	r.ChildAspectRatio = childAspectRatio
	return r
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

// SliverGridBuilderAttributes defines attributes for SliverGridBuilder widget.
// See: https://api.flutter.dev/flutter/widgets/SliverGrid/SliverGrid.builder.html
type SliverGridBuilderAttributes struct {
	*ValueReferenceHolder        `gen:"wrap"`
	*ThemeConsumer               `gen:"wrap"`
	*Builder                     `gen:"wrap"`
	*DefaultSliverGridAttributes `gen:"wrap"`
	SliverGridDelegateKey        string                 `json:"sliverGridDelegateKey"`
	SliverGridDelegateOptions    map[string]interface{} `json:"sliverGridDelegateOptions,omitempty"`
}

// NewSliverGridBuilderAttributes creates a new SliverGridBuilderAttributes instance.
func NewSliverGridBuilderAttributes(data *SliverGridBuilderAttributes) *SliverGridAttributes {
	data.DefaultSliverGridAttributes.Constructor = duit_utils.TristateFrom[SliverGridConstructor](SliverGridBuilder)
	return &SliverGridAttributes{data: data}
}

// SetSliverGridDelegateKey sets the sliverGridDelegateKey for the sliver grid builder widget.
func (r *SliverGridBuilderAttributes) SetSliverGridDelegateKey(sliverGridDelegateKey string) *SliverGridBuilderAttributes {
	r.SliverGridDelegateKey = sliverGridDelegateKey
	return r
}

// SetSliverGridDelegateOptions sets the sliverGridDelegateOptions for the sliver grid builder widget.
func (r *SliverGridBuilderAttributes) SetSliverGridDelegateOptions(sliverGridDelegateOptions map[string]interface{}) *SliverGridBuilderAttributes {
	r.SliverGridDelegateOptions = sliverGridDelegateOptions
	return r
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

// SliverGridExtentAttributes defines attributes for SliverGridExtent widget.
// See: https://api.flutter.dev/flutter/widgets/SliverGrid/SliverGrid.extent.html
type SliverGridExtentAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	*DefaultSliverGridAttributes
	MaxCrossAxisExtent float32 `json:"maxCrossAxisExtent"`
	MainAxisSpacing    float32 `json:"mainAxisSpacing,omitempty"`
	CrossAxisSpacing   float32 `json:"crossAxisSpacing,omitempty"`
	ChildAspectRatio   float32 `json:"childAspectRatio,omitempty"`
}

// NewSliverGridExtentAttributes creates a new SliverGridExtentAttributes instance.
func NewSliverGridExtentAttributes(data *SliverGridExtentAttributes) *SliverGridAttributes {
	data.DefaultSliverGridAttributes.Constructor = duit_utils.TristateFrom[SliverGridConstructor](SliverGridExtent)
	return &SliverGridAttributes{data: data}
}

// SetMaxCrossAxisExtent sets the maxCrossAxisExtent for the sliver grid extent widget.
func (r *SliverGridExtentAttributes) SetMaxCrossAxisExtent(maxCrossAxisExtent float32) *SliverGridExtentAttributes {
	r.MaxCrossAxisExtent = maxCrossAxisExtent
	return r
}

// SetMainAxisSpacing sets the mainAxisSpacing for the sliver grid extent widget.
func (r *SliverGridExtentAttributes) SetMainAxisSpacing(mainAxisSpacing float32) *SliverGridExtentAttributes {
	r.MainAxisSpacing = mainAxisSpacing
	return r
}

// SetCrossAxisSpacing sets the crossAxisSpacing for the sliver grid extent widget.
func (r *SliverGridExtentAttributes) SetCrossAxisSpacing(crossAxisSpacing float32) *SliverGridExtentAttributes {
	r.CrossAxisSpacing = crossAxisSpacing
	return r
}

// SetChildAspectRatio sets the childAspectRatio for the sliver grid extent widget.
func (r *SliverGridExtentAttributes) SetChildAspectRatio(childAspectRatio float32) *SliverGridExtentAttributes {
	r.ChildAspectRatio = childAspectRatio
	return r
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

type SliverGridAttributes struct {
	data any
}

// NewSliverGridAttributes creates a new SliverGridAttributes instance.
func NewSliverGridAttributes(data any) *SliverGridAttributes {
	switch data := data.(type) {
	case *SliverGridCommonAttributes:
		return NewSliverGridCommonAttributes(data)
	case *SliverGridCountAttributes:
		return NewSliverGridCountAttributes(data)
	case *SliverGridBuilderAttributes:
		return NewSliverGridBuilderAttributes(data)
	case *SliverGridExtentAttributes:
		return NewSliverGridExtentAttributes(data)
	}

	return nil
}

func (r *SliverGridAttributes) Validate() error {
	switch data := r.data.(type) {
	case *SliverGridCommonAttributes:
		return data.Validate()
	case *SliverGridCountAttributes:
		return data.Validate()
	case *SliverGridBuilderAttributes:
		return data.Validate()
	case *SliverGridExtentAttributes:
		return data.Validate()
	}
	return nil
}

func (r *SliverGridAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.data)
}
