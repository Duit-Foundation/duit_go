package duit_props

import (
	"encoding/json"
	"errors"
	"strconv"
)

type EdgeInsetsAll = float32

// Assign EdgeInsets in order -> left, top, right, bottom
type EdgeInsetsLTRB = [4]float32

// Assign EdgeInsets in order -> vertival, horizontal
type EdgeInsetsSymmentric = [2]float32

type EdgeInsets interface {
	EdgeInsetsAll | EdgeInsetsLTRB | EdgeInsetsSymmentric
}

type EdgeInsetsV2 struct {
	data any
}

func (e EdgeInsetsV2) All(value float32) *EdgeInsetsV2 {
	e.data = EdgeInsetsAll(value)
	return &e
}

func (e EdgeInsetsV2) LTRB(left, top, right, bottom float32) *EdgeInsetsV2 {
	e.data = EdgeInsetsLTRB{left, top, right, bottom}
	return &e
}

func (e EdgeInsetsV2) Symmentric(vertical, horizontal float32) *EdgeInsetsV2 {
	e.data = EdgeInsetsSymmentric{vertical, horizontal}
	return &e
}

func (e *EdgeInsetsV2) Validate() error {
	switch e.data.(type) {
	case EdgeInsetsAll:
		if e.data.(EdgeInsetsAll) < 0 {
			return errors.New("edge insets must be greater than 0")
		}
	case EdgeInsetsLTRB:
		for index, value := range e.data.(EdgeInsetsLTRB) {
			if value < 0 {
				return errors.New("edge insets must be greater than 0, error at index: " + strconv.Itoa(index))
			}
		}
	case EdgeInsetsSymmentric:
		for index, value := range e.data.(EdgeInsetsSymmentric) {
			if value < 0 {
				return errors.New("edge insets must be greater than 0, error at index: " + strconv.Itoa(index))
			}
		}
	default:
		return errors.New("invalid edge insets type")
	}

	return nil
}

func (e *EdgeInsetsV2) MarshalJSON() ([]byte, error) { return json.Marshal(e.data) }
