package main

import (
	"testing"

	"github.com/matryer/is"
)

func Test_partOne(t *testing.T) {
	tests := []struct {
		name string
		line string
		want int
	}{
		{
			name: "happy test input",
			line: GetTestInput(),
			want: 161,
		},
		{
			name: "Happy Path",
			line: GetInput(),
			want: 164730528,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partOne(tt.line)
			is.Equal(got, tt.want)
		})
	}
}

func Test_partTwo(t *testing.T) {
	tests := []struct {
		name  string
		lines string
		want  int
	}{
		{
			name:  "check test input length",
			lines: GetTestInput(),
			want:  71,
		},
		{
			name:  "check length",
			lines: GetInput(),
			want:  17699,
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
