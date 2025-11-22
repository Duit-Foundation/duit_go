package duit_attributes

import (
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type BadgeVariant = uint8

const (
	BadgeCommon = iota
	BadgeCount
)

type BadgeAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*duit_props.AnimatedPropertyOwner `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Variant                           duit_utils.Tristate[BadgeVariant] `json:"variant,omitempty"`
	BackgroundColor                   *duit_props.Color                 `json:"backgroundColor,omitempty"`
	TextColor                         *duit_props.Color                 `json:"textColor,omitempty"`
	SmallSize                         duit_utils.Tristate[float32]      `json:"smallSize,omitempty"`
	LargeSize                         duit_utils.Tristate[float32]      `json:"largeSize,omitempty"`
	Padding                           *duit_props.EdgeInsets            `json:"padding,omitempty"`
	Alignment                         *duit_props.Alignment             `json:"alignment,omitempty"`
	Offset                            *duit_props.Offset                `json:"offset,omitempty"`
	IsLabelVisible                    duit_utils.TBool                  `json:"isLabelVisible,omitempty"`
	Count                             duit_utils.Tristate[uint]         `json:"count,omitempty"`
}

//bindgen:exclude
func (r *BadgeAttributes) Validate() error {
	if err := r.ThemeConsumer.Validate(); err != nil {
		return err
	}

	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	if err := r.AnimatedPropertyOwner.Validate(); err != nil {
		return err
	}

	if r.Padding != nil {
		if err := r.Padding.Validate(); err != nil {
			return err
		}
	}

	if r.Variant != nil {
		if *r.Variant == BadgeCount && r.Count == nil {
			return errors.New("count property can`t be ")
		}
	}

	return nil
}

// NewBadgeAttributes creates a new BadgeAttributes instance.
func NewBadgeAttributes() *BadgeAttributes {
	return &BadgeAttributes{}
}

// SetVariant sets the variant property and returns the BadgeAttributes instance for method chaining.
func (r *BadgeAttributes) SetVariant(variant BadgeVariant) *BadgeAttributes {
	r.Variant = duit_utils.TristateFrom[BadgeVariant](variant)
	return r
}

// SetBackgroundColor sets the backgroundColor property and returns the BadgeAttributes instance for method chaining.
func (r *BadgeAttributes) SetBackgroundColor(backgroundColor *duit_props.Color) *BadgeAttributes {
	r.BackgroundColor = backgroundColor
	return r
}

// SetTextColor sets the textColor property and returns the BadgeAttributes instance for method chaining.
func (r *BadgeAttributes) SetTextColor(textColor *duit_props.Color) *BadgeAttributes {
	r.TextColor = textColor
	return r
}

// SetSmallSize sets the smallSize property and returns the BadgeAttributes instance for method chaining.
func (r *BadgeAttributes) SetSmallSize(smallSize float32) *BadgeAttributes {
	r.SmallSize = duit_utils.Float32Value(smallSize)
	return r
}

// SetLargeSize sets the largeSize property and returns the BadgeAttributes instance for method chaining.
func (r *BadgeAttributes) SetLargeSize(largeSize float32) *BadgeAttributes {
	r.LargeSize = duit_utils.Float32Value(largeSize)
	return r
}

// SetPadding sets the padding property and returns the BadgeAttributes instance for method chaining.
func (r *BadgeAttributes) SetPadding(padding *duit_props.EdgeInsets) *BadgeAttributes {
	r.Padding = padding
	return r
}

// SetAlignment sets the alignment property and returns the BadgeAttributes instance for method chaining.
func (r *BadgeAttributes) SetAlignment(alignment *duit_props.Alignment) *BadgeAttributes {
	r.Alignment = alignment
	return r
}

// SetOffset sets the offset property and returns the BadgeAttributes instance for method chaining.
func (r *BadgeAttributes) SetOffset(offset *duit_props.Offset) *BadgeAttributes {
	r.Offset = offset
	return r
}

// SetIsLabelVisible sets the isLabelVisible property and returns the BadgeAttributes instance for method chaining.
func (r *BadgeAttributes) SetIsLabelVisible(isLabelVisible bool) *BadgeAttributes {
	r.IsLabelVisible = duit_utils.BoolValue(isLabelVisible)
	return r
}

// SetCount sets the count property and returns the BadgeAttributes instance for method chaining.
func (r *BadgeAttributes) SetCount(count uint) *BadgeAttributes {
	r.Count = duit_utils.UintValue(count)
	return r
}
