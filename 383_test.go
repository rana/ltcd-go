package main

import "testing"

// Time complexity: O(m + n), m is the length of magazine. n is the length of note. We iterate through each string once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/25b2204e-ba1d-4a47-a8f1-0a1196b3ac98
func canConstruct383(note string, mag string) bool {
	// Ransom Note
	// Given two string `note` and `mag`.
	// Determine whether letters from mag can construct note.
	// Conditions:
	//	* Only lowercase English letters are used.
	// Use a frequency counter of letters.

	// Create a frequency map for magazine characters.
	// Increment counter
	var frq [26]int
	for _, chr := range mag {
		frq[chr-'a']++
	}

	// Check frequency of characters in note.
	// Decrement counter.
	for _, chr := range note {
		frq[chr-'a']--
		if frq[chr-'a'] < 0 {
			return false
		}
	}

	return true
}

func TestCanConstruct383(t *testing.T) {
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
			got := canConstruct383(tc.note, tc.mag)
			if got != tc.want {
				t.Errorf("canConstruct(%q, %q) = %v, want %v",
					tc.note, tc.mag, got, tc.want)
			}
		})
	}
}
