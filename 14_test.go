package main

import (
	"testing"
)

// Time complexity: O(s), s is the sum of all characters in all strings.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/a529f1b8-aacf-4cc2-a47a-f9641fcf616f
func longestCommonPrefix14(strs []string) string {
	// Longest Common Prefix
	// Given a string array.
	// Determine the longest common prefix.
	// Return the longest common prefix; or, empty string.
	// Key insight is that the longest common prefix
	// is in all strings.

	// Check input minimum edge cases.
	if len(strs) == 1 {
		return strs[0]
	}

	// Initialize variable.
	fstWrd := strs[0]

	// Iterate through each character of the first word.
	for idx := range fstWrd {
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

func TestLongestCommonPrefix14(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Example 1",
			input:    []string{"flower", "flow", "flight"},
			expected: "fl",
		},
		{
			name:     "Example 2",
			input:    []string{"dog", "racecar", "car"},
			expected: "",
		},
		{
			name:     "Single string",
			input:    []string{"hello"},
			expected: "hello",
		},
		{
			name:     "Identical strings",
			input:    []string{"interstellar", "interstellar", "interstellar"},
			expected: "interstellar",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := longestCommonPrefix14(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
