package ling

// Update updates all the elements of the slice following the predicate.
func (chainer SliceChainer[T]) Update(in []T, fn func(T) T) SliceChainer[T] {
	for i, e := range in {
		in[i] = fn(e)
	}

	return in
}
