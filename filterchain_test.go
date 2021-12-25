package ling

import (
	"reflect"
	"testing"
)

func TestSliceChainer_Filter(t *testing.T) {
	type args struct {
		fn func(int) bool
	}
	tests := []struct {
		name    string
		chainer SliceChainer[int]
		args    args
		want    SliceChainer[int]
	}{
		{
			name: "base",
			chainer: SliceChainer[int]{
				1, 2, 3, 4, 5,
			},
			args: args{
				func(i int) bool {
					return i > 3
				},
			},
			want: SliceChainer[int]{
				4, 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.chainer.Filter(tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
