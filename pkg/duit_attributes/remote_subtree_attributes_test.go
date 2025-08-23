package duit_attributes_test

import (
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
)

func TestRemoteSubtreeAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.RemoteSubtreeAttributes{
		DownloadPath: "/path/to/resource",
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestRemoteSubtreeAttributes_Validate_WithoutDownloadPath(t *testing.T) {
	attrs := &duit_attributes.RemoteSubtreeAttributes{}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error when downloadPath is not set")
	}
}