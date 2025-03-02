package main

import (
	"testing"
)

// Time complexity:
// Space complexity:
func addBoldTag616b(s string, wrds []string) string {
	return ""
}

func TestAddBoldTag616b(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		words    []string
		expected string
	}{
		{
			name:     "Example 1",
			s:        "abcxyz123",
			words:    []string{"abc", "123"},
			expected: "<b>abc</b>xyz<b>123</b>",
		},
		{
			name:     "Example 2",
			s:        "aaabbb",
			words:    []string{"aa", "b"},
			expected: "<b>aaabbb</b>",
		},
		{
			name:     "Empty string",
			s:        "",
			words:    []string{"abc"},
			expected: "",
		},
		{
			name:     "No words",
			s:        "hello world",
			words:    []string{},
			expected: "hello world",
		},
		{
			name:     "Non-overlapping words",
			s:        "abc def ghi",
			words:    []string{"abc", "ghi"},
			expected: "<b>abc</b> def <b>ghi</b>",
		},
		{
			name:     "Word not in string",
			s:        "hello",
			words:    []string{"world"},
			expected: "hello",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := addBoldTag616b(tc.s, tc.words)
			if got != tc.expected {
				t.Errorf("addBoldTag(%q, %v) = %q, want %q",
					tc.s, tc.words, got, tc.expected)
			}
		})
	}
}
