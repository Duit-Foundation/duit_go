package duit_props

type TextHeightBehavior struct {
	ApplyHeightToFirstAscent bool                     `json:"applyHeightToFirstAscent,omitempty"`
	ApplyHeightToLastDescent bool                     `json:"applyHeightToLastDescent,omitempty"`
	LeadingDistribution      TextLeadingDistribution `json:"leadingDistribution,omitempty"`
}
