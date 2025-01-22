package duit_core

type ComponentDescription struct {
	Tag        string                      `json:"tag"`
	LayoutRoot *DuitElementModel `json:"layoutRoot"`
}

func NewComponentDescription(tag string, layoutRoot *DuitElementModel) *ComponentDescription {
	return &ComponentDescription{
		Tag:        tag,
		LayoutRoot: layoutRoot,
	}
}