package ling

import "runtime"

var numCpu = runtime.NumCPU()

// SetMaxParallel sets the maximum number of goroutines ran at a time in the parallel functions. Defaults to
// runtime.NumCPU().
func SetMaxParallel(max int) {
	if max <= 0 {
		numCpu = 1
		return
	}

	numCpu = max
}

func getChunkSize(length int) int {
	return length/numCpu + 1
}
