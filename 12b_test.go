package main

import (
	"strings"
	"testing"
)

func intToRoman12b(num int) string {
	// Integer to Roman
	// Given an integer num.
	// Transform the integer to a roman numeral string.
	// Return the roman numeral string.
	// Use local optimization "greedy" approach.
	// Use a sort-descending integer-symbol table.
	// Table includes numeral combinations.
	// Roman numerals: I V X L C D M

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
	for _, vp := range vps {
		for num >= vp.val {
			res.WriteString(vp.sym)
			num -= vp.val
		}
	}
	return res.String()
}

func TestIntToRoman12b(t *testing.T) {
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
		res := intToRoman12b(tst.inp)
		if res != tst.exp {
			t.Errorf("intToRoman(%d) = %s; want %s", tst.inp, res, tst.exp)
		}
	}
}

func TestIntToRomanEdgeCases12b(t *testing.T) {
	// Test lower bound
	if res := intToRoman12b(1); res != "I" {
		t.Errorf("intToRoman(1) = %s; want I", res)
	}

	// Test upper bound
	if res := intToRoman12(3999); res != "MMMCMXCIX" {
		t.Errorf("intToRoman(3999) = %s; want MMMCMXCIX", res)
	}
}
