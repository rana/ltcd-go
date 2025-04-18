package main

import "testing"

// Time complexity: O(m * n), n is the length of string `hay`. m is the length of string `ned`.
// Space complexity: O(1), constant additional space used with several variables.
// https://claude.ai/chat/4ec0998a-2ef9-41f8-960d-af1a0b090d1b
func strStr28(hay string, ned string) int {
	// Find the Index of the First Occurrence in a String
	// Given two strings `hay` and `ned`.
	// Determine the first index of `ned` in `hay`.
	// Return the index; or, -1.
	// Use a comparison-based two-pointer technique.

	hayLen, nedLen := len(hay), len(ned)
	if hayLen < nedLen {
		return -1
	}
	hayIdx, nedIdx := 0, 0

	// Iterate through hay.
	for hayIdx <= hayLen-nedLen {
		nedIdx = 0
		// Iterate through ned.
		for nedIdx < nedLen && hay[hayIdx+nedIdx] == ned[nedIdx] {
			nedIdx++
		}
		// Check whole word match
		if nedIdx == nedLen {
			return hayIdx
		}
		hayIdx++
	}
	return -1
}

func TestStrStr28(t *testing.T) {
	// Test case 1: Normal case with match at beginning
	if res := strStr28("sadbutsad", "sad"); res != 0 {
		t.Errorf("Test case 1 failed: expected 0, got %d", res)
	}

	// Test case 2: No match case
	if res := strStr28("leetcode", "leeto"); res != -1 {
		t.Errorf("Test case 2 failed: expected -1, got %d", res)
	}

	// Test case 3: Match at end
	if res := strStr28("hello", "llo"); res != 2 {
		t.Errorf("Test case 3 failed: expected 2, got %d", res)
	}

	// Test case 4: Single character strings
	if res := strStr28("a", "a"); res != 0 {
		t.Errorf("Test case 4 failed: expected 0, got %d", res)
	}

	// Test case 5: Needle longer than haystack
	if res := strStr28("hi", "hello"); res != -1 {
		t.Errorf("Test case 5 failed: expected -1, got %d", res)
	}

	// Test case 6: Multiple occurrences
	if res := strStr28("mississippi", "issi"); res != 1 {
		t.Errorf("Test case 6 failed: expected 1, got %d", res)
	}
}
