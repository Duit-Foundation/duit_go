package duit_attributes_test

import (
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestSliverFillViewportAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverFillViewportAttributes{
		ViewportFraction:       1.0,
		PadEnds:                duit_utils.BoolValue(true),
		IsBuilderDelegate:      duit_utils.BoolValue(false),
		AddAutomaticKeepAlives: duit_utils.BoolValue(true),
		AddRepaintBoundaries:   duit_utils.BoolValue(false),
		AddSemanticIndexes:     duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverFillViewportAttributes_Validate_NegativeViewportFraction(t *testing.T) {
	attrs := &duit_attributes.SliverFillViewportAttributes{
		ViewportFraction: -0.5,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for negative viewportFraction")
	}
}

func TestSliverFillViewportAttributes_Validate_BuilderDelegateWithChildren(t *testing.T) {
	childElement := &duit_core.DuitElementModel{}

	attrs := &duit_attributes.SliverFillViewportAttributes{
		ViewportFraction:  1.0,
		IsBuilderDelegate: duit_utils.BoolValue(true),
		ChildObjects:      []*duit_core.DuitElementModel{childElement},
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error when isBuilderDelegate is true and childObjects is non-empty, got:", err)
	}
}

func TestSliverFillViewportAttributes_Validate_BuilderDelegateWithoutChildren(t *testing.T) {
	attrs := &duit_attributes.SliverFillViewportAttributes{
		ViewportFraction:  1.0,
		IsBuilderDelegate: duit_utils.BoolValue(true),
		ChildObjects:      []*duit_core.DuitElementModel{},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when isBuilderDelegate is true but childObjects is empty")
	}
}

func TestSliverFillViewportAttributes_Validate_NonBuilderDelegateWithChildren(t *testing.T) {
	childElement := &duit_core.DuitElementModel{}

	attrs := &duit_attributes.SliverFillViewportAttributes{
		ViewportFraction:  1.0,
		IsBuilderDelegate: duit_utils.BoolValue(false),
		ChildObjects:      []*duit_core.DuitElementModel{childElement},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when isBuilderDelegate is false but childObjects is non-empty")
	}
}

func TestSliverFillViewportAttributes_Validate_NilBuilderDelegateWithChildren(t *testing.T) {
	childElement := &duit_core.DuitElementModel{}

	attrs := &duit_attributes.SliverFillViewportAttributes{
		ViewportFraction: 1.0,
		ChildObjects:     []*duit_core.DuitElementModel{childElement},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when isBuilderDelegate is nil but childObjects is non-empty")
	}
}
