package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_builder"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// GridViewCommonAttributes tests

func TestGridViewCommonAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.GridViewCommonAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridCommon),
		},
		SliverGridDelegateKey: "test-key",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestGridViewCommonAttributes_Validate_MissingConstructor(t *testing.T) {
	attrs := &duit_attributes.GridViewCommonAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{},
		SliverGridDelegateKey:     "test-key",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing Constructor")
	}

	if !strings.Contains(err.Error(), "DefautlGridViewAttributes.Constructor is required") {
		t.Fatalf("expected error message about required Constructor, got: %s", err.Error())
	}
}

func TestGridViewCommonAttributes_Validate_NilDefautlGridViewAttributes(t *testing.T) {
	attrs := &duit_attributes.GridViewCommonAttributes[duit_props.EdgeInsetsAll]{
		SliverGridDelegateKey: "test-key",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil DefautlGridViewAttributes")
	}
}

func TestGridViewCommonAttributes_Validate_WrongConstructor(t *testing.T) {
	attrs := &duit_attributes.GridViewCommonAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridCount),
		},
		SliverGridDelegateKey: "test-key",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for wrong Constructor")
	}
}

func TestGridViewCommonAttributes_Validate_MissingSliverGridDelegateKey(t *testing.T) {
	attrs := &duit_attributes.GridViewCommonAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridCommon),
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing SliverGridDelegateKey")
	}
}

func TestGridViewCommonAttributes_Validate_NilConstructor(t *testing.T) {
	attrs := &duit_attributes.GridViewCommonAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		},
		SliverGridDelegateKey: "test-key",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil Constructor")
	}
}

// GridViewCountAttributes tests

func TestGridViewCountAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.GridViewCountAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridCount),
		},
		CrossAxisCount: 2,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestGridViewCountAttributes_Validate_MissingCrossAxisCount(t *testing.T) {
	attrs := &duit_attributes.GridViewCountAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridCount),
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing CrossAxisCount")
	}

	if !strings.Contains(err.Error(), "crossAxisCount is required and must be greater than 0") {
		t.Fatalf("expected error message about required CrossAxisCount, got: %s", err.Error())
	}
}

func TestGridViewCountAttributes_Validate_NilDefautlGridViewAttributes(t *testing.T) {
	attrs := &duit_attributes.GridViewCountAttributes[duit_props.EdgeInsetsAll]{
		CrossAxisCount: 2,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil DefautlGridViewAttributes")
	}
}

func TestGridViewCountAttributes_Validate_WrongConstructor(t *testing.T) {
	attrs := &duit_attributes.GridViewCountAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridCommon),
		},
		CrossAxisCount: 2,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for wrong Constructor")
	}
}

func TestGridViewCountAttributes_Validate_NilConstructor(t *testing.T) {
	attrs := &duit_attributes.GridViewCountAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		},
		CrossAxisCount: 2,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil Constructor")
	}
}

// GridViewBuilderAttributes tests

func TestGridViewBuilderAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.GridViewBuilderAttributes[duit_props.EdgeInsetsAll]{
		Builder: &duit_builder.Builder{},
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridBuilder),
		},
		SliverGridDelegateKey: "test-key",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestGridViewBuilderAttributes_Validate_MissingBuilder(t *testing.T) {
	attrs := &duit_attributes.GridViewBuilderAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridBuilder),
		},
		SliverGridDelegateKey: "test-key",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing Builder")
	}
}

func TestGridViewBuilderAttributes_Validate_NilDefautlGridViewAttributes(t *testing.T) {
	attrs := &duit_attributes.GridViewBuilderAttributes[duit_props.EdgeInsetsAll]{
		Builder:               &duit_builder.Builder{},
		SliverGridDelegateKey: "test-key",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil DefautlGridViewAttributes")
	}
}

func TestGridViewBuilderAttributes_Validate_MissingSliverGridDelegateKey(t *testing.T) {
	attrs := &duit_attributes.GridViewBuilderAttributes[duit_props.EdgeInsetsAll]{
		Builder: &duit_builder.Builder{},
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridBuilder),
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing SliverGridDelegateKey")
	}
}

func TestGridViewBuilderAttributes_Validate_WrongConstructor(t *testing.T) {
	attrs := &duit_attributes.GridViewBuilderAttributes[duit_props.EdgeInsetsAll]{
		Builder: &duit_builder.Builder{},
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridCommon),
		},
		SliverGridDelegateKey: "test-key",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for wrong Constructor")
	}
}

