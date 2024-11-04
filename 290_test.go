package main

import (
	"strings"
	"testing"
)

// Time complexity: O(n), n is the length of string s. We traverse pat and s once.
// Space complexity: O(n), we place each word in two maps.
// https://claude.ai/chat/6cded66a-0574-4637-87d7-8634153d0357
func wordPattern290(pat string, s string) bool {
	// Word Pattern
	// Given two string pat and s.
	// pat is a character pattern.
	// s is words seprated by a space.
	// Determine whether there is a bijective one-to-one mapping
	// between each character and each word.
	// Return true if condition met.
	// Conditions:
	//	* Same character and word exclusively map
	//	* chr -> wrd == wrd -> chr
	// Use two maps to solve.

	// Split s into words.
	wrds := strings.Split(s, " ")

	// Check input lengths match.
	if len(pat) != len(wrds) {
		return false
	}

	// Initialize two maps.
	chrToWrd := make(map[byte]string)
	wrdToChr := make(map[string]byte)

	// Iterate through each character and word.
	for idx := 0; idx < len(pat); idx++ {
		// Get the character and word.
		chr := pat[idx]
		wrd := wrds[idx]

		// Check chr -> wrd mapping.
		if mapWrd, exists := chrToWrd[chr]; exists {
			if mapWrd != wrd {
				return false
			}
		} else {
			// Add chr -> wrd mapping.
			chrToWrd[chr] = wrd
		}

		// Check wrd -> chr mapping.
		if mapChr, exists := wrdToChr[wrd]; exists {
			if mapChr != chr {
				return false
			}
		} else {
			// Add wrd -> chr mapping.
			wrdToChr[wrd] = chr
		}
	}

	return true
}

// Test cases
func TestWordPattern290(t *testing.T) {
	tests := []struct {
		name string
		pat  string
		s    string
		want bool
	}{
		{
			name: "Example 1: Valid pattern abba",
			pat:  "abba",
			s:    "dog cat cat dog",
			want: true,
		},
		{
			name: "Example 2: Invalid pattern abba",
			pat:  "abba",
			s:    "dog cat cat fish",
			want: false,
		},
		{
			name: "Example 3: Invalid pattern aaaa",
			pat:  "aaaa",
			s:    "dog cat cat dog",
			want: false,
		},
		{
			name: "Single char pattern",
			pat:  "a",
			s:    "dog",
			want: true,
		},
		{
			name: "Length mismatch",
			pat:  "abc",
			s:    "dog cat",
			want: false,
		},
		{
			name: "Same word different patterns",
			pat:  "ab",
			s:    "dog dog",
			want: false,
		},
		{
			name: "Different words same pattern",
			pat:  "aa",
			s:    "dog cat",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := wordPattern290(tc.pat, tc.s)
			if got != tc.want {
				t.Errorf("wordPattern(%q, %q) = %v, want %v",
					tc.pat, tc.s, got, tc.want)
			}
		})
	}
}
