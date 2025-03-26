package duit_material

type FabLocation uint8

const (
	CenterDocked FabLocation = iota
	CenterFloat
	CenterTop
	EndDocked
	EndFloat
	EndTop
	StartDocked
	StartFloat
	StartTop
	MiniCenterDocked
	MiniCenterFloat
	MiniCenterTop
	MiniEndDocked
	MiniEndFloat
	MiniEndTop
	MiniStartDocked
	MiniStartFloat
	MiniStartTop
)
