package main

import (
	"testing"
)

func longestCommonPrefix14b(strs []string) string {
	return ""
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
