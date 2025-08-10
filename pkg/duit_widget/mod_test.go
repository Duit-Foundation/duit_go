package duit_widget

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_attributes/duit_color"
	"github.com/Duit-Foundation/duit_go/v3/pkg/duit_core"
)

func TextAnimatedCrossFadeChildValidation(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	//no child, must panic
	AnimatedCrossFade(
		&duit_attributes.AnimatedCrossFadeAttributes{}, "", []*duit_core.DuitElementModel{},
	)

	AnimatedCrossFade(
		&duit_attributes.AnimatedCrossFadeAttributes{}, "", []*duit_core.DuitElementModel{
			Text(&duit_attributes.TextAttributes[duit_color.ColorString]{}, "text1", false),
			Text(&duit_attributes.TextAttributes[duit_color.ColorString]{}, "text2", false),
			Text(&duit_attributes.TextAttributes[duit_color.ColorString]{}, "text3", false),
		},
	)
}

func TestFragment(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	Fragment("", "")
}
