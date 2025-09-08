package duit_props_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

// Tests for NewButtonStyle constructor
func TestNewButtonStyle(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()

	if buttonStyle == nil {
		t.Fatal("expected non-nil ButtonStyle")
	}

	// Check that all fields are nil/zero values
	if buttonStyle.TextStyle != nil {
		t.Fatal("expected TextStyle to be nil")
	}
	if buttonStyle.BackgroundColor != nil {
		t.Fatal("expected BackgroundColor to be nil")
	}
	if buttonStyle.ForegroundColor != nil {
		t.Fatal("expected ForegroundColor to be nil")
	}
	if buttonStyle.OverlayColor != nil {
		t.Fatal("expected OverlayColor to be nil")
	}
	if buttonStyle.ShadowColor != nil {
		t.Fatal("expected ShadowColor to be nil")
	}
	if buttonStyle.IconColor != nil {
		t.Fatal("expected IconColor to be nil")
	}
	if buttonStyle.Elevation != nil {
		t.Fatal("expected Elevation to be nil")
	}
	if buttonStyle.IconSize != nil {
		t.Fatal("expected IconSize to be nil")
	}
	if buttonStyle.Padding != nil {
		t.Fatal("expected Padding to be nil")
	}
	if buttonStyle.MinimumSize != nil {
		t.Fatal("expected MinimumSize to be nil")
	}
	if buttonStyle.MaximumSize != nil {
		t.Fatal("expected MaximumSize to be nil")
	}
	if buttonStyle.Side != nil {
		t.Fatal("expected Side to be nil")
	}
	if buttonStyle.Shape != nil {
		t.Fatal("expected Shape to be nil")
	}
	if buttonStyle.VisualDensity != nil {
		t.Fatal("expected VisualDensity to be nil")
	}
	if buttonStyle.TapTargetSize != "" {
		t.Fatal("expected TapTargetSize to be empty string")
	}
	if buttonStyle.AnimationDuration != 0 {
		t.Fatal("expected AnimationDuration to be zero")
	}
}

// Tests for SetTextStyle method
func TestButtonStyle_SetTextStyle(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	textStyle := duit_props.NewWidgetStateTextStyle()
	
	result := buttonStyle.SetTextStyle(textStyle)

	if result != buttonStyle {
		t.Fatal("expected SetTextStyle to return the same instance for chaining")
	}

	if buttonStyle.TextStyle != textStyle {
		t.Fatal("expected TextStyle to be set")
	}
}

// Tests for SetBackgroundColor method
func TestButtonStyle_SetBackgroundColor(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	backgroundColor := duit_props.NewWidgetStateColor()
	
	result := buttonStyle.SetBackgroundColor(backgroundColor)

	if result != buttonStyle {
		t.Fatal("expected SetBackgroundColor to return the same instance for chaining")
	}

	if buttonStyle.BackgroundColor != backgroundColor {
		t.Fatal("expected BackgroundColor to be set")
	}
}

// Tests for SetForegroundColor method
func TestButtonStyle_SetForegroundColor(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	foregroundColor := duit_props.NewWidgetStateColor()
	
	result := buttonStyle.SetForegroundColor(foregroundColor)

	if result != buttonStyle {
		t.Fatal("expected SetForegroundColor to return the same instance for chaining")
	}

	if buttonStyle.ForegroundColor != foregroundColor {
		t.Fatal("expected ForegroundColor to be set")
	}
}

// Tests for SetOverlayColor method
func TestButtonStyle_SetOverlayColor(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	overlayColor := duit_props.NewWidgetStateColor()
	
	result := buttonStyle.SetOverlayColor(overlayColor)

	if result != buttonStyle {
		t.Fatal("expected SetOverlayColor to return the same instance for chaining")
	}

	if buttonStyle.OverlayColor != overlayColor {
		t.Fatal("expected OverlayColor to be set")
	}
}

