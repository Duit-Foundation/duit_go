package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_text_properties"
)

func TestRichTextAttributes_Validate_ValidAttributes(t *testing.T) {
	textSpan := &duit_text_properties.TextSpan[duit_color.ColorString]{
		Text: "Test text",
	}

	attrs := &duit_attributes.RichTextAttributes[duit_color.ColorString]{
		TextSpan: textSpan,
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestRichTextAttributes_Validate_MissingTextSpan(t *testing.T) {
	attrs := &duit_attributes.RichTextAttributes[duit_color.ColorString]{
		TextSpan: nil,
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for missing textSpan")
	}
}
