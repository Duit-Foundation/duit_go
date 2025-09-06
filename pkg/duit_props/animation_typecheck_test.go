package duit_props_test

import (
	"fmt"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

var (
	textStyleBeginString = duit_props.TextStyle{
		Color: duit_props.NewColorString("#ff00ff"),
	}
	textStyleEndString = duit_props.TextStyle{
		Color: duit_props.NewColorString("#00ff00"),
	}

	textStyleBeginRGBO = duit_props.TextStyle{
		Color: duit_props.NewColorRGBO([3]uint8{10, 20, 30}, 1),
	}
	textStyleEndRGBO = duit_props.TextStyle{
		Color: duit_props.NewColorRGBO([3]uint8{30, 20, 10}, 1),
	}

	decorationBeginString = duit_props.BoxDecoration{
		Color: duit_props.NewColorString("#123456"),
	}
	decorationEndString = duit_props.BoxDecoration{
		Color: duit_props.NewColorString("#abcdef"),
	}

	decorationBeginRGBO = duit_props.BoxDecoration{
		Color: duit_props.NewColorRGBO([3]uint8{1, 2, 3}, 1),
	}
	decorationEndRGBO = duit_props.BoxDecoration{
		Color: duit_props.NewColorRGBO([3]uint8{3, 2, 1}, 1),
	}
)

func TestCheckTweenType(t *testing.T) {

	arr := []any{
		// float
		duit_props.Tween("opacity", 0.0, 1.0, 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		// colors
		duit_props.ColorTween("color", duit_props.NewColorString("#000000"), duit_props.NewColorString("#ffffff"), 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		duit_props.ColorTween("color", duit_props.NewColorRGBO([3]uint8{0, 0, 0}, 1), duit_props.NewColorRGBO([3]uint8{255, 255, 255}, 1), 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		// text styles
		duit_props.TextStyleTween("textStyle", &textStyleBeginString, &textStyleEndString, 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		duit_props.TextStyleTween("textStyle", &textStyleBeginRGBO, &textStyleEndRGBO, 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		// decorations
		duit_props.DecorationTween("decoration", &decorationBeginString, &decorationEndString, 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		duit_props.DecorationTween("decoration", &decorationBeginRGBO, &decorationEndRGBO, 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		// alignment
		duit_props.AlignmentTween("alignment", "center", "topLeft", 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		// edge insets
		duit_props.EdgeInsetsTween("paddingAll", duit_props.NewEdgeInsetsAll(4), duit_props.NewEdgeInsetsAll(8), 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		duit_props.EdgeInsetsTween("paddingLTRB", duit_props.NewEdgeInsetsLTRB(1, 2, 3, 4), duit_props.NewEdgeInsetsLTRB(4, 3, 2, 1), 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		duit_props.EdgeInsetsTween("paddingSym", duit_props.NewEdgeInsetsSymmetric(5, 10), duit_props.NewEdgeInsetsSymmetric(10, 5), 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		// constraints and size
		duit_props.BoxConstraintsTween("constraints", duit_props.BoxConstraints{MinWidth: 0, MinHeight: 0}, duit_props.BoxConstraints{MaxWidth: 100, MaxHeight: 100}, 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		duit_props.SizeTween("size", duit_props.Size{Width: 10, Height: 20}, duit_props.Size{Width: 20, Height: 40}, 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		// border
		duit_props.BorderTween("border", duit_props.BorderOutline, duit_props.BorderUnderline, 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		// group
		duit_props.TweenGroup([]any{
			duit_props.Tween("opacity", 0.0, 1.0, 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
		}, "group1", 1, duit_props.AnimationTriggerOnEnter, duit_props.AnimationMethodForward, false, ""),
	}

	for _, tween := range arr {
		fmt.Println(tween)
		if err := duit_props.CheckTweenType(tween); err != nil {
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
		err := duit_props.CheckTweenType(invalidType)
		if err == nil {
			t.Fatalf("expected error for invalid type at index %d (%T), but got nil", i, invalidType)
		}
	}
}

func TestCheckTweenType_NilValue(t *testing.T) {
	err := duit_props.CheckTweenType(nil)
	if err != nil {
		t.Fatalf("expected no error for nil value, got: %s", err.Error())
	}
}
