package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// AnimatedBuilderAttributes represents attributes for AnimatedBuilder widget.
// See: https://api.flutter.dev/flutter/widgets/AnimatedBuilder-class.html
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

// NewAnimatedBuilderAttributes creates a new instance of AnimatedBuilderAttributes.
func NewAnimatedBuilderAttributes() *AnimatedBuilderAttributes {
	return &AnimatedBuilderAttributes{}
}

// SetTweens sets the tweens property and returns the AnimatedBuilderAttributes instance for method chaining.
func (r *AnimatedBuilderAttributes) SetTweens(tweens []any) *AnimatedBuilderAttributes {
	r.Tweens = tweens
	return r
}

func (r *AnimatedBuilderAttributes) AddTween(tween any) *AnimatedBuilderAttributes {
	r.Tweens = append(r.Tweens, tween)
	return r
}

// SetPersistentId sets the persistentId property and returns the AnimatedBuilderAttributes instance for method chaining.
func (r *AnimatedBuilderAttributes) SetPersistentId(persistentId string) *AnimatedBuilderAttributes {
	r.PersistentId = persistentId
	return r
}
