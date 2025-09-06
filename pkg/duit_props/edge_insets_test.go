package duit_props_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for NewEdgeInsetsAll constructor
func TestNewEdgeInsetsAll(t *testing.T) {
	edgeInsets := duit_props.NewEdgeInsetsAll(10.0)

	if edgeInsets == nil {
		t.Fatal("expected non-nil EdgeInsets")
	}

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for valid All() value, got:", err)
	}
}

func TestNewEdgeInsetsAll_ZeroValue(t *testing.T) {
	edgeInsets := duit_props.NewEdgeInsetsAll(0.0)

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for zero value, got:", err)
	}
}

func TestNewEdgeInsetsAll_NegativeValue(t *testing.T) {
	edgeInsets := duit_props.NewEdgeInsetsAll(-5.0)

	err := edgeInsets.Validate()
	if err == nil {
		t.Fatal("expected error for negative All() value")
	}
}

// Tests for NewEdgeInsetsLTRB constructor
func TestNewEdgeInsetsLTRB(t *testing.T) {
	edgeInsets := duit_props.NewEdgeInsetsLTRB(5.0, 10.0, 15.0, 20.0)

	if edgeInsets == nil {
		t.Fatal("expected non-nil EdgeInsets")
	}

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for valid LTRB() values, got:", err)
	}
}

func TestNewEdgeInsetsLTRB_ZeroValues(t *testing.T) {
	edgeInsets := duit_props.NewEdgeInsetsLTRB(0.0, 0.0, 0.0, 0.0)

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for zero LTRB() values, got:", err)
	}
}

func TestNewEdgeInsetsLTRB_NegativeValues(t *testing.T) {
	testCases := []struct {
		name   string
		left   float32
		top    float32
		right  float32
		bottom float32
	}{
		{"negative left", -1.0, 5.0, 5.0, 5.0},
		{"negative top", 5.0, -1.0, 5.0, 5.0},
		{"negative right", 5.0, 5.0, -1.0, 5.0},
		{"negative bottom", 5.0, 5.0, 5.0, -1.0},
		{"all negative", -1.0, -2.0, -3.0, -4.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			edgeInsets := duit_props.NewEdgeInsetsLTRB(tc.left, tc.top, tc.right, tc.bottom)

			err := edgeInsets.Validate()
			if err == nil {
				t.Fatalf("expected error for %s", tc.name)
			}
		})
	}
}

// Tests for NewEdgeInsetsSymmetric constructor
func TestNewEdgeInsetsSymmetric(t *testing.T) {
	edgeInsets := duit_props.NewEdgeInsetsSymmetric(10.0, 5.0)

	if edgeInsets == nil {
		t.Fatal("expected non-nil EdgeInsets")
	}

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for valid Symmentric() values, got:", err)
	}
}

func TestNewEdgeInsetsSymmetric_ZeroValues(t *testing.T) {
	edgeInsets := duit_props.NewEdgeInsetsSymmetric(0.0, 0.0)

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for zero Symmentric() values, got:", err)
	}
}

func TestNewEdgeInsetsSymmetric_NegativeValues(t *testing.T) {
	testCases := []struct {
		name       string
		vertical   float32
		horizontal float32
	}{
		{"negative vertical", -1.0, 5.0},
		{"negative horizontal", 5.0, -1.0},
		{"both negative", -1.0, -2.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			edgeInsets := duit_props.NewEdgeInsetsSymmetric(tc.vertical, tc.horizontal)

			err := edgeInsets.Validate()
			if err == nil {
				t.Fatalf("expected error for %s", tc.name)
			}
		})
	}
}

// Tests for EdgeInsets.Validate() method with nil pointer
func TestEdgeInsets_Validate_NilPointer(t *testing.T) {
	var edgeInsets *duit_props.EdgeInsets = nil

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for nil EdgeInsets, got:", err)
	}
}

// Tests for EdgeInsets.Validate() method with uninitialized data
func TestEdgeInsets_Validate_UninitializedData(t *testing.T) {
	edgeInsets := &duit_props.EdgeInsets{}

	err := edgeInsets.Validate()
	if err == nil {
		t.Fatal("expected error for uninitialized EdgeInsets")
	}
}

// Tests for EdgeInsets.Validate() method with different valid types
func TestEdgeInsets_Validate_AllTypes(t *testing.T) {
	testCases := []struct {
		name       string
		edgeInsets *duit_props.EdgeInsets
	}{
		{
			name:       "All type",
			edgeInsets: duit_props.NewEdgeInsetsAll(8.0),
		},
		{
			name:       "LTRB type",
			edgeInsets: duit_props.NewEdgeInsetsLTRB(1.0, 2.0, 3.0, 4.0),
		},
		{
			name:       "Symmetric type",
			edgeInsets: duit_props.NewEdgeInsetsSymmetric(6.0, 4.0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.edgeInsets.Validate()
			if err != nil {
				t.Fatalf("expected no error for valid %s, got: %v", tc.name, err)
			}
		})
	}
}

// Tests for EdgeInsets.MarshalJSON() method
func TestNewEdgeInsetsAll_MarshalJSON(t *testing.T) {
	edgeInsets := duit_props.NewEdgeInsetsAll(10.0)

	jsonData, err := edgeInsets.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	expected := "10"
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}

func TestNewEdgeInsetsLTRB_MarshalJSON(t *testing.T) {
	edgeInsets := duit_props.NewEdgeInsetsLTRB(1.0, 2.0, 3.0, 4.0)

	jsonData, err := edgeInsets.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	expected := "[1,2,3,4]"
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}

func TestNewEdgeInsetsSymmetric_MarshalJSON(t *testing.T) {
	edgeInsets := duit_props.NewEdgeInsetsSymmetric(5.0, 3.0)

	jsonData, err := edgeInsets.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	expected := "[5,3]"
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}

func TestEdgeInsets_MarshalJSON_Uninitialized(t *testing.T) {
	edgeInsets := &duit_props.EdgeInsets{}

	jsonData, err := edgeInsets.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() on uninitialized data, got:", err)
	}

	expected := "null"
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}
