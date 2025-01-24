package duit_core

import (
	"encoding/json"
	"errors"
)

type UiBuilder struct {
	root    *DuitElementModel
	widgets []*DuitElementModel
}

// Build generates a JSON representation of the UiBuilder.
//
// It returns a byte slice and an error.
func (builder *UiBuilder) Build() ([]byte, error) {
	content := map[string]interface{}{
		"root": builder.root,
	}

	json, err := json.Marshal(content)

	if err != nil {
		return nil, errors.New("Failed to build JSON: " + err.Error())
	}

	return json, nil
}

func (builder *UiBuilder) BuildMultiview() ([]byte, error) {
	if len(builder.widgets) == 0 {
		return nil, errors.New("no widgets provided")
	}
 
	content := map[string]interface{}{
		"widgets": builder.widgets,
	}

	json, err := json.Marshal(content)

	if err != nil {
		return nil, errors.New("Failed to build JSON: " + err.Error())
	}

	return json, nil
}

// Build generates a JSON representation of the UiBuilder without wrapping it in a map.
//
// It returns a byte slice and an error.
func (builder *UiBuilder) BuildUnwrapped() ([]byte, error) {
	json, err := json.Marshal(builder.root)

	if err != nil {
		return nil, errors.New("Failed to build JSON: " + err.Error())
	}

	return json, nil
}

// CreateRoot creates a root DuitElementModel for the UiBuilder.
//
// No parameters.
// Returns a pointer to the created root Subtree widget model.
func (builder *UiBuilder) CreateRoot() *DuitElementModel {
	builder.root = &DuitElementModel{
		ElementType: Subtree,
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

func (builder *UiBuilder) AddWidgets(widgets ...*DuitElementModel) {
	builder.widgets = append(builder.widgets, widgets...)
}