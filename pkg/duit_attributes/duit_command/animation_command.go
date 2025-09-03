package duit_command

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"

type AnimationCommand struct {
	Type            string                           `json:"type"`
	AnimatedPropKey string                           `json:"animatedPropKey"`
	Method          duit_props.AnimationMethod  `json:"method"`
	Trigger         duit_props.AnimationTrigger `json:"trigger"`
}

func NewAnimationCommand(animatedPropKey string, method duit_props.AnimationMethod, trigger duit_props.AnimationTrigger) *AnimationCommand {
	return &AnimationCommand{
		Type:            "animation",
		AnimatedPropKey: animatedPropKey,
		Method:          method,
		Trigger:         trigger,
	}
}
