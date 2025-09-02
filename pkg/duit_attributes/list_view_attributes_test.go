package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_builder"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_edge_insets"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_gestures"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// Tests for ListViewBaseAttributes
func TestListViewBaseAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestListViewBaseAttributes_Validate_MissingType(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing type")
	}
}

func TestListViewBaseAttributes_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:                    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		Reverse:                 duit_utils.BoolValue(true),
		Primary:                 duit_utils.BoolValue(false),
		ShrinkWrap:              duit_utils.BoolValue(true),
		AddAutomaticKeepAlives:  duit_utils.BoolValue(false),
		AddRepaintBoundaries:    duit_utils.BoolValue(true),
		AddSemanticIndexes:      duit_utils.BoolValue(false),
		ScrollDirection:         duit_flex.Vertical,
		ScrollPhysics:           duit_gestures.AlwaysScrollableScrollPhysics,
		CacheExtent:             100.0,
		Anchor:                  0.5,
		SemantickChildCount:     10,
		Padding:                 duit_utils.TristateFrom[duit_edge_insets.EdgeInsetsAll](duit_edge_insets.EdgeInsetsAll(16.0)),
		ItemExtent:              50.0,
		ClipBehavior:            duit_props.ClipAntiAlias,
		RestorationId:           "test_restoration",
		DragStarnBehavior:       duit_gestures.Start,
		KeyboardDismissBehavior: duit_gestures.Manual,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for ListViewBuilderAttributes
func TestListViewBuilderAttributes_Validate_ValidAttributes(t *testing.T) {
	baseAttrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Builder),
	}
	
	builder := &duit_builder.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	attrs := &duit_attributes.ListViewBuilderAttributes[duit_edge_insets.EdgeInsetsAll]{
		ListViewBaseAttributes: baseAttrs,
		Builder:                builder,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestListViewBuilderAttributes_Validate_MissingBaseAttributes(t *testing.T) {
	builder := &duit_builder.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	attrs := &duit_attributes.ListViewBuilderAttributes[duit_edge_insets.EdgeInsetsAll]{
		ListViewBaseAttributes: nil,
		Builder:                builder,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing base attributes")
	}
}

func TestListViewBuilderAttributes_Validate_MissingBuilder(t *testing.T) {
	baseAttrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Builder),
	}

	attrs := &duit_attributes.ListViewBuilderAttributes[duit_edge_insets.EdgeInsetsAll]{
		ListViewBaseAttributes: baseAttrs,
		Builder:                nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing builder")
	}
}

// Tests for ListViewSeparatedAttributes
func TestListViewSeparatedAttributes_Validate_ValidAttributes(t *testing.T) {
	baseAttrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Separated),
	}
	
	builder := &duit_builder.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	builderAttrs := &duit_attributes.ListViewBuilderAttributes[duit_edge_insets.EdgeInsetsAll]{
		ListViewBaseAttributes: baseAttrs,
		Builder:                builder,
	}

	separator := &duit_core.DuitElementModel{}

	attrs := &duit_attributes.ListViewSeparatedAttributes[duit_edge_insets.EdgeInsetsAll]{
		ListViewBuilderAttributes: builderAttrs,
		Separator:                 separator,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestListViewSeparatedAttributes_Validate_MissingBuilderAttributes(t *testing.T) {
	separator := &duit_core.DuitElementModel{}

	attrs := &duit_attributes.ListViewSeparatedAttributes[duit_edge_insets.EdgeInsetsAll]{
		ListViewBuilderAttributes: nil,
		Separator:                 separator,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing builder attributes")
	}
}

func TestListViewSeparatedAttributes_Validate_MissingSeparator(t *testing.T) {
	baseAttrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Separated),
	}
	
	builder := &duit_builder.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	builderAttrs := &duit_attributes.ListViewBuilderAttributes[duit_edge_insets.EdgeInsetsAll]{
		ListViewBaseAttributes: baseAttrs,
		Builder:                builder,
	}

	attrs := &duit_attributes.ListViewSeparatedAttributes[duit_edge_insets.EdgeInsetsAll]{
		ListViewBuilderAttributes: builderAttrs,
		Separator:                 nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing separator")
	}
}

