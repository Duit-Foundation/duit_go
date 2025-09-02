package duit_animations_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
)

var (
	textStyleBeginString = duit_text_properties.TextStyle[duit_props.ColorString]{
		Color: duit_props.ColorString("#ff00ff"),
	}
	textStyleEndString = duit_text_properties.TextStyle[duit_props.ColorString]{
		Color: duit_props.ColorString("#00ff00"),
	}

	textStyleBeginRGBO = duit_text_properties.TextStyle[*duit_props.ColorRGBO]{
		Color: &duit_props.ColorRGBO{R: 10, G: 20, B: 30, O: 1},
	}
	textStyleEndRGBO = duit_text_properties.TextStyle[*duit_props.ColorRGBO]{
		Color: &duit_props.ColorRGBO{R: 30, G: 20, B: 10, O: 1},
	}

	decorationBeginString = duit_decoration.BoxDecoration[duit_props.ColorString]{
		Color: duit_props.ColorString("#123456"),
	}
	decorationEndString = duit_decoration.BoxDecoration[duit_props.ColorString]{
		Color: duit_props.ColorString("#abcdef"),
	}

	decorationBeginRGBO = duit_decoration.BoxDecoration[*duit_props.ColorRGBO]{
		Color: &duit_props.ColorRGBO{R: 1, G: 2, B: 3, O: 1},
	}
	decorationEndRGBO = duit_decoration.BoxDecoration[*duit_props.ColorRGBO]{
		Color: &duit_props.ColorRGBO{R: 3, G: 2, B: 1, O: 1},
	}
)

func TestCheckTweenType(t *testing.T) {

	arr := []any{
		// float
		duit_animations.Tween("opacity", 0.0, 1.0, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		// colors
		duit_animations.ColorTween("color", duit_props.ColorString("#000000"), duit_props.ColorString("#ffffff"), 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		duit_animations.ColorTween("color", &duit_props.ColorRGBO{R: 0, G: 0, B: 0, O: 1}, &duit_props.ColorRGBO{R: 255, G: 255, B: 255, O: 1}, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		// text styles
		duit_animations.TextStyleTween("textStyle", textStyleBeginString, textStyleEndString, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		duit_animations.TextStyleTween("textStyle", textStyleBeginRGBO, textStyleEndRGBO, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		// decorations
		duit_animations.DecorationTween("decoration", decorationBeginString, decorationEndString, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		duit_animations.DecorationTween("decoration", decorationBeginRGBO, decorationEndRGBO, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		// alignment
		duit_animations.AlignmentTween("alignment", "center", "topLeft", 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		// edge insets
		duit_animations.EdgeInsetsTween("paddingAll", duit_props.EdgeInsetsAll(4), duit_props.EdgeInsetsAll(8), 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		duit_animations.EdgeInsetsTween("paddingLTRB", duit_props.EdgeInsetsLTRB{1, 2, 3, 4}, duit_props.EdgeInsetsLTRB{4, 3, 2, 1}, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		duit_animations.EdgeInsetsTween("paddingSym", duit_props.EdgeInsetsSymmentric{5, 10}, duit_props.EdgeInsetsSymmentric{10, 5}, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		// constraints and size
		duit_animations.BoxConstraintsTween("constraints", duit_flex.BoxConstraints{MinWidth: 0, MinHeight: 0}, duit_flex.BoxConstraints{MaxWidth: 100, MaxHeight: 100}, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		duit_animations.SizeTween("size", duit_flex.Size{Width: 10, Height: 20}, duit_flex.Size{Width: 20, Height: 40}, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		// border
		duit_animations.BorderTween("border", duit_decoration.Outline, duit_decoration.Underline, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		// group
		duit_animations.TweenGroup([]any{
			duit_animations.Tween("opacity", 0.0, 1.0, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		}, "group1", 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
	}

	for _, tween := range arr {
		if err := duit_animations.CheckTweenType(tween); err != nil {
			t.Fatal(err)
		}
	}
}

func TestCheckTweenType_InvalidTypes(t *testing.T) {
	invalidTypes := []any{
		"invalid string",
		123,
		struct{ field string }{field: "test"},
		[]string{"test"},
		map[string]int{"key": 1},
		func() {},
		&struct{ field int }{field: 42},
	}

	for i, invalidType := range invalidTypes {
		err := duit_animations.CheckTweenType(invalidType)
		if err == nil {
			t.Fatalf("expected error for invalid type at index %d (%T), but got nil", i, invalidType)
		}
		
		if err.Error() != "invalid tween type. Must be instance of tweenBase or tweenGroup or nil" {
			t.Fatalf("expected specific error message for invalid type at index %d, got: %s", i, err.Error())
		}
	}
}

func TestCheckTweenType_NilValue(t *testing.T) {
	err := duit_animations.CheckTweenType(nil)
	if err != nil {
		t.Fatalf("expected no error for nil value, got: %s", err.Error())
	}
}
