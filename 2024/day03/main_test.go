package main

import (
	"slices"
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
			line: GetTestInputPartOne(),
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
		wrong []int
		want  int
	}{
		{
			name:  "happy test input",
			lines: GetTestInputPartTwo(),
			want:  48,
		},
		{
			name:  "Too high",
			lines: GetInput(),
			wrong: []int{101252657},
		},
		{
			name:  "check length",
			lines: GetInput(),
			want:  70478672,
		},
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
