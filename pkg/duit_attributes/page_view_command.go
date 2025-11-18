package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type pageViewGenericCommand struct {
	Type     string                                 `json:"type"`
	Duration duit_utils.Tristate[uint]              `json:"duration,omitempty"`
	Curve    duit_utils.Tristate[duit_props.Curves] `json:"curve,omitempty"`
	Value    duit_utils.Tristate[float32]           `json:"value,omitempty"`
	Offset   duit_utils.Tristate[float32]           `json:"offset,omitempty"`
	Page     duit_utils.Tristate[uint]              `json:"page,omitempty"`
}

func NewPageViewJumpToCommand(value float32) *pageViewGenericCommand {
	return &pageViewGenericCommand{
		Type:  "jumpTo",
		Value: duit_utils.Float32Value(value),
	}
}

func NewPageViewJumpToPageCommand(page uint) *pageViewGenericCommand {
	return &pageViewGenericCommand{
		Type: "jumpToPage",
		Page: duit_utils.UintValue(page),
	}
}

func NewPageViewNextPageCommand(duration uint, curve duit_props.Curves) *pageViewGenericCommand {
	return &pageViewGenericCommand{
		Type:     "nextPage",
		Duration: duit_utils.UintValue(duration),
		Curve:    duit_utils.StringValue(curve),
	}
}

func NewPageViewPrevPageCommand(duration uint, curve duit_props.Curves) *pageViewGenericCommand {
	return &pageViewGenericCommand{
		Type:     "previousPage",
		Duration: duit_utils.UintValue(duration),
		Curve:    duit_utils.StringValue(curve),
	}
}

func NewPageViewAnimateToCommand(offset float32, duration uint, curve duit_props.Curves) *pageViewGenericCommand {
	return &pageViewGenericCommand{
		Type:     "animateTo",
		Offset:   duit_utils.Float32Value(offset),
		Duration: duit_utils.UintValue(duration),
		Curve:    duit_utils.StringValue(curve),
	}
}

func NewPageViewAnimateToPageCommand(page, duration uint, curve duit_props.Curves) *pageViewGenericCommand {
	return &pageViewGenericCommand{
		Type:     "animateToPage",
		Page:     duit_utils.UintValue(page),
		Duration: duit_utils.UintValue(duration),
		Curve:    duit_utils.StringValue(curve),
	}
}
