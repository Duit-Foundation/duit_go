package duit_attributes

type Clip string

const (
	None                   Clip = "none"
	HardEdge               Clip = "hardEdge"
	AntiAlias              Clip = "antiAlias"
	AntiAliasWithSaveLayer Clip = "antiAliasWithSaveLayer"
)
