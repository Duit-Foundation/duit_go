package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"

type FocusCapability struct {
	FocusNode *duit_props.FocusNode `json:"focusNode,omitempty"`
}

func (r *FocusCapability) SetFocusNode(node *duit_props.FocusNode) {
	r.FocusNode = node
}
