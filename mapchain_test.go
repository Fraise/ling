package ling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceChainer_Update(t *testing.T) {
	type args struct {
		in []int
		fn func(int) int
	}
	tests := []struct {
		name    string
		chainer SliceChainer[int]
		args    args
		want    SliceChainer[int]
	}{
		{
			name: "base",
			args: args{
				in: SliceChainer[int]{
					1, 2, 3,
				},
				fn: func(i int) int {
					return i + 1
				},
			},
			want: SliceChainer[int]{
				2, 3, 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.chainer.Update(tt.args.in, tt.args.fn), "Update(%v, %v)", tt.args.in, tt.args.fn)
		})
	}
}
