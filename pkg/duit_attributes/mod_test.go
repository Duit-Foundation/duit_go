package duit_attributes

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
)

func TestAlignAttributesValidation(t *testing.T) {
	attrs := &AlignAttributes{
		WidthFactor: 1.1,
		ValueReferenceHolder: &ValueReferenceHolder{
			Refs: []ValueRef{},
		},
		AnimatedPropertyOwner: &duit_animations.AnimatedPropertyOwner{
			AffectedProperties: []string{},
		},
		ThemeConsumer: &ThemeConsumer{},
	}

	err := attrs.Validate()

	if err == nil {
		t.Fatal("must throw error")
	}

	attrs = &AlignAttributes{
		Alignment: "center",
	}

	err = attrs.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func TestAnimatedAlignAttributesValidation(t *testing.T) {
	attrs := &AnimatedAlignAttributes{
		Alignment: "center",
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 1,
		},
	}

	err := attrs.Validate()

	if err != nil {
		t.Fatal(err)
	}

	attrs = &AnimatedAlignAttributes{
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 1,
		},
	}

	err = attrs.Validate()

	if err == nil {
		t.Fatal("must throw error")
	}

	attrs = &AnimatedAlignAttributes{
		ImplicitAnimatable: &duit_animations.ImplicitAnimatable{
			Duration: 1,
		},
	}

	err = attrs.Validate()

	if err == nil {
		t.Fatal("must throw error")
	}
}

func TestAnimatedBuilderAttributesValidation(t *testing.T) {
	attrs := &AnimatedBuilderAttributes{
		Tweens: []any{
			duit_animations.Tween("opacity", 0.0, 1.0, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		},
	}

	err := attrs.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func TestAnimatedBuilderAttributes_Validate_EmptyTweens(t *testing.T) {
	attrs := &AnimatedBuilderAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("must throw error for empty tweens")
	}
}

func TestAnimatedBuilderAttributes_Validate_InvalidTweenType(t *testing.T) {
	attrs := &AnimatedBuilderAttributes{
		Tweens: []any{
			"invalid",
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("must throw error for invalid tween type")
	}
}

func TestAnimatedBuilderAttributes_Validate_AllSupportedTweens(t *testing.T) {
	textStyleBeginString := duit_text_properties.TextStyle[duit_color.ColorString]{
		Color: duit_color.ColorString("#ff00ff"),
	}
	textStyleEndString := duit_text_properties.TextStyle[duit_color.ColorString]{
		Color: duit_color.ColorString("#00ff00"),
	}

	textStyleBeginRGBO := duit_text_properties.TextStyle[*duit_color.ColorRGBO]{
		Color: &duit_color.ColorRGBO{R: 10, G: 20, B: 30, O: 1},
	}
	textStyleEndRGBO := duit_text_properties.TextStyle[*duit_color.ColorRGBO]{
		Color: &duit_color.ColorRGBO{R: 30, G: 20, B: 10, O: 1},
	}

	decorationBeginString := duit_decoration.BoxDecoration[duit_color.ColorString]{
		Color: duit_color.ColorString("#123456"),
	}
	decorationEndString := duit_decoration.BoxDecoration[duit_color.ColorString]{
		Color: duit_color.ColorString("#abcdef"),
	}

	decorationBeginRGBO := duit_decoration.BoxDecoration[*duit_color.ColorRGBO]{
		Color: &duit_color.ColorRGBO{R: 1, G: 2, B: 3, O: 1},
	}
	decorationEndRGBO := duit_decoration.BoxDecoration[*duit_color.ColorRGBO]{
		Color: &duit_color.ColorRGBO{R: 3, G: 2, B: 1, O: 1},
	}

	attrs := &AnimatedBuilderAttributes{
		Tweens: []any{
			// float
			duit_animations.Tween("opacity", 0.0, 1.0, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			// colors
			duit_animations.ColorTween("color", duit_color.ColorString("#000000"), duit_color.ColorString("#ffffff"), 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			duit_animations.ColorTween("color", &duit_color.ColorRGBO{R: 0, G: 0, B: 0, O: 1}, &duit_color.ColorRGBO{R: 255, G: 255, B: 255, O: 1}, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			// text styles
			duit_animations.TextStyleTween("textStyle", textStyleBeginString, textStyleEndString, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			duit_animations.TextStyleTween("textStyle", textStyleBeginRGBO, textStyleEndRGBO, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			// decorations
			duit_animations.DecorationTween("decoration", decorationBeginString, decorationEndString, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			duit_animations.DecorationTween("decoration", decorationBeginRGBO, decorationEndRGBO, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			// alignment
			duit_animations.AlignmentTween("alignment", "center", "topLeft", 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			// edge insets
			duit_animations.EdgeInsetsTween("paddingAll", duit_edge_insets.EdgeInsetsAll(4), duit_edge_insets.EdgeInsetsAll(8), 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			duit_animations.EdgeInsetsTween("paddingLTRB", duit_edge_insets.EdgeInsetsLTRB{1, 2, 3, 4}, duit_edge_insets.EdgeInsetsLTRB{4, 3, 2, 1}, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			duit_animations.EdgeInsetsTween("paddingSym", duit_edge_insets.EdgeInsetsSymmentric{5, 10}, duit_edge_insets.EdgeInsetsSymmentric{10, 5}, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			// constraints and size
			duit_animations.BoxConstraintsTween("constraints", duit_flex.BoxConstraints{MinWidth: 0, MinHeight: 0}, duit_flex.BoxConstraints{MaxWidth: 100, MaxHeight: 100}, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			duit_animations.SizeTween("size", duit_flex.Size{Width: 10, Height: 20}, duit_flex.Size{Width: 20, Height: 40}, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			// border
			duit_animations.BorderTween("border", duit_decoration.Outline, duit_decoration.Underline, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			// group
			duit_animations.TweenGroup([]any{
				duit_animations.Tween("opacity", 0.0, 1.0, 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
			}, "group1", 1, duit_animations.OnEnter, duit_animations.Forward, false, ""),
		},
	}

	if err := attrs.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestAnimatedBuilderAttributes_Validate_NilTweenIsOk(t *testing.T) {
	attrs := &AnimatedBuilderAttributes{
		Tweens: []any{nil},
	}

	if err := attrs.Validate(); err != nil {
		t.Fatal(err)
	}
}
