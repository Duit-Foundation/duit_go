package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	decorations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestSliverAppBarAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverAppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		BottomOpacity:  duit_utils.Float32Value(0.5),
		ToolbarOpacity: duit_utils.Float32Value(0.8),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverAppBarAttributes_Validate_ValidOpacityValues(t *testing.T) {
	testCases := []struct {
		name    string
		opacity float32
	}{
		{"Zero opacity", 0.0},
		{"Low opacity", 0.1},
		{"Half opacity", 0.5},
		{"High opacity", 0.9},
		{"Full opacity", 1.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"_BottomOpacity", func(t *testing.T) {
			attrs := &duit_attributes.SliverAppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, decorations.RoundedRectangleBorder[duit_props.ColorString]]{
				BottomOpacity: duit_utils.Float32Value(tc.opacity),
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for bottomOpacity %f, got: %v", tc.opacity, err)
			}
		})

		t.Run(tc.name+"_ToolbarOpacity", func(t *testing.T) {
			attrs := &duit_attributes.SliverAppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, decorations.RoundedRectangleBorder[duit_props.ColorString]]{
				ToolbarOpacity: duit_utils.Float32Value(tc.opacity),
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for toolbarOpacity %f, got: %v", tc.opacity, err)
			}
		})
	}
}

func TestSliverAppBarAttributes_Validate_InvalidOpacityValues(t *testing.T) {
	testCases := []struct {
		name    string
		opacity float32
	}{
		{"Negative opacity", -0.1},
		{"Very high opacity", 2.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name+"_BottomOpacity", func(t *testing.T) {
			attrs := &duit_attributes.SliverAppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, decorations.RoundedRectangleBorder[duit_props.ColorString]]{
				BottomOpacity: duit_utils.Float32Value(tc.opacity),
			}

			err := attrs.Validate()
			if err == nil {
				t.Fatalf("expected error for bottomOpacity %f", tc.opacity)
			}
		})

		t.Run(tc.name+"_ToolbarOpacity", func(t *testing.T) {
			attrs := &duit_attributes.SliverAppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, decorations.RoundedRectangleBorder[duit_props.ColorString]]{
				ToolbarOpacity: duit_utils.Float32Value(tc.opacity),
			}

			err := attrs.Validate()
			if err == nil {
				t.Fatalf("expected error for toolbarOpacity %f", tc.opacity)
			}
		})
	}
}
