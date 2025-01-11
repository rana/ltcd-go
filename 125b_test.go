package main

import (
	"testing"
	"unicode"
)

func isValidPalindrome125b(s string) bool {
	// Valid Palindrome
	// Given a string s.
	// Determine if string is palindrome.
	// Return true if condition met.
	// Conditions:
	// * Remove non-alphanumeric.
	// * Transform to lowercase for comparison.
	// Use a helper function isAlphanumeric.
	// Use two-pointer technique.
	lft, rht := 0, len(s)-1
	for lft < rht {
		for lft < rht && !isAlphanumeric125b(s[lft]) {
			lft++
		}
		for lft < rht && !isAlphanumeric125b(s[rht]) {
			rht--
		}
		if unicode.ToLower(rune(s[lft])) != unicode.ToLower(rune(s[rht])) {
			return false
		}
		lft++
		rht--
	}
	return true
}

func isAlphanumeric125b(chr byte) bool {
	return unicode.IsLetter(rune(chr)) || unicode.IsDigit(rune(chr))
}

func TestIsValidPalindrome125b(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "Valid palindrome with spaces and punctuation",
			str:  "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "Invalid palindrome",
			str:  "race a car",
			want: false,
		},
		{
			name: "Empty string with space",
			str:  " ",
			want: true,
		},
		{
			name: "Single character",
			str:  "a",
			want: true,
		},
		{
			name: "Mixed case palindrome",
			str:  "Ab1bA",
			want: true,
		},
		{
			name: "Numbers only palindrome",
			str:  "12321",
			want: true,
		},
		{
			name: "Special characters only",
			str:  ".,",
			want: true,
		},
		{
			name: "Complex palindrome with numbers",
			str:  "A1b,2. 2b1a",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidPalindrome125b(tt.str); got != tt.want {
				t.Errorf("isValidPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
