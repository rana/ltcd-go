package main

import (
	"testing"
)

// Time complexity: O(s), s is the sum of all characters in all strings.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/a529f1b8-aacf-4cc2-a47a-f9641fcf616f
func longestCommonPrefix(strs []string) string {
	// Longest Common Prefix
	// Given a string array.
	// Determine the longest common prefix.
	// Return the longest common prefix; or, empty string.
	// Key insight is that the longest common prefix
	// is in all strings.

	// Check input minimum edge cases.
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	// Initialize variable.
	fstWrd := strs[0]

	// Iterate through each character of the first word.
	for idx := 0; idx < len(fstWrd); idx++ {
		// Get the first character of the first word.
		curChr := fstWrd[idx]

		// Compare character to position in other words.
		for _, str := range strs[1:] {
			// Check whether comparison string too long,
			// or for character mismatch.
			if idx >= len(str) || curChr != str[idx] {
				// Return first word up to index as prefix.
				return fstWrd[:idx]
			}
		}
	}

	// First word is prefix here.
	return fstWrd
}

// Test cases
func TestLongestCommonPrefix(t *testing.T) {
	// Test case 1: Example from problem statement
	t.Run("Example 1", func(t *testing.T) {
		strs := []string{"flower", "flow", "flight"}
		exp := "fl"
		res := longestCommonPrefix(strs)
		if res != exp {
			t.Errorf("Expected %s, got %s", exp, res)
		}
	})

	// Test case 2: No common prefix
	t.Run("Example 2", func(t *testing.T) {
		strs := []string{"dog", "racecar", "car"}
		exp := ""
		res := longestCommonPrefix(strs)
		if res != exp {
			t.Errorf("Expected %s, got %s", exp, res)
		}
	})

	// Test case 3: Empty array
	t.Run("Empty array", func(t *testing.T) {
		strs := []string{}
		exp := ""
		res := longestCommonPrefix(strs)
		if res != exp {
			t.Errorf("Expected %s, got %s", exp, res)
		}
	})

	// Test case 4: Single string
	t.Run("Single string", func(t *testing.T) {
		strs := []string{"hello"}
		exp := "hello"
		res := longestCommonPrefix(strs)
		if res != exp {
			t.Errorf("Expected %s, got %s", exp, res)
		}
	})

	// Test case 5: All identical strings
	t.Run("Identical strings", func(t *testing.T) {
		strs := []string{"interstellar", "interstellar", "interstellar"}
		exp := "interstellar"
		res := longestCommonPrefix(strs)
		if res != exp {
			t.Errorf("Expected %s, got %s", exp, res)
		}
	})
}
