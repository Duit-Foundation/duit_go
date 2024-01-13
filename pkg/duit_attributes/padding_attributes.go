package duit_attributes

import "github.com/lesleysin/duit_go/pkg/duit_attributes/duit_edge_insets"

type PaddingAttributes[T duit_edge_insets.EdgeInsets] struct {
	ValueReferenceHolder
	Padding T `json:"padding,omitempty"`
}
