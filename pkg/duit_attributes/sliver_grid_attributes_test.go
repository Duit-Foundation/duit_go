package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_builder"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

// Tests for DefaultSliverGridAttributes
func TestDefaultSliverGridAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.DefaultSliverGridAttributes{
		Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCommon),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestDefaultSliverGridAttributes_Validate_MissingConstructor(t *testing.T) {
	attrs := &duit_attributes.DefaultSliverGridAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when constructor is not set")
	}
}

// Tests for Constructor Tristate[SliverGridConstructor] property serialization
func TestDefaultSliverGridAttributes_Constructor_JSON_SliverGridCommon(t *testing.T) {
	attrs := &duit_attributes.DefaultSliverGridAttributes{
		Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCommon),
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

func TestDefaultSliverGridAttributes_Constructor_JSON_SliverGridCount(t *testing.T) {
	attrs := &duit_attributes.DefaultSliverGridAttributes{
		Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCount),
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

func TestDefaultSliverGridAttributes_Constructor_JSON_SliverGridBuilder(t *testing.T) {
	attrs := &duit_attributes.DefaultSliverGridAttributes{
		Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridBuilder),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"constructor":2`) {
		t.Fatalf("expected JSON to contain '\"constructor\":2', got: %s", jsonStr)
	}
}

func TestDefaultSliverGridAttributes_Constructor_JSON_SliverGridExtent(t *testing.T) {
	attrs := &duit_attributes.DefaultSliverGridAttributes{
		Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridExtent),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"constructor":3`) {
		t.Fatalf("expected JSON to contain '\"constructor\":3', got: %s", jsonStr)
	}
}

func TestDefaultSliverGridAttributes_Constructor_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.DefaultSliverGridAttributes{
		Constructor: duit_utils.Nillable[duit_attributes.SliverGridConstructor](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"constructor":null`) {
		t.Fatalf("expected JSON to contain '\"constructor\":null', got: %s", jsonStr)
	}
}

// Tests for SliverGridCommonAttributes
func TestSliverGridCommonAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverGridCommonAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCommon),
		},
		SliverGridDelegateKey: "testKey",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverGridCommonAttributes_Validate_MissingDefaultSliverGridAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverGridCommonAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		SliverGridDelegateKey: "testKey",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when defaultSliverGridAttributes is not set")
	}
}

func TestSliverGridCommonAttributes_Validate_WrongConstructor(t *testing.T) {
	attrs := &duit_attributes.SliverGridCommonAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCount),
		},
		SliverGridDelegateKey: "testKey",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when constructor is not SliverGridCommon")
	}
}

func TestSliverGridCommonAttributes_Validate_MissingSliverGridDelegateKey(t *testing.T) {
	attrs := &duit_attributes.SliverGridCommonAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCommon),
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when SliverGridDelegateKey is empty")
	}
}

// Tests for SliverGridCountAttributes
func TestSliverGridCountAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverGridCountAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCount),
		},
		CrossAxisCount: 2,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverGridCountAttributes_Validate_MissingDefaultSliverGridAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverGridCountAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		CrossAxisCount:       2,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when defaultSliverGridAttributes is not set")
	}
}

func TestSliverGridCountAttributes_Validate_WrongConstructor(t *testing.T) {
	attrs := &duit_attributes.SliverGridCountAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCommon),
		},
		CrossAxisCount: 2,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when constructor is not SliverGridCount")
	}
}

func TestSliverGridCountAttributes_Validate_ZeroCrossAxisCount(t *testing.T) {
	attrs := &duit_attributes.SliverGridCountAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCount),
		},
		CrossAxisCount: 0,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when crossAxisCount is 0")
	}
}

func TestSliverGridCountAttributes_Validate_ValidCrossAxisCounts(t *testing.T) {
	testCases := []struct {
		name           string
		crossAxisCount int
	}{
		{"Single column", 1},
		{"Two columns", 2},
		{"Three columns", 3},
		{"Many columns", 10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.SliverGridCountAttributes{
				ValueReferenceHolder: nil,
				ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
				DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
					Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCount),
				},
				CrossAxisCount: tc.crossAxisCount,
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for crossAxisCount %d, got: %v", tc.crossAxisCount, err)
			}
		})
	}
}

// Tests for SliverGridBuilderAttributes
func TestSliverGridBuilderAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverGridBuilderAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		Builder:              &duit_builder.Builder{},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridBuilder),
		},
		SliverGridDelegateKey: "testKey",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverGridBuilderAttributes_Validate_MissingDefaultSliverGridAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverGridBuilderAttributes{
		ValueReferenceHolder:  &duit_attributes.ValueReferenceHolder{},
		ThemeConsumer:         &duit_attributes.ThemeConsumer{},
		Builder:               &duit_builder.Builder{},
		SliverGridDelegateKey: "testKey",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when defaultSliverGridAttributes is not set")
	}
}

func TestSliverGridBuilderAttributes_Validate_WrongConstructor(t *testing.T) {
	attrs := &duit_attributes.SliverGridBuilderAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		Builder:              &duit_builder.Builder{},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCount),
		},
		SliverGridDelegateKey: "testKey",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when constructor is not SliverGridBuilder")
	}
}

func TestSliverGridBuilderAttributes_Validate_MissingBuilder(t *testing.T) {
	attrs := &duit_attributes.SliverGridBuilderAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridBuilder),
		},
		SliverGridDelegateKey: "testKey",
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when builder is not set")
	}
}

func TestSliverGridBuilderAttributes_Validate_MissingSliverGridDelegateKey(t *testing.T) {
	attrs := &duit_attributes.SliverGridBuilderAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		Builder:              &duit_builder.Builder{},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridBuilder),
		},
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when sliverGridDelegateKey is empty")
	}
}

// Tests for SliverGridExtentAttributes
func TestSliverGridExtentAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverGridExtentAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridExtent),
		},
		MaxCrossAxisExtent: 150.0,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestSliverGridExtentAttributes_Validate_MissingDefaultSliverGridAttributes(t *testing.T) {
	attrs := &duit_attributes.SliverGridExtentAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		MaxCrossAxisExtent:   150.0,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when defaultSliverGridAttributes is not set")
	}
}

func TestSliverGridExtentAttributes_Validate_WrongConstructor(t *testing.T) {
	attrs := &duit_attributes.SliverGridExtentAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridCount),
		},
		MaxCrossAxisExtent: 150.0,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when constructor is not SliverGridExtent")
	}
}

func TestSliverGridExtentAttributes_Validate_ZeroMaxCrossAxisExtent(t *testing.T) {
	attrs := &duit_attributes.SliverGridExtentAttributes{
		ValueReferenceHolder: nil,
		ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
		DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
			Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridExtent),
		},
		MaxCrossAxisExtent: 0.0,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when MaxCrossAxisExtent is 0")
	}
}

func TestSliverGridExtentAttributes_Validate_ValidMaxCrossAxisExtents(t *testing.T) {
	testCases := []struct {
		name               string
		maxCrossAxisExtent float32
	}{
		{"Small extent", 50.0},
		{"Medium extent", 150.0},
		{"Large extent", 300.0},
		{"Very small extent", 1.0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			attrs := &duit_attributes.SliverGridExtentAttributes{
				ValueReferenceHolder: nil,
				ThemeConsumer:        &duit_attributes.ThemeConsumer{IgnoreTheme: true},
				DefaultSliverGridAttributes: &duit_attributes.DefaultSliverGridAttributes{
					Constructor: duit_utils.TristateFrom[duit_attributes.SliverGridConstructor](duit_attributes.SliverGridExtent),
				},
				MaxCrossAxisExtent: tc.maxCrossAxisExtent,
			}

			err := attrs.Validate()
			if err != nil {
				t.Fatalf("expected no error for maxCrossAxisExtent %f, got: %v", tc.maxCrossAxisExtent, err)
			}
		})
	}
}
