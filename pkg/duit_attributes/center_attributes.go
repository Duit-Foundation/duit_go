package duit_attributes

type CenterAttributes struct {
	*ValueReferenceHolder
	*ThemeConsumer
	WidgthFactor float32 `json:"widthFactor,omitempty"`
	HeightFactor float32 `json:"heightFactor,omitempty"`
}

func (r *CenterAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}
	
	return nil
}
