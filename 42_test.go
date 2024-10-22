package main

import "testing"

// Time complexity: O(n), n is the length of the array hgts. We traverse the aray once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/869a17ef-28d2-49a8-bf9b-640c2a30c95b
func trap(hgts []int) int {
	// Trapping Rain Water
	// Given an integer array hgts.
	// Array represent elevations.
	// Each integer represents a bar with width 1.
	// Compute total collected water ni the elevation map.
	// Use a two-pointer technique.
	// Key insight is to use the minimum of the
	//	maximum left and right elevation,
	// minus the current elevation.

	// Initialize two indexes.
	lft, rht := 0, len(hgts)-1

	// Initialize max left and right elevations.
	maxLft, maxRht := 0, 0

	// Initialize total water collected.
	totWtr := 0

	// Loop until and right meet.
	for lft < rht {

		// Check for smaller side of elevation.
		if hgts[lft] < hgts[rht] {
			// Left-side elevation is smaller.

			// Check for the maximum left elevation.
			if hgts[lft] > maxLft {
				// Set maximum left elevation.
				maxLft = hgts[lft]
			} else {
				// Accumulate left water.
				totWtr += maxLft - hgts[lft]
			}

			// Advance left index.
			lft++
		} else {
			// Right-side elevation is smaller.

			// Check for maximum left elevation.
			if hgts[rht] > maxRht {
				// Set a new maximum right elevation.
				maxRht = hgts[rht]
			} else {
				// Accumulate right-side water.
				totWtr += maxRht - hgts[rht]
			}

			// Advance right index.
			rht--
		}
	}

	// Return total water collected.
	return totWtr
}

// Unit tests
func TestTrapRainWater(t *testing.T) {
	// Test cases: each contains an input height array and expected output
	tcs := []struct {
		hgt []int // Input height array
		exp int   // Expected water trapped
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6}, // Example 1 from problem statement
		{[]int{4, 2, 0, 3, 2, 5}, 9},                   // Example 2 from problem statement
		{[]int{}, 0},                                   // Edge case: empty array
		{[]int{1}, 0},                                  // Edge case: single element
		{[]int{1, 1, 1, 1}, 0},                         // Case where no water can be trapped
		{[]int{3, 0, 2, 0, 4}, 7},                      // Additional test case
	}

	// Run each test case
	for _, tc := range tcs {
		res := trap(tc.hgt)
		// Check if the result matches the expected output
		if res != tc.exp {
			t.Errorf("For hgt %v, exp %d, but got %d", tc.hgt, tc.exp, res)
		}
	}
}
