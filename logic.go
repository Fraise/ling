package ling

// Contains returns true if the slice contains the value.
func Contains[T comparable](in []T, val T) bool {
	for _, v := range in {
		if v == val {
			return true
		}
	}

	return false
}

// Intersection returns the intersection of the 2 slices.
func Intersection[T comparable](in1 []T, in2 []T) (out []T) {
	out = make([]T, 0)
	in1Map := make(map[T]bool, len(in1))

	for _, v := range in1 {
		in1Map[v] = true
	}

	for _, v := range in2 {
		if in1Map[v] {
			out = append(out, v)
		}
	}

	return out
}
