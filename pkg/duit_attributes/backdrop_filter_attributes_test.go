package duit_attributes_test

import (
	"strings"
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestBackdropFilterAttributes_Validate_ValidAttributes(t *testing.T) {
	blurFilter := duit_props.NewBlurImageFilter(5.0, 5.0, duit_props.TileModeClamp)
	attrs := &duit_attributes.BackdropFilterAttributes[duit_props.BlurImageFilter]{
		Filter:    blurFilter,
		BlendMode: duit_props.BlendModeSrcOver,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestBackdropFilterAttributes_Validate_MissingFilter(t *testing.T) {
	attrs := &duit_attributes.BackdropFilterAttributes[duit_props.BlurImageFilter]{
		Filter:    nil,
		BlendMode: duit_props.BlendModeSrcOver,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing filter")
	}

	if !strings.Contains(err.Error(), "filter property is required") {
		t.Fatalf("expected error message about required filter, got: %s", err.Error())
	}
}

func TestBackdropFilterAttributes_Validate_MissingBlendMode(t *testing.T) {
	blurFilter := duit_props.NewBlurImageFilter(5.0, 5.0, duit_props.TileModeClamp)
	attrs := &duit_attributes.BackdropFilterAttributes[duit_props.BlurImageFilter]{
		Filter:    blurFilter,
		BlendMode: "",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing blend mode")
	}

	if !strings.Contains(err.Error(), "blendMode property is required") {
		t.Fatalf("expected error message about required blendMode, got: %s", err.Error())
	}
}

func TestBackdropFilterAttributes_Validate_MissingBoth(t *testing.T) {
	attrs := &duit_attributes.BackdropFilterAttributes[duit_props.BlurImageFilter]{
		Filter:    nil,
		BlendMode: "",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing both filter and blend mode")
	}

	// Should fail on filter first
	if !strings.Contains(err.Error(), "filter property is required") {
		t.Fatalf("expected error message about required filter, got: %s", err.Error())
	}
}
