package duit_attributes

type TextFieldAttributes[T EdgeInsets] struct {
	Style              TextStyle          `json:"style,omitempty"`
	TextAlign          TextAlign          `json:"textAlign,omitempty"`
	TextDirection      TextDirection      `json:"textDirection,omitempty"`
	Autocorrect        bool               `json:"autocorrect,omitempty"`
	EnableSuggestions  bool               `json:"enableSuggestions,omitempty"`
	Expands            bool               `json:"expands,omitempty"`
	ReadOnly           bool               `json:"readOnly,omitempty"`
	ShowCursor         bool               `json:"showCursor,omitempty"`
	Enabled            bool               `json:"enabled,omitempty"`
	ObscureText        bool               `json:"obscureText,omitempty"`
	Autofocus          bool               `json:"autofocus,omitempty"`
	ObscuringCharacter string             `json:"obscuringCharacter,omitempty"`
	MaxLines           int                `json:"maxLines,omitempty"`
	MinLines           int                `json:"minLines,omitempty"`
	MaxLength          int                `json:"maxLength,omitempty"`
	Decoration         InputDecoration[T] `json:"decoration,omitempty"`
}
