package duit_widget_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_widget"
)

func TestSliverVisibility_Attributes_Nil(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when attributes is nil, but did not panic")
		}
	}()

	duit_widget.SliverVisibility(nil, "test", false, nil, nil)
}
