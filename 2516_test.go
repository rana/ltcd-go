package main

import "testing"

// Time complexity: O(n), n is length of string s.
// Space complexity: O(1), constant additional space used.
// https://console.anthropic.com/workbench/70406951-a2cf-47b0-8510-b8929c90956b
func takeCharacters2516(s string, k int) int {
	// Take K of Each Character From Left and Right
	// Given a string s composed of a b c.
	// Given a positive integer k.
	// Select a left or right character each minute.
	// Determine min minutes to select k of a b c.
	// Return min minutes.
	// Reframe to what to leave behind from middle.
	// Use two-pointer technique.
	// Use sliding window.
	// Frequency count chars in window and total.

	// Early exit if we don't need any characters
	if k == 0 {
		return 0
	}

	// Count total characters
	totCnt := [3]int{}
	for _, c := range s {
		totCnt[c-'a']++
	}

	// Check if it's possible to have k of each
	for i := range totCnt {
		if totCnt[i] < k {
			return -1
		}
	}

	// The problem can be reformulated:
	// Find the longest middle subarray where we don't exceed
	// the maximum allowed count for each character: (total - k)
	maxLen, lft := 0, 0
	winCnt := [3]int{} // Counts in the window
	sLen := len(s)

	for rht := range sLen {
		// Add current character to window
		winCnt[s[rht]-'a']++

		// Shrink from left if we exceed allowed count for any character
		for lft <= rht {
			valid := true
			for i := range totCnt {
				if winCnt[i] > totCnt[i]-k {
					valid = false
					break
				}
			}

			if valid {
				break
			}

			winCnt[s[lft]-'a']--
			lft++
		}

		// Update max length
		maxLen = max(maxLen, rht-lft+1)
	}

	// Total characters to take = total length - longest middle
	return sLen - maxLen
}

func TestTakeCharacters2516(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		// Examples from the problem description
		{
			name:     "Example 1",
			s:        "aabaaaacaabc",
			k:        2,
			expected: 8,
		},
		{
			name:     "Example 2",
			s:        "a",
			k:        1,
			expected: -1,
		},

		// Base cases
		{
			name:     "No characters needed",
			s:        "abc",
			k:        0,
			expected: 0,
		},
		{
			name:     "Single character of each type",
			s:        "abc",
			k:        1,
			expected: 3,
		},
		{
			name:     "Need all characters",
			s:        "aabbcc",
			k:        2,
			expected: 6,
		},

		// Edge cases
		{
			name:     "Empty string with k=0",
			s:        "",
			k:        0,
			expected: 0,
		},
		{
			name:     "Empty string with k>0",
			s:        "",
			k:        1,
			expected: -1,
		},
		{
			name:     "Long string with small k",
			s:        "abcabcabcabcabcabcabcabcabc",
			k:        1,
			expected: 3,
		},

		// Character distribution patterns
		{
			name:     "One character missing",
			s:        "aabb",
			k:        1,
			expected: -1, // No 'c' characters
		},
		{
			name:     "Uneven distribution - more on one side",
			s:        "aaabbbcca",
			k:        2,
			expected: 6,
		},

		// Special cases
		{
			name:     "All same character",
			s:        "aaaaaaaa",
			k:        4,
			expected: -1, // Can't get 4 of 'b' or 'c'
		},
		{
			name:     "Exactly k of each character",
			s:        "aabbcc",
			k:        2,
			expected: 6,
		},
		{
			name:     "Mixed distribution requiring both sides",
			s:        "abcabacbc",
			k:        2,
			expected: 6,
		},
		{
			name:     "Multiple optimal solutions",
			s:        "abcabcabc",
			k:        1,
			expected: 3, // Could take "abc" from either end
		},

		// Large k values
		{
			name:     "Large k relative to string",
			s:        "abcabcabc",
			k:        3,
			expected: 9, // Need entire string
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := takeCharacters2516(tc.s, tc.k)
			if got != tc.expected {
				t.Errorf("takeCharacters(%q, %d) = %d; want %d",
					tc.s, tc.k, got, tc.expected)
			}
		})
	}
}
