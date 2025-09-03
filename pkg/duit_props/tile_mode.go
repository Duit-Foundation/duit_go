package duit_props

type TileMode string

const (
	TileModeClamp  TileMode = "clamp"
	TileModeRepeat TileMode = "repeat"
	TileModeMirror TileMode = "mirror"
	TileModeDecal  TileMode = "decal"
)
