package main

import "testing"

func strStr28b(hay string, ned string) int {
	return 0
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
