package duit_painting

import (
	"encoding/json"
	"testing"
)

func TestFiltersConstructors(t *testing.T) {
	dilate := NewDilateImageFilter(1, 2)
	blur := NewBlurImageFilter(1, 2, "")
	erode := NewErodeImageFilter(1, 2)
	matrix := NewMatrixImageFilter([16]float64{}, "")
	compose := NewComposeImageFilter(dilate, blur)

	arr := []any{
		dilate,
		blur,
		erode,
		matrix,
		compose,
	}

	for _, v := range arr {
		_, err := json.Marshal(v)

		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestComposeFilterValidation(t *testing.T) {
	defer func() { _ = recover() }()

	//positive case
	filter := NewComposeImageFilter(NewBlurImageFilter(1, 2, ""), nil)

	_, err := json.Marshal(filter)

	if err != nil {
		t.Fatal(err)
	}

	//negative case
	filter2 := NewComposeImageFilter(BlurImageFilter{}, nil)

	_, e := json.Marshal(filter2)

	if e == nil {
		t.Fatal(err)
	}
}
