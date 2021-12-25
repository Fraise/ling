package ling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		in  []TestType
		val TestType
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "base",
			args: args{
				in: []TestType{
					{1, "alice"},
					{1, "bob"},
					{2, "carol"},
				},
				val: TestType{1, "bob"},
			},
			want: true,
		},
		{
			name: "base_false",
			args: args{
				in: []TestType{
					{1, "alice"},
					{1, "bob"},
					{2, "carol"},
				},
				val: TestType{2, "bob"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Contains(tt.args.in, tt.args.val), "Contains(%v, %v)", tt.args.in, tt.args.val)
		})
	}
}

func TestIntersection(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		in1 []TestType
		in2 []TestType
	}
	tests := []struct {
		name    string
		args    args
		wantOut []TestType
	}{
		{
			name: "base",
			args: args{
				in1: []TestType{
					{1, "alice"},
					{1, "bob"},
					{2, "carol"},
				},
				in2: []TestType{
					{1, "alice"},
					{2, "bob"},
					{2, "carol"},
				},
			},
			wantOut: []TestType{
				{1, "alice"},
				{2, "carol"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantOut, Intersection(tt.args.in1, tt.args.in2), "Intersection(%v, %v)", tt.args.in1, tt.args.in2)
		})
	}
}

func TestDistinct(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		in []TestType
	}
	tests := []struct {
		name    string
		args    args
		wantOut []TestType
	}{
		{
			name: "base",
			args: args{
				in: []TestType{
					{1, "alice"},
					{2, "bob"},
					{2, "bob"},
					{2, "carol"},
					{2, "carol"},
					{2, "carol"},
				},
			},
			wantOut: []TestType{
				{1, "alice"},
				{2, "bob"},
				{2, "carol"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantOut, Distinct(tt.args.in), "Distinct(%v)", tt.args.in)
		})
	}
}
