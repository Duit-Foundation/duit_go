package duit_attributes

import (
	duit_decoration "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_decorations"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"
	"github.com/lesleysin/duit_go/pkg/duit_attributes/duit_text_properties"
)

type TextFieldAttributes[T duit_edge_insets.EdgeInsets] struct {
	Style              duit_text_properties.TextStyle     `json:"style,omitempty"`
	TextAlign          duit_text_properties.TextAlign     `json:"textAlign,omitempty"`
	TextDirection      duit_text_properties.TextDirection `json:"textDirection,omitempty"`
	Autocorrect        bool                               `json:"autocorrect,omitempty"`
	EnableSuggestions  bool                               `json:"enableSuggestions,omitempty"`
	Expands            bool                               `json:"expands,omitempty"`
	ReadOnly           bool                               `json:"readOnly,omitempty"`
	ShowCursor         bool                               `json:"showCursor,omitempty"`
	Enabled            bool                               `json:"enabled,omitempty"`
	ObscureText        bool                               `json:"obscureText,omitempty"`
	Autofocus          bool                               `json:"autofocus,omitempty"`
	ObscuringCharacter string                             `json:"obscuringCharacter,omitempty"`
	MaxLines           int                                `json:"maxLines,omitempty"`
	MinLines           int                                `json:"minLines,omitempty"`
	MaxLength          int                                `json:"maxLength,omitempty"`
	Decoration         duit_decoration.InputDecoration[T] `json:"decoration,omitempty"`
}
