package duit_attributes

import (
	"encoding/json"
	"testing"

	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_painting"
)

func TestBackdropFilter(t *testing.T) {
	blur := duit_painting.NewBlurImageFilter(1, 2, "")

	attr := &BackdropFilterAttributes[duit_painting.BlurImageFilter]{
		Filter: blur,
		BlendMode: duit_painting.Clear,
	}

	_, err := json.Marshal(attr)

	if err != nil {
		t.Fatal(err)
	}
}