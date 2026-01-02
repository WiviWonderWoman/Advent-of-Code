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
			name:  "check test input length",
			lines: GetTestInput(),
			want:  21,
		},
		{
			name:  "check length",
			lines: GetInput(),
			want:  1681,
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
			want:  40,
		},
		// {
		// 	name:  "Too low",
		// 	lines: GetInput(),
		// 	wrong: []int{3073},
		// },
		// {
		// 	name:  "check length",
		// 	lines: GetInput(),
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partTwo(tt.lines)
			if len(tt.wrong) > 0 {
				if tt.name == "Too low" {
					for _, w := range tt.wrong {
						is.Equal(slices.Compare([]int{got}, []int{w}), 1)
					}
				}

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
