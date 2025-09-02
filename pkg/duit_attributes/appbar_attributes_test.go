package duit_attributes_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	duit_decorations "github.com/Duit-Foundation/duit_go/v4/pkg/duit_attributes/duit_decorations"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

func TestAppBarAttributes_Validate_ValidAttributes(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid attributes, got:", err)
	}
}

func TestAppBarAttributes_Validate_ValidOpacityRanges(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		BottomOpacity:  duit_utils.Float32Value(0.5),
		ToolbarOpacity: duit_utils.Float32Value(0.8),
	}

	err := attrs.Validate()
	if err != nil {
		t.Fatal("expected no error for valid opacity ranges, got:", err)
	}
}

func TestAppBarAttributes_Validate_InvalidBottomOpacityLow(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		BottomOpacity: duit_utils.Float32Value(-0.1),
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for BottomOpacity below 0")
	}
}

func TestAppBarAttributes_Validate_InvalidBottomOpacityHigh(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		BottomOpacity: duit_utils.Float32Value(1.1),
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for BottomOpacity above 1")
	}
}

func TestAppBarAttributes_Validate_InvalidToolbarOpacityLow(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ToolbarOpacity: duit_utils.Float32Value(-0.1),
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for ToolbarOpacity below 0")
	}
}

func TestAppBarAttributes_Validate_InvalidToolbarOpacityHigh(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ToolbarOpacity: duit_utils.Float32Value(1.1),
	}

	err := attrs.Validate()
	if err == nil {
		t.Fatal("expected error for ToolbarOpacity above 1")
	}
}

// Tests for BottomOpacity Tristate[float32] property serialization
func TestAppBarAttributes_BottomOpacity_JSON_Value(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		BottomOpacity: duit_utils.Float32Value(0.5),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"bottomOpacity":0.5`) {
		t.Fatalf("expected JSON to contain '\"bottomOpacity\":0.5', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_BottomOpacity_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		BottomOpacity: duit_utils.Nillable[float32](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"bottomOpacity"`) {
		t.Fatalf("expected JSON to not contain 'bottomOpacity' field when nil, got: %s", jsonStr)
	}
}

func TestAppBarAttributes_BottomOpacity_Unmarshal_Value(t *testing.T) {
	jsonStr := `{"bottomOpacity": 0.7}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.BottomOpacity == nil {
		t.Fatal("expected bottomOpacity to have a value")
	}

	if *attrs.BottomOpacity != 0.7 {
		t.Fatalf("expected bottomOpacity value to be 0.7, got: %f", *attrs.BottomOpacity)
	}
}

func TestAppBarAttributes_BottomOpacity_Unmarshal_Omitted(t *testing.T) {
	jsonStr := `{}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.BottomOpacity != nil {
		t.Fatal("expected bottomOpacity to not have a value when omitted")
	}
}

// Tests for ToolbarOpacity Tristate[float32] property serialization
func TestAppBarAttributes_ToolbarOpacity_JSON_Value(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ToolbarOpacity: duit_utils.Float32Value(0.3),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"toolbarOpacity":0.3`) {
		t.Fatalf("expected JSON to contain '\"toolbarOpacity\":0.3', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_ToolbarOpacity_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ToolbarOpacity: duit_utils.Nillable[float32](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"toolbarOpacity"`) {
		t.Fatalf("expected JSON to not contain 'toolbarOpacity' field when nil, got: %s", jsonStr)
	}
}

func TestAppBarAttributes_ToolbarOpacity_Unmarshal_Value(t *testing.T) {
	jsonStr := `{"toolbarOpacity": 0.9}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.ToolbarOpacity == nil {
		t.Fatal("expected toolbarOpacity to have a value")
	}

	if *attrs.ToolbarOpacity != 0.9 {
		t.Fatalf("expected toolbarOpacity value to be 0.9, got: %f", *attrs.ToolbarOpacity)
	}
}

func TestAppBarAttributes_ToolbarOpacity_Unmarshal_Omitted(t *testing.T) {
	jsonStr := `{}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.ToolbarOpacity != nil {
		t.Fatal("expected toolbarOpacity to not have a value when omitted")
	}
}

// Tests for ForceMaterialTransparency Tristate[bool] property serialization
func TestAppBarAttributes_ForceMaterialTransparency_JSON_True(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ForceMaterialTransparency: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"forceMaterialTransparency":true`) {
		t.Fatalf("expected JSON to contain '\"forceMaterialTransparency\":true', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_ForceMaterialTransparency_JSON_False(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ForceMaterialTransparency: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"forceMaterialTransparency":false`) {
		t.Fatalf("expected JSON to contain '\"forceMaterialTransparency\":false', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_ForceMaterialTransparency_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ForceMaterialTransparency: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"forceMaterialTransparency"`) {
		t.Fatalf("expected JSON to not contain 'forceMaterialTransparency' field when nil, got: %s", jsonStr)
	}
}

