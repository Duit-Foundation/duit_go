package duit_attributes

import (
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type TextFieldAttributes[TInsets duit_props.EdgeInsets, TColor duit_props.Color] struct {
	*ValueReferenceHolder
	*ThemeConsumer
	Style              *duit_props.TextStyle[TColor]           `json:"style,omitempty"`
	TextAlign          duit_props.TextAlign                    `json:"textAlign,omitempty"`
	TextDirection      duit_props.TextDirection                `json:"textDirection,omitempty"`
	Autocorrect        duit_utils.Tristate[bool]                         `json:"autocorrect,omitempty"`
	EnableSuggestions  duit_utils.Tristate[bool]                         `json:"enableSuggestions,omitempty"`
	Expands            duit_utils.Tristate[bool]                         `json:"expands,omitempty"`
	ReadOnly           duit_utils.Tristate[bool]                         `json:"readOnly,omitempty"`
	ShowCursor         duit_utils.Tristate[bool]                         `json:"showCursor,omitempty"`
	Enabled            duit_utils.Tristate[bool]                         `json:"enabled,omitempty"`
	ObscureText        duit_utils.Tristate[bool]                         `json:"obscureText,omitempty"`
	Autofocus          duit_utils.Tristate[bool]                         `json:"autofocus,omitempty"`
	ObscuringCharacter string                                            `json:"obscuringCharacter,omitempty"`
	MaxLines           uint                                              `json:"maxLines,omitempty"`
	MinLines           uint                                              `json:"minLines,omitempty"`
	MaxLength          uint                                              `json:"maxLength,omitempty"`
	Decoration         *duit_decoration.InputDecoration[TInsets, TColor] `json:"decoration,omitempty"`
}

func (r *TextFieldAttributes[TInsets, TColor]) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}
