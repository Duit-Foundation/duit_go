package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestContainerAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.ContainerAttributes{
		Width:  100.0,
		Height: 200.0,
		Color:  duit_props.NewColorString("#FF0000"),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}
