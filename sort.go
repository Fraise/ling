package ling

import "sort"

// sortable implements the sort.Interface used by the Go implementation to sort a slice.
type sortable[T any] struct {
	Slice  []T
	LessFn func(T, T) bool
}

func (s sortable[T]) Len() int {
	return len(s.Slice)
}

func (s sortable[T]) Less(i, j int) bool {
	return s.LessFn(s.Slice[i], s.Slice[j])
}

func (s sortable[T]) Swap(i, j int) {
	s.Slice[i], s.Slice[j] = s.Slice[j], s.Slice[i]
}

// Sort sorts the passed slice in ascending order using the predicate. If the predicate returns true, the first
// parameter will be placed before the second one. It uses the default Go implementation sort.Sort.
func Sort[T any](in []T, isLess func(first T, second T) bool) (out []T) {
	s := sortable[T]{
		Slice:  in,
		LessFn: isLess,
	}

	sort.Sort(s)

	return s.Slice
}
