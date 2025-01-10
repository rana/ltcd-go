package main

import "testing"

// Time complexity: O(n), n is the length of the string s.
// Space complexity: O(1), constant additional space used. Fixed-size map used for roman numeral to integer mapping.
// https://claude.ai/chat/7b14892e-57ab-4b97-9c1b-86aa99415c56
func romanToInt13(s string) int {
	// Roman to Integer
	// Given a string s.
	// String contains roman numeral characters.
	// Transform string s to an integer.
	// Return the integer.
	// Key insight is to traverse from right to left
	// allowing comparison of the current character to
	// the previous character.
	// A linear scan and lookup.

	// Initialize variables.
	rmn := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	prvVal := 0

	// Iterate string from right to left.
	for n := len(s) - 1; n >= 0; n-- {
		// Get the current integer value.
		curVal := rmn[s[n]]

		// Check whether to add or subtract to result.
		if curVal >= prvVal {
			res += curVal
		} else {
			res -= curVal
		}

		// Update previous value for next iteration.
		prvVal = curVal
	}

	return res
}

func TestRomanToInt13(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"III", "III", 3},
		{"LVIII", "LVIII", 58},
		{"MCMXCIV", "MCMXCIV", 1994},
		{"IV", "IV", 4},
		{"IX", "IX", 9},
		{"XL", "XL", 40},
		{"XC", "XC", 90},
		{"CD", "CD", 400},
		{"CM", "CM", 900},
		{"MMMCMXCIX", "MMMCMXCIX", 3999},
		{"Empty string", "", 0},
		{"Invalid order", "IIII", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := romanToInt13(tt.input)
			if got != tt.expected {
				t.Errorf("romanToInt(%s) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}
