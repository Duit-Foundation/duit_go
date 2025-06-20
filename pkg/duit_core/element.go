package duit_core

import (
	"github.com/google/uuid"
)

type DuitWidget interface {
	CreateElement(elemType string, elemId string, tag string, attributes interface{}, action any, controlled bool, mayHaveChildElements uint8) *DuitElementModel
}

type DuitElementModel struct {
	ElementType DuitElementType `json:"type"`
	Id          string          `json:"id"`
	Controlled  bool            `json:"controlled"`
	Attributes  interface{}     `json:"attributes"`
	Action      any             `json:"action,omitempty"`

	//Special component tag.
	//
	//Necessary when using custom components!
	//
	//It helps determine how exactly the model should be interpreted,
	//rendered and converted attributes on the side of the front-end framework
	Tag string `json:"tag,omitempty"`
	//Data set for populating components
	Data     interface{}         `json:"data,omitempty"`
	Child    *DuitElementModel   `json:"child,omitempty"`
	Children []*DuitElementModel `json:"children,omitempty"`

	//A parameter that determines whether the model can contain child elements
	//
	// 0 - cannot have children
	//
	// 1 - single child model
	//
	// 2 - multi child model
	MayHaveChildElements uint8 `json:"-"`
}

type DuitViewChild interface {
	[1]*DuitElementModel | []*DuitElementModel
}

// CreateElement creates a new instance of DuitElement.
//
// It takes the following parameters:
// - elemType: the type of the element
// - elemId: the ID of the element (optional)
// - tag: the tag of the element
// - attributes: the attributes of the element
// - action: the action associated with the element (optional)
// - controlled: a boolean indicating whether the element is controlled
//
// It returns a pointer to the upgraded created DuitElement.
func (e *DuitElementModel) CreateElement(elemType DuitElementType, elemId string, tag string, attributes interface{}, action any, controlled bool, mayHaveChildElements uint8, data interface{}, subviews ...*DuitElementModel) *DuitElementModel {
	var id string

	if elemId == "" {
		id = uuid.NewString()
	} else {
		id = elemId
	}

	switch act := action.(type) {
	case *RemoteAction:
	case *LocalAction:
	case *ScriptAction:
	case nil:
		e.Action = act
	default:
		panic("Invalid action type. Must be instance of duit_core.Action or nil")
	}

	switch mayHaveChildElements {
	case 1:
		e.Child = subviews[0]
	case 2:
		e.Children = append(e.Children, subviews...)
	}

	e.Id = id
	e.ElementType = elemType
	e.Action = action
	e.Attributes = attributes
	e.Tag = tag
	e.Controlled = controlled
	e.MayHaveChildElements = mayHaveChildElements
	e.Data = data
	return e
}

// AddChild adds a child element to the DuitElement.
//
// The child parameter is the element to be added as a child.
// The function returns the modified DuitElement.
func (element *DuitElementModel) AddChild(child *DuitElementModel) *DuitElementModel {
	switch element.MayHaveChildElements {
	case 1:
		element.Child = child
	case 2:
		element.Children = append(element.Children, child)
	}

	return element
}

// AddChildren adds the specified children to the DuitElement.
//
// It accepts a slice of *DuitElement as the children parameter.
// It returns a *DuitElement, which is the modified DuitElement.
//
// If DuitElement may contains only one child element, last element of slice will be added as child
func (element *DuitElementModel) AddChildren(children []*DuitElementModel) *DuitElementModel {
	switch element.MayHaveChildElements {
	case 1:
		element.Child = children[len(children)-1]
	case 2:
		element.Children = append(element.Children, children...)
	}

	return element
}
