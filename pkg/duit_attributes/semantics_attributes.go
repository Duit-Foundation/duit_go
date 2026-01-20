package duit_attributes

import "github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"

type SemanticsAttributes struct {
	*ValueReferenceHolder  `gen:"wrap"`
	*ThemeConsumer         `gen:"wrap"`
	Label                  string           `json:"label,omitempty"`
	Hint                   string           `json:"hint,omitempty"`
	Value                  string           `json:"value,omitempty"`
	Tooltip                string           `json:"tooltip,omitempty"`
	Enabled                duit_utils.TBool `json:"enabled,omitempty"`
	Checked                duit_utils.TBool `json:"checked,omitempty"`
	Selected               duit_utils.TBool `json:"selected,omitempty"`
	Button                 duit_utils.TBool `json:"button,omitempty"`
	Link                   duit_utils.TBool `json:"link,omitempty"`
	Header                 duit_utils.TBool `json:"header,omitempty"`
	TextField              duit_utils.TBool `json:"textField,omitempty"`
	Image                  duit_utils.TBool `json:"image,omitempty"`
	LiveRegion             duit_utils.TBool `json:"liveRegion,omitempty"`
	Container              duit_utils.TBool `json:"container,omitempty"`
	ExplicitChildNodes     duit_utils.TBool `json:"explicitChildNodes,omitempty"`
	ExcludeSemantics       duit_utils.TBool `json:"excludeSemantics,omitempty"`
	CustomSemanticsActions duit_utils.TBool `json:"customSemanticsActions,omitempty"`
}

func NewSemanticsAttributes() *SemanticsAttributes {
	return &SemanticsAttributes{}
}

// SetLabel sets the label property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetLabel(label string) *SemanticsAttributes {
	r.Label = label
	return r
}

// SetHint sets the hint property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetHint(hint string) *SemanticsAttributes {
	r.Hint = hint
	return r
}

// SetValue sets the value property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetValue(value string) *SemanticsAttributes {
	r.Value = value
	return r
}

// SetTooltip sets the tooltip property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetTooltip(tooltip string) *SemanticsAttributes {
	r.Tooltip = tooltip
	return r
}

// SetEnabled sets the enabled property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetEnabled(enabled bool) *SemanticsAttributes {
	r.Enabled = duit_utils.BoolValue(enabled)
	return r
}

// SetChecked sets the checked property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetChecked(checked bool) *SemanticsAttributes {
	r.Checked = duit_utils.BoolValue(checked)
	return r
}

// SetSelected sets the selected property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetSelected(selected bool) *SemanticsAttributes {
	r.Selected = duit_utils.BoolValue(selected)
	return r
}

// SetButton sets the button property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetButton(button bool) *SemanticsAttributes {
	r.Button = duit_utils.BoolValue(button)
	return r
}

// SetLink sets the link property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetLink(link bool) *SemanticsAttributes {
	r.Link = duit_utils.BoolValue(link)
	return r
}

// SetHeader sets the header property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetHeader(header bool) *SemanticsAttributes {
	r.Header = duit_utils.BoolValue(header)
	return r
}

// SetTextField sets the textField property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetTextField(textField bool) *SemanticsAttributes {
	r.TextField = duit_utils.BoolValue(textField)
	return r
}

// SetImage sets the image property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetImage(image bool) *SemanticsAttributes {
	r.Image = duit_utils.BoolValue(image)
	return r
}

// SetLiveRegion sets the liveRegion property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetLiveRegion(liveRegion bool) *SemanticsAttributes {
	r.LiveRegion = duit_utils.BoolValue(liveRegion)
	return r
}

// SetContainer sets the container property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetContainer(container bool) *SemanticsAttributes {
	r.Container = duit_utils.BoolValue(container)
	return r
}

// SetExplicitChildNodes sets the explicitChildNodes property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetExplicitChildNodes(explicitChildNodes bool) *SemanticsAttributes {
	r.ExplicitChildNodes = duit_utils.BoolValue(explicitChildNodes)
	return r
}

// SetExcludeSemantics sets the excludeSemantics property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetExcludeSemantics(excludeSemantics bool) *SemanticsAttributes {
	r.ExcludeSemantics = duit_utils.BoolValue(excludeSemantics)
	return r
}

// SetCustomSemanticsActions sets the customSemanticsActions property and returns the SemanticsAttributes instance for method chaining.
func (r *SemanticsAttributes) SetCustomSemanticsActions(customSemanticsActions bool) *SemanticsAttributes {
	r.CustomSemanticsActions = duit_utils.BoolValue(customSemanticsActions)
	return r
}

func (r *SemanticsAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}
