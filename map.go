package ling

func Map[T1, T2 any](e1 Enumerable[T1], fn func(T1) T2) (e2 Enumerable[T2]) {
	e2.slice = make([]T2, len(e1.slice))

	for i, e := range e1.slice {
		e2.slice[i] = fn(e)
	}

	return e2
}
