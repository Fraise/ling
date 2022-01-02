package ling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	type args struct {
		in []int
		fn func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "base",
			args: args{
				in: []int{
					1, 2, 3, 4, 5,
				},
				fn: func(i int) bool {
					return i > 3
				},
			},
			want: []int{
				4, 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Filter(tt.args.in, tt.args.fn), "Filter(%v, %v)", tt.args.in, tt.args.fn)
		})
	}
}

func BenchmarkFilter1000000(b *testing.B) {
	size := 1000000
	slice := make([]int, 0, size)

	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Filter(slice, func(i int) bool {
			return i < size/2
		})
	}
}

func BenchmarkFilter1000000000(b *testing.B) {
	size := 1000000000
	slice := make([]int, 0, size)

	for i := 0; i < size; i++ {
		slice = append(slice, i)
	}

	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Filter(slice, func(i int) bool {
			return i < size/2
		})
	}
}
