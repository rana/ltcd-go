package main

import "testing"

// Time complexity: O(n), n is the length of string hay. We traverse the string once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/3fcaeb96-7cff-48da-b438-946c0c5379f1
func minWindow(hay, ndl string) string {
	// Minimum Window Substring
	// Given two strings hay and ndl.
	// Find ndl in hay as a minimum substring.
	// Return the substring; or, empty string.
	// Conditions:
	//	* ndl may have duplicates
	//	* Uppercase and lowecase English letters only.
	// Use a frequency map counter for each hay and ndl.
	// Use a sliding window two-pointer technique.

	// Check for input minimum edge cases.
	if len(hay) < len(ndl) || len(ndl) == 0 {
		return ""
	}

	// Initialize frequency maps.
	var ndlFrq [128]int
	var curFrq [128]int

	// Fill the needle frequency map.
	for _, chr := range ndl {
		ndlFrq[chr]++
	}

	// Initialize sliding window variables.
	lft, cnt := 0, 0
	minLen := len(hay) + 1
	minStart := 0

	// Expand window on right side.
	for rht := 0; rht < len(hay); rht++ {
		// Get the current character.
		chr := hay[rht]
		// Increment the current character count.
		curFrq[chr]++

		// Check to increment window count.
		// Support skipping duplicates.
		if curFrq[chr] <= ndlFrq[chr] {
			cnt++
		}

		// Contract window while valid.
		for cnt == len(ndl) {
			// Update minimum window.
			curLen := rht - lft + 1
			if curLen < minLen {
				minLen = curLen
				minStart = lft
			}

			// Contract window.
			chr = hay[lft]
			curFrq[chr]--

			// Check whether required character removed.
			if curFrq[chr] < ndlFrq[chr] {
				cnt--
			}

			lft++
		}
	}

	// Check for no match condition.
	if minLen > len(hay) {
		return ""
	}

	// Return minimum substring.
	return hay[minStart : minStart+minLen]
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		hay  string
		ndl  string
		want string
	}{
		{
			name: "Example 1: Normal case",
			hay:  "ADOBECODEBANC",
			ndl:  "ABC",
			want: "BANC",
		},
		{
			name: "Example 2: Single character",
			hay:  "a",
			ndl:  "a",
			want: "a",
		},
		{
			name: "Example 3: Impossible case",
			hay:  "a",
			ndl:  "aa",
			want: "",
		},
		{
			name: "Empty strings",
			hay:  "",
			ndl:  "",
			want: "",
		},
		{
			name: "No valid window",
			hay:  "xyz",
			ndl:  "abc",
			want: "",
		},
		{
			name: "Multiple valid windows",
			hay:  "ADOBECODEBANCA",
			ndl:  "ABC",
			want: "BANC",
		},
		{
			name: "Case sensitivity",
			hay:  "aaAAAA",
			ndl:  "AA",
			want: "AA",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := minWindow(tc.hay, tc.ndl)
			if got != tc.want {
				t.Errorf("minWindow(%q, %q) = %q; want %q",
					tc.hay, tc.ndl, got, tc.want)
			}
		})
	}
}
