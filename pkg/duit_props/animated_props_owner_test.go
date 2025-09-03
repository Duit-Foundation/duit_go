package duit_props_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
)

func TestAnimatedPropertyOwner_Validate(t *testing.T) {
	
	attrs := &duit_props.AnimatedPropertyOwner{
		ParentBuilderId: "parentBuilderId",
		AffectedProperties: []string{"affectedProperties"},
	}

	err := attrs.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
