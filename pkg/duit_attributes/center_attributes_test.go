package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
)

func TestCenterAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.CenterAttributes{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestCenterAttributes_MarshalJSON_ZeroFactors(t *testing.T) {
	attrs := &duit_attributes.CenterAttributes{
		WidgthFactor: 0.0,
		HeightFactor: 0.0,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	// Zero values should be omitted due to omitempty tag
	if strings.Contains(jsonStr, `"widthFactor"`) {
		t.Fatalf("expected JSON to not contain 'widthFactor' field when zero, got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"heightFactor"`) {
		t.Fatalf("expected JSON to not contain 'heightFactor' field when zero, got: %s", jsonStr)
	}
}

func TestCenterAttributes_MarshalJSON_OnlyWidthFactor(t *testing.T) {
	attrs := &duit_attributes.CenterAttributes{
		WidgthFactor: 0.75,
		HeightFactor: 0.0,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"widthFactor":0.75`) {
		t.Fatalf("expected JSON to contain '\"widthFactor\":0.75', got: %s", jsonStr)
	}
	if strings.Contains(jsonStr, `"heightFactor"`) {
		t.Fatalf("expected JSON to not contain 'heightFactor' field when zero, got: %s", jsonStr)
	}
}

func TestCenterAttributes_MarshalJSON_OnlyHeightFactor(t *testing.T) {
	attrs := &duit_attributes.CenterAttributes{
		WidgthFactor: 0.0,
		HeightFactor: 0.9,
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"widthFactor"`) {
		t.Fatalf("expected JSON to not contain 'widthFactor' field when zero, got: %s", jsonStr)
	}
	if !strings.Contains(jsonStr, `"heightFactor":0.9`) {
		t.Fatalf("expected JSON to contain '\"heightFactor\":0.9', got: %s", jsonStr)
	}
}