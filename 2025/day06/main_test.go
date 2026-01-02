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
			want:  4277556,
		},
		{
			name:  "Happy Path",
			lines: GetInput(),
			want:  5784380717354,
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
			want:  3263827,
		},
		{
			name:  "Too low",
			lines: GetInput(),
			wrong: []int{3101160432201},
		},
		// ! latest: 8060033472429
		// {
		// 	name:  "Too high",
		// 	lines: GetInput(),
		// 	wrong: []int{7998609941928, 8060033472429},
		// },
		{
			name:  "Wrong",
			lines: GetInput(),
			wrong: []int{79962213186450},
		},
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
