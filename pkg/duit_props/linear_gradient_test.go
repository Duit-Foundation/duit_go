package duit_props_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for NewLinearGradient constructor
func TestNewLinearGradient(t *testing.T) {
	gradient := duit_props.NewLinearGradient()

	if gradient == nil {
		t.Fatal("expected non-nil LinearGradient")
	}

	if gradient.Colors != nil {
		t.Fatal("expected Colors to be nil for empty gradient")
	}

	if gradient.RotationAngle != 0 {
		t.Fatalf("expected RotationAngle to be 0, got %f", gradient.RotationAngle)
	}

	if gradient.Stops != nil {
		t.Fatal("expected Stops to be nil for empty gradient")
	}
}

// Tests for SetColors method
func TestLinearGradient_SetColors(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	colors := []*duit_props.Color{
		duit_props.NewColorString("#ff0000"),
		duit_props.NewColorString("#0000ff"),
	}
	gradient = gradient.SetColors(colors)

	if len(gradient.Colors) != 2 {
		t.Fatalf("expected 2 colors, got %d", len(gradient.Colors))
	}

	if gradient.Colors[0] != colors[0] {
		t.Fatal("expected first color to be set correctly")
	}

	if gradient.Colors[1] != colors[1] {
		t.Fatal("expected second color to be set correctly")
	}
}

func TestLinearGradient_SetColors_Empty(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	gradient = gradient.SetColors([]*duit_props.Color{})

	if len(gradient.Colors) != 0 {
		t.Fatalf("expected 0 colors, got %d", len(gradient.Colors))
	}
}

func TestLinearGradient_SetColors_Nil(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	gradient = gradient.SetColors(nil)

	if gradient.Colors != nil {
		t.Fatal("expected Colors to be nil")
	}
}

// Tests for AddColor method
func TestLinearGradient_AddColor(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	color := duit_props.NewColorString("#ff0000")
	gradient = gradient.AddColor(color)

	if len(gradient.Colors) != 1 {
		t.Fatalf("expected 1 color, got %d", len(gradient.Colors))
	}

	if gradient.Colors[0] != color {
		t.Fatal("expected color to be added correctly")
	}
}

func TestLinearGradient_AddColor_Multiple(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	color1 := duit_props.NewColorString("#ff0000")
	color2 := duit_props.NewColorString("#0000ff")
	gradient = gradient.AddColor(color1).AddColor(color2)

	if len(gradient.Colors) != 2 {
		t.Fatalf("expected 2 colors, got %d", len(gradient.Colors))
	}

	if gradient.Colors[0] != color1 {
		t.Fatal("expected first color to be added correctly")
	}

	if gradient.Colors[1] != color2 {
		t.Fatal("expected second color to be added correctly")
	}
}

func TestLinearGradient_AddColor_Nil(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	gradient = gradient.AddColor(nil)

	if len(gradient.Colors) != 1 {
		t.Fatalf("expected 1 color, got %d", len(gradient.Colors))
	}

	if gradient.Colors[0] != nil {
		t.Fatal("expected nil color to be added")
	}
}

// Tests for SetRotationAngle method
func TestLinearGradient_SetRotationAngle(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	angle := float32(1.57) // π/2 radians (90 degrees)
	gradient = gradient.SetRotationAngle(angle)

	if gradient.RotationAngle != angle {
		t.Fatalf("expected RotationAngle %f, got %f", angle, gradient.RotationAngle)
	}
}

func TestLinearGradient_SetRotationAngle_Zero(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	gradient = gradient.SetRotationAngle(0.0)

	if gradient.RotationAngle != 0.0 {
		t.Fatalf("expected RotationAngle 0.0, got %f", gradient.RotationAngle)
	}
}

func TestLinearGradient_SetRotationAngle_Negative(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	angle := float32(-1.57)
	gradient = gradient.SetRotationAngle(angle)

	if gradient.RotationAngle != angle {
		t.Fatalf("expected RotationAngle %f, got %f", angle, gradient.RotationAngle)
	}
}

