package duit_core

import (
	"encoding/json"
	"testing"
)

func TestBuildWrapper(t *testing.T) {
	view := DuitView{}
	builder := view.Builder();

	builder.CreateRoot();

	j, err := builder.Build()

	if err != nil {
		t.Fatal(err)
	}

	val := map[string]interface{}{}

	err = json.Unmarshal(j, &val)

	if err != nil {
		t.Fatal(err)
	}

	if (val["root"] == nil) {
		t.Fatal("root is nil")
	}
}

func TestBuildWithoutWrapper(t *testing.T) {
	view := DuitView{}
	builder := view.Builder();

	builder.CreateRoot();

	j, err := builder.BuildUnwrapped()

	if err != nil {
		t.Fatal(err)
	}

	val := map[string]interface{}{}

	err = json.Unmarshal(j, &val)

	if err != nil {
		t.Fatal(err)
	}

	if (val["root"] != nil) {
		t.Fatal("root is not nil")
	}
}