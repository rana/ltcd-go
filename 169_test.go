package main

import "testing"

// Time complexity: O(n), n is the length of array nums. We traverse the array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/67141a8d-0ae4-8002-a543-87e59e0782de
func majorityElement169(nums []int) int {
	// Majority Element
	// Given and array nums.
	// Determine the majority element.
	// Return the majority element.
	// Use the Boyer-Moore Voting algorithm.

	// Initialize a candidate and count.
	can, cnt := 0, 0

	// Iterate through each element of nums.
	for _, num := range nums {
		// if cnt is zero, set candidate to current num.
		if cnt == 0 {
			can = num
		}

		// Increment or decrement count based on candidate match.
		if can == num {
			cnt++
		} else {
			cnt--
		}
	}

	// Candidate is now majority element.
	// Return majority element.
	return can
}

func TestMajorityElement169(t *testing.T) {
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
			got := majorityElement169(tt.nums)
			if got != tt.want {
				t.Errorf("majorityElement169() = %d, want %d", got, tt.want)
			}
		})
	}
}
