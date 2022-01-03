package ling

import "sync"

// ParallelMap maps all the elements of the slice into another slice. It does not preserve the order of the initial
// slice.
//
// This function is much slower than Map for simple filter operations. It should only be used when the
// mapping operation is slow (IO...).
func ParallelMap[T1, T2 any](in []T1, fn func(T1) T2) []T2 {
	slice := make([]T2, 0)
	chunkSize := getChunkSize(len(in))
	receiver := make(chan T2, 1000)
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
				receiver <- fn(element)
			}
		}(i)
	}

	wg.Wait()
	close(receiver)
	<-done

	return slice
}
