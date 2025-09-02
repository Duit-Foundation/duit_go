package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestAnimatedPhysicalModelAttributes_Validate_ValidAttributes(t *testing.T) {
	color := duit_color.ColorV2{}.FromString("#FF0000")
	shadowColor := duit_color.ColorV2{}.FromString("#000000")

	attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
		Elevation:   duit_utils.Float32Value(4.0),
		Color:       color,
		ShadowColor: shadowColor,
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

func TestAnimatedPhysicalModelAttributes_Validate_ValidElevationValues(t *testing.T) {
	testCases := []struct {
		name      string
		elevation float32
	}{
		{
			name:      "Zero elevation",
			elevation: 0.0,
		},
		{
			name:      "Small elevation",
			elevation: 2.0,
		},
		{
			name:      "Medium elevation",
			elevation: 8.0,
		},
		{
			name:      "Large elevation",
			elevation: 24.0,
		},
		{
			name:      "Very small elevation",
			elevation: 0.5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			color := duit_color.ColorV2{}.FromString("#FF0000")
			shadowColor := duit_color.ColorV2{}.FromString("#000000")

			attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
				Elevation:   duit_utils.Float32Value(tc.elevation),
				Color:       color,
				ShadowColor: shadowColor,
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

func TestAnimatedPhysicalModelAttributes_Validate_InvalidElevation(t *testing.T) {
	color := duit_color.ColorV2{}.FromString("#FF0000")
	shadowColor := duit_color.ColorV2{}.FromString("#000000")

	attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
		Elevation:   duit_utils.Float32Value(-1.0), // Invalid negative elevation
		Color:       color,
		ShadowColor: shadowColor,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for negative elevation")
	}

	expected := "elevation must be greater than 0"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestAnimatedPhysicalModelAttributes_Validate_NilElevation(t *testing.T) {
	color := duit_color.ColorV2{}.FromString("#FF0000")
	shadowColor := duit_color.ColorV2{}.FromString("#000000")

	attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
		Elevation:   nil, // Required field
		Color:       color,
		ShadowColor: shadowColor,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil elevation")
	}

	expected := "elevation property is required"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestAnimatedPhysicalModelAttributes_Validate_NilColor(t *testing.T) {
	shadowColor := duit_color.ColorV2{}.FromString("#000000")

	attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
		Elevation:   duit_utils.Float32Value(4.0),
		Color:       nil, // Required field
		ShadowColor: shadowColor,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil color")
	}

	expected := "color property is required"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestAnimatedPhysicalModelAttributes_Validate_NilShadowColor(t *testing.T) {
	color := duit_color.ColorV2{}.FromString("#FF0000")

	attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
		Elevation:   duit_utils.Float32Value(4.0),
		Color:       color,
		ShadowColor: nil, // Required field
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil shadow color")
	}

	expected := "shadowColor property is required"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestAnimatedPhysicalModelAttributes_Validate_WithOptionalFields(t *testing.T) {
	color := duit_color.ColorV2{}.FromString("#FF0000")
	shadowColor := duit_color.ColorV2{}.FromString("#000000")

	borderRadius := duit_decoration.BorderRadiusV2{}.All(
		duit_decoration.RadiusV2{}.Circular(8.0),
	)

	attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
		Elevation:          duit_utils.Float32Value(4.0),
		Color:              color,
		ShadowColor:        shadowColor,
		ClipBehavior:       duit_props.ClipAntiAlias,
		BorderRadius:       borderRadius,
		Shape:              duit_decoration.Rectangle,
		AnimateColor:       duit_utils.BoolValue(true),
		AnimateShadowColor: duit_utils.BoolValue(false),
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for attributes with optional fields, got:", err)
	}
}

// JSON Serialization tests for Tristate fields

func TestAnimatedPhysicalModelAttributes_MarshalJSON_ElevationField(t *testing.T) {
	color := duit_color.ColorV2{}.FromString("#FF0000")
	shadowColor := duit_color.ColorV2{}.FromString("#000000")

	attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
		Elevation:   duit_utils.Float32Value(6.0),
		Color:       color,
		ShadowColor: shadowColor,
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
	if !strings.Contains(jsonStr, `"elevation":6`) {
		t.Fatalf("expected JSON to contain '\"elevation\":6', got: %s", jsonStr)
	}
}

func TestAnimatedPhysicalModelAttributes_MarshalJSON_AnimateColorField(t *testing.T) {
	color := duit_color.ColorV2{}.FromString("#FF0000")
	shadowColor := duit_color.ColorV2{}.FromString("#000000")

	attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
		Elevation:    duit_utils.Float32Value(4.0),
		Color:        color,
		ShadowColor:  shadowColor,
		AnimateColor: duit_utils.BoolValue(true),
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
	if !strings.Contains(jsonStr, `"animateColor":true`) {
		t.Fatalf("expected JSON to contain '\"animateColor\":true', got: %s", jsonStr)
	}
}

func TestAnimatedPhysicalModelAttributes_MarshalJSON_AnimateShadowColorField(t *testing.T) {
	color := duit_color.ColorV2{}.FromString("#FF0000")
	shadowColor := duit_color.ColorV2{}.FromString("#000000")

	attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
		Elevation:          duit_utils.Float32Value(4.0),
		Color:              color,
		ShadowColor:        shadowColor,
		AnimateShadowColor: duit_utils.BoolValue(false),
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
	if !strings.Contains(jsonStr, `"animateShadowColor":false`) {
		t.Fatalf("expected JSON to contain '\"animateShadowColor\":false', got: %s", jsonStr)
	}
}

func TestAnimatedPhysicalModelAttributes_MarshalJSON_NilTristateFields(t *testing.T) {
	color := duit_color.ColorV2{}.FromString("#FF0000")
	shadowColor := duit_color.ColorV2{}.FromString("#000000")

	attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
		Elevation:          duit_utils.Float32Value(4.0),
		Color:              color,
		ShadowColor:        shadowColor,
		AnimateColor:       nil,
		AnimateShadowColor: nil,
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 300,
			Curve:    duit_animations.Ease,
		},
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() with nil tristate fields, got:", err)
	}

	jsonStr := string(jsonData)
	// Check that omitempty works - nil fields should not be present
	if strings.Contains(jsonStr, `"animateColor"`) {
		t.Fatalf("expected JSON to NOT contain 'animateColor' field for nil value, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"animateShadowColor"`) {
		t.Fatalf("expected JSON to NOT contain 'animateShadowColor' field for nil value, got: %s", jsonStr)
	}
}

// Tests for ImplicitAnimatable embedded struct

func TestAnimatedPhysicalModelAttributes_Validate_NilImplicitAnimatable(t *testing.T) {
	color := duit_color.ColorV2{}.FromString("#FF0000")
	shadowColor := duit_color.ColorV2{}.FromString("#000000")

	attrs := &duit_attributes.AnimatedPhysicalModelAttributes{
		Elevation:          duit_utils.Float32Value(4.0),
		Color:              color,
		ShadowColor:        shadowColor,
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