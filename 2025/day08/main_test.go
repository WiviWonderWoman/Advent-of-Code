package main

import (
	"testing"

	"github.com/matryer/is"
)

func Test_partOne(t *testing.T) {
	lines := GetInput()
	tests := []struct {
		name    string
		lines   []string
		mutiply int
		wrong   []int
		want    int
	}{
		{
			name:    "Test input",
			lines:   GetTestInput(),
			mutiply: 10,
			want:    40,
		},
		{
			name:    "Happy path",
			lines:   lines,
			mutiply: 1000,
			want:    129564,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partOne(tt.lines, tt.mutiply)

			is.Equal(got, tt.want)
		})
	}
}

func Test_partTwo(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			name:  "check test input length",
			lines: GetTestInput(),
			want:  20,
		},
		{
			name:  "check length",
			lines: GetInput(),
			want:  1000,
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
