package duit_edge_insets

import (
	"encoding/json"
	"errors"
)

type EdgeInsetsAll float32

// Assign EdgeInsets in order -> left, top, right, bottom
type EdgeInsetsLTRB [4]float32

// Assign EdgeInsets in order -> vertival, horizontal
type EdgeInsetsSymmentric [2]float32

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
		if e.data.(EdgeInsetsLTRB)[0] < 0 || e.data.(EdgeInsetsLTRB)[1] < 0 || e.data.(EdgeInsetsLTRB)[2] < 0 || e.data.(EdgeInsetsLTRB)[3] < 0 {
			return errors.New("edge insets must be greater than 0")
		}
	case EdgeInsetsSymmentric:
		if e.data.(EdgeInsetsSymmentric)[0] < 0 || e.data.(EdgeInsetsSymmentric)[1] < 0 {
			return errors.New("edge insets must be greater than 0")
		}
	default:
		return errors.New("invalid edge insets type")
	}

	return nil
}

func (e *EdgeInsetsV2) MarshalJSON() ([]byte, error) { return json.Marshal(e.data) }
