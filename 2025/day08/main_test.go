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
		laps  int
		wrong []int
		want  int
	}{
		{
			name:  "Test input",
			lines: GetTestInput(),
			laps:  11,
			want:  40,
		},
		// {
		// 	name:  "Too low",
		// 	lines: GetInput(),
		// 	laps:  1000,
		// 	wrong: []int{252, 16896},
		// },
		// {
		// 	name:  "Happy path",
		// 	lines: GetInput(),
		// 	laps:  1000,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partOne(tt.lines, tt.laps)
			if len(tt.wrong) > 0 {
				if tt.name == "Too low" {
					for _, w := range tt.wrong {
						is.Equal(slices.Compare([]int{got}, []int{w}), 1)
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partTwo(tt.lines)
			is.Equal(got, tt.want)
		})
	}
}