// Tests for SetStops method
func TestLinearGradient_SetStops(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	stops := []float32{0.0, 0.5, 1.0}
	gradient = gradient.SetStops(stops)

	if len(gradient.Stops) != 3 {
		t.Fatalf("expected 3 stops, got %d", len(gradient.Stops))
	}

	for i, expectedStop := range stops {
		if gradient.Stops[i] != expectedStop {
			t.Fatalf("expected stop[%d] to be %f, got %f", i, expectedStop, gradient.Stops[i])
		}
	}
}

func TestLinearGradient_SetStops_Empty(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	gradient = gradient.SetStops([]float32{})

	if len(gradient.Stops) != 0 {
		t.Fatalf("expected 0 stops, got %d", len(gradient.Stops))
	}
}

func TestLinearGradient_SetStops_Nil(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	gradient = gradient.SetStops(nil)

	if gradient.Stops != nil {
		t.Fatal("expected Stops to be nil")
	}
}

// Tests for AddStop method
func TestLinearGradient_AddStop(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	stop := float32(0.5)
	gradient = gradient.AddStop(stop)

	if len(gradient.Stops) != 1 {
		t.Fatalf("expected 1 stop, got %d", len(gradient.Stops))
	}

	if gradient.Stops[0] != stop {
		t.Fatalf("expected stop to be %f, got %f", stop, gradient.Stops[0])
	}
}

func TestLinearGradient_AddStop_Multiple(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	stop1 := float32(0.0)
	stop2 := float32(1.0)
	gradient = gradient.AddStop(stop1).AddStop(stop2)

	if len(gradient.Stops) != 2 {
		t.Fatalf("expected 2 stops, got %d", len(gradient.Stops))
	}

	if gradient.Stops[0] != stop1 {
		t.Fatalf("expected first stop to be %f, got %f", stop1, gradient.Stops[0])
	}

	if gradient.Stops[1] != stop2 {
		t.Fatalf("expected second stop to be %f, got %f", stop2, gradient.Stops[1])
	}
}

func TestLinearGradient_AddStop_OutOfRange(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	// Testing with values outside typical 0.0-1.0 range
	gradient = gradient.AddStop(-0.5).AddStop(1.5)

	if len(gradient.Stops) != 2 {
		t.Fatalf("expected 2 stops, got %d", len(gradient.Stops))
	}

	if gradient.Stops[0] != -0.5 {
		t.Fatalf("expected first stop to be -0.5, got %f", gradient.Stops[0])
	}

	if gradient.Stops[1] != 1.5 {
		t.Fatalf("expected second stop to be 1.5, got %f", gradient.Stops[1])
	}
}

// Tests for SetBegin method
func TestLinearGradient_SetBegin(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	begin := duit_props.AlignmentTopLeft
	gradient = gradient.SetBegin(begin)

	if gradient.Begin != begin {
		t.Fatalf("expected Begin to be %s, got %s", begin, gradient.Begin)
	}
}

func TestLinearGradient_SetBegin_Center(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	begin := duit_props.AlignmentCenter
	gradient = gradient.SetBegin(begin)

	if gradient.Begin != begin {
		t.Fatalf("expected Begin to be %s, got %s", begin, gradient.Begin)
	}
}

// Tests for SetEnd method
func TestLinearGradient_SetEnd(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	end := duit_props.AlignmentBottomRight
	gradient = gradient.SetEnd(end)

	if gradient.End != end {
		t.Fatalf("expected End to be %s, got %s", end, gradient.End)
	}
}

func TestLinearGradient_SetEnd_Center(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	end := duit_props.AlignmentCenter
	gradient = gradient.SetEnd(end)

	if gradient.End != end {
		t.Fatalf("expected End to be %s, got %s", end, gradient.End)
	}
}

