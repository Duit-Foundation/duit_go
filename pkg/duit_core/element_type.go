package duit_core

// DuitElementType represents the type of a Duit element.
type DuitElementType string

const (
	SizedBox        DuitElementType = "SizedBox"
	Row             DuitElementType = "Row"
	Column          DuitElementType = "Column"
	Stack           DuitElementType = "Stack"
	ColoredBox      DuitElementType = "ColoredBox"
	Text            DuitElementType = "Text"
	TextField       DuitElementType = "TextField"
	Empty           DuitElementType = "Empty"
	Custom          DuitElementType = "Custom"
	Center          DuitElementType = "Center"
	ElevatedButton  DuitElementType = "ElevatedButton"
	Expanded        DuitElementType = "Expanded"
	Padding         DuitElementType = "Padding"
	Positioned      DuitElementType = "Positioned"
	DecoratedBox    DuitElementType = "DecoratedBox"
	CheckBox        DuitElementType = "CheckBox"
	Container       DuitElementType = "Container"
	Image           DuitElementType = "Image"
	GestureDetector DuitElementType = "GestureDetector"
	Align           DuitElementType = "Align"
	Transform       DuitElementType = "Transform"
)
