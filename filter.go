package ling

// Filter removes the elements from the slice when the predicate is false.
func Filter[T any](in []T, fn func(T) bool) []T {
	slice := make([]T, 0)

	for _, element := range in {
		if fn(element) {
			slice = append(slice, element)
		}
	}

	return slice
}
