package main

import (
	"testing"

	"github.com/matryer/is"
)

func Test_partOne(t *testing.T) {
	tests := []struct {
		name  string // description of this test case
		lines []string
		want  int
	}{
		{
			name:  "Test Input",
			lines: GetTestInput(),
			want:  3,
		},
		{
			name:  "Happy Path",
			lines: GetInput(),
			want:  1089,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partOne(tt.lines)
			is.Equal(got, tt.want)
		})
	}
}

func Test_partTwo(t *testing.T) {
	tests := []struct {
		name  string // description of this test case
		lines []string
		want  int
	}{
		{
			name:  "Test Input",
			lines: GetTestInput(),
			want:  6,
		},
		{
			name:  "Happy Path",
			lines: GetInput(),
			want:  6530,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partTwo(tt.lines)
			is.Equal(got, tt.want)
		})
	}
}
