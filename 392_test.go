package main

import "testing"

// Time complexity: O(n), n is the length of string t. We traverse string t once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/cf1db7a2-fa64-4594-8c14-a765d7cdff80
func isSubsequence392(s string, t string) bool {
	// Is Subsequence
	// Given strings s and t made of lowercase English letters.
	// Determine whether string s is a subsequence of string t.
	// Key insight is that s and t have the same characters
	// in the same order.
	// Use a two-pointer technique.

	// Initialize string lengths.
	sLen, tLen := len(s), len(t)

	// Check input minimum edge cases.
	if sLen == 0 {
		return true
	}
	if tLen == 0 {
		return false
	}

	// Initialize two indexes.
	sIdx, tIdx := 0, 0

	// Iterate through all the characters in t.
	for sIdx < sLen && tIdx < tLen {

		// Check for character match in s and t.
		if s[sIdx] == t[tIdx] {
			// Increment s index.
			sIdx++
		}

		// Increment t index.
		tIdx++
	}

	// Check for whole subsequence match.
	return sIdx == sLen
}

func TestIsSubsequence392(t *testing.T) {
	// Test case 1: Basic valid subsequence
	if res := isSubsequence392("abc", "ahbgdc"); !res {
		t.Errorf("Expected true for s='abc', t='ahbgdc', got false")
	}

	// Test case 2: Invalid subsequence
	if res := isSubsequence392("axc", "ahbgdc"); res {
		t.Errorf("Expected false for s='axc', t='ahbgdc', got true")
	}

	// Test case 3: Empty string s
	if res := isSubsequence392("", "ahbgdc"); !res {
		t.Errorf("Expected true for empty string s, got false")
	}

	// Test case 4: Empty string t
	if res := isSubsequence392("abc", ""); res {
		t.Errorf("Expected false for empty string t, got true")
	}

	// Test case 5: Single character match
	if res := isSubsequence392("a", "a"); !res {
		t.Errorf("Expected true for s='a', t='a', got false")
	}

	// Test case 6: Longer valid subsequence
	if res := isSubsequence392("ace", "abcde"); !res {
		t.Errorf("Expected true for s='ace', t='abcde', got false")
	}
}
