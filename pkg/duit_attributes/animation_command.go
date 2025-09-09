package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"

type AnimationCommand struct {
	Type            string                      `json:"type"`
	AnimatedPropKey string                      `json:"animatedPropKey"`
	Method          duit_props.AnimationMethod  `json:"method"`
	Trigger         duit_props.AnimationTrigger `json:"trigger"`
}

func NewAnimationCommand() *AnimationCommand {
	return &AnimationCommand{
		Type: "animation",
	}
}

// SetAnimatedPropKey sets the animated prop key for the animation command.
func (r *AnimationCommand) SetAnimatedPropKey(animatedPropKey string) *AnimationCommand {
	r.AnimatedPropKey = animatedPropKey
	return r
}

// SetMethod sets the method for the animation command.
func (r *AnimationCommand) SetMethod(method duit_props.AnimationMethod) *AnimationCommand {
	r.Method = method
	return r
}

// SetTrigger sets the trigger for the animation command.
func (r *AnimationCommand) SetTrigger(trigger duit_props.AnimationTrigger) *AnimationCommand {
	r.Trigger = trigger
	return r
}
