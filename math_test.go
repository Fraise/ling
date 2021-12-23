package ling

import (
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	type args struct {
		in *SliceChainer[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base",
			args: args{
				in: &SliceChainer[int]{
					1, 3, 2, -1,
				},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		in *SliceChainer[int]
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base",
			args: args{
				in: &SliceChainer[int]{
					1, 3, 2, -1,
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
