package duit_attributes

import (
	"errors"
	"fmt"

	"github.com/Duit-Foundation/duit_go/v4/pkg/duit_utils"
)

type Validatable interface {
	Validate() error
}

func RangeValidator[T duit_utils.NumConstraint](value T, min T, max T) error {
	if value < min || value > max {
		return errors.New("value must be between the allowed range: min=" + fmt.Sprint(min) + ", max=" + fmt.Sprint(max))
	}
	return nil
}
