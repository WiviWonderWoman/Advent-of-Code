package main

import (
	"slices"
	"testing"

	"github.com/matryer/is"
)

func Test_partOne(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		wrong []int
		want  int
	}{
		{
			name:  "happy test input",
			lines: GetTestInput(),
			want:  3749,
		},
		{
			name:  "Too low",
			lines: GetInput(),
			wrong: []int{84421892209, 84421962527, 84425653363, 84425834370, 88615284744, 94165549118, 99685075168},
		},
		{
			name:  "Happy Path",
			lines: GetInput(),
			want:  1545311493300,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partOne(tt.lines)
			if len(tt.wrong) > 0 {
				if tt.name == "Too low" {
					for _, w := range tt.wrong {
						is.Equal(slices.Compare([]int{got}, []int{w}), 1)
					}
				}

				if tt.name == "Too high" {
					for _, w := range tt.wrong {
						is.Equal(slices.Compare([]int{w}, []int{got}), 1)
					}
				}

				is.True(!slices.Contains(tt.wrong, got))
				return
			}
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
			want:  9,
		},
		{
			name:  "check length",
			lines: GetInput(),
			want:  850,
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
