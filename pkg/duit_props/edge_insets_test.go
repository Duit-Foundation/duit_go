package duit_props_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for EdgeInsetsV2.All() method
func TestEdgeInsetsV2_All(t *testing.T) {
	edgeInsets := duit_props.EdgeInsetsV2{}.All(10.0)

	if edgeInsets == nil {
		t.Fatal("expected non-nil EdgeInsetsV2")
	}

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for valid All() value, got:", err)
	}
}

func TestEdgeInsetsV2_All_ZeroValue(t *testing.T) {
	edgeInsets := duit_props.EdgeInsetsV2{}.All(0.0)

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for zero value, got:", err)
	}
}

func TestEdgeInsetsV2_All_NegativeValue(t *testing.T) {
	edgeInsets := duit_props.EdgeInsetsV2{}.All(-5.0)

	err := edgeInsets.Validate()
	if err == nil {
		t.Fatal("expected error for negative All() value")
	}
}

// Tests for EdgeInsetsV2.LTRB() method
func TestEdgeInsetsV2_LTRB(t *testing.T) {
	edgeInsets := duit_props.EdgeInsetsV2{}.LTRB(5.0, 10.0, 15.0, 20.0)

	if edgeInsets == nil {
		t.Fatal("expected non-nil EdgeInsetsV2")
	}

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for valid LTRB() values, got:", err)
	}
}

func TestEdgeInsetsV2_LTRB_ZeroValues(t *testing.T) {
	edgeInsets := duit_props.EdgeInsetsV2{}.LTRB(0.0, 0.0, 0.0, 0.0)

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for zero LTRB() values, got:", err)
	}
}

func TestEdgeInsetsV2_LTRB_NegativeValues(t *testing.T) {
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
			edgeInsets := duit_props.EdgeInsetsV2{}.LTRB(tc.left, tc.top, tc.right, tc.bottom)

			err := edgeInsets.Validate()
			if err == nil {
				t.Fatalf("expected error for %s", tc.name)
			}
		})
	}
}

// Tests for EdgeInsetsV2.Symmentric() method
func TestEdgeInsetsV2_Symmentric(t *testing.T) {
	edgeInsets := duit_props.EdgeInsetsV2{}.Symmentric(10.0, 5.0)

	if edgeInsets == nil {
		t.Fatal("expected non-nil EdgeInsetsV2")
	}

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for valid Symmentric() values, got:", err)
	}
}

func TestEdgeInsetsV2_Symmentric_ZeroValues(t *testing.T) {
	edgeInsets := duit_props.EdgeInsetsV2{}.Symmentric(0.0, 0.0)

	err := edgeInsets.Validate()
	if err != nil {
		t.Fatal("expected no error for zero Symmentric() values, got:", err)
	}
}

func TestEdgeInsetsV2_Symmentric_NegativeValues(t *testing.T) {
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
			edgeInsets := duit_props.EdgeInsetsV2{}.Symmentric(tc.vertical, tc.horizontal)

			err := edgeInsets.Validate()
			if err == nil {
				t.Fatalf("expected error for %s", tc.name)
			}
		})
	}
}

// Tests for EdgeInsetsV2.Validate() method with uninitialized data
func TestEdgeInsetsV2_Validate_UninitializedData(t *testing.T) {
	edgeInsets := &duit_props.EdgeInsetsV2{}

	err := edgeInsets.Validate()
	if err == nil {
		t.Fatal("expected error for uninitialized EdgeInsetsV2")
	}
}

// Tests for EdgeInsetsV2.Validate() method with different valid types
func TestEdgeInsetsV2_Validate_AllTypes(t *testing.T) {
	testCases := []struct {
		name       string
		edgeInsets *duit_props.EdgeInsetsV2
	}{
		{
			name:       "All type",
			edgeInsets: duit_props.EdgeInsetsV2{}.All(8.0),
		},
		{
			name:       "LTRB type",
			edgeInsets: duit_props.EdgeInsetsV2{}.LTRB(1.0, 2.0, 3.0, 4.0),
		},
		{
			name:       "Symmentric type",
			edgeInsets: duit_props.EdgeInsetsV2{}.Symmentric(6.0, 4.0),
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

// Tests for EdgeInsetsV2.MarshalJSON() method
func TestEdgeInsetsV2_MarshalJSON_All(t *testing.T) {
	edgeInsets := duit_props.EdgeInsetsV2{}.All(10.0)

	jsonData, err := edgeInsets.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	expected := "10"
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}

func TestEdgeInsetsV2_MarshalJSON_LTRB(t *testing.T) {
	edgeInsets := duit_props.EdgeInsetsV2{}.LTRB(1.0, 2.0, 3.0, 4.0)

	jsonData, err := edgeInsets.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	expected := "[1,2,3,4]"
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}

func TestEdgeInsetsV2_MarshalJSON_Symmentric(t *testing.T) {
	edgeInsets := duit_props.EdgeInsetsV2{}.Symmentric(5.0, 3.0)

	jsonData, err := edgeInsets.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	expected := "[5,3]"
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}

func TestEdgeInsetsV2_MarshalJSON_Uninitialized(t *testing.T) {
	edgeInsets := &duit_props.EdgeInsetsV2{}

	jsonData, err := edgeInsets.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() on uninitialized data, got:", err)
	}

	expected := "null"
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}
