package main

import (
	"sort"
	"testing"
)

// Time complexity: O(n * m)
// distributeEvenly1 distributes the given amount
// among recipients as evenly as possible
// while respecting the maximum owed amounts
func distributeEvenly1(amt int, owed map[string]int) map[string]int {
	// Distribute Money
	// Given integer amt.
	// Given string-int map owed.
	// Evenly as possible distribute money to receipients.
	// Return map string-int of distribution.
	// Conditions:
	// * Order recipients.
	// * Do distribution in round robin
	// * Skip recipient if owed amount already reached.

	// Initialize distribution map with all recipients starting at 0
	dst := make(map[string]int)

	// Create an ordered list of recipients
	// Sort the recipients to ensure deterministic order
	var rcps []string
	for rcp := range owed {
		rcps = append(rcps, rcp)
	}
	sort.Strings(rcps)

	// Distribute money in rounds
	for amt > 0 {
		dstAny := false
		for _, rcp := range rcps {
			if dst[rcp] < owed[rcp] {
				dst[rcp]++
				amt--
				dstAny = true
				if amt == 0 {
					break
				}
			}
		}
		// If we couldn't distribute any money in this round, break the loop
		if !dstAny {
			break
		}
	}

	return dst
}

func TestDistributeEvenly1(t *testing.T) {
	cases := []struct {
		name     string
		amt      int
		owed     map[string]int
		expected map[string]int
	}{
		{
			name: "Example case",
			amt:  45,
			owed: map[string]int{
				"Recipient A": 30,
				"Recipient B": 10,
				"Recipient C": 10,
				"Recipient D": 14,
			},
			expected: map[string]int{
				"Recipient A": 13,
				"Recipient B": 10,
				"Recipient C": 10,
				"Recipient D": 12,
			},
		},
		{
			name: "Exact amount case",
			amt:  20,
			owed: map[string]int{
				"A": 5,
				"B": 5,
				"C": 10,
			},
			expected: map[string]int{
				"A": 5,
				"B": 5,
				"C": 10,
			},
		},
		{
			name: "Not enough money",
			amt:  10,
			owed: map[string]int{
				"X": 10,
				"Y": 10,
				"Z": 10,
			},
			expected: map[string]int{
				"X": 4,
				"Y": 3,
				"Z": 3,
			},
		},
		{
			name: "Too much money",
			amt:  100,
			owed: map[string]int{
				"P1": 10,
				"P2": 20,
			},
			expected: map[string]int{
				"P1": 10,
				"P2": 20,
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := distributeEvenly1(tc.amt, tc.owed)
			// Check if the maps are equal
			if len(got) != len(tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
				return
			}
			for rcp, amt := range tc.expected {
				if got[rcp] != amt {
					t.Errorf("got %v, want %v", got, tc.expected)
					return
				}
			}
		})
	}
}
