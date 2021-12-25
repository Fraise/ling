package ling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	type outType struct {
		Val int
	}

	type args struct {
		in []int
		fn func(int) outType
	}
	tests := []struct {
		name    string
		args    args
		wantOut []outType
	}{
		{
			name: "base",
			args: args{
				in: []int{
					1, 2,
				},
				fn: func(i int) outType {
					return outType{
						Val: i,
					}
				},
			},
			wantOut: []outType{
				{Val: 1},
				{Val: 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut := Map(tt.args.in, tt.args.fn)
			assert.Equal(t, tt.wantOut, gotOut)
		})
	}
}

func TestAggregateToSlice(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		in []TestType
		fn func(TestType) (int, TestType)
	}
	tests := []struct {
		name    string
		args    args
		wantOut [][]TestType
	}{
		{
			name: "base",
			args: args{
				in: []TestType{
					{1, "alice"},
					{2, "bob"},
					{1, "carol"},
				},
				fn: func(first TestType) (int, TestType) {
					return first.int, first
				},
			},
			wantOut: [][]TestType{
				{{1, "alice"}, {1, "carol"}},
				{{2, "bob"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut := AggregateToSlice(tt.args.in, tt.args.fn)
			assert.Equal(t, tt.wantOut, gotOut)
		})
	}
}

func TestAggregate(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		in []TestType
		fn func(TestType) (int, TestType)
	}
	tests := []struct {
		name    string
		args    args
		wantOut map[int][]TestType
	}{
		{
			name: "base",
			args: args{
				in: []TestType{
					{1, "alice"},
					{2, "bob"},
					{1, "carol"},
				},
				fn: func(first TestType) (int, TestType) {
					return first.int, first
				},
			},
			wantOut: map[int][]TestType{
				1: {{1, "alice"}, {1, "carol"}},
				2: {{2, "bob"}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantOut, Aggregate(tt.args.in, tt.args.fn), "Aggregate(%v, %v)", tt.args.in, tt.args.fn)
		})
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		in []int
		fn func(int) int
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
					1, 2, 3,
				},
				fn: func(i int) int {
					return i + 1
				},
			},
			want: []int{
				2, 3, 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Update(tt.args.in, tt.args.fn), "Update(%v, %v)", tt.args.in, tt.args.fn)
		})
	}
}
