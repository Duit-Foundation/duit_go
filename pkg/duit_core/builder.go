package duit_core

import (
	"encoding/json"
	"errors"
)

type UiBuilder struct {
	root             *DuitElementModel
	widgets          []*DuitElementModel // Deprecated
	embedded         []*ComponentDescription
	multiviewWidgets map[string]*DuitElementModel
}

// Build generates a JSON representation of the UiBuilder.
//
// It returns a byte slice and an error.
func (builder *UiBuilder) Build() ([]byte, error) {
	content := map[string]interface{}{
		"root":     builder.root,
		"embedded": builder.embedded,
	}

	json, err := json.Marshal(content)

	if err != nil {
		return nil, errors.New("Failed to build JSON: " + err.Error())
	}

	return json, nil
}

func (builder *UiBuilder) BuildMultiview() ([]byte, error) {
	if len(builder.multiviewWidgets) == 0 {
		return nil, errors.New("no widgets provided")
	}

	content := map[string]interface{}{
		"widgets": builder.multiviewWidgets,
		"embedded": builder.embedded,
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
	if builder.embedded != nil {
		return nil, errors.New("embedded components not supported")
	}

	if builder.multiviewWidgets != nil {
		return nil, errors.New("multiview widgets not supported")
	}

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

// Deprecated: AddWidgets is broken, use AddWidget instead
func (builder *UiBuilder) AddWidgets(widgets ...*DuitElementModel) {
	builder.widgets = append(builder.widgets, widgets...)
}

// AddWidget adds a single widget to the UiBuilder.
// The widget is appended to the list of widgets managed by the builder.
func (builder *UiBuilder) AddWidget(tag string, widget *DuitElementModel) *UiBuilder {
	if builder.multiviewWidgets == nil {
		builder.multiviewWidgets = map[string]*DuitElementModel{}
		builder.multiviewWidgets[tag] = widget
		return  builder
	}
	builder.multiviewWidgets[tag] = widget
	return  builder
}

// AddComponents adds components to the UiBuilder.
//
// Components are arbitrary data structures that are stored in the "embedded" field of the JSON representation of the UiBuilder.
//
// components: The components to add.
func (builder *UiBuilder) AddComponents(components ...*ComponentDescription) {
	builder.embedded = append(builder.embedded, components...)
}
