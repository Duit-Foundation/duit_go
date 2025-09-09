package duit_attributes

// RepaintBoundaryAttributes defines attributes for RepaintBoundary widget.
// See: https://api.flutter.dev/flutter/widgets/RepaintBoundary-class.html
type RepaintBoundaryAttributes struct {
	*ThemeConsumer `gen:"wrap"`
	ChildIndex     int `json:"childIndex,omitempty"`
}

//bindgen:exclude
func (r *RepaintBoundaryAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	return nil
}

// NewRepaintBoundaryAttributes creates a new RepaintBoundaryAttributes instance.
func NewRepaintBoundaryAttributes() *RepaintBoundaryAttributes {
	return &RepaintBoundaryAttributes{}
}

// SetChildIndex sets the child index for the repaint boundary widget.
func (r *RepaintBoundaryAttributes) SetChildIndex(childIndex int) *RepaintBoundaryAttributes {
	r.ChildIndex = childIndex
	return r
}
