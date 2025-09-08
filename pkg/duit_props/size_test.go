package duit_props_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for NewSize constructor
func TestNewSize(t *testing.T) {
	size := duit_props.NewSize(100.0, 200.0)

	if size == nil {
		t.Fatal("expected non-nil Size")
	}

	err := size.Validate()
	if err != nil {
		t.Fatal("expected no error for valid Size, got:", err)
	}
}

func TestNewSize_ZeroValues(t *testing.T) {
	size := duit_props.NewSize(0.0, 0.0)

	err := size.Validate()
	if err == nil {
		t.Fatal("expected error for zero Size values")
	}
}

func TestNewSize_NegativeValues(t *testing.T) {
	testCases := []struct {
		name   string
		width  float32
		height float32
	}{
		{"negative width", -10.0, 20.0},
		{"negative height", 10.0, -20.0},
		{"both negative", -10.0, -20.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			size := duit_props.NewSize(tc.width, tc.height)

			err := size.Validate()
			if err == nil {
				t.Fatalf("expected error for %s", tc.name)
			}
		})
	}
}

// Tests for NewSizeSquare constructor
func TestNewSizeSquare(t *testing.T) {
	size := duit_props.NewSizeSquare(50)

	if size == nil {
		t.Fatal("expected non-nil Size")
	}

	err := size.Validate()
	if err != nil {
		t.Fatal("expected no error for valid square Size, got:", err)
	}
}

func TestNewSizeSquare_ZeroValue(t *testing.T) {
	size := duit_props.NewSizeSquare(0)

	err := size.Validate()
	if err == nil {
		t.Fatal("expected error for zero square Size value")
	}
}

// Tests for NewSizeFrom constructor
func TestNewSizeFrom(t *testing.T) {
	size := duit_props.NewSizeFrom(150.0, duit_props.AxisHorizontal)

	if size == nil {
		t.Fatal("expected non-nil Size")
	}

	err := size.Validate()
	if err != nil {
		t.Fatal("expected no error for valid SizeFrom, got:", err)
	}
}

func TestNewSizeFrom_ZeroValue(t *testing.T) {
	size := duit_props.NewSizeFrom(0.0, duit_props.AxisVertical)

	err := size.Validate()
	if err == nil {
		t.Fatal("expected error for zero SizeFrom value")
	}
}

func TestNewSizeFrom_NegativeValue(t *testing.T) {
	size := duit_props.NewSizeFrom(-25.0, duit_props.AxisHorizontal)

	err := size.Validate()
	if err == nil {
		t.Fatal("expected error for negative SizeFrom value")
	}
}

// Tests for Size.Validate() method with nil pointer
func TestSize_Validate_NilPointer(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on nil receiver, but did not panic")
		}
	}()

	var size *duit_props.Size = nil

	size.Validate()
}

// Tests for Size.Validate() method with uninitialized data
func TestSize_Validate_UninitializedData(t *testing.T) {
	size := &duit_props.Size{}

	err := size.Validate()
	if err != nil {
		t.Fatal("expected no error for uninitialized Size, got:", err)
	}
}

// Tests for Size.Validate() method with different valid types
func TestSize_Validate_AllTypes(t *testing.T) {
	testCases := []struct {
		name string
		size *duit_props.Size
	}{
		{
			name: "WH type",
			size: duit_props.NewSize(100.0, 200.0),
		},
		{
			name: "Square type",
			size: duit_props.NewSizeSquare(75),
		},
		{
			name: "From type",
			size: duit_props.NewSizeFrom(120.0, duit_props.AxisVertical),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.size.Validate()
			if err != nil {
				t.Fatalf("expected no error for valid %s, got: %v", tc.name, err)
			}
		})
	}
}

// Tests for Size.MarshalJSON() method
func TestNewSize_MarshalJSON(t *testing.T) {
	size := duit_props.NewSize(100.0, 200.0)

	jsonData, err := size.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	expected := `{"width":100,"height":200}`
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}

func TestNewSizeSquare_MarshalJSON(t *testing.T) {
	size := duit_props.NewSizeSquare(50)

	jsonData, err := size.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	expected := "50"
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}

func TestNewSizeFrom_MarshalJSON(t *testing.T) {
	size := duit_props.NewSizeFrom(150.0, duit_props.AxisHorizontal)

	jsonData, err := size.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	expected := `{"value":150,"mainAxis":"horizontal"}`
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}

func TestSize_MarshalJSON_Uninitialized(t *testing.T) {
	size := &duit_props.Size{}

	jsonData, err := size.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON() on uninitialized data, got:", err)
	}

	expected := "null"
	if string(jsonData) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(jsonData))
	}
}
