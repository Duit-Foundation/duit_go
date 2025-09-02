package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestAnimatedContainerAttributes_Validate_ValidAttributes(t *testing.T) {
	// Use concrete types that implement the interfaces
	attrs := &duit_attributes.AnimatedContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_props.ColorString]{
		Width:  100.0,
		Height: 200.0,
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

func TestAnimatedContainerAttributes_Validate_WithDimensions(t *testing.T) {
	testCases := []struct {
		name   string
		width  float32
		height float32
	}{
		{
			name:   "Positive dimensions",
			width:  150.0,
			height: 250.0,
		},
		{
			name:   "Zero dimensions",
			width:  0.0,
			height: 0.0,
		},
		{
			name:   "Only width",
			width:  100.0,
			height: 0.0,
		},
		{
			name:   "Only height",
			width:  0.0,
			height: 150.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_props.ColorString]{
				Width:  tc.width,
				Height: tc.height,
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

func TestAnimatedContainerAttributes_Validate_WithColorAndClip(t *testing.T) {
	attrs := &duit_attributes.AnimatedContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_props.ColorString]{
		Width:        100.0,
		Height:       200.0,
		Color:        "#FF0000",
		ClipBehavior: duit_props.ClipAntiAlias,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for attributes with color and clip, got:", err)
	}
}

func TestAnimatedContainerAttributes_Validate_WithPaddingAndMargin(t *testing.T) {
	attrs := &duit_attributes.AnimatedContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_props.ColorString]{
		Width:   100.0,
		Height:  200.0,
		Padding: 10.0,
		Margin:  5.0,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for attributes with padding and margin, got:", err)
	}
}

func TestAnimatedContainerAttributes_Validate_WithAlignment(t *testing.T) {
	attrs := &duit_attributes.AnimatedContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_props.ColorString]{
		Width:              100.0,
		Height:             200.0,
		Alignment:          duit_props.AlignmentCenter,
		TransformAlignment: duit_props.AlignmentTopLeft,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for attributes with alignment, got:", err)
	}
}

func TestAnimatedContainerAttributes_Validate_WithConstraints(t *testing.T) {
	constraints := &duit_flex.BoxConstraints{
		MinWidth:  50.0,
		MaxWidth:  200.0,
		MinHeight: 100.0,
		MaxHeight: 300.0,
	}

	attrs := &duit_attributes.AnimatedContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_props.ColorString]{
		Width:       100.0,
		Height:      200.0,
		Constraints: constraints,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for attributes with constraints, got:", err)
	}
}

// Tests for ImplicitAnimatable embedded struct

func TestAnimatedContainerAttributes_Validate_NilImplicitAnimatable(t *testing.T) {
	attrs := &duit_attributes.AnimatedContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_props.ColorString]{
		Width:              100.0,
		Height:             200.0,
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



func TestAnimatedContainerAttributes_MarshalJSON_WithAlignment(t *testing.T) {
	attrs := &duit_attributes.AnimatedContainerAttributes[duit_edge_insets.EdgeInsetsAll, duit_props.ColorString]{
		Width:              100.0,
		Height:             200.0,
		Alignment:          duit_props.AlignmentCenter,
		TransformAlignment: duit_props.AlignmentTopLeft,
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
	if !strings.Contains(jsonStr, `"alignment":"center"`) {
		t.Fatalf("expected JSON to contain '\"alignment\":\"center\"', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"transformAlignment":"topLeft"`) {
		t.Fatalf("expected JSON to contain '\"transformAlignment\":\"topLeft\"', got: %s", jsonStr)
	}
}