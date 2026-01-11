package main

import (
	"fmt"
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
			want:  14,
		},
		{
			name:  "Too Low",
			lines: GetInput(),
			wrong: []int{219},
		},
		{
			name:  "Too high",
			lines: GetInput(),
			wrong: []int{386},
		},
		{
			name:  "check length",
			lines: GetInput(),
			want:  320,
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
		wrong []int
		want  int
	}{
		{
			name:  "happy test input",
			lines: GetTestInput(),
			want:  34,
		},
		{
			name: "test input",
			lines: []string{
				"T.........",
				"...T......",
				".T........",
				"..........",
				"..........",
				"..........",
				"..........",
				"..........",
				"..........",
				"..........",
			},
			want: 9,
		},
		{
			name:  "Too Low",
			lines: GetInput(),
			wrong: []int{1137},
		},
		{
			name:  "Too high",
			lines: GetInput(),
			wrong: []int{2974},
		},
		{
			name:  "check length",
			lines: GetInput(),
			want:  1157,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := partTwo(tt.lines)
			if len(tt.wrong) > 0 {
				if tt.name == "Too low" {
					fmt.Println("GOT: ", got)
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
