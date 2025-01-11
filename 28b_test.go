package main

import "testing"

// Time complexity: O(m * n), m is the length of string hay. n is the length of string ned.
// Space complexity: O(1), constant additional space used.
func strStr28b(hay string, ned string) int {
	// Find the Index of the First Occurrence in a String
	// Given strings hay and ned.
	// Determine whether ned is in hay.
	// Return the first index of ned; or, -1.
	// Use a two-pointer technique.
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

func TestStrStr28b(t *testing.T) {
	// Test case 1: Normal case with match at beginning
	if res := strStr28b("sadbutsad", "sad"); res != 0 {
		t.Errorf("Test case 1 failed: expected 0, got %d", res)
	}

	// Test case 2: No match case
	if res := strStr28b("leetcode", "leeto"); res != -1 {
		t.Errorf("Test case 2 failed: expected -1, got %d", res)
	}

	// Test case 3: Match at end
	if res := strStr28b("hello", "llo"); res != 2 {
		t.Errorf("Test case 3 failed: expected 2, got %d", res)
	}

	// Test case 4: Single character strings
	if res := strStr28b("a", "a"); res != 0 {
		t.Errorf("Test case 4 failed: expected 0, got %d", res)
	}

	// Test case 5: Needle longer than haystack
	if res := strStr28b("hi", "hello"); res != -1 {
		t.Errorf("Test case 5 failed: expected -1, got %d", res)
	}

	// Test case 6: Multiple occurrences
	if res := strStr28b("mississippi", "issi"); res != 1 {
		t.Errorf("Test case 6 failed: expected 1, got %d", res)
	}
}
