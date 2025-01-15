package main

import "testing"

func lengthOfLastWord58b(s string) int {
	return 0
}

func TestLengthOfLastWord58b(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{
			name: "Basic two words",
			str:  "Hello World",
			want: 5,
		},
		{
			name: "Multiple spaces between words",
			str:  "   fly me   to   the moon  ",
			want: 4,
		},
		{
			name: "No trailing spaces",
			str:  "luffy is still joyboy",
			want: 6,
		},
		{
			name: "Single character",
			str:  "a",
			want: 1,
		},
		{
			name: "Single word with trailing spaces",
			str:  "Hello    ",
			want: 5,
		},
		{
			name: "Single word with leading spaces",
			str:  "    Hello",
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord58b(tt.str); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
