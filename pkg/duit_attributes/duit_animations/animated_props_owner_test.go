package duit_animations_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_animations"
)

func TestAnimatedPropertyOwner_Validate(t *testing.T) {
	
	attrs := &duit_animations.AnimatedPropertyOwner{
		ParentBuilderId: "parentBuilderId",
		AffectedProperties: []string{"affectedProperties"},
	}

	err := attrs.Validate()

	if err != nil {
		t.Fatal(err)
	}
}