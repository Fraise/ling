package ling

type Enumerable[T any] struct {
	slice []T
}

type SliceChainer[T any] []T

func CopyIntoChainer[T any](in []T) *SliceChainer[T] {
	chainer := make(SliceChainer[T], len(in))

	copy(chainer, in)

	return &chainer
}

func IntoChainer[T any](in *[]T) *SliceChainer[T] {
	return (*SliceChainer[T])(in)
}

func (chainer *SliceChainer[T]) Filter(functions ...func(T) bool) *SliceChainer[T] {
	slice := make(SliceChainer[T], 0)

	for _, element := range *chainer {

		for _, fn := range functions {
			if fn(element) {
				slice = append(slice, element)
			}
		}
	}

	return &slice
}
