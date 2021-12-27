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

func TestFirst(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		in []TestType
		fn func(testType TestType) bool
	}
	tests := []struct {
		name    string
		args    args
		want    TestType
		wantErr bool
	}{
		{
			name: "base",
			args: args{
				in: []TestType{
					{1, "alice"},
					{2, "bob"},
					{3, "bob"},
					{2, "carol"},
				},
				fn: func(testType TestType) bool {
					return testType.string == "bob"
				},
			},
			want:    TestType{2, "bob"},
			wantErr: false,
		},
		{
			name: "base_error",
			args: args{
				in: []TestType{
					{1, "alice"},
					{2, "bob"},
					{3, "bob"},
					{2, "carol"},
				},
				fn: func(testType TestType) bool {
					return testType.string == "eve"
				},
			},
			want:    TestType{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := First(tt.args.in, tt.args.fn)

			assert.Equal(t, err != nil, tt.wantErr)
			assert.Equalf(t, tt.want, got, "First(%v, %v)", tt.args.in, tt.args.fn)
		})
	}
}

func TestFirstOr(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		in         []TestType
		fn         func(testType TestType) bool
		defaultVal TestType
	}
	tests := []struct {
		name string
		args args
		want TestType
	}{
		{
			name: "base",
			args: args{
				in: []TestType{
					{1, "alice"},
					{2, "bob"},
					{3, "bob"},
					{2, "carol"},
				},
				fn: func(testType TestType) bool {
					return testType.string == "bob"
				},
				defaultVal: TestType{-1, "default"},
			},
			want: TestType{2, "bob"},
		},
		{
			name: "base_default",
			args: args{
				in: []TestType{
					{1, "alice"},
					{2, "bob"},
					{3, "bob"},
					{2, "carol"},
				},
				fn: func(testType TestType) bool {
					return testType.string == "eve"
				},
				defaultVal: TestType{-1, "default"},
			},
			want: TestType{-1, "default"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FirstOr(tt.args.in, tt.args.fn, tt.args.defaultVal), "FirstOr(%v, %v, %v)", tt.args.in, tt.args.fn, tt.args.defaultVal)
		})
	}
}
