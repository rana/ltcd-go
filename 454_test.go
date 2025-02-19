package main

import "testing"

// Time Complexity: O(n²), where n is the length of the input arrays
// Space Complexity: O(n²), The hash map stores at most n² different sums from the first two arrays
// https://console.anthropic.com/workbench/4afccec2-42d5-4865-91a1-8b71137a30c0
func fourSumCount454(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	frqSum := make(map[int]int)
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			frqSum[n1+n2]++
		}
	}
	cnt := 0
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			cnt += frqSum[-(n3 + n4)]
		}
	}
	return cnt
}

func TestFourSumCount454(t *testing.T) {
	cases := []struct {
		name     string
		nums1    []int
		nums2    []int
		nums3    []int
		nums4    []int
		expected int
	}{
		{
			name:     "Example 1",
			nums1:    []int{1, 2},
			nums2:    []int{-2, -1},
			nums3:    []int{-1, 2},
			nums4:    []int{0, 2},
			expected: 2,
		},
		{
			name:     "Example 2",
			nums1:    []int{0},
			nums2:    []int{0},
			nums3:    []int{0},
			nums4:    []int{0},
			expected: 1,
		},
		{
			name:     "No valid combinations",
			nums1:    []int{1},
			nums2:    []int{1},
			nums3:    []int{1},
			nums4:    []int{1},
			expected: 0,
		},
		{
			name:     "Multiple same combinations",
			nums1:    []int{1, 1},
			nums2:    []int{-1, -1},
			nums3:    []int{0, 0},
			nums4:    []int{0, 0},
			expected: 16,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := fourSumCount454(tc.nums1, tc.nums2, tc.nums3, tc.nums4)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
