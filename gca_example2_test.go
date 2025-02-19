package main

import "testing"

// Module 2 – Data Manipulation
// You are given two strings: pattern and source. The first string pattern contains only the symbols 0 and 1, and
// the second string source contains only lowercase English letters. Your task is to calculate the number of
// substrings of source that match pattern.
// We’ll say that a substring source[l..r] matches pattern if the following three conditions are met:
// • The pattern and substring are equal in length.
// • Where there is a 0 in the pattern, there is a vowel in the substring.
// • Where there is a 1 in the pattern, there is a consonant in the substring.
// Vowels are ["a", "e", "i", "o", "u", "y"]. All other letters are consonants.

// Example
// For pattern = "010" and source = "amazing", the output should be solution(pattern, source) = 2.
// • "010" matches source[0..2] = "ama". The pattern specifies "vowel, consonant, vowel". "ama"
// matches this pattern: 0 matches a, 1 matches m, and 0 matches a.
// • "010" doesn’t match source[1..3] = "maz"
// 5"010" matches source[2..4] = "azi"
// "010" doesn’t match source[3..5] = "zin"
// "010" doesn’t match source[4..6] = "ing"
// So, there are 2 matches.

// For pattern = "100" and source = "codesignal", the output should be solution(pattern, source) = 0.
// • There are no double vowels in the string "codesignal", so it’s not possible for any of its substrings to
// match this pattern.

// Guaranteed constraints:
// • 1 ≤ source.length ≤ 103
// • 1 ≤ pattern.length ≤ 103

// Time Complexity: O(n * m) where n is length of source string, m is length of pattern
// Space Complexity: O(1)
// https://console.anthropic.com/workbench/f83050e1-437e-46a2-9efb-1c0e1cf5a505
func solutionGCA2(pat string, src string) int {
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
			if (pat[idx] == '0' && !isVow) || (pat[idx] == '1' && isVow) {
				return false
			}
		}
		return true
	}

	cnt := 0
	patLen := len(pat)
	for idx := 0; idx <= len(src)-patLen; idx++ {
		if mtcPat(src[idx:idx+patLen], pat) {
			cnt++
		}
	}
	return cnt
}

func TestSolutionGCA2(t *testing.T) {
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
			got := solutionGCA2(tc.pat, tc.src)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
