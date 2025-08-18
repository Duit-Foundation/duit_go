package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
)

func TestAlignAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AlignAttributes{
		Alignment: duit_alignment.Center,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestAlignAttributes_Validate_EmptyAlignment(t *testing.T) {
	attrs := &duit_attributes.AlignAttributes{
		Alignment: "",
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

func TestAlignAttributes_Validate_DifferentAlignments(t *testing.T) {
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
			attrs := &duit_attributes.AlignAttributes{
				Alignment: tc.alignment,
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for %s, got: %v", tc.name, err)
			}
		})
	}
}

func TestAlignAttributes_Validate_WithFactors(t *testing.T) {
	testCases := []struct {
		name         string
		widthFactor  float32
		heightFactor float32
	}{
		{
			name:         "Both factors set",
			widthFactor:  0.5,
			heightFactor: 0.8,
		},
		{
			name:         "Only width factor",
			widthFactor:  0.7,
			heightFactor: 0.0,
		},
		{
			name:         "Only height factor",
			widthFactor:  0.0,
			heightFactor: 0.6,
		},
		{
			name:         "Zero factors",
			widthFactor:  0.0,
			heightFactor: 0.0,
		},
		{
			name:         "Factor of 1",
			widthFactor:  1.0,
			heightFactor: 1.0,
		},
		{
			name:         "Large factors",
			widthFactor:  2.0,
			heightFactor: 1.5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AlignAttributes{
				Alignment:    duit_alignment.Center,
				WidthFactor:  tc.widthFactor,
				HeightFactor: tc.heightFactor,
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for %s, got: %v", tc.name, err)
			}
		})
	}
}

func TestAlignAttributes_Validate_WithAnimatedPropertyOwner(t *testing.T) {
	attrs := &duit_attributes.AlignAttributes{
		Alignment: duit_alignment.Center,
		AnimatedPropertyOwner: &duit_animations.AnimatedPropertyOwner{
			ParentBuilderId:    "test-builder",
			AffectedProperties: []string{"alignment"},
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for attributes with AnimatedPropertyOwner, got:", err)
	}
}

func TestAlignAttributes_Validate_WithThemeConsumer(t *testing.T) {
	attrs := &duit_attributes.AlignAttributes{
		Alignment: duit_alignment.Center,
		ThemeConsumer: &duit_attributes.ThemeConsumer{
			Theme: "primaryTheme",
		},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for attributes with ThemeConsumer, got:", err)
	}
}

func TestAlignAttributes_MarshalJSON_WithFactors(t *testing.T) {
	attrs := &duit_attributes.AlignAttributes{
		Alignment:    duit_alignment.Center,
		WidthFactor:  0.5,
		HeightFactor: 0.8,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"alignment":"center"`) {
		t.Fatalf("expected JSON to contain '\"alignment\":\"center\"', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"widthFactor":0.5`) {
		t.Fatalf("expected JSON to contain '\"widthFactor\":0.5', got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"heightFactor":0.8`) {
		t.Fatalf("expected JSON to contain '\"heightFactor\":0.8', got: %s", jsonStr)
	}
}

func TestAlignAttributes_MarshalJSON_ZeroFactors(t *testing.T) {
	attrs := &duit_attributes.AlignAttributes{
		Alignment:    duit_alignment.TopLeft,
		WidthFactor:  0.0,
		HeightFactor: 0.0,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"alignment":"topLeft"`) {
		t.Fatalf("expected JSON to contain '\"alignment\":\"topLeft\"', got: %s", jsonStr)
	}
	// Zero values should be omitted due to omitempty tag
	if strings.Contains(jsonStr, `"widthFactor"`) {
		t.Fatalf("expected JSON to not contain 'widthFactor' field when zero, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"heightFactor"`) {
		t.Fatalf("expected JSON to not contain 'heightFactor' field when zero, got: %s", jsonStr)
	}
}

func TestAlignAttributes_UnmarshalJSON_WithFactors(t *testing.T) {
	jsonStr := `{"alignment":"center","widthFactor":0.5,"heightFactor":0.8}`

	var attrs duit_attributes.AlignAttributes
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.Alignment != duit_alignment.Center {
		t.Fatalf("expected alignment to be 'center', got: %s", attrs.Alignment)
	}
	if attrs.WidthFactor != 0.5 {
		t.Fatalf("expected widthFactor to be 0.5, got: %f", attrs.WidthFactor)
	}
	if attrs.HeightFactor != 0.8 {
		t.Fatalf("expected heightFactor to be 0.8, got: %f", attrs.HeightFactor)
	}
}

func TestAlignAttributes_UnmarshalJSON_WithoutFactors(t *testing.T) {
	jsonStr := `{"alignment":"bottomRight"}`

	var attrs duit_attributes.AlignAttributes
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.Alignment != duit_alignment.BottomRight {
		t.Fatalf("expected alignment to be 'bottomRight', got: %s", attrs.Alignment)
	}
	if attrs.WidthFactor != 0.0 {
		t.Fatalf("expected widthFactor to be 0.0 when not provided, got: %f", attrs.WidthFactor)
	}
	if attrs.HeightFactor != 0.0 {
		t.Fatalf("expected heightFactor to be 0.0 when not provided, got: %f", attrs.HeightFactor)
	}
}
