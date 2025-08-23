package duit_attributes

type RepaintBoundaryAttributes struct {
	*ThemeConsumer
	ChildIndex int `json:"childIndex,omitempty"`
}

func (r *RepaintBoundaryAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}
