package main

import "testing"

// Time complexity: O(n), n is the length of both strings. We iterate through each string once in a single loop.
// Space complexity: O(1), constant additional space used. We use two maps of characters. The character set is fixed-size 256 characters.
func isIsomorphic(s string, t string) bool {
	// Isomorphic Strings
	// Given two string s and t.
	// Determine whether the strings are isomorphic.
	// Return true if condition met.
	// Condition:
	//	* s and t are ascii characres.
	// Isomorphic is a bi-directional one-to-one mapping.
	// chrS -> chrT == chrT -> chrS
	// Use two maps to determine.

	// Check input lengths are equal.
	if len(s) != len(t) {
		return false
	}

	// Create maps.
	sTOt := make(map[byte]byte)
	tTOs := make(map[byte]byte)

	// Iterate through each string once.
	for idx := 0; idx < len(s); idx++ {
		// Get each character.
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

func TestIsIsomorphic(t *testing.T) {
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
			if got := isIsomorphic(tc.str, tc.txt); got != tc.want {
				t.Errorf("isIsomorphic(%q, %q) = %v, want %v",
					tc.str, tc.txt, got, tc.want)
			}
		})
	}
}
