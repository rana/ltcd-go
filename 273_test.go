package main

import (
	"strings"
	"testing"
)

// Time complexity: O(1), for constrinat max constant num; O(log n) for unconstrained.
// Space complexity: O(1), for constrinat max constant num; O(log n) for unconstrained.
// https://console.anthropic.com/workbench/3c326162-62ea-4687-b2d5-e61af4c46427
func numberToWords273(num int) string {
	if num == 0 {
		return "Zero"
	}

	ones := []string{"", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten",
		"Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	grp := []string{"", "Thousand", "Million", "Billion"}

	// Convert 0-999
	toWords := func(n int) string {
		if n == 0 {
			return ""
		}

		var sb strings.Builder

		// Hundreds
		if n >= 100 {
			sb.WriteString(ones[n/100] + " Hundred")
			if n%100 != 0 {
				sb.WriteString(" ")
			}
		}

		// Tens and ones
		n %= 100
		if n > 0 {
			if n < 20 {
				sb.WriteString(ones[n])
			} else {
				sb.WriteString(tens[n/10])
				if n%10 != 0 {
					sb.WriteString(" " + ones[n%10])
				}
			}
		}

		return sb.String()
	}

	var res strings.Builder

	// Process from billions to ones
	for i := 3; i >= 0; i-- {
		div := 1
		for range i {
			div *= 1000
		}

		cur := (num / div) % 1000

		if cur > 0 {
			if res.Len() > 0 {
				res.WriteString(" ")
			}
			res.WriteString(toWords(cur))
			if i > 0 {
				res.WriteString(" " + grp[i])
			}
		}
	}

	return res.String()
}

func TestNumberToWords273(t *testing.T) {
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
			got := numberToWords273(tc.num)
			if got != tc.expected {
				t.Errorf("numberToWords(%d) = %q, want %q", tc.num, got, tc.expected)
			}
		})
	}
}
