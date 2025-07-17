package duit_command

import "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_animations"

type AnimationCommand struct {
	Type            string                           `json:"type"`
	AnimatedPropKey string                           `json:"animatedPropKey"`
	Method          duit_animations.AnimationMethod  `json:"method"`
	Trigger         duit_animations.AnimationTrigger `json:"trigger"`
}

func NewAnimationCommand(animatedPropKey string, method duit_animations.AnimationMethod, trigger duit_animations.AnimationTrigger) *AnimationCommand {
	return &AnimationCommand{
		Type:            "animation",
		AnimatedPropKey: animatedPropKey,
		Method:          method,
		Trigger:         trigger,
	}
}
