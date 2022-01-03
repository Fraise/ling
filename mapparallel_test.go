package ling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParallelMap(t *testing.T) {
	type outType struct {
		Val int
	}

	contains := func(in []outType, val outType) bool {
		for _, v := range in {
			if v == val {
				return true
			}
		}

		return false
	}

	type args struct {
		in []int
		fn func(int) outType
	}
	tests := []struct {
		name string
		args args
		want []outType
	}{
		{
			name: "base",
			args: args{
				in: []int{
					1, 2, 3, 4, 5, 6, 7,
				},
				fn: func(i int) outType {
					return outType{
						Val: i,
					}
				},
			},
			want: []outType{
				{Val: 1},
				{Val: 2},
				{Val: 3},
				{Val: 4},
				{Val: 5},
				{Val: 6},
				{Val: 7},
			},
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := ParallelMap(tt.args.in, tt.args.fn)
			assert.Equal(t, len(tt.want), len(res))

			for _, v := range tt.want {
				assert.True(t, contains(res, v))
			}
		})
	}
}
