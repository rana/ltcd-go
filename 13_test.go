package main

import "testing"

// Time complexity: O(n), n is the length of the string s.
// Space complexity: O(1), constant additional space used. Fixed-size map used for roman numeral to integer mapping.
// https://claude.ai/chat/7b14892e-57ab-4b97-9c1b-86aa99415c56
func romanToInt(s string) int {
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

// Test cases for valid inputs
func TestRomanToInt_III(t *testing.T) {
	if got := romanToInt("III"); got != 3 {
		t.Errorf("romanToInt(III) = %d; want 3", got)
	}
}

func TestRomanToInt_LVIII(t *testing.T) {
	if got := romanToInt("LVIII"); got != 58 {
		t.Errorf("romanToInt(LVIII) = %d; want 58", got)
	}
}

func TestRomanToInt_MCMXCIV(t *testing.T) {
	if got := romanToInt("MCMXCIV"); got != 1994 {
		t.Errorf("romanToInt(MCMXCIV) = %d; want 1994", got)
	}
}

// Test cases for subtraction rules
func TestRomanToInt_IV(t *testing.T) {
	if got := romanToInt("IV"); got != 4 {
		t.Errorf("romanToInt(IV) = %d; want 4", got)
	}
}

func TestRomanToInt_IX(t *testing.T) {
	if got := romanToInt("IX"); got != 9 {
		t.Errorf("romanToInt(IX) = %d; want 9", got)
	}
}

func TestRomanToInt_XL(t *testing.T) {
	if got := romanToInt("XL"); got != 40 {
		t.Errorf("romanToInt(XL) = %d; want 40", got)
	}
}

func TestRomanToInt_XC(t *testing.T) {
	if got := romanToInt("XC"); got != 90 {
		t.Errorf("romanToInt(XC) = %d; want 90", got)
	}
}

func TestRomanToInt_CD(t *testing.T) {
	if got := romanToInt("CD"); got != 400 {
		t.Errorf("romanToInt(CD) = %d; want 400", got)
	}
}

func TestRomanToInt_CM(t *testing.T) {
	if got := romanToInt("CM"); got != 900 {
		t.Errorf("romanToInt(CM) = %d; want 900", got)
	}
}

// Test case for maximum valid input
func TestRomanToInt_MMMCMXCIX(t *testing.T) {
	if got := romanToInt("MMMCMXCIX"); got != 3999 {
		t.Errorf("romanToInt(MMMCMXCIX) = %d; want 3999", got)
	}
}

// Test cases for invalid inputs
func TestRomanToInt_EmptyString(t *testing.T) {
	if got := romanToInt(""); got != 0 {
		t.Errorf("romanToInt(\"\") = %d; want 0", got)
	}
}

func TestRomanToInt_InvalidOrder(t *testing.T) {
	if got := romanToInt("IIII"); got != 4 {
		t.Errorf("romanToInt(IIII) = %d; want 4", got)
	}
}
