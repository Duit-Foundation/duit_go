package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// ImageAttributes defines attributes for Image widget.
// See: https://api.flutter.dev/flutter/widgets/Image-class.html
type ImageAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Width                             float32                   `json:"width,omitempty"`
	Height                            float32                   `json:"height,omitempty"`
	CacheWidth                        int                       `json:"cacheWidth,omitempty"`
	CacheHeight                       int                       `json:"cacheHeight,omitempty"`
	Scale                             float32                   `json:"scale,omitempty"`
	Repeat                            duit_props.ImageRepeat    `json:"repeat,omitempty"`
	Src                               string                    `json:"src"`
	ByteData                          *duit_props.ImageBuffer   `json:"byteData"`
	IsAntiAlias                       duit_utils.Tristate[bool] `json:"isAntiAlias,omitempty"`
	MatchTextDirection                duit_utils.Tristate[bool] `json:"matchTextDirection,omitempty"`
	GaplessPlayback                   duit_utils.Tristate[bool] `json:"gaplessPlayback,omitempty"`
	ExcludeFromSemantics              duit_utils.Tristate[bool] `json:"excludeFromSemantics,omitempty"`
	FilterQuality                     duit_props.FilterQuality  `json:"filterQuality,omitempty"`
	ColorBlendMode                    duit_props.BlendMode      `json:"colorBlendMode,omitempty"`
	Color                             *duit_props.Color         `json:"color,omitempty"`
	Fit                               duit_props.BoxFit         `json:"fit,omitempty"`
	Alignment                         duit_props.Alignment      `json:"alignment,omitempty"`
	Type                              duit_props.ImageType      `json:"type,omitempty"`
}

//bindgen:exclude
func (r *ImageAttributes) Validate() error {
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
	case duit_props.ImageTypeAsset, duit_props.ImageTypeNetwork:
		if r.Src == "" {
			return errors.New("src property is required")
		}
	case duit_props.ImageTypeMemory:
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

// NewImageAttributes creates a new ImageAttributes instance.
func NewImageAttributes() *ImageAttributes {
	return &ImageAttributes{}
}

// SetWidth sets the width for the image widget.
func (r *ImageAttributes) SetWidth(width float32) *ImageAttributes {
	r.Width = width
	return r
}

// SetHeight sets the height for the image widget.
func (r *ImageAttributes) SetHeight(height float32) *ImageAttributes {
	r.Height = height
	return r
}

// SetCacheWidth sets the cache width for the image widget.
func (r *ImageAttributes) SetCacheWidth(cacheWidth int) *ImageAttributes {
	r.CacheWidth = cacheWidth
	return r
}

// SetCacheHeight sets the cache height for the image widget.
func (r *ImageAttributes) SetCacheHeight(cacheHeight int) *ImageAttributes {
	r.CacheHeight = cacheHeight
	return r
}

// SetScale sets the scale for the image widget.
func (r *ImageAttributes) SetScale(scale float32) *ImageAttributes {
	r.Scale = scale
	return r
}

// SetRepeat sets the repeat for the image widget.
func (r *ImageAttributes) SetRepeat(repeat duit_props.ImageRepeat) *ImageAttributes {
	r.Repeat = repeat
	return r
}

// SetSrc sets the src for the image widget.
func (r *ImageAttributes) SetSrc(src string) *ImageAttributes {
	r.Src = src
	return r
}

// SetByteData sets the byte data for the image widget.
func (r *ImageAttributes) SetByteData(byteData *duit_props.ImageBuffer) *ImageAttributes {
	r.ByteData = byteData
	return r
}

// SetIsAntiAlias sets the is anti alias for the image widget.
func (r *ImageAttributes) SetIsAntiAlias(isAntiAlias duit_utils.Tristate[bool]) *ImageAttributes {
	r.IsAntiAlias = isAntiAlias
	return r
}

// SetMatchTextDirection sets the match text direction for the image widget.
func (r *ImageAttributes) SetMatchTextDirection(matchTextDirection duit_utils.Tristate[bool]) *ImageAttributes {
	r.MatchTextDirection = matchTextDirection
	return r
}

// SetGaplessPlayback sets the gapless playback for the image widget.
func (r *ImageAttributes) SetGaplessPlayback(gaplessPlayback duit_utils.Tristate[bool]) *ImageAttributes {
	r.GaplessPlayback = gaplessPlayback
	return r
}

// SetExcludeFromSemantics sets the exclude from semantics for the image widget.
func (r *ImageAttributes) SetExcludeFromSemantics(excludeFromSemantics duit_utils.Tristate[bool]) *ImageAttributes {
	r.ExcludeFromSemantics = excludeFromSemantics
	return r
}

// SetFilterQuality sets the filter quality for the image widget.
func (r *ImageAttributes) SetFilterQuality(filterQuality duit_props.FilterQuality) *ImageAttributes {
	r.FilterQuality = filterQuality
	return r
}

// SetColorBlendMode sets the color blend mode for the image widget.
func (r *ImageAttributes) SetColorBlendMode(colorBlendMode duit_props.BlendMode) *ImageAttributes {
	r.ColorBlendMode = colorBlendMode
	return r
}

// SetColor sets the color for the image widget.
func (r *ImageAttributes) SetColor(color *duit_props.Color) *ImageAttributes {
	r.Color = color
	return r
}

// SetFit sets the fit for the image widget.
func (r *ImageAttributes) SetFit(fit duit_props.BoxFit) *ImageAttributes {
	r.Fit = fit
	return r
}

// SetAlignment sets the alignment for the image widget.
func (r *ImageAttributes) SetAlignment(alignment duit_props.Alignment) *ImageAttributes {
	r.Alignment = alignment
	return r
}

// SetType sets the type for the image widget.
func (r *ImageAttributes) SetType(imageType duit_props.ImageType) *ImageAttributes {
	r.Type = imageType
	return r
}
