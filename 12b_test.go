package main

import (
	"testing"
)

func intToRoman12b(num int) string {
	return ""
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
