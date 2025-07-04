package duit_attributes

import (
	"encoding/json"
	"testing"

	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_painting"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_utils"
)

func TestBackdropFilter(t *testing.T) {
	blur := duit_painting.NewBlurImageFilter(1, 2, "")

	attr := &BackdropFilterAttributes[duit_painting.BlurImageFilter]{
		Filter:    blur,
		BlendMode: duit_painting.Clear,
	}

	_, err := json.Marshal(attr)

	if err != nil {
		t.Fatal(err)
	}
}

func TestContainerAttributes(t *testing.T) {
	cont := &ContainerAttributes[duit_edge_insets.EdgeInsetsAll, *duit_color.ColorRGBO]{
		Decoration: &duit_decoration.BoxDecoration[*duit_color.ColorRGBO]{
			Color: &duit_color.ColorRGBO{R: 255, G: 0, B: 0, O: 1},
		},
	}

	j, err := json.Marshal(cont)

	if err != nil {
		t.Fatal(err)
	}

	val := map[string]interface{}{}

	err = json.Unmarshal(j, &val)

	if err != nil {
		t.Fatal(err)
	}

	if val["decoration"] == nil {
		t.Fatal("decoration is nil")
	}

	if val["color"] != nil {
		t.Fatal("color is not nil")
	}
}

func TestSafeAreaAttributes(t *testing.T) {
	cont := &SafeAreaAttributes[duit_edge_insets.EdgeInsetsAll]{
		Top:    duit_utils.BoolPtr(false),
		Bottom: duit_utils.BoolPtr(true),
	}

	j, err := json.Marshal(cont)

	if err != nil {
		t.Fatal(err)
	}

	val := map[string]interface{}{}

	err = json.Unmarshal(j, &val)

	if err != nil {
		t.Fatal(err)
	}

	if val["top"] != false {
		t.Fatal("top is nil")
	}

	if val["bottom"] != true {
		t.Fatal("botton is nil")
	}

	if val["right"] != nil {
		t.Fatal("right is not nil")
	}
}

func TestSliverSafeAreaAttributes(t *testing.T) {
	cont := &SliverSafeAreaAttributes[duit_edge_insets.EdgeInsetsAll]{
		Top:    duit_utils.BoolPtr(false),
		Bottom: duit_utils.BoolPtr(true),
		Left:   duit_utils.BoolPtr(true),
		Right:  duit_utils.BoolPtr(false),
	}

	j, err := json.Marshal(cont)

	if err != nil {
		t.Fatal(err)
	}

	val := map[string]interface{}{}

	err = json.Unmarshal(j, &val)

	if err != nil {
		t.Fatal(err)
	}

	if val["top"] != false {
		t.Fatal("top is not false")
	}

	if val["bottom"] != true {
		t.Fatal("bottom is not true")
	}

	if val["left"] != true {
		t.Fatal("left is not true")
	}

	if val["right"] != false {
		t.Fatal("right is not false")
	}
}

func TestGridViewAttributes(t *testing.T) {
	attrsCommonWithValidationErr := &GridViewCommonAttributes[duit_edge_insets.EdgeInsetsAll]{}

	_, err := json.Marshal(attrsCommonWithValidationErr)

	if err == nil {
		t.Fatal(err)
	}

	if attrsCommonWithValidationErr.Constructor != GridCommon {
		t.Fatal("constructor is not GridCommon")
	}

	attrsCountWithValidationErr := &GridViewCountAttributes[duit_edge_insets.EdgeInsetsAll]{}

	_, err = json.Marshal(attrsCountWithValidationErr)

	if err == nil {
		t.Fatal(err)
	}

	if attrsCountWithValidationErr.Constructor != GridCount {
		t.Fatal(attrsCountWithValidationErr.Constructor)
		t.Fatal("constructor is not GridCount")
	}

	attrsExtentWithValidationErr := &GridViewExtentAttributes[duit_edge_insets.EdgeInsetsAll]{}

	_, err = json.Marshal(attrsExtentWithValidationErr)

	if err == nil {
		t.Fatal(err)
	}

	if attrsExtentWithValidationErr.Constructor != GridExtent {
		t.Fatal("constructor is not GridExtent")
	}

	attrsBuilderWithValidationErr := &GridViewBuilderAttributes[duit_edge_insets.EdgeInsetsAll]{}

	_, err = json.Marshal(attrsBuilderWithValidationErr)

	if err == nil {
		t.Fatal(err)
	}

	if attrsBuilderWithValidationErr.Constructor != GridBuilder {
		t.Fatal("constructor is not GridBuilder")
	}

}

func TestAppBarAttributes(t *testing.T) {
	attrs := &AppBarAttributes[duit_color.ColorString, duit_edge_insets.EdgeInsetsAll, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{}

	_, err := json.Marshal(attrs)

	if err != nil {
		t.Fatal(err)
	}

	if *attrs.AutomaticallyImplyLeading == false {
		t.Fatal("automaticallyImplyLeading is not true")
	}

	if *attrs.Primary == false {
		t.Fatal("primary is not true")
	}

	invalidAttrs := &AppBarAttributes[duit_color.ColorString, duit_edge_insets.EdgeInsetsAll, duit_decoration.RoundedRectangleBorder[duit_color.ColorString]]{
		ToolbarOpacity: 1.1,
	}

	_, err = json.Marshal(invalidAttrs)

	if err == nil {
		t.Fatal("must throw error")
	}
}

func TestAnimatedScaleAttributes(t *testing.T) {
	attrs := &AnimatedScaleAttributes{
		Scale: 0.0,
	}

	_, err := json.Marshal(attrs)

	if err == nil {
		t.Fatal("must throw error")
	}
}

func TestSliverOpacityAttributes(t *testing.T) {
	attrs := &SliverOpacityAttributes{
		Opacity: 1.0,
	}

	_, err := json.Marshal(attrs)

	if err != nil {
		t.Fatal(err)
	}

	if attrs.Opacity != 1.0 {
		t.Fatal("opacity is not 1.0")
	}

	invalidAttrs := &SliverOpacityAttributes{
		Opacity: 1.1,
	}

	_, err = json.Marshal(invalidAttrs)

	if err == nil {
		t.Fatal("must throw error")
	}
}
