package duit_props_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for NewWidgetStateProperty constructor
func TestNewWidgetStateProperty(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[string]()

	if property == nil {
		t.Fatal("expected non-nil WidgetStateProperty")
	}

	// Check that all fields are zero values
	if property.Disabled != "" {
		t.Fatal("expected Disabled to be empty string")
	}
	if property.Hovered != "" {
		t.Fatal("expected Hovered to be empty string")
	}
	if property.Focused != "" {
		t.Fatal("expected Focused to be empty string")
	}
	if property.Pressed != "" {
		t.Fatal("expected Pressed to be empty string")
	}
	if property.Dragged != "" {
		t.Fatal("expected Dragged to be empty string")
	}
	if property.Selected != "" {
		t.Fatal("expected Selected to be empty string")
	}
	if property.Error != "" {
		t.Fatal("expected Error to be empty string")
	}
}

// Tests for NewWidgetStateColor constructor
func TestNewWidgetStateColor(t *testing.T) {
	property := duit_props.NewWidgetStateColor()

	if property == nil {
		t.Fatal("expected non-nil WidgetStateColor")
	}

	// Check that all fields are nil
	if property.Disabled != nil {
		t.Fatal("expected Disabled to be nil")
	}
	if property.Hovered != nil {
		t.Fatal("expected Hovered to be nil")
	}
	if property.Focused != nil {
		t.Fatal("expected Focused to be nil")
	}
	if property.Pressed != nil {
		t.Fatal("expected Pressed to be nil")
	}
	if property.Dragged != nil {
		t.Fatal("expected Dragged to be nil")
	}
	if property.Selected != nil {
		t.Fatal("expected Selected to be nil")
	}
	if property.Error != nil {
		t.Fatal("expected Error to be nil")
	}
}

// Tests for NewWidgetStateEdgeInsets constructor
func TestNewWidgetStateEdgeInsets(t *testing.T) {
	property := duit_props.NewWidgetStateEdgeInsets()

	if property == nil {
		t.Fatal("expected non-nil WidgetStateEdgeInsets")
	}

	// Check that all fields are nil
	if property.Disabled != nil {
		t.Fatal("expected Disabled to be nil")
	}
	if property.Hovered != nil {
		t.Fatal("expected Hovered to be nil")
	}
}

// Tests for NewWidgetStateFloat32 constructor
func TestNewWidgetStateFloat32(t *testing.T) {
	property := duit_props.NewWidgetStateFloat32()

	if property == nil {
		t.Fatal("expected non-nil WidgetStateFloat32")
	}

	// Check that all fields are zero values
	if property.Disabled != 0 {
		t.Fatal("expected Disabled to be 0")
	}
	if property.Hovered != 0 {
		t.Fatal("expected Hovered to be 0")
	}
}

// Tests for NewWidgetStateBorderSide constructor
func TestNewWidgetStateBorderSide(t *testing.T) {
	property := duit_props.NewWidgetStateBorderSide()

	if property == nil {
		t.Fatal("expected non-nil WidgetStateBorderSide")
	}
}

// Tests for NewWidgetStateShapeBorder constructor
func TestNewWidgetStateShapeBorder(t *testing.T) {
	property := duit_props.NewWidgetStateShapeBorder()

	if property == nil {
		t.Fatal("expected non-nil WidgetStateShapeBorder")
	}

	// Check that all fields are nil
	if property.Disabled != nil {
		t.Fatal("expected Disabled to be nil")
	}
}

// Tests for NewWidgetStateTextStyle constructor
func TestNewWidgetStateTextStyle(t *testing.T) {
	property := duit_props.NewWidgetStateTextStyle()

	if property == nil {
		t.Fatal("expected non-nil WidgetStateTextStyle")
	}
}

// Tests for NewWidgetStateVisualDensity constructor
func TestNewWidgetStateVisualDensity(t *testing.T) {
	property := duit_props.NewWidgetStateVisualDensity()

	if property == nil {
		t.Fatal("expected non-nil WidgetStateVisualDensity")
	}
}

// Tests for SetDisabled method
func TestWidgetStateProperty_SetDisabled(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[string]()
	testValue := "disabled"
	
	result := property.SetDisabled(testValue)

	if result != property {
		t.Fatal("expected SetDisabled to return the same instance for chaining")
	}

	if property.Disabled != testValue {
		t.Fatalf("expected Disabled to be '%s', got '%s'", testValue, property.Disabled)
	}
}

// Tests for SetHovered method
func TestWidgetStateProperty_SetHovered(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[string]()
	testValue := "hovered"
	
	result := property.SetHovered(testValue)

	if result != property {
		t.Fatal("expected SetHovered to return the same instance for chaining")
	}

	if property.Hovered != testValue {
		t.Fatalf("expected Hovered to be '%s', got '%s'", testValue, property.Hovered)
	}
}

