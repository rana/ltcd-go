package main

import (
	"strings"
	"testing"
)

// Time complexity: O(n), n is the string array. We traverse the array three times. Once to split by whitespace. And a second time to reverse. A third time to join a string.
// Space complexity: O(n), n is the lnegth of a split array holding intermediate words.
// https://claude.ai/chat/c15cc2ad-5d6a-40d4-91ae-8b6c3810cb01
func reverseWords(str string) string {
	// Reverse Words in a String
	// Given a string str.
	// String holds words and whitespaces.
	// Reverse the words.
	// Return a string of reversed word delimited by spaces.
	// Use a two-pointer technique.

	// Trim leading and trailing spaces.
	// Split words by whitespace.
	wrds := strings.Fields(str)

	// Initialize two-pointer variables.
	lft, rht := 0, len(wrds)-1

	// Iterate until lft and rht meet.
	for lft < rht {
		// Reverse words.
		wrds[lft], wrds[rht] = wrds[rht], wrds[lft]

		// Advance indexes.
		lft++
		rht--
	}

	// Return reversed word strign joined by space.
	return strings.Join(wrds, " ")
}

// Test cases
func TestReverseWords(t *testing.T) {
	tests := []struct {
		nam string // test name
		inp string // input string
		exp string // expected result
	}{
		{
			nam: "Example 1: Basic case",
			inp: "the sky is blue",
			exp: "blue is sky the",
		},
		{
			nam: "Example 2: Leading and trailing spaces",
			inp: "  hello world  ",
			exp: "world hello",
		},
		{
			nam: "Example 3: Multiple spaces between words",
			inp: "a good   example",
			exp: "example good a",
		},
		{
			nam: "Single word",
			inp: "hello",
			exp: "hello",
		},
		{
			nam: "Multiple spaces everywhere",
			inp: "   multiple   spaces   everywhere   ",
			exp: "everywhere spaces multiple",
		},
		{
			nam: "Special characters",
			inp: "123 abc!@# def456",
			exp: "def456 abc!@# 123",
		},
	}

	for _, tst := range tests {
		t.Run(tst.nam, func(t *testing.T) {
			got := reverseWords(tst.inp)
			if got != tst.exp {
				t.Errorf("reverseWords(%q) = %q; want %q", tst.inp, got, tst.exp)
			}
		})
	}
}
