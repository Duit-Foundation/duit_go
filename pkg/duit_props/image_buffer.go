package duit_props

import (
	"encoding/json"
)

const (
	BufferTypeByteBuffer = "Buffer" //same name with JS Buffer {type: 'Buffer', data: []}
)

type ImageBuffer struct {
	Type string `json:"type"`
	Data []byte `json:"data"`
}

func NewImageBuffer(data []byte) *ImageBuffer {
	return &ImageBuffer{
		Type: BufferTypeByteBuffer,
		Data: data,
	}
}

func (r *ImageBuffer) MarshalJSON() ([]byte, error) { // Custom struct marshal function
	if r.Data == nil {
		return json.Marshal(struct {
			Type string `json:"type"`
			Data any    `json:"data"`
		}{
			Type: r.Type,
			Data: nil,
		})
	}

	// Более эффективное преобразование в массив
	dataArray := make([]int, len(r.Data))
	for i, b := range r.Data {
		dataArray[i] = int(b)
	}

	return json.Marshal(struct {
		Type string `json:"type"`
		Data []int  `json:"data"`
	}{
		Type: r.Type,
		Data: dataArray,
	})
}