// Tests for SetFocused method
func TestWidgetStateProperty_SetFocused(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[string]()
	testValue := "focused"
	
	result := property.SetFocused(testValue)

	if result != property {
		t.Fatal("expected SetFocused to return the same instance for chaining")
	}

	if property.Focused != testValue {
		t.Fatalf("expected Focused to be '%s', got '%s'", testValue, property.Focused)
	}
}

// Tests for SetPressed method
func TestWidgetStateProperty_SetPressed(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[string]()
	testValue := "pressed"
	
	result := property.SetPressed(testValue)

	if result != property {
		t.Fatal("expected SetPressed to return the same instance for chaining")
	}

	if property.Pressed != testValue {
		t.Fatalf("expected Pressed to be '%s', got '%s'", testValue, property.Pressed)
	}
}

// Tests for SetDragged method
func TestWidgetStateProperty_SetDragged(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[string]()
	testValue := "dragged"
	
	result := property.SetDragged(testValue)

	if result != property {
		t.Fatal("expected SetDragged to return the same instance for chaining")
	}

	if property.Dragged != testValue {
		t.Fatalf("expected Dragged to be '%s', got '%s'", testValue, property.Dragged)
	}
}

// Tests for SetSelected method
func TestWidgetStateProperty_SetSelected(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[string]()
	testValue := "selected"
	
	result := property.SetSelected(testValue)

	if result != property {
		t.Fatal("expected SetSelected to return the same instance for chaining")
	}

	if property.Selected != testValue {
		t.Fatalf("expected Selected to be '%s', got '%s'", testValue, property.Selected)
	}
}

// Tests for SetError method
func TestWidgetStateProperty_SetError(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[string]()
	testValue := "error"
	
	result := property.SetError(testValue)

	if result != property {
		t.Fatal("expected SetError to return the same instance for chaining")
	}

	if property.Error != testValue {
		t.Fatalf("expected Error to be '%s', got '%s'", testValue, property.Error)
	}
}

// Tests for method chaining
func TestWidgetStateProperty_MethodChaining(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[string]().
		SetDisabled("disabled").
		SetHovered("hovered").
		SetFocused("focused").
		SetPressed("pressed").
		SetDragged("dragged").
		SetSelected("selected").
		SetError("error")

	if property.Disabled != "disabled" {
		t.Fatal("expected Disabled to be 'disabled'")
	}
	if property.Hovered != "hovered" {
		t.Fatal("expected Hovered to be 'hovered'")
	}
	if property.Focused != "focused" {
		t.Fatal("expected Focused to be 'focused'")
	}
	if property.Pressed != "pressed" {
		t.Fatal("expected Pressed to be 'pressed'")
	}
	if property.Dragged != "dragged" {
		t.Fatal("expected Dragged to be 'dragged'")
	}
	if property.Selected != "selected" {
		t.Fatal("expected Selected to be 'selected'")
	}
	if property.Error != "error" {
		t.Fatal("expected Error to be 'error'")
	}
}

// Tests with different types
func TestWidgetStateProperty_IntType(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[int]()
	testValue := 42
	
	property.SetDisabled(testValue)

	if property.Disabled != testValue {
		t.Fatalf("expected Disabled to be %d, got %d", testValue, property.Disabled)
	}
}

func TestWidgetStateProperty_BoolType(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[bool]()
	testValue := true
	
	property.SetFocused(testValue)

	if property.Focused != testValue {
		t.Fatalf("expected Focused to be %t, got %t", testValue, property.Focused)
	}
}

// Tests with WidgetStateColor
func TestWidgetStateColor_SettersWithColor(t *testing.T) {
	property := duit_props.NewWidgetStateColor()
	redColor := duit_props.NewColorString("#ff0000")
	blueColor := duit_props.NewColorString("#0000ff")
	
	property.SetDisabled(redColor).SetHovered(blueColor)

	if property.Disabled != redColor {
		t.Fatal("expected Disabled color to be set")
	}
	if property.Hovered != blueColor {
		t.Fatal("expected Hovered color to be set")
	}
}

// Tests with WidgetStateFloat32
func TestWidgetStateFloat32_SettersWithFloat(t *testing.T) {
	property := duit_props.NewWidgetStateFloat32()
	
	property.SetDisabled(0.5).SetHovered(1.0)

	if property.Disabled != 0.5 {
		t.Fatalf("expected Disabled to be 0.5, got %f", property.Disabled)
	}
	if property.Hovered != 1.0 {
		t.Fatalf("expected Hovered to be 1.0, got %f", property.Hovered)
	}
}

// Test overwriting values
func TestWidgetStateProperty_OverwriteValues(t *testing.T) {
	property := duit_props.NewWidgetStateProperty[string]()
	
	property.SetDisabled("first")
	property.SetDisabled("second")

	if property.Disabled != "second" {
		t.Fatalf("expected Disabled to be 'second', got '%s'", property.Disabled)
	}
}
