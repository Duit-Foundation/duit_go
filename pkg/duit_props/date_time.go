package duit_props

import (
	"encoding/json"
	"time"
)

type DateTime struct {
	data any
}

func NewDateTimeFromString(value string) *DateTime {
	return &DateTime{data: value}
}

func NewDateTimeFrom(value time.Time) *DateTime {
	return &DateTime{data: value}
}

func NewDateTimeFromMicrosecondsSinceEpoch(value int64) *DateTime {
	return &DateTime{data: value}
}

func (r *DateTime) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }
