package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// InteractiveViewerAttributes represents attributes for InteractiveViewer widget.
// https://api.flutter.dev/flutter/widgets/InteractiveViewer-class.html
type InteractiveViewerAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	ClipBehavior                      duit_props.Clip        `json:"clipBehavior,omitempty"`
	BoundaryMargin                    *duit_props.EdgeInsets `json:"boundaryMargin,omitempty"`
	MaxScale                          float32                `json:"maxScale,omitempty"`
	MinScale                          float32                `json:"minScale,omitempty"`
	InteractionEndFrictionCoefficient float32                `json:"interactionEndFrictionCoefficient,omitempty"`
	OnInteractionEnd                  any                    `json:"onInteractionEnd,omitempty"`
	OnInteractionStart                any                    `json:"onInteractionStart,omitempty"`
	OnInteractionUpdate               any                    `json:"onInteractionUpdate,omitempty"`
	PanEnabled                        duit_utils.TBool       `json:"panEnabled,omitempty"`
	ScaleEnabled                      duit_utils.TBool       `json:"scaleEnabled,omitempty"`
	ScaleFactor                       float32                `json:"scaleFactor,omitempty"`
	Alignment                         duit_props.Alignment   `json:"alignment,omitempty"`
	TrackpadScrollCausesScale         duit_utils.TBool       `json:"trackpadScrollCausesScale,omitempty"`
	Constrained                       duit_utils.TBool       `json:"constrained,omitempty"`
}

// NewInteractiveViewerAttributes creates a new instance of InteractiveViewerAttributes.
func NewInteractiveViewerAttributes() *InteractiveViewerAttributes {
	return &InteractiveViewerAttributes{}
}

// SetClipBehavior sets the clipBehavior property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetClipBehavior(clipBehavior duit_props.Clip) *InteractiveViewerAttributes {
	r.ClipBehavior = clipBehavior
	return r
}

// SetBoundaryMargin sets the boundaryMargin property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetBoundaryMargin(boundaryMargin *duit_props.EdgeInsets) *InteractiveViewerAttributes {
	r.BoundaryMargin = boundaryMargin
	return r
}

// SetMaxScale sets the maxScale property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetMaxScale(maxScale float32) *InteractiveViewerAttributes {
	r.MaxScale = maxScale
	return r
}

// SetMinScale sets the minScale property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetMinScale(minScale float32) *InteractiveViewerAttributes {
	r.MinScale = minScale
	return r
}

// SetInteractionEndFrictionCoefficient sets the interactionEndFrictionCoefficient property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetInteractionEndFrictionCoefficient(interactionEndFrictionCoefficient float32) *InteractiveViewerAttributes {
	r.InteractionEndFrictionCoefficient = interactionEndFrictionCoefficient
	return r
}

// SetOnInteractionEnd sets the onInteractionEnd property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetOnInteractionEnd(onInteractionEnd any) *InteractiveViewerAttributes {
	r.OnInteractionEnd = onInteractionEnd
	return r
}

// SetOnInteractionStart sets the onInteractionStart property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetOnInteractionStart(onInteractionStart any) *InteractiveViewerAttributes {
	r.OnInteractionStart = onInteractionStart
	return r
}

// SetOnInteractionUpdate sets the onInteractionUpdate property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetOnInteractionUpdate(onInteractionUpdate any) *InteractiveViewerAttributes {
	r.OnInteractionUpdate = onInteractionUpdate
	return r
}

// SetPanEnabled sets the panEnabled property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetPanEnabled(panEnabled bool) *InteractiveViewerAttributes {
	r.PanEnabled = duit_utils.BoolValue(panEnabled)
	return r
}

// SetScaleEnabled sets the scaleEnabled property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetScaleEnabled(scaleEnabled bool) *InteractiveViewerAttributes {
	r.ScaleEnabled = duit_utils.BoolValue(scaleEnabled)
	return r
}

// SetScaleFactor sets the scaleFactor property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetScaleFactor(scaleFactor float32) *InteractiveViewerAttributes {
	r.ScaleFactor = scaleFactor
	return r
}

// SetAlignment sets the alignment property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetAlignment(alignment duit_props.Alignment) *InteractiveViewerAttributes {
	r.Alignment = alignment
	return r
}

// SetTrackpadScrollCausesScale sets the trackpadScrollCausesScale property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetTrackpadScrollCausesScale(trackpadScrollCausesScale bool) *InteractiveViewerAttributes {
	r.TrackpadScrollCausesScale = duit_utils.BoolValue(trackpadScrollCausesScale)
	return r
}

// SetConstrained sets the constrained property and returns the InteractiveViewerAttributes instance for method chaining.
func (r *InteractiveViewerAttributes) SetConstrained(constrained bool) *InteractiveViewerAttributes {
	r.Constrained = duit_utils.BoolValue(constrained)
	return r
}

//bindgen:exclude
func (r *InteractiveViewerAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	callbacks := []any{
		r.OnInteractionEnd,
		r.OnInteractionStart,
		r.OnInteractionUpdate,
	}

	for _, callback := range callbacks {
		if callback != nil {
			if err := duit_action.CheckActionType(callback); err != nil {
				return err
			}
		}
	}

	return nil
}
