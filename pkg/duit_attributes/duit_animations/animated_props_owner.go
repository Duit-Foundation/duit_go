package duit_animations

type AnimatedPropertyOwner struct {
	ParentBuilderId    string    `json:"parentBuilderId,omitempty"`
	AffectedProperties *[]string `json:"affectedProperties,omitempty"`
}
