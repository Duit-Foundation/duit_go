package duit_utils

// Tristate represents a pointer to type T (can be nil).
//
// Tristate is used in the project in two main contexts:
//
// 1. JSON Serialization with omitempty for bool fields:
//    Solves the problem where Go's omitempty excludes false values from JSON.
//    With Tristate[bool]:
//    - nil: field is omitted from JSON
//    - &false: field appears in JSON as "field":false
//    - &true: field appears in JSON as "field":true
//
//    Example:
//        type Widget struct {
//            Enabled Tristate[bool] `json:"enabled,omitempty"`
//        }
//
// 2. Validation of required fields:
//    Allows distinguishing between "field not provided" and "field has default value".
//    Used in Validate() methods to check if required fields were explicitly set:
//
//    Example:
//        func (w *Widget) Validate() error {
//            if w.Value == nil {
//                return errors.New("value property is required")
//            }
//            return nil
//        }
type Tristate[T any] *T

// TristateFrom converts an input value of any type to a Tristate of type T.
// The function handles different types of input data and returns appropriate
// Tristate representations:
//
// - If value is nil, returns nil Tristate
// - If value is a pointer to T (*T), returns this pointer as is
// - If value is a value of type T, creates a pointer to this value
//   and returns it as Tristate
// - If value doesn't match any of the above cases,
//   returns nil Tristate
//
// Parameters:
//   value - input value of any type for conversion
//
// Returns:
//   Tristate[T] - pointer to value of type T or nil
func TristateFrom[T any](value any) Tristate[T] {
	if value == nil {
		return nil
	}
	
	if ptr, ok := value.(*T); ok {
		return Tristate[T](ptr)
	}
	
	if val, ok := value.(T); ok {
		return Tristate[T](&val)
	}
	
	return nil
}

func StringValue(value string) Tristate[string] {
	return TristateFrom[string](&value)
}

func BoolValue(value bool) Tristate[bool] {
	return TristateFrom[bool](&value)
}

func IntValue(value int) Tristate[int] {
	return TristateFrom[int](&value)
}

func Int8Value(value int8) Tristate[int8] {
	return TristateFrom[int8](&value)
}

func Int16Value(value int16) Tristate[int16] {
	return TristateFrom[int16](&value)
}

func Int32Value(value int32) Tristate[int32] {
	return TristateFrom[int32](&value)
}

func Int64Value(value int64) Tristate[int64] {
	return TristateFrom[int64](&value)
}

func UintValue(value uint) Tristate[uint] {
	return TristateFrom[uint](&value)
}

func Uint8Value(value uint8) Tristate[uint8] {
	return TristateFrom[uint8](&value)
}

func Uint16Value(value uint16) Tristate[uint16] {
	return TristateFrom[uint16](&value)
}

func Uint32Value(value uint32) Tristate[uint32] {
	return TristateFrom[uint32](&value)
}

func Uint64Value(value uint64) Tristate[uint64] {
	return TristateFrom[uint64](&value)
}

func Float32Value(value float32) Tristate[float32] {
	return TristateFrom[float32](&value)
}

func Float64Value(value float64) Tristate[float64] {
	return TristateFrom[float64](&value)
}

func Complex64Value(value complex64) Tristate[complex64] {
	return TristateFrom[complex64](&value)
}

func Complex128Value(value complex128) Tristate[complex128] {
	return TristateFrom[complex128](&value)
}

func Nillable[T any]() Tristate[T] {
	return TristateFrom[T](nil)
}