package duit_props

type FabLocation = uint8

const (
	FabLocationCenterDocked FabLocation = iota
	FabLocationCenterFloat
	FabLocationCenterTop
	FabLocationEndDocked
	FabLocationEndFloat
	FabLocationEndTop
	FabLocationStartDocked
	FabLocationStartFloat
	FabLocationStartTop
	FabLocationMiniCenterDocked
	FabLocationMiniCenterFloat
	FabLocationMiniCenterTop
	FabLocationMiniEndDocked
	FabLocationMiniEndFloat
	FabLocationMiniEndTop
	FabLocationMiniStartDocked
	FabLocationMiniStartFloat
	FabLocationMiniStartTop
)
