package duit_attributes

import (
	"errors"

	animations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_painting"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type ImageAttributes[TColor duit_props.Color] struct {
	*ValueReferenceHolder
	*animations.AnimatedPropertyOwner
	*ThemeConsumer
	Width                float32                     `json:"width,omitempty"`
	Height               float32                     `json:"height,omitempty"`
	CacheWidth           int                         `json:"cacheWidth,omitempty"`
	CacheHeight          int                         `json:"cacheHeight,omitempty"`
	Scale                float32                     `json:"scale,omitempty"`
	Repeat               duit_painting.ImageRepeat   `json:"repeat,omitempty"`
	Src                  string                      `json:"src"`
	ByteData             *duit_painting.ImageBuffer  `json:"byteData"`
	IsAntiAlias          duit_utils.Tristate[bool]   `json:"isAntiAlias,omitempty"`
	MatchTextDirection   duit_utils.Tristate[bool]   `json:"matchTextDirection,omitempty"`
	GaplessPlayback      duit_utils.Tristate[bool]   `json:"gaplessPlayback,omitempty"`
	ExcludeFromSemantics duit_utils.Tristate[bool]   `json:"excludeFromSemantics,omitempty"`
	FilterQuality        duit_painting.FilterQuality `json:"filterQuality,omitempty"`
	ColorBlendMode       duit_painting.BlendMode     `json:"colorBlendMode,omitempty"`
	Color                TColor                      `json:"color,omitempty"`
	Fit                  duit_flex.BoxFit            `json:"fit,omitempty"`
	Alignment            duit_props.Alignment    `json:"alignment,omitempty"`
	Type                 duit_painting.ImageType     `json:"type,omitempty"`
}

func (r *ImageAttributes[TColor]) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}
	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	switch r.Type {
	case duit_painting.Asset, duit_painting.Network:
		if r.Src == "" {
			return errors.New("src property is required")
		}
	case duit_painting.Memory:
		if r.ByteData == nil {
			return errors.New("byteData property is required")
		}
	case "":
		return errors.New("type property is required")
	default:
		return errors.New("invalid image type: " + string(r.Type))
	}

	return nil
}
