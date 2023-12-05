package duit_attributes

import (
	duit_decoration "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_decorations"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_text_properties"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_utility_styles"
)

type TextFieldAttributes[TInsets duit_edge_insets.EdgeInsets, TColor duit_utility_styles.Color] struct {
	Style              duit_text_properties.TextStyle[TColor]           `json:"style,omitempty"`
	TextAlign          duit_text_properties.TextAlign                   `json:"textAlign,omitempty"`
	TextDirection      duit_text_properties.TextDirection               `json:"textDirection,omitempty"`
	Autocorrect        bool                                             `json:"autocorrect,omitempty"`
	EnableSuggestions  bool                                             `json:"enableSuggestions,omitempty"`
	Expands            bool                                             `json:"expands,omitempty"`
	ReadOnly           bool                                             `json:"readOnly,omitempty"`
	ShowCursor         bool                                             `json:"showCursor,omitempty"`
	Enabled            bool                                             `json:"enabled,omitempty"`
	ObscureText        bool                                             `json:"obscureText,omitempty"`
	Autofocus          bool                                             `json:"autofocus,omitempty"`
	ObscuringCharacter string                                           `json:"obscuringCharacter,omitempty"`
	MaxLines           uint                                             `json:"maxLines,omitempty"`
	MinLines           uint                                             `json:"minLines,omitempty"`
	MaxLength          uint                                             `json:"maxLength,omitempty"`
	Decoration         duit_decoration.InputDecoration[TInsets, TColor] `json:"decoration,omitempty"`
}
