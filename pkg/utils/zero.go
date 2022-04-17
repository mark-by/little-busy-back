package utils

import "golang.org/x/exp/constraints"

func DropNil[T any](value *T) T {
	if value == nil {
		var zero T
		return zero
	}
	return *value
}

func Nullable[T constraints.Ordered](value T) *T {
	var zero T
	if value == zero {
		return nil
	}
	return &value
}
