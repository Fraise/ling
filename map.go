package ling

// Map maps all the elements of the slice into another slice.
func Map[T1, T2 any](in []T1, fn func(T1) T2) (out []T2) {
	out = make([]T2, len(in))

	for i, e := range in {
		out[i] = fn(e)
	}

	return out
}

// AggregateToSlice groups together elements the key K returned by the predicate and return them as a slice of slices.
func AggregateToSlice[T any, K comparable](in []T, fn func(T) (K, T)) (out [][]T) {
	aggregate := Aggregate(in, fn)

	out = make([][]T, 0, len(aggregate))

	for _, v := range aggregate {
		out = append(out, v)
	}

	return out
}

// Aggregate groups together elements in a map by the key K returned by the predicate.
func Aggregate[T any, K comparable](in []T, fn func(T) (K, T)) (out map[K][]T) {
	out = make(map[K][]T)

	for _, v := range in {
		key, val := fn(v)

		if out[key] == nil {
			out[key] = make([]T, 0)
		}

		out[key] = append(out[key], val)
	}

	return out
}

// Update updates all the elements of the slice following the predicate.
func Update[T any](in []T, fn func(T) T) []T {
	for i, e := range in {
		in[i] = fn(e)
	}

	return in
}
