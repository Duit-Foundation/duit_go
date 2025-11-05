package duit_utils

import (
	"testing"
)

// Tests for RangeValidator
func TestRangeValidator_IntValidRange(t *testing.T) {
	err := RangeValidator(5, 0, 10)
	if err != nil {
		t.Fatal("expected no error for valid int range, got:", err)
	}
}

func TestRangeValidator_IntBoundaryValues(t *testing.T) {
	// Test minimum boundary
	err := RangeValidator(0, 0, 10)
	if err != nil {
		t.Fatal("expected no error for minimum boundary value, got:", err)
	}

	// Test maximum boundary
	err = RangeValidator(10, 0, 10)
	if err != nil {
		t.Fatal("expected no error for maximum boundary value, got:", err)
	}
}

func TestRangeValidator_IntBelowRange(t *testing.T) {
	err := RangeValidator(-1, 0, 10)
	if err == nil {
		t.Fatal("expected error for value below range")
	}

	expected := "value must be between the allowed range: min=0, max=10"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestRangeValidator_IntAboveRange(t *testing.T) {
	err := RangeValidator(11, 0, 10)
	if err == nil {
		t.Fatal("expected error for value above range")
	}

	expected := "value must be between the allowed range: min=0, max=10"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestRangeValidator_FloatValidRange(t *testing.T) {
	err := RangeValidator(2.5, 0.0, 5.0)
	if err != nil {
		t.Fatal("expected no error for valid float range, got:", err)
	}
}

func TestRangeValidator_FloatBelowRange(t *testing.T) {
	err := RangeValidator(-0.1, 0.0, 1.0)
	if err == nil {
		t.Fatal("expected error for float value below range")
	}

	expected := "value must be between the allowed range: min=0, max=1"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestRangeValidator_FloatAboveRange(t *testing.T) {
	err := RangeValidator(1.1, 0.0, 1.0)
	if err == nil {
		t.Fatal("expected error for float value above range")
	}

	expected := "value must be between the allowed range: min=0, max=1"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestRangeValidator_UintValidRange(t *testing.T) {
	var value uint = 5
	var min uint = 0
	var max uint = 10

	err := RangeValidator(value, min, max)
	if err != nil {
		t.Fatal("expected no error for valid uint range, got:", err)
	}
}

func TestRangeValidator_Int64ValidRange(t *testing.T) {
	var value int64 = 100
	var min int64 = 0
	var max int64 = 1000

	err := RangeValidator(value, min, max)
	if err != nil {
		t.Fatal("expected no error for valid int64 range, got:", err)
	}
}

func TestRangeValidator_NegativeRange(t *testing.T) {
	err := RangeValidator(-5, -10, -1)
	if err != nil {
		t.Fatal("expected no error for valid negative range, got:", err)
	}
}

func TestRangeValidator_NegativeRangeOutside(t *testing.T) {
	err := RangeValidator(0, -10, -1)
	if err == nil {
		t.Fatal("expected error for value outside negative range")
	}

	expected := "value must be between the allowed range: min=-10, max=-1"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}
