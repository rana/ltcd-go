package main

import "testing"

// Time complexity: O(n)
// Space complexity: O(min(m, n)), m is lenght of charset. n is length of string s.
func lengthOfLongestSubstring3b(s string) int {
	return 0
}

func TestLengthOfLongestSubstring3b(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty string", "", 0},
		{"single char", "a", 1},
		{"all same chars", "bbbbb", 1},
		{"no repeating chars", "abc", 3},
		{"repeating chars", "abcabcbb", 3},
		{"mixed pattern", "pwwkew", 3},
		{"with spaces", "ab c d", 4}, // Corrected from 5 to 4
		{"with digits", "ab123cd", 7},
		{"with symbols", "ab@#$cd", 7},
		{"complex case", "dvdf", 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := lengthOfLongestSubstring3b(tc.s)
			if got != tc.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d; want %d",
					tc.s, got, tc.want)
			}
		})
	}
}
