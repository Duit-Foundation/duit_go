package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// Tests for SliverList
func TestSliverList_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type: duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverList_Validate_MissingType(t *testing.T) {
	attrs := &duit_attributes.SliverList{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when type is not set")
	}
}

func TestSliverList_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type:                   duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
		AddAutomaticKeepAlives: duit_utils.BoolValue(true),
		AddRepaintBoundaries:   duit_utils.BoolValue(false),
		AddSemanticIndexes:     duit_utils.BoolValue(true),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for SliverListBuilderAttributes
func TestSliverListBuilderAttributes_Validate_ValidAttributes(t *testing.T) {
	sliverList := &duit_attributes.SliverList{
		Type: duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListBuilder),
	}

	builder := &duit_attributes.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	attrs := &duit_attributes.SliverListBuilderAttributes{
		SliverList: sliverList,
		Builder:    builder,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverListBuilderAttributes_Validate_MissingSliverList(t *testing.T) {
	builder := &duit_attributes.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	attrs := &duit_attributes.SliverListBuilderAttributes{
		SliverList: nil,
		Builder:    builder,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when sliverList is not set")
	}
}

func TestSliverListBuilderAttributes_Validate_WrongType(t *testing.T) {
	sliverList := &duit_attributes.SliverList{
		Type: duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
	}

	builder := &duit_attributes.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	attrs := &duit_attributes.SliverListBuilderAttributes{
		SliverList: sliverList,
		Builder:    builder,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when type is not SliverListBuilder")
	}
}

func TestSliverListBuilderAttributes_Validate_MissingBuilder(t *testing.T) {
	sliverList := &duit_attributes.SliverList{
		Type: duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListBuilder),
	}

	attrs := &duit_attributes.SliverListBuilderAttributes{
		SliverList: sliverList,
		Builder:    nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when builder is not set")
	}
}

// Tests for SliverListSeparatedAttributes
func TestSliverListSeparatedAttributes_Validate_ValidAttributes(t *testing.T) {
	sliverList := &duit_attributes.SliverList{
		Type: duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListBuilder),
	}

	builder := &duit_attributes.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	builderAttrs := &duit_attributes.SliverListBuilderAttributes{
		SliverList: sliverList,
		Builder:    builder,
	}

	separator := &duit_core.DuitElementModel{}

	attrs := &duit_attributes.SliverListSeparatedAttributes{
		SliverListBuilderAttributes: builderAttrs,
		Separator:                   separator,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverListSeparatedAttributes_Validate_MissingBuilderAttributes(t *testing.T) {
	separator := &duit_core.DuitElementModel{}

	attrs := &duit_attributes.SliverListSeparatedAttributes{
		SliverListBuilderAttributes: nil,
		Separator:                   separator,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when SliverListBuilderAttributes is not set")
	}
}

func TestSliverListSeparatedAttributes_Validate_WrongType(t *testing.T) {
	sliverList := &duit_attributes.SliverList{
		Type: duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
	}

	builder := &duit_attributes.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	builderAttrs := &duit_attributes.SliverListBuilderAttributes{
		SliverList: sliverList,
		Builder:    builder,
	}

	separator := &duit_core.DuitElementModel{}

	attrs := &duit_attributes.SliverListSeparatedAttributes{
		SliverListBuilderAttributes: builderAttrs,
		Separator:                   separator,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when type is not SliverListBuilder")
	}
}

func TestSliverListSeparatedAttributes_Validate_MissingSeparator(t *testing.T) {
	sliverList := &duit_attributes.SliverList{
		Type: duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListBuilder),
	}

	builder := &duit_attributes.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	builderAttrs := &duit_attributes.SliverListBuilderAttributes{
		SliverList: sliverList,
		Builder:    builder,
	}

	attrs := &duit_attributes.SliverListSeparatedAttributes{
		SliverListBuilderAttributes: builderAttrs,
		Separator:                   nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when separator is not set")
	}
}

// Tests for Tristate[SliverListKind] property serialization
func TestSliverList_Type_JSON_SliverListCommon(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type: duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"type":0`) {
		t.Fatalf("expected JSON to contain '\"type\":0', got: %s", jsonStr)
	}
}

func TestSliverList_Type_JSON_SliverListBuilder(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type: duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListBuilder),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"type":1`) {
		t.Fatalf("expected JSON to contain '\"type\":1', got: %s", jsonStr)
	}
}

func TestSliverList_Type_JSON_SliverListSeparated(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type: duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListSeparated),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"type":2`) {
		t.Fatalf("expected JSON to contain '\"type\":2', got: %s", jsonStr)
	}
}

func TestSliverList_Type_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type: duit_utils.Nillable[duit_attributes.SliverListKind](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"type":null`) {
		t.Fatalf("expected JSON to contain '\"type\":null', got: %s", jsonStr)
	}
}

// Tests for Tristate[bool] properties serialization
func TestSliverList_AddAutomaticKeepAlives_JSON_True(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type:                   duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
		AddAutomaticKeepAlives: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"addAutomaticKeepAlives":true`) {
		t.Fatalf("expected JSON to contain '\"addAutomaticKeepAlives\":true', got: %s", jsonStr)
	}
}

func TestSliverList_AddAutomaticKeepAlives_JSON_False(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type:                   duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
		AddAutomaticKeepAlives: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"addAutomaticKeepAlives":false`) {
		t.Fatalf("expected JSON to contain '\"addAutomaticKeepAlives\":false', got: %s", jsonStr)
	}
}

func TestSliverList_AddAutomaticKeepAlives_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type:                   duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
		AddAutomaticKeepAlives: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"addAutomaticKeepAlives"`) {
		t.Fatalf("expected JSON to not contain 'addAutomaticKeepAlives' field when nil, got: %s", jsonStr)
	}
}

func TestSliverList_AddRepaintBoundaries_JSON_True(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type:                 duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
		AddRepaintBoundaries: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"addRepaintBoundaries":true`) {
		t.Fatalf("expected JSON to contain '\"addRepaintBoundaries\":true', got: %s", jsonStr)
	}
}

func TestSliverList_AddRepaintBoundaries_JSON_False(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type:                 duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
		AddRepaintBoundaries: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"addRepaintBoundaries":false`) {
		t.Fatalf("expected JSON to contain '\"addRepaintBoundaries\":false', got: %s", jsonStr)
	}
}

func TestSliverList_AddRepaintBoundaries_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type:                 duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
		AddRepaintBoundaries: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"addRepaintBoundaries"`) {
		t.Fatalf("expected JSON to not contain 'addRepaintBoundaries' field when nil, got: %s", jsonStr)
	}
}

func TestSliverList_AddSemanticIndexes_JSON_True(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type:               duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
		AddSemanticIndexes: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"addSemanticIndexes":true`) {
		t.Fatalf("expected JSON to contain '\"addSemanticIndexes\":true', got: %s", jsonStr)
	}
}

func TestSliverList_AddSemanticIndexes_JSON_False(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type:               duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
		AddSemanticIndexes: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"addSemanticIndexes":false`) {
		t.Fatalf("expected JSON to contain '\"addSemanticIndexes\":false', got: %s", jsonStr)
	}
}

func TestSliverList_AddSemanticIndexes_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.SliverList{
		Type:               duit_utils.TristateFrom[duit_attributes.SliverListKind](duit_attributes.SliverListCommon),
		AddSemanticIndexes: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"addSemanticIndexes"`) {
		t.Fatalf("expected JSON to not contain 'addSemanticIndexes' field when nil, got: %s", jsonStr)
	}
}
