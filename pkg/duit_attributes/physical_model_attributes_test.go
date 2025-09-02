package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	duit_decoration "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestPhysicalModelAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.PhysicalModelAttributes[duit_color.ColorString]{
		Elevation: duit_utils.Float32Value(8.0),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestPhysicalModelAttributes_Validate_WithoutElevation(t *testing.T) {
	attrs := &duit_attributes.PhysicalModelAttributes[duit_color.ColorString]{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when elevation is not set")
	}
}

func TestPhysicalModelAttributes_Validate_WithAllProperties(t *testing.T) {
	attrs := &duit_attributes.PhysicalModelAttributes[duit_color.ColorString]{
		Elevation:    duit_utils.Float32Value(16.0),
		Color:        duit_color.ColorString("#FF0000"),
		ShadowColor:  duit_color.ColorString("#000000"),
		ClipBehavior: duit_props.ClipAntiAlias,
		BorderRadius: &duit_decoration.BorderRadius{},
		Shape:        duit_decoration.Rectangle,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for attributes with all properties, got:", err)
	}
}

// Tests for Tristate[float32] property serialization
func TestPhysicalModelAttributes_Tristate_JSON_ValidValue(t *testing.T) {
	attrs := &duit_attributes.PhysicalModelAttributes[duit_color.ColorString]{
		Elevation: duit_utils.Float32Value(12.0),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"elevation":12`) {
		t.Fatalf("expected JSON to contain '\"elevation\":12', got: %s", jsonStr)
	}
}

func TestPhysicalModelAttributes_Tristate_JSON_ZeroValue(t *testing.T) {
	attrs := &duit_attributes.PhysicalModelAttributes[duit_color.ColorString]{
		Elevation: duit_utils.Float32Value(0.0),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"elevation":0`) {
		t.Fatalf("expected JSON to contain '\"elevation\":0', got: %s", jsonStr)
	}
}

func TestPhysicalModelAttributes_Tristate_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.PhysicalModelAttributes[duit_color.ColorString]{
		Elevation: duit_utils.Nillable[float32](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	// When Tristate is nil/unset, it should be omitted from JSON due to omitempty tag
	if strings.Contains(jsonStr, `"elevation"`) {
		t.Fatalf("expected JSON to not contain 'elevation' field when nil, got: %s", jsonStr)
	}
}