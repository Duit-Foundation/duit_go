package duit_attributes_test

import (
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestOverflowBoxAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.OverflowBoxAttributes{
		MinWidth:  100.0,
		MaxWidth:  200.0,
		MinHeight: 150.0,
		MaxHeight: 250.0,
		Alignment: duit_props.AlignmentCenter,
		Fit:       duit_props.OverflowBoxFitMax,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}
