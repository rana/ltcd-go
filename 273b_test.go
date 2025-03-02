package main

import (
	"testing"
)

// Time complexity: O(1), for constrinat max constant num; O(log n) for unconstrained.
// Space complexity: O(1), for constrinat max constant num; O(log n) for unconstrained.
// https://console.anthropic.com/workbench/3c326162-62ea-4687-b2d5-e61af4c46427
func numberToWords273b(num int) string {
	return ""
}

func TestNumberToWords273b(t *testing.T) {
	cases := []struct {
		name     string
		num      int
		expected string
	}{
		{
			name:     "Example 1: 123",
			num:      123,
			expected: "One Hundred Twenty Three",
		},
		{
			name:     "Example 2: 12345",
			num:      12345,
			expected: "Twelve Thousand Three Hundred Forty Five",
		},
		{
			name:     "Example 3: 1234567",
			num:      1234567,
			expected: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		},
		{
			name:     "Zero",
			num:      0,
			expected: "Zero",
		},
		{
			name:     "Single digit",
			num:      5,
			expected: "Five",
		},
		{
			name:     "Teens",
			num:      16,
			expected: "Sixteen",
		},
		{
			name:     "Round numbers",
			num:      1000000,
			expected: "One Million",
		},
		{
			name:     "Max int",
			num:      2147483647,
			expected: "Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Seven",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := numberToWords273b(tc.num)
			if got != tc.expected {
				t.Errorf("numberToWords(%d) = %q, want %q", tc.num, got, tc.expected)
			}
		})
	}
}
