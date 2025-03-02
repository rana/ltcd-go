package main

import (
	"sort"
	"testing"
)

func distributeEvenly1b(amt int, owed map[string]int) map[string]int {
	// Distribute Money Evenly
	// Given an integer amt.
	// Given a string-int map owed.
	// Determine an even distribution of amt to owed.
	// Conditions:
	// * Order the recipients
	// * Do distribution in rounds.
	// * If owed amount is reached, add no more.
	// Create an array of recipients to maintain distribution order.
	// Sort the array of recipients.

	dst := make(map[string]int)
	var rcps []string
	for rcp := range owed {
		rcps = append(rcps, rcp)
	}
	sort.Strings(rcps)

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
		if !dstAny {
			break
		}
	}

	return dst
}

func TestDistributeEvenly1b(t *testing.T) {
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
			got := distributeEvenly1b(tc.amt, tc.owed)
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
