package main

import (
	"strings"
	"testing"
)

// Time complexity: O(w × n × m). w is chars in words. n is length of s. max substring of longest word.
// Space complexity: O(n), for length of bool array.
// https://console.anthropic.com/workbench/6900248e-ef13-4115-b90f-0b84e75b9e13
func addBoldTag616(s string, wrds []string) string {
	// Add Bold Tag in String
	// Given a string s.
	// Given an array of strings wrds.
	// Add closed pair of bold tags <b></b>
	// to string around words in s.
	// Return processed string.
	// Conditions:
	// * Overlapping substrings use a single bold tag
	// * Two consecutive strings use a single bold tag
	// Use a []bool
	// Use two steps:
	// 1. Mark characters which are bold
	// 2. Use a strings.Builder to write processed string.

	// Step 1: Mark which characters should be bold
	sLen := len(s)
	bold := make([]bool, sLen)
	for i := range sLen {
		for _, wrd := range wrds {
			wrdLen := len(wrd)
			if i+wrdLen <= sLen && s[i:i+wrdLen] == wrd {
				for j := range wrdLen {
					bold[i+j] = true
				}
			}
		}
	}

	// Step 2: Build the result string with bold tags
	var res strings.Builder
	i := 0
	for i < sLen {
		if !bold[i] {
			res.WriteByte(s[i])
			i++
		} else {
			res.WriteString("<b>")
			for i < sLen && bold[i] {
				res.WriteByte(s[i])
				i++
			}
			res.WriteString("</b>")
		}
	}

	return res.String()
}

func TestAddBoldTag616(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		words    []string
		expected string
	}{
		{
			name:     "Example 1",
			s:        "abcxyz123",
			words:    []string{"abc", "123"},
			expected: "<b>abc</b>xyz<b>123</b>",
		},
		{
			name:     "Example 2",
			s:        "aaabbb",
			words:    []string{"aa", "b"},
			expected: "<b>aaabbb</b>",
		},
		{
			name:     "Empty string",
			s:        "",
			words:    []string{"abc"},
			expected: "",
		},
		{
			name:     "No words",
			s:        "hello world",
			words:    []string{},
			expected: "hello world",
		},
		{
			name:     "Non-overlapping words",
			s:        "abc def ghi",
			words:    []string{"abc", "ghi"},
			expected: "<b>abc</b> def <b>ghi</b>",
		},
		{
			name:     "Word not in string",
			s:        "hello",
			words:    []string{"world"},
			expected: "hello",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := addBoldTag616(tc.s, tc.words)
			if got != tc.expected {
				t.Errorf("addBoldTag(%q, %v) = %q, want %q",
					tc.s, tc.words, got, tc.expected)
			}
		})
	}
}
