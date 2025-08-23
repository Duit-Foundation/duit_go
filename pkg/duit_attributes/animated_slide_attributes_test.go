package duit_attributes_test

import (
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
)

func TestAnimatedSlideAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AnimatedSlideAttributes{
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
		Offset: &duit_flex.Offset{
			Dx: 0.5,
			Dy: 1.0,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestAnimatedSlideAttributes_Validate_MissingImplicitAnimatable(t *testing.T) {
	attrs := &duit_attributes.AnimatedSlideAttributes{
		ImplicitAnimatable: nil,
		Offset: &duit_flex.Offset{
			Dx: 0.5,
			Dy: 1.0,
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing ImplicitAnimatable")
	}

	if !strings.Contains(err.Error(), "implicitAnimatable property is required on implicit animated widgets") {
		t.Fatalf("expected error message about required ImplicitAnimatable, got: %s", err.Error())
	}
}

func TestAnimatedSlideAttributes_Validate_MissingOffset(t *testing.T) {
	attrs := &duit_attributes.AnimatedSlideAttributes{
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
		Offset: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing Offset")
	}

	if !strings.Contains(err.Error(), "offset property is required") {
		t.Fatalf("expected error message about required Offset, got: %s", err.Error())
	}
}

func TestAnimatedSlideAttributes_Validate_MissingBoth(t *testing.T) {
	attrs := &duit_attributes.AnimatedSlideAttributes{
		ImplicitAnimatable: nil,
		Offset:             nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing both ImplicitAnimatable and Offset")
	}

	// Should fail on ImplicitAnimatable first
	if !strings.Contains(err.Error(), "implicitAnimatable property is required on implicit animated widgets") {
		t.Fatalf("expected error message about required ImplicitAnimatable, got: %s", err.Error())
	}
}

func TestAnimatedSlideAttributes_Validate_InvalidImplicitAnimatable(t *testing.T) {
	attrs := &duit_attributes.AnimatedSlideAttributes{
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 0, // Invalid duration
			Curve:    duit_animations.Ease,
		},
		Offset: &duit_flex.Offset{
			Dx: 0.5,
			Dy: 1.0,
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for invalid ImplicitAnimatable")
	}

	if !strings.Contains(err.Error(), "duration property is required and must be greater than 0") {
		t.Fatalf("expected error message about invalid duration, got: %s", err.Error())
	}
}
