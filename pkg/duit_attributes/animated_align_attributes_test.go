package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
)

func TestAnimatedAlignAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AnimatedAlignAttributes{
		Alignment: duit_alignment.Center,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestAnimatedAlignAttributes_Validate_EmptyAlignment(t *testing.T) {
	attrs := &duit_attributes.AnimatedAlignAttributes{
		Alignment: "",
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for empty alignment")
	}

	expected := "alignment property is required"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestAnimatedAlignAttributes_Validate_DifferentAlignments(t *testing.T) {
	testCases := []struct {
		name      string
		alignment duit_alignment.Alignment
	}{
		{
			name:      "Center alignment",
			alignment: duit_alignment.Center,
		},
		{
			name:      "TopLeft alignment",
			alignment: duit_alignment.TopLeft,
		},
		{
			name:      "TopRight alignment",
			alignment: duit_alignment.TopRight,
		},
		{
			name:      "BottomLeft alignment",
			alignment: duit_alignment.BottomLeft,
		},
		{
			name:      "BottomRight alignment",
			alignment: duit_alignment.BottomRight,
		},
		{
			name:      "CenterLeft alignment",
			alignment: duit_alignment.CenterLeft,
		},
		{
			name:      "CenterRight alignment",
			alignment: duit_alignment.CenterRight,
		},
		{
			name:      "TopCenter alignment",
			alignment: duit_alignment.TopCenter,
		},
		{
			name:      "BottomCenter alignment",
			alignment: duit_alignment.BottomCenter,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedAlignAttributes{
				Alignment: tc.alignment,
				ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
					Duration: 300,
					Curve:    duit_animations.Ease,
				},
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for %s, got: %v", tc.name, err)
			}
		})
	}
}

func TestAnimatedAlignAttributes_Validate_WithFactors(t *testing.T) {
	attrs := &duit_attributes.AnimatedAlignAttributes{
		Alignment:    duit_alignment.Center,
		WidthFactor:  0.5,
		HeightFactor: 0.8,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for attributes with factors, got:", err)
	}
}

// Tests for ImplicitAnimatable embedded struct

func TestAnimatedAlignAttributes_Validate_NilImplicitAnimatable(t *testing.T) {
	attrs := &duit_attributes.AnimatedAlignAttributes{
		Alignment:          duit_alignment.Center,
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

func TestAnimatedAlignAttributes_MarshalJSON_WithFactors(t *testing.T) {
	attrs := &duit_attributes.AnimatedAlignAttributes{
		Alignment:    duit_alignment.Center,
		WidthFactor:  0.5,
		HeightFactor: 0.8,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"widthFactor":0.5`) {
		t.Fatalf("expected JSON to contain '\"widthFactor\":0.5', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"heightFactor":0.8`) {
		t.Fatalf("expected JSON to contain '\"heightFactor\":0.8', got: %s", jsonStr)
	}
}
