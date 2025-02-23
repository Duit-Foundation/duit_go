package duit_utils

func BoolPtr(v bool) *bool {
	ptr := new(bool)

	*ptr = v
	return ptr
}