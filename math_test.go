package ling

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base",
			args: args{
				in: []int{
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
		in []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base",
			args: args{
				in: []int{
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

func TestCount(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		in []TestType
		fn func(TestType) bool
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "base",
			args: args{
				in: []TestType{
					{1, "alice"},
					{1, "bob"},
					{2, "carole"},
					{3, "eve"},
				},
				fn: func(testType TestType) bool {
					return testType.int == 1 || testType.string == "eve"
				},
			},
			wantCount: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantCount, Count(tt.args.in, tt.args.fn), "Count(%v, %v)", tt.args.in, tt.args.fn)
		})
	}
}
