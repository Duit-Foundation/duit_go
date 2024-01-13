package duit_core

import (
	"encoding/json"
	"errors"
)

type UiBuilder struct {
	root *DuitElementModel
}

// Build generates a JSON representation of the UiBuilder.
//
// It returns a byte slice and an error.
func (builder *UiBuilder) Build() ([]byte, error) {
	json, err := json.Marshal(*builder.root)

	if err != nil {
		return nil, errors.New("Failed to build JSON: " + err.Error())
	}

	return json, nil
}

// CreateRoot creates a root DuitElementModel for the UiBuilder.
//
// No parameters.
// Returns a pointer to the created root DuitElementModel.
func (builder *UiBuilder) CreateRoot() *DuitElementModel {
	builder.root = &DuitElementModel{
		ElementType: Column,
	}
	return builder.root
}

// RootFrom sets the root element of the UiBuilder.
//
// rootElement: The root element to set.
// returns: The updated root element.
func (builder *UiBuilder) RootFrom(rootElement *DuitElementModel) *DuitElementModel {
	builder.root = rootElement
	return builder.root
}

// CreateRootOfExactType creates a root element of the specified type and sets it as the root of the UiBuilder.
//
// Parameters:
// - elType: The type of the root element.
// - attributes: The attributes of the root element.
// - id: The ID of the root element.
//
// Returns:
// - *DuitElementModel: The root element created.
func (builder *UiBuilder) CreateRootOfExactType(elType DuitElementType, attributes interface{}, id string) *DuitElementModel {
	switch elType {
	case Column:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 2, nil)
	case Row:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 2, nil)
	case Stack:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 2, nil)
	case Center:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 1, nil)
	case ColoredBox:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 1, nil)
	case Container:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 1, nil)
	case DecoratedBox:
		builder.root = new(DuitElementModel).CreateElement(elType, id, "", attributes, nil, false, 1, nil)
	default:
		panic("Invalid element type")
	}

	return builder.root
}
