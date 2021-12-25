package ling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceChainer_Sort(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		isLess func(first TestType, second TestType) bool
	}
	tests := []struct {
		name    string
		chainer SliceChainer[TestType]
		args    args
		want    SliceChainer[TestType]
	}{
		{
			name: "base",
			chainer: SliceChainer[TestType]{
				{3, "alice"},
				{1, "bob"},
				{2, "carol"},
			},
			args: args{

				isLess: func(first TestType, second TestType) bool {
					return first.int < second.int
				},
			},
			want: SliceChainer[TestType]{
				{1, "bob"},
				{2, "carol"},
				{3, "alice"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.chainer.Sort(tt.args.isLess), "Sort(%v)", tt.args.isLess)
		})
	}
}
