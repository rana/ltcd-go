package main

import (
	"testing"
)

func reverseWords151b(s string) string {
	return ""
}

func TestReverseWords151b(t *testing.T) {
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
			got := reverseWords151b(tst.inp)
			if got != tst.exp {
				t.Errorf("reverseWords(%q) = %q; want %q", tst.inp, got, tst.exp)
			}
		})
	}
}
