package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestStackAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.StackAttributes{
		TextDirection: duit_props.TextDirectionLtr,
		ClipBehavior:  duit_props.ClipAntiAlias,
		Alignment:     duit_props.AlignmentCenter,
		Fit:           duit_flex.Loose,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}
