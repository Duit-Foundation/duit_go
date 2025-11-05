package duit_props

type SliderInteraction = string

const (
	SliderInteractionTapAndSlide SliderInteraction = "tapAndSlide"
	SliderInteractionTapOnly     SliderInteraction = "tapOnly"
	SliderInteractionSlideOnly   SliderInteraction = "slideOnly"
	SliderInteractionSlideThumb  SliderInteraction = "slideThumb"
)
