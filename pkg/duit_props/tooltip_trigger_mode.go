package duit_props

// TooltipTriggerMode defines how a tooltip is triggered.
// See: https://api.flutter.dev/flutter/material/TooltipTriggerMode.html
type TooltipTriggerMode string

const (
	// TooltipTriggerModeManual - tooltip will only be shown by calling ensureTooltipVisible.
	TooltipTriggerModeManual TooltipTriggerMode = "manual"
	// TooltipTriggerModeLongPress - tooltip will be shown after a long press.
	TooltipTriggerModeLongPress TooltipTriggerMode = "longPress"
	// TooltipTriggerModeTap - tooltip will be shown after a single tap.
	TooltipTriggerModeTap TooltipTriggerMode = "tap"
)
