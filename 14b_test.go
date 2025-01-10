package main

import (
	"testing"
)

func longestCommonPrefix14b(strs []string) string {
	// Longest Common Prefix
	// Given an array of strings strs.
	// Determine the longest common prefix in all strings.
	// Return the longest common prefix.

	// Check input min edge cases.
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	fstWrd := strs[0]
	// Iterate through all character of first word.
	for idx := 0; idx < len(fstWrd); idx++ {
		curChr := fstWrd[idx]
		// Iterate through remaining strings.
		for _, str := range strs {
			// Check exit conditions.
			if idx >= len(str) || curChr != str[idx] {
				return fstWrd[:idx]
			}
		}
	}
	return fstWrd
}

func TestLongestCommonPrefix14b(t *testing.T) {
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
			name:     "Empty array",
			input:    []string{},
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
			result := longestCommonPrefix14b(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
