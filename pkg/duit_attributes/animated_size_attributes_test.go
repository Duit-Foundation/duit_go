package duit_attributes_test

import (
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestAnimatedSizeAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AnimatedSizeAttributes{
		ClipBehavior: duit_props.ClipAntiAlias,
		Alignment:    duit_props.AlignmentCenter,
		ImplicitAnimatable: &duit_props.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_props.CurveEase,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

// Tests for ImplicitAnimatable embedded struct

func TestAnimatedSizeAttributes_Validate_NilImplicitAnimatable(t *testing.T) {
	attrs := &duit_attributes.AnimatedSizeAttributes{
		ClipBehavior:       duit_props.ClipAntiAlias,
		Alignment:          duit_props.AlignmentCenter,
		ImplicitAnimatable: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil ImplicitAnimatable")
	}

	expected := "implicitAnimatable property is required on implicit animated widgets"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}
