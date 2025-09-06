package duit_props

import (
	"encoding/json"
	"errors"
	"strconv"
)

type EdgeInsets struct {
	data any
}

func NewEdgeInsetsAll(value float32) *EdgeInsets {
	return &EdgeInsets{data: value}
}

func NewEdgeInsetsLTRB(left, top, right, bottom float32) *EdgeInsets {
	return &EdgeInsets{data: [4]float32{left, top, right, bottom}}
}

func NewEdgeInsetsSymmetric(vertical, horizontal float32) *EdgeInsets {
	return &EdgeInsets{data: [2]float32{vertical, horizontal}}
}

func (r *EdgeInsets) Validate() error {
	if r == nil {
		return nil
	}

	switch r.data.(type) {
	case float32:
		if r.data.(float32) < 0 {
			return errors.New("edge insets must be greater than 0")
		}
	case [4]float32:
		for index, value := range r.data.([4]float32) {
			if value < 0 {
				return errors.New("edge insets must be greater than 0, error at index: " + strconv.Itoa(index))
			}
		}
	case [2]float32:
		for index, value := range r.data.([2]float32) {
			if value < 0 {
				return errors.New("edge insets must be greater than 0, error at index: " + strconv.Itoa(index))
			}
		}
	default:
		return errors.New("invalid edge insets type")
	}

	return nil
}

func (r *EdgeInsets) MarshalJSON() ([]byte, error) { return json.Marshal(r.data) }
