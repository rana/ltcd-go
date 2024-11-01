package main

import "testing"

// Time complexity: O(n), n is the length of array rts. We traverse the array twice.
// Space complexity: O(n), n additional space used for candy accumulation.
// https://claude.ai/chat/369ef6a2-2505-4c98-909a-2b7a41283e3f
func candy(rts []int) int {
	// Candy
	// Given an integer array rts.
	// Determine the minimum sum of candy.
	// Return the minimum sum of candy.
	// Conditions:
	//	* Each element receives at least one candy.
	//	* Each element rating larger than an adjacent,
	//	receives one more candy.
	// Use a two-pass approach.
	// Pass one, compare left-adjacent ratings.
	// Pass two, compare right-adjacent ratings.
	// Use a local optimization "greedy" approach.

	// Initialize variables.
	len := len(rts)
	cnd := make([]int, len)
	for n := range cnd {
		cnd[n] = 1
	}

	// Pass one: lft to rht: compare left-adjacent ratings.
	for n := 1; n < len; n++ {
		if rts[n] > rts[n-1] {
			cnd[n] = cnd[n-1] + 1
		}
	}

	// Pass two: rht to lft: compare right-adjacent ratings.
	for n := len - 2; n >= 0; n-- {
		if rts[n] > rts[n+1] {
			cnd[n] = max_135(cnd[n], cnd[n+1]+1)
		}
	}

	// Sum candies.
	sum := 0
	for _, c := range cnd {
		sum += c
	}

	return sum
}

func max_135(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestCandy(t *testing.T) {
	tsts := []struct {
		nme string
		rts []int
		wnt int
	}{
		{"Example 1", []int{1, 0, 2}, 5},
		{"Example 2", []int{1, 2, 2}, 4},
		{"Single child", []int{1}, 1},
		{"All equal", []int{1, 1, 1, 1}, 4},
		{"Increasing", []int{1, 2, 3, 4, 5}, 15},
		{"Decreasing", []int{5, 4, 3, 2, 1}, 15},
		{"Valley", []int{3, 2, 1, 2, 3}, 11}, // Corrected expected output
		{"Peak", []int{1, 3, 5, 3, 1}, 9},
	}

	for _, tt := range tsts {
		t.Run(tt.nme, func(t *testing.T) {
			if got := candy(tt.rts); got != tt.wnt {
				t.Errorf("candy() = %v, want %v", got, tt.wnt)
			}
		})
	}
}
