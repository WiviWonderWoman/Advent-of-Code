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
			want:  50,
		},
		{
			name:  "Happy Path",
			lines: GetInput(),
			want:  4737096935,
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
			name:  "happy test input",
			lines: GetTestInput(),
			want:  24,
		},
		// {
		// 	name:  "To low",
		// 	lines: GetInput(),
		// 	wrong: []int{853446},
		// },
		// {
		// 	name:  "Happy Path",
		// 	lines: GetInput(),
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partTwo(tt.lines)
			if len(tt.wrong) > 0 {
				if tt.name == "To low" {
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
