package duit_attributes

import (
	"encoding/json"
	"errors"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_core"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type ListKind uint8

const (
	ListCommon ListKind = iota
	ListBuilder
	ListSeparated
)

type ListViewBaseAttributes struct {
	Type                    duit_utils.Tristate[ListKind]                `json:"type,omitempty"`
	Reverse                 duit_utils.Tristate[bool]                    `json:"reverse,omitempty"`
	Primary                 duit_utils.Tristate[bool]                    `json:"primary,omitempty"`
	ShrinkWrap              duit_utils.Tristate[bool]                    `json:"shrinkWrap,omitempty"`
	AddAutomaticKeepAlives  duit_utils.Tristate[bool]                    `json:"addAutomaticKeepAlives,omitempty"`
	AddRepaintBoundaries    duit_utils.Tristate[bool]                    `json:"addRepaintBoundaries,omitempty"`
	AddSemanticIndexes      duit_utils.Tristate[bool]                    `json:"addSemanticIndexes,omitempty"`
	ScrollDirection         duit_props.Axis                              `json:"scrollDirection,omitempty"`
	ScrollPhysics           duit_props.ScrollPhysics                     `json:"scrollPhysics,omitempty"`
	CacheExtent             float32                                      `json:"cacheExtent,omitempty"`
	Anchor                  float32                                      `json:"anchors,omitempty"`
	SemantickChildCount     int                                          `json:"semanticChildCount,omitempty"`
	Padding                 *duit_props.EdgeInsets                       `json:"padding,omitempty"`
	ItemExtent              float32                                      `json:"itemExtent,omitempty"`
	ClipBehavior            duit_props.Clip                              `json:"clipBehavior,omitempty"`
	RestorationId           string                                       `json:"restorationId,omitempty"`
	DragStarnBehavior       duit_props.DragStartBehavior                 `json:"dragStartBehavior,omitempty"`
	KeyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior `json:"keyboardDismissBehavior,omitempty"`
}

func (r *ListViewBaseAttributes) SetType(listKind ListKind) {
	r.Type = duit_utils.TristateFrom[ListKind](listKind)
}

func (r *ListViewBaseAttributes) SetReverse(reverse bool) {
	r.Reverse = duit_utils.TristateFrom[bool](reverse)
}

func (r *ListViewBaseAttributes) SetPrimary(primary bool) {
	r.Primary = duit_utils.TristateFrom[bool](primary)
}

func (r *ListViewBaseAttributes) SetShrinkWrap(shrinkWrap bool) {
	r.ShrinkWrap = duit_utils.TristateFrom[bool](shrinkWrap)
}

func (r *ListViewBaseAttributes) SetAddAutomaticKeepAlives(addAutomaticKeepAlives bool) {
	r.AddAutomaticKeepAlives = duit_utils.TristateFrom[bool](addAutomaticKeepAlives)
}

func (r *ListViewBaseAttributes) SetAddRepaintBoundaries(addRepaintBoundaries bool) {
	r.AddRepaintBoundaries = duit_utils.TristateFrom[bool](addRepaintBoundaries)
}

func (r *ListViewBaseAttributes) SetAddSemanticIndexes(addSemanticIndexes bool) {
	r.AddSemanticIndexes = duit_utils.TristateFrom[bool](addSemanticIndexes)
}

func (r *ListViewBaseAttributes) SetScrollDirection(scrollDirection duit_props.Axis) {
	r.ScrollDirection = scrollDirection
}

func (r *ListViewBaseAttributes) SetScrollPhysics(scrollPhysics duit_props.ScrollPhysics) {
	r.ScrollPhysics = scrollPhysics
}

func (r *ListViewBaseAttributes) SetCacheExtent(cacheExtent float32) {
	r.CacheExtent = cacheExtent
}

func (r *ListViewBaseAttributes) SetAnchor(anchor float32) {
	r.Anchor = anchor
}

func (r *ListViewBaseAttributes) SetSemanticChildCount(semanticChildCount int) {
	r.SemantickChildCount = semanticChildCount
}

func (r *ListViewBaseAttributes) SetPadding(padding *duit_props.EdgeInsets) {
	r.Padding = padding
}

func (r *ListViewBaseAttributes) SetItemExtent(itemExtent float32) {
	r.ItemExtent = itemExtent
}

func (r *ListViewBaseAttributes) SetClipBehavior(clipBehavior duit_props.Clip) {
	r.ClipBehavior = clipBehavior
}

func (r *ListViewBaseAttributes) SetRestorationId(restorationId string) {
	r.RestorationId = restorationId
}

func (r *ListViewBaseAttributes) SetDragStarnBehavior(dragStarnBehavior duit_props.DragStartBehavior) {
	r.DragStarnBehavior = dragStarnBehavior
}

func (r *ListViewBaseAttributes) SetKeyboardDismissBehavior(keyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior) {
	r.KeyboardDismissBehavior = keyboardDismissBehavior
}

func (r *ListViewBaseAttributes) Validate() error {
	if r.Type != nil {

	} else {
		return errors.New("type is required")
	}

	return nil
}

type ListViewCommonAttributes struct {
	*ListViewBaseAttributes
}

func NewListViewCommonAttributes(data *ListViewCommonAttributes) *ListViewAttributes {
	data.ListViewBaseAttributes.Type = duit_utils.TristateFrom[ListKind](ListCommon)
	return &ListViewAttributes{data: data}
}

type ListViewBuilderAttributes struct {
	*ListViewBaseAttributes
	*Builder
}

func NewListViewBuilderAttributes(data *ListViewBuilderAttributes) *ListViewAttributes {
	data.ListViewBaseAttributes.Type = duit_utils.TristateFrom[ListKind](ListBuilder)
	return &ListViewAttributes{data: data}
}

func (r *ListViewBuilderAttributes) Validate() error {
	if r.ListViewBaseAttributes != nil {
		if err := r.ListViewBaseAttributes.Validate(); err != nil {
			return err
		}

	} else {
		return errors.New("listView property is required")
	}

	if r.Builder == nil {
		return errors.New("builder property is required")
	}

	return nil
}

type ListViewSeparatedAttributes struct {
	*ListViewBaseAttributes
	*Builder
	Separator *duit_core.DuitElementModel `json:"separator"`
}

func NewListViewSeparatedAttributes(data *ListViewSeparatedAttributes) *ListViewAttributes {
	data.ListViewBaseAttributes.Type = duit_utils.TristateFrom[ListKind](ListSeparated)
	return &ListViewAttributes{data: data}
}

func (r *ListViewSeparatedAttributes) SetSeparator(separator *duit_core.DuitElementModel) *ListViewSeparatedAttributes {
	r.Separator = separator
	return r
}

func (r *ListViewSeparatedAttributes) Validate() error {
	if r.ListViewBaseAttributes != nil {
		if err := r.ListViewBaseAttributes.Validate(); err != nil {
			return err
		}
	} else {
		return errors.New("listViewBuilderAttributes property is required")
	}

	if r.Separator == nil {
		return errors.New("separator property is required")
	}

	return nil
}

type ListViewAttributes struct {
	data any
}

func NewListViewAttributes(data any) *ListViewAttributes {
	switch data := data.(type) {
	case *ListViewCommonAttributes:
		return NewListViewCommonAttributes(data)
	case *ListViewBuilderAttributes:
		return NewListViewBuilderAttributes(data)
	case *ListViewSeparatedAttributes:
		return NewListViewSeparatedAttributes(data)
	}
	return nil
}

func (r *ListViewAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.data)
}

func (r *ListViewAttributes) Validate() error {
	switch data := r.data.(type) {
	case *ListViewCommonAttributes:
		return data.Validate()
	case *ListViewBuilderAttributes:
		return data.Validate()
	case *ListViewSeparatedAttributes:
		return data.Validate()
	}
	return nil
}
