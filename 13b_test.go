package main

import "testing"

func romanToInt13b(s string) int {
	return 0
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
