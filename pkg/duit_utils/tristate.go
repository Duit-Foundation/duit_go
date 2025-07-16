package duit_utils

// Tristate представляет указатель на тип T (может быть nil)
type Tristate[T any] *T

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