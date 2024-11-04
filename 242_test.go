package main

import (
	"testing"
)

// Time complexity: O(n), n is the length of both strings. We traverse both strings once.
// Space complexity: O(1), constant additional space used. We use a fixed-size array for frequency counting.
// https://claude.ai/chat/e004b443-5692-4aad-a4b8-a9feb7b7e7ea
func isAnagram242(s, t string) bool {
	// Valid Anagram
	// Given two string s and t.
	// Determine if s is an anagram of t.
	// Return true if condition met.
	// Condition:
	//	* Identical characters used.
	//	* Identical frequency of characters used.
	//	* Lowercase English letter.
	// Use a frequency map based on a fixed-size array.

	// Check input strings have equal lengths.
	if len(s) != len(t) {
		return false
	}

	// Initialize a frequency array.
	var frq [26]int

	// Iterate through each string.
	// Increment and decrement.
	for idx := 0; idx < len(s); idx++ {
		frq[s[idx]-'a']++
		frq[t[idx]-'a']--
	}

	// Check for equal occurrences.
	for idx := 0; idx < len(frq); idx++ {
		if frq[idx] != 0 {
			return false
		}
	}

	return true
}

// TestIsAnagram_ValidCases tests various valid anagram scenarios
func TestIsAnagram242_ValidCases(t *testing.T) {
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
			got := isAnagram242(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("isAnagram(%q, %q) = %v, want %v",
					tc.s, tc.t, got, tc.want)
			}
		})
	}
}
