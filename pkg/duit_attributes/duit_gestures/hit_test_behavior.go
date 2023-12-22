package duit_gestures

type HitTestBehavior string

const (
	DeferToChild HitTestBehavior = "deferToChild"
	Opaque       HitTestBehavior = "opaque"
	Translucent  HitTestBehavior = "translucent"
)
