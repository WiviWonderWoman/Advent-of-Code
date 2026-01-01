package main

import (
	"testing"

	"github.com/matryer/is"
)

func Test_dayFour(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		total int
		swap  bool
		want  int
	}{
		{
			name:  "happy test input part 1",
			lines: GetTestInput(),
			swap:  false,
			total: 0,
			want:  13,
		},
		{
			name:  "Happy Path Part 1",
			lines: GetInput(),
			swap:  false,
			total: 0,
			want:  1344,
		},
		{
			name:  "happy test input part 2",
			lines: GetTestInput(),
			swap:  true,
			total: 0,
			want:  43,
		},
		{
			name:  "Happy Path Part 2",
			lines: GetInput(),
			swap:  true,
			total: 0,
			want:  8112,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := dayFour(tt.lines, tt.total, tt.swap)
			is.Equal(got, tt.want)
		})
	}
}
