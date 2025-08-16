package duit_animations

import "errors"

type AnimatedPropertyOwner struct {
	ParentBuilderId    string   `json:"parentBuilderId,omitempty"`
	AffectedProperties []string `json:"affectedProperties,omitempty"`
}

func (r *AnimatedPropertyOwner) Validate() error {
	if r == nil {
		return nil
	}

	if len(r.AffectedProperties) == 0 {
		return errors.New("affectedProperties property must not be empty slice")
	}

	if len(r.ParentBuilderId) == 0 {
		return errors.New("parentBuilderId property must not be empty string")
	}

	return nil
}
