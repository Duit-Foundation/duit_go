package duit_widget

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"

// Fragment creates a Fragment widget element.
//
// The parameter tag is required and must be a non-empty string. The id parameter
// is optional and may be an empty string.
//
// *If tag is empty the function panics!*
//
// Example:
//  f := Fragment("profile", "profile-1")
//  _ = f
func Fragment(tag string, id string) *duit_core.DuitElementModel {
	if tag == "" {
		panic("Tag is required in Fragment definition")
	}

	return new(duit_core.DuitElementModel).CreateElement(duit_core.Fragment, id, tag, nil, nil, false, 0, nil, nil)
}
