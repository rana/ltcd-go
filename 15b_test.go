package main

import (
	"fmt"
	"sort"
	"testing"
)

// Time complexity:
// Space complexity:
func threeSum15b(nums []int) [][]int {
	return nil
}

func TestThreeSum15b(t *testing.T) {
	compareResults := func(got, want [][]int) bool {
		if len(got) != len(want) {
			return false
		}

		// Create maps for comparison
		gotMap := make(map[string]bool)
		wantMap := make(map[string]bool)

		for _, triplet := range got {
			sort.Ints(triplet)
			key := fmt.Sprintf("%d,%d,%d", triplet[0], triplet[1], triplet[2])
			gotMap[key] = true
		}

		for _, triplet := range want {
			sort.Ints(triplet)
			key := fmt.Sprintf("%d,%d,%d", triplet[0], triplet[1], triplet[2])
			wantMap[key] = true
		}

		for k := range gotMap {
			if !wantMap[k] {
				return false
			}
		}

		return true
	}
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "Example 3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Empty array",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "Array with duplicates",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum15(tt.nums)
			if !compareResults(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
