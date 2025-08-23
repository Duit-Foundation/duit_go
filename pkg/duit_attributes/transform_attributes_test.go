package duit_attributes_test

import (
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_transform"
)

func TestTransformAttributes_Validate_ValidAttributes(t *testing.T) {
	data := duit_transform.TranslateTransfrom{}

	attrs := &duit_attributes.TransfromAttributes[duit_transform.TranslateTransfrom]{
		Type: "translate",
		Data: &data,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestTransformAttributes_Validate_MissingType(t *testing.T) {
	data := duit_transform.RotateTransfrom{}

	attrs := &duit_attributes.TransfromAttributes[duit_transform.RotateTransfrom]{
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
	attrs := &duit_attributes.TransfromAttributes[duit_transform.ScaleTransfrom]{
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
