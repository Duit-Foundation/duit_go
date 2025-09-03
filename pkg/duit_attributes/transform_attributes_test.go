package duit_attributes_test

import (
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestTransformAttributes_Validate_ValidAttributes(t *testing.T) {
	data := duit_props.TranslateTransform{}

	attrs := &duit_attributes.TransfromAttributes[duit_props.TranslateTransform]{
		Type: "translate",
		Data: &data,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestTransformAttributes_Validate_MissingType(t *testing.T) {
	data := duit_props.TranslateTransform{}

	attrs := &duit_attributes.TransfromAttributes[duit_props.TranslateTransform]{
		Type: "",
		Data: &data,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing type")
	}

	if !strings.Contains(err.Error(), "type property is required") {
		t.Fatalf("expected error message about required type, got: %s", err.Error())
	}
}

func TestTransformAttributes_Validate_MissingData(t *testing.T) {
	attrs := &duit_attributes.TransfromAttributes[duit_props.TranslateTransform]{
		Type: "rotate",
		Data: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing data")
	}

	if !strings.Contains(err.Error(), "data property is required") {
		t.Fatalf("expected error message about required data, got: %s", err.Error())
	}
}