func TestAppBarAttributes_ForceMaterialTransparency_Unmarshal_True(t *testing.T) {
	jsonStr := `{"forceMaterialTransparency": true}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.ForceMaterialTransparency == nil {
		t.Fatal("expected forceMaterialTransparency to have a value")
	}

	if !*attrs.ForceMaterialTransparency {
		t.Fatal("expected forceMaterialTransparency value to be true")
	}
}

func TestAppBarAttributes_ForceMaterialTransparency_Unmarshal_False(t *testing.T) {
	jsonStr := `{"forceMaterialTransparency": false}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.ForceMaterialTransparency == nil {
		t.Fatal("expected forceMaterialTransparency to have a value")
	}

	if *attrs.ForceMaterialTransparency {
		t.Fatal("expected forceMaterialTransparency value to be false")
	}
}

func TestAppBarAttributes_ForceMaterialTransparency_Unmarshal_Omitted(t *testing.T) {
	jsonStr := `{}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.ForceMaterialTransparency != nil {
		t.Fatal("expected forceMaterialTransparency to not have a value when omitted")
	}
}

// Tests for CenterTitle Tristate[bool] property serialization
func TestAppBarAttributes_CenterTitle_JSON_True(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		CenterTitle: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"centerTitle":true`) {
		t.Fatalf("expected JSON to contain '\"centerTitle\":true', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_CenterTitle_JSON_False(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		CenterTitle: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"centerTitle":false`) {
		t.Fatalf("expected JSON to contain '\"centerTitle\":false', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_CenterTitle_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		CenterTitle: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"centerTitle"`) {
		t.Fatalf("expected JSON to not contain 'centerTitle' field when nil, got: %s", jsonStr)
	}
}

func TestAppBarAttributes_CenterTitle_Unmarshal_True(t *testing.T) {
	jsonStr := `{"centerTitle": true}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.CenterTitle == nil {
		t.Fatal("expected centerTitle to have a value")
	}

	if !*attrs.CenterTitle {
		t.Fatal("expected centerTitle value to be true")
	}
}

func TestAppBarAttributes_CenterTitle_Unmarshal_False(t *testing.T) {
	jsonStr := `{"centerTitle": false}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.CenterTitle == nil {
		t.Fatal("expected centerTitle to have a value")
	}

	if *attrs.CenterTitle {
		t.Fatal("expected centerTitle value to be false")
	}
}

func TestAppBarAttributes_CenterTitle_Unmarshal_Omitted(t *testing.T) {
	jsonStr := `{}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.CenterTitle != nil {
		t.Fatal("expected centerTitle to not have a value when omitted")
	}
}

// Tests for AutomaticallyImplyLeading Tristate[bool] property serialization
func TestAppBarAttributes_AutomaticallyImplyLeading_JSON_True(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		AutomaticallyImplyLeading: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"automaticallyImplyLeading":true`) {
		t.Fatalf("expected JSON to contain '\"automaticallyImplyLeading\":true', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_AutomaticallyImplyLeading_JSON_False(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		AutomaticallyImplyLeading: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"automaticallyImplyLeading":false`) {
		t.Fatalf("expected JSON to contain '\"automaticallyImplyLeading\":false', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_AutomaticallyImplyLeading_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		AutomaticallyImplyLeading: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"automaticallyImplyLeading"`) {
		t.Fatalf("expected JSON to not contain 'automaticallyImplyLeading' field when nil, got: %s", jsonStr)
	}
}

func TestAppBarAttributes_AutomaticallyImplyLeading_Unmarshal_True(t *testing.T) {
	jsonStr := `{"automaticallyImplyLeading": true}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.AutomaticallyImplyLeading == nil {
		t.Fatal("expected automaticallyImplyLeading to have a value")
	}

	if !*attrs.AutomaticallyImplyLeading {
		t.Fatal("expected automaticallyImplyLeading value to be true")
	}
}

func TestAppBarAttributes_AutomaticallyImplyLeading_Unmarshal_False(t *testing.T) {
	jsonStr := `{"automaticallyImplyLeading": false}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.AutomaticallyImplyLeading == nil {
		t.Fatal("expected automaticallyImplyLeading to have a value")
	}

	if *attrs.AutomaticallyImplyLeading {
		t.Fatal("expected automaticallyImplyLeading value to be false")
	}
}

func TestAppBarAttributes_AutomaticallyImplyLeading_Unmarshal_Omitted(t *testing.T) {
	jsonStr := `{}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.AutomaticallyImplyLeading != nil {
		t.Fatal("expected automaticallyImplyLeading to not have a value when omitted")
	}
}

// Tests for ExcludeHeaderSemantics Tristate[bool] property serialization
func TestAppBarAttributes_ExcludeHeaderSemantics_JSON_True(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ExcludeHeaderSemantics: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"excludeHeaderSemantics":true`) {
		t.Fatalf("expected JSON to contain '\"excludeHeaderSemantics\":true', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_ExcludeHeaderSemantics_JSON_False(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ExcludeHeaderSemantics: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"excludeHeaderSemantics":false`) {
		t.Fatalf("expected JSON to contain '\"excludeHeaderSemantics\":false', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_ExcludeHeaderSemantics_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		ExcludeHeaderSemantics: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"excludeHeaderSemantics"`) {
		t.Fatalf("expected JSON to not contain 'excludeHeaderSemantics' field when nil, got: %s", jsonStr)
	}
}

func TestAppBarAttributes_ExcludeHeaderSemantics_Unmarshal_True(t *testing.T) {
	jsonStr := `{"excludeHeaderSemantics": true}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.ExcludeHeaderSemantics == nil {
		t.Fatal("expected excludeHeaderSemantics to have a value")
	}

	if !*attrs.ExcludeHeaderSemantics {
		t.Fatal("expected excludeHeaderSemantics value to be true")
	}
}

func TestAppBarAttributes_ExcludeHeaderSemantics_Unmarshal_False(t *testing.T) {
	jsonStr := `{"excludeHeaderSemantics": false}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.ExcludeHeaderSemantics == nil {
		t.Fatal("expected excludeHeaderSemantics to have a value")
	}

	if *attrs.ExcludeHeaderSemantics {
		t.Fatal("expected excludeHeaderSemantics value to be false")
	}
}

func TestAppBarAttributes_ExcludeHeaderSemantics_Unmarshal_Omitted(t *testing.T) {
	jsonStr := `{}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.ExcludeHeaderSemantics != nil {
		t.Fatal("expected excludeHeaderSemantics to not have a value when omitted")
	}
}

// Tests for Primary Tristate[bool] property serialization
func TestAppBarAttributes_Primary_JSON_True(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		Primary: duit_utils.BoolValue(true),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"primary":true`) {
		t.Fatalf("expected JSON to contain '\"primary\":true', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_Primary_JSON_False(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		Primary: duit_utils.BoolValue(false),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if !strings.Contains(jsonStr, `"primary":false`) {
		t.Fatalf("expected JSON to contain '\"primary\":false', got: %s", jsonStr)
	}
}

func TestAppBarAttributes_Primary_JSON_Nil(t *testing.T) {
	attrs := &duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]{
		Primary: duit_utils.Nillable[bool](),
	}

	jsonData, err := json.Marshal(attrs)
	if err != nil {
		t.Fatal("expected no error for MarshalJSON(), got:", err)
	}

	jsonStr := string(jsonData)
	if strings.Contains(jsonStr, `"primary"`) {
		t.Fatalf("expected JSON to not contain 'primary' field when nil, got: %s", jsonStr)
	}
}

func TestAppBarAttributes_Primary_Unmarshal_True(t *testing.T) {
	jsonStr := `{"primary": true}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.Primary == nil {
		t.Fatal("expected primary to have a value")
	}

	if !*attrs.Primary {
		t.Fatal("expected primary value to be true")
	}
}

func TestAppBarAttributes_Primary_Unmarshal_False(t *testing.T) {
	jsonStr := `{"primary": false}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.Primary == nil {
		t.Fatal("expected primary to have a value")
	}

	if *attrs.Primary {
		t.Fatal("expected primary value to be false")
	}
}

func TestAppBarAttributes_Primary_Unmarshal_Omitted(t *testing.T) {
	jsonStr := `{}`
	
	var attrs duit_attributes.AppBarAttributes[duit_props.ColorString, duit_props.EdgeInsetsAll, duit_decorations.RoundedRectangleBorder[duit_props.ColorString]]
	err := json.Unmarshal([]byte(jsonStr), &attrs)
	if err != nil {
		t.Fatal("expected no error for UnmarshalJSON(), got:", err)
	}

	if attrs.Primary != nil {
		t.Fatal("expected primary to not have a value when omitted")
	}
}
