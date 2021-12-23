package ling

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	type outType struct {
		Val int
	}

	type args struct {
		e1 Enumerable[int]
		fn func(int) outType
	}
	tests := []struct {
		name   string
		args   args
		wantE2 Enumerable[outType]
	}{
		{
			name: "base",
			args: args{
				e1: Enumerable[int]{
					slice: []int{
						1, 2,
					},
				},
				fn: func(i int) outType {
					return outType{
						Val: i,
					}
				},
			},
			wantE2: Enumerable[outType]{
				slice: []outType{
					{Val: 1},
					{Val: 2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotE2 := Map(tt.args.e1, tt.args.fn); !reflect.DeepEqual(gotE2, tt.wantE2) {
				t.Errorf("Map() = %v, want %v", gotE2, tt.wantE2)
			}
		})
	}
}
