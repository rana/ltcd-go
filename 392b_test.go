package main

import "testing"

func isSubsequence392b(s string, t string) bool {
	// Is Subsequence
	// Given strings s and t.
	// Determine whether s is a subsequence in t.
	// Return true if condition met.
	// Use a two-pointer technique.
	sLen, tLen := len(s), len(t)
	if sLen == 0 {
		return true
	}
	if tLen == 0 {
		return false
	}
	sIdx, tIdx := 0, 0
	for sIdx < sLen && tIdx < tLen {
		if s[sIdx] == t[tIdx] {
			sIdx++
		}
		tIdx++
	}
	return sIdx == sLen
}

func TestIsSubsequence392b(t *testing.T) {
	// Test case 1: Basic valid subsequence
	if res := isSubsequence392b("abc", "ahbgdc"); !res {
		t.Errorf("Expected true for s='abc', t='ahbgdc', got false")
	}

	// Test case 2: Invalid subsequence
	if res := isSubsequence392b("axc", "ahbgdc"); res {
		t.Errorf("Expected false for s='axc', t='ahbgdc', got true")
	}

	// Test case 3: Empty string s
	if res := isSubsequence392b("", "ahbgdc"); !res {
		t.Errorf("Expected true for empty string s, got false")
	}

	// Test case 4: Empty string t
	if res := isSubsequence392b("abc", ""); res {
		t.Errorf("Expected false for empty string t, got true")
	}

	// Test case 5: Single character match
	if res := isSubsequence392b("a", "a"); !res {
		t.Errorf("Expected true for s='a', t='a', got false")
	}

	// Test case 6: Longer valid subsequence
	if res := isSubsequence392b("ace", "abcde"); !res {
		t.Errorf("Expected true for s='ace', t='abcde', got false")
	}
}
