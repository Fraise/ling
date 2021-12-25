package ling

import "sort"

// Sort sorts the passed slice in ascending order using the predicate. If the predicate returns true, the first
// parameter will be placed before the second one. It uses the default Go implementation sort.Sort.
func (chainer SliceChainer[T]) Sort(isLess func(first T, second T) bool) SliceChainer[T] {
	s := sortable[T]{
		Slice:  chainer,
		LessFn: isLess,
	}

	sort.Sort(s)

	return chainer
}
