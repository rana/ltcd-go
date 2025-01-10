package main

import "testing"

func romanToInt13b(s string) int {
	// Roman to Integer
	// Given a string of roman numerals.
	// Transform the roman numerals to an integer.
	// Return the integer.
	// Conditions:
	// * Apply roman numeral rules
	// Use a map of byte-integer
	// Traverse string right to left.
	// Compare to previous value for add or subtract rule.
	// Roman numerals: I V X L C D M

	rmn := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	res, prvVal := 0, 0

	// Iterate through characters right to left.
	for idx := len(s) - 1; idx >= 0; idx-- {
		curVal := rmn[s[idx]]
		if curVal >= prvVal {
			res += curVal
		} else {
			res -= curVal
		}
		prvVal = curVal
	}
	return res
}

func TestRomanToInt13b(t *testing.T) {
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
			got := romanToInt13b(tt.input)
			if got != tt.expected {
				t.Errorf("romanToInt(%s) = %d; want %d", tt.input, got, tt.expected)
			}
		})
	}
}
