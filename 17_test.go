package main

import (
	"reflect"
	"sort"
	"testing"
)

// Time complexity: O(4^n * n), n is the length of the digit string.
// Space complexity: O(n), for the recursion stack.
// https://claude.ai/chat/ff66dd71-d8e1-439b-a04e-dac96f897072
func letterCombinations17(digs string) []string {
	// Letter Combinations of a Phone Number
	// Given a string of digits.
	// A digit can map to three or four letters.
	// Create all letter combinations from a digit string.
	// Return all letter combinations.
	// Conditions:
	// 	* Digits are 2-9
	//	* Digit maps to three or four letters.
	// Use recursive backtracking.

	// Check input min edge case.
	if len(digs) == 0 {
		return []string{}
	}

	// Create a digit to letters mapping.
	numMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	// Initialize result.
	var res []string

	// Create a closure for recursive backtracking
	// combination building.
	var bckTrk func(pos int, cur string)
	bckTrk = func(pos int, cur string) {
		// Check base case: completed combination building
		if pos == len(digs) {
			res = append(res, cur)
			return
		}

		// Get digit and related letters.
		dig := digs[pos]
		ltrs := numMap[dig]

		// Build combinations with letters recursively.
		for _, ltr := range ltrs {
			bckTrk(pos+1, cur+string(ltr))
		}
	}

	// Build combinations from empty string.
	bckTrk(0, "")

	return res
}

// Tests
func TestLetterCombinations17(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:     "Example 1 - two digits",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "Example 2 - empty string",
			digits:   "",
			expected: []string{},
		},
		{
			name:     "Example 3 - single digit",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
		{
			name:   "Four digits",
			digits: "2345",
			expected: []string{
				"adgj", "adgk", "adgl", "adhj", "adhk", "adhl", "adij", "adik", "adil",
				"aegj", "aegk", "aegl", "aehj", "aehk", "aehl", "aeij", "aeik", "aeil",
				"afgj", "afgk", "afgl", "afhj", "afhk", "afhl", "afij", "afik", "afil",
				"bdgj", "bdgk", "bdgl", "bdhj", "bdhk", "bdhl", "bdij", "bdik", "bdil",
				"begj", "begk", "begl", "behj", "behk", "behl", "beij", "beik", "beil",
				"bfgj", "bfgk", "bfgl", "bfhj", "bfhk", "bfhl", "bfij", "bfik", "bfil",
				"cdgj", "cdgk", "cdgl", "cdhj", "cdhk", "cdhl", "cdij", "cdik", "cdil",
				"cegj", "cegk", "cegl", "cehj", "cehk", "cehl", "ceij", "ceik", "ceil",
				"cfgj", "cfgk", "cfgl", "cfhj", "cfhk", "cfhl", "cfij", "cfik", "cfil",
			},
		},
		{
			name:     "Digits with 4 letters (7)",
			digits:   "7",
			expected: []string{"p", "q", "r", "s"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := letterCombinations17(tc.digits)
			// Sort both slices for comparison since order doesn't matter
			sort.Strings(got)
			sort.Strings(tc.expected)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
