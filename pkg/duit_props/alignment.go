package duit_props

type Alignment = string
type AlignmentDirectional = string

const (
	AlignmentBottomCenter Alignment = "bottomCenter"
	AlignmentBottomLeft   Alignment = "bottomLeft"
	AlignmentBottomRight  Alignment = "bottomRight"
	AlignmentCenter       Alignment = "center"
	AlignmentCenterLeft   Alignment = "centerLeft"
	AlignmentCenterRight  Alignment = "centerRight"
	AlignmentTopCenter    Alignment = "topCenter"
	AlignmentTopLeft      Alignment = "topLeft"
	AlignmentTopRight     Alignment = "topRight"
)

const (
	AlignmentDirectionalBottomCenter AlignmentDirectional = "bottomCenter"
	AlignmentDirectionalBottomEnd    AlignmentDirectional = "bottomEnd"
	AlignmentDirectionalBottomStart  AlignmentDirectional = "bottomStart"
	AlignmentDirectionalCenter       AlignmentDirectional = "center"
	AlignmentDirectionalCenterStart  AlignmentDirectional = "centerStart"
	AlignmentDirectionalCenterEnd    AlignmentDirectional = "centerEnd"
	AlignmentDirectionalTopCenter    AlignmentDirectional = "topCenter"
	AlignmentDirectionalTopEnd       AlignmentDirectional = "topEnd"
	AlignmentDirectionalTopStart     AlignmentDirectional = "topStart"
)