// Tests for Tristate[bool] properties serialization
func TestListViewBaseAttributes_Reverse_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		Reverse: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"reverse":true`) {
		t.Fatalf("expected JSON to contain '\"reverse\":true', got: %s", jsonStr)
	}
}

func TestListViewBaseAttributes_Reverse_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		Reverse: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"reverse":false`) {
		t.Fatalf("expected JSON to contain '\"reverse\":false', got: %s", jsonStr)
	}
}

func TestListViewBaseAttributes_Reverse_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		Reverse: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"reverse"`) {
		t.Fatalf("expected JSON to not contain 'reverse' field when nil, got: %s", jsonStr)
	}
}

func TestListViewBaseAttributes_Primary_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		Primary: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"primary":true`) {
		t.Fatalf("expected JSON to contain '\"primary\":true', got: %s", jsonStr)
	}
}

func TestListViewBaseAttributes_Primary_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		Primary: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"primary":false`) {
		t.Fatalf("expected JSON to contain '\"primary\":false', got: %s", jsonStr)
	}
}

func TestListViewBaseAttributes_Primary_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		Primary: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"primary"`) {
		t.Fatalf("expected JSON to not contain 'primary' field when nil, got: %s", jsonStr)
	}
}

func TestListViewBaseAttributes_ShrinkWrap_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:       duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		ShrinkWrap: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"shrinkWrap":true`) {
		t.Fatalf("expected JSON to contain '\"shrinkWrap\":true', got: %s", jsonStr)
	}
}

func TestListViewBaseAttributes_ShrinkWrap_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:       duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		ShrinkWrap: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"shrinkWrap":false`) {
		t.Fatalf("expected JSON to contain '\"shrinkWrap\":false', got: %s", jsonStr)
	}
}

func TestListViewBaseAttributes_ShrinkWrap_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:       duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		ShrinkWrap: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"shrinkWrap"`) {
		t.Fatalf("expected JSON to not contain 'shrinkWrap' field when nil, got: %s", jsonStr)
	}
}

func TestListViewBaseAttributes_AddAutomaticKeepAlives_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:                    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		AddAutomaticKeepAlives:  duit_utils.BoolValue(true),
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

func TestListViewBaseAttributes_AddAutomaticKeepAlives_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:                    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		AddAutomaticKeepAlives:  duit_utils.BoolValue(false),
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

func TestListViewBaseAttributes_AddAutomaticKeepAlives_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:                    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		AddAutomaticKeepAlives:  duit_utils.Nillable[bool](),
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

func TestListViewBaseAttributes_AddRepaintBoundaries_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:                  duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		AddRepaintBoundaries:  duit_utils.BoolValue(true),
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

func TestListViewBaseAttributes_AddRepaintBoundaries_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:                  duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		AddRepaintBoundaries:  duit_utils.BoolValue(false),
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

func TestListViewBaseAttributes_AddRepaintBoundaries_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:                  duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		AddRepaintBoundaries:  duit_utils.Nillable[bool](),
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

func TestListViewBaseAttributes_AddSemanticIndexes_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:                  duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		AddSemanticIndexes:    duit_utils.BoolValue(true),
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

func TestListViewBaseAttributes_AddSemanticIndexes_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:                  duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		AddSemanticIndexes:    duit_utils.BoolValue(false),
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

func TestListViewBaseAttributes_AddSemanticIndexes_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type:                  duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
		AddSemanticIndexes:    duit_utils.Nillable[bool](),
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

// Tests for Tristate[ListKind] property serialization
func TestListViewBaseAttributes_Type_JSON_Common(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Common),
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

func TestListViewBaseAttributes_Type_JSON_Builder(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Builder),
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

func TestListViewBaseAttributes_Type_JSON_Separated(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.Separated),
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

func TestListViewBaseAttributes_Type_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes[duit_edge_insets.EdgeInsetsAll]{
		Type: nil,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"type"`) {
		t.Fatalf("expected JSON to not contain 'type' field when nil, got: %s", jsonStr)
	}
}
