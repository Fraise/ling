package ling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParallelFilter(t *testing.T) {
	type TestType struct {
		int
		string
	}

	testSlice := make([]TestType, 0, 1000)

	for i := 0; i < 1000; i++ {
		testSlice = append(testSlice, TestType{
			int:    i,
			string: "bob",
		})
	}

	type args struct {
		in []TestType
		fn func(testType TestType) bool
	}
	tests := []struct {
		name string
		args args
		want []TestType
	}{
		{
			name: "base",
			args: args{
				in: []TestType{
					{1, "alice"},
					{2, "bob"},
					{3, "bob"},
					{2, "carol"},
				},
				fn: func(i TestType) bool {
					return i.int > 2
				},
			},
			want: []TestType{{3, "bob"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ParallelFilter(tt.args.in, tt.args.fn), "ParallelFilter(%v, %v)", tt.args.in, tt.args.fn)
		})
	}
}

func BenchmarkParallelFilter1000000(b *testing.B) {
	size := 1000000
	slice := make([]int, 0, size)

	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ParallelFilter(slice, func(i int) bool {
			return i < size/2
		})
	}
}

func BenchmarkParallelFilter1000000000(b *testing.B) {
	size := 1000000000
	slice := make([]int, 0, size)

	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ParallelFilter(slice, func(i int) bool {
			return i < size/2
		})
	}
}
