package duit_painting

import (
	"fmt"
	"strings"
)

type BufferType string

const (
	ByteBuffer BufferType = "Buffer" //same name with JS Buffer {type: 'Buffer', data: []}
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

// func (s ImageBuffer) x() ([]byte, error) {
// 	type innerStruct struct {
// 		F1 string
// 		F2 []int
// 	}

// 	bytes2int := func(in []byte) []int {
// 		out := make([]int, 0, len(in))

// 		for _, n := range in {
// 			out = append(out, int(n))
// 		}

// 		return out
// 	}

// 	i := innerStruct{
// 		F1: s.F1,
// 		F2: bytes2int(s.F2),
// 	}

// 	return json.Marshal(i)
// }
