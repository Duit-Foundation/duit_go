package duit_attributes

import (
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_alignment"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_flex"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_painting"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"
)

type ImageAttributes[TColor duit_utility_styles.Color] struct {
	Width                float32                     `json:"width,omitempty"`
	Height               float32                     `json:"height,omitempty"`
	CacheWidth           int                         `json:"cacheWidth,omitempty"`
	CacheHeight          int                         `json:"cacheHeight,omitempty"`
	Scale                float32                     `json:"scale,omitempty"`
	Repeat               duit_painting.ImageRepeat   `json:"repeat,omitempty"`
	Src                  string                      `json:"src"`
	ByteData             []byte                      `json:"byteData,omitempty"`
	IsAntiAlias          bool                        `json:"isAntiAlias,omitempty"`
	MatchTextDirection   bool                        `json:"matchTextDirection,omitempty"`
	GaplessPlayback      bool                        `json:"gaplessPlayback,omitempty"`
	ExcludeFromSemantics bool                        `json:"excludeFromSemantics,omitempty"`
	FilterQuality        duit_painting.FilterQuality `json:"filterQuality,omitempty"`
	ColorBlendMode       duit_painting.BlendMode     `json:"colorBlendMode,omitempty"`
	Color                TColor                      `json:"color,omitempty"`
	Fit                  duit_flex.BoxFit            `json:"fit,omitempty"`
	Alignment            duit_alignment.Alignment    `json:"alignment,omitempty"`
	Type                 duit_painting.ImageType     `json:"type,omitempty"`
}
