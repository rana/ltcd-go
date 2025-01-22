package main

import "testing"

// Time complexity:
// Space complexity:
func canConstruct383b(note string, mag string) bool {
	return false
}

func TestCanConstruct383b(t *testing.T) {
	tests := []struct {
		name string
		note string
		mag  string
		want bool
	}{
		{
			name: "Example 1: Single letter mismatch",
			note: "a",
			mag:  "b",
			want: false,
		},
		{
			name: "Example 2: Insufficient letters",
			note: "aa",
			mag:  "ab",
			want: false,
		},
		{
			name: "Example 3: Sufficient letters",
			note: "aa",
			mag:  "aab",
			want: true,
		},
		{
			name: "Empty note",
			note: "",
			mag:  "abc",
			want: true,
		},
		{
			name: "Same strings",
			note: "abc",
			mag:  "abc",
			want: true,
		},
		{
			name: "Multiple occurrences",
			note: "aabbcc",
			mag:  "aabbccdd",
			want: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := canConstruct383b(tc.note, tc.mag)
			if got != tc.want {
				t.Errorf("canConstruct(%q, %q) = %v, want %v",
					tc.note, tc.mag, got, tc.want)
			}
		})
	}
}
