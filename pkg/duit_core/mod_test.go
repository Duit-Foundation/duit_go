package duit_core

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestBuildWrapper(t *testing.T) {
	view := DuitView{}
	builder := view.Builder()

	builder.CreateRoot()

	j, err := builder.Build()

	if err != nil {
		t.Fatal(err)
	}

	val := map[string]interface{}{}

	err = json.Unmarshal(j, &val)

	if err != nil {
		t.Fatal(err)
	}

	if val["root"] == nil {
		t.Fatal("root is nil")
	}
}

func TestBuildWithoutWrapper(t *testing.T) {
	view := DuitView{}
	builder := view.Builder()

	builder.CreateRoot()

	j, err := builder.BuildUnwrapped()

	if err != nil {
		t.Fatal(err)
	}

	val := map[string]interface{}{}

	err = json.Unmarshal(j, &val)

	if err != nil {
		t.Fatal(err)
	}

	if val["root"] != nil {
		t.Fatal("root is not nil")
	}
}

func TestBuildMultiView(t *testing.T) {
	view := DuitView{}
	builder := view.Builder()

	_, err := builder.BuildMultiview()

	if err == nil {
		t.Fatal("expected error")
	}

	v1 := &DuitElementModel{
		ElementType: Text,
		Attributes:  nil,
		Controlled:  false,
		Id:          "v1",
	}

	v2 := &DuitElementModel{
		ElementType: Text,
		Attributes:  nil,
		Controlled:  false,
		Id:          "v2",
	}

	builder.AddWidgets(v1, v2)

	j, err := builder.BuildMultiview()

	if err != nil {
		t.Fatal(err)
	}

	val := map[string]interface{}{}

	err = json.Unmarshal(j, &val)

	if err != nil {
		t.Fatal(err)
	}

	if len(val["widgets"].([]interface{})) != 2 {
		t.Fatal("expected 2 widgets")
	}
}

func TestEmbeddedComponents(t *testing.T) {
	view := DuitView{}
	builder := view.Builder()

	v1 := &DuitElementModel{
		ElementType: Text,
		Attributes:  nil,
		Controlled:  false,
		Id:          "v1",
	}

	builder.AddComponents(&ComponentDescription{
		Tag:        "test",
		LayoutRoot: v1,
	})

	layout, err := builder.Build()

	if err != nil {
		t.Fatal(err)
	}

	val := map[string]interface{}{}

	err = json.Unmarshal(layout, &val)

	if err != nil {
		t.Fatal(err)
	}

	if val["embedded"] == nil {
		t.Fatal(errors.New("embedded component is nil"))
	}

	_, err = builder.BuildUnwrapped()

	if err == nil {
		t.Fatal(errors.New("expected error"))
	}
}
