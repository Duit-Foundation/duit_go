package duit_core

import (
	"os"
)

type DuitView struct{}

// Static reads the content of a file specified by the given file path.
//
// It takes a single parameter, filePath, which is a string representing the path to the file.
// The function returns a string containing the content of the file and an error if any occurred.
func (view *DuitView) Static(filePath string) (string, error) {
	jsonObj, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	return string(jsonObj), nil
}

// Builder returns a new UiBuilder instance.
//
// No parameters.
// *UiBuilder - a pointer to the new UiBuilder instance.
func (view *DuitView) Builder() *UiBuilder {
	return &UiBuilder{}
}
