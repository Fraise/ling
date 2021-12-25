package ling

import "constraints"

// Min returns the minimum value of the slice.
func Min[T constraints.Ordered](in []T) (min T) {
	for i, e := range in {
		if i == 0 {
			min = e
			continue
		}
		if e < min {
			min = e
		}
	}

	return min
}

// Max returns the maximum value of the slice.
func Max[T constraints.Ordered](in []T) (max T) {
	for i, e := range in {
		if i == 0 {
			max = e
			continue
		}
		if e > max {
			max = e
		}
	}

	return max
}

// Count returns the number of elements of the slice satisfying a predicate.
func Count[T any](in []T, fn func(T) bool) (count int) {
	for _, v := range in {
		if fn(v) {
			count++
		}
	}

	return count
}
