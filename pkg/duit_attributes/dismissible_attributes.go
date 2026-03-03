package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_action"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// DismissibleAttributes defines attributes for Dismissible widget.
// See: https://api.flutter.dev/flutter/widgets/Dismissible-class.html
//
// Children slots (multichild widget):
//   - children[0]: child (required) - main content
//   - children[1]: background (optional) - shown when swiping in primary direction
//   - children[2]: secondaryBackground (optional) - shown when swiping in secondary direction
type DismissibleAttributes struct {
	*ValueReferenceHolder `gen:"wrap"`
	*ThemeConsumer        `gen:"wrap"`
	Direction             duit_props.DismissDirection             `json:"direction,omitempty"`
	ResizeDuration        duit_utils.Tristate[uint]               `json:"resizeDuration,omitempty"`
	MovementDuration      duit_utils.Tristate[uint]               `json:"movementDuration,omitempty"`
	CrossAxisEndOffset    duit_utils.Tristate[float32]            `json:"crossAxisEndOffset,omitempty"`
	DismissThresholds     map[duit_props.DismissDirection]float32 `json:"dismissThresholds,omitempty"`
	HitTestBehavior       duit_props.HitTestBehavior              `json:"hitTestBehavior,omitempty"`
	DragStartBehavior     duit_props.DragStartBehavior            `json:"dragStartBehavior,omitempty"`
	OnDismissed           any                                     `json:"onDismissed,omitempty"`
	OnResize              any                                     `json:"onResize,omitempty"`
	OnUpdate              any                                     `json:"onUpdate,omitempty"`
}

// Validate validates the DismissibleAttributes.
//
//bindgen:exclude
func (r *DismissibleAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if r.OnDismissed != nil {
		if err := duit_action.CheckActionType(r.OnDismissed); err != nil {
			return err
		}
	}

	if r.OnResize != nil {
		if err := duit_action.CheckActionType(r.OnResize); err != nil {
			return err
		}
	}

	if r.OnUpdate != nil {
		if err := duit_action.CheckActionType(r.OnUpdate); err != nil {
			return err
		}
	}

	return nil
}

// NewDismissibleAttributes creates a new DismissibleAttributes instance.
func NewDismissibleAttributes() *DismissibleAttributes {
	return &DismissibleAttributes{}
}

// SetDirection sets the direction property and returns DismissibleAttributes for method chaining.
func (r *DismissibleAttributes) SetDirection(direction duit_props.DismissDirection) *DismissibleAttributes {
	r.Direction = direction
	return r
}

// SetResizeDuration sets the resizeDuration property (milliseconds) and returns DismissibleAttributes for method chaining.
func (r *DismissibleAttributes) SetResizeDuration(resizeDuration uint) *DismissibleAttributes {
	r.ResizeDuration = duit_utils.UintValue(resizeDuration)
	return r
}

// SetMovementDuration sets the movementDuration property (milliseconds) and returns DismissibleAttributes for method chaining.
func (r *DismissibleAttributes) SetMovementDuration(movementDuration uint) *DismissibleAttributes {
	r.MovementDuration = duit_utils.UintValue(movementDuration)
	return r
}

// SetCrossAxisEndOffset sets the crossAxisEndOffset property and returns DismissibleAttributes for method chaining.
func (r *DismissibleAttributes) SetCrossAxisEndOffset(crossAxisEndOffset float32) *DismissibleAttributes {
	r.CrossAxisEndOffset = duit_utils.Float32Value(crossAxisEndOffset)
	return r
}

// SetDismissThresholds sets the dismissThresholds property and returns DismissibleAttributes for method chaining.
// Keys: vertical, horizontal, up, down, startToEnd, endToStart, none. Values are fractional offset (0.0-1.0).
func (r *DismissibleAttributes) SetDismissThresholds(dismissThresholds map[string]float32) *DismissibleAttributes {
	r.DismissThresholds = dismissThresholds
	return r
}

// SetHitTestBehavior sets the hitTestBehavior property and returns DismissibleAttributes for method chaining.
func (r *DismissibleAttributes) SetHitTestBehavior(value duit_props.HitTestBehavior) *DismissibleAttributes {
	r.HitTestBehavior = value
	return r
}

// SetDragStartBehavior sets the dragStartBehavior property and returns DismissibleAttributes for method chaining.
func (r *DismissibleAttributes) SetDragStartBehavior(value duit_props.DragStartBehavior) *DismissibleAttributes {
	r.DragStartBehavior = value
	return r
}
