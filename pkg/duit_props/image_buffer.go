package duit_props

import (
	"fmt"
	"strings"
)

type BufferType string

const (
	BufferTypeByteBuffer BufferType = "Buffer" //same name with JS Buffer {type: 'Buffer', data: []}
)

type ImageBuffer struct {
	Type BufferType `json:"type"`
	Data []byte     `json:"data"`
}

func (ib *ImageBuffer) MarshalJSON() ([]byte, error) { // Custom struct marshal function
	var array string
	if ib.Data == nil {
		array = "null"
	} else {
		array = strings.Join(strings.Fields(fmt.Sprintf("%d", ib.Data)), ",")
	}
	jsonResult := fmt.Sprintf(`{"type":%q,"data":%s}`, ib.Type, array)

	return []byte(jsonResult), nil
}
