package main

import "testing"

// Time complexity: O(n), n is the length of string s. We iterate through it once.
// Space complexity: O(min(m, n) m is the character set size. n is the string length.
// https://claude.ai/chat/a04662a3-d932-4a0f-81d4-4a778fffbe23
func lengthOfLongestSubstring3(s string) int {
	// Longest Substring Without Repeating Characters
	// Given a string s.
	// Find the maximum length of a substring with condition.
	// Condition:
	//	* No duplicate characters in substring.
	// Return the max length.
	// Use a two-pointer technique.
	// Use a map to track character last seen position.

	// Initialize variables.
	maxLen := 0
	lft := 0
	chrs := make(map[byte]int)

	// Iterate through right pointer of each character.
	for rht := 0; rht < len(s); rht++ {
		// Get previous matching character.
		prv, ok := chrs[s[rht]]
		// Check whether character was seen
		// and in the window.
		if ok && prv >= lft {
			// Update the left pointer beyond prvevious.
			// Current and previous characters are
			// duplicates in window.
			lft = prv + 1
		}

		// Store the position of the current character.
		chrs[s[rht]] = rht

		// Calculate current substring length.
		curLen := rht - lft + 1

		// Determine max length.
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	// Return max length.
	return maxLen
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"empty string", "", 0},
		{"single char", "a", 1},
		{"all same chars", "bbbbb", 1},
		{"no repeating chars", "abc", 3},
		{"repeating chars", "abcabcbb", 3},
		{"mixed pattern", "pwwkew", 3},
		{"with spaces", "ab c d", 4}, // Corrected from 5 to 4
		{"with digits", "ab123cd", 7},
		{"with symbols", "ab@#$cd", 7},
		{"complex case", "dvdf", 3},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := lengthOfLongestSubstring3(tc.s)
			if got != tc.want {
				t.Errorf("lengthOfLongestSubstring(%q) = %d; want %d",
					tc.s, got, tc.want)
			}
		})
	}
}
