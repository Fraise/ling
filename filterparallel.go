package ling

import "sync"

// ParallelFilter removes the elements from the slice when the predicate is false using multiple goroutines.
// This function does not keep the ordering of the input slice.
//
// This function is much slower than Filter (~ 10 times) for simple filter operations. It should only be used when the
// filter operation is slow (IO...).
func ParallelFilter[T any](in []T, fn func(T) bool) []T {
	slice := make([]T, 0)
	chunkSize := getChunkSize(len(in))
	receiver := make(chan T, 1000)
	done := make(chan bool)
	var wg sync.WaitGroup

	go func() {
		for val := range receiver {
			slice = append(slice, val)
		}
		done <- true
	}()

	for i := 0; i < len(in); i += chunkSize {
		wg.Add(1)
		go func(iVal int) {
			defer wg.Done()

			end := iVal + chunkSize

			if iVal+chunkSize >= len(in) {
				end = len(in)
			}

			for _, element := range in[iVal:end] {
				if fn(element) {
					receiver <- element
				}
			}
		}(i)
	}

	wg.Wait()
	close(receiver)
	<-done

	return slice
}
