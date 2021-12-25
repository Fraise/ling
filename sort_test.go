package ling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		in     []TestType
		isLess func(first TestType, second TestType) bool
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
					{3, "alice"},
					{1, "bob"},
					{2, "carol"},
				},
				isLess: func(first TestType, second TestType) bool {
					return first.int < second.int
				},
			},
			wantOut: []TestType{
				{1, "bob"},
				{2, "carol"},
				{3, "alice"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantOut, Sort(tt.args.in, tt.args.isLess), "Sort(%v, %v)", tt.args.in, tt.args.isLess)
		})
	}
}
