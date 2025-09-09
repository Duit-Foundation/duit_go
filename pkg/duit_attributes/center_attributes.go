package duit_attributes

// CenterAttributes represents attributes for Center widget.
// https://api.flutter.dev/flutter/widgets/Center-class.html
type CenterAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	WidthFactor           float32 `json:"widthFactor,omitempty"`
	HeightFactor          float32 `json:"heightFactor,omitempty"`
}

// NewCenterAttributes creates a new instance of CenterAttributes.
func NewCenterAttributes() *CenterAttributes {
	return &CenterAttributes{}
}

// SetWidthFactor sets the widthFactor property and returns the CenterAttributes instance for method chaining.
func (r *CenterAttributes) SetWidthFactor(widthFactor float32) *CenterAttributes {
	r.WidthFactor = widthFactor
	return r
}

// SetHeightFactor sets the heightFactor property and returns the CenterAttributes instance for method chaining.
func (r *CenterAttributes) SetHeightFactor(heightFactor float32) *CenterAttributes {
	r.HeightFactor = heightFactor
	return r
}

//bindgen:exclude
func (r *CenterAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}
	
	return nil
}
