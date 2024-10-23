package main

import "testing"

// Time complexity: O(m * n), n is the length of string `hay`. m is the length of string `ned`.
// Space complexity: O(1), constant additional space used with several variables.
// https://claude.ai/chat/4ec0998a-2ef9-41f8-960d-af1a0b090d1b
func strStr(hay string, ned string) int {
	// Find the Index of the First Occurrence in a String
	// Given two strings `hay` and `ned`.
	// Determine the first index of `ned` in `hay`.
	// Return the index; or, -1.
	// Use a comparison-based two-pointer technique.

	// Initialize length variables.
	hayLen, nedLen := len(hay), len(ned)

	// Check input minimum edge case.
	if hayLen < nedLen {
		return -1
	}

	// Initialize two-pointer variables.
	lft, rht := 0, 0

	// Iterate through each character in hay (up to nedLen).
	for lft <= hayLen-nedLen {
		// Reset rht pointer for new hay comparison state.
		rht = 0

		// Iterate through each character in ned for all matches.
		for rht < nedLen && hay[lft+rht] == ned[rht] {
			rht++
		}

		// Check for whole needle match.
		if rht == nedLen {
			return lft
		}

		// Move the left pointer right.
		lft++
	}

	// No needle match.
	return -1
}

func TestStrStr(t *testing.T) {
	// Test case 1: Normal case with match at beginning
	if res := strStr("sadbutsad", "sad"); res != 0 {
		t.Errorf("Test case 1 failed: expected 0, got %d", res)
	}

	// Test case 2: No match case
	if res := strStr("leetcode", "leeto"); res != -1 {
		t.Errorf("Test case 2 failed: expected -1, got %d", res)
	}

	// Test case 3: Match at end
	if res := strStr("hello", "llo"); res != 2 {
		t.Errorf("Test case 3 failed: expected 2, got %d", res)
	}

	// Test case 4: Single character strings
	if res := strStr("a", "a"); res != 0 {
		t.Errorf("Test case 4 failed: expected 0, got %d", res)
	}

	// Test case 5: Needle longer than haystack
	if res := strStr("hi", "hello"); res != -1 {
		t.Errorf("Test case 5 failed: expected -1, got %d", res)
	}

	// Test case 6: Multiple occurrences
	if res := strStr("mississippi", "issi"); res != 1 {
		t.Errorf("Test case 6 failed: expected 1, got %d", res)
	}
}
