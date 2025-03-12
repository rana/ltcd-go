package main

import (
	"testing"
	"unicode"
)

// Time complexity: O(n), n is the length of the string. We traverse the string once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/6f5b68df-2313-4b4b-a30e-0cb09dab3a7e
func isValidPalindrome125(s string) bool {
	// Valid Palindrome
	// Given a string s with printable ASCII characters.
	// Determine whether the string is a palindrome.
	// Palindrome conditions:
	//	* Remove non-alphanumeric characters.
	//	* Transform characters to lowercase.
	// Use a one-pass approach which cleans while checks.
	// Use a two-pointer technique.

	isLtrDig := func(chr rune) bool {
		return unicode.IsLetter(chr) || unicode.IsDigit(chr)
	}

	// Initialize two pointers.
	lft, rht := 0, len(s)-1

	// Iterate until left and right meet.
	for lft < rht {
		// Skip left non-alphanumeric characters.
		for lft < rht && !isLtrDig(rune(s[lft])) {
			lft++
		}

		// Skip right non-alphanumeric characters.
		for lft < rht && !isLtrDig(rune(s[rht])) {
			rht--
		}

		// Compare case-insensitive characters.
		if unicode.ToLower(rune(s[lft])) != unicode.ToLower(rune(s[rht])) {
			return false
		}

		// Move pointers.
		lft++
		rht--
	}

	return true
}

// Unit Tests
func TestIsValidPalindrome125(t *testing.T) {
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
			if got := isValidPalindrome125(tt.str); got != tt.want {
				t.Errorf("isValidPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
