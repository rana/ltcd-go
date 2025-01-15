package main

import (
	"reflect"
	"testing"
)

// Time complexity:
// Space complexity:
func removeDuplicates80b(nums []int) int {
	return 0
}

func TestRemoveDuplicates80b(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		wantLen int
		wantRes []int
	}{
		{
			name:    "Example 1",
			nums:    []int{1, 1, 1, 2, 2, 3},
			wantLen: 5,
			wantRes: []int{1, 1, 2, 2, 3},
		},
		{
			name:    "Example 2",
			nums:    []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			wantLen: 7,
			wantRes: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name:    "No duplicates",
			nums:    []int{1, 2, 3, 4, 5},
			wantLen: 5,
			wantRes: []int{1, 2, 3, 4, 5},
		},
		{
			name:    "All duplicates",
			nums:    []int{1, 1, 1, 1, 1},
			wantLen: 2,
			wantRes: []int{1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			gotLen := removeDuplicates80b(numsCopy)
			if gotLen != tt.wantLen {
				t.Errorf("got length %d, want %d", gotLen, tt.wantLen)
			}
			if !reflect.DeepEqual(numsCopy[:gotLen], tt.wantRes) {
				t.Errorf("got nums %v, want %v", numsCopy[:gotLen], tt.wantRes)
			}
		})
	}
}
