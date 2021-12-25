package ling

type SliceChainer[T any] []T

// CopyIntoChainer creates a copy of the passed slice and creates a SliceChainer that point to it.
func CopyIntoChainer[T any](in []T) SliceChainer[T] {
	chainer := make(SliceChainer[T], len(in))

	copy(chainer, in)

	return chainer
}

// IntoChainer creates a SliceChainer with the passed slice as underlying type.
func IntoChainer[T any](in []T) SliceChainer[T] {
	return in
}

// ToSlice returns the slice pointed by the *SliceChainer.
func (chainer SliceChainer[T]) ToSlice() []T {
	return chainer
}