func TestGridViewBuilderAttributes_Validate_NilConstructor(t *testing.T) {
	attrs := &duit_attributes.GridViewBuilderAttributes[duit_props.EdgeInsetsAll]{
		Builder:                   &duit_builder.Builder{},
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		},
		SliverGridDelegateKey: "test-key",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil Constructor")
	}
}

// GridViewExtentAttributes tests

func TestGridViewExtentAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.GridViewExtentAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridExtent),
		},
		MaxCrossAxisExtent: 100.0,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestGridViewExtentAttributes_Validate_MissingMaxCrossAxisExtent(t *testing.T) {
	attrs := &duit_attributes.GridViewExtentAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridExtent),
		},
		MaxCrossAxisExtent: 0,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing MaxCrossAxisExtent")
	}
}

func TestGridViewExtentAttributes_Validate_NilDefautlGridViewAttributes(t *testing.T) {
	attrs := &duit_attributes.GridViewExtentAttributes[duit_props.EdgeInsetsAll]{
		MaxCrossAxisExtent: 100.0,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil DefautlGridViewAttributes")
	}
}

func TestGridViewExtentAttributes_Validate_WrongConstructor(t *testing.T) {
	attrs := &duit_attributes.GridViewExtentAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
			Constructor: duit_utils.TristateFrom[duit_attributes.GridConstructor](duit_attributes.GridCommon),
		},
		MaxCrossAxisExtent: 100.0,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for wrong Constructor")
	}
}

func TestGridViewExtentAttributes_Validate_NilConstructor(t *testing.T) {
	attrs := &duit_attributes.GridViewExtentAttributes[duit_props.EdgeInsetsAll]{
		DefautlGridViewAttributes: &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{},
		MaxCrossAxisExtent:        100.0,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for nil Constructor")
	}
}

// Tests for Tristate properties serialization in DefautlGridViewAttributes
// Following the rule: Always add tests for a struct property if the property is of type duit_utils.Tristate[T]

func TestDefautlGridViewAttributes_Constructor_JSON_GridCommon(t *testing.T) {
	constructor := duit_attributes.GridCommon
	attrs := &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		Constructor: &constructor,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"constructor":0`) {
		t.Fatalf("expected JSON to contain '\"constructor\":0', got: %s", jsonStr)
	}
}

func TestDefautlGridViewAttributes_Constructor_JSON_GridCount(t *testing.T) {
	constructor := duit_attributes.GridCount
	attrs := &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		Constructor: &constructor,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"constructor":1`) {
		t.Fatalf("expected JSON to contain '\"constructor\":1', got: %s", jsonStr)
	}
}

func TestDefautlGridViewAttributes_Reverse_JSON_True(t *testing.T) {
	constructor := duit_attributes.GridCommon
	attrs := &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		Constructor: &constructor,
		Reverse:     duit_utils.BoolValue(true),
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

func TestDefautlGridViewAttributes_Reverse_JSON_False(t *testing.T) {
	constructor := duit_attributes.GridCommon
	attrs := &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		Constructor: &constructor,
		Reverse:     duit_utils.BoolValue(false),
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

func TestDefautlGridViewAttributes_Primary_JSON_True(t *testing.T) {
	constructor := duit_attributes.GridCommon
	attrs := &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		Constructor: &constructor,
		Primary:     duit_utils.BoolValue(true),
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

func TestDefautlGridViewAttributes_ShrinkWrap_JSON_True(t *testing.T) {
	constructor := duit_attributes.GridCommon
	attrs := &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		Constructor: &constructor,
		ShrinkWrap:  duit_utils.BoolValue(true),
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

func TestDefautlGridViewAttributes_AddAutomaticKeepAlives_JSON_True(t *testing.T) {
	constructor := duit_attributes.GridCommon
	attrs := &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		Constructor:            &constructor,
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

func TestDefautlGridViewAttributes_AddRepaintBoundaries_JSON_True(t *testing.T) {
	constructor := duit_attributes.GridCommon
	attrs := &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		Constructor:          &constructor,
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

func TestDefautlGridViewAttributes_AddSemanticIndexes_JSON_True(t *testing.T) {
	constructor := duit_attributes.GridCommon
	attrs := &duit_attributes.DefautlGridViewAttributes[duit_props.EdgeInsetsAll]{
		Constructor:        &constructor,
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
