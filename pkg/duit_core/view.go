package duit_core

import (
	"os"
)

type DuitView struct{}

func (view *DuitView) Static(filePath string) (string, error) {
	jsonObj, err := os.ReadFile(filePath)

	if err != nil {
		return "", err
	}

	return string(jsonObj), nil
}

func (view *DuitView) Builder() *UiBuilder {
	return &UiBuilder{}
}
