package duit_utils

import (
	"errors"
	"fmt"
)

func RangeValidator[T NumConstraint](value T, min T, max T) error {
	if value < min || value > max {
		return errors.New("value must be between the allowed range: min=" + fmt.Sprint(min) + ", max=" + fmt.Sprint(max))
	}
	return nil
}
