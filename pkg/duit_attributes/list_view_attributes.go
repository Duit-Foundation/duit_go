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

// ListViewBaseAttributes defines attributes for ListViewBase widget.
// See: https://api.flutter.dev/flutter/widgets/ListView/ListView.html
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

// NewListViewBaseAttributes creates a new ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetType(listKind ListKind) {
	r.Type = duit_utils.TristateFrom[ListKind](listKind)
}

// SetReverse sets the reverse for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetReverse(reverse bool) {
	r.Reverse = duit_utils.TristateFrom[bool](reverse)
}

// SetPrimary sets the primary for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetPrimary(primary bool) {
	r.Primary = duit_utils.TristateFrom[bool](primary)
}

// SetShrinkWrap sets the shrink wrap for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetShrinkWrap(shrinkWrap bool) {
	r.ShrinkWrap = duit_utils.TristateFrom[bool](shrinkWrap)
}

// SetAddAutomaticKeepAlives sets the add automatic keep alive for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetAddAutomaticKeepAlives(addAutomaticKeepAlives bool) {
	r.AddAutomaticKeepAlives = duit_utils.TristateFrom[bool](addAutomaticKeepAlives)
}

// SetAddRepaintBoundaries sets the add repaint boundaries for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetAddRepaintBoundaries(addRepaintBoundaries bool) {
	r.AddRepaintBoundaries = duit_utils.TristateFrom[bool](addRepaintBoundaries)
}

// SetAddSemanticIndexes sets the add semantic indexes for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetAddSemanticIndexes(addSemanticIndexes bool) {
	r.AddSemanticIndexes = duit_utils.TristateFrom[bool](addSemanticIndexes)
}

// SetScrollDirection sets the scroll direction for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetScrollDirection(scrollDirection duit_props.Axis) {
	r.ScrollDirection = scrollDirection
}

// SetScrollPhysics sets the scroll physics for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetScrollPhysics(scrollPhysics duit_props.ScrollPhysics) {
	r.ScrollPhysics = scrollPhysics
}

// SetCacheExtent sets the cache extent for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetCacheExtent(cacheExtent float32) {
	r.CacheExtent = cacheExtent
}

// SetAnchor sets the anchor for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetAnchor(anchor float32) {
	r.Anchor = anchor
}

// SetSemanticChildCount sets the semantic child count for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetSemanticChildCount(semanticChildCount int) {
	r.SemantickChildCount = semanticChildCount
}

// SetPadding sets the padding for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetPadding(padding *duit_props.EdgeInsets) {
	r.Padding = padding
}

// SetItemExtent sets the item extent for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetItemExtent(itemExtent float32) {
	r.ItemExtent = itemExtent
}

// SetClipBehavior sets the clip behavior for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetClipBehavior(clipBehavior duit_props.Clip) {
	r.ClipBehavior = clipBehavior
}

// SetRestorationId sets the restoration id for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetRestorationId(restorationId string) {
	r.RestorationId = restorationId
}

// SetDragStarnBehavior sets the drag start behavior for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetDragStarnBehavior(dragStarnBehavior duit_props.DragStartBehavior) {
	r.DragStarnBehavior = dragStarnBehavior
}

// SetKeyboardDismissBehavior sets the keyboard dismiss behavior for the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) SetKeyboardDismissBehavior(keyboardDismissBehavior duit_props.ScrollViewKeyboardDismissBehavior) {
	r.KeyboardDismissBehavior = keyboardDismissBehavior
}

// Validate validates the ListViewBaseAttributes instance.
func (r *ListViewBaseAttributes) Validate() error {
	if r.Type != nil {

	} else {
		return errors.New("type is required")
	}

	return nil
}

// ListViewCommonAttributes defines attributes for ListViewCommon widget.
// See: https://api.flutter.dev/flutter/widgets/ListView/ListView.html
type ListViewCommonAttributes struct {
	*ListViewBaseAttributes
}

// NewListViewCommonAttributes creates a new ListViewCommonAttributes instance.
func NewListViewCommonAttributes(data *ListViewCommonAttributes) *ListViewAttributes {
	data.ListViewBaseAttributes.Type = duit_utils.TristateFrom[ListKind](ListCommon)
	return &ListViewAttributes{data: data}
}

// ListViewBuilderAttributes defines attributes for ListViewBuilder widget.
// See: https://api.flutter.dev/flutter/widgets/ListView/ListView.builder.html
type ListViewBuilderAttributes struct {
	*ListViewBaseAttributes
	*Builder
}

// NewListViewBuilderAttributes creates a new ListViewBuilderAttributes instance.
func NewListViewBuilderAttributes(data *ListViewBuilderAttributes) *ListViewAttributes {
	data.ListViewBaseAttributes.Type = duit_utils.TristateFrom[ListKind](ListBuilder)
	return &ListViewAttributes{data: data}
}

// Validate validates the ListViewBuilderAttributes instance.
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

// ListViewSeparatedAttributes defines attributes for ListViewSeparated widget.
// See: https://api.flutter.dev/flutter/widgets/ListView/ListView.separated.html
type ListViewSeparatedAttributes struct {
	*ListViewBaseAttributes
	*Builder
	Separator *duit_core.DuitElementModel `json:"separator"`
}

// NewListViewSeparatedAttributes creates a new ListViewSeparatedAttributes instance.
func NewListViewSeparatedAttributes(data *ListViewSeparatedAttributes) *ListViewAttributes {
	data.ListViewBaseAttributes.Type = duit_utils.TristateFrom[ListKind](ListSeparated)
	return &ListViewAttributes{data: data}
}

// SetSeparator sets the separator for the ListViewSeparatedAttributes instance.
func (r *ListViewSeparatedAttributes) SetSeparator(separator *duit_core.DuitElementModel) *ListViewSeparatedAttributes {
	r.Separator = separator
	return r
}

// Validate validates the ListViewSeparatedAttributes instance.
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

// ListViewAttributes defines attributes for ListView widget.
// See: https://api.flutter.dev/flutter/widgets/ListView/ListView.html
type ListViewAttributes struct {
	data any
}

// NewListViewAttributes creates a new ListViewAttributes instance.
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