// Tests for SetShadowColor method
func TestButtonStyle_SetShadowColor(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	shadowColor := duit_props.NewWidgetStateColor()
	
	result := buttonStyle.SetShadowColor(shadowColor)

	if result != buttonStyle {
		t.Fatal("expected SetShadowColor to return the same instance for chaining")
	}

	if buttonStyle.ShadowColor != shadowColor {
		t.Fatal("expected ShadowColor to be set")
	}
}

// Tests for SetIconColor method
func TestButtonStyle_SetIconColor(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	iconColor := duit_props.NewWidgetStateColor()
	
	result := buttonStyle.SetIconColor(iconColor)

	if result != buttonStyle {
		t.Fatal("expected SetIconColor to return the same instance for chaining")
	}

	if buttonStyle.IconColor != iconColor {
		t.Fatal("expected IconColor to be set")
	}
}

// Tests for SetElevation method
func TestButtonStyle_SetElevation(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	elevation := duit_props.NewWidgetStateFloat32()
	
	result := buttonStyle.SetElevation(elevation)

	if result != buttonStyle {
		t.Fatal("expected SetElevation to return the same instance for chaining")
	}

	if buttonStyle.Elevation != elevation {
		t.Fatal("expected Elevation to be set")
	}
}

// Tests for SetIconSize method
func TestButtonStyle_SetIconSize(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	iconSize := duit_props.NewWidgetStateFloat32()
	
	result := buttonStyle.SetIconSize(iconSize)

	if result != buttonStyle {
		t.Fatal("expected SetIconSize to return the same instance for chaining")
	}

	if buttonStyle.IconSize != iconSize {
		t.Fatal("expected IconSize to be set")
	}
}

// Tests for SetPadding method
func TestButtonStyle_SetPadding(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	padding := duit_props.NewWidgetStateEdgeInsets()
	
	result := buttonStyle.SetPadding(padding)

	if result != buttonStyle {
		t.Fatal("expected SetPadding to return the same instance for chaining")
	}

	if buttonStyle.Padding != padding {
		t.Fatal("expected Padding to be set")
	}
}

// Tests for SetMinimumSize method
func TestButtonStyle_SetMinimumSize(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	minimumSize := duit_props.NewWidgetStateSize()
	
	result := buttonStyle.SetMinimumSize(minimumSize)

	if result != buttonStyle {
		t.Fatal("expected SetMinimumSize to return the same instance for chaining")
	}

	if buttonStyle.MinimumSize != minimumSize {
		t.Fatal("expected MinimumSize to be set")
	}
}

// Tests for SetMaximumSize method
func TestButtonStyle_SetMaximumSize(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	maximumSize := duit_props.NewWidgetStateSize()
	
	result := buttonStyle.SetMaximumSize(maximumSize)

	if result != buttonStyle {
		t.Fatal("expected SetMaximumSize to return the same instance for chaining")
	}

	if buttonStyle.MaximumSize != maximumSize {
		t.Fatal("expected MaximumSize to be set")
	}
}

// Tests for SetSide method
func TestButtonStyle_SetSide(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	side := duit_props.NewWidgetStateBorderSide()
	
	result := buttonStyle.SetSide(side)

	if result != buttonStyle {
		t.Fatal("expected SetSide to return the same instance for chaining")
	}

	if buttonStyle.Side != side {
		t.Fatal("expected Side to be set")
	}
}

// Tests for SetShape method
func TestButtonStyle_SetShape(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	shape := duit_props.NewWidgetStateShapeBorder()
	
	result := buttonStyle.SetShape(shape)

	if result != buttonStyle {
		t.Fatal("expected SetShape to return the same instance for chaining")
	}

	if buttonStyle.Shape != shape {
		t.Fatal("expected Shape to be set")
	}
}

// Tests for SetVisualDensity method
func TestButtonStyle_SetVisualDensity(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	visualDensity := duit_props.NewWidgetStateVisualDensity()
	
	result := buttonStyle.SetVisualDensity(visualDensity)

	if result != buttonStyle {
		t.Fatal("expected SetVisualDensity to return the same instance for chaining")
	}

	if buttonStyle.VisualDensity != visualDensity {
		t.Fatal("expected VisualDensity to be set")
	}
}

