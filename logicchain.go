package ling

import "errors"

// First returns the first element satisfying the predicate, or the default value of T and an error otherwise.
func (chainer SliceChainer[T]) First(fn func(T) bool) (T, error) {
	for _, v := range chainer {
		if fn(v) {
			return v, nil
		}
	}

	var noVal T

	return noVal, errors.New("value not found")
}

// FirstOr returns the first element satisfying the predicate, or the passed default value otherwise.
func (chainer SliceChainer[T]) FirstOr(fn func(T) bool, defaultVal T) T {
	for _, v := range chainer {
		if fn(v) {
			return v
		}
	}

	return defaultVal
}
