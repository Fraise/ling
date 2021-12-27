package ling

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceChainer_First(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		fn func(testType TestType) bool
	}
	tests := []struct {
		name    string
		chainer SliceChainer[TestType]
		args    args
		want    TestType
		wantErr bool
	}{
		{
			name: "base",
			chainer: SliceChainer[TestType]{
				{1, "alice"},
				{2, "bob"},
				{3, "bob"},
				{2, "carol"},
			},
			args: args{
				fn: func(testType TestType) bool {
					return testType.string == "bob"
				},
			},
			want:    TestType{2, "bob"},
			wantErr: false,
		},
		{
			name: "base_error",
			chainer: SliceChainer[TestType]{
				{1, "alice"},
				{2, "bob"},
				{3, "bob"},
				{2, "carol"},
			},
			args: args{
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
			got, err := tt.chainer.First(tt.args.fn)

			assert.Equal(t, err != nil, tt.wantErr)
			assert.Equalf(t, tt.want, got, "First(%v)", tt.args.fn)
		})
	}
}

func TestSliceChainer_FirstOr(t *testing.T) {
	type TestType struct {
		int
		string
	}

	type args struct {
		fn         func(testType TestType) bool
		defaultVal TestType
	}
	tests := []struct {
		name    string
		chainer SliceChainer[TestType]
		args    args
		want    TestType
	}{
		{
			name: "base",
			chainer: SliceChainer[TestType]{
				{1, "alice"},
				{2, "bob"},
				{3, "bob"},
				{2, "carol"},
			},
			args: args{
				fn: func(testType TestType) bool {
					return testType.string == "bob"
				},
				defaultVal: TestType{-1, "default"},
			},
			want: TestType{2, "bob"},
		},
		{
			name: "base_default",
			chainer: SliceChainer[TestType]{
				{1, "alice"},
				{2, "bob"},
				{3, "bob"},
				{2, "carol"},
			},
			args: args{
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
			assert.Equalf(t, tt.want, tt.chainer.FirstOr(tt.args.fn, tt.args.defaultVal), "FirstOr(%v, %v)", tt.args.fn, tt.args.defaultVal)
		})
	}
}
