package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_alignment"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_clip"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_flex"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
)

func TestStackAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.StackAttributes{
		TextDirection: duit_text_properties.Ltr,
		ClipBehavior:  duit_clip.AntiAlias,
		Alignment:     duit_alignment.Center,
		Fit:           duit_flex.Loose,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}
