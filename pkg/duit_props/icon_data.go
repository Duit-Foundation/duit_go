package duit_props

import "encoding/json"

type IconData struct {
	data any
}

func NewIcon(iconName string) *IconData {
	return &IconData{data: iconName}
}

func NewIconData(codePoint int, fontFamily, fontPackage string, matchTextDirection bool, fontFamilyFallback []string) *IconData {
	return &IconData{data: map[string]any{
		"codePoint":          codePoint,
		"fontFamily":         fontFamily,
		"fontPackage":        fontPackage,
		"matchTextDirection": matchTextDirection,
		"fontFamilyFallback": fontFamilyFallback,
	}}
}

func (r *IconData) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }