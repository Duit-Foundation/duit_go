package duit_utils

// Tristate represents a pointer to type T (can be nil)
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
		return ptr
	}
	
	if val, ok := value.(T); ok {
		return &val
	}
	
	return nil
}