// Tests for method chaining
func TestLinearGradient_MethodChaining(t *testing.T) {
	color1 := duit_props.NewColorString("#ff0000")
	color2 := duit_props.NewColorString("#0000ff")
	
	gradient := duit_props.NewLinearGradient().
		AddColor(color1).
		AddColor(color2).
		SetRotationAngle(1.57).
		AddStop(0.0).
		AddStop(1.0).
		SetBegin(duit_props.AlignmentTopLeft).
		SetEnd(duit_props.AlignmentBottomRight)

	if len(gradient.Colors) != 2 {
		t.Fatalf("expected 2 colors, got %d", len(gradient.Colors))
	}

	if gradient.Colors[0] != color1 {
		t.Fatal("expected first color to be set correctly")
	}

	if gradient.Colors[1] != color2 {
		t.Fatal("expected second color to be set correctly")
	}

	if gradient.RotationAngle != 1.57 {
		t.Fatalf("expected RotationAngle 1.57, got %f", gradient.RotationAngle)
	}

	if len(gradient.Stops) != 2 {
		t.Fatalf("expected 2 stops, got %d", len(gradient.Stops))
	}

	if gradient.Stops[0] != 0.0 {
		t.Fatalf("expected first stop to be 0.0, got %f", gradient.Stops[0])
	}

	if gradient.Stops[1] != 1.0 {
		t.Fatalf("expected second stop to be 1.0, got %f", gradient.Stops[1])
	}

	if gradient.Begin != duit_props.AlignmentTopLeft {
		t.Fatalf("expected Begin to be %s, got %s", duit_props.AlignmentTopLeft, gradient.Begin)
	}

	if gradient.End != duit_props.AlignmentBottomRight {
		t.Fatalf("expected End to be %s, got %s", duit_props.AlignmentBottomRight, gradient.End)
	}
}

// Tests for complex scenarios
func TestLinearGradient_ComplexGradient(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	
	// Create a gradient with multiple colors and custom stops
	colors := []*duit_props.Color{
		duit_props.NewColorString("#ff0000"), // Red
		duit_props.NewColorString("#00ff00"), // Green
		duit_props.NewColorString("#0000ff"), // Blue
	}
	stops := []float32{0.0, 0.3, 1.0}
	
	gradient = gradient.
		SetColors(colors).
		SetStops(stops).
		SetRotationAngle(0.785). // 45 degrees
		SetBegin(duit_props.AlignmentTopCenter).
		SetEnd(duit_props.AlignmentBottomCenter)

	if len(gradient.Colors) != 3 {
		t.Fatalf("expected 3 colors, got %d", len(gradient.Colors))
	}

	if len(gradient.Stops) != 3 {
		t.Fatalf("expected 3 stops, got %d", len(gradient.Stops))
	}

	if gradient.RotationAngle != 0.785 {
		t.Fatalf("expected RotationAngle 0.785, got %f", gradient.RotationAngle)
	}

	if gradient.Begin != duit_props.AlignmentTopCenter {
		t.Fatalf("expected Begin to be %s, got %s", duit_props.AlignmentTopCenter, gradient.Begin)
	}

	if gradient.End != duit_props.AlignmentBottomCenter {
		t.Fatalf("expected End to be %s, got %s", duit_props.AlignmentBottomCenter, gradient.End)
	}
}

// Tests for overwriting values
func TestLinearGradient_OverwriteValues(t *testing.T) {
	gradient := duit_props.NewLinearGradient()
	
	// Set initial values
	color1 := duit_props.NewColorString("#ff0000")
	gradient = gradient.AddColor(color1).SetRotationAngle(1.0)
	
	// Overwrite with new values
	colors := []*duit_props.Color{
		duit_props.NewColorString("#00ff00"),
		duit_props.NewColorString("#0000ff"),
	}
	gradient = gradient.SetColors(colors).SetRotationAngle(2.0)

	if len(gradient.Colors) != 2 {
		t.Fatalf("expected 2 colors after overwrite, got %d", len(gradient.Colors))
	}

	if gradient.RotationAngle != 2.0 {
		t.Fatalf("expected RotationAngle 2.0 after overwrite, got %f", gradient.RotationAngle)
	}
}
