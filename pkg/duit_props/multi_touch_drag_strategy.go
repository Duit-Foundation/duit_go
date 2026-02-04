package duit_props

type MultiTouchDragStrategy = string

const (
	MultiTouchDragStrategyLatestPointer           MultiTouchDragStrategy = "latestPointer"
	MultiTouchDragStrategyAverageBoundaryPointers MultiTouchDragStrategy = "averageBoundaryPointers"
	MultiTouchDragStrategySumAllPointers          MultiTouchDragStrategy = "sumAllPointers"
)
