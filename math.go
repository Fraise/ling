package ling

import "constraints"

func Min[T constraints.Ordered](in *SliceChainer[T]) T {
	var min T

	for i, e := range *in {
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

func Max[T constraints.Ordered](in *SliceChainer[T]) T {
	var max T

	for i, e := range *in {
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
