package duit_props

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"

// Go representation of FocusNode object.
// https://api.flutter.dev/flutter/widgets/FocusNode-class.html
type FocusNode struct {
	DebugLabel                duit_utils.TString `json:"debugLabel,omitempty"`
	SkipTraversal             duit_utils.TBool   `json:"skipTraversal,omitempty"`
	CanRequestFocus           duit_utils.TBool   `json:"canRequestFocus,omitempty"`
	DescendantsAreFocusable   duit_utils.TBool   `json:"descendantsAreFocusable,omitempty"`
	DescendantsAreTraversable duit_utils.TBool   `json:"descendantsAreTraversable,omitempty"`
}

// NewFocusNode creates an empty FocusNode instance.
func NewFocusNode() *FocusNode {
	return &FocusNode{}
}

// SetDebugLabel sets a debug label for the FocusNode.
func (r *FocusNode) SetDebugLabel(label string) *FocusNode {
	r.DebugLabel = duit_utils.TristateFrom[string](label)
	return r
}

// SetSkipTraversal specifies whether the FocusNode should be skipped during traversal.
func (r *FocusNode) SetSkipTraversal(value bool) *FocusNode {
	r.SkipTraversal = duit_utils.TristateFrom[bool](value)
	return r
}

// SetCanRequestFocus sets whether the FocusNode can request focus.
func (r *FocusNode) SetCanRequestFocus(value bool) *FocusNode {
	r.CanRequestFocus = duit_utils.TristateFrom[bool](value)
	return r
}

// SetDescendantsAreFocusable specifies if descendants of this node are focusable.
func (r *FocusNode) SetDescendantsAreFocusable(value bool) *FocusNode {
	r.DescendantsAreFocusable = duit_utils.TristateFrom[bool](value)
	return r
}

// SetDescendantsAreTraversable specifies if descendants can be traversed for focus.
func (r *FocusNode) SetDescendantsAreTraversable(value bool) *FocusNode {
	r.DescendantsAreTraversable = duit_utils.TristateFrom[bool](value)
	return r
}
