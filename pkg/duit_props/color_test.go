package duit_props_test

import (
	"encoding/json"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for NewColorString constructor
func TestNewColorString_ValidColor(t *testing.T) {
	color := duit_props.NewColorString("#ff0000")

	if color == nil {
		t.Fatal("expected non-nil Color")
	}

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for valid color string, got:", err)
	}
}

func TestNewColorString_ValidColorWithUppercase(t *testing.T) {
	color := duit_props.NewColorString("#FF0000")

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for valid uppercase color string, got:", err)
	}
}

func TestNewColorString_ValidColorMixed(t *testing.T) {
	color := duit_props.NewColorString("#aB12Cd")

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for valid mixed case color string, got:", err)
	}
}

func TestNewColorString_InvalidColorTooShort(t *testing.T) {
	color := duit_props.NewColorString("#ff00")

	err := color.Validate()
	if err == nil {
		t.Fatal("expected error for color string too short")
	}

	expected := "invalid color string"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestNewColorString_InvalidColorTooLong(t *testing.T) {
	color := duit_props.NewColorString("#ff0000aa")

	err := color.Validate()
	if err == nil {
		t.Fatal("expected error for color string too long")
	}

	expected := "invalid color string"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestNewColorString_InvalidColorNoHash(t *testing.T) {
	color := duit_props.NewColorString("ff00000")

	err := color.Validate()
	if err == nil {
		t.Fatal("expected error for color string without #")
	}

	expected := "invalid color string. Must start with #"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestNewColorString_InvalidColorEmptyString(t *testing.T) {
	color := duit_props.NewColorString("")

	err := color.Validate()
	if err == nil {
		t.Fatal("expected error for empty color string")
	}

	expected := "invalid color string"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

// Tests for NewColorRGBO constructor
func TestNewColorRGBO_ValidValues(t *testing.T) {
	color := duit_props.NewColorRGBO([3]uint8{255, 128, 0}, 1.0)

	if color == nil {
		t.Fatal("expected non-nil Color")
	}

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for valid RGBO values, got:", err)
	}
}

func TestNewColorRGBO_ZeroValues(t *testing.T) {
	color := duit_props.NewColorRGBO([3]uint8{0, 0, 0}, 0.0)

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for zero RGBO values, got:", err)
	}
}

func TestNewColorRGBO_MaxValues(t *testing.T) {
	color := duit_props.NewColorRGBO([3]uint8{255, 255, 255}, 1.0)

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for max RGBO values, got:", err)
	}
}

func TestNewColorRGBO_PartialOpacity(t *testing.T) {
	color := duit_props.NewColorRGBO([3]uint8{100, 150, 200}, 0.5)

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for partial opacity RGBO values, got:", err)
	}
}

// Tests for NewColorRGBA constructor
func TestNewColorRGBA_ValidValues(t *testing.T) {
	color := duit_props.NewColorRGBA([4]uint8{255, 128, 0, 255})

	if color == nil {
		t.Fatal("expected non-nil Color")
	}

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for valid RGBA values, got:", err)
	}
}

func TestNewColorRGBA_ZeroValues(t *testing.T) {
	color := duit_props.NewColorRGBA([4]uint8{0, 0, 0, 0})

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for zero RGBA values, got:", err)
	}
}

func TestNewColorRGBA_MaxValues(t *testing.T) {
	color := duit_props.NewColorRGBA([4]uint8{255, 255, 255, 255})

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for max RGBA values, got:", err)
	}
}

func TestNewColorRGBA_PartialAlpha(t *testing.T) {
	color := duit_props.NewColorRGBA([4]uint8{100, 150, 200, 128})

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for partial alpha RGBA values, got:", err)
	}
}

// Tests for Color.Validate() with nil receiver
func TestColor_Validate_NilReceiver(t *testing.T) {
	var color *duit_props.Color

	err := color.Validate()
	if err != nil {
		t.Fatal("expected no error for nil Color, got:", err)
	}
}

// Tests for Color.Validate() with different types using table-driven tests
func TestColor_Validate_AllTypes(t *testing.T) {
	testCases := []struct {
		name  string
		color *duit_props.Color
	}{
		{
			name:  "String type",
			color: duit_props.NewColorString("#ff0000"),
		},
		{
			name:  "RGBO type",
			color: duit_props.NewColorRGBO([3]uint8{255, 128, 0}, 0.8),
		},
		{
			name:  "RGBA type",
			color: duit_props.NewColorRGBA([4]uint8{100, 150, 200, 255}),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.color.Validate()
			if err != nil {
				t.Fatalf("expected no error for valid %s, got: %v", tc.name, err)
			}
		})
	}
}

// Tests for Color.MarshalJSON() method
func TestColor_MarshalJSON_String(t *testing.T) {
	color := duit_props.NewColorString("#ff0000")

	data, err := color.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling string color, got:", err)
	}

	expected := "\"#ff0000\""
	if string(data) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(data))
	}
}

func TestColor_MarshalJSON_RGBO(t *testing.T) {
	color := duit_props.NewColorRGBO([3]uint8{255, 128, 0}, 1.0)

	data, err := color.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling RGBO color, got:", err)
	}

	expected := "[255,128,0,255]"
	if string(data) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(data))
	}
}

func TestColor_MarshalJSON_RGBO_PartialOpacity(t *testing.T) {
	color := duit_props.NewColorRGBO([3]uint8{100, 150, 200}, 0.5)

	data, err := color.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling RGBO color with partial opacity, got:", err)
	}

	expected := "[100,150,200,127]" // 0.5 * 255 = 127.5 -> 127 as uint8
	if string(data) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(data))
	}
}

func TestColor_MarshalJSON_RGBA(t *testing.T) {
	color := duit_props.NewColorRGBA([4]uint8{255, 128, 0, 200})

	data, err := color.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling RGBA color, got:", err)
	}

	expected := "[255,128,0,200]"
	if string(data) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(data))
	}
}

func TestColor_MarshalJSON_AllConstructors(t *testing.T) {
	testCases := []struct {
		name     string
		color    *duit_props.Color
		expected string
	}{
		{
			name:     "String constructor",
			color:    duit_props.NewColorString("#aabbcc"),
			expected: "\"#aabbcc\"",
		},
		{
			name:     "RGBO constructor",
			color:    duit_props.NewColorRGBO([3]uint8{170, 187, 204}, 0.6),
			expected: "[170,187,204,153]", // 0.6 * 255 = 153
		},
		{
			name:     "RGBA constructor", 
			color:    duit_props.NewColorRGBA([4]uint8{170, 187, 204, 153}),
			expected: "[170,187,204,153]",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := tc.color.MarshalJSON()
			if err != nil {
				t.Fatal("expected no error marshaling color, got:", err)
			}

			if string(data) != tc.expected {
				t.Fatalf("expected JSON '%s', got '%s'", tc.expected, string(data))
			}

			// Verify that the JSON is valid
			var result interface{}
			err = json.Unmarshal(data, &result)
			if err != nil {
				t.Fatalf("expected valid JSON, got unmarshal error: %v", err)
			}
		})
	}
}
