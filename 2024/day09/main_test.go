package main

import (
	"slices"
	"testing"

	"github.com/matryer/is"
)

func Test_partOne(t *testing.T) {
	tests := []struct {
		name  string
		line  string
		wrong []int
		want  int
	}{
		{
			name: "happy test input",
			line: GetTestInput(),
			want: 1928,
		},
		{
			name:  "Too low",
			line:  GetInput(),
			wrong: []int{90165373410},
		},
		{
			name: "check length",
			line: GetInput(),
			want: 6323641412437,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partOne(tt.line)
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
		name string
		line string
		want int
	}{
		{
			name: "check test input length",
			line: GetTestInput(),
			want: 19,
		},
		{
			name: "check length",
			line: GetInput(),
			want: 19999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partTwo(tt.line)
			is.Equal(got, tt.want)
		})
	}
}
