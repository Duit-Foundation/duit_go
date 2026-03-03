package duit_props

// DismissDirection specifies which direction the Dismissible widget can be dismissed.
// Matches Flutter's DismissDirection enum.
type DismissDirection = string

const (
	DismissDirectionVertical    DismissDirection = "vertical"
	DismissDirectionHorizontal  DismissDirection = "horizontal"
	DismissDirectionUp          DismissDirection = "up"
	DismissDirectionDown        DismissDirection = "down"
	DismissDirectionStartToEnd  DismissDirection = "startToEnd"
	DismissDirectionEndToStart   DismissDirection = "endToStart"
	DismissDirectionNone        DismissDirection = "none"
)
