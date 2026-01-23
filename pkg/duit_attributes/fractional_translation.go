package duit_attributes

import (
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_props"
	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type FractionalTranslationAttributes struct {
	*ValueReferenceHolder             `gen:"wrap"`
	*ThemeConsumer                    `gen:"wrap"`
	Translation                       *duit_props.Offset `json:"translation,omitempty"`
	TransformHitTests                 duit_utils.TBool   `json:"transformHitTests,omitempty"`
}

func NewFractionalTranslationAttributes() *FractionalTranslationAttributes {
	return &FractionalTranslationAttributes{}
}

func (r *FractionalTranslationAttributes) SetTranslation(translation *duit_props.Offset) *FractionalTranslationAttributes {
	r.Translation = translation
	return r
}

func (r *FractionalTranslationAttributes) SetTransformHitTests(transformHitTests bool) *FractionalTranslationAttributes {
	r.TransformHitTests = duit_utils.BoolValue(transformHitTests)
	return r
}

func (r *FractionalTranslationAttributes) Validate() error {
	if err := r.ValueReferenceHolder.Validate(); err != nil {
		return err
	}

	return nil
}