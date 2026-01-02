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
		want  int
	}{
		{
			name:  "happy test input",
			lines: GetTestInput(),
			want:  18,
		},
		{
			name:  "Happy Path",
			lines: GetInput(),
			want:  2685,
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
		name  string
		lines []string
		wrong []int
		want  int
	}{
		{
			name:  "check test input length",
			lines: GetTestInput(),
			want:  19,
		},
		{
			name:  "Too high",
			lines: GetInput(),
			wrong: []int{2132, 2190, 2685},
		},
		// {
		// 	name:  "check length",
		// 	lines: GetInput(),
		// 	want:  140,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partTwo(tt.lines)
			if len(tt.wrong) > 0 {
				if tt.name == "Too high" {
					is.Equal(slices.Compare(tt.wrong, []int{got}), 1)
				}
				is.True(!slices.Contains(tt.wrong, got))
				return
			}
			is.Equal(got, tt.want)
		})
	}
}
