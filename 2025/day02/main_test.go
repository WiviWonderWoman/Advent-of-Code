package main

import (
	"testing"

	"github.com/matryer/is"
)

func Test_dayTwo(t *testing.T) {
	tests := []struct {
		name  string
		lines []string
		part  int
		want  int
	}{
		{
			name:  "happy test input part 1",
			lines: GetTestInput(),
			part:  1,
			want:  1227775554,
		},
		{
			name:  "Happy Path part 1",
			lines: GetInput(),
			part:  1,
			want:  52316131093,
		},
		{
			name:  "TEST input part 2",
			lines: []string{"565653-565659", "824824821-824824827", "2121212118-2121212124"},
			part:  2,
			want:  2946602601,
		},
		{
			name:  "happy test input part 2",
			lines: GetTestInput(),
			part:  2,
			want:  4174379265,
		},
		{
			name:  "Happy Path part 2",
			lines: GetInput(),
			part:  2,
			want:  69564213293,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)
			got := dayTwo(tt.lines, tt.part)
			is.Equal(got, tt.want)
		})
	}
}
