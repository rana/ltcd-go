package main

import "testing"

// Time complexity: O(n), n is the length of the input string s an t. We taverse each string once.
// Space complexity: O(1), constant additional space used. We use 256 ascii characters in two maps.
// https://claude.ai/chat/78fcfb73-a4a0-4ed4-8746-b9b997bbdb29
func isIsomorphic205(s, t string) bool {
	// Isomorphic Strings
	// Given two string s and t.
	// Determine whether the two strings are isomorphic.
	// Return true if the condition is met.
	// Conditions:
	//	* One-to-one mapping between each character.
	//	* Bi-directional mapping is identical s -> t == t -> s.
	// Use two maps to solve.

	// Check input minimum edge case.
	if len(s) != len(t) {
		return false
	}

	// Initialize two maps.
	sTOt := make(map[byte]byte)
	tTOs := make(map[byte]byte)

	// Iterate through each character.
	for idx := 0; idx < len(s); idx++ {
		// Get characters from each string.
		chrS := s[idx]
		chrT := t[idx]

		// Check s -> t mapping.
		if mapChrT, exists := sTOt[chrS]; exists {
			if mapChrT != chrT {
				return false
			}
		} else {
			// Add s -> t mapping.
			sTOt[chrS] = chrT
		}

		// Check t -> s mapping.
		if mapChrS, exists := tTOs[chrT]; exists {
			if mapChrS != chrS {
				return false
			}
		} else {
			// Add t -> s mapping.
			tTOs[chrT] = chrS
		}
	}

	return true
}

func TestIsIsomorphic205(t *testing.T) {
	tests := []struct {
		name string
		str  string
		txt  string
		want bool
	}{
		{
			name: "Example 1: Valid mapping",
			str:  "egg",
			txt:  "add",
			want: true,
		},
		{
			name: "Example 2: Invalid mapping",
			str:  "foo",
			txt:  "bar",
			want: false,
		},
		{
			name: "Example 3: Paper-Title",
			str:  "paper",
			txt:  "title",
			want: true,
		},
		{
			name: "Empty strings",
			str:  "",
			txt:  "",
			want: true,
		},
		{
			name: "Single character",
			str:  "a",
			txt:  "b",
			want: true,
		},
		{
			name: "Different lengths",
			str:  "abc",
			txt:  "ab",
			want: false,
		},
		{
			name: "Same characters different pattern",
			str:  "badc",
			txt:  "baba",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := isIsomorphic205(tc.str, tc.txt); got != tc.want {
				t.Errorf("isIsomorphic(%q, %q) = %v, want %v",
					tc.str, tc.txt, got, tc.want)
			}
		})
	}
}
