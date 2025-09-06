package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

type AnimatedBuilderAttributes struct {
	*ValueReferenceHolder
	Tweens       []any  `json:"tweens"`
	PersistentId string `json:"persistentId,omitempty"`
}

func (r *AnimatedBuilderAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if r.Tweens == nil {
		return errors.New("tweens property is required and must be non-empty")
	}

	for _, tween := range r.Tweens {
		if err := duit_props.CheckTweenType(tween); err != nil {
			return err
		}
	}

	return nil
}
