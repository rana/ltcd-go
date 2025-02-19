package main

import "testing"

// Time Complexity: O(n * m) where n is length of source string, m is length of pattern
// Space Complexity: O(1)
// https://console.anthropic.com/workbench/f83050e1-437e-46a2-9efb-1c0e1cf5a505
func solutionGCA2b(pat string, src string) int {
	// Data Manipulation
	// Given string pat and string src.
	// Determine quantity of substring in src matching pattern of 0 1 in pat.
	// Return quantity.
	// Conditions:
	// * Length equal for pat and substring
	// * 0 maps to a vowel
	// * 1 maps to a consonant
	// * Vowels: a e i o u y
	// Use a sliding window approach.
	isVwl := func(chr byte) bool {
		switch chr {
		case 'a', 'e', 'i', 'o', 'u', 'y':
			return true
		}
		return false
	}
	mtcPat := func(sub, pat string) bool {
		for idx := range sub {
			isVow := isVwl(sub[idx])
			if pat[idx] == '0' && !isVow || pat[idx] == '1' && isVow {
				return false
			}
		}
		return true
	}

	cnt := 0
	patLen, srcLen := len(pat), len(src)
	for idx := 0; idx <= srcLen-patLen; idx++ {
		if mtcPat(src[idx:idx+patLen], pat) {
			cnt++
		}
	}
	return cnt
}

func TestSolutionGCA2b(t *testing.T) {
	cases := []struct {
		name     string
		pat      string
		src      string
		expected int
	}{
		{
			name:     "Example 1: amazing with 010",
			pat:      "010",
			src:      "amazing",
			expected: 2,
		},
		{
			name:     "Example 2: codesignal with 100",
			pat:      "100",
			src:      "codesignal",
			expected: 0,
		},
		{
			name:     "Empty source",
			pat:      "01",
			src:      "",
			expected: 0,
		},
		{
			name:     "Pattern longer than source",
			pat:      "0101",
			src:      "abc",
			expected: 0,
		},
		{
			name:     "Single character pattern",
			pat:      "0",
			src:      "aei",
			expected: 3,
		},
		{
			name:     "All vowels source",
			pat:      "00",
			src:      "aei",
			expected: 2,
		},
		{
			name:     "All consonants source",
			pat:      "11",
			src:      "bcd",
			expected: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := solutionGCA2b(tc.pat, tc.src)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
