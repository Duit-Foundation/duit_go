package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_cross_axis"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_main_axis"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestWrapAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.WrapAttributes{
		TextDirection:      duit_props.TextDirectionLtr,
		VerticalDirection:  duit_props.VerticalDirectionDown,
		Alignment:          duit_main_axis.Start,
		RunAlignment:       duit_main_axis.Center,
		CrossAxisAlignment: duit_cross_axis.Start,
		Spacing:            10.0,
		RunSpacing:         5.0,
		ClipBehavior:       duit_props.ClipAntiAlias,
		Direction:          duit_props.AxisHorizontal,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}
