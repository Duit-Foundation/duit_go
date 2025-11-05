package duit_props_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for NewBoxConstraints constructor
func TestNewBoxConstraints(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()

	if constraints == nil {
		t.Fatal("expected non-nil BoxConstraints")
	}

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error for valid empty constraints, got:", err)
	}
}

// Tests for SetMinWidth method
func TestBoxConstraints_SetMinWidth(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMinWidth(100.0).SetMaxWidth(200.0)

	if constraints.MinWidth != 100.0 {
		t.Fatalf("expected MinWidth 100.0, got %f", constraints.MinWidth)
	}

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error for valid MinWidth, got:", err)
	}
}

func TestBoxConstraints_SetMinWidth_ZeroValue(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMinWidth(0.0)

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error for zero MinWidth, got:", err)
	}
}

func TestBoxConstraints_SetMinWidth_NegativeValue(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMinWidth(-10.0)

	err := constraints.Validate()
	if err == nil {
		t.Fatal("expected error for negative MinWidth")
	}
}

// Tests for SetMaxWidth method
func TestBoxConstraints_SetMaxWidth(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMaxWidth(200.0)

	if constraints.MaxWidth != 200.0 {
		t.Fatalf("expected MaxWidth 200.0, got %f", constraints.MaxWidth)
	}

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error for valid MaxWidth, got:", err)
	}
}

func TestBoxConstraints_SetMaxWidth_ZeroValue(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMaxWidth(0.0)

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error for zero MaxWidth, got:", err)
	}
}

func TestBoxConstraints_SetMaxWidth_NegativeValue(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMaxWidth(-10.0)

	err := constraints.Validate()
	if err == nil {
		t.Fatal("expected error for negative MaxWidth")
	}
}

// Tests for SetMinHeight method
func TestBoxConstraints_SetMinHeight(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMinHeight(150.0).SetMaxHeight(300.0)

	if constraints.MinHeight != 150.0 {
		t.Fatalf("expected MinHeight 150.0, got %f", constraints.MinHeight)
	}

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error for valid MinHeight, got:", err)
	}
}

func TestBoxConstraints_SetMinHeight_ZeroValue(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMinHeight(0.0)

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error for zero MinHeight, got:", err)
	}
}

func TestBoxConstraints_SetMinHeight_NegativeValue(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMinHeight(-10.0)

	err := constraints.Validate()
	if err == nil {
		t.Fatal("expected error for negative MinHeight")
	}
}

// Tests for SetMaxHeight method
func TestBoxConstraints_SetMaxHeight(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMaxHeight(300.0)

	if constraints.MaxHeight != 300.0 {
		t.Fatalf("expected MaxHeight 300.0, got %f", constraints.MaxHeight)
	}

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error for valid MaxHeight, got:", err)
	}
}

func TestBoxConstraints_SetMaxHeight_ZeroValue(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMaxHeight(0.0)

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error for zero MaxHeight, got:", err)
	}
}

func TestBoxConstraints_SetMaxHeight_NegativeValue(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMaxHeight(-10.0)

	err := constraints.Validate()
	if err == nil {
		t.Fatal("expected error for negative MaxHeight")
	}
}

// Tests for Validate method with constraint relationships
func TestBoxConstraints_Validate_MinWidthGreaterThanMaxWidth(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMinWidth(200.0).SetMaxWidth(100.0)

	err := constraints.Validate()
	if err == nil {
		t.Fatal("expected error when MinWidth > MaxWidth")
	}
}

func TestBoxConstraints_Validate_MinHeightGreaterThanMaxHeight(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.SetMinHeight(200.0).SetMaxHeight(100.0)

	err := constraints.Validate()
	if err == nil {
		t.Fatal("expected error when MinHeight > MaxHeight")
	}
}

func TestBoxConstraints_Validate_ValidConstraints(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.
		SetMinWidth(100.0).
		SetMaxWidth(200.0).
		SetMinHeight(150.0).
		SetMaxHeight(300.0)

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error for valid constraints, got:", err)
	}
}

func TestBoxConstraints_Validate_EqualMinMax(t *testing.T) {
	constraints := duit_props.NewBoxConstraints()
	constraints = constraints.
		SetMinWidth(100.0).
		SetMaxWidth(100.0).
		SetMinHeight(150.0).
		SetMaxHeight(150.0)

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error when Min equals Max, got:", err)
	}
}

// Tests for Validate method with nil receiver
func TestBoxConstraints_Validate_NilReceiver(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on nil receiver, but did not panic")
		}
	}()

	var constraints *duit_props.BoxConstraints

	constraints.Validate()
}

// Tests for method chaining
func TestBoxConstraints_MethodChaining(t *testing.T) {
	constraints := duit_props.NewBoxConstraints().
		SetMinWidth(50.0).
		SetMaxWidth(100.0).
		SetMinHeight(75.0).
		SetMaxHeight(150.0)

	if constraints.MinWidth != 50.0 {
		t.Fatalf("expected MinWidth 50.0, got %f", constraints.MinWidth)
	}
	if constraints.MaxWidth != 100.0 {
		t.Fatalf("expected MaxWidth 100.0, got %f", constraints.MaxWidth)
	}
	if constraints.MinHeight != 75.0 {
		t.Fatalf("expected MinHeight 75.0, got %f", constraints.MinHeight)
	}
	if constraints.MaxHeight != 150.0 {
		t.Fatalf("expected MaxHeight 150.0, got %f", constraints.MaxHeight)
	}

	err := constraints.Validate()
	if err != nil {
		t.Fatal("expected no error for chained method calls, got:", err)
	}
}
