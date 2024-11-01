package main

import "testing"

// Time complexity: O(n * m * k), n is the length of string s. m is the length of a word (all equal length). k is the number of words.
// Space complexity: O(k), we store a frequency map of words.
// https://claude.ai/chat/4c9842a1-9e15-4702-bc09-bb8201a504f2
func findSubstring(s string, wrds []string) []int {
	// Substring with Concatenation of All Words
	// Given a string `s`.
	// Given a string array `wrds`.
	// Find concatenated permutations of `wrds` in string `s`.
	// Return an array of starting indexes.
	// Notice words have the same length.
	// Use a word frequency counter map.
	// Use a sliding window approach.
	// The sliding window is the length of a whole permutation.

	// Check input minimum edge cases.
	if len(s) == 0 || len(wrds) == 0 {
		return []int{}
	}

	// Initialize variables.
	var res []int
	wrdLen := len(wrds[0])
	totLen := wrdLen * len(wrds)
	strLen := len(s)

	// Create frequency counter map for words.
	wrdMap := make(map[string]int)
	for _, wrd := range wrds {
		wrdMap[wrd]++
	}

	// Iterate through possible starting indexes.
	for idx := 0; idx <= strLen-totLen; idx++ {
		// Copy the word map for decrementing.
		curMap := make(map[string]int)
		for key, val := range wrdMap {
			curMap[key] = val
		}

		allWrdsMatch := true

		// Iterate through all word positions.
		for pos := 0; pos < len(wrds); pos++ {
			// Get the current word in string s.
			curWrd := s[idx+pos*wrdLen : idx+(pos+1)*wrdLen]

			// Get current word frequency count.
			cnt, exists := curMap[curWrd]

			// Check failure conditions.
			if !exists || cnt == 0 {
				allWrdsMatch = false
				break
			}

			// Decrement word count.
			curMap[curWrd]--
		}

		// Check for whole permutation match.
		if allWrdsMatch {
			// Sotre starting index.
			res = append(res, idx)
		}
	}

	// Return starting indexes
	return res
}

// Test cases
func TestFindSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		words    []string
		expected []int
	}{
		{
			name:     "Example 1",
			s:        "barfoothefoobarman",
			words:    []string{"foo", "bar"},
			expected: []int{0, 9},
		},
		{
			name:     "Example 2",
			s:        "wordgoodgoodgoodbestword",
			words:    []string{"word", "good", "best", "word"},
			expected: []int{},
		},
		{
			name:     "Example 3",
			s:        "barfoofoobarthefoobarman",
			words:    []string{"bar", "foo", "the"},
			expected: []int{6, 9, 12},
		},
		{
			name:     "Empty string",
			s:        "",
			words:    []string{"foo"},
			expected: []int{},
		},
		{
			name:     "Empty words",
			s:        "foo",
			words:    []string{},
			expected: []int{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findSubstring(tc.s, tc.words)
			if !equalSlices(got, tc.expected) {
				t.Errorf("findSubstring(%v, %v) = %v; want %v",
					tc.s, tc.words, got, tc.expected)
			}
		})
	}
}

// Helper function to compare slices
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
