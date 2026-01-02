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
			name:  "Happy Path Test",
			lines: GetTestInput(),
			want:  3,
		},
		{
			name:  "Happy Path",
			lines: GetInput(),
			want:  737,
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
		want  int
		wrong []int
	}{
		{
			name:  "Happy Path Test",
			lines: GetTestInput(),
			want:  14,
		},
		{
			name:  "Happy Path",
			lines: GetInput(),
			want:  357485433193284,
		},
		{
			name:  "Too high",
			lines: GetInput(),
			wrong: []int{465145238590222, 465145238590030},
		},
		{
			name:  "Wrong",
			lines: GetInput(),
			wrong: []int{461367712657565},
		},
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
