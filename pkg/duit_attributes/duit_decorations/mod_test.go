package duit_decoration

import (
	"encoding/json"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_color"
)

func TestRoundedRectangleBorderMarshaling(t *testing.T) {
	sb := &RoundedRectangleBorder[duit_color.ColorString]{}

	v, err := json.Marshal(sb)

	if err != nil {
		t.Fatal("Failed to marshal shape border")
	} else {
		data := map[string]interface{}{}

		json.Unmarshal(v, &data)

		if tp, ok := data["type"]; !ok {
			t.Fatal("type is missing")
		} else {
			if len(tp.(string)) == 0 {
				t.Log(tp)
				t.Fatal("type is not RoundedRectangleBorderType")
			}
		}
	}
}

func TestBeveledRectangleBorderMarshaling(t *testing.T) {
	sb := &BeveledRectangleBorder[duit_color.ColorString]{}

	v, err := json.Marshal(sb)

	if err != nil {
		t.Fatal("Failed to marshal shape border")
	} else {
		data := map[string]interface{}{}

		json.Unmarshal(v, &data)

		if tp, ok := data["type"]; !ok {
			t.Fatal("type is missing")
		} else {
			if len(tp.(string)) == 0 {
				t.Log(tp)
				t.Fatal("type is not RoundedRectangleBorderType")
			}
		}
	}
}

func TestContinuousRectangleBorderMarshaling(t *testing.T) {
	sb := &ContinuousRectangleBorder[duit_color.ColorString]{}

	v, err := json.Marshal(sb)

	if err != nil {
		t.Fatal("Failed to marshal shape border")
	} else {
		data := map[string]interface{}{}

		json.Unmarshal(v, &data)

		if tp, ok := data["type"]; !ok {
			t.Fatal("type is missing")
		} else {
			if len(tp.(string)) == 0 {
				t.Log(tp)
				t.Fatal("type is not RoundedRectangleBorderType")
			}
		}
	}
}

func TestCircleBorderMarshaling(t *testing.T) {
	sb := &CircleBorder[duit_color.ColorString]{}

	v, err := json.Marshal(sb)

	if err != nil {
		t.Fatal("Failed to marshal shape border")
	} else {
		data := map[string]interface{}{}

		json.Unmarshal(v, &data)

		if tp, ok := data["type"]; !ok {
			t.Fatal("type is missing")
		} else {
			if len(tp.(string)) == 0 {
				t.Log(tp)
				t.Fatal("type is not RoundedRectangleBorderType")
			}
		}
	}
}

func TestStadiumBorderMarshaling(t *testing.T) {
	sb := &StadiumBorder[duit_color.ColorString]{}

	v, err := json.Marshal(sb)

	if err != nil {
		t.Fatal("Failed to marshal shape border")
	} else {
		data := map[string]interface{}{}

		json.Unmarshal(v, &data)

		if tp, ok := data["type"]; !ok {
			t.Fatal("type is missing")
		} else {
			if len(tp.(string)) == 0 {
				t.Log(tp)
				t.Fatal("type is not RoundedRectangleBorderType")
			}
		}
	}
}
