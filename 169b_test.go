package main

import "testing"

// Time complexity: O(n), n is the length of the array.
// Space complexity: O(1), constant additional space used.
func majorityElement169b(nums []int) int {
	// Majority Element
	// Given and array nums.
	// Determine the majority element.
	// Return the majority element.
	// Majority element appears more than len(nums)/2.
	// Use Boyer-Moore Voting algorithm.

	// Initialize variables.
	can, cnt := nums[0], 0

	// Iterate through nums.
	for _, num := range nums {
		// Check whether to set candidate.
		if cnt == 0 {
			can = num
		}

		// Check count adjustment.
		if can == num {
			cnt++
		} else {
			cnt--
		}
	}

	// Candidate is now majority element.
	return can
}

func TestMajorityElement169b(t *testing.T) {
	// Define test cases
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "three numbers with majority",
			nums: []int{3, 2, 3},
			want: 3,
		},
		{
			name: "seven numbers with majority",
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		{
			name: "single number",
			nums: []int{1},
			want: 1,
		},
		{
			name: "three numbers with majority at end",
			nums: []int{6, 5, 5},
			want: 5,
		},
		{
			name: "five numbers with majority at start",
			nums: []int{1, 1, 1, 2, 2},
			want: 1,
		},
	}

	// Run all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := majorityElement169b(tt.nums)
			if got != tt.want {
				t.Errorf("majorityElement169() = %d, want %d", got, tt.want)
			}
		})
	}
}
