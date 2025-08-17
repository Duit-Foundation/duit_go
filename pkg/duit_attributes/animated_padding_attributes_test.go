package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
)

func TestAnimatedPaddingAttributes_Validate_ValidPadding(t *testing.T) {
	padding := duit_edge_insets.EdgeInsetsV2{}.All(10.0)
	
	attrs := &duit_attributes.AnimatedPaddingAttributes{
		Padding: padding,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid padding, got:", err)
	}
}

func TestAnimatedPaddingAttributes_Validate_NilPadding(t *testing.T) {
	attrs := &duit_attributes.AnimatedPaddingAttributes{
		Padding: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil padding")
	}

	expected := "padding property is required"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestAnimatedPaddingAttributes_Validate_InvalidPadding(t *testing.T) {
	// Create invalid padding with negative values
	padding := duit_edge_insets.EdgeInsetsV2{}.All(-5.0)
	
	attrs := &duit_attributes.AnimatedPaddingAttributes{
		Padding: padding,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for invalid padding")
	}
}

func TestAnimatedPaddingAttributes_Validate_DifferentPaddingTypes(t *testing.T) {
	testCases := []struct {
		name    string
		padding *duit_edge_insets.EdgeInsetsV2
	}{
		{
			name:    "All padding",
			padding: duit_edge_insets.EdgeInsetsV2{}.All(8.0),
		},
		{
			name:    "LTRB padding", 
			padding: duit_edge_insets.EdgeInsetsV2{}.LTRB(5.0, 10.0, 5.0, 10.0),
		},
		{
			name:    "Symmetric padding",
			padding: duit_edge_insets.EdgeInsetsV2{}.Symmentric(12.0, 8.0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.AnimatedPaddingAttributes{
				Padding: tc.padding,
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for %s, got: %v", tc.name, err)
			}
		})
	}
}