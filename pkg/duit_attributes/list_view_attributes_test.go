package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// ListViewAttributes tests

func TestListViewAttributes_MarshalJSON_ListViewCommonAttributes(t *testing.T) {
	commonAttrs := &duit_attributes.ListViewCommonAttributes{
		ListViewBaseAttributes: &duit_attributes.ListViewBaseAttributes{
			Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
		},
	}

	attrs := duit_attributes.NewListViewCommonAttributes(commonAttrs)
	
	jsonData, err := attrs.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"type":0`) {
		t.Fatalf("expected JSON to contain type, got: %s", jsonStr)
	}
}

func TestListViewAttributes_MarshalJSON_ListViewBuilderAttributes(t *testing.T) {
	builderAttrs := &duit_attributes.ListViewBuilderAttributes{
		ListViewBaseAttributes: &duit_attributes.ListViewBaseAttributes{
			Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListBuilder),
		},
		Builder: &duit_attributes.Builder{
			ChildObjects: []*duit_core.DuitElementModel{},
		},
	}

	attrs := duit_attributes.NewListViewBuilderAttributes(builderAttrs)
	
	jsonData, err := attrs.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"type":1`) {
		t.Fatalf("expected JSON to contain type, got: %s", jsonStr)
	}
}

func TestListViewAttributes_MarshalJSON_ListViewSeparatedAttributes(t *testing.T) {
	separatedAttrs := &duit_attributes.ListViewSeparatedAttributes{
		ListViewBaseAttributes: &duit_attributes.ListViewBaseAttributes{
			Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListSeparated),
		},
		Builder: &duit_attributes.Builder{
			ChildObjects: []*duit_core.DuitElementModel{},
		},
		Separator: &duit_core.DuitElementModel{},
	}

	attrs := duit_attributes.NewListViewSeparatedAttributes(separatedAttrs)
	
	jsonData, err := attrs.MarshalJSON()
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"type":2`) {
		t.Fatalf("expected JSON to contain type, got: %s", jsonStr)
	}
}

func TestNewListViewAttributes_WithValidTypes(t *testing.T) {
	testCases := []struct {
		name string
		data any
	}{
		{
			"ListViewCommonAttributes",
			&duit_attributes.ListViewCommonAttributes{
				ListViewBaseAttributes: &duit_attributes.ListViewBaseAttributes{
					Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
				},
			},
		},
		{
			"ListViewBuilderAttributes",
			&duit_attributes.ListViewBuilderAttributes{
				ListViewBaseAttributes: &duit_attributes.ListViewBaseAttributes{
					Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListBuilder),
				},
				Builder: &duit_attributes.Builder{},
			},
		},
		{
			"ListViewSeparatedAttributes",
			&duit_attributes.ListViewSeparatedAttributes{
				ListViewBaseAttributes: &duit_attributes.ListViewBaseAttributes{
					Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListSeparated),
				},
				Builder:   &duit_attributes.Builder{},
				Separator: &duit_core.DuitElementModel{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := duit_attributes.NewListViewAttributes(tc.data)
			if attrs == nil {
				t.Fatalf("expected non-nil ListViewAttributes for %s", tc.name)
			}
		})
	}
}

func TestNewListViewAttributes_WithInvalidType(t *testing.T) {
	attrs := duit_attributes.NewListViewAttributes("invalid")
	if attrs != nil {
		t.Fatal("expected nil ListViewAttributes for invalid type")
	}
}

// Tests for ListViewBaseAttributes
func TestListViewBaseAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestListViewBaseAttributes_Validate_MissingType(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing type")
	}
}

func TestListViewBaseAttributes_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:                    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
		Reverse:                 duit_utils.BoolValue(true),
		Primary:                 duit_utils.BoolValue(false),
		ShrinkWrap:              duit_utils.BoolValue(true),
		AddAutomaticKeepAlives:  duit_utils.BoolValue(false),
		AddRepaintBoundaries:    duit_utils.BoolValue(true),
		AddSemanticIndexes:      duit_utils.BoolValue(false),
		ScrollDirection:         duit_props.AxisVertical,
		ScrollPhysics:           duit_props.ScrollPhysicsAlwaysScrollable,
		CacheExtent:             100.0,
		Anchor:                  0.5,
		SemantickChildCount:     10,
		Padding:                 duit_utils.TristateFrom[duit_props.EdgeInsets](duit_props.NewEdgeInsetsAll(16.0)),
		ItemExtent:              50.0,
		ClipBehavior:            duit_props.ClipAntiAlias,
		RestorationId:           "test_restoration",
		DragStarnBehavior:       duit_props.DragStartBehaviorStart,
		KeyboardDismissBehavior: duit_props.ScrollViewKeyboardDismissBehaviorManual,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes with all properties, got:", err)
	}
}

// Tests for ListViewBuilderAttributes
func TestListViewBuilderAttributes_Validate_ValidAttributes(t *testing.T) {
	baseAttrs := &duit_attributes.ListViewBaseAttributes{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListBuilder),
	}

	builder := &duit_attributes.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	attrs := &duit_attributes.ListViewBuilderAttributes{
		ListViewBaseAttributes: baseAttrs,
		Builder:                builder,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestListViewBuilderAttributes_Validate_MissingBaseAttributes(t *testing.T) {
	builder := &duit_attributes.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	attrs := &duit_attributes.ListViewBuilderAttributes{
		ListViewBaseAttributes: nil,
		Builder:                builder,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing base attributes")
	}
}

func TestListViewBuilderAttributes_Validate_MissingBuilder(t *testing.T) {
	baseAttrs := &duit_attributes.ListViewBaseAttributes{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListBuilder),
	}

	attrs := &duit_attributes.ListViewBuilderAttributes{
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
	baseAttrs := &duit_attributes.ListViewBaseAttributes{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListSeparated),
	}

	builder := &duit_attributes.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	separator := &duit_core.DuitElementModel{}

	attrs := &duit_attributes.ListViewSeparatedAttributes{
		ListViewBaseAttributes: baseAttrs,
		Builder:                builder,
		Separator:              separator,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestListViewSeparatedAttributes_Validate_MissingBuilderAttributes(t *testing.T) {
	separator := &duit_core.DuitElementModel{}

	attrs := &duit_attributes.ListViewSeparatedAttributes{
		ListViewBaseAttributes: nil,
		Separator:              separator,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing builder attributes")
	}
}

func TestListViewSeparatedAttributes_Validate_MissingSeparator(t *testing.T) {
	baseAttrs := &duit_attributes.ListViewBaseAttributes{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListSeparated),
	}

	builder := &duit_attributes.Builder{
		ChildObjects: []*duit_core.DuitElementModel{},
	}

	attrs := &duit_attributes.ListViewSeparatedAttributes{
		ListViewBaseAttributes: baseAttrs,
		Builder:                builder,
		Separator:              nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing separator")
	}
}

// Tests for Tristate[bool] properties serialization
func TestListViewBaseAttributes_Reverse_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:    duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:       duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:       duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:       duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:                   duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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

func TestListViewBaseAttributes_AddAutomaticKeepAlives_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:                   duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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

func TestListViewBaseAttributes_AddAutomaticKeepAlives_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:                   duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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

func TestListViewBaseAttributes_AddRepaintBoundaries_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:                 duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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

func TestListViewBaseAttributes_AddRepaintBoundaries_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:                 duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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

func TestListViewBaseAttributes_AddRepaintBoundaries_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:                 duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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

func TestListViewBaseAttributes_AddSemanticIndexes_JSON_True(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:               duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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

func TestListViewBaseAttributes_AddSemanticIndexes_JSON_False(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:               duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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

func TestListViewBaseAttributes_AddSemanticIndexes_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type:               duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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

// Tests for Tristate[ListKind] property serialization
func TestListViewBaseAttributes_Type_JSON_Common(t *testing.T) {
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListCommon),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListBuilder),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
		Type: duit_utils.TristateFrom[duit_attributes.ListKind](duit_attributes.ListSeparated),
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
	attrs := &duit_attributes.ListViewBaseAttributes{
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
