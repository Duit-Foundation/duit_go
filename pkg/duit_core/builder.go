package duit_core

import (
	"encoding/json"
	"errors"
)

type UiBuilder struct {
	root *DuitElementModel
}

func (builder *UiBuilder) Build() ([]byte, error) {
	json, err := json.Marshal(*builder.root)

	if err != nil {
		return nil, errors.New("Failed to build JSON: " + err.Error())
	}

	return json, nil
}

func (builder *UiBuilder) CreateRoot() *DuitElementModel {
	builder.root = &DuitElementModel{
		ElementType: Column,
	}
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
		builder.root = new(DuitElementModel).CreateElement(Column, id, "", attributes, nil, false, 2)
	case Row:
		builder.root = new(DuitElementModel).CreateElement(Row, id, "", attributes, nil, false, 2)
	case Stack:
		builder.root = new(DuitElementModel).CreateElement(Stack, id, "", attributes, nil, false, 2)
	}

	return builder.root
}
