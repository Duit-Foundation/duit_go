package duit_props

type HitTestBehavior = string

const (
	HitTestBehaviorDeferToChild HitTestBehavior = "deferToChild"
	HitTestBehaviorOpaque       HitTestBehavior = "opaque"
	HitTestBehaviorTranslucent  HitTestBehavior = "translucent"
)
