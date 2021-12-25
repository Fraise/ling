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
