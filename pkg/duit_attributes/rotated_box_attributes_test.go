package duit_attributes_test

import (
	"strings"
	"testing"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestRotatedBoxAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.RotatedBoxAttributes{
		QuarterTurns: duit_utils.IntValue(2),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestRotatedBoxAttributes_Validate_MissingQuarterTurns(t *testing.T) {
	attrs := &duit_attributes.RotatedBoxAttributes{
		QuarterTurns: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing quarterTurns")
	}

	if !strings.Contains(err.Error(), "quarterTurns property is required") {
		t.Fatalf("expected error message about required quarterTurns, got: %s", err.Error())
	}
}
