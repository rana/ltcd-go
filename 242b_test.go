package main

import "testing"

// Time complexity: O(n), n is the length of both strings. We traverse the length once.
// Space complexity: O(1), constant additional space used. We use a 26 length array for frequency counting.
func isAnagram(s, t string) bool {
	// Valid Anagram
	// Given two strings s and t.
	// Determine whether s is an anagram of t.
	// Return true if condition met.
	// Condition:
	//	* s and t are lowercase English letters.
	// Key insight is that an anagram has identical characters
	// in each string, and an identical frequency of characters
	// in each string.
	// Use an array for the map due to the small character set.

	// Check input lengths match.
	if len(s) != len(t) {
		return false
	}

	// Initialize a character frequency counter.
	var frq [26]int

	// Iterate through both strings simultaneously.
	// Increment and decrement count.
	for idx := 0; idx < len(s); idx++ {
		frq[s[idx]-'a']++
		frq[t[idx]-'a']--
	}

	// Validate all character counts are zero.
	for idx := 0; idx < len(frq); idx++ {
		if frq[idx] != 0 {
			return false
		}
	}

	return true
}

func TestIsAnagram_ValidCases(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "example 1 - simple anagram",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "example 2 - not anagram",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "same strings",
			s:    "hello",
			t:    "hello",
			want: true,
		},
		{
			name: "empty strings",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "different lengths",
			s:    "ab",
			t:    "abc",
			want: false,
		},
		{
			name: "same letters different counts",
			s:    "aaab",
			t:    "aaba",
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isAnagram(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("isAnagram(%q, %q) = %v, want %v",
					tc.s, tc.t, got, tc.want)
			}
		})
	}
}
