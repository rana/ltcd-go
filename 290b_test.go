package main

import (
	"testing"
)

// Time complexity:
// Space complexity:
func wordPattern290b(pat string, s string) bool {
	return false
}

func TestWordPattern290b(t *testing.T) {
	tests := []struct {
		name string
		pat  string
		s    string
		want bool
	}{
		{
			name: "Example 1: Valid pattern abba",
			pat:  "abba",
			s:    "dog cat cat dog",
			want: true,
		},
		{
			name: "Example 2: Invalid pattern abba",
			pat:  "abba",
			s:    "dog cat cat fish",
			want: false,
		},
		{
			name: "Example 3: Invalid pattern aaaa",
			pat:  "aaaa",
			s:    "dog cat cat dog",
			want: false,
		},
		{
			name: "Single char pattern",
			pat:  "a",
			s:    "dog",
			want: true,
		},
		{
			name: "Length mismatch",
			pat:  "abc",
			s:    "dog cat",
			want: false,
		},
		{
			name: "Same word different patterns",
			pat:  "ab",
			s:    "dog dog",
			want: false,
		},
		{
			name: "Different words same pattern",
			pat:  "aa",
			s:    "dog cat",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := wordPattern290b(tc.pat, tc.s)
			if got != tc.want {
				t.Errorf("wordPattern(%q, %q) = %v, want %v",
					tc.pat, tc.s, got, tc.want)
			}
		})
	}
}
