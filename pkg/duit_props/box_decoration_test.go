package duit_props_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for NewBoxDecoration constructor
func TestNewBoxDecoration(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()

	if decoration == nil {
		t.Fatal("expected non-nil BoxDecoration")
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for valid empty decoration, got:", err)
	}
}

// Tests for SetColor method
func TestBoxDecoration_SetColor(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	color := duit_props.NewColorString("#ff0000")
	decoration = decoration.SetColor(color)

	if decoration.Color != color {
		t.Fatal("expected Color to be set")
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for valid color, got:", err)
	}
}

func TestBoxDecoration_SetColor_Nil(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	decoration = decoration.SetColor(nil)

	if decoration.Color != nil {
		t.Fatal("expected Color to be nil")
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for nil color, got:", err)
	}
}

// Tests for SetBorder method
func TestBoxDecoration_SetBorder(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	border := &duit_props.BorderSide{}
	decoration = decoration.SetBorder(border)

	if decoration.Border != border {
		t.Fatal("expected Border to be set")
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for valid border, got:", err)
	}
}

func TestBoxDecoration_SetBorder_Nil(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	decoration = decoration.SetBorder(nil)

	if decoration.Border != nil {
		t.Fatal("expected Border to be nil")
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for nil border, got:", err)
	}
}

// Tests for SetBorderRadius method
func TestBoxDecoration_SetBorderRadius(t *testing.T) {
	borderRadius := duit_props.NewBorderRadiusAll(duit_props.NewRadiusCircular(10.0))
	decoration := duit_props.NewBoxDecoration()
	decoration = decoration.SetBorderRadius(borderRadius)

	if decoration.BorderRadius != borderRadius {
		t.Fatalf("expected BorderRadius 10.0, got %v", decoration.BorderRadius)
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for valid border radius, got:", err)
	}
}

func TestBoxDecoration_SetBorderRadius_Zero(t *testing.T) {
	borderRadius := duit_props.NewBorderRadiusAll(duit_props.NewRadiusCircular(0.0))
	decoration := duit_props.NewBoxDecoration()
	decoration = decoration.SetBorderRadius(borderRadius)

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for zero border radius, got:", err)
	}
}

// Tests for SetShape method
func TestBoxDecoration_SetShape(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	decoration = decoration.SetShape(duit_props.BoxShapeCircle)

	if decoration.Shape != duit_props.BoxShapeCircle {
		t.Fatalf("expected Shape %s, got %s", duit_props.BoxShapeCircle, decoration.Shape)
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for valid shape, got:", err)
	}
}

func TestBoxDecoration_SetShape_Rectangle(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	decoration = decoration.SetShape(duit_props.BoxShapeRectangle)

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for rectangle shape, got:", err)
	}
}

// Tests for SetBoxShadow method
func TestBoxDecoration_SetBoxShadow(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	shadows := []duit_props.BoxShadow{
		{BlurRadius: 5.0, SpreadRadius: 2.0},
		{BlurRadius: 10.0, SpreadRadius: 1.0},
	}
	decoration = decoration.SetBoxShadow(shadows)

	if len(decoration.BoxShadow) != 2 {
		t.Fatalf("expected 2 shadows, got %d", len(decoration.BoxShadow))
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for valid box shadows, got:", err)
	}
}

func TestBoxDecoration_SetBoxShadow_Empty(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	decoration = decoration.SetBoxShadow([]duit_props.BoxShadow{})

	if len(decoration.BoxShadow) != 0 {
		t.Fatal("expected empty box shadows")
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for empty box shadows, got:", err)
	}
}

func TestBoxDecoration_SetBoxShadow_Nil(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	decoration = decoration.SetBoxShadow(nil)

	if decoration.BoxShadow != nil {
		t.Fatal("expected nil box shadows")
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for nil box shadows, got:", err)
	}
}

// Tests for AddBoxShadow method
func TestBoxDecoration_AddBoxShadow(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	shadow := duit_props.BoxShadow{BlurRadius: 5.0, SpreadRadius: 2.0}
	decoration = decoration.AddBoxShadow(shadow)

	if len(decoration.BoxShadow) != 1 {
		t.Fatalf("expected 1 shadow, got %d", len(decoration.BoxShadow))
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for valid added shadow, got:", err)
	}
}

func TestBoxDecoration_AddBoxShadow_Multiple(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	shadow1 := duit_props.BoxShadow{BlurRadius: 5.0, SpreadRadius: 2.0}
	shadow2 := duit_props.BoxShadow{BlurRadius: 10.0, SpreadRadius: 1.0}
	decoration = decoration.AddBoxShadow(shadow1).AddBoxShadow(shadow2)

	if len(decoration.BoxShadow) != 2 {
		t.Fatalf("expected 2 shadows, got %d", len(decoration.BoxShadow))
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for multiple added shadows, got:", err)
	}
}

// Tests for SetGradient method
func TestBoxDecoration_SetGradient(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	gradient := &duit_props.LinearGradient{
		Colors: []duit_props.Color{
			*duit_props.NewColorString("#ff0000"),
			*duit_props.NewColorString("#0000ff"),
		},
	}
	decoration = decoration.SetGradient(gradient)

	if decoration.Gradient != gradient {
		t.Fatal("expected Gradient to be set")
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for valid gradient, got:", err)
	}
}

func TestBoxDecoration_SetGradient_Nil(t *testing.T) {
	decoration := duit_props.NewBoxDecoration()
	decoration = decoration.SetGradient(nil)

	if decoration.Gradient != nil {
		t.Fatal("expected Gradient to be nil")
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for nil gradient, got:", err)
	}
}

// Tests for Validate method with nil receiver
func TestBoxDecoration_Validate_NilReceiver(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic on nil receiver, but did not panic")
		}
	}()

	var decoration *duit_props.BoxDecoration

	decoration.Validate()
}

// Tests for method chaining
func TestBoxDecoration_MethodChaining(t *testing.T) {
	borderRadius := duit_props.NewBorderRadiusAll(duit_props.NewRadiusCircular(10.0))
	decoration := duit_props.NewBoxDecoration().
		SetColor(duit_props.NewColorString("#ff0000")).
		SetBorderRadius(borderRadius).
		SetShape(duit_props.BoxShapeCircle).
		AddBoxShadow(duit_props.BoxShadow{BlurRadius: 5.0})

	if decoration.Color == nil {
		t.Fatal("expected Color to be set")
	}

	if decoration.BorderRadius != borderRadius {
		t.Fatalf("expected BorderRadius 10.0, got %f", decoration.BorderRadius)
	}
	if decoration.Shape != duit_props.BoxShapeCircle {
		t.Fatalf("expected Shape %s, got %s", duit_props.BoxShapeCircle, decoration.Shape)
	}
	if len(decoration.BoxShadow) != 1 {
		t.Fatalf("expected 1 shadow, got %d", len(decoration.BoxShadow))
	}

	err := decoration.Validate()
	if err != nil {
		t.Fatal("expected no error for chained method calls, got:", err)
	}
}