// Tests for SetTapTargetSize method
func TestButtonStyle_SetTapTargetSize(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	tapTargetSize := duit_props.MaterialTapTargetSizePadded
	
	result := buttonStyle.SetTapTargetSize(tapTargetSize)

	if result != buttonStyle {
		t.Fatal("expected SetTapTargetSize to return the same instance for chaining")
	}

	if buttonStyle.TapTargetSize != tapTargetSize {
		t.Fatalf("expected TapTargetSize to be %v, got %v", tapTargetSize, buttonStyle.TapTargetSize)
	}
}

// Tests for SetAnimationDuration method
func TestButtonStyle_SetAnimationDuration(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	animationDuration := uint(300)
	
	result := buttonStyle.SetAnimationDuration(animationDuration)

	if result != buttonStyle {
		t.Fatal("expected SetAnimationDuration to return the same instance for chaining")
	}

	if buttonStyle.AnimationDuration != animationDuration {
		t.Fatalf("expected AnimationDuration to be %d, got %d", animationDuration, buttonStyle.AnimationDuration)
	}
}

// Tests for method chaining
func TestButtonStyle_MethodChaining(t *testing.T) {
	textStyle := duit_props.NewWidgetStateTextStyle()
	backgroundColor := duit_props.NewWidgetStateColor()
	foregroundColor := duit_props.NewWidgetStateColor()
	elevation := duit_props.NewWidgetStateFloat32()
	padding := duit_props.NewWidgetStateEdgeInsets()
	tapTargetSize := duit_props.MaterialTapTargetSizePadded
	animationDuration := uint(250)

	buttonStyle := duit_props.NewButtonStyle().
		SetTextStyle(textStyle).
		SetBackgroundColor(backgroundColor).
		SetForegroundColor(foregroundColor).
		SetElevation(elevation).
		SetPadding(padding).
		SetTapTargetSize(tapTargetSize).
		SetAnimationDuration(animationDuration)

	if buttonStyle.TextStyle != textStyle {
		t.Fatal("expected TextStyle to be set")
	}
	if buttonStyle.BackgroundColor != backgroundColor {
		t.Fatal("expected BackgroundColor to be set")
	}
	if buttonStyle.ForegroundColor != foregroundColor {
		t.Fatal("expected ForegroundColor to be set")
	}
	if buttonStyle.Elevation != elevation {
		t.Fatal("expected Elevation to be set")
	}
	if buttonStyle.Padding != padding {
		t.Fatal("expected Padding to be set")
	}
	if buttonStyle.TapTargetSize != tapTargetSize {
		t.Fatal("expected TapTargetSize to be set")
	}
	if buttonStyle.AnimationDuration != animationDuration {
		t.Fatal("expected AnimationDuration to be set")
	}
}

// Test overwriting values
func TestButtonStyle_OverwriteValues(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	firstDuration := uint(100)
	secondDuration := uint(200)
	
	buttonStyle.SetAnimationDuration(firstDuration)
	buttonStyle.SetAnimationDuration(secondDuration)

	if buttonStyle.AnimationDuration != secondDuration {
		t.Fatalf("expected AnimationDuration to be %d, got %d", secondDuration, buttonStyle.AnimationDuration)
	}
}

// Test with nil values
func TestButtonStyle_SetNilValues(t *testing.T) {
	buttonStyle := duit_props.NewButtonStyle()
	
	// Setting nil values should work without panicking
	buttonStyle.SetTextStyle(nil)
	buttonStyle.SetBackgroundColor(nil)
	buttonStyle.SetElevation(nil)

	if buttonStyle.TextStyle != nil {
		t.Fatal("expected TextStyle to be nil")
	}
	if buttonStyle.BackgroundColor != nil {
		t.Fatal("expected BackgroundColor to be nil")
	}
	if buttonStyle.Elevation != nil {
		t.Fatal("expected Elevation to be nil")
	}
}
