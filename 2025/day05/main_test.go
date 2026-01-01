package main

import (
	"adventofcode2025/utils"
	"slices"
	"testing"

	"github.com/matryer/is"
)

func Test_partOne(t *testing.T) {
	lines, err := utils.ReadInput("day05")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name  string
		lines []string
		want  int
	}{
		{
			name: "Happy Path Test",
			lines: []string{
				"3-5",
				"10-14",
				"16-20",
				"12-18",
				"",
				"1",
				"5",
				"8",
				"11",
				"17",
				"32",
			},
			want: 3,
		},
		{
			name:  "Happy Path",
			lines: lines,
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
	lines, err := utils.ReadInput("day05")
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name  string
		lines []string
		want  int
		wrong []int
	}{
		{
			name: "Happy Path Test",
			lines: []string{
				"3-5",
				"10-14",
				"16-20",
				"12-18",
				"",
				"1",
				"5",
				"8",
				"11",
				"17",
				"32",
			},
			want: 14,
		},
		{
			name:  "Happy Path",
			lines: lines,
			want:  357485433193284,
		},
		{
			name:  "To high",
			lines: lines,
			wrong: []int{465145238590222, 465145238590030},
		},
		{
			name:  "Wrong",
			lines: lines,
			wrong: []int{461367712657565},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partTwo(tt.lines)
			if len(tt.wrong) > 0 {
				if tt.name == "To high" {
					is.Equal(slices.Compare(tt.wrong, []int{got}), 1)
				}
				is.True(!slices.Contains(tt.wrong, got))
				return
			}

			is.Equal(got, tt.want)
		})
	}
}
