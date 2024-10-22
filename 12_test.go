package main

import (
	"strings"
	"testing"
)

// Time complexity: O(1), fixed-size number of operations. Maximum input integer is 3999.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/f1d0b601-0832-40eb-a6e2-d1bd6fc852e9
func intToRoman(num int) string {
	// Integer to Roman
	// Given an integer num.
	// Transform the integer to a string of roman numerals.
	// Use a local optimization "greedy" approach.
	// Use a value-symbol table sorted descending.
	// Roman numerals: I V X L C D M

	// Initialize variables.
	vps := []struct {
		val int
		sym string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	var res strings.Builder

	// Iterate through value-pair table.
	for _, vp := range vps {
		// Check whther input integer is greater or
		// equal to current table value.
		for num >= vp.val {
			// Store the roman numeral.
			res.WriteString(vp.sym)
			// Decrement the current table value.
			num -= vp.val
		}
	}

	return res.String()
}

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		inp int
		exp string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
		{3999, "MMMCMXCIX"}, // Maximum value
		{1, "I"},            // Minimum value
	}

	for _, tst := range tests {
		res := intToRoman(tst.inp)
		if res != tst.exp {
			t.Errorf("intToRoman(%d) = %s; want %s", tst.inp, res, tst.exp)
		}
	}
}

func TestIntToRomanEdgeCases(t *testing.T) {
	// Test lower bound
	if res := intToRoman(1); res != "I" {
		t.Errorf("intToRoman(1) = %s; want I", res)
	}

	// Test upper bound
	if res := intToRoman(3999); res != "MMMCMXCIX" {
		t.Errorf("intToRoman(3999) = %s; want MMMCMXCIX", res)
	}
}
