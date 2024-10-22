package main

import (
	"testing"
)

// Time complexity: O(n), n is the length of the string s. We traverse the string once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/db0a7f32-697a-4d1e-b6fa-6e36812f1d14
func lengthOfLastWord(s string) int {
	// Length of Last Word
	// Given a string s.
	// Measure the length of the last word in the string.
	// Return the length of the last word.
	// Consider there may be trailing spaces.
	// Words are delimited by spaces.
	// Iterate from the end of the string.

	// Initialize variables.
	cnt := 0
	idx := len(s) - 1

	// Skip any trailing spaces.
	for idx >= 0 && s[idx] == ' ' {
		idx--
	}

	// Measure the length of the last word.
	for idx >= 0 && s[idx] != ' ' {
		cnt++
		idx--
	}

	return cnt
}

// TestLengthOfLastWord_Basic tests the example cases
func TestLengthOfLastWord_Basic(t *testing.T) {
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.str); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestLengthOfLastWord_Edge tests edge cases
func TestLengthOfLastWord_Edge(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
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
			if got := lengthOfLastWord(tt.str); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
