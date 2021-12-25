package ling

// Filter removes the elements from the *SliceChainer when the predicate is false.
func (chainer SliceChainer[T]) Filter(fn func(T) bool) SliceChainer[T] {
	slice := make(SliceChainer[T], 0)

	for _, element := range chainer {
		if fn(element) {
			slice = append(slice, element)
		}
	}

	return slice
}
