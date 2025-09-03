package duit_props_test

import (
	"encoding/json"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for RadiusV2.Circular() method
func TestRadiusV2_Circular_ValidValue(t *testing.T) {
	radius := duit_props.RadiusV2{}.Circular(10.0)
	
	if radius == nil {
		t.Fatal("expected non-nil RadiusV2")
	}
	
	err := radius.Validate()
	if err != nil {
		t.Fatal("expected no error for valid circular radius, got:", err)
	}
}

func TestRadiusV2_Circular_ZeroValue(t *testing.T) {
	radius := duit_props.RadiusV2{}.Circular(0.0)
	
	err := radius.Validate()
	if err != nil {
		t.Fatal("expected no error for zero circular radius, got:", err)
	}
}

func TestRadiusV2_Circular_NegativeValue(t *testing.T) {
	radius := duit_props.RadiusV2{}.Circular(-5.0)
	
	err := radius.Validate()
	if err == nil {
		t.Fatal("expected error for negative circular radius")
	}
	
	expected := "radius must be greater than 0"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

// Tests for RadiusV2.Elliptical() method
func TestRadiusV2_Elliptical_ValidValues(t *testing.T) {
	radius := duit_props.RadiusV2{}.Elliptical([2]float32{10.0, 15.0})
	
	if radius == nil {
		t.Fatal("expected non-nil RadiusV2")
	}
	
	err := radius.Validate()
	if err != nil {
		t.Fatal("expected no error for valid elliptical radius, got:", err)
	}
}

func TestRadiusV2_Elliptical_ZeroValues(t *testing.T) {
	radius := duit_props.RadiusV2{}.Elliptical([2]float32{0.0, 0.0})
	
	err := radius.Validate()
	if err != nil {
		t.Fatal("expected no error for zero elliptical radius, got:", err)
	}
}

func TestRadiusV2_Elliptical_NegativeFirstValue(t *testing.T) {
	radius := duit_props.RadiusV2{}.Elliptical([2]float32{-5.0, 10.0})
	
	err := radius.Validate()
	if err == nil {
		t.Fatal("expected error for negative first elliptical radius value")
	}
	
	expected := "radius value at index 0 must be greater than 0"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestRadiusV2_Elliptical_NegativeSecondValue(t *testing.T) {
	radius := duit_props.RadiusV2{}.Elliptical([2]float32{10.0, -5.0})
	
	err := radius.Validate()
	if err == nil {
		t.Fatal("expected error for negative second elliptical radius value")
	}
	
	expected := "radius value at index 1 must be greater than 0"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestRadiusV2_Elliptical_BothNegativeValues(t *testing.T) {
	radius := duit_props.RadiusV2{}.Elliptical([2]float32{-3.0, -7.0})
	
	err := radius.Validate()
	if err == nil {
		t.Fatal("expected error for both negative elliptical radius values")
	}
	
	expected := "radius value at index 0 must be greater than 0"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

// Tests for RadiusV2.Validate() with nil receiver
func TestRadiusV2_Validate_NilReceiver(t *testing.T) {
	var radius *duit_props.RadiusV2
	
	err := radius.Validate()
	if err != nil {
		t.Fatal("expected no error for nil RadiusV2, got:", err)
	}
}

// Tests for RadiusV2.MarshalJSON() method
func TestRadiusV2_MarshalJSON_Circular(t *testing.T) {
	radius := duit_props.RadiusV2{}.Circular(10.0)
	
	data, err := radius.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling circular radius, got:", err)
	}
	
	expected := "10"
	if string(data) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(data))
	}
}

func TestRadiusV2_MarshalJSON_Elliptical(t *testing.T) {
	radius := duit_props.RadiusV2{}.Elliptical([2]float32{5.0, 8.0})
	
	data, err := radius.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling elliptical radius, got:", err)
	}
	
	expected := "[5,8]"
	if string(data) != expected {
		t.Fatalf("expected JSON '%s', got '%s'", expected, string(data))
	}
}

// Tests for BorderRadiusV2.All() method
func TestBorderRadiusV2_All_ValidRadius(t *testing.T) {
	radius := duit_props.RadiusV2{}.Circular(10.0)
	borderRadius := duit_props.BorderRadiusV2{}.All(radius)
	
	if borderRadius == nil {
		t.Fatal("expected non-nil BorderRadiusV2")
	}
	
	err := borderRadius.Validate()
	if err != nil {
		t.Fatal("expected no error for valid All() border radius, got:", err)
	}
}

func TestBorderRadiusV2_All_NilRadius(t *testing.T) {
	borderRadius := duit_props.BorderRadiusV2{}.All(nil)
	
	err := borderRadius.Validate()
	if err != nil {
		t.Fatal("expected no error for nil radius in All(), got:", err)
	}
}

func TestBorderRadiusV2_All_InvalidRadius(t *testing.T) {
	radius := duit_props.RadiusV2{}.Circular(-5.0)
	borderRadius := duit_props.BorderRadiusV2{}.All(radius)
	
	err := borderRadius.Validate()
	if err == nil {
		t.Fatal("expected error for invalid radius in All()")
	}
}

// Tests for BorderRadiusV2.Only() method
func TestBorderRadiusV2_Only_AllValidRadius(t *testing.T) {
	topLeft := duit_props.RadiusV2{}.Circular(5.0)
	topRight := duit_props.RadiusV2{}.Circular(10.0)
	bottomLeft := duit_props.RadiusV2{}.Elliptical([2]float32{3.0, 7.0})
	bottomRight := duit_props.RadiusV2{}.Elliptical([2]float32{8.0, 12.0})
	
	borderRadius := duit_props.BorderRadiusV2{}.Only(topLeft, topRight, bottomLeft, bottomRight)
	
	if borderRadius == nil {
		t.Fatal("expected non-nil BorderRadiusV2")
	}
	
	err := borderRadius.Validate()
	if err != nil {
		t.Fatal("expected no error for valid Only() border radius, got:", err)
	}
}

func TestBorderRadiusV2_Only_SomeNilRadius(t *testing.T) {
	topLeft := duit_props.RadiusV2{}.Circular(5.0)
	topRight := duit_props.RadiusV2{}.Circular(10.0)
	
	borderRadius := duit_props.BorderRadiusV2{}.Only(topLeft, topRight, nil, nil)
	
	err := borderRadius.Validate()
	if err != nil {
		t.Fatal("expected no error for some nil radius in Only(), got:", err)
	}
}

func TestBorderRadiusV2_Only_InvalidTopLeft(t *testing.T) {
	topLeft := duit_props.RadiusV2{}.Circular(-5.0)
	topRight := duit_props.RadiusV2{}.Circular(10.0)
	
	borderRadius := duit_props.BorderRadiusV2{}.Only(topLeft, topRight, nil, nil)
	
	err := borderRadius.Validate()
	if err == nil {
		t.Fatal("expected error for invalid topLeft radius in Only()")
	}
}

func TestBorderRadiusV2_Only_InvalidTopRight(t *testing.T) {
	topLeft := duit_props.RadiusV2{}.Circular(5.0)
	topRight := duit_props.RadiusV2{}.Elliptical([2]float32{-3.0, 7.0})
	
	borderRadius := duit_props.BorderRadiusV2{}.Only(topLeft, topRight, nil, nil)
	
	err := borderRadius.Validate()
	if err == nil {
		t.Fatal("expected error for invalid topRight radius in Only()")
	}
}

func TestBorderRadiusV2_Only_InvalidBottomLeft(t *testing.T) {
	bottomLeft := duit_props.RadiusV2{}.Elliptical([2]float32{5.0, -8.0})
	bottomRight := duit_props.RadiusV2{}.Circular(10.0)
	
	borderRadius := duit_props.BorderRadiusV2{}.Only(nil, nil, bottomLeft, bottomRight)
	
	err := borderRadius.Validate()
	if err == nil {
		t.Fatal("expected error for invalid bottomLeft radius in Only()")
	}
}

func TestBorderRadiusV2_Only_InvalidBottomRight(t *testing.T) {
	bottomLeft := duit_props.RadiusV2{}.Circular(5.0)
	bottomRight := duit_props.RadiusV2{}.Circular(-10.0)
	
	borderRadius := duit_props.BorderRadiusV2{}.Only(nil, nil, bottomLeft, bottomRight)
	
	err := borderRadius.Validate()
	if err == nil {
		t.Fatal("expected error for invalid bottomRight radius in Only()")
	}
}

// Tests for BorderRadiusV2.Vertical() method
func TestBorderRadiusV2_Vertical_ValidRadius(t *testing.T) {
	top := duit_props.RadiusV2{}.Circular(8.0)
	bottom := duit_props.RadiusV2{}.Elliptical([2]float32{5.0, 12.0})
	
	borderRadius := duit_props.BorderRadiusV2{}.Vertical(top, bottom)
	
	if borderRadius == nil {
		t.Fatal("expected non-nil BorderRadiusV2")
	}
	
	err := borderRadius.Validate()
	if err != nil {
		t.Fatal("expected no error for valid Vertical() border radius, got:", err)
	}
}

func TestBorderRadiusV2_Vertical_NilValues(t *testing.T) {
	borderRadius := duit_props.BorderRadiusV2{}.Vertical(nil, nil)
	
	err := borderRadius.Validate()
	if err != nil {
		t.Fatal("expected no error for nil values in Vertical(), got:", err)
	}
}

func TestBorderRadiusV2_Vertical_InvalidTop(t *testing.T) {
	top := duit_props.RadiusV2{}.Circular(-8.0)
	bottom := duit_props.RadiusV2{}.Circular(5.0)
	
	borderRadius := duit_props.BorderRadiusV2{}.Vertical(top, bottom)
	
	err := borderRadius.Validate()
	if err == nil {
		t.Fatal("expected error for invalid top radius in Vertical()")
	}
}

func TestBorderRadiusV2_Vertical_InvalidBottom(t *testing.T) {
	top := duit_props.RadiusV2{}.Circular(8.0)
	bottom := duit_props.RadiusV2{}.Elliptical([2]float32{5.0, -12.0})
	
	borderRadius := duit_props.BorderRadiusV2{}.Vertical(top, bottom)
	
	err := borderRadius.Validate()
	if err == nil {
		t.Fatal("expected error for invalid bottom radius in Vertical()")
	}
}

// Tests for BorderRadiusV2.Horizontal() method
func TestBorderRadiusV2_Horizontal_ValidRadius(t *testing.T) {
	left := duit_props.RadiusV2{}.Elliptical([2]float32{6.0, 9.0})
	right := duit_props.RadiusV2{}.Circular(12.0)
	
	borderRadius := duit_props.BorderRadiusV2{}.Horizontal(left, right)
	
	if borderRadius == nil {
		t.Fatal("expected non-nil BorderRadiusV2")
	}
	
	err := borderRadius.Validate()
	if err != nil {
		t.Fatal("expected no error for valid Horizontal() border radius, got:", err)
	}
}

func TestBorderRadiusV2_Horizontal_NilValues(t *testing.T) {
	borderRadius := duit_props.BorderRadiusV2{}.Horizontal(nil, nil)
	
	err := borderRadius.Validate()
	if err != nil {
		t.Fatal("expected no error for nil values in Horizontal(), got:", err)
	}
}

func TestBorderRadiusV2_Horizontal_InvalidLeft(t *testing.T) {
	left := duit_props.RadiusV2{}.Elliptical([2]float32{-6.0, 9.0})
	right := duit_props.RadiusV2{}.Circular(12.0)
	
	borderRadius := duit_props.BorderRadiusV2{}.Horizontal(left, right)
	
	err := borderRadius.Validate()
	if err == nil {
		t.Fatal("expected error for invalid left radius in Horizontal()")
	}
}

func TestBorderRadiusV2_Horizontal_InvalidRight(t *testing.T) {
	left := duit_props.RadiusV2{}.Circular(6.0)
	right := duit_props.RadiusV2{}.Circular(-12.0)
	
	borderRadius := duit_props.BorderRadiusV2{}.Horizontal(left, right)
	
	err := borderRadius.Validate()
	if err == nil {
		t.Fatal("expected error for invalid right radius in Horizontal()")
	}
}

// Tests for BorderRadiusV2.Validate() with nil receiver
func TestBorderRadiusV2_Validate_NilReceiver(t *testing.T) {
	var borderRadius *duit_props.BorderRadiusV2
	
	err := borderRadius.Validate()
	if err != nil {
		t.Fatal("expected no error for nil BorderRadiusV2, got:", err)
	}
}

// Tests for BorderRadiusV2.Validate() with different types using table-driven tests
func TestBorderRadiusV2_Validate_AllTypes(t *testing.T) {
	testCases := []struct {
		name         string
		borderRadius *duit_props.BorderRadiusV2
	}{
		{
			name:         "All type",
			borderRadius: duit_props.BorderRadiusV2{}.All(duit_props.RadiusV2{}.Circular(8.0)),
		},
		{
			name: "Only type",
			borderRadius: duit_props.BorderRadiusV2{}.Only(
				duit_props.RadiusV2{}.Circular(5.0),
				duit_props.RadiusV2{}.Elliptical([2]float32{3.0, 7.0}),
				duit_props.RadiusV2{}.Circular(10.0),
				duit_props.RadiusV2{}.Elliptical([2]float32{6.0, 4.0}),
			),
		},
		{
			name: "Vertical type",
			borderRadius: duit_props.BorderRadiusV2{}.Vertical(
				duit_props.RadiusV2{}.Circular(6.0),
				duit_props.RadiusV2{}.Elliptical([2]float32{4.0, 8.0}),
			),
		},
		{
			name: "Horizontal type",
			borderRadius: duit_props.BorderRadiusV2{}.Horizontal(
				duit_props.RadiusV2{}.Elliptical([2]float32{7.0, 3.0}),
				duit_props.RadiusV2{}.Circular(9.0),
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.borderRadius.Validate()
			if err != nil {
				t.Fatalf("expected no error for valid %s, got: %v", tc.name, err)
			}
		})
	}
}

// Tests for BorderRadiusV2.MarshalJSON() method
func TestBorderRadiusV2_MarshalJSON_All(t *testing.T) {
	radius := duit_props.RadiusV2{}.Circular(10.0)
	borderRadius := duit_props.BorderRadiusV2{}.All(radius)
	
	data, err := borderRadius.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling All() border radius, got:", err)
	}
	
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal("expected valid JSON from All() border radius, got:", err)
	}
	
	if result["radius"] == nil {
		t.Fatal("expected 'radius' field in All() JSON output")
	}
}

