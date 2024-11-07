package main

import (
	"strconv"
	"strings"
	"testing"
)

// Time complexity: O(n), n is the length of nums. We traverse the array once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/af8f11c1-69bb-45b3-99d4-56e6457b4499
func summaryRanges228(nums []int) []string {
	// Summary Ranges
	// Given an integer array nums.
	// Determine consecutive sequences.
	// Return an array of summary strings.
	// Conditions:
	//	* All nums are unique.
	//	* Array is sort-ascending.
	//	* Single integer is considered a sequence.
	//	* Range is inclusive of first element and last element.
	// Use a sliding window two-pointer approach.
	// Store sequence start, and search for sequence end.
	// Use a helper function to create a summary string.

	// Initialize variables.
	var res []string
	numLen := len(nums)

	// Check input minimum edge case.
	if numLen == 0 {
		return res
	}

	// Iterate through each element of the nums array.
	// Use a two-pointer approach.
	lft := 0
	for rht := 0; rht < numLen; rht++ {
		// Check for the array end, and sequence end.
		if rht+1 == numLen || nums[rht+1] != nums[rht]+1 {
			// Sequence end.
			// Store range.
			res = append(res, formatRange228(nums[lft], nums[rht]))
			// Start a new sequence.
			lft = rht + 1
		}
	}

	return res
}

func formatRange228(fst, lst int) string {
	if fst == lst {
		return strconv.Itoa(fst)
	}

	// Use a string builder.
	var bld strings.Builder
	bld.WriteString(strconv.Itoa(fst))
	bld.WriteString("->")
	bld.WriteString(strconv.Itoa(lst))

	return bld.String()
}

func TestSummaryRanges228(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []string
	}{
		{
			name: "Example 1",
			nums: []int{0, 1, 2, 4, 5, 7},
			want: []string{"0->2", "4->5", "7"},
		},
		{
			name: "Example 2",
			nums: []int{0, 2, 3, 4, 6, 8, 9},
			want: []string{"0", "2->4", "6", "8->9"},
		},
		{
			name: "Empty array",
			nums: []int{},
			want: []string{},
		},
		{
			name: "Single element",
			nums: []int{1},
			want: []string{"1"},
		},
		{
			name: "All continuous",
			nums: []int{1, 2, 3, 4, 5},
			want: []string{"1->5"},
		},
		{
			name: "All separate",
			nums: []int{1, 3, 5, 7, 9},
			want: []string{"1", "3", "5", "7", "9"},
		},
		{
			name: "Negative numbers",
			nums: []int{-3, -2, -1, 1, 2},
			want: []string{"-3->-1", "1->2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := summaryRanges228(tt.nums)
			if len(got) != len(tt.want) {
				t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("summaryRanges() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
