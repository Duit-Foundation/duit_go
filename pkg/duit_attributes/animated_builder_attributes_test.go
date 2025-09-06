package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestAnimatedBuilderAttributes_Validate_ValidAttributes(t *testing.T) {
	tweens := []any{
		duit_props.Tween("opacity", 0.0, 1.0, 300, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, duit_props.CurveEase),
	}

	attrs := &duit_attributes.AnimatedBuilderAttributes{
		Tweens:       tweens,
		PersistentId: "test-id",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestAnimatedBuilderAttributes_Validate_NilTweens(t *testing.T) {
	attrs := &duit_attributes.AnimatedBuilderAttributes{
		Tweens:       nil,
		PersistentId: "test-id",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil tweens")
	}

	expected := "tweens property is required and must be non-empty"
	if err.Error() != expected {
		t.Fatalf("expected error message '%s', got '%s'", expected, err.Error())
	}
}

func TestAnimatedBuilderAttributes_Validate_EmptyTweens(t *testing.T) {
	attrs := &duit_attributes.AnimatedBuilderAttributes{
		Tweens:       []any{},
		PersistentId: "test-id",
	}

	err := attrs.Validate()
	// Empty slice doesn't trigger error, only nil does
	if err != nil {
		t.Fatal("expected no error for empty tweens slice, got:", err)
	}
}

func TestAnimatedBuilderAttributes_Validate_ValidTweenTypes(t *testing.T) {
	testCases := []struct {
		name  string
		tween any
	}{
		{
			name:  "DoubleTween",
			tween: duit_props.Tween("opacity", 0.0, 1.0, 300, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, duit_props.CurveEase),
		},
		{
			name:  "ColorTween",
			tween: duit_props.ColorTween("color", duit_props.NewColorString("#FF0000"), duit_props.NewColorString("#00FF00"), 300, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, duit_props.CurveEase),
		},
		{
			name: "SizeTween",
			tween: duit_props.SizeTween("size",
				duit_props.Size{Width: 100.0, Height: 100.0},
				duit_props.Size{Width: 200.0, Height: 200.0},
				300, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, duit_props.CurveEase),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tweens := []any{tc.tween}

			attrs := &duit_attributes.AnimatedBuilderAttributes{
				Tweens:       tweens,
				PersistentId: "test-id",
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for %s, got: %v", tc.name, err)
			}
		})
	}
}

func TestAnimatedBuilderAttributes_Validate_MultipleTweens(t *testing.T) {
	tweens := []any{
		duit_props.Tween("opacity", 0.0, 1.0, 300, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, duit_props.CurveEase),
		duit_props.ColorTween("color", duit_props.NewColorString("#FF0000"), duit_props.NewColorString("#00FF00"), 300, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, duit_props.CurveEase),
		duit_props.SizeTween("size",
			duit_props.Size{Width: 100.0, Height: 100.0},
			duit_props.Size{Width: 200.0, Height: 200.0},
			300, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, duit_props.CurveEase),
	}

	attrs := &duit_attributes.AnimatedBuilderAttributes{
		Tweens:       tweens,
		PersistentId: "test-id",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for multiple valid tweens, got:", err)
	}
}

func TestAnimatedBuilderAttributes_Validate_InvalidTweenType(t *testing.T) {
	tweens := []any{
		"invalid-tween-type",
	}

	attrs := &duit_attributes.AnimatedBuilderAttributes{
		Tweens:       tweens,
		PersistentId: "test-id",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for invalid tween type")
	}

	// The error message will come from duit_props.CheckTweenType
	if !strings.Contains(err.Error(), "invalid tween type") {
		t.Fatalf("expected error to contain 'invalid tween type', got: %s", err.Error())
	}
}

func TestAnimatedBuilderAttributes_Validate_WithoutPersistentId(t *testing.T) {
	tweens := []any{
		duit_props.Tween("opacity", 0.0, 1.0, 300, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, duit_props.CurveEase),
	}

	attrs := &duit_attributes.AnimatedBuilderAttributes{
		Tweens: tweens,
		// PersistentId is omitted
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error when PersistentId is omitted, got:", err)
	}
}

func TestAnimatedBuilderAttributes_MarshalJSON_WithPersistentId(t *testing.T) {
	tweens := []any{
		duit_props.Tween("opacity", 0.0, 1.0, 300, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, duit_props.CurveEase),
	}

	attrs := &duit_attributes.AnimatedBuilderAttributes{
		Tweens:       tweens,
		PersistentId: "test-persistent-id",
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"persistentId":"test-persistent-id"`) {
		t.Fatalf("expected JSON to contain '\"persistentId\":\"test-persistent-id\"', got: %s", jsonStr)
	}
}

func TestAnimatedBuilderAttributes_MarshalJSON_WithoutPersistentId(t *testing.T) {
	tweens := []any{
		duit_props.Tween("opacity", 0.0, 1.0, 300, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, duit_props.CurveEase),
	}

	attrs := &duit_attributes.AnimatedBuilderAttributes{
		Tweens: tweens,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	// PersistentId should be omitted from JSON due to omitempty tag
	if strings.Contains(jsonStr, `"persistentId"`) {
		t.Fatalf("expected JSON to not contain 'persistentId' field when empty, got: %s", jsonStr)
	}
}

func TestAnimatedBuilderAttributes_MarshalJSON_WithTweens(t *testing.T) {
	tweens := []any{
		duit_props.Tween("opacity", 0.0, 1.0, 300, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, duit_props.CurveEase),
	}

	attrs := &duit_attributes.AnimatedBuilderAttributes{
		Tweens:       tweens,
		PersistentId: "test-id",
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"tweens"`) {
		t.Fatalf("expected JSON to contain 'tweens' field, got: %s", jsonStr)
	}
}