func TestBorderRadiusV2_MarshalJSON_Only(t *testing.T) {
	topLeft := duit_props.RadiusV2{}.Circular(5.0)
	topRight := duit_props.RadiusV2{}.Elliptical([2]float32{3.0, 7.0})
	
	borderRadius := duit_props.BorderRadiusV2{}.Only(topLeft, topRight, nil, nil)
	
	data, err := borderRadius.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling Only() border radius, got:", err)
	}
	
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal("expected valid JSON from Only() border radius, got:", err)
	}
	
	if result["topLeft"] == nil {
		t.Fatal("expected 'topLeft' field in Only() JSON output")
	}
	if result["topRight"] == nil {
		t.Fatal("expected 'topRight' field in Only() JSON output")
	}
}

func TestBorderRadiusV2_MarshalJSON_Vertical(t *testing.T) {
	top := duit_props.RadiusV2{}.Circular(8.0)
	bottom := duit_props.RadiusV2{}.Elliptical([2]float32{4.0, 6.0})
	
	borderRadius := duit_props.BorderRadiusV2{}.Vertical(top, bottom)
	
	data, err := borderRadius.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling Vertical() border radius, got:", err)
	}
	
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal("expected valid JSON from Vertical() border radius, got:", err)
	}
	
	if result["top"] == nil {
		t.Fatal("expected 'top' field in Vertical() JSON output")
	}
	if result["bottom"] == nil {
		t.Fatal("expected 'bottom' field in Vertical() JSON output")
	}
}

func TestBorderRadiusV2_MarshalJSON_Horizontal(t *testing.T) {
	left := duit_props.RadiusV2{}.Elliptical([2]float32{7.0, 3.0})
	right := duit_props.RadiusV2{}.Circular(9.0)
	
	borderRadius := duit_props.BorderRadiusV2{}.Horizontal(left, right)
	
	data, err := borderRadius.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error marshaling Horizontal() border radius, got:", err)
	}
	
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatal("expected valid JSON from Horizontal() border radius, got:", err)
	}
	
	if result["left"] == nil {
		t.Fatal("expected 'left' field in Horizontal() JSON output")
	}
	if result["right"] == nil {
		t.Fatal("expected 'right' field in Horizontal() JSON output")
	}
}
