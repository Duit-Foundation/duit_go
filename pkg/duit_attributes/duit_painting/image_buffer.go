package duit_painting

type BufferType string

const (
	ByteBuffer BufferType = "Buffer" //same name with JS Buffer {type: 'Buffer', data: []}
)

type ImageBuffer struct {
	Type BufferType `json:"type"`
	Data []byte     `json:"data"`
}
