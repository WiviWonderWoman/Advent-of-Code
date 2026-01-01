package main

import (
	"testing"

	"github.com/matryer/is"
)

func Test_dayThree(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		turns int
		want  int
	}{
		{
			name:  "happy test input part 1",
			lines: GetTestInput(),
			turns: 1,
			want:  357,
		},
		{
			name:  "Happy Path Part 1",
			lines: GetInput(),
			turns: 1,
			want:  16927,
		},
		{
			name:  "happy test input part 2",
			lines: GetTestInput(),
			turns: 11,
			want:  3121910778619,
		},
		{
			name:  "Happy Path Part 2",
			lines: GetInput(),
			turns: 11,
			want:  167384358365132,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := dayThree(tt.lines, tt.turns)
			is.Equal(got, tt.want)
		})
	}
}
