package duit_props

import (
	"encoding/json"
)

type Locale struct {
	data any
}

func NewLocaleString(value string) *Locale {
	return &Locale{data: value}
}

func NewLocaleTags(languageCode, countryCode string) *Locale {
	return &Locale{data: map[string]string{"languageCode": languageCode, "countryCode": countryCode}}
}

func (r *Locale) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }
