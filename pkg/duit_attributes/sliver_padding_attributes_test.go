package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestSliverPaddingAttributes_Validate_ValidAttributes(t *testing.T) {
	padding := &duit_props.EdgeInsetsV2{}
	padding = padding.LTRB(10.0, 20.0, 15.0, 25.0)

	attrs := &duit_attributes.SliverPaddingAttributes{
		Padding: padding,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverPaddingAttributes_Validate_WithoutPadding(t *testing.T) {
	attrs := &duit_attributes.SliverPaddingAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when padding is not set")
	}
}

func TestSliverPaddingAttributes_Validate_WithZeroPadding(t *testing.T) {
	padding := &duit_props.EdgeInsetsV2{}
	padding = padding.LTRB(0.0, 0.0, 0.0, 0.0)

	attrs := &duit_attributes.SliverPaddingAttributes{
		Padding: padding,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for zero padding, got:", err)
	}
}

func TestSliverPaddingAttributes_Validate_WithPartialPadding(t *testing.T) {
	padding := &duit_props.EdgeInsetsV2{}
	padding = padding.LTRB(10.0, 20.0, 0.0, 0.0)

	attrs := &duit_attributes.SliverPaddingAttributes{
		Padding: padding,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for partial padding, got:", err)
	}
}

func TestSliverPaddingAttributes_Validate_WithSymmetricPadding(t *testing.T) {
	padding := &duit_props.EdgeInsetsV2{}
	padding = padding.Symmentric(10.0, 20.0)

	attrs := &duit_attributes.SliverPaddingAttributes{
		Padding: padding,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for symmetric padding, got:", err)
	}
}

func TestSliverPaddingAttributes_Validate_WithAllPadding(t *testing.T) {
	padding := &duit_props.EdgeInsetsV2{}
	padding = padding.All(15.0)

	attrs := &duit_attributes.SliverPaddingAttributes{
		Padding: padding,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for all padding, got:", err)
	}
}
