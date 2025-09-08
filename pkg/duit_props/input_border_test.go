package duit_props_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for NewInputBorder constructor
func TestNewInputBorder(t *testing.T) {
	border := duit_props.NewInputBorder("outline")

	if border == nil {
		t.Fatal("expected non-nil InputBorder")
	}

	if border.Options == nil {
		t.Fatal("expected non-nil Options")
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for valid input border, got:", err)
	}
}

func TestNewInputBorder_EmptyType(t *testing.T) {
	border := duit_props.NewInputBorder("")

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for empty border type, got:", err)
	}
}

// Tests for SetGapPadding method
func TestInputBorder_SetGapPadding(t *testing.T) {
	border := duit_props.NewInputBorder("outline")
	border = border.SetGapPadding(5.0)

	if border.Options.GapPadding != 5.0 {
		t.Fatalf("expected GapPadding 5.0, got %f", border.Options.GapPadding)
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for valid gap padding, got:", err)
	}
}

func TestInputBorder_SetGapPadding_Zero(t *testing.T) {
	border := duit_props.NewInputBorder("outline")
	border = border.SetGapPadding(0.0)

	if border.Options.GapPadding != 0.0 {
		t.Fatalf("expected GapPadding 0.0, got %f", border.Options.GapPadding)
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for zero gap padding, got:", err)
	}
}

func TestInputBorder_SetGapPadding_Negative(t *testing.T) {
	border := duit_props.NewInputBorder("outline")
	border = border.SetGapPadding(-2.0)

	if border.Options.GapPadding != -2.0 {
		t.Fatalf("expected GapPadding -2.0, got %f", border.Options.GapPadding)
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for negative gap padding, got:", err)
	}
}

// Tests for SetBorderRadius method
func TestInputBorder_SetBorderRadius(t *testing.T) {
	border := duit_props.NewInputBorder("outline")
	border = border.SetBorderRadius(10.0)

	if border.Options.BorderRadius != 10.0 {
		t.Fatalf("expected BorderRadius 10.0, got %f", border.Options.BorderRadius)
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for valid border radius, got:", err)
	}
}

func TestInputBorder_SetBorderRadius_Zero(t *testing.T) {
	border := duit_props.NewInputBorder("outline")
	border = border.SetBorderRadius(0.0)

	if border.Options.BorderRadius != 0.0 {
		t.Fatalf("expected BorderRadius 0.0, got %f", border.Options.BorderRadius)
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for zero border radius, got:", err)
	}
}

// Tests for SetBorderSide method
func TestInputBorder_SetBorderSide(t *testing.T) {
	border := duit_props.NewInputBorder("outline")
	borderSide := duit_props.NewBorderSide().SetWidth(2.0)
	border = border.SetBorderSide(borderSide)

	if border.Options.BorderSide != borderSide {
		t.Fatal("expected BorderSide to be set")
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for valid border side, got:", err)
	}
}

func TestInputBorder_SetBorderSide_Nil(t *testing.T) {
	border := duit_props.NewInputBorder("outline")
	border = border.SetBorderSide(nil)

	if border.Options.BorderSide != nil {
		t.Fatal("expected BorderSide to be nil")
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for nil border side, got:", err)
	}
}

// Tests for InputBorder.Validate() method
func TestInputBorder_Validate_NilOptions(t *testing.T) {
	border := &duit_props.InputBorder{
		Type:    "outline",
		Options: nil,
	}

	err := border.Validate()
	if err == nil {
		t.Fatal("expected error for nil options")
	}
}

func TestInputBorder_Validate_NilReceiver(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on nil receiver, but did not panic")
		}
	}()

	var border *duit_props.InputBorder
	border.Validate()
}

// Tests for method chaining
func TestInputBorder_MethodChaining(t *testing.T) {
	borderSide := duit_props.NewBorderSide().SetWidth(1.5)
	border := duit_props.NewInputBorder("outline").
		SetGapPadding(4.0).
		SetBorderRadius(8.0).
		SetBorderSide(borderSide)

	if border.Options.GapPadding != 4.0 {
		t.Fatalf("expected GapPadding 4.0, got %f", border.Options.GapPadding)
	}

	if border.Options.BorderRadius != 8.0 {
		t.Fatalf("expected BorderRadius 8.0, got %f", border.Options.BorderRadius)
	}

	if border.Options.BorderSide != borderSide {
		t.Fatal("expected BorderSide to be set")
	}

	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for chained method calls, got:", err)
	}
}

// Tests for InputBorderOptions.Validate() method with nil receiver
func TestInputBorderOptions_Validate_NilReceiver(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on nil receiver, but did not panic")
		}
	}()

	var options *duit_props.InputBorderOptions
	options.Validate()
}

// Tests for different border types using table-driven tests
func TestInputBorder_DifferentTypes(t *testing.T) {
	testCases := []struct {
		name       string
		borderType string
	}{
		{
			name:       "Outline border",
			borderType: "outline",
		},
		{
			name:       "Underline border",
			borderType: "underline",
		},
		{
			name:       "None border",
			borderType: "none",
		},
		{
			name:       "Empty border type",
			borderType: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			border := duit_props.NewInputBorder(tc.borderType)

			if border.Type != tc.borderType {
				t.Fatalf("expected Type '%s', got '%s'", tc.borderType, border.Type)
			}

			err := border.Validate()
			if err != nil {
				t.Fatalf("expected no error for %s, got: %v", tc.name, err)
			}
		})
	}
}

// Tests for complex configurations
func TestInputBorder_ComplexConfiguration(t *testing.T) {
	color := duit_props.NewColorString("#ff0000")
	borderSide := duit_props.NewBorderSide().
		SetColor(color).
		SetWidth(2.0)

	border := duit_props.NewInputBorder("outline").
		SetGapPadding(6.0).
		SetBorderRadius(12.0).
		SetBorderSide(borderSide)

	// Validate the complete configuration
	err := border.Validate()
	if err != nil {
		t.Fatal("expected no error for complex configuration, got:", err)
	}

	// Verify all properties are set correctly
	if border.Type != "outline" {
		t.Fatalf("expected Type 'outline', got '%s'", border.Type)
	}

	if border.Options.GapPadding != 6.0 {
		t.Fatalf("expected GapPadding 6.0, got %f", border.Options.GapPadding)
	}

	if border.Options.BorderRadius != 12.0 {
		t.Fatalf("expected BorderRadius 12.0, got %f", border.Options.BorderRadius)
	}

	if border.Options.BorderSide == nil {
		t.Fatal("expected BorderSide to be set")
	}
}
